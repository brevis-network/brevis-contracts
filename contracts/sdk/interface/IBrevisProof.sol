// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../lib/Lib.sol";

interface IBrevisProof {
    function submitProof(uint64 _chainId, bytes calldata _proofWithPubInputs) external returns (bytes32 _requestId);

    function hasProof(bytes32 _requestId) external view returns (bool);

    // return appCommitHash and appVkHash
    function getProofAppData(bytes32 _requestId) external view returns (bytes32, bytes32);

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
