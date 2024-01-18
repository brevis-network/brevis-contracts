// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/access/Ownable.sol";
import "../../verifiers/interfaces/ITxVerifier.sol";
import {BytesLib} from "solidity-bytes-utils/contracts/BytesLib.sol";

contract VerifyNumberTx is Ownable {
    ITxVerifier public txVerifier;
    mapping(uint64 => address) public srcContract;
    bytes4 constant funcSelector = bytes4(keccak256(bytes("sendNumber(uint256)")));

    event VerifiedNumber(uint64 chainId, uint64 blknum, address from, uint256 number);

    constructor(ITxVerifier _txVerifier) {
        txVerifier = _txVerifier;
    }

    function submitNumberTxProof(
        bytes calldata _tx,
        bytes calldata _proof,
        bytes calldata _auxiBlkVerifyInfo
    ) external {
        ITxVerifier.TxInfo memory txInfo = txVerifier.verifyTxAndLog(_tx, _proof, _auxiBlkVerifyInfo);
        require(txInfo.to == srcContract[txInfo.chainId], "invalid sender contract");
        uint256 number = decodeCalldata(txInfo.data);
        emit VerifiedNumber(txInfo.chainId, txInfo.blkNum, txInfo.from, number);
    }

    function decodeCalldata(bytes memory _data) private pure returns (uint256 number) {
        bytes4 method;
        assembly {
            method := mload(add(_data, 32))
        }
        require(method == funcSelector, "wrong method");
        bytes memory argdata = BytesLib.slice(_data, 4, _data.length - 4);
        number = abi.decode(argdata, (uint256));
    }

    function setReceiptVerifier(ITxVerifier _txVerifier) external onlyOwner {
        txVerifier = _txVerifier;
    }

    function setSrcContract(uint64 _chainId, address _contract) external onlyOwner {
        srcContract[_chainId] = _contract;
    }
}
