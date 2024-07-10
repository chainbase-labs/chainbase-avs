// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import {IAVS} from "./interface/IAVS.sol";
import {AVSStorage} from "./AVSStorage.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IAVSDirectory} from "@eigenlayer/contracts/interfaces/IAVSDirectory.sol";
import {ISignatureUtils} from "@eigenlayer/contracts/interfaces/ISignatureUtils.sol";
import "../Script/StrategyParams.s.sol";

import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

contract AVS is IAVS, OwnableUpgradeable, AVSStorage {
    /// @notice EigenLayer core DelegationManager
    IDelegationManager internal immutable _delegationManager;

    /// @notice EigenLayer core AVSDirectory
    IAVSDirectory internal immutable _avsDirectory;

    constructor(IDelegationManager delegationManager, IAVSDirectory avsDirectory) {
        _delegationManager = delegationManager;
        _avsDirectory = avsDirectory;

        _disableInitializers();
    }

    /**
     * @notice Initialize the Chainbase AVS  contract.
     * @param metadataURI_      Metadata URI for the AVS
     */
    function initialize(string calldata metadataURI_) external initializer {
        _setStrategyParams(StrategyParams.holesky());

        if (bytes(metadataURI_).length > 0) {
            _avsDirectory.updateAVSMetadataURI(metadataURI_);
        }
    }

    /// ************ view function ************ ///
    /// ****************************************** ///

    function operators() external view returns (address[] memory) {
        return _operators;
    }

    function strategyParams() external view returns (IAVS.StrategyParam[] memory) {
        return _strategyParams;
    }

    /**
     * @notice Returns true if the operator is in the allowlist.
     * @param operator The operator to check
     */
    function isInAllowlist(address operator) external view returns (bool) {
        return _allowlist[operator];
    }

    function canRegister(address operator) external view returns (bool, string memory) {
        if (_operatorRegistered[operator]) {
            return (false, "Operator already registered");
        }
        return (true, "");
    }

    /// ************ external function ************ ///
    /// ****************************************** ///
    function registerOperator(ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature) external {
        address operator = msg.sender;
        _avsDirectory.registerOperatorToAVS(operator, operatorSignature);
        _operators.push(operator);
        _operatorRegistered[operator] = true;
        emit OperatorAdded(operator);
    }

    function getRestakeableStrategies() external view returns (address[] memory) {
        return _getRestakeableStrategies();
    }

    /**
     * @notice Returns the list of strategies that the operator has potentially restaked on the AVS
     * @dev Implemented to match IServiceManager interface - required for compatibility with
     *      eigenlayer frontend.
     *
     *      This function is intended to be called off-chain
     *
     *      No guarantee is made on whether the operator has shares for a strategy. The off-chain
     *      service should do that validation separately. This matches the behavior defined in
     *      eigenlayer-middleware's ServiceManagerBase.
     *
     * @param operator The address of the operator to get restaked strategies for
     */
    function getOperatorRestakedStrategies(address operator) external view returns (address[] memory) {
        if (!_operatorRegistered[operator]) return new address[](0);
        return _getRestakeableStrategies();
    }

    /// ************ internal function ************ ///
    /// ****************************************** ///

    /**
     * @notice Returns the list of restakeable strategy addresses
     */
    function _getRestakeableStrategies() internal view returns (address[] memory) {
        address[] memory strategies = new address[](_strategyParams.length);
        for (uint256 i = 0; i < _strategyParams.length;) {
            strategies[i] = address(_strategyParams[i].strategy);
            unchecked {
                i++;
            }
        }
        return strategies;
    }

    /**
     * @notice Set the strategy parameters.
     * @param params The strategy parameters
     */
    function _setStrategyParams(StrategyParam[] memory params) private {
        delete _strategyParams;

        for (uint256 i = 0; i < params.length;) {
            require(address(params[i].strategy) != address(0), "AVS: no zero strategy");
            require(params[i].multiplier > 0, "AVS: no zero multiplier");

            // ensure no duplicates
            for (uint256 j = i + 1; j < params.length;) {
                require(address(params[i].strategy) != address(params[j].strategy), "AVS: no duplicate strategy");
                unchecked {
                    j++;
                }
            }

            _strategyParams.push(params[i]);
            unchecked {
                i++;
            }
        }

        emit StrategyParamsSet(params);
    }
}
