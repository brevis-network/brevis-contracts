// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./interfaces/IMessageBridge.sol";
import "./interfaces/IMessageReceiverApp.sol";
import "./libraries/RLPReader.sol";
import "./libraries/MerkleProofTree.sol";
import "./libraries/MsgLib.sol";
import "../interfaces/IEthereumLightClient.sol";
import "../verifiers/interfaces/ISlotValueVerifier.sol";

contract MessageBridge is IMessageBridge, ReentrancyGuard, Ownable {
    using RLPReader for RLPReader.RLPItem;
    using RLPReader for bytes;

    /* Sender side (source chain) storage */
    mapping(uint64 => bytes32) public sentMessages; // nonce -> messageId
    uint256 constant SENT_MESSAGES_STORAGE_SLOT = 2;
    uint64 public nonce;

    /* Receiver side (dest chain) storage */
    mapping(bytes32 => MessageStatus) public receivedMessages; // messageId -> status
    mapping(uint256 => IEthereumLightClient) public lightClients; // chainId -> light client
    mapping(uint256 => address) public remoteMessageBridges; // chainId -> source chain bridge
    mapping(uint256 => bytes32) public remoteMessageBridgeHashes;
    ISlotValueVerifier public slotValueVerifier;
    // minimum amount of gas needed by this contract before it tries to deliver a message to the target.
    uint256 public preExecuteMessageGasUsage;

    /****************************************
     * Sender side (source chain) functions *
     ****************************************/

    function sendMessage(uint64 _dstChainId, address _receiver, bytes calldata _message) external returns (bytes32) {
        bytes32 messageId = MsgLib.computeMessageId(
            nonce,
            msg.sender,
            _receiver,
            uint64(block.chainid),
            _dstChainId,
            _message
        );
        sentMessages[nonce] = messageId;
        emit MessageSent(messageId, nonce++, _dstChainId, msg.sender, _receiver, _message);
        return messageId;
    }

    /****************************************
     * Receiver side (dest chain) functions *
     ****************************************/

    function executeMessage(
        uint64 _srcChainId,
        uint64 _nonce,
        address _sender,
        address _receiver,
        bytes calldata _message,
        bytes[] calldata _accountProof,
        bytes[] calldata _storageProof
    ) external nonReentrant returns (bool success) {
        (bytes32 messageId, bytes32 slotKeyHash) = _getSlotAndMessageId(
            _srcChainId,
            _nonce,
            _sender,
            _receiver,
            _message
        );
        _verifyAccountAndStorageProof(_srcChainId, messageId, slotKeyHash, _accountProof, _storageProof);
        return _executeMessage(messageId, _srcChainId, _nonce, _sender, _receiver, _message);
    }

    function executeMessageWithZkProof(
        uint64 _srcChainId,
        uint64 _nonce,
        address _sender,
        address _receiver,
        bytes calldata _message,
        bytes calldata _zkProofData,
        bytes calldata _blkVerifyInfo
    ) external nonReentrant returns (bool success) {
        (bytes32 messageId, bytes32 slotKeyHash) = _getSlotAndMessageId(
            _srcChainId,
            _nonce,
            _sender,
            _receiver,
            _message
        );
        _verifyZkSlotValueProof(_srcChainId, messageId, slotKeyHash, _zkProofData, _blkVerifyInfo);
        return _executeMessage(messageId, _srcChainId, _nonce, _sender, _receiver, _message);
    }

    function setLightClient(uint64 _chainId, address _lightClient) external onlyOwner {
        lightClients[_chainId] = IEthereumLightClient(_lightClient);
    }

    function setSlotValueVerifier(address _slotValueVerifier) external onlyOwner {
        slotValueVerifier = ISlotValueVerifier(_slotValueVerifier);
    }

    function setRemoteMessageBridge(uint64 _chainId, address _remoteMessageBridge) external onlyOwner {
        remoteMessageBridges[_chainId] = _remoteMessageBridge;
        remoteMessageBridgeHashes[_chainId] = keccak256(abi.encodePacked(_remoteMessageBridge));
    }

    function setPreExecuteMessageGasUsage(uint256 _usage) public onlyOwner {
        preExecuteMessageGasUsage = _usage;
    }

    function getExecutionStateRootAndSlot(uint64 _chainId) public view returns (bytes32 root, uint64 slot) {
        return lightClients[_chainId].optimisticExecutionStateRootAndSlot();
    }

    function _getSlotAndMessageId(
        uint64 _srcChainId,
        uint64 _nonce,
        address _sender,
        address _receiver,
        bytes calldata _message
    ) private view returns (bytes32 messageId, bytes32 slotKeyHash) {
        messageId = MsgLib.computeMessageId(_nonce, _sender, _receiver, _srcChainId, uint64(block.chainid), _message);
        require(receivedMessages[messageId] == MessageStatus.Null, "MessageBridge: message already executed");
        slotKeyHash = keccak256(abi.encode(keccak256(abi.encode(_nonce, SENT_MESSAGES_STORAGE_SLOT))));
    }

    function _verifyAccountAndStorageProof(
        uint64 _srcChainId,
        bytes32 _messageId,
        bytes32 _slotKeyHash,
        bytes[] calldata _accountProof,
        bytes[] calldata _storageProof
    ) private view {
        require(
            _retrieveStorageRoot(_srcChainId, _accountProof) == keccak256(_storageProof[0]),
            "MessageBridge: invalid storage root"
        );
        bytes memory proof = MerkleProofTree.read(_slotKeyHash, _storageProof);
        require(bytes32(proof.toRlpItem().toUint()) == _messageId, "MessageBridge: invalid message hash");
    }

    function _retrieveStorageRoot(uint64 _srcChainId, bytes[] calldata _accountProof) private view returns (bytes32) {
        // verify accountProof and get storageRoot
        (bytes32 executionStateRoot, ) = getExecutionStateRootAndSlot(_srcChainId);
        require(executionStateRoot != bytes32(0), "MessageBridge: execution state root not found");
        require(executionStateRoot == keccak256(_accountProof[0]), "MessageBridge: invalid account proof root");

        // get storageRoot
        bytes memory accountInfo = MerkleProofTree.read(remoteMessageBridgeHashes[_srcChainId], _accountProof);
        RLPReader.RLPItem[] memory items = accountInfo.toRlpItem().toList();
        require(items.length == 4, "MessageBridge: invalid account decoded from RLP");
        return bytes32(items[2].toUint());
    }

    function _verifyZkSlotValueProof(
        uint64 _srcChainId,
        bytes32 _messageId,
        bytes32 _slotKeyHash,
        bytes calldata _zkProofData,
        bytes calldata _blkVerifyInfo
    ) private view {
        ISlotValueVerifier.SlotInfo memory slotInfo = slotValueVerifier.verifySlotValue(
            _srcChainId,
            _zkProofData,
            _blkVerifyInfo
        );
        require(slotInfo.slotKeyHash == _slotKeyHash, "MessageBridge: slot key not match");
        require(slotInfo.slotValue == _messageId, "MessageBridge: slot value not match");
        require(slotInfo.addrHash == remoteMessageBridgeHashes[_srcChainId], "MessageBridge: src contract not match");
    }

    function _executeMessage(
        bytes32 _messageId,
        uint64 _srcChainId,
        uint64 _nonce,
        address _sender,
        address _receiver,
        bytes calldata _message
    ) private returns (bool success) {
        // execute message
        bytes memory recieveCall = abi.encodeWithSelector(
            IMessageReceiverApp.executeMessage.selector,
            _srcChainId,
            _sender,
            _message,
            msg.sender
        );
        uint256 gasLeftBeforeExecution = gasleft();
        (bool ok, bytes memory res) = _receiver.call(recieveCall);
        if (ok) {
            success = abi.decode((res), (bool));
        } else {
            _handleExecutionRevert(_messageId, gasLeftBeforeExecution, res);
        }
        receivedMessages[_messageId] = success ? MessageStatus.Success : MessageStatus.Fail;
        emit MessageExecuted(_messageId, _nonce, _srcChainId, _sender, _receiver, _message, success);
        return success;
    }

    function _handleExecutionRevert(
        bytes32 messageId,
        uint256 _gasLeftBeforeExecution,
        bytes memory _returnData
    ) private {
        uint256 gasLeftAfterExecution = gasleft();
        uint256 maxTargetGasLimit = block.gaslimit - preExecuteMessageGasUsage;
        if (_gasLeftBeforeExecution < maxTargetGasLimit && gasLeftAfterExecution <= _gasLeftBeforeExecution / 64) {
            // if this happens, the execution must have not provided sufficient gas limit,
            // then the tx should revert instead of recording a non-retryable failure status
            // https://github.com/wolflo/evm-opcodes/blob/main/gas.md#aa-f-gas-to-send-with-call-operations
            assembly {
                invalid()
            }
        }
        string memory revertMsg = MsgLib.checkRevertMsg(_returnData);
        // otherwiase, emit revert message, return and mark the execution as failed (non-retryable)
        emit MessageCallReverted(messageId, revertMsg);
    }
}
