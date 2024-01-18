// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "./zk-verifiers/BlsSigVerifier.sol";
import "./zk-verifiers/CommitteeRootMappingVerifier.sol";
import "./interfaces/IBeaconVerifier.sol";

contract BeaconVerifier is IBeaconVerifier, BlsSigVerifier, CommitteeRootMappingVerifier {
    function verifySignatureProof(
        bytes32 signingRoot,
        bytes32 syncCommitteePoseidonRoot,
        uint256 participation,
        uint256 commitment,
        Proof memory p
    ) public view returns (bool) {
        uint256[35] memory input;
        uint256 root = uint256(signingRoot);
        // slice the signing root into 32 individual bytes and assign them in order to the first 32 slots of input[]
        for (uint256 i = 0; i < 32; i++) {
            input[(32 - 1 - i)] = root % 256;
            root = root / 256;
        }
        input[32] = participation;
        input[33] = uint256(syncCommitteePoseidonRoot);
        input[34] = commitment;
        return verifyBlsSigProof(p.a, p.b, p.c, p.commitment, input);
    }

    function verifySyncCommitteeRootMappingProof(
        bytes32 sszRoot,
        bytes32 poseidonRoot,
        Proof memory p
    ) public view returns (bool) {
        uint256[33] memory input;
        uint256 root = uint256(sszRoot);
        for (uint256 i = 0; i < 32; i++) {
            input[(32 - 1 - i)] = root % 256;
            root = root / 256;
        }
        input[32] = uint256(poseidonRoot);
        return verifyCommitteeRootMappingProof(p.a, p.b, p.c, input);
    }
}
