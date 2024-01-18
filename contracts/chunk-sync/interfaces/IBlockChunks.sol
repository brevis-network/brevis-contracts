// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

interface IBlockChunks {
    // historicalRoots(chainId, startBlockNumber) is 0 unless (startBlockNumber % 128 == 0)
    // historicalRoots(chainId, startBlockNumber) holds the hash of
    //   prevHash || root || numFinal
    // where
    // - prevHash is the parent hash of block startBlockNumber
    // - root is the partial Merkle root of blockhashes of block numbers
    //   [startBlockNumber, startBlockNumber + 128)
    //   where unconfirmed block hashes are 0's
    // - numFinal is the number of confirmed consecutive roots in [startBlockNumber, startBlockNumber + 128)
    function historicalRoots(uint64 chainId, uint32 startBlockNumber) external view returns (bytes32);

    event UpdateEvent(uint64 chainId, uint32 startBlockNumber, bytes32 prevHash, bytes32 root, uint32 numFinal);

    struct BlockHashWitness {
        uint64 chainId;
        uint32 blkNum;
        bytes32 claimedBlkHash;
        bytes32 prevHash;
        uint32 numFinal;
        bytes32[7] merkleProof;
    }

    // update blocks in the "backward" direction, anchoring on a "recent" end blockhash from anchor contract
    // * startBlockNumber must be a multiple of 128
    // * for now always endBlockNumber = startBlockNumber + 127 (full update on every 128 blocks chunk)
    function updateRecent(uint64 chainId, bytes calldata proofData) external;

    // update older blocks in "backwards" direction, anchoring on more recent trusted blockhash
    // must be batch of 128 blocks
    function updateOld(uint64 chainId, bytes32 nextRoot, uint32 nextNumFinal, bytes calldata proofData) external;

    function isBlockHashValid(BlockHashWitness calldata witness) external view returns (bool);
}
