// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../libraries/MerkleProofTree.sol";

contract MockMerkleProofTree {
    function mockRead(bytes32 key, bytes[] memory proof) external pure returns (bytes memory result) {
        return MerkleProofTree.read(key, proof);
    }
}
