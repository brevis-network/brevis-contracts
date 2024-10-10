// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

interface IAnchorBlocks {
    function blockHashs(uint256 blockNum) external view returns (bytes32);
    function blockNums(bytes32 blockHash) external view returns (uint256);
}
