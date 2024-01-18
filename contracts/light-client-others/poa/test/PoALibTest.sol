// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../libraries/RLPWriter.sol";
import "../libraries/Memory.sol";
import "../libraries/ECDSA.sol";

contract PoALibTest {
    function mockRange(bytes memory source, uint256 from, uint256 to) public pure returns (bytes memory) {
        return Memory.range(source, from, to);
    }

    function mockCopy(bytes memory _source, uint256 _length) public pure returns (bytes memory) {
        bytes memory dest = new bytes(_length);
        uint256 destPtr;
        assembly {
            destPtr := add(dest, 0x20)
        }

        uint256 srcPtr;
        assembly {
            srcPtr := add(_source, 0x20)
        }

        Memory.copy(destPtr, srcPtr, _length);
        return dest;
    }

    function mockWriteUint(uint256 _input) public pure returns (bytes memory) {
        return RLPWriter.writeUint(_input);
    }

    function mockWriteAddress(address _input) public pure returns (bytes memory) {
        return RLPWriter.writeAddress(_input);
    }

    function mockWriteRLPList(bytes[] memory _input) public pure returns (bytes memory) {
        return RLPWriter.writeList(_input);
    }

    function mockWriteBool(bool _input) public pure returns (bytes memory) {
        return RLPWriter.writeBool(_input);
    }

    function mockWriteString(string calldata _input) public pure returns (bytes memory) {
        return RLPWriter.writeString(_input);
    }

    function mockWriteBytes(bytes memory _input) public pure returns (bytes memory) {
        return RLPWriter.writeBytes(_input);
    }

    function mockToBinary(uint256 _input) public pure returns (bytes memory) {
        return RLPWriter.toBinary(_input);
    }

    function mockUint256MaxToBinary() public pure returns (bytes memory) {
        return RLPWriter.toBinary(115792089237316195423570985008687907853269984665640564039457584007913129639935);
    }

    function mockRecoverAddress(bytes32 message, bytes memory signature) public pure returns (address) {
        bytes32 h = ECDSA.toEthSignedMessageHash(message);
        return ECDSA.recover(h, signature);
    }
}
