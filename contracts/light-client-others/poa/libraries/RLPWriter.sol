// SPDX-License-Identifier: MIT
// Inspired: https://github.com/ethereum-optimism/optimism/blob/v1.0.9/packages/contracts-bedrock/contracts/libraries/rlp/RLPWriter.sol

pragma solidity ^0.8.18;

import "./Memory.sol";

// import "hardhat/console.sol";

library RLPWriter {
    /**
     * RLP encodes bool
     * @param _input The bool value to be encoded
     * @return RLP encoded bool value in bytes
     */
    function writeBool(bool _input) internal pure returns (bytes memory) {
        bytes memory encoded = new bytes(1);
        encoded[0] = (_input ? bytes1(0x01) : bytes1(0x80));
        return encoded;
    }

    /**
     * RLP encodes bytes
     * @param _input The byte string to be encoded
     * @return RLP encoded string in bytes
     */
    function writeBytes(bytes memory _input) internal pure returns (bytes memory) {
        bytes memory encoded;

        // input ∈ [0x00, 0x7f]
        if (_input.length == 1 && uint8(_input[0]) < 128) {
            encoded = _input;
        } else {
            // Offset 0x80
            encoded = abi.encodePacked(_writeLength(_input.length, 128), _input);
        }

        return encoded;
    }

    /**
     * RLP encodes a list of RLP encoded items
     * @param _input The list of RLP encoded items
     * @return RLP encoded list of items in bytes
     */
    function writeList(bytes[] memory _input) internal pure returns (bytes memory) {
        bytes memory flatten = _flatten(_input);
        // offset 0xc0
        return abi.encodePacked(_writeLength(flatten.length, 192), flatten);
    }

    /**
     * RLP encodes a string
     * @param _input The string to be encoded
     * @return RLP encoded string in bytes
     */
    function writeString(string memory _input) internal pure returns (bytes memory) {
        return writeBytes(bytes(_input));
    }

    /**
     * RLP encodes an address
     * @param _input The address to be encoded
     * @return RLP encoded address in bytes
     */
    function writeAddress(address _input) internal pure returns (bytes memory) {
        return writeBytes(abi.encodePacked(_input));
    }

    /**
     * RLP encodes a uint256 value
     * @param _input The uint256 to be encoded
     * @return RLP encoded uint256 in bytes
     */
    function writeUint(uint256 _input) internal pure returns (bytes memory) {
        return writeBytes(_toBinary(_input));
    }

    /**
     * Encode offset + length as first byte, followed by length in hex display if needed
     * _offset: 0x80 for single item, 0xc0/192 for list
     * If length is greater than 55, offset should add 55. 0xb7 for single item, 0xf7 for list
     * @param _length The length of single item or list
     * @param _offset Type indicator
     * @return RLP encoded bytes
     */
    function _writeLength(uint256 _length, uint256 _offset) private pure returns (bytes memory) {
        bytes memory encoded;

        if (_length < 56) {
            encoded = new bytes(1);
            encoded[0] = bytes1(uint8(_offset) + uint8(_length));
        } else {
            uint256 hexLengthForInputLength = 0;
            uint256 index = 1;
            while (_length / index != 0) {
                index *= 256;
                hexLengthForInputLength++;
            }
            encoded = new bytes(hexLengthForInputLength + 1);

            // 0x80 + 55 = 0xb7
            // 0xc0 + 55 = 0xf7
            encoded[0] = bytes1(uint8(_offset) + 55 + uint8(hexLengthForInputLength));
            for (index = 1; index <= hexLengthForInputLength; index++) {
                encoded[index] = bytes1(uint8((_length / (256 ** (hexLengthForInputLength - index))) % 256));
            }
        }

        return encoded;
    }

    function toBinary(uint256 _input) internal pure returns (bytes memory) {
        return _toBinary(_input);
    }

    /**
     * Encode integer into big endian without leading zeros
     * @param _input The integer to be encoded
     * @return RLP encoded bytes
     */
    function _toBinary(uint256 _input) private pure returns (bytes memory) {
        // if input value is 0, return 0x00
        if (_input == 0) {
            bytes memory zeroResult = new bytes(1);
            zeroResult[0] = 0;
            return zeroResult;
        }

        bytes memory data = abi.encodePacked(_input);

        uint8 index = 0;
        for (; index < 32; ) {
            if (data[index] != 0) {
                break;
            }

            unchecked {
                ++index;
            }
        }

        bytes memory result = new bytes(32 - index);
        uint256 resultPtr;
        assembly {
            resultPtr := add(result, 0x20)
        }

        uint256 dataPtr;
        assembly {
            dataPtr := add(data, 0x20)
        }

        Memory.copy(resultPtr, dataPtr + index, 32 - index);

        return result;
    }

    /**
     * Flattens a list of byte strings into one byte string.
     * @notice From: https://github.com/sammayo/solidity-rlp-encoder/blob/master/RLPEncode.sol.
     * @param _list List of byte strings to flatten.
     * @return The flattened byte string.
     */
    function _flatten(bytes[] memory _list) private pure returns (bytes memory) {
        if (_list.length == 0) {
            return new bytes(0);
        }

        uint256 length = 0;
        uint256 index = 0;

        for (; index < _list.length; ) {
            length += _list[index].length;
            unchecked {
                ++index;
            }
        }

        bytes memory flattened = new bytes(length);
        uint256 flattenedPtr;
        assembly {
            flattenedPtr := add(flattened, 0x20)
        }

        for (index = 0; index < _list.length; ) {
            bytes memory item = _list[index];
            uint256 itemPtr;
            assembly {
                itemPtr := add(item, 0x20)
            }

            Memory.copy(flattenedPtr, itemPtr, item.length);
            flattenedPtr += _list[index].length;

            unchecked {
                ++index;
            }
        }

        return flattened;
    }
}
