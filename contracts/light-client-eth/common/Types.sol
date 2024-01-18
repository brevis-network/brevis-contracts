// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../../verifiers/interfaces/IBeaconVerifier.sol";

struct LightClientOptimisticUpdate {
    // Header attested to by the sync committee
    HeaderWithExecution attestedHeader;
    // Sync committee aggregate signature participation & zk proof
    SyncAggregate syncAggregate;
    // Slot at which the aggregate signature was created (untrusted)
    uint64 signatureSlot;
}

struct LightClientUpdate {
    // Header attested to by the sync committee
    HeaderWithExecution attestedHeader;
    HeaderWithExecution finalizedHeader;
    // merkle branch from finalized beacon header root to attestedHeader.stateRoot
    bytes32[] finalityBranch;
    bytes32 nextSyncCommitteeRoot;
    bytes32[] nextSyncCommitteeBranch;
    bytes32 nextSyncCommitteePoseidonRoot;
    IBeaconVerifier.Proof nextSyncCommitteeRootMappingProof;
    // Sync committee aggregate signature participation & zk proof
    SyncAggregate syncAggregate;
    // Slot at which the aggregate signature was created (untrusted)
    uint64 signatureSlot;
}

struct HeaderWithExecution {
    BeaconBlockHeader beacon;
    ExecutionPayload execution;
    // merkle branch from execution payload root to beacon block root
    LeafWithBranch executionRoot;
}

function isEmpty(HeaderWithExecution memory header) pure returns (bool) {
    return header.beacon.stateRoot == bytes32(0);
}

// only contains the fields we care about in execution payload
struct ExecutionPayload {
    // merkle branch from execution state root to execution payload root
    LeafWithBranch stateRoot;
    // merkle branch from execution block hash to execution payload root
    LeafWithBranch blockHash;
    // merkle branch from execution block number to execution payload root
    LeafWithBranch blockNumber;
}

function isEmpty(ExecutionPayload memory payload) pure returns (bool) {
    return
        payload.stateRoot.leaf == bytes32(0) &&
        payload.blockHash.leaf == bytes32(0) &&
        payload.blockNumber.leaf == bytes32(0);
}

struct LeafWithBranch {
    bytes32 leaf;
    bytes32[] branch;
}

struct BeaconBlockHeader {
    uint64 slot;
    uint64 proposerIndex;
    bytes32 parentRoot;
    bytes32 stateRoot;
    bytes32 bodyRoot;
}

struct SyncAggregate {
    uint64 participation;
    bytes32 poseidonRoot;
    uint256 commitment;
    IBeaconVerifier.Proof proof;
}
