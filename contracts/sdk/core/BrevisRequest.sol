// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "./FeeVault.sol";
import "../interface/IBrevisRequest.sol";
import "../interface/IBrevisProof.sol";
import "../interface/IBrevisApp.sol";
import "../../interfaces/ISigsVerifier.sol";
import "../lib/Lib.sol";

contract BrevisRequest is IBrevisRequest, FeeVault {
    IBrevisProof public brevisProof;
    ISigsVerifier public immutable sigsVerifier;

    uint256 public requestTimeout;
    uint256 public challengeWindow;
    uint256 public responseTimeout;

    mapping(bytes32 => Request) public requests;
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
        Fee memory fee = Fee(msg.value, _refundee);
        Callback memory callback = Callback(_callback, 0); // todo: support gas limit
        requests[requestKey] = Request(uint64(block.timestamp), status, fee, callback);
        emit RequestSent(_requestId, msg.sender, msg.value, _callback, _option);
    }

    // fulfill onchain request
    function fulfillRequest(
        bytes32 _requestId,
        uint64 _chainId,
        bytes calldata _proof,
        bytes calldata _appCircuitOutput
    ) external {
        _fulfillRequest(_requestId, _chainId, _proof, _appCircuitOutput, address(0), RequestStatus.ZkPending);
    }

    // fulfill offchain request
    function fulfillRequest(
        bytes32 _requestId,
        uint64 _chainId,
        bytes calldata _proof,
        bytes calldata _appCircuitOutput,
        address callback
    ) external {
        _fulfillRequest(_requestId, _chainId, _proof, _appCircuitOutput, callback, RequestStatus.Null);
    }

    // fulfill aggregated requests
    function fulfillRequests(
        bytes32[] calldata _requestIds,
        uint64 _chainId,
        bytes calldata _proof,
        Brevis.ProofData[] calldata _proofDataArray,
        bytes[] calldata _appCircuitOutputs,
        address[] calldata callbacks
    ) external {
        _fulfillRequests(_requestIds, _chainId, _proof, _proofDataArray, _appCircuitOutputs, callbacks, address(0));
    }

    // fulfill aggregated requests with batchCallback
    function fulfillRequests(
        bytes32[] calldata _requestIds,
        uint64 _chainId,
        bytes calldata _proof,
        Brevis.ProofData[] calldata _proofDataArray,
        bytes[] calldata _appCircuitOutputs,
        address batchCallback
    ) external {
        address[] memory empty;
        _fulfillRequests(_requestIds, _chainId, _proof, _proofDataArray, _appCircuitOutputs, empty, batchCallback);
    }

    function refund(bytes32 _requestId) external {
        // TODO: refund for op request
        bytes32 requestKey = _requestId; // todo: keccak256(abi.encodePacked(_requestId, _nonce));
        Request storage request = requests[requestKey];
        require(
            request.status == RequestStatus.ZkPending || request.status == RequestStatus.OpPending,
            "invalid request status"
        );
        require(block.timestamp > request.timestamp + requestTimeout);
        require(!IBrevisProof(brevisProof).hasProof(_requestId), "proof already generated");
        (bool sent, ) = request.fee.refundee.call{value: request.fee.amount, gas: 50000}("");
        require(sent, "send native failed");
        request.status = RequestStatus.Refunded;
        emit RequestRefunded(_requestId);
    }

    // --------------------- optimistic workflow functions ---------------------

    function fulfillOpRequests(
        bytes32[] calldata _requestIds,
        bytes[] calldata _dataURLs,
        bytes[] calldata _sigs,
        address[] calldata _signers,
        uint256[] calldata _powers
    ) external {
        require(_requestIds.length > 0, "invalid requestIds");
        require(_requestIds.length == _dataURLs.length);

        bytes32 domain = keccak256(abi.encodePacked(block.chainid, address(this), "FulfillRequests"));
        sigsVerifier.verifySigs(abi.encodePacked(domain, _requestIds), _sigs, _signers, _powers); // todo: nonces

        uint256 timestamp = block.timestamp;
        for (uint i = 0; i < _requestIds.length; i++) {
            bytes32 requestKey = _requestIds[i]; // todo: keccak256(abi.encodePacked(_requestIds[i], _nonces[i]));
            Request storage request = requests[requestKey];
            require(request.status == RequestStatus.OpPending, "invalid request status");
            request.status = RequestStatus.OpSubmitted;
            request.timestamp = uint64(timestamp);
        }

        emit OpRequestsFulfilled(_requestIds, _dataURLs);
    }

    function askForRequestData(bytes32 _requestId) external payable {
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
        // TODO: msg.value should be larger than a configurable value
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
        // TODO: msg.value should be larger than a configurable value
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
        (bytes32 requestId, , ) = IBrevisProof(brevisProof).submitProof(_chainId, _proof);
        require(_requestId == requestId, "requestId and proof not match");

        bytes32 requestKey = _requestId; // todo: keccak256(abi.encodePacked(_requestId, _nonce));
        requests[requestKey].status = RequestStatus.ZkAttested;
        disputes[requestKey].status = DisputeStatus.DataValidityProofPosted;

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

    /*********************
     * Private Functions *
     *********************/

    function _fulfillRequest(
        bytes32 _requestId,
        uint64 _chainId,
        bytes calldata _proof,
        bytes calldata _appCircuitOutput,
        address _callback,
        RequestStatus _expStatus
    ) private {
        (bytes32 requestId, bytes32 appCommitHash, bytes32 appVkHash) = IBrevisProof(brevisProof).submitProof(
            _chainId,
            _proof
        ); // revert for invalid proof
        require(_requestId == requestId, "requestId and proof not match");

        bytes32 requestKey = _requestId; // todo: keccak256(abi.encodePacked(_requestId, _nonce));
        Request storage request = requests[requestKey];
        require(request.status == _expStatus, "invalid request status");
        request.status = RequestStatus.ZkAttested;

        emit RequestFulfilled(_requestId);

        if (_expStatus == RequestStatus.ZkPending) {
            _callback = request.callback.target;
        }
        if (_callback != address(0)) {
            require(appCommitHash == keccak256(_appCircuitOutput), "failed to open output commitment");
            uint256 gas;
            if (_expStatus == RequestStatus.ZkPending) {
                gas = request.callback.gas;
            }
            if (gas == 0) {
                gas = gasleft();
            }
            // If the call failed due to insufficient gas, anyone can still call the app.applyBrevisProof directly to proceed
            (bool success, ) = _callback.call{gas: gas}(
                abi.encodeWithSelector(IBrevisApp.brevisCallback.selector, _requestId, appVkHash, _appCircuitOutput)
            );
            if (!success) {
                emit RequestCallbackFailed(_requestId);
            }
        }
    }

    function _fulfillRequests(
        bytes32[] calldata _requestIds,
        uint64 _chainId,
        bytes calldata _proof,
        Brevis.ProofData[] calldata _proofDataArray,
        bytes[] calldata _appCircuitOutputs,
        address[] memory _callbacks,
        address _batchCallback
    ) private {
        IBrevisProof(brevisProof).mustSubmitAggProof(_chainId, _requestIds, _proof);
        bool checkAppData;
        if (_proofDataArray.length > 0 && _batchCallback != address(0)) {
            checkAppData = true;
            IBrevisProof(brevisProof).mustValidateRequests(_chainId, _proofDataArray);
            require(_proofDataArray.length == _appCircuitOutputs.length, "length mismatch");
        }

        uint256 numFulfilled;
        for (uint256 i = 0; i < _requestIds.length; i++) {
            bytes32 requestKey = _requestIds[i]; // todo: keccak256(abi.encodePacked(_requestIds[i], _nonces[i]));
            Request storage request = requests[requestKey];
            RequestStatus status = request.status;
            if (status == RequestStatus.ZkPending || status == RequestStatus.Null) {
                request.status = RequestStatus.ZkAttested;
                numFulfilled++;

                if (checkAppData) {
                    require(
                        _proofDataArray[i].appCommitHash == keccak256(_appCircuitOutputs[i]),
                        "failed to open output commitment"
                    );
                    address callback = _callbacks[i];
                    if (status == RequestStatus.ZkPending) {
                        require(request.callback.target == callback, "mismatch callback");
                    }

                    if (callback != address(0)) {
                        uint256 gas;
                        if (status == RequestStatus.ZkPending) {
                            gas = request.callback.gas;
                        }
                        if (gas == 0) {
                            gas = gasleft();
                        }
                        (bool success, ) = request.callback.target.call{gas: gas}(
                            abi.encodeWithSelector(
                                IBrevisApp.brevisCallback.selector,
                                _requestIds[i],
                                _proofDataArray[i].appVkHash,
                                _appCircuitOutputs[i]
                            )
                        );
                        if (!success) {
                            emit RequestCallbackFailed(_requestIds[i]);
                        }
                    }
                } else if (_batchCallback != address(0) && status == RequestStatus.ZkPending) {
                    require(request.callback.target == _batchCallback, "mismatch callback");
                    require(request.callback.gas == 0, "nozero callback gas");
                }
            }
        }
        if (_batchCallback != address(0)) {
            (bool success, ) = _callbacks[0].call(
                abi.encodeWithSelector(
                    IBrevisApp.brevisBatchCallback.selector,
                    _chainId,
                    _proofDataArray,
                    _appCircuitOutputs
                )
            );
            if (!success) {
                emit RequestsCallbackFailed(_requestIds);
            }
        }
        require(numFulfilled > 0, "no fulfilled requests");
        emit RequestsFulfilled(_requestIds);
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
