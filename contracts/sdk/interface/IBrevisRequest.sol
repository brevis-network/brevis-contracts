// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../lib/Lib.sol";

interface IBrevisRequest {
    enum RequestStatus {
        Null,
        Pending,
        ZkAttested,
        Refunded,
        OpSubmitted,
        OpDisputing,
        OpDisputed,
        OpAttested
    }

    enum Option {
        ZkMode,
        OpMode_MIMC,
        OpMode_KECCAK
    }

    struct Request {
        uint256 timestamp;
        uint256 fee;
        address refundee;
        address callback;
        RequestStatus status;
        Option option; // TODO: remove this
    }

    enum DisputeStatus {
        Null,
        WaitingForQueryData,
        QueryDataPosted,
        WaitingForZkProof
    }

    struct Dispute {
        DisputeStatus status;
        uint256 responseDeadline;
    }

    event RequestSent(bytes32 requestId, address sender, uint256 fee, address callback, Option option);
    event RequestFulfilled(bytes32 requestId);
    event RequestsFulfilled(bytes32[] requestIds);
    event RequestRefunded(bytes32 requestId);
    event RequestCallbackFailed(bytes32 requestId);
    event RequestsCallbackFailed(bytes32[] requestIds);

    event OpRequestsFulfilled(bytes32[] requestIds, bytes[] queryURLs);
    event Challenge(bytes32 indexed requestId, DisputeStatus status, address from);
    event QueryDataPost(bytes32 indexed requestId);
    event ProofPost(bytes32 indexed requestId);

    event RequestTimeoutUpdated(uint256 from, uint256 to);
    event ChallengeTimeoutUpdated(uint256 from, uint256 to);
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

    function fulfillOpRequests(
        bytes32[] calldata _requestIds,
        bytes[] calldata _queryURLs,
        bytes[] calldata _sigs,
        address[] calldata _signers,
        uint256[] calldata _powers
    ) external;

    function askForQueryData(bytes32 _requestId) external payable;

    function postQueryData(bytes32 _requestId, bytes calldata _queryData) external;

    function challengeQueryData(bytes calldata _proof) external;

    function askForProof(bytes32 _requestId) external payable;

    function refund(bytes32 _requestId) external;

    function postProof(bytes32 _requestId, uint64 _chainId, bytes calldata _proof) external;

    function queryRequestStatus(bytes32 _requestId) external view returns (RequestStatus);
}
