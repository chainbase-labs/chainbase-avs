// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

// solhint-disable no-console
// solhint-disable var-name-mixedcase
// solhint-disable max-states-count
// solhint-disable state-visibility

import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";

contract DeployProxyAdmin is Script {
    function run() external {
        uint256 deployerKey = vm.envUint("PROXY_ADMIN_DEPLOYER_KEY");
        address ownerAddress = vm.envAddress("PROXY_ADMIN_OWNER_ADDRESS");

        // require(block.chainid == 1, "Only mainnet deployment.");

        vm.startBroadcast(deployerKey);
        (ProxyAdmin proxyAdmin) = deploy(ownerAddress);
        vm.stopBroadcast();

        console.log("ProxyAdmin deployed at: ", address(proxyAdmin));
        console.log("ProxyAdmin owner: ", ownerAddress);
    }

    function deploy(address owner) public returns (ProxyAdmin proxyAdmin) {
        require(owner != address(0), "Owner not set");
        proxyAdmin = new ProxyAdmin();
        proxyAdmin.transferOwnership(owner);
        require(address(proxyAdmin) != address(0), "ProxyAdmin not deployed");
        require(proxyAdmin.owner() == owner, "ProxyAdmin owner not set");
    }
}
