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
        WaitingForQueryData,
        QueryDataPosted,
        WaitingForQueryDataProof,
        QueryDataProofPosted,
        WaitingForValidityProof,
        ValidityProofPosted
    }

    struct Dispute {
        DisputeStatus status;
        bytes32 queryDataHash;
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

    event OpRequestsFulfilled(bytes32[] requestIds, uint256[] nonces, bytes[] queryURLs);
    event AskFor(bytes32 indexed requestId, uint256 nonce, DisputeStatus status, address from);
    event QueryDataPosted(bytes32 indexed requestId, uint256 nonce);
    event QueryDataProofPosted(bytes32 indexed requestId, uint256 nonce);
    event ValidityProofProofPosted(bytes32 indexed requestId, uint256 nonce);

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
        bytes[] calldata _queryURLs,
        bytes[] calldata _sigs,
        address[] calldata _signers,
        uint256[] calldata _powers
    ) external;

    function refund(bytes32 _requestId, uint256 _nonce) external;

    function askForQueryData(bytes32 _requestId, uint256 _nonce) external payable;

    function postQueryData(bytes32 _requestId, uint256 _nonce, bytes calldata _queryData) external;

    function askForQueryDataProof(bytes32 _requestId, uint256 _nonce) external payable;

    function postQueryDataProof(bytes32 _requestId, uint256 _nonce, bytes calldata _proof) external;

    function askForValidityProof(bytes32 _requestId, uint256 _nonce) external payable;

    function postValidityProof(bytes32 _requestId, uint256 _nonce, uint64 _chainId, bytes calldata _proof) external;

    function queryRequestStatus(bytes32 _requestId, uint256 _nonce) external view returns (RequestStatus);

    function queryRequestStatus(
        bytes32 _requestId,
        uint256 _nonce,
        uint256 _appChallengeWindow
    ) external view returns (RequestStatus);
}
