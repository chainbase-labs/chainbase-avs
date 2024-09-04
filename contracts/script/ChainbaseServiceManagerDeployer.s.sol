// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "forge-std/Script.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@eigenlayer/test/mocks/EmptyContract.sol";
import "@eigenlayer-middleware/src/StakeRegistry.sol";
import "@eigenlayer-middleware/src/BLSApkRegistry.sol";
import "@eigenlayer-middleware/src/IndexRegistry.sol";
import "@eigenlayer-middleware/src/RegistryCoordinator.sol";
import "@eigenlayer-middleware/src/OperatorStateRetriever.sol";
import "@eigenlayer-middleware/lib/eigenlayer-contracts/src/contracts/permissions/PauserRegistry.sol";

import "../src/ChainbaseServiceManager.sol";
import {EigenHoleSkyDeployments} from "./EigenDeployments.s.sol";

contract ChainbaseServiceManagerDeployer is Script {
    address public multiSigManager;
    address public aggregator;
    address public generator;

    ProxyAdmin public proxyAdmin;
    PauserRegistry public pauserRegistry;
    OperatorStateRetriever public operatorStateRetriever;

    IRegistryCoordinator public registryCoordinatorProxy;
    IRegistryCoordinator public registryCoordinatorImplementation;

    IBLSApkRegistry public blsApkRegistryProxy;
    IBLSApkRegistry public blsApkRegistryImplementation;

    IIndexRegistry public indexRegistryProxy;
    IIndexRegistry public indexRegistryImplementation;

    IStakeRegistry public stakeRegistryProxy;
    IStakeRegistry public stakeRegistryImplementation;

    IChainbaseServiceManager public chainbaseServiceManagerProxy;
    IChainbaseServiceManager public chainbaseServiceManagerImplementation;

    //forge script --chain holesky script/ChainbaseServiceManagerDeployer.s.sol --rpc-url $HOLESKY_RPC_URL --broadcast --verify -vvvv
    //forge script script/ChainbaseServiceManagerDeployer.s.sol --rpc-url http://localhost:8545 --broadcast -vvvv
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        multiSigManager = vm.envAddress("SERVICE_MANAGER");
        aggregator = vm.envAddress("SERVICE_AGGREGATOR");
        generator = vm.envAddress("SERVICE_GENERATOR");

        address delegationManager = vm.envAddress("DELEGATION_MANAGER");
        address avsDirectory = vm.envAddress("AVS_DIRECTORY");
        _deployChainbaseServiceManagerContracts(IDelegationManager(delegationManager), IAVSDirectory(avsDirectory));

        vm.stopBroadcast();
    }

    function _deployChainbaseServiceManagerContracts(IDelegationManager delegationManager, IAVSDirectory avsDirectory)
        internal
    {
        IStrategy[1] memory deployedStrategyArray = [IStrategy(EigenHoleSkyDeployments.cbETHStrategy)];
        uint256 numStrategies = deployedStrategyArray.length;

        // deploy proxy admin for ability to upgrade proxy contracts
        proxyAdmin = new ProxyAdmin();

        // deploy pauser registry
        {
            address[] memory pausers = new address[](2);
            pausers[0] = msg.sender;
            pausers[1] = multiSigManager;
            address unpauser = multiSigManager;
            pauserRegistry = new PauserRegistry(pausers, unpauser);
        }

        // deploy operator state retriever
        operatorStateRetriever = new OperatorStateRetriever();

        EmptyContract emptyContract = new EmptyContract();
        /**
         * First, deploy upgradeable proxy contracts that **will point** to the implementations. Since the implementation contracts are
         * not yet deployed, we give these proxies an empty contract as the initial implementation, to act as if they have no code.
         */
        blsApkRegistryProxy =
            IBLSApkRegistry(address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), "")));

        indexRegistryProxy =
            IIndexRegistry(address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), "")));

        stakeRegistryProxy =
            IStakeRegistry(address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), "")));

        registryCoordinatorProxy = RegistryCoordinator(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
        );

        chainbaseServiceManagerProxy = ChainbaseServiceManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
        );

        /**
         * Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
         */
        {
            blsApkRegistryImplementation = new BLSApkRegistry(registryCoordinatorProxy);

            proxyAdmin.upgrade(
                ITransparentUpgradeableProxy(payable(address(blsApkRegistryProxy))),
                address(blsApkRegistryImplementation)
            );

            indexRegistryImplementation = new IndexRegistry(registryCoordinatorProxy);

            proxyAdmin.upgrade(
                ITransparentUpgradeableProxy(payable(address(indexRegistryProxy))), address(indexRegistryImplementation)
            );

            stakeRegistryImplementation = new StakeRegistry(registryCoordinatorProxy, delegationManager);

            proxyAdmin.upgrade(
                ITransparentUpgradeableProxy(payable(address(stakeRegistryProxy))), address(stakeRegistryImplementation)
            );
        }

        registryCoordinatorImplementation = new RegistryCoordinator(
            IServiceManager(address(chainbaseServiceManagerProxy)),
            IStakeRegistry(address(stakeRegistryProxy)),
            IBLSApkRegistry(address(blsApkRegistryProxy)),
            IIndexRegistry(address(indexRegistryProxy))
        );

        {
            uint256 numQuorums = 1;
            // for each quorum to setup, we need to define
            // QuorumOperatorSetParam, minimumStakeForQuorum, and strategyParams
            IRegistryCoordinator.OperatorSetParam[] memory quorumsOperatorSetParams =
                new IRegistryCoordinator.OperatorSetParam[](numQuorums);
            for (uint256 i = 0; i < numQuorums; i++) {
                // hard code these for now
                quorumsOperatorSetParams[i] = IRegistryCoordinator.OperatorSetParam({
                    maxOperatorCount: 10000,
                    kickBIPsOfOperatorStake: 15000,
                    kickBIPsOfTotalStake: 100
                });
            }
            // set to 0 for every quorum
            uint96[] memory quorumsMinimumStake = new uint96[](numQuorums);
            IStakeRegistry.StrategyParams[][] memory quorumsStrategyParams =
                new IStakeRegistry.StrategyParams[][](numQuorums);
            for (uint256 i = 0; i < numQuorums; i++) {
                quorumsStrategyParams[i] = new IStakeRegistry.StrategyParams[](numStrategies);
                for (uint256 j = 0; j < numStrategies; j++) {
                    quorumsStrategyParams[i][j] = IStakeRegistry.StrategyParams({
                        strategy: deployedStrategyArray[j],
                        // setting this to 1 ether since the divisor is also 1 ether
                        // therefore this allows an operator to register with even just 1 token
                        // see https://github.com/Layr-Labs/eigenlayer-middleware/blob/m2-mainnet/src/StakeRegistry.sol#L484
                        //    weight += uint96(sharesAmount * strategyAndMultiplier.multiplier / WEIGHTING_DIVISOR);
                        multiplier: 1 ether
                    });
                }
            }
            proxyAdmin.upgradeAndCall(
                ITransparentUpgradeableProxy(payable(address(registryCoordinatorProxy))),
                address(registryCoordinatorImplementation),
                abi.encodeWithSelector(
                    RegistryCoordinator.initialize.selector,
                    multiSigManager,
                    multiSigManager,
                    multiSigManager,
                    pauserRegistry,
                    0, // 0 initialPausedStatus means everything unpaused
                    quorumsOperatorSetParams,
                    quorumsMinimumStake,
                    quorumsStrategyParams
                )
            );
        }

        chainbaseServiceManagerImplementation =
            new ChainbaseServiceManager(avsDirectory, registryCoordinatorProxy, stakeRegistryProxy);

        proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(chainbaseServiceManagerProxy))),
            address(chainbaseServiceManagerImplementation),
            abi.encodeWithSelector(
                ChainbaseServiceManager.initialize.selector, pauserRegistry, multiSigManager, aggregator, generator
            )
        );

        // WRITE JSON DATA
        string memory parent_object = "parent object";
        string memory deployed_addresses = "addresses";

        vm.serializeAddress(deployed_addresses, "proxyAdmin", address(proxyAdmin));
        vm.serializeAddress(deployed_addresses, "pauserRegistry", address(pauserRegistry));
        vm.serializeAddress(deployed_addresses, "operatorStateRetriever", address(operatorStateRetriever));

        vm.serializeAddress(deployed_addresses, "blsApkRegistryProxy", address(blsApkRegistryProxy));
        vm.serializeAddress(deployed_addresses, "blsApkRegistryImplementation", address(blsApkRegistryImplementation));

        vm.serializeAddress(deployed_addresses, "indexRegistryProxy", address(indexRegistryProxy));
        vm.serializeAddress(deployed_addresses, "indexRegistryImplementation", address(indexRegistryImplementation));

        vm.serializeAddress(deployed_addresses, "stakeRegistryProxy", address(stakeRegistryProxy));
        vm.serializeAddress(deployed_addresses, "stakeRegistryImplementation", address(stakeRegistryImplementation));

        vm.serializeAddress(deployed_addresses, "registryCoordinator", address(registryCoordinatorProxy));
        vm.serializeAddress(
            deployed_addresses, "registryCoordinatorImplementation", address(registryCoordinatorImplementation)
        );

        vm.serializeAddress(deployed_addresses, "chainbaseServiceManagerProxy", address(chainbaseServiceManagerProxy));
        string memory deployed_addresses_output = vm.serializeAddress(
            deployed_addresses, "chainbaseServiceManagerImplementation", address(chainbaseServiceManagerImplementation)
        );

        // serialize all the data
        string memory finalJson = vm.serializeString(parent_object, deployed_addresses, deployed_addresses_output);
        writeOutput(finalJson, "chainbase_avs_deployment_output");
    }

    function writeOutput(string memory outputJson, string memory outputFileName) internal {
        string memory outputDir = string.concat(vm.projectRoot(), "/script/output/");
        string memory chainDir = string.concat(vm.toString(block.chainid), "/");
        string memory outputFilePath = string.concat(outputDir, chainDir, outputFileName, ".json");
        vm.writeJson(outputJson, outputFilePath);
    }
}
