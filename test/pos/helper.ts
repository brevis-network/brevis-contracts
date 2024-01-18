import { Gindex, ProofType, SingleProof, createProof } from '@chainsafe/persistent-merkle-tree';
import { ByteVectorType, ContainerType, VectorCompositeType } from '@chainsafe/ssz';
import dotenv from 'dotenv';
import { BigNumberish } from 'ethers';
import { BytesLike, arrayify } from 'ethers/lib/utils';
import { ethers } from 'hardhat';
import { ExecutionPayloadStruct, LeafWithBranchStruct } from '../../typechain/AnchorBlocks';
import { BeaconBlockHeaderStruct, LightClientUpdateStruct } from '../../typechain/EthereumLightClient';
import { ssz } from '../../vendor/lodestar/types';
import { ProofStruct } from './../../typechain/ILightClientZkVerifier';
import { HeaderWithExecutionStruct } from './../../typechain/LightClientStore';
import proof from './proof_638.json';
import update from './update_638.json';

dotenv.config();

export const ZERO_BYTES_32 = '0x0000000000000000000000000000000000000000000000000000000000000000';

export type Groth16Proof = typeof proof.bls_sig_proof.proof;

export type BLSProof = typeof proof.bls_sig_proof;

export type UpdateProof = typeof proof;

export type LightClientUpdate = typeof update.data;

export function sumCommitteeBits(bits: string): number {
  return ethers.utils.arrayify(bits).reduce((acc, curr) => {
    // curr is 0~255
    let sum = 0;
    // curr.toString(2) is 00000000~11111111
    for (const digit of curr.toString(2)) {
      const d = parseInt(digit);
      sum += d;
    }
    return acc + sum;
  }, 0);
}

export function getSyncPeriodBySlot(slot: number): number {
  return Math.floor(slot / 32 / 256);
}

export function getSyncCommitteeRoot(pubkeys: string[], aggregatePubkey: string) {
  const pubkeyType = new ByteVectorType(48);
  const pubkeysType = new VectorCompositeType(pubkeyType, 512);
  const rootType = new ContainerType({
    pubkeys: pubkeysType,
    aggregatePubkey: pubkeyType
  });
  const root = rootType.hashTreeRoot({
    pubkeys: pubkeys.map((key) => arrayify(key)),
    aggregatePubkey: arrayify(aggregatePubkey)
  });
  return ethers.utils.hexlify(root);
}

export type HeaderWithExecution = typeof update.data.attested_header;
export type BeaconBlockHeader = typeof update.data.attested_header.beacon;
export type ExecutionPayloadHeader = typeof update.data.attested_header.execution;

function getExecutionPayloadRoot(eh: ExecutionPayloadHeader) {
  const t = ssz.capella.ExecutionPayloadHeader;
  const v = ssz.capella.ExecutionPayloadHeader.fromJson(eh);
  const root = ethers.utils.hexlify(t.hashTreeRoot(v));
  return root;
}

function getTreeProof(eh: ExecutionPayloadHeader, gindex: bigint) {
  const t = ssz.capella.ExecutionPayloadHeader;
  const v = ssz.capella.ExecutionPayloadHeader.fromJson(eh);
  const tree = t.value_toTree(v);
  const p = createProof(tree, {
    type: ProofType.single,
    gindex: gindex as Gindex
  }) as SingleProof;
  return p;
}

export function newLeafWithBranch(l: BytesLike, b: BytesLike[]): LeafWithBranchStruct {
  return { leaf: l, branch: b };
}
export function emptyLeafWithBranch(): LeafWithBranchStruct {
  return { leaf: ZERO_BYTES_32, branch: [] };
}
export function newBeaconBlockHeader(b: BeaconBlockHeader): BeaconBlockHeaderStruct {
  return {
    bodyRoot: b.body_root,
    parentRoot: b.parent_root,
    proposerIndex: b.proposer_index,
    slot: b.slot,
    stateRoot: b.state_root
  };
}

