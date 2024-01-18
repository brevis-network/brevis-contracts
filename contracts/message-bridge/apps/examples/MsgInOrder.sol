// SPDX-License-Identifier: MIT
pragma solidity >=0.8.9;

import "../../framework/MessageApp.sol";

// a simple example to enforce in-order message delivery
contract MsgInOrder is MessageApp {
    event MessageReceived(uint64 srcChainId, address srcContract, address sender, uint64 seq, bytes message);

    // map at source chain. (dstChainId, dstContract) -> seq
    mapping(uint64 => mapping(address => uint64)) public sendSeq;

    // map at destination chain (srcChainId, srcContract) -> seq
    mapping(uint64 => mapping(address => uint64)) public recvSeq;

    constructor(IMessageBridge _messageBridge) MessageApp(_messageBridge) {}

    // called by user on source chain to send cross-chain message
    function sendMessage(uint64 _dstChainId, address _dstContract, bytes calldata _message) external payable {
        uint64 seq = sendSeq[_dstChainId][_dstContract];
        bytes memory message = abi.encode(msg.sender, seq, _message);
        _sendMessage(_dstChainId, _dstContract, message);
        sendSeq[_dstChainId][_dstContract] += 1;
    }

    // called by MessageBridge on destination chain to receive message
    function _handleMessage(
        uint64 _srcChainId,
        address _srcContract,
        bytes calldata _message,
        address // execution
    ) internal override {
        (address sender, uint64 seq, bytes memory message) = abi.decode((_message), (address, uint64, bytes));
        uint64 expectedSeq = recvSeq[_srcChainId][_srcContract];
        require(seq == expectedSeq, _abortReason("sequence number not expected")); // let execution retry later.
        emit MessageReceived(_srcChainId, _srcContract, sender, seq, message);
        recvSeq[_srcChainId][_srcContract] += 1;
    }
}
