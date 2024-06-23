// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import {IStrategy} from "@eigenlayer/contracts/interfaces/IStrategy.sol";
import {ISignatureUtils} from "@eigenlayer/contracts/interfaces/ISignatureUtils.sol";

/**
 * @title IAVS
 * @notice Interface for the AVS contract.
 */
interface IAVS {
    /**
     * @notice Emitted when an operator is added to the AVS.
     * @param operator The address of the operator
     */
    event OperatorAdded(address indexed operator);

    /**
     * @notice Emitted when an operator is removed from the AVS.
     * @param operator The address of the operator
     */
    event OperatorRemoved(address indexed operator);

    /**
     * @notice Struct representing an AVS operator
     * @custom:field addr       The operator's ethereum address
     * @custom:field pubkey     The operator's 64 byte uncompressed secp256k1 public key
     * @custom:field delegated  The total amount delegated, not including operator stake
     * @custom:field staked     The total amount staked by the operator, not including delegations
     */
    struct Operator {
        address addr;
        bytes pubkey;
        uint96 delegated;
        uint96 staked;
    }

    /**
     * @notice Represents a single supported strategy.
     * @custom:field strategy   The strategy contract
     * @custom:field multiplier The stake multiplier, to weight strategy against others
     */
    struct StrategyParam {
        IStrategy strategy;
        uint96 multiplier;
    }

    /**
     * @notice Returns the currrent address of operator registered as AVS.
     */
    function operators() external view returns (address[] memory);

    /**
     * @notice Returns the current strategy parameters.
     */
    function strategyParams() external view returns (StrategyParam[] memory);

    /**
     * @notice Register an operator with the AVS. Forwards call to EigenLayer' AVSDirectory.
     * @param pubkey            64 byte uncompressed secp256k1 public key (no 0x04 prefix)
     *                          Pubkey must match operator's address (msg.sender)
     * @param operatorSignature The signature, salt, and expiry of the operator's signature.
     */
    // function registerOperator(
    //     bytes calldata pubkey,
    //     ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    // ) external;

    /**
     * @notice Check if an operator can register to the AVS.
     *         Returns true, with no reason, if the operator can register to the AVS.
     *         Returns false, with a reason, if the operator cannot register to the AVS.
     * @dev This function is intented to be called off-chain.
     * @param operator The operator to check
     * @return canRegister True if the operator can register, false otherwise
     * @return reason      The reason the operator cannot register. Empty if canRegister is true.
     */
    function canRegister(address operator) external view returns (bool, string memory);
}
