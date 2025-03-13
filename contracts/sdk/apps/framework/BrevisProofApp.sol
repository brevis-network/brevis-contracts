// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

// App that directly interact with the BrevisProof contract.B
abstract contract BrevisProofApp {
    IBrevisProof public brevisProof;

    function _checkBrevisProof(
        uint64 _chainId,
        bytes calldata _proof,
        bytes calldata _appOutput,
        bytes32 _appVkHash
    ) internal {
        (, bytes32 appCommitHash, bytes32 appVkHash) = brevisProof.submitProof(_chainId, _proof);
        require(appVkHash == _appVkHash, "mismatch vkhash");
        require(appCommitHash == keccak256(_appOutput), "invalid circuit output");
    }

    function _checkBrevisAggProof(
        uint64 _chainId,
        bytes32[] calldata _proofIds,
        bytes calldata _proofWithPubInputs,
        IBrevisProof.ProofData[] calldata _proofDataArray
    ) internal {
        brevisProof.submitAggProof(_chainId, _proofIds, _proofWithPubInputs);
        brevisProof.validateAggProofData(_chainId, _proofDataArray);
    }
}

interface IBrevisProof {
    struct ProofData {
        bytes32 commitHash;
        bytes32 appCommitHash; // zk-program computing circuit commit hash
        bytes32 appVkHash; // zk-program computing circuit Verify Key hash
        bytes32 smtRoot;
        bytes32 dummyInputCommitment; // zk-program computing circuit dummy input commitment
    }

    function submitProof(
        uint64 _chainId,
        bytes calldata _proofWithPubInputs
    ) external returns (bytes32 proofId, bytes32 appCommitHash, bytes32 appVkHash);

    function submitAggProof(uint64 _chainId, bytes32[] calldata _proofIds, bytes calldata _proofWithPubInputs) external;

    function validateAggProofData(uint64 _chainId, ProofData[] calldata _proofDataArray) external view;
}
