// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../lib/Lib.sol";

interface IBrevisRequest {
    enum RequestStatus {
        Null,
        ZkPending,
        ZkAttested,
        OpPending,
        OpSubmitted,
        OpDisputing,
        OpDisputed,
        OpAttested,
        Refunded
    }

    struct Request {
        uint256 timestamp;
        uint256 fee;
        address refundee;
        address callback;
        RequestStatus status;
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

    event RequestSent(bytes32 requestId, uint256 _nonce, address sender, uint256 fee, address callback, bool zk);
    event RequestFulfilled(bytes32 requestId, uint256 nonce);
    event RequestsFulfilled(bytes32[] requestIds, uint256[] nonces);
    event RequestRefunded(bytes32 requestId, uint256 nonce);
    event RequestCallbackFailed(bytes32 requestId, uint256 nonce);
    event RequestsCallbackFailed(bytes32[] requestIds, uint256[] nonces);

    event OpRequestsFulfilled(bytes32[] requestIds, uint256[] nonces, bytes[] queryURLs);
    event Challenge(bytes32 indexed requestId, uint256 nonce, DisputeStatus status, address from);
    event QueryDataPost(bytes32 indexed requestId, uint256 nonce);
    event ProofPost(bytes32 indexed requestId, uint256 nonce);

    event RequestTimeoutUpdated(uint256 from, uint256 to);
    event ChallengeWindowUpdated(uint256 from, uint256 to);
    event ResponseTimeoutUpdated(uint256 from, uint256 to);

    function sendRequest(
        bytes32 _requestId,
        uint256 _nonce,
        address _refundee,
        address _callback,
        bool _zk
    ) external payable;

    function fulfillRequest(
        bytes32 _requestId,
        uint256 _nonce,
        uint64 _chainId,
        bytes calldata _proof,
        bytes calldata _appCircuitOutput
    ) external;

    function fulfillRequests(
        bytes32[] calldata _requestIds,
        uint256[] calldata _nonces,
        uint64 _chainId,
        bytes calldata _proof,
        Brevis.ProofData[] calldata _proofDataArray,
        bytes[] calldata _appCircuitOutputs,
        address _callback
    ) external;

    function fulfillOpRequests(
        bytes32[] calldata _requestIds,
        uint256[] calldata _nonces,
        bytes[] calldata _queryURLs,
        bytes[] calldata _sigs,
        address[] calldata _signers,
        uint256[] calldata _powers
    ) external;

    function refund(bytes32 _requestId, uint256 _nonce) external;

    function askForQueryData(bytes32 _requestId, uint256 _nonce) external payable;

    function postQueryData(bytes32 _requestId, uint256 _nonce, bytes calldata _queryData) external;

    function challengeQueryData(bytes calldata _proof, uint256 _nonce) external;

    function askForProof(bytes32 _requestId, uint256 _nonce) external payable;

    function postProof(bytes32 _requestId, uint256 _nonce, uint64 _chainId, bytes calldata _proof) external;

    function queryRequestStatus(bytes32 _requestId, uint256 _nonce) external view returns (RequestStatus);
}
