// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

interface IVerifier {
    struct Proof {
        uint256[2] a;
        uint256[2][2] b;
        uint256[2] c;
        uint256[2] commitment;
    }

    function verifyProof(
        uint256[2] memory a,
        uint256[2][2] memory b,
        uint256[2] memory c,
        uint256[2] memory commit,
        uint256[10] calldata input
    ) external view returns (bool r);
}
