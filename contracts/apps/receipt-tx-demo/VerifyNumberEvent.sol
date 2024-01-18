// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/access/Ownable.sol";
import "../../verifiers/interfaces/IReceiptVerifier.sol";
import {BytesLib} from "solidity-bytes-utils/contracts/BytesLib.sol";

contract VerifyNumberEvent is Ownable {
    IReceiptVerifier public receiptVerifier;
    mapping(uint64 => address) public srcContract;
    bytes32 constant eventTopic = keccak256(bytes("SendNumber(address,uint256)"));

    event VerifiedNumber(uint64 chainId, uint64 blknum, address from, uint256 number);

    constructor(IReceiptVerifier _receiptVerifier) {
        receiptVerifier = _receiptVerifier;
    }

    function submitNumberReceiptProof(
        bytes calldata _receipt,
        bytes calldata _proof,
        bytes calldata _auxiBlkVerifyInfo
    ) external {
        // retrieve verified event
        IReceiptVerifier.ReceiptInfo memory receiptInfo = receiptVerifier.verifyReceiptAndLog(
            _receipt,
            _proof,
            _auxiBlkVerifyInfo
        );
        IReceiptVerifier.LogInfo memory log = receiptInfo.logs[0];

        // compare expected and verified values
        require(receiptInfo.success, "tx failed");
        require(log.addr == srcContract[receiptInfo.chainId], "invalid sender contract");
        require(log.topics[0] == eventTopic, "invalid event");

        // decode event data
        address from = address(bytes20(BytesLib.slice(log.data, 12, 20)));
        uint256 number = uint256(bytes32(BytesLib.slice(log.data, 32, 32)));
        emit VerifiedNumber(receiptInfo.chainId, receiptInfo.blkNum, from, number);
    }

    function setReceiptVerifier(IReceiptVerifier _receiptVerifier) external onlyOwner {
        receiptVerifier = _receiptVerifier;
    }

    function setSrcContract(uint64 _chainId, address _contract) external onlyOwner {
        srcContract[_chainId] = _contract;
    }
}
