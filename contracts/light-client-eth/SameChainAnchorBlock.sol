// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./interfaces/IAnchorBlocks.sol";

contract SameChainAnchorBlocks is IAnchorBlocks {
   function blocks(uint256 blockNum) external view returns (bytes32) {
        return blockhash(blockNum);
    }
}
