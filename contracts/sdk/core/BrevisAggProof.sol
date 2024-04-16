// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/access/Ownable.sol";
import "../lib/Lib.sol";
import "../../interfaces/ISMT.sol";
import "../../verifiers/interfaces/IZkpVerifier.sol";

contract BrevisAggProof is Ownable {
    ISMT public smtContract;

    constructor(ISMT _smtContract) {
        smtContract = _smtContract;
    }

    mapping(bytes32 => bool) public merkleRoots;
    mapping(bytes32 => bool) public requestIds;
    mapping(uint64 => IZkpVerifier) public aggProofVerifierAddress;
    event SmtContractUpdated(ISMT smtContract);
    event AggProofVerifierAddressesUpdated(uint64[] chainIds, IZkpVerifier[] newAddresses);

    uint32 constant PUBLIC_BYTES_START_IDX = 12 * 32; // the first 12 32bytes are groth16 proof (A/B/C/Commitment/CommitmentPOK)
    uint8 constant TREE_DEPTH = 4;
    uint256 constant LEAF_NODES_LEN = 2 ** TREE_DEPTH;

    function mustValidateRequest(
        uint64 _chainId,
        Brevis.ProofData calldata _proofData,
        bytes32 _merkleRoot,
        bytes32[] calldata _merkleProof,
        uint8 _nodeIndex
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
            if ((_nodeIndex >> depth) & 1 == 0) {
                root = keccak256(abi.encodePacked(root, _merkleProof[depth]));
            } else {
                root = keccak256(abi.encodePacked(_merkleProof[depth], root));
            }
        }
        require(_merkleRoot == root, "invalid data");
    }

    function mustValidateRequests(uint64 _chainId, Brevis.ProofData[] calldata _proofDataArray) external view {
        uint dataLen = _proofDataArray.length;
        require(dataLen <= LEAF_NODES_LEN, "size exceeds");
        bytes32[2 * LEAF_NODES_LEN - 1] memory hashes;
        for (uint i = 0; i < dataLen; i++) {
            require(smtContract.isSmtRootValid(_chainId, _proofDataArray[i].smtRoot), "invalid smt root");
            hashes[i] = keccak256(
                abi.encodePacked(
                    _proofDataArray[i].commitHash,
                    _proofDataArray[i].smtRoot,
                    _proofDataArray[i].vkHash,
                    _proofDataArray[i].appCommitHash,
                    _proofDataArray[i].appVkHash
                )
            );
        }
        // note, hashes[dataLen] to hashes[LEAF_NODES_LEN - 1] defaults to 0
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

    function mustSubmitAggProof(
        uint64 _chainId,
        bytes32[] calldata _requestIds,
        bytes calldata _proofWithPubInputs
    ) external {
        IZkpVerifier verifier = aggProofVerifierAddress[_chainId];
        require(address(verifier) != address(0), "chain agg proof verifier not set");
        require(verifier.verifyRaw(_proofWithPubInputs), "proof not valid");

        (bytes32 root, bytes32 commitHash) = unpack(_proofWithPubInputs);
        require(keccak256(abi.encodePacked(_requestIds)) == commitHash, "requestIds not right");
        merkleRoots[root] = true;
        for (uint i = 0; i < _requestIds.length; i++) {
            requestIds[_requestIds[i]] = true;
        }
    }

    function inAgg(bytes32 _requestId) public view returns (bool) {
        return requestIds[_requestId];
    }

    function unpack(bytes calldata _proofWithPubInputs) internal pure returns (bytes32 merkleRoot, bytes32 commitHash) {
        merkleRoot = bytes32(_proofWithPubInputs[PUBLIC_BYTES_START_IDX:PUBLIC_BYTES_START_IDX + 32]);
        commitHash = bytes32(_proofWithPubInputs[PUBLIC_BYTES_START_IDX + 32:PUBLIC_BYTES_START_IDX + 2 * 32]);
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
