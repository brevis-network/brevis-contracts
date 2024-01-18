// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../interfaces/IMessageReceiverApp.sol";
import "../libraries/MsgLib.sol";
import "./MessageBridgeAddress.sol";

abstract contract MessageReceiverApp is IMessageReceiverApp, MessageBridgeAddress {
    modifier onlyMessageBridge() {
        require(msg.sender == address(messageBridge), "caller is not message bridge");
        _;
    }

    /**
     * @notice Called by MessageBridge to execute a message
     * @param srcChainId The source chain ID where the message is originated from
     * @param sender The address of the source app contract
     * @param message Arbitrary message bytes originated from and encoded by the source app contract
     * @param executor Address who called the MessageBridge execution function
     * @return true Always return true if _handleMessage is not reverted
     */
    function executeMessage(
        uint64 srcChainId,
        address sender,
        bytes calldata message,
        address executor
    ) external onlyMessageBridge returns (bool) {
        _handleMessage(srcChainId, sender, message, executor);
        return true;
    }

    /**
     * @notice Internally called by executeMessage function to execute a message
     * @param srcChainId The source chain ID where the message is originated from
     * @param sender The address of the source app contract
     * @param message Arbitrary message bytes originated from and encoded by the source app contract
     * @param executor Address who called the MessageBridge execution function
     */
    function _handleMessage(
        uint64 srcChainId,
        address sender,
        bytes calldata message,
        address executor
    ) internal virtual;

    // Add abort prefix in the reason string for require or revert.
    // This will abort (revert) the message execution without markig it as failed state,
    // making it possible to retry later.
    function _abortReason(string memory reason) internal pure returns (string memory) {
        return string.concat(MsgLib.ABORT_PREFIX, reason);
    }
}
