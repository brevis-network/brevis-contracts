// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/access/Ownable.sol";
import "../light-client/interfaces/IAnchorBlocks.sol";
import "../interfaces/ISMT.sol";

contract SMT is ISMT, Ownable {
    event SmtRootUpdated(bytes32 smtRoot, uint64 endBlockNum, uint8 bufferIndex);
    event AnchorProviderUpdated(uint64 chainId, address anchorProvider);
    event VerifierUpdated(uint64 chainId, address verifier);

    uint8 public constant BUFFER_SIZE = 16;

    mapping(uint64 => IAnchorBlocks) public anchorProviders;
    mapping(uint64 => IVerifier) public verifiers;

    mapping(uint64 => bytes32[BUFFER_SIZE]) public smtRoots;
    mapping(uint64 => uint8) public curBufferIndices;

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
            smtRoots[chid][0] = _initRoots[i];
        }
    }

    function getLatestRoot(uint64 chainId) public view returns (bytes32 root, uint8 bufferIndex) {
        bytes32[BUFFER_SIZE] memory roots = smtRoots[chainId];
        uint8 index = curBufferIndices[chainId];
        return (roots[index], index);
    }

    function getRoot(uint64 chainId, uint8 bufferIndex) public view returns (bytes32 root) {
        return smtRoots[chainId][bufferIndex];
    }

    function isSmtRootValid(uint64 chainId, bytes32 smtRoot) public view returns (bool) {
        bytes32[BUFFER_SIZE] memory roots = smtRoots[chainId];
        for (uint256 i = 0; i < roots.length; i++) {
            if (roots[i] == smtRoot) {
                return true;
            }
        }
        return false;
    }

    function updateRoot(uint64 chainId, SmtUpdate memory u) external {
        // If nextChunkMerkleRoot is empty, it means the zk proof bypasses checking if the updated chunk anchors to a known chunk.
        // Instead, the responsibility of checking the validity of endBlockHash is deferred to this contract.
        if (u.nextChunkMerkleRoot == 0) {
            IAnchorBlocks anchorProvider = anchorProviders[chainId];
            require(address(anchorProvider) != address(0), "unknown anchor provider");
            bytes32 anchorHash = anchorProvider.blocks(u.endBlockNum);
            require(anchorHash == u.endBlockHash, "anchor check failed");
        }
        uint8 curIndex = curBufferIndices[chainId];
        bytes32 root = smtRoots[chainId][curIndex];
        bool success = verifyProof(chainId, root, u);
        require(success, "invalid zk proof");

        curIndex = (curIndex + 1) % BUFFER_SIZE;
        smtRoots[chainId][curIndex] = u.newSmtRoot;
        curBufferIndices[chainId] = curIndex;
        emit SmtRootUpdated(u.newSmtRoot, u.endBlockNum, curIndex);
    }

    function verifyProof(uint64 chainId, bytes32 oldSmtRoot, SmtUpdate memory u) private view returns (bool) {
        IVerifier verifier = verifiers[chainId];
        require(address(verifier) != address(0), "no verifier for chainId");

        uint256[10] memory input;
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
        input[9] = uint256(u.commitPub);

        return verifier.verifyProof(u.proof.a, u.proof.b, u.proof.c, u.proof.commitment, input);
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
