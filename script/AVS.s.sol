// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import {Script, console} from "forge-std/Script.sol";
import {AVS} from "../src/AVS.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IAVSDirectory} from "@eigenlayer/contracts/interfaces/IAVSDirectory.sol";

import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract DeployAVS is Script {
    // TODO: right now just for test. when launched at mainnet, need fixed.
    address proxyAdmin = 0xdFbD62c5d8C5739852f67F2D7d2148FC5Bf2ce8E;

    address delegationManagerAddress = vm.envAddress("DELEGATION_MANAGER_ADDRESS");
    address avsDirectoryAddress = vm.envAddress("AVS_DIRECTORY_ADDRESS");
    string metadataURI = vm.envString("METADATA_URI");

    // set by deploy(...)
    address _impl;
    bytes _implConstructorArgs;
    address _proxy;
    bytes _proxyConstructorArgs;

    function run() external {
        vm.startBroadcast();

        AVS avs = deploy();

        vm.stopBroadcast();
        console.log("OmniAVS deployed at: ", address(avs));
        console.log("Implementation: ", _impl);
        console.log("Implementation Constructor Args: ");
        console.logBytes(_implConstructorArgs);
        console.log("Proxy: ", _proxy);
        console.log("Proxy Constructor Args: ");
        console.logBytes(_proxyConstructorArgs);

        console.log("AVS deployed to:", address(avs));
    }

    function deploy() public returns (AVS) {
        address impl =
            address(new AVS(IDelegationManager(delegationManagerAddress), IAVSDirectory(avsDirectoryAddress)));

        _impl = impl;
        _implConstructorArgs =
            abi.encode(IDelegationManager(delegationManagerAddress), IAVSDirectory(avsDirectoryAddress));

        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            impl, proxyAdmin, abi.encodeWithSelector(AVS.initialize.selector, metadataURI)
        );

        _proxy = address(proxy);
        _proxyConstructorArgs =
            abi.encode(impl, proxyAdmin, abi.encodeWithSelector(AVS.initialize.selector, metadataURI));

        return AVS(address(proxy));
    }
}
