// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../../lib/RLPReader.sol";

library Brevis {
    // retrieved from proofData, to align the logs with circuit...
    struct ProofData {
        bytes32 commitHash;
        bytes32 appCommitHash; // zk-program computing circuit commit hash
        bytes32 appVkHash; // zk-program computing circuit Verify Key hash
        bytes32 smtRoot;
        bytes32 dummyInputCommitment; // zk-program computing circuit dummy input commitment
    }

    struct ProofAppData {
        bytes32 appCommitHash;
        bytes32 appVkHash;
    }
}

library Tx {
    using RLPReader for bytes;
    using RLPReader for uint;
    using RLPReader for RLPReader.RLPItem;

    struct TxInfo {
        uint64 chainId;
        uint64 nonce;
        uint256 gasTipCap;
        uint256 gasFeeCap;
        uint256 gas;
        address to;
        uint256 value;
        bytes data;
        address from; // calculate from V R S
    }

    // support DynamicFeeTxType for now
    function decodeTx(bytes calldata txRaw) public pure returns (TxInfo memory info) {
        uint8 txType = uint8(txRaw[0]);
        require(txType == 2, "not a DynamicFeeTxType");

        bytes memory rlpData = txRaw[1:];
        RLPReader.RLPItem[] memory values = rlpData.toRlpItem().toList();
        info.chainId = uint64(values[0].toUint());
        info.nonce = uint64(values[1].toUint());
        info.gasTipCap = values[2].toUint();
        info.gasFeeCap = values[3].toUint();
        info.gas = values[4].toUint();
        info.to = values[5].toAddress();
        info.value = values[6].toUint();
        info.data = values[7].toBytes();

        (uint8 v, bytes32 r, bytes32 s) = (
            uint8(values[9].toUint()),
            bytes32(values[10].toBytes()),
            bytes32(values[11].toBytes())
        );
        // remove r,s,v and adjust length field
        bytes memory unsignedTxRaw;
        uint16 unsignedTxRawDataLength;
        uint8 prefix = uint8(txRaw[1]);
        uint8 lenBytes = prefix - 0xf7; // assume lenBytes won't larger than 2, means the tx rlp data size won't exceed 2^16
        if (lenBytes == 1) {
            unsignedTxRawDataLength = uint8(bytes1(txRaw[2:3])) - 67; //67 is the bytes of r,s,v
        } else {
            unsignedTxRawDataLength = uint16(bytes2(txRaw[2:2 + lenBytes])) - 67;
        }
        if (unsignedTxRawDataLength <= 55) {
            unsignedTxRaw = abi.encodePacked(txRaw[:2], txRaw[3:txRaw.length - 67]);
            unsignedTxRaw[1] = bytes1(0xc0 + uint8(unsignedTxRawDataLength));
        } else {
            if (unsignedTxRawDataLength <= 255) {
                unsignedTxRaw = abi.encodePacked(
                    txRaw[0],
                    bytes1(0xf8),
                    bytes1(uint8(unsignedTxRawDataLength)),
                    txRaw[2 + lenBytes:txRaw.length - 67]
                );
            } else {
                unsignedTxRaw = abi.encodePacked(
                    txRaw[0],
                    bytes1(0xf9),
                    bytes2(unsignedTxRawDataLength),
                    txRaw[2 + lenBytes:txRaw.length - 67]
                );
            }
        }
        info.from = recover(keccak256(unsignedTxRaw), r, s, v);
    }

    function recover(bytes32 message, bytes32 r, bytes32 s, uint8 v) internal pure returns (address) {
        if (v < 27) {
            v += 27;
        }
        return ecrecover(message, v, r, s);
    }
}
