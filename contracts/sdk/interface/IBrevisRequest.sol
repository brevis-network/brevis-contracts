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

    // TODO: save gas
    struct Request {
        uint64 timestamp;
        RequestStatus status;
        Callback callback;
        bytes32 feeHash; // keccak256(abi.encodePacked(amount, refundee));
    }

    struct Fee {
        uint256 amount;
        address refundee;
    }

    struct Callback {
        address target;
        uint64 gas;
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
        address refundee,
        uint256 fee,
        address callback,
        uint64 gas,
        RequestOption option
    );
    event RequestFulfilled(bytes32 requestId);
    event RequestsFulfilled(bytes32[] requestIds);
    event RequestRefunded(bytes32 requestId);
    event RequestCallbackFailed(bytes32 requestId);
    event RequestsCallbackFailed(bytes32[] requestIds);

    event OpRequestsFulfilled(bytes32[] requestIds, bytes32[] appCommitHashes, bytes32[] appVkHashes, bytes[] urls);
    event AskFor(bytes32 indexed requestId, DisputeStatus status, address from);
    event RequestDataPosted(bytes32 indexed requestId);
    event DataAvailabilityProofPosted(bytes32 indexed requestId);
    event DataValidityProofProofPosted(bytes32 indexed requestId);

    event RequestTimeoutUpdated(uint256 from, uint256 to);
    event ChallengeWindowUpdated(uint256 from, uint256 to);
    event ResponseTimeoutUpdated(uint256 from, uint256 to);

    function sendRequest(
        bytes32 _requestId,
        address _refundee,
        address _callback,
        uint64 _gas,
        RequestOption option
    ) external payable;

    function fulfillRequest(
        bytes32 _requestId,
        uint64 _chainId,
        bytes calldata _proof,
        bytes calldata _appCircuitOutput,
        address _callback
    ) external;

    function fulfillRequests(
        bytes32[] calldata _requestIds,
        uint64 _chainId,
        bytes calldata _proof,
        Brevis.ProofData[] calldata _proofDataArray,
        bytes[] calldata _appCircuitOutputs,
        address[] calldata callbacks
    ) external;

    function fulfillOpRequests(
        bytes32[] calldata _requestIds,
        bytes32[] calldata _appCommitHashes,
        bytes32[] calldata _appVkHashes,
        bytes[] calldata _dataURLs,
        bytes[] calldata _sigs,
        address[] calldata _signers,
        uint256[] calldata _powers
    ) external;

    function refund(bytes32 _requestId, uint256 _amount, address _refundee) external;

    function askForRequestData(bytes32 _requestId) external payable;

    function postRequestData(bytes32 _requestId, bytes calldata _requestData) external;

    function askForDataAvailabilityProof(bytes32 _requestId) external payable;

    function postDataAvailabilityProof(bytes32 _requestId, bytes calldata _proof) external;

    function askForDataValidityProof(bytes32 _requestId) external payable;

    function postDataValidityProof(bytes32 _requestId, uint64 _chainId, bytes calldata _proof) external;

    function queryRequestStatus(bytes32 _requestId) external view returns (RequestStatus);

    function queryRequestStatus(bytes32 _requestId, uint256 _appChallengeWindow) external view returns (RequestStatus);

    function validateRequestOpData(
        bytes32 _requestId,
        bytes32 _appCommitHash,
        bytes32 _appVkHash
    ) external view returns (bool);
}
