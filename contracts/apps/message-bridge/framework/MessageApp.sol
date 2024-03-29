// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

import "./MessageSenderApp.sol";
import "./MessageReceiverApp.sol";

abstract contract MessageApp is MessageSenderApp, MessageReceiverApp {
    constructor(IMessageBridge _messageBridge) {
        messageBridge = _messageBridge;
    }
}
