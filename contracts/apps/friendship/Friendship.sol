// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/access/Ownable.sol";
import "../../verifiers/interfaces/ITxVerifier.sol";

contract Friendship is Ownable {
    ITxVerifier public txVerifier;

    // (chainId, sender, receiver) -> timestamp of latest tx
    mapping(uint64 => mapping(address => mapping(address => uint64))) public lastestTxTimestamps;

    event VerifiedFriendship(uint64 chainId, address from, address to, uint64 timestamp);

    constructor(ITxVerifier _txVerifier) {
        txVerifier = _txVerifier;
    }

    function submitFriendshipProof(
        bytes calldata _tx,
        bytes calldata _proof,
        bytes calldata _auxiBlkVerifyInfo
    ) external {
        ITxVerifier.TxInfo memory txInfo = txVerifier.verifyTxAndLog(_tx, _proof, _auxiBlkVerifyInfo);
        require(txInfo.blkTime > lastestTxTimestamps[txInfo.chainId][txInfo.from][txInfo.to], "not latest tx");
        lastestTxTimestamps[txInfo.chainId][txInfo.from][txInfo.to] = txInfo.blkTime;
        emit VerifiedFriendship(txInfo.chainId, txInfo.from, txInfo.to, txInfo.blkTime);
    }

    function setTxVerifier(ITxVerifier _txVerifier) external onlyOwner {
        txVerifier = _txVerifier;
    }
}
