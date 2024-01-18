// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../light-client/interfaces/IAnchorBlocks.sol";

contract MockAnchorBlocks is IAnchorBlocks {
    mapping(uint256 => bytes32) public blocks;

    function update(uint256 blockNum, bytes32 blockHash) external {
        blocks[blockNum] = blockHash;
    }
}
