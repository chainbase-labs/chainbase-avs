// SPDX-License-Identifier: UNLICENSED
pragma solidity = 0.8.24;

import "@eigenlayer/contracts/permissions/Pausable.sol";
import "@eigenlayer-middleware/src/ServiceManagerBase.sol";
import "@eigenlayer-middleware/src/BLSSignatureChecker.sol";
import "@eigenlayer-middleware/src/OperatorStateRetriever.sol";

import "./ChainbaseServiceManagerStorage.sol";

contract ChainbaseServiceManager is
    Pausable,
    BLSSignatureChecker,
    OperatorStateRetriever,
    ServiceManagerBase,
    ChainbaseServiceManagerStorage
{
    //=========================================================================
    //                                MODIFIERS
    //=========================================================================
    /**
     * @dev Modifier to restrict function calls to the aggregator only
     */
    modifier onlyAggregator() {
        require(msg.sender == aggregator, "ChainbaseServiceManager: aggregator must be the caller");
        _;
    }

    //=========================================================================
    //                                CONSTRUCTOR
    //=========================================================================
    /**
     * @dev Constructor to initialize the contract with required addresses
     * @param _avsDirectory Address of the AVSDirectory
     * @param _registryCoordinator Address of the RegistryCoordinator
     * @param _stakeRegistry Address of the StakeRegistry
     */
    constructor(
        IAVSDirectory _avsDirectory, // AVSDirectory address
        IRegistryCoordinator _registryCoordinator, // RegistryCoordinator address
        IStakeRegistry _stakeRegistry // StakeRegistry address
    )
        BLSSignatureChecker(_registryCoordinator)
        ServiceManagerBase(_avsDirectory, _registryCoordinator, _stakeRegistry)
    {}

    //=========================================================================
    //                                INITIALIZE
    //=========================================================================
    /**
     * @dev Initialize function to set the initial owner and other contract addresses
     * @param _pauserRegistry Address of the PauserRegistry
     * @param initialOwner Address of the initial owner
     * @param _aggregator Address of the aggregator
     * @param _generator Address of the generator
     */
    function initialize(IPauserRegistry _pauserRegistry, address initialOwner, address _aggregator, address _generator)
        public
        initializer
    {
        _initializePauser(_pauserRegistry, UNPAUSE_ALL);
        __ServiceManagerBase_init(initialOwner);
        aggregator = _aggregator;
        generator = _generator;
    }

    //=========================================================================
    //                                EXTERNAL
    //=========================================================================
    /**
     * @notice External function to create a new task
     * @param taskDescription Description of the task
     * @param quorumThresholdPercentage Threshold percentage for quorum
     * @param quorumNumbers List of quorum numbers
     */
    function createNewTask(
        string calldata taskDescription,
        uint32 quorumThresholdPercentage,
        bytes calldata quorumNumbers
    ) external override {
        // create a new task struct
        Task memory newTask;
        newTask.taskDescription = taskDescription;
        newTask.taskCreatedBlock = uint32(block.number);
        newTask.quorumThresholdPercentage = quorumThresholdPercentage;
        newTask.quorumNumbers = quorumNumbers;

        // store hash of task onchain, emit event, and increase taskNum
        allTaskHashes[latestTaskNum] = keccak256(abi.encode(newTask));
        emit NewTaskCreated(latestTaskNum, newTask);
        latestTaskNum = latestTaskNum + 1;
    }

    /**
     * @notice Returns the current 'taskNumber' of the middleware
     * @return The current task number
     */
    function taskNumber() external view override returns (uint32) {
        return latestTaskNum;
    }

    //=========================================================================
    //                                AGGREGATOR
    //=========================================================================
    /**
     * @notice External function to respond to existing tasks
     * @param task Task struct containing task details
     * @param taskResponse Task response struct containing the response details
     * @param nonSignerStakesAndSignature Struct containing non-signer stakes and the signature
     */
    function respondToTask(
        Task calldata task,
        TaskResponse calldata taskResponse,
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external onlyAggregator {
        uint32 taskCreatedBlock = task.taskCreatedBlock;
        bytes calldata quorumNumbers = task.quorumNumbers;
        uint32 quorumThresholdPercentage = task.quorumThresholdPercentage;
        // check that the task is valid, hasn't been responded yet, and is being responded in time
        require(
            keccak256(abi.encode(task)) == allTaskHashes[taskResponse.referenceTaskIndex],
            "ChainbaseServiceManager: supplied task does not match the one recorded in the contract"
        );
        require(
            allTaskResponses[taskResponse.referenceTaskIndex] == bytes32(0),
            "ChainbaseServiceManager: aggregator has already responded to the task"
        );
        require(
            uint32(block.number) <= taskCreatedBlock + TASK_RESPONSE_WINDOW_BLOCK,
            "ChainbaseServiceManager: aggregator has responded to the task too late"
        );

        /* CHECKING SIGNATURES & WHETHER THRESHOLD IS MET OR NOT */
        // calculate message which operators signed
        bytes32 message = keccak256(abi.encode(taskResponse));

        // check the BLS signature
        (QuorumStakeTotals memory quorumStakeTotals, bytes32 hashOfNonSigners) =
            checkSignatures(message, quorumNumbers, taskCreatedBlock, nonSignerStakesAndSignature);

        // check that signatories own at least a threshold percentage of each quorum
        for (uint256 i = 0; i < quorumNumbers.length; i++) {
            // we don't check that the quorumThresholdPercentages are not >100 because a greater value would trivially fail the check, implying
            // signed stake > total stake
            require(
                quorumStakeTotals.signedStakeForQuorum[i] * _THRESHOLD_DENOMINATOR
                    >= quorumStakeTotals.totalStakeForQuorum[i] * uint8(quorumThresholdPercentage),
                "ChainbaseServiceManager: signatories do not own at least threshold percentage of a quorum"
            );
        }

        TaskResponseMetadata memory taskResponseMetadata = TaskResponseMetadata(uint32(block.number), hashOfNonSigners);
        // updating the storage with task response
        allTaskResponses[taskResponse.referenceTaskIndex] = keccak256(abi.encode(taskResponse, taskResponseMetadata));

        // emitting event
        emit TaskResponded(taskResponse, taskResponseMetadata);
    }
}
