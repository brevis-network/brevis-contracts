// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../light-client-eth/interfaces/IAnchorBlocks.sol";

contract MockAnchorBlocks is IAnchorBlocks {
    mapping(uint256 => bytes32) public blockHashs;
    mapping(bytes32 => uint256) public blockNums;


    function update(uint256 blockNum, bytes32 blockHash) external {
        blockHashs[blockNum] = blockHash;
        blockNums[blockHash] = blockNum;
    }
}
