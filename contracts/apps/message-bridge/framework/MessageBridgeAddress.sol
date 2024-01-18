// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../interfaces/IMessageBridge.sol";

abstract contract MessageBridgeAddress {
    IMessageBridge public messageBridge;
}
