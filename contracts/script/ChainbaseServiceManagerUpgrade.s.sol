// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "forge-std/Script.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import "../src/ChainbaseServiceManager.sol";

contract ChainbaseServiceManagerUpgrade is Script {
    //forge script --chain holesky script/ChainbaseServiceManagerUpgrade.s.sol --rpc-url $HOLESKY_RPC_URL --broadcast --verify -vvvv
    //forge script script/ChainbaseServiceManagerUpgrade.s.sol --rpc-url http://localhost:8545 --broadcast -vvvv
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        address avsDirectory = vm.envAddress("AVS_DIRECTORY");

        string memory chainbaseDeployedContracts = readOutput(
            "chainbase_avs_deployment_output"
        );

        address proxyAdmin = stdJson.readAddress(
            chainbaseDeployedContracts,
            ".addresses.proxyAdmin"
        );

        address chainbaseServiceManagerProxy = stdJson.readAddress(
            chainbaseDeployedContracts,
            ".addresses.chainbaseServiceManagerProxy"
        );
        address registryCoordinatorProxy = stdJson.readAddress(
            chainbaseDeployedContracts,
            ".addresses.registryCoordinator"
        );
        address stakeRegistryProxy = stdJson.readAddress(
            chainbaseDeployedContracts,
            ".addresses.stakeRegistryProxy"
        );
        ChainbaseServiceManager newImplementation = new ChainbaseServiceManager(
            IAVSDirectory(avsDirectory),
            IRegistryCoordinator(registryCoordinatorProxy),
            IStakeRegistry(stakeRegistryProxy)
        );

        ProxyAdmin(proxyAdmin).upgrade(ITransparentUpgradeableProxy(chainbaseServiceManagerProxy), address(newImplementation));

        vm.stopBroadcast();
    }

    function readOutput(
        string memory outputFileName
    ) internal view returns (string memory) {
        string memory inputDir = string.concat(
            vm.projectRoot(),
            "/script/output/"
        );
        string memory chainDir = string.concat(vm.toString(block.chainid), "/");
        string memory file = string.concat(outputFileName, ".json");
        return vm.readFile(string.concat(inputDir, chainDir, file));
    }
}
