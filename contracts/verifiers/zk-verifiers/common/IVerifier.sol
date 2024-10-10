// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

interface IVerifier {
    function verifyProof(
        uint256[8] calldata proof,
        uint256[2] calldata commit,
        uint256[2] calldata knowledgeProof,
        uint256[8] calldata input
    ) external view returns (bool r);
}
