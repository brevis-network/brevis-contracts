// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/Ownable.sol";

import "./interfaces/IAnchorBlocks.sol";

contract SameChainAnchorBlocks is IAnchorBlocks, Ownable {
    function blocks(uint256 blockNum) external view returns (bytes32) {
        return blockhash(blockNum);
    }
}
