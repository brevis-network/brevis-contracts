// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "./FeeVault.sol";
import "../interface/IBrevisRequest.sol";
import "../interface/IBrevisProof.sol";
import "../interface/IBrevisApp.sol";
import "../../interfaces/ISigsVerifier.sol";
import "../lib/Lib.sol";

contract BrevisRequest is IBrevisRequest, FeeVault {
    // common workflow
    IBrevisProof public brevisProof;
    uint256 public requestTimeout;
    mapping(bytes32 => Request) public requests;

    // optimistic workflow
    ISigsVerifier public immutable sigsVerifier;
    uint256 public challengeWindow;
    uint256 public responseTimeout;
    uint256 public depositAskForData;
    uint256 public depositAskForProof;
    mapping(bytes32 => bytes32) public opdata; // keccak256(abi.encodePacked(appCommitHash, appVkHash));
    mapping(bytes32 => Dispute) public disputes;

    constructor(address _feeCollector, IBrevisProof _brevisProof, ISigsVerifier _sigsVerifier) FeeVault(_feeCollector) {
        brevisProof = _brevisProof;
        sigsVerifier = _sigsVerifier;
    }

    /*********************************
     * External and Public Functions *
     *********************************/

    function sendRequest(
        bytes32 _requestId,
        address _refundee,
        address _callback,
        uint64 _gas,
        RequestOption _option
    ) external payable {
        bytes32 requestKey = _requestId; // todo: keccak256(abi.encodePacked(_requestId, _nonce));
        require(requests[requestKey].status == RequestStatus.Null, "invalid request status");
        if (_refundee == address(0)) {
            _refundee = msg.sender;
        }
        RequestStatus status;
        if (_option == RequestOption.Zk) {
            status == RequestStatus.ZkPending;
        } else if (_option == RequestOption.Zk) {
            status = RequestStatus.OpPending;
        } else {
            revert("invalid request option");
        }
        bytes32 feeHash = keccak256(abi.encodePacked(msg.value, _refundee)); // todo: option to add fee
        Callback memory callback = Callback(_callback, _gas);
        requests[requestKey] = Request(uint64(block.timestamp), status, callback, feeHash);
        emit RequestSent(_requestId, _refundee, msg.value, _callback, _gas, _option);
    }

    function fulfillRequest(
        bytes32 _requestId,
        uint64 _chainId,
        bytes calldata _proof,
        bytes calldata _appCircuitOutput,
        address _callback
    ) external {
        (bytes32 requestId, bytes32 appCommitHash, bytes32 appVkHash) = IBrevisProof(brevisProof).submitProof(
            _chainId,
            _proof
        ); // revert for invalid proof
        require(_requestId == requestId, "requestId and proof not match");

        bytes32 requestKey = _requestId; // todo: keccak256(abi.encodePacked(_requestId, _nonce));
        Request storage request = requests[requestKey];
        RequestStatus status = request.status;
        require(status == RequestStatus.ZkPending || status == RequestStatus.Null, "invalid request status");
        request.status = RequestStatus.ZkAttested;

        if (_appCircuitOutput.length > 0) {
            require(appCommitHash == keccak256(_appCircuitOutput), "failed to open output commitment");
        }
        bool success = _brevisCallback(_callback, appVkHash, _appCircuitOutput, request, status);
        if (!success) {
            emit RequestCallbackFailed(_requestId);
        }
        emit RequestFulfilled(_requestId);
    }

    // fulfill batch requests with aggProof
    function fulfillRequests(
        bytes32[] calldata _requestIds,
        uint64 _chainId,
        bytes calldata _proof,
        Brevis.ProofData[] calldata _proofDataArray,
        bytes[] calldata _appCircuitOutputs,
        address[] calldata _callbacks
    ) external {
        IBrevisProof(brevisProof).mustSubmitAggProof(_chainId, _requestIds, _proof);
        if (_callbacks.length > 0) {
            IBrevisProof(brevisProof).mustValidateRequests(_chainId, _proofDataArray);
            require(
                _requestIds.length == _proofDataArray.length && _requestIds.length == _appCircuitOutputs.length,
                "length mismatch"
            );
            require(_callbacks.length == 1 || _callbacks.length == _requestIds.length, "length mismtach");
        }

        uint256 numFulfilled;
        for (uint256 i = 0; i < _requestIds.length; i++) {
            bytes32 requestKey = _requestIds[i]; // todo: keccak256(abi.encodePacked(_requestIds[i], _nonces[i]));
            Request storage request = requests[requestKey];
            RequestStatus status = request.status;
            if (status == RequestStatus.ZkPending || status == RequestStatus.Null) {
                request.status = RequestStatus.ZkAttested;
                numFulfilled++;

                if (_callbacks.length > 0) {
                    require(
                        _proofDataArray[i].appCommitHash == keccak256(_appCircuitOutputs[i]),
                        "failed to open output commitment"
                    );
                    if (_callbacks.length > 1) {
                        bool success = _brevisCallback(
                            _callbacks[i],
                            _proofDataArray[i].appVkHash,
                            _appCircuitOutputs[i],
                            request,
                            status
                        );
                        if (!success) {
                            emit RequestCallbackFailed(_requestIds[i]);
                        }
                    } else if (status == RequestStatus.ZkPending) {
                        require(request.callback.target == _callbacks[0], "callback mismatch");
                        require(request.callback.gas == 0, "invalid gas for batch callback");
                    }
                }
            }
        }
        if (_callbacks.length == 1) {
            bytes32[] memory appVkHashes = new bytes32[](_proofDataArray.length);
            for (uint256 i = 0; i < appVkHashes.length; i++) {
                appVkHashes[i] = _proofDataArray[i].appVkHash;
            }
            (bool success, ) = _callbacks[0].call(
                abi.encodeWithSelector(IBrevisApp.brevisBatchCallback.selector, appVkHashes, _appCircuitOutputs)
            );
            if (!success) {
                emit RequestsCallbackFailed(_requestIds);
            }
        }
        require(numFulfilled > 0, "no fulfilled requests");
        emit RequestsFulfilled(_requestIds);
    }

    function refund(bytes32 _requestId, uint256 _amount, address _refundee) external {
        bytes32 requestKey = _requestId; // todo: keccak256(abi.encodePacked(_requestId, _nonce));
        Request storage request = requests[requestKey];
        require(
            request.status == RequestStatus.ZkPending || request.status == RequestStatus.OpPending,
            "invalid request status"
        );
        require(block.timestamp > request.timestamp + requestTimeout);

        require(request.feeHash == keccak256(abi.encodePacked(_amount, _refundee)), "invalid input");
        (bool sent, ) = _refundee.call{value: _amount, gas: 50000}("");
        require(sent, "send native failed");
        request.status = RequestStatus.Refunded;
        emit RequestRefunded(_requestId);
    }

    // --------------------- optimistic workflow functions ---------------------

    function fulfillOpRequests(
        bytes32[] calldata _requestIds,
        bytes32[] calldata _appCommitHashes,
        bytes32[] calldata _appVkHashes,
        bytes[] calldata _dataURLs,
        bytes[] calldata _sigs,
        address[] calldata _signers,
        uint256[] calldata _powers
    ) external {
        uint256 dataNum = _requestIds.length;
        require(
            dataNum == _appCommitHashes.length && dataNum == _appVkHashes.length && dataNum == _dataURLs.length,
            "length mismatch"
        );

        bytes32 domain = keccak256(abi.encodePacked(block.chainid, address(this), "FulfillRequests"));
        sigsVerifier.verifySigs(
            abi.encodePacked(domain, _requestIds, _appCommitHashes, _appVkHashes), // todo: nonces, urls
            _sigs,
            _signers,
            _powers
        );

        uint256 timestamp = block.timestamp;
        for (uint i = 0; i < _requestIds.length; i++) {
            bytes32 requestKey = _requestIds[i]; // todo: keccak256(abi.encodePacked(_requestIds[i], _nonces[i]));
            Request storage request = requests[requestKey];
            require(
                request.status == RequestStatus.OpPending || request.status == RequestStatus.Null,
                "invalid request status"
            );
            request.status = RequestStatus.OpSubmitted;
            request.timestamp = uint64(timestamp);
            opdata[requestKey] = keccak256(abi.encodePacked(_appCommitHashes[i], _appVkHashes[i]));
        }

        emit OpRequestsFulfilled(_requestIds, _appCommitHashes, _appVkHashes, _dataURLs);
    }

    function askForRequestData(bytes32 _requestId) external payable {
        require(msg.value > depositAskForData, "insufficient deposit");
        // TODO: msg.value should be larger than a configurable value
        bytes32 requestKey = _requestId; // todo: keccak256(abi.encodePacked(_requestId, _nonce));
        Request storage request = requests[requestKey];
        Dispute storage dispute = disputes[requestKey];
        require(request.status == RequestStatus.OpSubmitted, "not in a disputable status");
        require(request.timestamp + challengeWindow > block.timestamp, "pass challenge window");

        request.status = RequestStatus.OpDisputing;
        dispute.status = DisputeStatus.WaitingForRequestData;
        dispute.responseDeadline = block.timestamp + responseTimeout;

        emit AskFor(_requestId, DisputeStatus.WaitingForRequestData, msg.sender);
    }

    function postRequestData(bytes32 _requestId, bytes calldata _requestData) external {
        bytes32 requestKey = _requestId; // todo: keccak256(abi.encodePacked(_requestId, _nonce));
        Request storage request = requests[requestKey];
        Dispute storage dispute = disputes[requestKey];
        require(request.status == RequestStatus.OpDisputing, "invalid request status");
        require(dispute.status == DisputeStatus.WaitingForRequestData, "invalid dispute status");

        disputes[requestKey].requestDataHash = keccak256(_requestData); // TODO: use claimed mimc?
        disputes[requestKey].status = DisputeStatus.RequestDataPosted;
        emit RequestDataPosted(_requestId);
    }

    function askForDataAvailabilityProof(bytes32 _requestId) external payable {
        require(msg.value > depositAskForProof, "insufficient deposit");
        bytes32 requestKey = _requestId; // todo: keccak256(abi.encodePacked(_requestId, _nonce));
        Request storage request = requests[requestKey];
        Dispute storage dispute = disputes[requestKey];
        require(
            request.status == RequestStatus.OpDisputing && dispute.status == DisputeStatus.RequestDataPosted,
            "invalid states"
        );
        require(request.timestamp + challengeWindow > block.timestamp, "pass challenge window");

        request.status = RequestStatus.OpDisputing;
        dispute.status = DisputeStatus.WaitingForDataAvailabilityProof;
        dispute.responseDeadline = block.timestamp + responseTimeout;

        emit AskFor(_requestId, DisputeStatus.WaitingForDataAvailabilityProof, msg.sender);
    }

    function postDataAvailabilityProof(bytes32 _requestId, bytes calldata _proof) external {
        bytes32 requestKey = _requestId; // todo: keccak256(abi.encodePacked(_requestId, _nonce));
        Request storage request = requests[requestKey];
        Dispute storage dispute = disputes[requestKey];
        require(
            request.status == RequestStatus.OpDisputing &&
                dispute.status == DisputeStatus.WaitingForDataAvailabilityProof,
            "invalid states"
        );
        disputes[requestKey].status = DisputeStatus.DataAvailabilityProofPosted;
        // TODO: check proof

        emit DataAvailabilityProofPosted(_requestId);
    }

    function askForDataValidityProof(bytes32 _requestId) external payable {
        require(msg.value > depositAskForProof, "insufficient deposit");
        bytes32 requestKey = _requestId; // todo: keccak256(abi.encodePacked(_requestId, _nonce));
        Request storage request = requests[requestKey];
        Dispute storage dispute = disputes[requestKey];
        require(
            request.status == RequestStatus.OpSubmitted ||
                (request.status == RequestStatus.OpDisputing &&
                    dispute.status != DisputeStatus.WaitingForDataValidityProof),
            "invalid states"
        );
        require(request.timestamp + challengeWindow > block.timestamp, "pass challenge window");

        request.status = RequestStatus.OpDisputing;
        dispute.status = DisputeStatus.WaitingForDataValidityProof;
        dispute.responseDeadline = block.timestamp + responseTimeout;

        emit AskFor(_requestId, DisputeStatus.WaitingForDataValidityProof, msg.sender);
    }

    function postDataValidityProof(bytes32 _requestId, uint64 _chainId, bytes calldata _proof) external {
        bytes32 requestKey = _requestId; // todo: keccak256(abi.encodePacked(_requestId, _nonce));
        Request storage request = requests[requestKey];
        Dispute storage dispute = disputes[requestKey];
        require(
            request.status == RequestStatus.OpDisputing && dispute.status == DisputeStatus.WaitingForDataValidityProof,
            "invalid states"
        );

        (bytes32 requestId, bytes32 appCommitHash, bytes32 appVkHash) = IBrevisProof(brevisProof).submitProof(
            _chainId,
            _proof
        );
        require(_requestId == requestId, "invalid proof: requestId");
        require(opdata[requestKey] == keccak256(abi.encodePacked(appCommitHash, appVkHash)), "invalid proof: appHash");
        request.status = RequestStatus.ZkAttested;
        dispute.status = DisputeStatus.DataValidityProofPosted;

        emit DataValidityProofProofPosted(_requestId);
    }

    // --------------------- owner functions ---------------------

    function setRequestTimeout(uint256 _timeout) external onlyOwner {
        uint256 oldTimeout = requestTimeout;
        requestTimeout = _timeout;
        emit RequestTimeoutUpdated(oldTimeout, _timeout);
    }

    function setChallengeWindow(uint256 _challengeWindow) external onlyOwner {
        uint256 oldChallengeWindow = challengeWindow;
        challengeWindow = _challengeWindow;
        emit ChallengeWindowUpdated(oldChallengeWindow, _challengeWindow);
    }

    function setDisputeDeposits(uint256 _amtAskForData, uint256 _amtAskForProof) external onlyOwner {
        depositAskForData = _amtAskForData;
        depositAskForProof = _amtAskForProof;
        emit DisputeDepositsUpdated(_amtAskForData, _amtAskForProof);
    }

    function setResponseTimeout(uint256 _responseTimeout) external onlyOwner {
        uint256 oldResponseTimeout = responseTimeout;
        responseTimeout = _responseTimeout;
        emit ResponseTimeoutUpdated(oldResponseTimeout, _responseTimeout);
    }

    /**************************
     *  Public View Functions *
     **************************/

    function queryRequestStatus(bytes32 _requestId) external view returns (RequestStatus) {
        return _queryRequestStatus(_requestId, challengeWindow);
    }

    function queryRequestStatus(bytes32 _requestId, uint256 _appChallengeWindow) external view returns (RequestStatus) {
        return _queryRequestStatus(_requestId, _appChallengeWindow);
    }

    function queryRequestTimestamp(bytes32 _requestId) external view returns (uint256) {
        bytes32 requestKey = _requestId; // todo: keccak256(abi.encodePacked(_requestId, _nonce));
        return requests[requestKey].timestamp;
    }

    function validateRequestOpData(
        bytes32 _requestId,
        bytes32 _appCommitHash,
        bytes32 _appVkHash
    ) external view returns (bool) {
        bytes32 requestKey = _requestId; // todo: keccak256(abi.encodePacked(_requestId, _nonce));
        return opdata[requestKey] == keccak256(abi.encodePacked(_appCommitHash, _appVkHash));
    }

    /*********************
     * Private Functions *
     *********************/

    function _brevisCallback(
        address _callback,
        bytes32 _appVkHash,
        bytes calldata _appCircuitOutput,
        Request storage _request,
        RequestStatus _status
    ) private returns (bool) {
        if (_status == RequestStatus.ZkPending) {
            require(_request.callback.target == _callback, "callback mismatch");
        }
        if (_callback != address(0)) {
            uint256 gas;
            if (_status == RequestStatus.ZkPending) {
                gas = _request.callback.gas;
            }
            if (gas == 0) {
                gas = gasleft();
            }
            // If the call failed due to insufficient gas, anyone can still call the app.applyBrevisProof directly to proceed
            (bool success, ) = _callback.call{gas: gas}(
                abi.encodeWithSelector(IBrevisApp.brevisCallback.selector, _appVkHash, _appCircuitOutput)
            );
            if (!success) {
                return false;
            }
        }
        return true;
    }

    function _queryRequestStatus(bytes32 _requestId, uint256 _challengeWindow) private view returns (RequestStatus) {
        bytes32 requestKey = _requestId; // todo: keccak256(abi.encodePacked(_requestId, _nonce));
        Request storage request = requests[requestKey];
        if (request.status == RequestStatus.OpSubmitted) {
            if (request.timestamp + _challengeWindow < block.timestamp) {
                return RequestStatus.OpAttested;
            }
        } else if (request.status == RequestStatus.OpDisputing) {
            Dispute storage dispute = disputes[requestKey];
            if (
                dispute.status == DisputeStatus.RequestDataPosted ||
                dispute.status == DisputeStatus.DataAvailabilityProofPosted
            ) {
                if (request.timestamp + _challengeWindow < block.timestamp) {
                    return RequestStatus.OpAttested;
                }
            } else if (dispute.responseDeadline < block.timestamp) {
                // WaitingForRequestData || WaitingForDataValidityProof || DataValidityProofPosted
                return RequestStatus.OpDisputed;
            }
        }
        return requests[requestKey].status;
    }
}
