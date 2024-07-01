// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../light-client-eth/interfaces/IAnchorBlocks.sol";
import "../interfaces/ISMT.sol";
import "../safeguard/BrevisAccess.sol";

contract SMT is ISMT, BrevisAccess {
    event SmtRootUpdated(bytes32 smtRoot, uint64 endBlockNum, uint64 chainId);
    event AnchorProviderUpdated(uint64 chainId, address anchorProvider);
    event VerifierUpdated(uint64 chainId, address verifier);

    mapping(uint64 => IAnchorBlocks) public anchorProviders;
    mapping(uint64 => IVerifier) public verifiers;

    mapping(uint64 => mapping(bytes32 => bool)) public smtRoots;
    mapping(uint64 => bytes32) public latestRoots;

    constructor(
        uint64[] memory _chainIds,
        address[] memory _anchorProviders,
        address[] memory _verifiers,
        bytes32[] memory _initRoots
    ) {
        require(_chainIds.length == _anchorProviders.length, "len mismatch");
        require(_chainIds.length == _verifiers.length, "len mismatch");
        require(_chainIds.length == _initRoots.length, "len mismatch");
        for (uint256 i = 0; i < _chainIds.length; i++) {
            uint64 chid = _chainIds[i];
            anchorProviders[chid] = IAnchorBlocks(_anchorProviders[i]);
            verifiers[chid] = IVerifier(_verifiers[i]);
            smtRoots[chid][_initRoots[i]] = true;
            latestRoots[chid] = _initRoots[i];
        }
    }

    function getLatestRoot(uint64 chainId) public view returns (bytes32) {
        return latestRoots[chainId];
    }

    function isSmtRootValid(uint64 chainId, bytes32 smtRoot) public view returns (bool) {
        return smtRoots[chainId][smtRoot];
    }

    function updateRoot(uint64 chainId, SmtUpdate memory u) external onlyActiveProver {
        // If nextChunkMerkleRoot is empty, it means the zk proof bypasses checking if the updated chunk anchors to a known chunk.
        // Instead, the responsibility of checking the validity of endBlockHash is deferred to this contract.
        if (u.nextChunkMerkleRoot == 0) {
            IAnchorBlocks anchorProvider = anchorProviders[chainId];
            require(address(anchorProvider) != address(0), "unknown anchor provider");
            bytes32 anchorHash = anchorProvider.blocks(u.endBlockNum);
            require(anchorHash == u.endBlockHash, "anchor check failed");
        }
        bytes32 root = latestRoots[chainId];
        bool success = verifyProof(chainId, root, u);
        require(success, "invalid zk proof");

        smtRoots[chainId][u.newSmtRoot] = true;
        latestRoots[chainId] = u.newSmtRoot;
        emit SmtRootUpdated(u.newSmtRoot, u.endBlockNum, chainId);
    }

    function verifyProof(uint64 chainId, bytes32 oldSmtRoot, SmtUpdate memory u) private view returns (bool) {
        IVerifier verifier = verifiers[chainId];
        require(address(verifier) != address(0), "no verifier for chainId");

        uint256[9] memory input;
        uint256 m = 1 << 128;
        input[0] = uint256(oldSmtRoot) >> 128;
        input[1] = uint256(oldSmtRoot) % m;
        input[2] = uint256(u.newSmtRoot) >> 128;
        input[3] = uint256(u.newSmtRoot) % m;
        input[4] = uint256(u.endBlockHash) >> 128;
        input[5] = uint256(u.endBlockHash) % m;
        input[6] = u.endBlockNum;
        input[7] = uint256(u.nextChunkMerkleRoot) >> 128;
        input[8] = uint256(u.nextChunkMerkleRoot) % m;
       
        return verifier.verifyProof(u.proof, u.commit, u.knowledgeProof, input);
    }

    function setAnchorProvider(uint64 chainId, address anchorProvider) external onlyOwner {
        anchorProviders[chainId] = IAnchorBlocks(anchorProvider);
        emit AnchorProviderUpdated(chainId, anchorProvider);
    }

    function setVerifier(uint64 chainId, address verifier) external onlyOwner {
        verifiers[chainId] = IVerifier(verifier);
        emit VerifierUpdated(chainId, verifier);
    }
}
