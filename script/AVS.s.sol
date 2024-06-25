// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import {Script, console} from "forge-std/Script.sol";
import {AVS} from "../src/AVS.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IAVSDirectory} from "@eigenlayer/contracts/interfaces/IAVSDirectory.sol";

contract DeployAVS is Script {
    function run() external {
        // https://github.com/Layr-Labs/eigenlayer-contracts holesky
        address delegationManagerAddress = vm.envAddress("DELEGATION_MANAGER_ADDRESS");
        address avsDirectoryAddress = vm.envAddress("AVS_DIRECTORY_ADDRESS");
        string memory metadataURI = vm.envString("METADATA_URI");

        vm.startBroadcast();

        AVS avs = new AVS(IDelegationManager(delegationManagerAddress), IAVSDirectory(avsDirectoryAddress), metadataURI);

        vm.stopBroadcast();

        console.log("AVS deployed to:", address(avs));
    }
}
