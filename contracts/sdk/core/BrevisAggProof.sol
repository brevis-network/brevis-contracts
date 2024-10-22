// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../lib/Lib.sol";
import "../../interfaces/ISMT.sol";
import "../../verifiers/interfaces/IZkpVerifier.sol";
import "../../safeguard/BrevisAccess.sol";

contract BrevisAggProof is BrevisAccess {
    uint32 constant PUBLIC_BYTES_START_IDX = 12 * 32; // the first 12 32bytes are groth16 proof (A/B/C/Commitment/CommitmentPOK)
    uint8 constant TREE_DEPTH = 4;
    uint256 constant LEAF_NODES_LEN = 2 ** TREE_DEPTH;

    ISMT public smtContract;

    mapping(bytes32 => bool) public merkleRoots;
    mapping(uint64 => IZkpVerifier) public aggProofVerifierAddress;
    mapping(uint64 => bytes32) public dummyInputCommitments;
    event SmtContractUpdated(address smtContract);
    event AggProofVerifierAddressesUpdated(uint64[] chainIds, IZkpVerifier[] newAddresses);
    event DummyInputCommitmentsUpdated(uint64[] chainIds, bytes32[] updatedDummyInputCommitments);

    constructor(ISMT _smtContract) {
        smtContract = _smtContract;
    }

    /*********************************
     * External and Public Functions *
     *********************************/

    function submitAggProof(
        uint64 _chainId,
        bytes32[] calldata _proofIds,
        bytes calldata _proofWithPubInputs
    ) external onlyActiveProver {
        IZkpVerifier verifier = aggProofVerifierAddress[_chainId];
        require(address(verifier) != address(0), "chain agg proof verifier not set");
        require(verifier.verifyRaw(_proofWithPubInputs), "proof not valid");

        (bytes32 root, bytes32 proofIdsCommit) = unpack(_proofWithPubInputs);

        uint dataLen = _proofIds.length;
        bytes32[LEAF_NODES_LEN] memory rIds;
        for (uint i = 0; i < dataLen; i++) {
            rIds[i] = _proofIds[i];
        }
        // note, to align with circuit, rIds[dataLen] to rIds[LEAF_NODES_LEN - 1] filled with last real one
        if (dataLen < LEAF_NODES_LEN) {
            for (uint i = dataLen; i < LEAF_NODES_LEN; i++) {
                rIds[i] = rIds[dataLen - 1];
            }
        }
        require(keccak256(abi.encodePacked(rIds)) == proofIdsCommit, "proofIds not right");
        merkleRoots[root] = true;
    }

    // validate all leaf nodes in the agg proof data
    function validateAggProofData(uint64 _chainId, Brevis.ProofData[] calldata _proofDataArray) external view {
        uint dataLen = _proofDataArray.length;
        require(dataLen <= LEAF_NODES_LEN, "size exceeds");
        bytes32[2 * LEAF_NODES_LEN - 1] memory hashes;
        for (uint i = 0; i < dataLen; i++) {
            require(smtContract.isSmtRootValid(_chainId, _proofDataArray[i].smtRoot), "invalid smt root");
            require(dummyInputCommitments[_chainId] == _proofDataArray[i].dummyInputCommitment, "invalid dummy input");
            hashes[i] = keccak256(
                abi.encodePacked(
                    _proofDataArray[i].commitHash,
                    _proofDataArray[i].smtRoot,
                    _proofDataArray[i].appCommitHash,
                    _proofDataArray[i].appVkHash,
                    _proofDataArray[i].dummyInputCommitment
                )
            );
        }
        // note, hashes[dataLen] to hashes[LEAF_NODES_LEN - 1] filled with last real one
        if (dataLen < LEAF_NODES_LEN) {
            for (uint i = dataLen; i < LEAF_NODES_LEN; i++) {
                hashes[i] = hashes[dataLen - 1];
            }
        }

        uint shift = 0;
        uint counter = LEAF_NODES_LEN;
        while (counter > 0) {
            for (uint i = 0; i < counter - 1; i += 2) {
                hashes[shift + counter + i / 2] = keccak256(abi.encodePacked(hashes[shift + i], hashes[shift + i + 1]));
            }
            shift += counter;
            counter /= 2;
        }

        require(merkleRoots[hashes[hashes.length - 1]], "merkle root not exists");
    }

    // validate a single leaf node in the agg proof data
    function validateAggProofData(
        uint64 _chainId,
        Brevis.ProofData calldata _proofData,
        bytes32 _merkleRoot,
        bytes32[] calldata _merkleProof,
        uint8 _nodeIndex
    ) external view {
        require(merkleRoots[_merkleRoot], "merkle root not exists");
        require(smtContract.isSmtRootValid(_chainId, _proofData.smtRoot), "invalid smt root");
        require(dummyInputCommitments[_chainId] == _proofData.dummyInputCommitment, "invalid dummy input");

        bytes32 proofDataHash = keccak256(
            abi.encodePacked(
                _proofData.commitHash,
                _proofData.smtRoot,
                _proofData.appCommitHash,
                _proofData.appVkHash,
                _proofData.dummyInputCommitment
            )
        );
        bytes32 root = proofDataHash;
        for (uint8 depth = 0; depth < TREE_DEPTH; depth++) {
            if ((_nodeIndex >> depth) & 1 == 0) {
                root = keccak256(abi.encodePacked(root, _merkleProof[depth]));
            } else {
                root = keccak256(abi.encodePacked(_merkleProof[depth], root));
            }
        }
        require(_merkleRoot == root, "invalid data");
    }

    // -------- owner functions --------

    function updateSmtContract(ISMT _smtContract) public onlyOwner {
        smtContract = _smtContract;
        emit SmtContractUpdated(address(smtContract));
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

    function setDummyInputCommitments(
        uint64[] calldata _chainIds,
        bytes32[] calldata _dummyInputCommitments
    ) public onlyOwner {
        require(_chainIds.length == _dummyInputCommitments.length, "length not match");
        for (uint256 i = 0; i < _chainIds.length; i++) {
            dummyInputCommitments[_chainIds[i]] = _dummyInputCommitments[i];
        }
        emit DummyInputCommitmentsUpdated(_chainIds, _dummyInputCommitments);
    }

    /**********************************
     * Internal and Private Functions *
     **********************************/

    function unpack(
        bytes calldata _proofWithPubInputs
    ) internal pure returns (bytes32 merkleRoot, bytes32 proofIdsCommit) {
        merkleRoot = bytes32(_proofWithPubInputs[PUBLIC_BYTES_START_IDX:PUBLIC_BYTES_START_IDX + 32]);
        proofIdsCommit = bytes32(_proofWithPubInputs[PUBLIC_BYTES_START_IDX + 32:PUBLIC_BYTES_START_IDX + 2 * 32]);
    }
}
