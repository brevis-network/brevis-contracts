// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../framework/SimpleReceiptApp.sol";
import "../../lib/Lib.sol";
import "../../interface/IBrevisProof.sol";

contract TestBrevisApp is SimpleReceiptApp {
    constructor(IBrevisProof _brevisProof) BrevisApp(_brevisProof) {}

    event Success();

    function submit(
        bytes32 _proofRequestId,
        uint64 _chainId,
        bytes32 _smtRoot,
        Brevis.ReceiptInfo calldata _rinfo,
        Brevis.StorageInfo calldata _sinfo,
        Brevis.TransactionInfo calldata _tinfo
    ) external {
        Brevis.ExtractInfos memory info;
        info.smtRoot = _smtRoot;
        info.receipts = new Brevis.ReceiptInfo[](1);
        info.receipts[0] = _rinfo;
        info.stores = new Brevis.StorageInfo[](1);
        info.stores[0] = _sinfo;
        info.txs = new Brevis.TransactionInfo[](1);
        info.txs[0] = _tinfo;

        validateRequest(_proofRequestId, _chainId, info);

        emit Success();
    }
}
