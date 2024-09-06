// SPDX-License-Identifier: UNLICENSED
pragma solidity = 0.8.24;

import "@eigenlayer-middleware/src/interfaces/IBLSSignatureChecker.sol";

interface IChainbaseServiceManager {
    //=========================================================================
    //                                 EVENT
    //=========================================================================
    event NewTaskCreated(uint32 indexed taskIndex, Task task);

    event TaskResponded(TaskResponse taskResponse, TaskResponseMetadata taskResponseMetadata);

    event TaskCompleted(uint32 indexed taskIndex);

    //=========================================================================
    //                                STRUCTS
    //=========================================================================
    struct Task {
        string taskDetails;
        uint32 taskCreatedBlock;
        bytes quorumNumbers;
        uint32 quorumThresholdPercentage;
    }

    // Task response is hashed and signed by operators.
    // these signatures are aggregated and sent to the contract as response.
    struct TaskResponse {
        // Can be obtained by the operator from the event NewTaskCreated.
        uint32 referenceTaskIndex;
        // This is just the response that the operator has to compute by itself.
        string taskResponse;
    }

    // Extra information related to taskResponse, which is filled inside the contract.
    // It thus cannot be signed by operators, so we keep it in a separate struct than TaskResponse
    // This metadata is needed by the challenger, so we emit it in the TaskResponded event
    struct TaskResponseMetadata {
        uint32 taskRespondedBlock;
        bytes32 hashOfNonSigners;
    }

    //=========================================================================
    //                                FUNCTIONS
    //=========================================================================
    function createNewTask(string calldata taskDetails, uint32 quorumThresholdPercentage, bytes calldata quorumNumbers)
        external;

    function respondToTask(
        Task calldata task,
        TaskResponse calldata taskResponse,
        IBLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external;

    function taskNumber() external view returns (uint32);
}
