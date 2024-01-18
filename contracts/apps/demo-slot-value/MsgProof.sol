// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../../verifiers/interfaces/ISlotValueVerifier.sol";

// source chain contract
contract MsgSender {
    uint64 public nonce; // slot 0
    mapping(uint64 => bytes32) public sent; // slot 1, nonce -> msgHash
    event MsgSent(uint64 nonce, address sender, bytes msg);

    function sendMsg(bytes calldata _msg) external {
        bytes32 msgHash = keccak256(abi.encodePacked(msg.sender, _msg));
        sent[nonce] = msgHash; // store in slot keccak256(abi.encode(_nonce, 1))
        emit MsgSent(nonce++, msg.sender, _msg);
    }
}

// destination chain contract
contract MsgReceiver {
    uint64 public senderChainId;
    bytes32 public senderContractHash;
    ISlotValueVerifier public slotValueVerifier;

    event MsgReceived(uint64 nonce, address sender, bytes msg);

    constructor(ISlotValueVerifier _verifier, uint64 _senderChainId, address _senderContract) {
        slotValueVerifier = _verifier;
        senderChainId = _senderChainId;
        senderContractHash = keccak256(abi.encodePacked(_senderContract));
    }

    function recvMsg(
        uint64 _nonce,
        address _sender,
        bytes calldata _msg,
        bytes calldata _proofData,
        bytes calldata _blkVerifyInfo
    ) external {
        // compute expected slot and msg hash, sender map slot is 1
        bytes32 slotKeyHash = keccak256(abi.encode(keccak256(abi.encode(_nonce, 1))));
        bytes32 msgHash = keccak256(abi.encodePacked(_sender, _msg));
        // retrieve zk verified slot info
        ISlotValueVerifier.SlotInfo memory slotInfo = slotValueVerifier.verifySlotValue(
            senderChainId,
            _proofData,
            _blkVerifyInfo
        );
        // compare expected and verified values
        require(slotInfo.slotKeyHash == slotKeyHash, "slot key not match");
        require(slotInfo.slotValue == msgHash, "slot value not match");
        require(slotInfo.addrHash == senderContractHash, "sender contract not match");
        emit MsgReceived(_nonce, _sender, _msg);
    }
}
