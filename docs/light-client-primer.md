# Light Client Primer

## Light Client & Sync Protocol

The light client is an Ethereum client that keeps track of a set of block headers. Without light clients, we can still verify the validity of data on the Ethereum chain, but it requires the knowledge of all validators and beacon committee assignments (essentially part of the job of a full client). To enable lower-spec machines to be able to verify blockchain's data with reasonable security, the sync protocol is implemented in Ethereum consensus.

### Sync Committee

Validators that participate in the sync protocol form a sync committee.

> The sync committee is the "flagship feature" of the Altair hard fork. This is a committee of 512 validators that is randomly selected every sync committee period (~1 day), and while a validator is part of the currently active sync committee they are expected to continually sign the block header that is the new head of the chain at each slot.

Note that a sync committee is different from a beacon committee whose job is to sign attestions to newly proposed beacon chain blocks at each slot.

## Light Client Implementation

The implementation of [succinct lab’s light client](https://github.com/succinctlabs/eth-proof-of-consensus/blob/main/contracts/src/lightclient/BeaconLightClient.sol) follows the light client design specified in the [Ethereum light client spec](https://github.com/ethereum/consensus-specs/blob/dev/specs/altair/light-client/sync-protocol.md#process_light_client_finality_update). We'll use this implementation as an example throughout the rest of the doc.

### Minimal Light Client State

In order to allow users of the light client to check whether some information is contained in a specific beacon block, a mapping of slot to beacon block header needs to be stored.

```c
struct BeaconBlockHeader {
    uint64 slot;
    bytes32 stateRoot;
    // ... other fields omitted for now
}

mapping(uint64 => BeaconBlockHeader) public headers;
mapping(uint256 => bytes32) public syncCommitteeRootByPeriod;
```

### `headers` State

The most valuable information in the beacon block header is the `stateRoot` since if you can provide some data and a merkle path to this `stateRoot` you can prove that the block contains the data.

The ultimate goal of the light client is to have every block (or maybe only the blocks we care about) that are on the canonical Ethereum beacon chain also appear in the `headers` state.

### `syncCommitteeRootByPeriod` State

Sync committees' ssz merkle roots are saved in this state, keyed by sync period. The saved sync committee root is used as the public input of the BLS signature proof in every subsequent `headers` update in the same sync period. (Note that in the implementation, this root is further mapped to a Poseidon merkle root for zk efficiency reasons. More on this later)

## Updating the Light Client's State

### Updating the Headers

```c
struct BLSAggregatedSignature {
    uint64 participation;
    Groth16Proof proof;
}

struct LightClientUpdate {
    BeaconBlockHeader attestedHeader;
    BeaconBlockHeader finalizedHeader;
    bytes32[] finalityBranch;
    bytes32 executionStateRoot;
    bytes32[] executionStateRootBranch;
    BLSAggregatedSignature signature;
    // ... other fields omitted for now
}
```

The first step in updating the light client's state is always to fetch existing record of the sync committee commitment from `syncCommitteeRootByPeriod` as it's used for verifying the signature

Let's take a look at each field for what purpose they serve.

### `attestedHeader` & `finalizedHeader`

The field we care about the most is `finalizedHeader`. A block is said to be "finalized" when it's state root appears in the beacon state trie index 105. We can know this by applying `ssz_root(finalizedHeader)` to the `finalityBranch` at [generalized index](https://ethereum.org/en/developers/docs/data-structures-and-encoding/ssz/#generalized-indices) = 105 and see if it computes `attestedHeader.stateRoot`. If it's true, then we know the beacon chain state corresponding to the `attestedHeader` contains finality info about that block. If this update also has `signature.participation` > 2/3 of the its corresponding sync committee members, the light client can safely put it into the `headers` state.

### `executionStateRoot`

This field is essentially what enables MessageBridge. `executionStateRoot` is basically eth1's state root. (TODO explain further)

### Why have both `attestedHeader` and `finalizedHeader`?

My understanding is that because block finality (as defined by the consensus spec) cannot be reached in the same epoch the block is proposed. It means in order to show that a block is finalized, another block in a subsequent epoch must be presented to the light client.

### Updating the Sync Committee

Since sync committee rotates every 256 epochs, we need a way to update our `syncCommitteeRootByPeriod`.

Previously we ommitted some fields in the `LightClientUpdate` struct. With sync committee related fields added in, the struct now look like this

```c
struct LightClientUpdate {
    BeaconBlockHeader attestedHeader;
    BeaconBlockHeader finalizedHeader;
    bytes32[] finalityBranch;
    bytes32 executionStateRoot;
    bytes32[] executionStateRootBranch;
    BLSAggregatedSignature signature;

    bytes32 nextSyncCommitteeRoot;
    bytes32[] nextSyncCommitteeBranch;
}
```

Since committee info only needs to be updated once per 256 epochs, there is a function `updateSyncCommittee()` to update it separate from the one that updates just a block. The steps to update sync commitee is the following:

1. compute the period number from `update.finalizedHeader.slot` (`update.attestedHeader.slot` if no finality proof is provided, more on this later).
2. verify the merkle branch `update.nextSyncCommitteebranch` from `update.nextSyncCommitteeRoot` to `update.finalizedHeader.stateRoot` (generalized index = 55). This proves that the new committe we are replacing the old one with exists in the latest beacon state. (Note: I can't figure out why Succinct Lab uses `update.finalizedHeader.stateRoot` as the target root here even though finality is not required when calling `updateSyncCommittee()`. It should be `update.attestedHeader` when finality proof is not present?)
3. if the update has a finality proof, the new sync commitee is added to the `syncCommitteeRootByPeriod` state, keyed by the new period number.
4. else, if the update is a "better" one comparing to the one stored in a temp storage, replace it. More on whatever "better" means and why we need to store this at all later.

#### What If the Beacon Chain Keeps Failing to Reach Finality

Since sync committee can only be updated upon finalized blocks, if the chain fails to reach finality for an extended period of time, the aforementioned "better" update can be used to force finalize a cached block and update the sync committee. So what is a better update?

### Sync Committee Commitment

#### Poseidon Merkle Root

Details about the sync committee commitment were omitted on purpose in the previous explanations for clarity. In fact, each time a committee's SSZ merkle root is saved into the state, a mapping from that root to a Poseidon merkle root is also created. This mapping is supported by a proof that the updater must provide in when calling `updateSyncCommittee()`. This additional proof is built from a circuit that takes in the 512 pubkeys of the committee and calculates both SSZ and Poseidon merkle roots. The point of this extra mapping is to save proof generation time so that when proving validity for each block that we want to update, the prover only needs to calculate a zk friendly Poseidon root of the pubkeys. The cost of this mapping is that there is an additional proof needed when updating the committee. So in short, heavier proof every 27 hours and lighter proof for every block.
