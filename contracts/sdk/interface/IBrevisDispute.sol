// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "./IBrevisTypes.sol";

interface IBrevisDispute is IBrevisTypes {
    event AskFor(bytes32 indexed proofId, uint64 nonce, DisputeStatus status, address from);
    event RequestDataPosted(bytes32 indexed proofId, uint64 nonce, bytes[] data, uint256 index, bool done);
    event DataAvailabilityProofPosted(bytes32 indexed proofId, uint64 nonce);
    event DataValidityProofProofPosted(bytes32 indexed proofId, uint64 nonce);

    event ChallengeWindowUpdated(uint256 from, uint256 to);
    event ResponseTimeoutUpdated(uint256 from, uint256 to);
    event DisputeDepositsUpdated(uint256 amtAskForData, uint256 amtAskForProof);

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

    function getChallengeWindow() external view returns (uint256);

    function getDisputeStatus(bytes32 _requestKey) external view returns (DisputeStatus);

    function getResponseDeadline(bytes32 _requestKey) external view returns (uint256);
}
