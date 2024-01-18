// SPDX-License-Identifier: MIT

pragma solidity ^0.8.18;

import "../../interfaces/IEthereumLightClient.sol";

interface IMessageBridge {
    enum MessageStatus {
        Null,
        Success,
        Fail
    }

    event MessageSent(
        bytes32 indexed messageId,
        uint256 indexed nonce,
        uint64 dstChainId,
        address sender,
        address receiver,
        bytes message
    );
    event MessageExecuted(
        bytes32 indexed messageId,
        uint256 indexed nonce,
        uint64 srcChainId,
        address sender,
        address receiver,
        bytes message,
        bool success
    );
    event MessageCallReverted(bytes32 messageId, string reason); // help debug

    function lightClients(uint256 chainId) external view returns (IEthereumLightClient);

    function sendMessage(uint64 dstChainId, address receiver, bytes calldata message) external returns (bytes32);

    function executeMessage(
        uint64 srcChainId,
        uint64 nonce,
        address sender,
        address receiver,
        bytes calldata message,
        bytes[] calldata accountProof,
        bytes[] calldata storageProof
    ) external returns (bool);

    function getExecutionStateRootAndSlot(uint64 chainId) external view returns (bytes32 root, uint64 slot);
}
