// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../lib/Lib.sol";

interface IBrevisRequest {
    enum RequestStatus {
        Pending,
        ZkAttested,
        OpSubmitted,
        OpQueryDataSubmitted,
        OpDisputing,
        OpDisputed,
        OpAttested,
        Refunded
    }

    enum Option {
        ZkMode,
        OpMode_MIMC,
        OpMode_KECCAK
    }

    enum AskForType {
        NULL,
        QueryData,
        Proof
    }

    struct Request {
        uint256 deadline;
        uint256 fee;
        address refundee;
        address callback;
        RequestStatus status;
        Option option;
    }

    struct RequestExt {
        uint256 canChallengeBefore;
        AskForType askFor;
        uint256 shouldRespondBefore;
    }

    event RequestSent(bytes32 requestId, address sender, uint256 fee, address callback, Option option);
    event RequestFulfilled(bytes32 requestId);
    event RequestsFulfilled(bytes32[] requestId);
    event RequestRefunded(bytes32 requestId);
    event RequestCallbackFailed(bytes32 requestId);
    event RequestsCallbackFailed(bytes32[] requestIds);

    event OpRequestsFulfilled(bytes32[] requestIds, bytes[] queryURLs);
    event AskFor(bytes32 indexed requestId, AskForType askFor, address from);
    event QueryDataPost(bytes32 indexed requestId);
    event ProofPost(bytes32 indexed requestId);

    event RequestTimeoutUpdated(uint256 from, uint256 to);
    event ChallengeWindowUpdated(uint256 from, uint256 to);
    event ResponseTimeoutUpdated(uint256 from, uint256 to);

    function sendRequest(bytes32 _requestId, address _refundee, address _callback, Option _option) external payable;

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

    function askForQueryData(bytes32 _requestId) external payable;

    function postQueryData(bytes32 _requestId, bytes calldata _queryData) external;

    function challengeQueryData(bytes calldata _proof) external;

    function askForProof(bytes32 _requestId) external payable;

    function postProof(bytes32 _requestId, uint64 _chainId, bytes calldata _proof) external;

    function queryRequestStatus(bytes32 _requestId) external view returns (RequestStatus);
}