export function emptyBeaconBlockHeader(): BeaconBlockHeaderStruct {
  return {
    bodyRoot: ZERO_BYTES_32,
    parentRoot: ZERO_BYTES_32,
    proposerIndex: 0,
    slot: 0,
    stateRoot: ZERO_BYTES_32
  };
}

export function emptyExecution(): ExecutionPayloadStruct {
  return {
    stateRoot: emptyLeafWithBranch(),
    blockNumber: emptyLeafWithBranch(),
    blockHash: emptyLeafWithBranch()
  };
}

export function newHeaderWithExecution(h: HeaderWithExecution): HeaderWithExecutionStruct {
  const stateRoot = getTreeProof(h.execution, BigInt(18));
  const blockNumber = getTreeProof(h.execution, BigInt(22));
  const blockHash = getTreeProof(h.execution, BigInt(28));
  return {
    beacon: newBeaconBlockHeader(h.beacon),
    execution: {
      stateRoot: newLeafWithBranch(stateRoot.leaf, stateRoot.witnesses),
      blockNumber: newLeafWithBranch(blockNumber.leaf, blockNumber.witnesses),
      blockHash: newLeafWithBranch(blockHash.leaf, blockHash.witnesses)
    },
    executionRoot: newLeafWithBranch(getExecutionPayloadRoot(h.execution), h.execution_branch)
  };
}

export function emptyHeaderWithExecution(): HeaderWithExecutionStruct {
  return {
    beacon: emptyBeaconBlockHeader(),
    execution: emptyExecution(),
    executionRoot: emptyLeafWithBranch()
  };
}

export function newSyncAggregate(participationBits: string, p: BLSProof) {
  return {
    participation: sumCommitteeBits(participationBits),
    poseidonRoot: p.poseidon_root,
    commitment: p.commitment_pub,
    proof: newGroth16Proof(p.proof)
  };
}

function newPoint(point: string[]): [BigNumberish, BigNumberish] {
  return [point[0], point[1]];
}

function newGroth16Proof(p: Groth16Proof): ProofStruct {
  return {
    a: newPoint(p.a),
    b: [newPoint(p.b[0]), newPoint(p.b[1])],
    c: newPoint(p.c),
    commitment: newPoint(p.commitment)
  };
}

function emptyGroth16Proof(): ProofStruct {
  return {
    a: newPoint(['0', '0']),
    b: [newPoint(['0', '0']), newPoint(['0', '0'])],
    c: newPoint(['0', '0']),
    commitment: newPoint(['0', '0'])
  };
}

export function newLightClientUpdate(u: LightClientUpdate, p: UpdateProof): LightClientUpdateStruct {
  return {
    attestedHeader: newHeaderWithExecution(u.attested_header),
    finalizedHeader: newHeaderWithExecution(u.finalized_header),
    finalityBranch: u.finality_branch,
    nextSyncCommitteeRoot: getSyncCommitteeRoot(u.next_sync_committee.pubkeys, u.next_sync_committee.aggregate_pubkey),
    nextSyncCommitteeBranch: u.next_sync_committee_branch,
    nextSyncCommitteePoseidonRoot: p.committee_map_proof.poseidon_ssz_root,
    nextSyncCommitteeRootMappingProof: newGroth16Proof(p.committee_map_proof.proof),
    signatureSlot: u.signature_slot,
    syncAggregate: newSyncAggregate(u.sync_aggregate.sync_committee_bits, p.bls_sig_proof)
  };
}

export function newOptimisticUpdate(u: LightClientUpdate, p: UpdateProof): LightClientUpdateStruct {
  return {
    attestedHeader: newHeaderWithExecution(u.attested_header),
    finalizedHeader: emptyHeaderWithExecution(),
    finalityBranch: [],
    nextSyncCommitteeRoot: ZERO_BYTES_32,
    nextSyncCommitteeBranch: [],
    nextSyncCommitteePoseidonRoot: ZERO_BYTES_32,
    nextSyncCommitteeRootMappingProof: emptyGroth16Proof(),
    signatureSlot: u.signature_slot,
    syncAggregate: newSyncAggregate(u.sync_aggregate.sync_committee_bits, p.bls_sig_proof)
  };
}
