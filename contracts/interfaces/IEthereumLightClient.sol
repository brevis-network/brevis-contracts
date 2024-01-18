// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

interface IEthereumLightClient {
    function finalizedExecutionStateRootAndSlot() external view returns (bytes32 root, uint64 slot);

    function optimisticExecutionStateRootAndSlot() external view returns (bytes32 root, uint64 slot);
}
