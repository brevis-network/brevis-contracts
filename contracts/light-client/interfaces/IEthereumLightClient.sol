// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../common/Types.sol";

interface IEthereumLightClient {
    function optimisticExecutionStateRootAndSlot() external view returns (bytes32 root, uint64 slot);

    function finalizedExecutionStateRootAndSlot() external view returns (bytes32 root, uint64 slot);

    // reverts if check fails
    function verifyCommitteeSignature(
        uint64 signatureSlot,
        BeaconBlockHeader memory header,
        SyncAggregate memory syncAggregate
    ) external view;
}
