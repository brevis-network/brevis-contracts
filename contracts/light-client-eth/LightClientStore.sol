// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "./common/Types.sol";
import "../verifiers/interfaces/IBeaconVerifier.sol";

abstract contract LightClientStore {
    // beacon chain genesis information
    uint256 immutable GENESIS_TIME;
    bytes32 immutable GENESIS_VALIDATOR_ROOT;

    uint64 public finalizedSlot;
    bytes32 public finalizedExecutionStateRoot;

    uint64 public optimisticSlot;
    bytes32 public optimisticExecutionStateRoot;

    bytes32 public currentSyncCommitteeRoot;
    bytes32 public currentSyncCommitteePoseidonRoot;
    bytes32 public nextSyncCommitteeRoot;
    bytes32 public nextSyncCommitteePoseidonRoot;

    LightClientUpdate public bestValidUpdate;

    // fork versions
    uint64[] public forkEpochs;
    bytes4[] public forkVersions;

    // zk verifier
    IBeaconVerifier public zkVerifier; // contract too big. need to move this one out

    constructor(
        uint256 genesisTime,
        bytes32 genesisValidatorsRoot,
        uint64[] memory _forkEpochs,
        bytes4[] memory _forkVersions,
        uint64 _finalizedSlot,
        bytes32 syncCommitteeRoot,
        bytes32 syncCommitteePoseidonRoot,
        address _zkVerifier
    ) {
        GENESIS_TIME = genesisTime;
        GENESIS_VALIDATOR_ROOT = genesisValidatorsRoot;
        forkEpochs = _forkEpochs;
        forkVersions = _forkVersions;
        finalizedSlot = _finalizedSlot;
        currentSyncCommitteeRoot = syncCommitteeRoot;
        currentSyncCommitteePoseidonRoot = syncCommitteePoseidonRoot;
        zkVerifier = IBeaconVerifier(_zkVerifier);
    }
}
