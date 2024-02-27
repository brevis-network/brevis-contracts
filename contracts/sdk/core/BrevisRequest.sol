// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "./FeeVault.sol";
import "../interface/IBrevisProof.sol";
import "../interface/IBrevisApp.sol";
import "../../interfaces/ISigsVerifier.sol";

contract BrevisRequest is FeeVault {
    uint256 public requestTimeout;
    IBrevisProof public brevisProof;
    ISigsVerifier public immutable sigsVerifier;

    enum RequestStatus {
        Pending,
        ZkAttested,
        OpSubmitted,
        OpQueryDataSubmitted,
        OpDisputing,
        OpDisputed,
        OpAttested,
        Refunded
    }

    enum Option {
        ZkMode,
        OpMode_MIMC,
        OpMode_KECCAK
    }

    struct Request {
        uint256 deadline;
        uint256 fee;
        address refundee;
        IBrevisApp callback;
        RequestStatus status;
        Option option;
    }
    mapping(bytes32 => Request) public requests; // TODO: store hash of request data to save gas cost

    event RequestTimeoutUpdated(uint256 from, uint256 to);
    event RequestSent(bytes32 requestId, address sender, uint256 fee, IBrevisApp callback, Option option);
    event RequestFulfilled(bytes32 requestId);
    event RequestRefunded(bytes32 requestId);

    constructor(address _feeCollector, IBrevisProof _brevisProof, ISigsVerifier _sigsVerifier) FeeVault(_feeCollector) {
        brevisProof = _brevisProof;
        sigsVerifier = _sigsVerifier;
    }

    function sendRequest(bytes32 _requestId, address _refundee, IBrevisApp _callback, Option _option) external payable {
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
        bool _withAppProof,
        bytes calldata _appCircuitOutput
    ) external {
        bytes32 reqIdFromProof = IBrevisProof(brevisProof).submitProof(_chainId, _proof, _withAppProof); // will be reverted when proof is not valid
        require(_requestId == reqIdFromProof, "requestId and proof not match");
        chargeFee(_requestId); // will be reverted when failed to charge fee
        requests[_requestId].status = RequestStatus.ZkAttested;

        emit RequestFulfilled(_requestId);

        address app = address(requests[_requestId].callback);
        if (app != address(0)) {
            // No matter if the call is success or not. The relayer should set correct gas limit.
            // If the call exceeds the gasleft(), as the proof data is saved ahead,
            // anyone can still call the app.callback directly to proceed
            app.call(abi.encodeWithSelector(IBrevisApp.brevisCallback.selector, _requestId, _appCircuitOutput));
        }
    }

    function chargeFee(bytes32 _requestId) public {
        require(requests[_requestId].deadline != 0, "request not in queue");
        require(IBrevisProof(brevisProof).hasProof(_requestId), "proof not generated");
        requests[_requestId].deadline = 0; //simply set deadline to 0, then fee is not able be refunded
    }

    function refund(bytes32 _requestId) public {
        require(requests[_requestId].deadline != 0, "request not in queue");
        require(block.timestamp > requests[_requestId].deadline);
        requests[_requestId].deadline = 0;
        (bool sent, ) = requests[_requestId].refundee.call{value: requests[_requestId].fee, gas: 50000}("");
        require(sent, "send native failed");
        requests[_requestId].status = RequestStatus.Refunded;
        emit RequestRefunded(_requestId);
    }

    function setRequestTimeout(uint256 _timeout) external onlyOwner {
        uint256 oldTimeout = requestTimeout;
        requestTimeout = _timeout;
        emit RequestTimeoutUpdated(oldTimeout, _timeout);
    }

    enum AskForType {
        NULL,
        QueryData,
        Proof
    }

    struct RequestExt {
        uint256 canChallengeBefore;
        AskForType askFor;
        uint256 shouldRespondBefore;
    }
    mapping(bytes32 => RequestExt) public requestExts;
    mapping(bytes32 => bytes32) public keccakToMimc;

    uint256 public challengeWindow; // in seconds
    // BVN responses an ask-for-data request
    uint256 public responseTimeout;
    event ChallengeWindowUpdated(uint256 from, uint256 to);
    event ResponseTimeoutUpdated(uint256 from, uint256 to);

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

    event OpRequestsFulfilled(bytes32[] requestIds, bytes[] queryURLs);

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
            if (i != _requestIds.length - 1) {
                signBytes = abi.encodePacked(signBytes, _requestIds[i], ",");
            } else {
                signBytes = abi.encodePacked(signBytes, _requestIds[i]);
            }
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

    event AskFor(bytes32 indexed requestId, AskForType askFor, address from);
    event QueryDataPost(bytes32 indexed requestId);
    event ProofPost(bytes32 indexed requestId);

    function askForQueryData(bytes32 _requestId) external payable {
        // TODO: msg.value should be larger than a configurable value

        require(requests[_requestId].status == RequestStatus.OpSubmitted, "not in a disputable status");

        requestExts[_requestId].askFor = AskForType.QueryData;
        requestExts[_requestId].shouldRespondBefore = block.timestamp + responseTimeout;
        requests[_requestId].status = RequestStatus.OpDisputing;

        emit AskFor(_requestId, AskForType.QueryData, msg.sender);
    }

    function postQueryData(bytes32 _requestId, bytes calldata _queryData) external {
        if (requests[_requestId].option == Option.OpMode_KECCAK) {
            require(keccak256(_queryData) == _requestId, "not valid queryData");

            requests[_requestId].status = RequestStatus.OpQueryDataSubmitted;
            requestExts[_requestId].canChallengeBefore = block.timestamp + challengeWindow; // extend the window for proof challenge
            emit QueryDataPost(_requestId);
        } else if (requests[_requestId].option == Option.OpMode_MIMC) {
            bytes32 dataHash = keccak256(_queryData);
            keccakToMimc[dataHash] = _requestId;

            requests[_requestId].status = RequestStatus.OpQueryDataSubmitted;
            requestExts[_requestId].canChallengeBefore = block.timestamp + challengeWindow; // extend the window for proof challenge
            emit QueryDataPost(_requestId);
        } else {
            revert("not a valid op request");
        }
    }

    // after postQueryData with OpMode_MIMC
    function challengeQueryData(bytes calldata _proof) external {
        (bytes32 myRequestId, bytes32 dataHash) = verifyQueryDataProofAndRetrieveKeys(_proof);
        bytes32 opRequestId = keccakToMimc[dataHash];
        require(opRequestId != bytes32(0), "query data not posted");

        if (myRequestId != opRequestId) {
            requests[opRequestId].status = RequestStatus.OpDisputed;
            // TODO slash flow
        }
    }

    function verifyQueryDataProofAndRetrieveKeys(
        bytes calldata _proof
    ) private returns (bytes32 _myRequestId, bytes32 _dataHash) {
        // TODO
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
        bytes32 reqIdFromProof = IBrevisProof(brevisProof).submitProof(_chainId, _proof, true);
        require(_requestId == reqIdFromProof, "requestId and proof not match");

        requests[_requestId].status = RequestStatus.ZkAttested;

        emit ProofPost(_requestId);
    }
}
