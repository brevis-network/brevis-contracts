// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./BrevisProof.sol";
import "../lib/Lib.sol";
import "../../interfaces/ISMT.sol";
import "../../verifiers/interfaces/IZkpVerifier.sol";

contract BrevisAggProof is Ownable {
    ISMT public smtContract;

    constructor(ISMT _smtContract) {
        smtContract = _smtContract;
    }

    mapping(bytes32 => bool) public merkleRoots;
    mapping(uint64 => IZkpVerifier) public aggProofVerifierAddress;
    event SmtContractUpdated(ISMT smtContract);
    event AggProofVerifierAddressesUpdated(uint64[] chainIds, IZkpVerifier[] newAddresses);

    uint32 constant PUBLIC_BYTES_START_IDX = 12 * 32; // the first 12 32bytes are groth16 proof (A/B/C/Commitment/CommitmentPOK)
    uint8 constant TREE_DEPTH = 5;

    function mustValidateRequest(
        uint64 _chainId,
        Brevis.ProofData calldata _proofData,
        bytes32 _merkleRoot,
        bytes32[TREE_DEPTH] calldata _merkleProof,
        bool _isLeftSide
    ) external view {
        require(merkleRoots[_merkleRoot], "merkle root not exists");
        require(smtContract.isSmtRootValid(_chainId, _proofData.smtRoot), "invalid smt root");

        bytes32 proofDataHash = keccak256(
            abi.encodePacked(
                _proofData.commitHash,
                _proofData.smtRoot,
                _proofData.vkHash,
                _proofData.appCommitHash,
                _proofData.appVkHash
            )
        );
        bytes32 root = proofDataHash;
        for (uint8 depth = 0; depth < TREE_DEPTH; depth++) {
            if (_isLeftSide) {
                root = keccak256(abi.encodePacked(root, _merkleProof[depth]));
            } else {
                root = keccak256(abi.encodePacked(_merkleProof[depth], root));
            }
        }
        require(_merkleRoot == root, "invalid data");
    }

    function submitAggProof(uint64 _chainId, bytes calldata _proofWithPubInputs) external {
        IZkpVerifier verifier = aggProofVerifierAddress[_chainId];
        require(address(verifier) != address(0), "chain agg proof verifier not set");
        require(verifier.verifyRaw(_proofWithPubInputs), "proof not valid");

        bytes32 root = unpackMerkleRoot(_proofWithPubInputs);
        merkleRoots[root] = true;
    }

    function unpackMerkleRoot(bytes calldata _proofWithPubInputs) internal pure returns (bytes32) {
        return bytes32(_proofWithPubInputs[PUBLIC_BYTES_START_IDX:PUBLIC_BYTES_START_IDX + 32]);
    }

    function updateSmtContract(ISMT _smtContract) public onlyOwner {
        smtContract = _smtContract;
        emit SmtContractUpdated(smtContract);
    }

    function updateAggProofVerifierAddresses(
        uint64[] calldata _chainIds,
        IZkpVerifier[] calldata _verifierAddresses
    ) public onlyOwner {
        require(_chainIds.length == _verifierAddresses.length, "length not match");
        for (uint256 i = 0; i < _chainIds.length; i++) {
            aggProofVerifierAddress[_chainIds[i]] = _verifierAddresses[i];
        }
        emit AggProofVerifierAddressesUpdated(_chainIds, _verifierAddresses);
    }
}
