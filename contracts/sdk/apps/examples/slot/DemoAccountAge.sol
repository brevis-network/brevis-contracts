// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../../framework/BrevisApp.sol";
import "../../../lib/Lib.sol";
import "../../../interface/IBrevisProof.sol";

// Single storage slot: Account age proof
// Use Brevis to prove the age of an address by proving its oldest storage slot.

contract DemoAccountAge is BrevisApp {
    mapping(address => uint64) public visibleSinceBlocks;

    constructor(IBrevisProof _brevisProof) BrevisApp(_brevisProof) {}

    function submitUserStorageInfo(
        bytes32 _proofRequestId,
        uint64 _chainId,
        bytes32 _smtRoot,
        Brevis.StorageInfo calldata _info
    ) external {
        Brevis.ExtractInfos memory info;
        info.smtRoot = _smtRoot;
        info.stores = new Brevis.StorageInfo[](1);
        info.stores[0] = _info;

        validateRequest(_proofRequestId, _chainId, info);

        visibleSinceBlocks[_info.account] = _info.blockNumber;
    }

    function getAccountAge(address _account) external view returns (uint64) {
        uint64 visibleSinceBlock = visibleSinceBlocks[_account];
        require(visibleSinceBlock > 0, "no proof yet");

        return uint64(block.number) - visibleSinceBlock; // if the proof is from another chain, then should get current block number from SMT
    }
}
