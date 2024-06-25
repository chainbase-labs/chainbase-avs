// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import {IAVS} from "./interface/IAVS.sol";
import {AVSStorage} from "./AVSStorage.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IAVSDirectory} from "@eigenlayer/contracts/interfaces/IAVSDirectory.sol";

contract AVS is IAVS, AVSStorage {
    /// @notice EigenLayer core DelegationManager
    IDelegationManager internal immutable _delegationManager;

    /// @notice EigenLayer core AVSDirectory
    IAVSDirectory internal immutable _avsDirectory;

    constructor(IDelegationManager delegationManager, IAVSDirectory avsDirectory, string memory metadataURI_) {
        _delegationManager = delegationManager;
        _avsDirectory = avsDirectory;

        if (bytes(metadataURI_).length > 0) {
            _avsDirectory.updateAVSMetadataURI(metadataURI_);
        }
    }

    /// view function

    function operators() external view returns (address[] memory) {
        return _operators;
    }

    function strategyParams() external view returns (IAVS.StrategyParam[] memory) {
        return _strategyparams;
    }

    /**
     * @notice Returns true if the operator is in the allowlist.
     * @param operator The operator to check
     */
    function isInAllowlist(address operator) external view returns (bool) {
        return _allowlist[operator];
    }

    function canRegister(address operator) external view returns (bool, string memory) {
        return (true, "");
    }
}
