// SPDX-License-Identifier: MIT
pragma solidity >=0.8.9;

import "../../framework/MessageApp.sol";

// A HelloWorld test example for basic cross-chain message passing
contract MsgTest is MessageApp {
    event MessageReceived(uint64 srcChainId, address srcContract, address sender, uint64 number);

    constructor(IMessageBridge _messageBridge) MessageApp(_messageBridge) {}

    // called by user on source chain to send cross-chain messages
    function sendMessage(uint64 _dstChainId, address _dstContract, uint64 _number) external {
        bytes memory message = abi.encode(msg.sender, _number);
        _sendMessage(_dstChainId, _dstContract, message);
    }

    // called by MessageBridge on destination chain to receive cross-chain messages
    function _handleMessage(
        uint64 _srcChainId,
        address _srcContract,
        bytes calldata _message,
        address // execution
    ) internal override {
        (address sender, uint64 number) = abi.decode((_message), (address, uint64));
        require(number != 1000, _abortReason("test abort"));
        require(number != 1001, "test revert");
        emit MessageReceived(_srcChainId, _srcContract, sender, number);
    }
}
