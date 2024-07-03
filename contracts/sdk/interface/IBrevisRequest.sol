// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../lib/Lib.sol";
import "../../interfaces/ISigsVerifier.sol";

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
        RequestStatus status;
        uint64 timestamp;
        uint8 option;
    }

    struct OnchainRequestInfo {
        bytes32 feeHash; // keccak256(abi.encodePacked(amount, refundee))
        Callback callback;
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

    struct RequestDataHash {
        bytes32[] hashes;
        bytes32 root;
    }

    struct Dispute {
        DisputeStatus status;
        address challenger;
        RequestDataHash requestDataHash;
        uint256 responseDeadline;
        uint256 deposit;
    }

    // todo: reduce event fields
    event RequestSent(bytes32 proofId, uint64 nonce, address refundee, uint256 fee, Callback callback, uint8 option);
    event RequestFulfilled(bytes32 proofId, uint64 nonce);
    event RequestsFulfilled(bytes32[] proofIds, uint64[] nonces);
    event RequestRefunded(bytes32 proofId, uint64 nonce);
    event RequestCallbackFailed(bytes32 proofId, uint64 nonce);
    event RequestsCallbackFailed(bytes32[] proofIds, uint64[] nonces);
    event RequestFeeIncreased(bytes32 proofId, uint64 nonce, uint256 gas, uint256 fee);

    event OpRequestsFulfilled(bytes32[] proofIds, uint64[] nonces, bytes32[] appCommitHashes, bytes32[] appVkHashes);
    event AskFor(bytes32 indexed proofId, uint64 nonce, DisputeStatus status, address from);
    event RequestDataPosted(bytes32 indexed proofId, uint64 nonce, bytes[] data, uint256 index, bool done);
    event DataAvailabilityProofPosted(bytes32 indexed proofId, uint64 nonce);
    event DataValidityProofProofPosted(bytes32 indexed proofId, uint64 nonce);

    event RequestTimeoutUpdated(uint256 from, uint256 to);
    event ChallengeWindowUpdated(uint256 from, uint256 to);
    event ResponseTimeoutUpdated(uint256 from, uint256 to);
    event BaseDataUrlUpdated(string from, string to);
    event BrevisProofUpdated(address from, address to);
    event BvnSigsVerifierUpdated(address from, address to);
    event AvsSigsVerifierUpdated(address from, address to);
    event DisputeDepositsUpdated(uint256 amtAskForData, uint256 amtAskForProof);

    function sendRequest(
        bytes32 _proofId,
        uint64 _nonce,
        address _refundee,
        Callback calldata _callback,
        uint8 _option
    ) external payable;

    function fulfillRequest(
        bytes32 _proofId,
        uint64 _nonce,
        uint64 _chainId,
        bytes calldata _proof,
        bytes calldata _appCircuitOutput,
        address _callbackTarget
    ) external;

    function fulfillRequests(
        bytes32[] calldata _proofIds,
        uint64[] calldata _nonces,
        uint64 _chainId,
        bytes calldata _proof,
        Brevis.ProofData[] calldata _proofDataArray,
        bytes[] calldata _appCircuitOutputs,
        address[] calldata _callbackTargets
    ) external;

    function fulfillOpRequests(
        bytes32[] calldata _proofIds,
        uint64[] calldata _nonces,
        bytes32[] calldata _appCommitHashes,
        bytes32[] calldata _appVkHashes,
        IBvnSigsVerifier.SigInfo calldata _bvnSigInfo,
        IAvsSigsVerifier.SigInfo calldata _eigenSigInfo
    ) external;

    function refund(bytes32 _proofId, uint64 _nonce, uint256 _amount, address _refundee) external;

    function increaseGasFee(
        bytes32 _proofId,
        uint64 _nonce,
        uint64 _addGas,
        uint256 _currentFee,
        address _refundee
    ) external payable;

    function askForRequestData(bytes32 _proofId, uint64 _nonce) external payable;

    function postRequestData(
        bytes32 _proofId,
        uint64 _nonce,
        bytes[] calldata _requestData,
        uint256 _index,
        bool _done
    ) external;

    function askForDataAvailabilityProof(bytes32 _proofId, uint64 _nonce) external payable;

    function postDataAvailabilityProof(bytes32 _proofId, uint64 _nonce, bytes calldata _proof) external;

    function askForDataValidityProof(bytes32 _proofId, uint64 _nonce) external payable;

    function postDataValidityProof(bytes32 _proofId, uint64 _nonce, uint64 _chainId, bytes calldata _proof) external;

    function queryRequestStatus(bytes32 _proofId, uint64 _nonce) external view returns (RequestStatus, uint8);

    function queryRequestStatus(
        bytes32 _proofId,
        uint64 _nonce,
        uint256 _appChallengeWindow
    ) external view returns (RequestStatus, uint8);

    function validateOpAppData(
        bytes32 _proofId,
        uint64 _nonce,
        bytes32 _appCommitHash,
        bytes32 _appVkHash,
        uint8 _option
    ) external view returns (bool);

    function validateOpAppData(
        bytes32 _proofId,
        uint64 _nonce,
        bytes32 _appCommitHash,
        bytes32 _appVkHash,
        uint256 _appChallengeWindow,
        uint8 _option
    ) external view returns (bool);

    function dataURL(bytes32 _proofId) external view returns (string memory);
}
