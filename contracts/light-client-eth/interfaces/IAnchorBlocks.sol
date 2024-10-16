// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

interface IAnchorBlocks {
    function blocks(uint256 blockNum) external view returns (bytes32);
}
