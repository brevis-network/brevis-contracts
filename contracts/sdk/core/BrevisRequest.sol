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
    uint256 public challengeWindow; // in seconds
    uint256 public responseTimeout; // BVN responses an ask-for-data request

    mapping(bytes32 => Request) public requests; // TODO: store hash of request data to save gas cost
    mapping(bytes32 => RequestExt) public requestExts;
    mapping(bytes32 => bytes32) public keccakToMimc;

    constructor(address _feeCollector, IBrevisProof _brevisProof, ISigsVerifier _sigsVerifier) FeeVault(_feeCollector) {
        brevisProof = _brevisProof;
        sigsVerifier = _sigsVerifier;
    }

    /*********************************
     * External and Public Functions *
     *********************************/

    function sendRequest(bytes32 _requestId, address _refundee, address _callback, Option _option) external payable {
        require(requests[_requestId].deadline == 0, "request already in queue");
        require(_refundee != address(0), "refundee not provided");
        requests[_requestId] = Request(
            block.timestamp + requestTimeout,
            msg.value,
            _refundee,
            _callback,
            RequestStatus.Pending,
            _option
        );
        emit RequestSent(_requestId, msg.sender, msg.value, _callback, _option);
    }

    function fulfillRequest(
        bytes32 _requestId,
        uint64 _chainId,
        bytes calldata _proof,
        bytes calldata _appCircuitOutput
    ) external {
        require(!IBrevisProof(brevisProof).hasProof(_requestId), "proof already generated");

        bytes32 reqIdFromProof = IBrevisProof(brevisProof).submitProof(_chainId, _proof); // revert for invalid proof
        require(_requestId == reqIdFromProof, "requestId and proof not match");
        requests[_requestId].status = RequestStatus.ZkAttested;

        emit RequestFulfilled(_requestId);

        address app = address(requests[_requestId].callback);
        if (app != address(0)) {
            // The relayer should set correct gas limit. If the call failed due to insufficient gasleft(),
            // anyone can still call the app.brevisCallback directly to proceed
            (bool success, ) = app.call(
                abi.encodeWithSelector(IBrevisApp.brevisCallback.selector, _requestId, _appCircuitOutput)
            );
            if (!success) {
                emit RequestCallbackFailed(_requestId);
            }
        }
    }

    function fulfillAggRequests(
        uint64 _chainId,
        bytes32[] calldata _requestIds,
        bytes calldata _proof,
        Brevis.ProofData[] calldata _proofDataArray,
        bytes[] calldata _appCircuitOutputs,
        address _callback
    ) external {
        IBrevisProof(brevisProof).mustSubmitAggProof(_chainId, _requestIds, _proof);

        for (uint8 i = 1; i < _requestIds.length; i++) {
            bytes32 requestId = _requestIds[i];
            requests[requestId].status = RequestStatus.ZkAttested;
        }

        emit RequestsFulfilled(_requestIds);

        if (_callback != address(0)) {
            (bool success, ) = _callback.call(
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
    }

    function refund(bytes32 _requestId) external {
        require(block.timestamp > requests[_requestId].deadline);
        require(!IBrevisProof(brevisProof).hasProof(_requestId), "proof already generated");
        require(requests[_requestId].deadline != 0, "request not in queue");
        requests[_requestId].deadline = 0; //reset deadline, then user is able to send request again
        (bool sent, ) = requests[_requestId].refundee.call{value: requests[_requestId].fee, gas: 50000}("");
        require(sent, "send native failed");
        requests[_requestId].status = RequestStatus.Refunded;
        emit RequestRefunded(_requestId);
    }

    // --------------------- optimistic workflow functions ---------------------

    // Op functions
    function fulfillOpRequests(
        bytes32[] calldata _requestIds,
        bytes[] calldata _queryURLs,
        bytes[] calldata _sigs,
        address[] calldata _signers,
        uint256[] calldata _powers
    ) external {
        require(_requestIds.length > 0, "invalid requestIds");
        require(_requestIds.length == _queryURLs.length);

        bytes memory signBytes = abi.encodePacked(block.chainid);
        for (uint256 i = 0; i < _requestIds.length; i++) {
            signBytes = abi.encodePacked(signBytes, _requestIds[i]);
        }
        bytes32 domain = keccak256(abi.encodePacked(block.chainid, address(this), "FulfillRequests"));
        sigsVerifier.verifySigs(abi.encodePacked(domain, signBytes), _sigs, _signers, _powers);

        for (uint i = 0; i < _requestIds.length; i++) {
            brevisProof.submitOpResult(_requestIds[i]);
            requests[_requestIds[i]].status = RequestStatus.OpSubmitted;
            requestExts[_requestIds[i]].canChallengeBefore = block.timestamp + challengeWindow;
        }

        emit OpRequestsFulfilled(_requestIds, _queryURLs);
    }

    function askForQueryData(bytes32 _requestId) external payable {
        // TODO: msg.value should be larger than a configurable value

        require(requests[_requestId].status == RequestStatus.OpSubmitted, "not in a disputable status");

        requestExts[_requestId].askFor = AskForType.QueryData;
        requestExts[_requestId].shouldRespondBefore = block.timestamp + responseTimeout;
        requests[_requestId].status = RequestStatus.OpDisputing;

        emit AskFor(_requestId, AskForType.QueryData, msg.sender);
    }

    function postQueryData(bytes32 _requestId, bytes calldata _queryData) external {
        if (requests[_requestId].option == Option.OpMode) {
            bytes32 dataHash = keccak256(_queryData);
            keccakToMimc[dataHash] = _requestId;

            requests[_requestId].status = RequestStatus.OpQueryDataSubmitted;
            requestExts[_requestId].canChallengeBefore = block.timestamp + challengeWindow; // extend the window for proof challenge
            emit QueryDataPost(_requestId);
        } else {
            revert("not a valid op request");
        }
    }

    // after postQueryData with OpMode
    function challengeQueryData(bytes calldata _proof) external {
        (bytes32 myRequestId, bytes32 dataHash) = verifyQueryDataProofAndRetrieveKeys(_proof);
        bytes32 opRequestId = keccakToMimc[dataHash];
        require(opRequestId != bytes32(0), "query data not posted");

        if (myRequestId != opRequestId) {
            requests[opRequestId].status = RequestStatus.OpDisputed;
            // TODO slash flow
        }
    }

    function askForProof(bytes32 _requestId) external payable {
        // TODO: msg.value should be larger than a configurable value
        require(
            requests[_requestId].status == RequestStatus.OpSubmitted ||
                requests[_requestId].status == RequestStatus.OpQueryDataSubmitted,
            "not in a disputable status"
        );

        requestExts[_requestId].askFor = AskForType.Proof;
        requestExts[_requestId].shouldRespondBefore = block.timestamp + responseTimeout;
        requests[_requestId].status = RequestStatus.OpDisputing;

        emit AskFor(_requestId, AskForType.Proof, msg.sender);
    }

    function postProof(bytes32 _requestId, uint64 _chainId, bytes calldata _proof) external {
        bytes32 reqIdFromProof = IBrevisProof(brevisProof).submitProof(_chainId, _proof);
        require(_requestId == reqIdFromProof, "requestId and proof not match");

        requests[_requestId].status = RequestStatus.ZkAttested;

        emit ProofPost(_requestId);
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
        if (
            (requests[_requestId].status == RequestStatus.OpSubmitted ||
                requests[_requestId].status == RequestStatus.OpQueryDataSubmitted) &&
            requestExts[_requestId].canChallengeBefore <= block.timestamp
        ) {
            return RequestStatus.OpAttested;
        }

        if (
            requests[_requestId].status == RequestStatus.OpDisputing &&
            requestExts[_requestId].shouldRespondBefore <= block.timestamp
        ) {
            return RequestStatus.OpDisputed;
        }

        return requests[_requestId].status;
    }

    /*********************
     * Private Functions *
     *********************/

    function verifyQueryDataProofAndRetrieveKeys(
        bytes calldata _proof
    ) private returns (bytes32 _myRequestId, bytes32 _dataHash) {
        // TODO
    }
}
