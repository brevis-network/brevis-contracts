// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

interface ITxVerifier {
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
        uint32 blkNum;
        bytes32 blkHash;
        uint64 blkTime;
    }

    // reverts if not verified
    // - txRaw: signed dynamic fee tx rlp encode data
    // - proofData: tx proof data
    // - auxiBlkVerifyInfo: auxiliary info for blk verify in chunk sync contract
    function verifyTx(
        bytes calldata txRaw,
        bytes calldata proofData,
        bytes calldata auxiBlkVerifyInfo
    ) external view returns (TxInfo memory txInfo);

    // verifyTx and emit event
    function verifyTxAndLog(
        bytes calldata txRaw,
        bytes calldata proofData,
        bytes calldata auxiBlkVerifyInfo
    ) external returns (TxInfo memory info);
}
