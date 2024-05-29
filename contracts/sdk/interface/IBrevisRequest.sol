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

    enum RequestOption {
        Zk,
        Op
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
        WaitingForRequestData,
        RequestDataPosted,
        WaitingForDataAvailabilityProof,
        DataAvailabilityProofPosted,
        WaitingForDataValidityProof,
        DataValidityProofPosted
    }

    struct Dispute {
        DisputeStatus status;
        bytes32 requestDataHash;
        uint256 responseDeadline;
    }

    event RequestSent(
        bytes32 requestId,
        uint256 _nonce,
        address sender,
        uint256 fee,
        address callback,
        RequestOption option
    );
    event RequestFulfilled(bytes32 requestId, uint256 nonce);
    event RequestsFulfilled(bytes32[] requestIds, uint256[] nonces);
    event RequestRefunded(bytes32 requestId, uint256 nonce);
    event RequestCallbackFailed(bytes32 requestId, uint256 nonce);

    event OpRequestsFulfilled(bytes32[] requestIds, uint256[] nonces, bytes[] URLs);
    event AskFor(bytes32 indexed requestId, uint256 nonce, DisputeStatus status, address from);
    event RequestDataPosted(bytes32 indexed requestId, uint256 nonce);
    event DataAvailabilityProofPosted(bytes32 indexed requestId, uint256 nonce);
    event DataValidityProofProofPosted(bytes32 indexed requestId, uint256 nonce);

    event RequestTimeoutUpdated(uint256 from, uint256 to);
    event ChallengeWindowUpdated(uint256 from, uint256 to);
    event ResponseTimeoutUpdated(uint256 from, uint256 to);

    function sendRequest(
        bytes32 _requestId,
        uint256 _nonce,
        address _refundee,
        address _callback,
        RequestOption option
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
        bytes calldata _proof
    ) external;

    function fulfillRequests(
        bytes32[] calldata _requestIds,
        uint256[] calldata _nonces,
        uint64 _chainId,
        bytes calldata _proof,
        Brevis.ProofData[] calldata _proofDataArray,
        bytes[] calldata _appCircuitOutputs
    ) external;

    function fulfillOpRequests(
        bytes32[] calldata _requestIds,
        uint256[] calldata _nonces,
        bytes[] calldata _dataURLs,
        bytes[] calldata _sigs,
        address[] calldata _signers,
        uint256[] calldata _powers
    ) external;

    function refund(bytes32 _requestId, uint256 _nonce) external;

    function askForRequestData(bytes32 _requestId, uint256 _nonce) external payable;

    function postRequestData(bytes32 _requestId, uint256 _nonce, bytes calldata _requestData) external;

    function askForDataAvailabilityProof(bytes32 _requestId, uint256 _nonce) external payable;

    function postDataAvailabilityProof(bytes32 _requestId, uint256 _nonce, bytes calldata _proof) external;

    function askForDataValidityProof(bytes32 _requestId, uint256 _nonce) external payable;

    function postDataValidityProof(bytes32 _requestId, uint256 _nonce, uint64 _chainId, bytes calldata _proof) external;

    function queryRequestStatus(bytes32 _requestId, uint256 _nonce) external view returns (RequestStatus);

    function queryRequestStatus(
        bytes32 _requestId,
        uint256 _nonce,
        uint256 _appChallengeWindow
    ) external view returns (RequestStatus);
}
