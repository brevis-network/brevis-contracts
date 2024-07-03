// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/utils/Strings.sol";

import "./FeeVault.sol";
import "../interface/IBrevisRequest.sol";
import "../interface/IBrevisProof.sol";
import "../interface/IBrevisApp.sol";
import "../lib/Lib.sol";
import "../../safeguard/BrevisAccess.sol";
import "../../interfaces/ISigsVerifier.sol";

contract BrevisRequest is IBrevisRequest, FeeVault, BrevisAccess {
    // common workflow
    IBrevisProof public brevisProof;
    uint256 public requestTimeout;
    mapping(bytes32 => Request) public requests; // requestKey => Request;
    mapping(bytes32 => OnchainRequestInfo) public onchainRequests; // requestKey => OnchainRequestInfo

    // optimistic workflow
    IBvnSigsVerifier public bvnSigsVerifier;
    IAvsSigsVerifier public avsSigsVerifier;
    uint256 public challengeWindow;
    uint256 public responseTimeout;
    uint256 public depositAskForData;
    uint256 public depositAskForProof;
    string public baseDataURL;
    mapping(bytes32 => bytes32) public opdata; // requestKey => keccak256(abi.encodePacked(appCommitHash, appVkHash))
    mapping(bytes32 => Dispute) public disputes; // requestKey => Dispute

    uint8 constant OPT_IDX_SIG_BVN = 0;
    uint8 constant OPT_IDX_SIG_AVS = 1;

    constructor(
        address _feeCollector,
        IBrevisProof _brevisProof,
        IBvnSigsVerifier _bvnSigsVerifier
    ) FeeVault(_feeCollector) {
        brevisProof = _brevisProof;
        bvnSigsVerifier = _bvnSigsVerifier;
    }

    /*********************************
     * External and Public Functions *
     *********************************/

    function sendRequest(
        bytes32 _proofId,
        uint64 _nonce,
        address _refundee,
        Callback calldata _callback,
        uint8 _option
    ) external payable {
        bytes32 requestKey = keccak256(abi.encodePacked(_proofId, _nonce));
        RequestStatus status = requests[requestKey].status;
        require(status == RequestStatus.Null, "invalid status");
        _option = _option & 0x03; // binary 00000011
        if (_option == 0) {
            status == RequestStatus.ZkPending;
        } else {
            status = RequestStatus.OpPending;
        }
        requests[requestKey] = Request(status, uint64(block.timestamp), _option);

        if (_refundee == address(0)) {
            _refundee = msg.sender;
        }
        bytes32 feeHash = keccak256(abi.encodePacked(msg.value, _refundee));
        onchainRequests[requestKey] = OnchainRequestInfo(feeHash, _callback);

        emit RequestSent(_proofId, _nonce, _refundee, msg.value, _callback, _option);
    }

    function fulfillRequest(
        bytes32 _proofId,
        uint64 _nonce,
        uint64 _chainId,
        bytes calldata _proof,
        bytes calldata _appCircuitOutput,
        address _callbackTarget
    ) external onlyActiveProver {
        (bytes32 proofId, bytes32 appCommitHash, bytes32 appVkHash) = brevisProof.submitProof(_chainId, _proof);
        require(_proofId == proofId, "invalid proofId");

        bytes32 requestKey = keccak256(abi.encodePacked(_proofId, _nonce));
        Request storage request = requests[requestKey];
        RequestStatus status = request.status;
        require(status == RequestStatus.ZkPending || status == RequestStatus.Null, "invalid status");
        request.status = RequestStatus.ZkAttested;

        if (_appCircuitOutput.length > 0) {
            require(appCommitHash == keccak256(_appCircuitOutput), "invalid circuit output");
        }
        bool success = _brevisCallback(_callbackTarget, appVkHash, _appCircuitOutput, requestKey, status);
        if (!success) {
            emit RequestCallbackFailed(_proofId, _nonce);
        }
        emit RequestFulfilled(_proofId, _nonce);
    }

    // fulfill batch requests with aggProof
    function fulfillRequests(
        bytes32[] calldata _proofIds,
        uint64[] calldata _nonces,
        uint64 _chainId,
        bytes calldata _proof,
        Brevis.ProofData[] calldata _proofDataArray,
        bytes[] calldata _appCircuitOutputs,
        address[] calldata _callbackTargets
    ) external onlyActiveProver {
        uint256 dataNum = _proofIds.length;
        brevisProof.submitAggProof(_chainId, _proofIds, _proof);
        brevisProof.validateAggProofData(_chainId, _proofDataArray);
        require(_callbackTargets.length == 1 || _callbackTargets.length == dataNum, "invalid callback length");

        uint256 numFulfilled;
        for (uint256 i = 0; i < dataNum; i++) {
            require(_proofDataArray[i].commitHash == _proofIds[i], "invalid proofId");
            require(_proofDataArray[i].appCommitHash == keccak256(_appCircuitOutputs[i]), "invalid circuit output");

            bytes32 requestKey = keccak256(abi.encodePacked(_proofIds[i], _nonces[i]));
            Request storage request = requests[requestKey];
            RequestStatus status = request.status;
            if (status == RequestStatus.ZkPending || status == RequestStatus.Null) {
                request.status = RequestStatus.ZkAttested;
                numFulfilled++;
                if (_callbackTargets.length > 1) {
                    bool success = _brevisCallback(
                        _callbackTargets[i],
                        _proofDataArray[i].appVkHash,
                        _appCircuitOutputs[i],
                        requestKey,
                        status
                    );
                    if (!success) {
                        emit RequestCallbackFailed(_proofIds[i], _nonces[i]);
                    }
                } else if (status == RequestStatus.ZkPending) {
                    Callback memory callback = onchainRequests[requestKey].callback;
                    require(callback.target == _callbackTargets[0], "callback mismatch");
                    require(callback.gas == 0, "invalid gas for batch callback");
                }
            }
        }
        require(numFulfilled > 0, "no fulfilled requests");
        if (_callbackTargets.length == 1) {
            bytes32[] memory appVkHashes = new bytes32[](_proofDataArray.length);
            for (uint256 i = 0; i < appVkHashes.length; i++) {
                appVkHashes[i] = _proofDataArray[i].appVkHash;
            }
            // If the call failed due any reason, anyone can trigger retry later by calling applyBrevisAggProof.
            (bool success, ) = _callbackTargets[0].call(
                abi.encodeWithSelector(IBrevisApp.brevisBatchCallback.selector, appVkHashes, _appCircuitOutputs)
            );
            if (!success) {
                emit RequestsCallbackFailed(_proofIds, _nonces);
            }
        }
        emit RequestsFulfilled(_proofIds, _nonces);
    }

    function increaseGasFee(
        bytes32 _proofId,
        uint64 _nonce,
        uint64 _addGas,
        uint256 _currentFee,
        address _refundee
    ) external payable {
        bytes32 requestKey = keccak256(abi.encodePacked(_proofId, _nonce));
        RequestStatus status = requests[requestKey].status;
        require(status == RequestStatus.ZkPending || status == RequestStatus.OpPending, "invalid status");

        OnchainRequestInfo storage info = onchainRequests[requestKey];
        require(info.feeHash == keccak256(abi.encodePacked(_currentFee, _refundee)), "invalid input");
        uint256 newFee = _currentFee + msg.value;
        info.feeHash == keccak256(abi.encodePacked(newFee, _refundee));
        if (_addGas > 0) {
            info.callback.gas += _addGas;
        }
        emit RequestFeeIncreased(_proofId, _nonce, info.callback.gas, newFee);
    }

    function refund(bytes32 _proofId, uint64 _nonce, uint256 _amount, address _refundee) external {
        bytes32 requestKey = keccak256(abi.encodePacked(_proofId, _nonce));
        Request memory request = requests[requestKey];
        RequestStatus status = request.status;
        require(status == RequestStatus.ZkPending || status == RequestStatus.OpPending, "invalid status");
        require(block.timestamp > request.timestamp + requestTimeout);

        bytes32 feeHash = onchainRequests[requestKey].feeHash;
        require(feeHash == keccak256(abi.encodePacked(_amount, _refundee)), "invalid input");
        (bool sent, ) = _refundee.call{value: _amount, gas: 50000}("");
        require(sent, "send native failed");
        requests[requestKey].status = RequestStatus.Refunded;
        emit RequestRefunded(_proofId, _nonce);
    }

    // --------------------- optimistic workflow functions ---------------------

    function fulfillOpRequests(
        bytes32[] calldata _proofIds,
        uint64[] calldata _nonces,
        bytes32[] calldata _appCommitHashes,
        bytes32[] calldata _appVkHashes,
        IBvnSigsVerifier.SigInfo calldata _bvnSigInfo,
        IAvsSigsVerifier.SigInfo calldata _avsSigInfo
    ) external onlyActiveProver {
        uint8 option = 0;
        require(_bvnSigInfo.sigs.length > 0 || _avsSigInfo.blockNum > 0, "empty sigs");
        bytes32 domain = keccak256(abi.encodePacked(block.chainid, address(this), "FulfillRequests"));
        bytes32 signedHash = keccak256(abi.encodePacked(domain, _proofIds, _nonces, _appCommitHashes, _appVkHashes));
        if (_bvnSigInfo.sigs.length > 0) {
            bvnSigsVerifier.verifySigs(signedHash, _bvnSigInfo.sigs, _bvnSigInfo.signers, _bvnSigInfo.powers);
            option = _bitSet(option, OPT_IDX_SIG_BVN);
        }
        if (_avsSigInfo.blockNum > 0) {
            avsSigsVerifier.verifySigs(signedHash, _avsSigInfo.blockNum, _avsSigInfo.params);
            option = _bitSet(option, OPT_IDX_SIG_AVS);
        }
        _submitOpStates(_proofIds, _nonces, _appCommitHashes, _appVkHashes, option); // to avoid "stack too deep"
        emit OpRequestsFulfilled(_proofIds, _nonces, _appCommitHashes, _appVkHashes);
    }

    function askForRequestData(bytes32 _proofId, uint64 _nonce) external payable {
        require(msg.value > depositAskForData, "insufficient deposit");
        bytes32 requestKey = keccak256(abi.encodePacked(_proofId, _nonce));
        Request storage request = requests[requestKey];
        Dispute storage dispute = disputes[requestKey];
        require(request.status == RequestStatus.OpSubmitted, "not in a disputable status");
        require(request.timestamp + challengeWindow > block.timestamp, "pass challenge window");

        request.status = RequestStatus.OpDisputing;
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
        Request storage request = requests[requestKey];
        Dispute storage dispute = disputes[requestKey];
        require(request.status == RequestStatus.OpDisputing, "invalid request status");
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
        dispute.challenger = msg.sender;
        dispute.deposit = msg.value;

        emit AskFor(_proofId, _nonce, DisputeStatus.WaitingForDataAvailabilityProof, msg.sender);
    }

    function postDataAvailabilityProof(
        bytes32 _proofId,
        uint64 _nonce,
        bytes calldata /*proof*/
    ) external onlyActiveProver {
        bytes32 requestKey = keccak256(abi.encodePacked(_proofId, _nonce));
        Request storage request = requests[requestKey];
        Dispute storage dispute = disputes[requestKey];
        require(
            request.status == RequestStatus.OpDisputing &&
                dispute.status == DisputeStatus.WaitingForDataAvailabilityProof,
            "invalid states"
        );
        disputes[requestKey].status = DisputeStatus.DataAvailabilityProofPosted;
        // todo: validate proof

        emit DataAvailabilityProofPosted(_proofId, _nonce);
    }

    function askForDataValidityProof(bytes32 _proofId, uint64 _nonce) external payable {
        require(msg.value > depositAskForProof, "insufficient deposit");
        bytes32 requestKey = keccak256(abi.encodePacked(_proofId, _nonce));
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
        Request storage request = requests[requestKey];
        Dispute storage dispute = disputes[requestKey];
        require(
            request.status == RequestStatus.OpDisputing && dispute.status == DisputeStatus.WaitingForDataValidityProof,
            "invalid states"
        );

        (bytes32 proofId, bytes32 appCommitHash, bytes32 appVkHash) = brevisProof.submitProof(_chainId, _proof);
        require(_proofId == proofId, "invalid proof: proofId");
        require(opdata[requestKey] == keccak256(abi.encodePacked(appCommitHash, appVkHash)), "invalid proof: appHash");
        request.status = RequestStatus.ZkAttested;
        dispute.status = DisputeStatus.DataValidityProofPosted;

        emit DataValidityProofProofPosted(_proofId, _nonce);
    }

    // --------------------- app helper functions ---------------------

    // apply proved data to a target contract
    function applyBrevisProof(
        bytes32 _proofId,
        bytes32 _appVkHash,
        bytes32 _appCommitHash,
        bytes calldata _appCircuitOutput,
        address _callbackTarget
    ) external {
        brevisProof.validateProofAppData(_proofId, _appCommitHash, _appVkHash);
        require(_appCommitHash == keccak256(_appCircuitOutput), "invalid circuit output");
        IBrevisApp(_callbackTarget).brevisCallback(_appVkHash, _appCircuitOutput);
    }

    // apply multiple proved data fulfilled through AggProof to a target contract
    function applyBrevisAggProof(
        uint64 _chainId,
        Brevis.ProofData[] calldata _proofDataArray,
        bytes[] calldata _appCircuitOutputs,
        address[] calldata _callbackTargets
    ) external {
        uint256 dataNum = _proofDataArray.length;
        brevisProof.validateAggProofData(_chainId, _proofDataArray);
        for (uint256 i = 0; i < dataNum; i++) {
            if (_callbackTargets[i] != address(0)) {
                require(_proofDataArray[i].appCommitHash == keccak256(_appCircuitOutputs[i]), "invalid circuit output");
                (bool success, bytes memory res) = _callbackTargets[i].call(
                    abi.encodeWithSelector(
                        IBrevisApp.brevisCallback.selector,
                        _proofDataArray[i].appVkHash,
                        _appCircuitOutputs[i]
                    )
                );
                require(success, string.concat(Strings.toString(i), ":", Utils.getRevertMsg(res)));
            }
        }
    }

    // apply single proved data fulfilled through AggProof to a target contract
    function applyBrevisAggProof(
        uint64 _chainId,
        Brevis.ProofData calldata _proofData,
        bytes32 _merkleRoot,
        bytes32[] calldata _merkleProof,
        uint8 _nodeIndex,
        bytes calldata _appCircuitOutput,
        address _callbackTarget
    ) external {
        brevisProof.validateAggProofData(_chainId, _proofData, _merkleRoot, _merkleProof, _nodeIndex);
        require(_proofData.appCommitHash == keccak256(_appCircuitOutput), "invalid circuit output");
        IBrevisApp(_callbackTarget).brevisCallback(_proofData.appVkHash, _appCircuitOutput);
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

    function setDisputeDeposits(uint256 _amtAskForData, uint256 _amtAskForProof) external onlyOwner {
        depositAskForData = _amtAskForData;
        depositAskForProof = _amtAskForProof;
        emit DisputeDepositsUpdated(_amtAskForData, _amtAskForProof);
    }

    function setBaseDataURL(string memory _url) external onlyOwner {
        string memory oldURL = baseDataURL;
        baseDataURL = _url;
        emit BaseDataUrlUpdated(oldURL, _url);
    }

    function setBrevisProof(address _brevisProof) external onlyOwner {
        address oldAddr = address(brevisProof);
        brevisProof = IBrevisProof(_brevisProof);
        emit BrevisProofUpdated(oldAddr, _brevisProof);
    }

    function setBvnSigsVerifier(address _bvnSigsVerifier) external onlyOwner {
        address oldAddr = address(bvnSigsVerifier);
        bvnSigsVerifier = IBvnSigsVerifier(_bvnSigsVerifier);
        emit BvnSigsVerifierUpdated(oldAddr, _bvnSigsVerifier);
    }

    function setAvsSigsVerifier(address _avsSigsVerifier) external onlyOwner {
        address oldAddr = address(avsSigsVerifier);
        avsSigsVerifier = IAvsSigsVerifier(_avsSigsVerifier);
        emit BvnSigsVerifierUpdated(oldAddr, _avsSigsVerifier);
    }

    /**************************
     *  Public View Functions *
     **************************/

    function queryRequestStatus(bytes32 _proofId, uint64 _nonce) external view returns (RequestStatus, uint8) {
        return _queryRequestStatus(_proofId, _nonce, challengeWindow);
    }

    function queryRequestStatus(
        bytes32 _proofId,
        uint64 _nonce,
        uint256 _appChallengeWindow
    ) external view returns (RequestStatus, uint8) {
        return _queryRequestStatus(_proofId, _nonce, _appChallengeWindow);
    }

    function validateOpAppData(
        bytes32 _proofId,
        uint64 _nonce,
        bytes32 _appCommitHash,
        bytes32 _appVkHash,
        uint8 _option
    ) external view returns (bool) {
        return _validateOpAppData(_proofId, _nonce, _appCommitHash, _appVkHash, challengeWindow, _option);
    }

    function validateOpAppData(
        bytes32 _proofId,
        uint64 _nonce,
        bytes32 _appCommitHash,
        bytes32 _appVkHash,
        uint256 _appChallengeWindow,
        uint8 _option
    ) external view returns (bool) {
        return _validateOpAppData(_proofId, _nonce, _appCommitHash, _appVkHash, _appChallengeWindow, _option);
    }

    function dataURL(bytes32 _proofId) external view returns (string memory) {
        if (bytes(baseDataURL).length == 0) {
            return "";
        }
        return string.concat(baseDataURL, Strings.toHexString(uint256(_proofId), 32));
    }

    /*********************
     * Private Functions *
     *********************/

    function _brevisCallback(
        address _callbackTarget,
        bytes32 _appVkHash,
        bytes calldata _appCircuitOutput,
        bytes32 _requestKey,
        RequestStatus _status
    ) private returns (bool) {
        uint256 gas;
        if (_status == RequestStatus.ZkPending /* is onchain request*/) {
            Callback memory callback = onchainRequests[_requestKey].callback;
            require(callback.target == _callbackTarget, "callback mismatch");
            gas = callback.gas;
        }
        if (_callbackTarget != address(0)) {
            if (gas == 0) {
                gas = gasleft();
            }
            // If the call failed due any reason, anyone can trigger retry later by calling applyBrevisProof.
            (bool success, ) = _callbackTarget.call{gas: gas}(
                abi.encodeWithSelector(IBrevisApp.brevisCallback.selector, _appVkHash, _appCircuitOutput)
            );
            if (!success) {
                return false;
            }
        }
        return true;
    }

    function _submitOpStates(
        bytes32[] calldata _proofIds,
        uint64[] calldata _nonces,
        bytes32[] calldata _appCommitHashes,
        bytes32[] calldata _appVkHashes,
        uint8 option
    ) private {
        uint64 timestamp = uint64(block.timestamp);
        for (uint i = 0; i < _proofIds.length; i++) {
            bytes32 requestKey = keccak256(abi.encodePacked(_proofIds[i], _nonces[i]));
            Request memory request = requests[requestKey];
            if (request.status == RequestStatus.OpPending) {
                require(request.option == option, "insufficient types of sigs");
            } else {
                require(request.status == RequestStatus.Null, "invalid status");
            }
            requests[requestKey] = Request(RequestStatus.OpSubmitted, timestamp, option);
            opdata[requestKey] = keccak256(abi.encodePacked(_appCommitHashes[i], _appVkHashes[i]));
        }
    }

    function _queryRequestStatus(
        bytes32 _proofId,
        uint64 _nonce,
        uint256 _challengeWindow
    ) private view returns (RequestStatus, uint8) {
        bytes32 requestKey = keccak256(abi.encodePacked(_proofId, _nonce));
        Request memory request = requests[requestKey];
        if (request.status == RequestStatus.OpSubmitted) {
            if (request.timestamp + _challengeWindow < block.timestamp) {
                return (RequestStatus.OpAttested, request.option);
            }
        } else if (request.status == RequestStatus.OpDisputing) {
            Dispute storage dispute = disputes[requestKey];
            DisputeStatus dstatus = dispute.status;
            if (dstatus == DisputeStatus.RequestDataPosted || dstatus == DisputeStatus.DataAvailabilityProofPosted) {
                if (request.timestamp + _challengeWindow < block.timestamp) {
                    return (RequestStatus.OpAttested, request.option);
                }
            } else if (dispute.responseDeadline < block.timestamp) {
                // did not respond in time for WaitingForXXX
                return (RequestStatus.OpDisputed, request.option);
            }
        }
        return (request.status, request.option);
    }

    function _validateOpAppData(
        bytes32 _proofId,
        uint64 _nonce,
        bytes32 _appCommitHash,
        bytes32 _appVkHash,
        uint256 _challengeWindow,
        uint8 _option
    ) private view returns (bool readyToUse) {
        bytes32 requestKey = keccak256(abi.encodePacked(_proofId, _nonce));
        require(opdata[requestKey] == keccak256(abi.encodePacked(_appCommitHash, _appVkHash)), "invalid data");
        (RequestStatus status, uint8 option) = _queryRequestStatus(_proofId, _nonce, _challengeWindow);
        if (status == RequestStatus.OpAttested) {
            return (_option & option) == _option;
        } else if (status == RequestStatus.ZkAttested) {
            return true;
        }
        return false;
    }

    function _bitSet(uint8 _value, uint8 _index) private pure returns (uint8) {
        uint8 mask = uint8(1 << _index);
        return _value | mask;
    }
}
