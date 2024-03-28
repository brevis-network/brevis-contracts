// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "./BrevisProof.sol";
import "../lib/Lib.sol";

contract BrevisAggProof is BrevisProof {

    constructor(ISMT _smtContract) BrevisProof(_smtContract) {}
    mapping(bytes32 => bool) public merkleRoots;

    function mustValidateRequest(Brevis.ProofData calldata _proofData, bytes32 _merkleRoot, bytes32[] calldata _merkleProof) external view {
        require(merkleRoots[_merkleRoot], "invlida merkle root");

        // TODO: 
        // 1. hash _proofData as h
        // 2. validate h is in _merkleRoot with _merklePRoof
    }
}