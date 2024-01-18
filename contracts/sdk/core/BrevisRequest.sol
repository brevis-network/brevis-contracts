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
        IBrevisApp callback;
        RequestStatus status;
    }
    mapping(bytes32 => Request) public requests;

    event RequestTimeoutUpdated(uint256 from, uint256 to);
    event RequestSent(bytes32 requestId, address sender, uint256 fee, IBrevisApp callback);
    event RequestFulfilled(bytes32 requestId);

    constructor(address _feeCollector, IBrevisProof _brevisProof) FeeVault(_feeCollector) {
        brevisProof = _brevisProof;
    }

    function sendRequest(bytes32 _requestId, address _refundee, IBrevisApp _callback) external payable {
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
            app.call(abi.encodeWithSelector(IBrevisApp.callback.selector, _requestId, _appCircuitOutput));
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
