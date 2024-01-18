// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

library Memory {
    /**
     * Copies a part of bytes.
     * @param source original bytes
     * @param from the first index to be copied, data included
     * @param to the last index(to be copied) + 1, data excluded.
     */
    function range(bytes memory source, uint256 from, uint256 to) internal pure returns (bytes memory) {
        if (from >= to) {
            return "";
        }

        require(from < source.length && from >= 0, "Memory: from out of bounds");
        require(to <= source.length && to >= 0, "Memory: to out of bounds");

        bytes memory result = new bytes(to - from);

        uint256 srcPtr;
        assembly {
            srcPtr := add(source, 0x20)
        }

        srcPtr += from;

        uint256 destPtr;
        assembly {
            destPtr := add(result, 0x20)
        }

        copy(destPtr, srcPtr, to - from);

        return result;
    }

    /**
     * Copies a piece of memory to another location
     * @notice From: https://github.com/Arachnid/solidity-stringutils/blob/master/src/strings.sol
     * @param _destPtr Destination location pointer
     * @param _srcPtr Source location pointer
     * @param _length Length of memory(in bytes) to be copied.
     */
    function copy(uint256 _destPtr, uint256 _srcPtr, uint256 _length) internal pure {
        uint256 destPtr = _destPtr;
        uint256 srcPtr = _srcPtr;
        uint256 remainingLength = _length;

        for (; remainingLength >= 32; remainingLength -= 32) {
            assembly {
                mstore(destPtr, mload(srcPtr))
            }
            destPtr += 32;
            srcPtr += 32;
        }

        uint256 mask;
        unchecked {
            mask = 256 ** (32 - remainingLength) - 1;
        }

        assembly {
            let srcPart := and(mload(srcPtr), not(mask))
            let destPart := and(mload(destPtr), mask)
            mstore(destPtr, or(destPart, srcPart))
        }
    }
}
