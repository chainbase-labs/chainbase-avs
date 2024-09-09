// SPDX-License-Identifier: UNLICENSED
pragma solidity = 0.8.24;

import {IChainbaseServiceManager} from "./IChainbaseServiceManager.sol";

abstract contract ChainbaseServiceManagerStorage is IChainbaseServiceManager {
    //=========================================================================
    //                                CONSTANT
    //=========================================================================
    // The number of blocks from the task initialization within which the aggregator has to respond to
    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 100;
    uint256 internal constant _THRESHOLD_DENOMINATOR = 100;

    //=========================================================================
    //                                STORAGE
    //=========================================================================
    // The latest task index
    uint32 public latestTaskNum;

    // mapping of task indices to all tasks hashes
    // when a task is created, task hash is stored here,
    // and responses need to pass the actual task,
    // which is hashed onchain and checked against this mapping
    mapping(uint32 => bytes32) public allTaskHashes;

    // mapping of task indices to hash of abi.encode(taskResponse, taskResponseMetadata)
    mapping(uint32 => bytes32) public allTaskResponses;

    address public aggregator;
    address public generator;
}
