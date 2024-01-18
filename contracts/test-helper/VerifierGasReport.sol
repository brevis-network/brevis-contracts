// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../verifiers/interfaces/ITxVerifier.sol";
import "../verifiers/interfaces/IReceiptVerifier.sol";
import "../verifiers/interfaces/ISlotValueVerifier.sol";

// used for test gas consumptions of view functions

interface IProofVerifier {
    function verifyProof(
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c,
        uint256[2] memory commit,
        uint256[6] memory input
    ) external view returns (bool r);

    function verifyProof(
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c,
        uint256[2] memory commit,
        uint256[9] memory input
    ) external view returns (bool r);

    function verifyProof(
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c,
        uint256[2] memory commit,
        uint256[8] memory input
    ) external view returns (bool r);

    function verifyRaw(bytes calldata proofData) external view returns (bool);
}

contract VerifierGasReport {
    address public verifier;

    event ProofVerified(bool success);
    event TxVerified(address from, bytes32 blkHash);

    constructor(address _verifier) {
        verifier = _verifier;
    }

    function transaction13VerifyProof(
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c,
        uint256[2] memory commit,
        uint256[6] memory input
    ) external {
        verifyProof(a, b, c, commit, input);
    }

    function transaction37VerifyProof(
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c,
        uint256[2] memory commit,
        uint256[6] memory input
    ) external {
        verifyProof(a, b, c, commit, input);
    }

    function transactionVerifyProof(
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c,
        uint256[2] memory commit,
        uint256[6] memory input
    ) external {
        verifyProof(a, b, c, commit, input);
    }

    function receiptVerifyProof(
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c,
        uint256[2] memory commit,
        uint256[6] memory input
    ) external {
        verifyProof(a, b, c, commit, input);
    }

    function ethStorageVerifyProof(
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c,
        uint256[2] memory commit,
        uint256[9] memory input
    ) external {
        verifyProof(a, b, c, commit, input);
    }

    function ethChunkOf4VerifyProof(
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c,
        uint256[2] memory commit,
        uint256[8] memory input
    ) external {
        verifyProof(a, b, c, commit, input);
    }

    function ethChunkOf128VerifyProof(
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c,
        uint256[2] memory commit,
        uint256[8] memory input
    ) external {
        verifyProof(a, b, c, commit, input);
    }

    function verifyTx(bytes calldata txRaw, bytes calldata proofData, bytes calldata auxiBlkVerifyInfo) external {
        ITxVerifier(verifier).verifyTx(txRaw, proofData, auxiBlkVerifyInfo);
        emit ProofVerified(true);
    }

    function verifyReceipt(
        bytes calldata receiptRaw,
        bytes calldata proofData,
        bytes calldata auxiBlkVerifyInfo
    ) external {
        IReceiptVerifier(verifier).verifyReceipt(receiptRaw, proofData, auxiBlkVerifyInfo);
        emit ProofVerified(true);
    }

    function verifyRaw(bytes calldata proofData) external {
        bool success = IProofVerifier(verifier).verifyRaw(proofData);
        emit ProofVerified(success);
    }

    function verifySlotValue(uint64 chainId, bytes calldata proofData, bytes calldata blkVerifyInfo) external {
        ISlotValueVerifier(verifier).verifySlotValue(chainId, proofData, blkVerifyInfo);
        emit ProofVerified(true);
    }

    function verifyProof(
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c,
        uint256[2] memory commit,
        uint256[6] memory input
    ) private {
        bool success = IProofVerifier(verifier).verifyProof(a, b, c, commit, input);
        emit ProofVerified(success);
    }

    function verifyProof(
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c,
        uint256[2] memory commit,
        uint256[9] memory input
    ) private {
        bool success = IProofVerifier(verifier).verifyProof(a, b, c, commit, input);
        emit ProofVerified(success);
    }

    function verifyProof(
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c,
        uint256[2] memory commit,
        uint256[8] memory input
    ) private {
        bool success = IProofVerifier(verifier).verifyProof(a, b, c, commit, input);
        emit ProofVerified(success);
    }

    function setVerifier(address _verifier) external {
        verifier = _verifier;
    }
}
