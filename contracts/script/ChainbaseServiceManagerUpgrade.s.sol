// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import "forge-std/Script.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import "../src/ChainbaseServiceManager.sol";
import {EigenHoleSkyDeployments} from "./EigenDeployments.s.sol";

contract ChainbaseServiceManagerUpgrade is Script {
    struct DeploymentAddresses {
        address proxyAdmin;
        address chainbaseServiceManagerProxy;
        address registryCoordinatorProxy;
        address stakeRegistryProxy;
    }
    //forge script --chain holesky script/ChainbaseServiceManagerUpgrade.s.sol --rpc-url $HOLESKY_RPC_URL --broadcast --verify -vvvv
    //forge script script/ChainbaseServiceManagerUpgrade.s.sol --rpc-url http://localhost:8545 --broadcast -vvvv

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        DeploymentAddresses memory addresses = getDeploymentAddresses();
        ChainbaseServiceManager newImplementation = createNewImplementation(addresses);

        ProxyAdmin(addresses.proxyAdmin).upgrade(
            ITransparentUpgradeableProxy(addresses.chainbaseServiceManagerProxy), address(newImplementation)
        );

        vm.stopBroadcast();
    }

    function getDeploymentAddresses() internal view returns (DeploymentAddresses memory) {
        string memory chainbaseDeployedContracts = readOutput("chainbase_avs_deployment_output");

        return DeploymentAddresses({
            proxyAdmin: stdJson.readAddress(chainbaseDeployedContracts, ".addresses.proxyAdmin"),
            chainbaseServiceManagerProxy: stdJson.readAddress(
                chainbaseDeployedContracts, ".addresses.chainbaseServiceManagerProxy"
            ),
            registryCoordinatorProxy: stdJson.readAddress(chainbaseDeployedContracts, ".addresses.registryCoordinator"),
            stakeRegistryProxy: stdJson.readAddress(chainbaseDeployedContracts, ".addresses.stakeRegistryProxy")
        });
    }

    function createNewImplementation(DeploymentAddresses memory addresses) internal returns (ChainbaseServiceManager) {
        return new ChainbaseServiceManager(
            IAVSDirectory(EigenHoleSkyDeployments.AVSDirectory),
            IRewardsCoordinator(EigenHoleSkyDeployments.RewardsCoordinator),
            IRegistryCoordinator(addresses.registryCoordinatorProxy),
            IStakeRegistry(addresses.stakeRegistryProxy),
            IAllocationManager(EigenHoleSkyDeployments.AllocationManager)
        );
    }

    function readOutput(string memory outputFileName) internal view returns (string memory) {
        string memory inputDir = string.concat(vm.projectRoot(), "/script/output/");
        string memory chainDir = string.concat(vm.toString(block.chainid), "/");
        string memory file = string.concat(outputFileName, ".json");
        return vm.readFile(string.concat(inputDir, chainDir, file));
    }
}
