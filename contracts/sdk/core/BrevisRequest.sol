// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "./FeeVault.sol";
import "../interface/IBrevisProof.sol";
import "../interface/IBrevisApp.sol";

contract BrevisRequest is FeeVault {
    uint256 public requestTimeout;
    IBrevisProof public brevisProof;

    enum RequestStatus {
        Pending,
        ZkAttested,
        Refunded
    }

    struct Request {
        uint256 deadline;
        uint256 fee;
        address refundee;
        address callback;
        RequestStatus status;
    }
    mapping(bytes32 => Request) public requests; // TODO: store hash of request data to save gas cost

    event RequestTimeoutUpdated(uint256 from, uint256 to);
    event RequestSent(bytes32 requestId, address sender, uint256 fee, address callback);
    event RequestFulfilled(bytes32 requestId);

    constructor(address _feeCollector, IBrevisProof _brevisProof) FeeVault(_feeCollector) {
        brevisProof = _brevisProof;
    }

    function sendRequest(bytes32 _requestId, address _refundee, address _callback) external payable {
        require(requests[_requestId].deadline == 0, "request already in queue");
        require(_refundee != address(0), "refundee not provided");
        requests[_requestId] = Request(
            block.timestamp + requestTimeout,
            msg.value,
            _refundee,
            _callback,
            RequestStatus.Pending
        );
        emit RequestSent(_requestId, msg.sender, msg.value, _callback);
    }

    function fulfillRequest(
        bytes32 _requestId,
        uint64 _chainId,
        bytes calldata _proof,
        bool _withAppProof,
        bytes calldata _appCircuitOutput
    ) external {
        require(!IBrevisProof(brevisProof).hasProof(_requestId), "proof already generated");

        bytes32 reqIdFromProof = IBrevisProof(brevisProof).submitProof(_chainId, _proof, _withAppProof); // will revert if proof is not valid
        require(_requestId == reqIdFromProof, "requestId and proof not match");
        requests[_requestId].status = RequestStatus.ZkAttested;

        emit RequestFulfilled(_requestId);

        address app = requests[_requestId].callback;
        if (app != address(0)) {
            // No matter if the call is success or not. The relayer should set correct gas limit.
            // If the call exceeds the gasleft(), as the proof data is saved ahead,
            // anyone can still call the app.callback directly to proceed
            app.call(abi.encodeWithSelector(IBrevisApp.brevisCallback.selector, _requestId, _appCircuitOutput));
        }
    }

    function refund(bytes32 _requestId) public {
        require(block.timestamp > requests[_requestId].deadline);
        require(!IBrevisProof(brevisProof).hasProof(_requestId), "proof already generated");
        requests[_requestId].deadline = 0; //reset deadline, then user is able to sent request again
        (bool sent, ) = requests[_requestId].refundee.call{value: requests[_requestId].fee, gas: 50000}("");
        require(sent, "send native failed");
        requests[_requestId].status = RequestStatus.Refunded;
    }

    function setRequestTimeout(uint256 _timeout) external onlyOwner {
        uint256 oldTimeout = requestTimeout;
        requestTimeout = _timeout;
        emit RequestTimeoutUpdated(oldTimeout, _timeout);
    }

    function queryRequestStatus(bytes32 _requestId) external view returns (RequestStatus) {
        return requests[_requestId].status;
    }
}
