// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

import "./RLPReader.sol";

library MerkleProofTree {
    using RLPReader for RLPReader.RLPItem;
    using RLPReader for bytes;

    function _gnaw(uint256 index, bytes32 key) private pure returns (uint256 gnaw) {
        assembly {
            gnaw := shr(mul(sub(63, index), 4), key)
        }
        return gnaw % 16;
    }

    function _pathLength(bytes memory path) private pure returns (uint256, bool) {
        uint256 gnaw = uint256(uint8(path[0])) / 16;
        return ((path.length - 1) * 2 + (gnaw % 2), gnaw > 1);
    }

    function read(bytes32 key, bytes[] memory proof) internal pure returns (bytes memory result) {
        bytes32 root;
        bytes memory node = proof[0];

        uint256 index = 0;
        uint256 pathLength = 0;

        while (true) {
            RLPReader.RLPItem[] memory items = node.toRlpItem().toList();
            if (items.length == 17) {
                uint256 gnaw = _gnaw(pathLength++, key);
                root = bytes32(items[gnaw].toUint());
            } else {
                require(items.length == 2, "MessageBridge: Iinvalid RLP list length");
                (uint256 nodePathLength, bool isLeaf) = _pathLength(items[0].toBytes());
                pathLength += nodePathLength;
                if (isLeaf) {
                    return items[1].toBytes();
                } else {
                    root = bytes32(items[1].toUint());
                }
            }

            node = proof[++index];
            require(root == keccak256(node), "MessageBridge: node hash mismatched");
        }
    }

    function restoreMerkleRoot(bytes32 leaf, uint256 index, bytes32[] memory proof) internal pure returns (bytes32) {
        bytes32 value = leaf;
        for (uint256 i = 0; i < proof.length; i++) {
            if ((index / (2 ** i)) % 2 == 1) {
                value = sha256(bytes.concat(proof[i], value));
            } else {
                value = sha256(bytes.concat(value, proof[i]));
            }
        }
        return value;
    }
}
