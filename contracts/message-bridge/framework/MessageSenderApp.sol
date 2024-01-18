// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../interfaces/IMessageBridge.sol";
import "./MessageBridgeAddress.sol";

abstract contract MessageSenderApp is MessageBridgeAddress {
    /**
     * @notice Send a message to a contract on another chain.
     * @param dstChainId The destination chain ID.
     * @param receiver The address of the destination app contract.
     * @param message Arbitrary message bytes to be decoded by the destination app contract.
     * @return messageId Message Id computed by MessageBridge
     */
    function _sendMessage(
        uint64 dstChainId,
        address receiver,
        bytes memory message
    ) internal returns (bytes32 messageId) {
        return messageBridge.sendMessage(dstChainId, receiver, message);
    }
}
