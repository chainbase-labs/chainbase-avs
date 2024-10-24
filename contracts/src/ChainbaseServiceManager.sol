// SPDX-License-Identifier: UNLICENSED
pragma solidity = 0.8.24;

import "@eigenlayer-middleware/src/ServiceManagerBase.sol";
import "@eigenlayer-middleware/src/BLSSignatureChecker.sol";

import "./ChainbaseServiceManagerStorage.sol";

contract ChainbaseServiceManager is BLSSignatureChecker, ServiceManagerBase, ChainbaseServiceManagerStorage {
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

    /**
     * @dev Modifier to restrict function calls to the generator only
     */
    modifier onlyGenerator() {
        require(msg.sender == generator, "ChainbaseServiceManager: generator must be the caller");
        _;
    }

    /**
     * @dev Modifier to restrict operator in whitelist
     */
    modifier onlyWhitelisted(address operator) {
        require(!whitelistEnabled || operatorWhitelist[operator], "ChainbaseServiceManager: operator not in whitelist");
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
     * @param initialOwner Address of the initial owner
     * @param _aggregator Address of the aggregator
     * @param _generator Address of the generator
     */
    function initialize(address initialOwner, address _aggregator, address _generator) public initializer {
        __ServiceManagerBase_init(initialOwner);
        aggregator = _aggregator;
        generator = _generator;
    }

    //=========================================================================
    //                                 MANAGE
    //=========================================================================
    /**
     * @dev Sets the aggregator.
     * @param _aggregator The address of the aggregator.
     */
    function setAggregator(address _aggregator) external onlyOwner {
        aggregator = _aggregator;
    }

    /**
     * @dev Sets the generator.
     * @param _generator The address of the generator.
     */
    function setGenerator(address _generator) external onlyOwner {
        generator = _generator;
    }

    /**
     * @dev Add the operator to the whitelist.
     * @param operators The address of the operators.
     */
    function addOperatorsToWhitelist(address[] calldata operators) external onlyOwner {
        for (uint256 i = 0; i < operators.length; i++) {
            operatorWhitelist[operators[i]] = true;
        }
    }

    /**
     * @dev Remove the operator from the whitelist.
     * @param operators The address of the operators.
     */
    function removeOperatorsFromWhitelist(address[] calldata operators) external onlyOwner {
        for (uint256 i = 0; i < operators.length; i++) {
            delete operatorWhitelist[operators[i]];
        }
    }

    /**
     * @dev Set the whitelist enabled status.
     * @param _whitelistEnabled The status to set.
     */
    function setWhitelistEnabled(bool _whitelistEnabled) external onlyOwner {
        whitelistEnabled = _whitelistEnabled;
    }

    //=========================================================================
    //                                EXTERNAL
    //=========================================================================
    /**
     * @notice Forwards a call to EigenLayer's AVSDirectory contract to confirm operator registration with the AVS
     * @param operator The address of the operator to register.
     * @param operatorSignature The signature, salt, and expiry of the operator's signature.
     */
    function registerOperatorToAVS(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) public override onlyWhitelisted(operator) {
        super.registerOperatorToAVS(operator, operatorSignature);
    }

    /**
     * @notice Forwards a call to EigenLayer's AVSDirectory contract to confirm operator deregistration from the AVS
     * @param operator The address of the operator to deregister.
     */
    function deregisterOperatorFromAVS(address operator) public override onlyWhitelisted(operator) {
        super.deregisterOperatorFromAVS(operator);
    }

    /**
     * @notice External function to create a new task
     * @param taskDetails Details of the task
     * @param quorumThresholdPercentage Threshold percentage for quorum
     * @param quorumNumbers List of quorum numbers
     */
    function createNewTask(string calldata taskDetails, uint32 quorumThresholdPercentage, bytes calldata quorumNumbers)
        external
        override
        onlyGenerator
    {
        // create a new task struct
        Task memory newTask;
        newTask.taskDetails = taskDetails;
        newTask.taskCreatedBlock = uint32(block.number);
        newTask.quorumThresholdPercentage = quorumThresholdPercentage;
        newTask.quorumNumbers = quorumNumbers;

        // store hash of task onchain, emit event, and increase taskNum
        allTaskHashes[latestTaskNum] = keccak256(abi.encode(newTask));
        emit NewTaskCreated(latestTaskNum, newTask);
        latestTaskNum = latestTaskNum + 1;
    }

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
    ) external override onlyAggregator {
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
                     >= quorumStakeTotals.totalStakeForQuorum[i] * quorumThresholdPercentage,
                "ChainbaseServiceManager: signatories do not own at least threshold percentage of a quorum"
            );
        }

        TaskResponseMetadata memory taskResponseMetadata = TaskResponseMetadata(uint32(block.number), hashOfNonSigners);
        // updating the storage with task response
        allTaskResponses[taskResponse.referenceTaskIndex] = keccak256(abi.encode(taskResponse, taskResponseMetadata));

        // emitting event
        emit TaskResponded(taskResponse, taskResponseMetadata);
    }

    /**
     * @notice Returns the current 'taskNumber' of the middleware
     * @return The current task number
     */
    function taskNumber() external view override returns (uint32) {
        return latestTaskNum;
    }
}
