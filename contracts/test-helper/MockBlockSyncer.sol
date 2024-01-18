// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../chunk-sync/interfaces/IBlockChunks.sol";

contract MockBlockChunks is IBlockChunks {
    function isBlockHashValid(BlockHashWitness calldata) external pure returns (bool) {
        return true;
    }

    function historicalRoots(uint64 chainId, uint32 startBlockNumber) external view returns (bytes32) {
        // nothing to do
    }

    function updateRecent(uint64 chainId, bytes calldata proofData) external {
        // nothing to do
    }

    function updateOld(uint64 chainId, bytes32 nextRoot, uint32 nextNumFinal, bytes calldata proofData) external {
        // nothing to do
    }
}
