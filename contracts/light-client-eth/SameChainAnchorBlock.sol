// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./interfaces/IAnchorBlocks.sol";
import "../safeguard/BrevisAccess.sol";

contract SameChainAnchorBlocks is IAnchorBlocks, BrevisAccess {
    // execution block number => execution block hash
    mapping(uint256 => bytes32) public blockHashs;
    // execution block hash => execution block number
    mapping(bytes32 => uint256) public blockNums;

    function processUpdate(uint256 blockNumber) external onlyActiveProver {
        require(blockHashs[blockNumber] == bytes32(0), "block hash already exists");
        bytes32 blockHash = blockhash(blockNumber);
        require(blockHash!= bytes32(0), "block hash not found");
        blockHashs[blockNumber] = blockHash;
        blockNums[blockHash] = blockNumber;
    }
}
