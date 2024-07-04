// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "./IBrevisTypes.sol";
import "../lib/Lib.sol";
import "../../interfaces/ISigsVerifier.sol";

interface IBrevisRequest is IBrevisTypes {
    // todo: reduce event fields
    event RequestSent(bytes32 proofId, uint64 nonce, address refundee, uint256 fee, Callback callback, uint8 option);

    event RequestFulfilled(bytes32 proofId, uint64 nonce);
    event RequestsFulfilled(bytes32[] proofIds, uint64[] nonces);
    event OpRequestsFulfilled(bytes32[] proofIds, uint64[] nonces, bytes32[] appCommitHashes, bytes32[] appVkHashes);

    event RequestRefunded(bytes32 proofId, uint64 nonce);
    event RequestCallbackFailed(bytes32 proofId, uint64 nonce);
    event RequestsCallbackFailed(bytes32[] proofIds, uint64[] nonces);
    event RequestFeeIncreased(bytes32 proofId, uint64 nonce, uint256 gas, uint256 fee);

    event RequestTimeoutUpdated(uint256 from, uint256 to);
    event BaseDataUrlUpdated(string from, string to);
    event BrevisProofUpdated(address from, address to);
    event BrevisDisputeUpdated(address from, address to);
    event BvnSigsVerifierUpdated(address from, address to);
    event AvsSigsVerifierUpdated(address from, address to);

    function sendRequest(
        bytes32 _proofId,
        uint64 _nonce,
        address _refundee,
        Callback calldata _callback,
        uint8 _option // bitmap 0: zk, 1: op bvn, 2: op avs, 3: op bvn and avs
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
