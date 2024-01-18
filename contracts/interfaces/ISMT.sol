// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../verifiers/zk-verifiers/common/IVerifier.sol";

interface ISMT {
    struct SmtUpdate {
        bytes32 newSmtRoot;
        uint64 endBlockNum;
        bytes32 endBlockHash;
        bytes32 nextChunkMerkleRoot;
        IVerifier.Proof proof;
        bytes32 commitPub;
    }

    function updateRoot(uint64 chainId, SmtUpdate memory u) external;

    function isSmtRootValid(uint64 chainId, bytes32 smtRoot) external view returns (bool);
}
