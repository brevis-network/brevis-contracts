// SPDX-License-Identifier: MIT
pragma solidity ^0.8.2;

import "../proto/TendermintHelper.sol";
import {SimpleValidator, Validator} from "../proto/TendermintLight.sol";

library MerkleTree {
    /**
     * @dev returns empty hash
     */
    function emptyHash() internal pure returns (bytes32) {
        return sha256(abi.encode());
    }

    /**
     * @dev returns tmhash(0x00 || leaf)
     *
     */
    function leafHash(bytes memory leaf) internal pure returns (bytes32) {
        uint8 leafPrefix = 0x00;
        return sha256(abi.encodePacked(leafPrefix, leaf));
    }

    /**
     * @dev returns tmhash(0x01 || left || right)
     */
    function innerHash(bytes32 leaf, bytes32 right) internal pure returns (bytes32) {
        uint8 innerPrefix = 0x01;
        return sha256(abi.encodePacked(innerPrefix, leaf, right));
    }

    /**
     * @dev returns the largest power of 2 less than length
     *
     * TODO: This function can be optimized with bit shifting approach:
     * https://www.baeldung.com/java-largest-power-of-2-less-than-number
     */
    function getSplitPoint(uint256 input) internal pure returns (uint256) {
        require(input > 1, "MerkleTree: invalid input");

        uint256 result = 1;
        for (uint256 i = input - 1; i > 1; i--) {
            if ((i & (i - 1)) == 0) {
                result = i;
                break;
            }
        }
        return result;
    }

    /**
     * @dev computes a Merkle tree where the leaves are validators, in the provided order
     * Follows RFC-6962
     */
    function merkleRootHash(
        Validator.Data[] memory validators,
        uint256 start,
        uint256 total
    ) internal pure returns (bytes32) {
        if (total == 0) {
            return emptyHash();
        } else if (total == 1) {
            bytes memory encodedValidator = SimpleValidator.encode(
                TendermintHelper.toSimpleValidator(validators[start])
            );
            return leafHash(encodedValidator);
        } else {
            uint256 k = getSplitPoint(total);
            bytes32 left = merkleRootHash(validators, start, k); // validators[:k]
            bytes32 right = merkleRootHash(validators, start + k, total - k); // validators[k:]
            return innerHash(left, right);
        }
    }

    /**
     * @dev computes a Merkle tree where the leaves are the byte slice in the provided order
     * Follows RFC-6962
     */
    function merkleRootHash(bytes[14] memory validators, uint256 start, uint256 total) internal pure returns (bytes32) {
        if (total == 0) {
            return emptyHash();
        } else if (total == 1) {
            return leafHash(validators[start]);
        } else {
            uint256 k = getSplitPoint(total);
            bytes32 left = merkleRootHash(validators, start, k); // validators[:k]
            bytes32 right = merkleRootHash(validators, start + k, total - k); // validators[k:]
            return innerHash(left, right);
        }
    }
}
