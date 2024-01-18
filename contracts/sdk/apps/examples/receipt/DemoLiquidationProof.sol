// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../../framework/SimpleReceiptApp.sol";
import "../../../lib/Lib.sol";
import "../../../interface/IBrevisProof.sol";

// Single receipt: Proof of liquidation
// Use Brevis to prove that an address has been liquidated in Compound v2.
// Specifically, prove there was a LiquidateBorrow event emitted by Compound v2 contract for an borrower address.

contract DemoLiquidationProof is SimpleReceiptApp {
    mapping(address => bool) public liquiationFlags;

    constructor(IBrevisProof _brevisProof) BrevisApp(_brevisProof) {}

    // should set correct FieldLocation in BrevisApp.setFieldLocations() to restrict
    // the first field address as from CompoundV2 and the topic as LiquidateBorrow
    function submitUserEvent(
        bytes32 _proofRequestId,
        uint64 _chainId,
        bytes32 _smtRoot,
        Brevis.ReceiptInfo calldata _info
    ) external {
        Brevis.ExtractInfos memory info;
        info.smtRoot = _smtRoot;
        info.receipts = new Brevis.ReceiptInfo[](1);
        info.receipts[0] = _info;

        validateRequest(_proofRequestId, _chainId, info);

        address borrower = address(bytes20(_info.logs[0].value));
        liquiationFlags[borrower] = true;
    }
}
