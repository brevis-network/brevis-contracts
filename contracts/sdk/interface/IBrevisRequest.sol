// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../lib/Lib.sol";

interface IBrevisRequest {
    enum RequestStatus {
        Pending,
        ZkAttested,
        Refunded
    }

    struct Request {
        uint256 deadline;
        uint256 fee;
        address refundee;
        address callback;
        RequestStatus status;
    }

    event RequestSent(bytes32 requestId, address sender, uint256 fee, address callback);
    event RequestFulfilled(bytes32 requestId);
    event RequestsFulfilled(bytes32[] requestId);
    event RequestRefunded(bytes32 requestId);
    event RequestCallbackFailed(bytes32 requestId);
    event RequestsCallbackFailed(bytes32[] requestIds);
    event RequestTimeoutUpdated(uint256 from, uint256 to);

    function sendRequest(bytes32 _requestId, address _refundee, address _callback) external payable;

    function fulfillRequest(
        bytes32 _requestId,
        uint64 _chainId,
        bytes calldata _proof,
        bool _withAppProof,
        bytes calldata _appCircuitOutput
    ) external;

    function fulfillAggRequests(
        uint64 _chainId,
        bytes32[] calldata _requestIds,
        bytes calldata _proof,
        Brevis.ProofData[] calldata _proofDataArray,
        bytes[] calldata _appCircuitOutputs,
        address _callback
    ) external;

    function refund(bytes32 _requestId) external;

    function queryRequestStatus(bytes32 _requestId) external view returns (RequestStatus);
}
