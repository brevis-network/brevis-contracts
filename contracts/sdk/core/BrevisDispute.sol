// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "./BrevisRequest.sol";
import "../interface/IBrevisDispute.sol";

// TODO: handle dispute fee
contract BrevisDispute is IBrevisDispute {
    BrevisRequest immutable brevisRequest;

    uint256 public challengeWindow;
    uint256 public responseTimeout;
    uint256 public depositAskForData;
    uint256 public depositAskForProof;

    mapping(bytes32 => Dispute) public disputes; // requestKey => Dispute

    modifier onlyActiveProver() {
        require(brevisRequest.isActiveProver(msg.sender), "invalid prover");
        _;
    }

    modifier onlyOwner() {
        require(brevisRequest.owner() == msg.sender, "invalid prover");
        _;
    }

    constructor(BrevisRequest _brevisRequest) {
        brevisRequest = _brevisRequest;
    }

    function askForRequestData(bytes32 _proofId, uint64 _nonce) external payable {
        require(msg.value > depositAskForData, "insufficient deposit");

        bytes32 requestKey = keccak256(abi.encodePacked(_proofId, _nonce));
        (RequestStatus status, uint64 timestamp, ) = brevisRequest.requests(requestKey);
        require(status == RequestStatus.OpSubmitted, "not in a disputable status");
        require(timestamp + challengeWindow > block.timestamp, "pass challenge window");
        brevisRequest.setRequestStatus(requestKey, RequestStatus.OpDisputing);

        Dispute storage dispute = disputes[requestKey];
        dispute.status = DisputeStatus.WaitingForRequestData;
        dispute.responseDeadline = block.timestamp + responseTimeout;
        dispute.challenger = msg.sender;
        dispute.deposit = msg.value;

        emit AskFor(_proofId, _nonce, DisputeStatus.WaitingForRequestData, msg.sender);
    }

    function postRequestData(
        bytes32 _proofId,
        uint64 _nonce,
        bytes[] calldata _data,
        uint256 _index,
        bool _done
    ) external onlyActiveProver {
        bytes32 requestKey = keccak256(abi.encodePacked(_proofId, _nonce));
        (RequestStatus status, , ) = brevisRequest.requests(requestKey);
        require(status == RequestStatus.OpDisputing, "invalid request status");

        Dispute storage dispute = disputes[requestKey];
        require(dispute.status == DisputeStatus.WaitingForRequestData, "invalid dispute status");
        RequestDataHash storage dataHash = dispute.requestDataHash;
        require(dataHash.hashes.length == _index, "invalid index");
        for (uint i = 0; i < _data.length; i++) {
            dataHash.hashes.push(keccak256(_data[i]));
        }
        if (_done) {
            dataHash.root = keccak256(abi.encodePacked(dataHash.hashes)); // todo: consider merkle
            disputes[requestKey].status = DisputeStatus.RequestDataPosted;
        }
        emit RequestDataPosted(_proofId, _nonce, _data, _index, _done);
    }

    function askForDataAvailabilityProof(bytes32 _proofId, uint64 _nonce) external payable {
        require(msg.value > depositAskForProof, "insufficient deposit");
        bytes32 requestKey = keccak256(abi.encodePacked(_proofId, _nonce));
        (RequestStatus status, uint64 timestamp, ) = brevisRequest.requests(requestKey);
        require(timestamp + challengeWindow > block.timestamp, "pass challenge window");
        brevisRequest.setRequestStatus(requestKey, RequestStatus.OpDisputing);

        Dispute storage dispute = disputes[requestKey];
        require(
            status == RequestStatus.OpDisputing && dispute.status == DisputeStatus.RequestDataPosted,
            "invalid states"
        );

        dispute.status = DisputeStatus.WaitingForDataAvailabilityProof;
        dispute.responseDeadline = block.timestamp + responseTimeout;
        dispute.challenger = msg.sender;
        dispute.deposit = msg.value;

        emit AskFor(_proofId, _nonce, DisputeStatus.WaitingForDataAvailabilityProof, msg.sender);
    }

    function postDataAvailabilityProof(
        bytes32 _proofId,
        uint64 _nonce,
        bytes calldata // proof
    ) external onlyActiveProver {
        bytes32 requestKey = keccak256(abi.encodePacked(_proofId, _nonce));
        (RequestStatus status, , ) = brevisRequest.requests(requestKey);
        Dispute storage dispute = disputes[requestKey];
        require(
            status == RequestStatus.OpDisputing && dispute.status == DisputeStatus.WaitingForDataAvailabilityProof,
            "invalid states"
        );
        disputes[requestKey].status = DisputeStatus.DataAvailabilityProofPosted;
        // todo: validate proof

        emit DataAvailabilityProofPosted(_proofId, _nonce);
    }

    function askForDataValidityProof(bytes32 _proofId, uint64 _nonce) external payable {
        require(msg.value > depositAskForProof, "insufficient deposit");

        bytes32 requestKey = keccak256(abi.encodePacked(_proofId, _nonce));
        (RequestStatus status, uint64 timestamp, ) = brevisRequest.requests(requestKey);
        require(timestamp + challengeWindow > block.timestamp, "pass challenge window");
        brevisRequest.setRequestStatus(requestKey, RequestStatus.OpDisputing);

        Dispute storage dispute = disputes[requestKey];
        require(
            status == RequestStatus.OpSubmitted ||
                (status == RequestStatus.OpDisputing && dispute.status != DisputeStatus.WaitingForDataValidityProof),
            "invalid states"
        );

        dispute.status = DisputeStatus.WaitingForDataValidityProof;
        dispute.responseDeadline = block.timestamp + responseTimeout;
        dispute.challenger = msg.sender;
        dispute.deposit = msg.value;

        emit AskFor(_proofId, _nonce, DisputeStatus.WaitingForDataValidityProof, msg.sender);
    }

    function postDataValidityProof(
        bytes32 _proofId,
        uint64 _nonce,
        uint64 _chainId,
        bytes calldata _proof
    ) external onlyActiveProver {
        bytes32 requestKey = keccak256(abi.encodePacked(_proofId, _nonce));
        (RequestStatus status, , ) = brevisRequest.requests(requestKey);
        Dispute storage dispute = disputes[requestKey];
        require(
            status == RequestStatus.OpDisputing && dispute.status == DisputeStatus.WaitingForDataValidityProof,
            "invalid states"
        );

        (bytes32 proofId, bytes32 appCommitHash, bytes32 appVkHash) = brevisRequest.brevisProof().submitProof(
            _chainId,
            _proof
        );
        require(_proofId == proofId, "invalid proof: proofId");
        require(
            brevisRequest.opdata(requestKey) == keccak256(abi.encodePacked(appCommitHash, appVkHash)),
            "invalid proof: appHash"
        );
        brevisRequest.setRequestStatus(requestKey, RequestStatus.ZkAttested);
        dispute.status = DisputeStatus.DataValidityProofPosted;

        emit DataValidityProofProofPosted(_proofId, _nonce);
    }

    function getChallengeWindow() external view returns (uint256) {
        return challengeWindow;
    }

    function getDisputeStatus(bytes32 _requestKey) external view returns (DisputeStatus) {
        return disputes[_requestKey].status;
    }

    function getResponseDeadline(bytes32 _requestKey) external view returns (uint256) {
        return disputes[_requestKey].responseDeadline;
    }

    // --------------------- owner functions ---------------------

    function setChallengeWindow(uint256 _challengeWindow) external onlyOwner {
        uint256 oldChallengeWindow = challengeWindow;
        challengeWindow = _challengeWindow;
        emit ChallengeWindowUpdated(oldChallengeWindow, _challengeWindow);
    }

    function setResponseTimeout(uint256 _responseTimeout) external onlyOwner {
        uint256 oldResponseTimeout = responseTimeout;
        responseTimeout = _responseTimeout;
        emit ResponseTimeoutUpdated(oldResponseTimeout, _responseTimeout);
    }

    function setDisputeDeposits(uint256 _amtAskForData, uint256 _amtAskForProof) external onlyOwner {
        depositAskForData = _amtAskForData;
        depositAskForProof = _amtAskForProof;
        emit DisputeDepositsUpdated(_amtAskForData, _amtAskForProof);
    }
}
