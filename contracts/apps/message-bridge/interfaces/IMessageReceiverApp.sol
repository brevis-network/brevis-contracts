// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

interface IMessageReceiverApp {
    /**
     * @notice Called by MessageBridge to execute a message
     * @param _srcChainId The source chain ID where the message is originated from
     * @param _sender The address of the source app contract
     * @param _message Arbitrary message bytes originated from and encoded by the source app contract
     * @param _executor Address who called the MessageBridge execution function
     */
    function executeMessage(
        uint64 _srcChainId,
        address _sender,
        bytes calldata _message,
        address _executor
    ) external returns (bool);
}
