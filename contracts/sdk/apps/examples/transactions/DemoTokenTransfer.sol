// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../../framework/BrevisApp.sol";
import "../../../lib/Lib.sol";
import "../../../interface/IBrevisProof.sol";

// Single tx: Proof of USDC transfer
// Use Brevis to prove address #1 has ever made a USDC transfer to address #2.
// This example is useful in the social recovery of account abstraction where a friend
// with proven interaction can help recover the lost wallet.

contract DemoTokenTransfer is BrevisApp {
    address public token;
    bytes4 public transferSelector;

    // (sender, receiver) -> timestamp of latest tx
    mapping(address => mapping(address => uint64)) public txTimestamps;

    constructor(IBrevisProof _brevisProof, address _token, bytes4 _transferSelector) BrevisApp(_brevisProof) {
        token = _token;
        transferSelector = _transferSelector;
    }

    function submitUserTx(
        bytes32 _proofRequestId,
        uint64 _chainId,
        bytes32 _smtRoot,
        Brevis.TransactionInfo calldata _info,
        bytes calldata _txRaw
    ) external {
        Brevis.ExtractInfos memory info;
        info.smtRoot = _smtRoot;
        info.txs = new Brevis.TransactionInfo[](1);
        info.txs[0] = _info;

        validateRequest(_proofRequestId, _chainId, info);

        bytes memory leafRlp = bytes.concat(_info.leafRlpPrefix, _txRaw);
        bytes32 leafHash = keccak256(leafRlp);
        require(leafHash == _info.leafHash, "leafHash not match");

        Tx.TxInfo memory txInfo = Tx.decodeTx(_txRaw);
        require(txInfo.to == token, "not a token transfer");
        (bytes4 sig, address receiver, ) = abiDecodeTransfer(txInfo.data);
        require(sig == transferSelector, "not transfer selector");
        txTimestamps[txInfo.from][receiver] = _info.blockTime;
    }

    function abiDecodeTransfer(bytes memory _data) private pure returns (bytes4 sig, address receiver, uint256 amount) {
        assembly {
            sig := mload(add(_data, 32))
            receiver := mload(add(_data, 36))
            amount := mload(add(_data, 68))
        }
    }
}
