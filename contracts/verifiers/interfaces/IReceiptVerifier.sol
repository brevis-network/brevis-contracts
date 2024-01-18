// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

interface IReceiptVerifier {
    struct ReceiptInfo {
        bool success;
        uint64 chainId;
        bytes32 blkHash;
        uint32 blkNum;
        uint64 blkTime;
        LogInfo[] logs;
        // TODO: add transaction index
    }

    struct LogInfo {
        address addr;
        bytes32[] topics;
        bytes data;
    }

    // reverts if not verified
    // - receiptRaw: signed dynamic fee receipt rlp encode data
    // - proofData: receipt proof data
    // - auxiBlkVerifyInfo: auxiliary info for blk verify in chunk sync contract
    function verifyReceipt(
        bytes calldata receiptRaw,
        bytes calldata proofData,
        bytes calldata auxiBlkVerifyInfo
    ) external view returns (ReceiptInfo memory receiptInfo);

    // verifyReceipt and emit event
    function verifyReceiptAndLog(
        bytes calldata receiptRaw,
        bytes calldata proofData,
        bytes calldata auxiBlkVerifyInfo
    ) external returns (ReceiptInfo memory receiptInfo);
}
