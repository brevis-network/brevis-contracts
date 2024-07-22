// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./interfaces/IAnchorBlocks.sol";
import "../safeguard/Ownable.sol";

contract SameChainAnchorBlocks is IAnchorBlocks, Ownable {
    function blocks(uint256 blockNum) external view returns (bytes32) {
        return blockhash(blockNum);
    }
}
