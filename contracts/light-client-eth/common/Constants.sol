// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

// light client security params
uint256 constant MIN_SYNC_COMMITTEE_PARTICIPANTS = 1;
uint256 constant UPDATE_TIMEOUT = 86400;

// beacon chain constants
uint256 constant FINALIZED_ROOT_INDEX = 105;
uint256 constant NEXT_SYNC_COMMITTEE_INDEX = 55;
uint256 constant SYNC_COMMITTEE_SIZE = 512;
uint64 constant SLOTS_PER_EPOCH = 32;
uint64 constant EPOCHS_PER_SYNC_COMMITTEE_PERIOD = 256;
bytes32 constant DOMAIN_SYNC_COMMITTEE = bytes32(uint256(0x07) << 248);
uint256 constant SLOT_LENGTH_SECONDS = 12;

// https://github.com/ethereum/consensus-specs/blob/dev/specs/capella/light-client/sync-protocol.md
// beaconBodyRoot -> stateRoot gindex: 2 << 7 | 9 * 2 << 3 | 2
uint256 constant EXECUTION_STATE_ROOT_INDEX = 402;
// beaconBodyRoot -> blockHash gindex: 2 << 7 | 9 * 2 << 3 | 12
uint256 constant EXECUTION_BLOCK_HASH_INDEX = 412;

// the following indices are gindices counting from the executionPayloadRoot
// beaconBodyRoot -> executionPayloadRoot gindex: 2 << 4 | 9
uint256 constant EXECUTION_PAYLOAD_ROOT_INDEX = 25;
// executionPayloadRoot -> stateRoot gindex: 2 << 4 | 2
uint256 constant EXECUTION_STATE_ROOT_LOCAL_INDEX = 18;
// executionPayloadRoot -> blockNumber gindex: 2 << 4 | 6
uint256 constant EXECUTION_BLOCK_NUMBER_LOCAL_INDEX = 22;
// executionPayloadRoot -> blockHash gindex: 2 << 4 | 12
uint256 constant EXECUTION_BLOCK_HASH_LOCAL_INDEX = 28;
