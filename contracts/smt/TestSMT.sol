// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/access/Ownable.sol";
import "../light-client-eth/interfaces/IAnchorBlocks.sol";
import "../interfaces/ISMT.sol";
import "./SMT.sol";

contract TestSMT is SMT {
    constructor(
        uint64[] memory _chainIds,
        address[] memory _anchorProviders,
        address[] memory _verifiers,
        bytes32[] memory _initRoots
    ) SMT(_chainIds, _anchorProviders, _verifiers, _initRoots) {}

    // function for testing convenience
    function addRootForTesting(uint64 chainId, bytes32 newRoot, uint64 endBlockNum) external onlyOwner {
        uint8 curIndex = curBufferIndices[chainId];
        curIndex = (curIndex + 1) % BUFFER_SIZE;
        smtRoots[chainId][curIndex] = newRoot;
        curBufferIndices[chainId] = curIndex;
        emit SmtRootUpdated(newRoot, endBlockNum, curIndex);
    }
}
