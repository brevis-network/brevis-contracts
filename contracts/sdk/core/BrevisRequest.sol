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
    uint256 public responseTimeout; // BVN responses time window a challenge

    mapping(bytes32 => Request) public requests; // TODO: store data hash to save gas cost
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
        uint256 _nonce,
        address _refundee,
        address _callback,
        RequestOption _option
    ) external payable {
        bytes32 requestKey = keccak256(abi.encodePacked(_requestId, _nonce));
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
        requests[requestKey] = Request(block.timestamp, msg.value, _refundee, _callback, status);
        emit RequestSent(_requestId, _nonce, msg.sender, msg.value, _callback, _option);
    }

    function fulfillRequest(
        bytes32 _requestId,
        uint256 _nonce,
        uint64 _chainId,
        bytes calldata _proof,
        bytes calldata _appCircuitOutput
    ) external {
        bytes32 requestKey = keccak256(abi.encodePacked(_requestId, _nonce));
        Request storage request = requests[requestKey];
        require(
            request.status != RequestStatus.Null && request.status != RequestStatus.ZkAttested,
            "invalid request status"
        );

        bytes32 reqIdFromProof = IBrevisProof(brevisProof).submitProof(_chainId, _proof); // revert for invalid proof
        require(_requestId == reqIdFromProof, "requestId and proof not match");
        request.status = RequestStatus.ZkAttested;

        emit RequestFulfilled(_requestId, _nonce);

        if (request.callback != address(0)) {
            // The relayer should set correct gas limit. If the call failed due to insufficient gasleft(),
            // anyone can still call the app.brevisCallback directly to proceed
            (bool success, ) = request.callback.call(
                abi.encodeWithSelector(IBrevisApp.brevisCallback.selector, _requestId, _appCircuitOutput)
            );
            if (!success) {
                emit RequestCallbackFailed(_requestId, _nonce);
            }
        }
    }

    function fulfillRequests(
        bytes32[] calldata _requestIds,
        uint256[] calldata _nonces,
        uint64 _chainId,
        bytes calldata _proof,
        Brevis.ProofData[] calldata _proofDataArray,
        bytes[] calldata _appCircuitOutputs,
        address _callback
    ) external {
        IBrevisProof(brevisProof).mustSubmitAggProof(_chainId, _requestIds, _proof);

        for (uint8 i = 1; i < _requestIds.length; i++) {
            bytes32 requestKey = keccak256(abi.encodePacked(_requestIds[i], _nonces[i]));
            requests[requestKey].status = RequestStatus.ZkAttested;
        }

        emit RequestsFulfilled(_requestIds, _nonces);

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
                emit RequestsCallbackFailed(_requestIds, _nonces);
            }
        }
    }

    function refund(bytes32 _requestId, uint256 _nonce) external {
        // TODO: refund for op request
        bytes32 requestKey = keccak256(abi.encodePacked(_requestId, _nonce));
        Request storage request = requests[requestKey];
        require(
            request.status == RequestStatus.ZkPending || request.status == RequestStatus.OpPending,
            "invalid request status"
        );
        require(block.timestamp > request.timestamp + requestTimeout);
        require(!IBrevisProof(brevisProof).hasProof(_requestId), "proof already generated");
        (bool sent, ) = request.refundee.call{value: request.fee, gas: 50000}("");
        require(sent, "send native failed");
        request.status = RequestStatus.Refunded;
        emit RequestRefunded(_requestId, _nonce);
    }

    // --------------------- optimistic workflow functions ---------------------

    function fulfillOpRequests(
        bytes32[] calldata _requestIds,
        uint256[] calldata _nonces,
        bytes[] calldata _queryURLs,
        bytes[] calldata _sigs,
        address[] calldata _signers,
        uint256[] calldata _powers
    ) external {
        require(_requestIds.length > 0, "invalid requestIds");
        require(_requestIds.length == _queryURLs.length);

        bytes32 domain = keccak256(abi.encodePacked(block.chainid, address(this), "FulfillRequests"));
        sigsVerifier.verifySigs(abi.encodePacked(domain, _requestIds, _nonces), _sigs, _signers, _powers);

        for (uint i = 0; i < _requestIds.length; i++) {
            brevisProof.submitOpResult(_requestIds[i]);
            bytes32 requestKey = keccak256(abi.encodePacked(_requestIds[i], _nonces[i]));
            Request storage request = requests[requestKey];
            require(request.status == RequestStatus.OpPending, "invalid request status");
            request.status = RequestStatus.OpSubmitted;
        }

        emit OpRequestsFulfilled(_requestIds, _nonces, _queryURLs);
    }

    function askForQueryData(bytes32 _requestId, uint256 _nonce) external payable {
        // TODO: msg.value should be larger than a configurable value
        bytes32 requestKey = keccak256(abi.encodePacked(_requestId, _nonce));
        Request storage request = requests[requestKey];
        Dispute storage dispute = disputes[requestKey];
        require(request.status == RequestStatus.OpSubmitted, "not in a disputable status");
        require(request.timestamp + challengeWindow > block.timestamp, "pass challenge window");

        request.status = RequestStatus.OpDisputing;
        dispute.status = DisputeStatus.WaitingForQueryData;
        dispute.responseDeadline = block.timestamp + responseTimeout;

        emit AskFor(_requestId, _nonce, DisputeStatus.WaitingForQueryData, msg.sender);
    }

    function postQueryData(bytes32 _requestId, uint256 _nonce, bytes calldata _queryData) external {
        bytes32 requestKey = keccak256(abi.encodePacked(_requestId, _nonce));
        Request storage request = requests[requestKey];
        Dispute storage dispute = disputes[requestKey];
        require(request.status == RequestStatus.OpDisputing, "invalid request status");
        require(dispute.status == DisputeStatus.WaitingForQueryData, "invalid dispute status");

        disputes[requestKey].queryDataHash = keccak256(_queryData); // TODO: use claimed mimc?
        disputes[requestKey].status = DisputeStatus.QueryDataPosted;
        emit QueryDataPosted(_requestId, _nonce);
    }

    function askForQueryDataProof(bytes32 _requestId, uint256 _nonce) external payable {
        // TODO: msg.value should be larger than a configurable value
        bytes32 requestKey = keccak256(abi.encodePacked(_requestId, _nonce));
        Request storage request = requests[requestKey];
        Dispute storage dispute = disputes[requestKey];
        require(
            request.status == RequestStatus.OpDisputing && dispute.status == DisputeStatus.QueryDataPosted,
            "invalid states"
        );
        require(request.timestamp + challengeWindow > block.timestamp, "pass challenge window");

        request.status = RequestStatus.OpDisputing;
        dispute.status = DisputeStatus.WaitingForQueryDataProof;
        dispute.responseDeadline = block.timestamp + responseTimeout;

        emit AskFor(_requestId, _nonce, DisputeStatus.WaitingForQueryDataProof, msg.sender);
    }

    function postQueryDataProof(bytes32 _requestId, uint256 _nonce, bytes calldata _proof) external {
        bytes32 requestKey = keccak256(abi.encodePacked(_requestId, _nonce));
        Request storage request = requests[requestKey];
        Dispute storage dispute = disputes[requestKey];
        require(
            request.status == RequestStatus.OpDisputing && dispute.status == DisputeStatus.WaitingForQueryDataProof,
            "invalid states"
        );
        disputes[requestKey].status = DisputeStatus.QueryDataProofPosted;
        // TODO: check proof

        emit QueryDataProofPosted(_requestId, _nonce);
    }

    function askForValidityProof(bytes32 _requestId, uint256 _nonce) external payable {
        // TODO: msg.value should be larger than a configurable value
        bytes32 requestKey = keccak256(abi.encodePacked(_requestId, _nonce));
        Request storage request = requests[requestKey];
        Dispute storage dispute = disputes[requestKey];
        require(
            request.status == RequestStatus.OpSubmitted ||
                (request.status == RequestStatus.OpDisputing &&
                    dispute.status != DisputeStatus.WaitingForValidityProof),
            "invalid states"
        );
        require(request.timestamp + challengeWindow > block.timestamp, "pass challenge window");

        request.status = RequestStatus.OpDisputing;
        dispute.status = DisputeStatus.WaitingForValidityProof;
        dispute.responseDeadline = block.timestamp + responseTimeout;

        emit AskFor(_requestId, _nonce, DisputeStatus.WaitingForValidityProof, msg.sender);
    }

    function postValidityProof(bytes32 _requestId, uint256 _nonce, uint64 _chainId, bytes calldata _proof) external {
        bytes32 reqIdFromProof = IBrevisProof(brevisProof).submitProof(_chainId, _proof);
        require(_requestId == reqIdFromProof, "requestId and proof not match");

        bytes32 requestKey = keccak256(abi.encodePacked(_requestId, _nonce));
        requests[requestKey].status = RequestStatus.ZkAttested;
        disputes[requestKey].status = DisputeStatus.ValidityProofPosted;

        emit ValidityProofProofPosted(_requestId, _nonce);
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

    function queryRequestStatus(bytes32 _requestId, uint256 _nonce) external view returns (RequestStatus) {
        bytes32 requestKey = keccak256(abi.encodePacked(_requestId, _nonce));
        Request storage request = requests[requestKey];
        if (request.status == RequestStatus.OpSubmitted) {
            if (request.timestamp + challengeWindow < block.timestamp) {
                return RequestStatus.OpAttested;
            }
        } else if (request.status == RequestStatus.OpDisputing) {
            Dispute storage dispute = disputes[requestKey];
            if (dispute.status == DisputeStatus.QueryDataPosted) {
                if (request.timestamp + challengeWindow < block.timestamp) {
                    return RequestStatus.OpAttested;
                }
            } else if (dispute.responseDeadline < block.timestamp) {
                // WaitingForQueryData || WaitingForZkProof
                return RequestStatus.OpDisputed;
            }
        }
        return requests[requestKey].status;
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
