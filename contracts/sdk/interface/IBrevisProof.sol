// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../lib/Lib.sol";

interface IBrevisProof {
    function submitProof(
        uint64 _chainId,
        bytes calldata _proofWithPubInputs
    ) external returns (bytes32 requestId, bytes32 appCommitHash, bytes32 appVkHash);

    function hasProof(bytes32 _requestId) external view returns (bool);

    function hasProofAppData(
        bytes32 _requestId,
        bytes32 _appCommitHash,
        bytes32 _appVkHash
    ) external view returns (bool);

    function mustValidateRequest(
        uint64 _chainId,
        Brevis.ProofData calldata _proofData,
        bytes32 _merkleRoot,
        bytes32[] calldata _merkleProof,
        uint8 _nodeIndex
    ) external view;

    function mustValidateRequests(uint64 _chainId, Brevis.ProofData[] calldata _proofDataArray) external view;

    function mustSubmitAggProof(
        uint64 _chainId,
        bytes32[] calldata _requestIds,
        bytes calldata _proofWithPubInputs
    ) external;

    function getRequestContract() external view returns (address);
}
