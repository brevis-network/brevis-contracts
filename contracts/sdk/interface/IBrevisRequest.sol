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

    // todo: reduce event fields
    event RequestSent(
        bytes32 proofId,
        uint64 nonce,
        address refundee,
        uint256 fee,
        address callback,
        uint64 gas,
        RequestOption option
    );
    event RequestFulfilled(bytes32 proofId, uint64 nonce);
    event RequestsFulfilled(bytes32[] proofIds, uint64[] nonces);
    event RequestRefunded(bytes32 proofId, uint64 nonce);
    event RequestCallbackFailed(bytes32 proofId, uint64 nonce);
    event RequestsCallbackFailed(bytes32[] proofIds, uint64[] nonces);

    event OpRequestsFulfilled(
        bytes32[] proofIds,
        uint64[] nonces,
        bytes32[] appCommitHashes,
        bytes32[] appVkHashes,
        bytes[] urls
    );
    event AskFor(bytes32 indexed proofId, uint64 nonce, DisputeStatus status, address from);
    event RequestDataPosted(bytes32 indexed proofId, uint64 nonce);
    event DataAvailabilityProofPosted(bytes32 indexed proofId, uint64 nonce);
    event DataValidityProofProofPosted(bytes32 indexed proofId, uint64 nonce);

    event RequestTimeoutUpdated(uint256 from, uint256 to);
    event ChallengeWindowUpdated(uint256 from, uint256 to);
    event ResponseTimeoutUpdated(uint256 from, uint256 to);
    event DisputeDepositsUpdated(uint256 amtAskForData, uint256 amtAskForProof);

    function sendRequest(
        bytes32 _proofId,
        uint64 _nonce,
        address _refundee,
        address _callback,
        uint64 _gas,
        RequestOption option
    ) external payable;

    function fulfillRequest(
        bytes32 _proofId,
        uint64 _nonce,
        uint64 _chainId,
        bytes calldata _proof,
        bytes calldata _appCircuitOutput,
        address _callback
    ) external;

    function fulfillRequests(
        bytes32[] calldata _proofIds,
        uint64[] calldata _nonces,
        uint64 _chainId,
        bytes calldata _proof,
        Brevis.ProofData[] calldata _proofDataArray,
        bytes[] calldata _appCircuitOutputs,
        address[] calldata callbacks
    ) external;

    function fulfillOpRequests(
        bytes32[] calldata _proofIds,
        uint64[] calldata _nonces,
        bytes32[] calldata _appCommitHashes,
        bytes32[] calldata _appVkHashes,
        bytes[] calldata _dataURLs,
        bytes[] calldata _sigs,
        address[] calldata _signers,
        uint256[] calldata _powers
    ) external;

    function refund(bytes32 _proofId, uint64 _nonce, uint256 _amount, address _refundee) external;

    function askForRequestData(bytes32 _proofId, uint64 _nonce) external payable;

    function postRequestData(bytes32 _proofId, uint64 _nonce, bytes calldata _requestData) external;

    function askForDataAvailabilityProof(bytes32 _proofId, uint64 _nonce) external payable;

    function postDataAvailabilityProof(bytes32 _proofId, uint64 _nonce, bytes calldata _proof) external;

    function askForDataValidityProof(bytes32 _proofId, uint64 _nonce) external payable;

    function postDataValidityProof(bytes32 _proofId, uint64 _nonce, uint64 _chainId, bytes calldata _proof) external;

    function queryRequestStatus(bytes32 _proofId, uint64 _nonce) external view returns (RequestStatus);

    function queryRequestStatus(
        bytes32 _proofId,
        uint64 _nonce,
        uint256 _appChallengeWindow
    ) external view returns (RequestStatus);

    function validateOpAppData(
        bytes32 _proofId,
        uint64 _nonce,
        bytes32 _appCommitHash,
        bytes32 _appVkHash
    ) external view returns (bool);

    function validateOpAppData(
        bytes32 _proofId,
        uint64 _nonce,
        bytes32 _appCommitHash,
        bytes32 _appVkHash,
        uint256 _appChallengeWindow
    ) external view returns (bool);
}
