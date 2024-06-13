// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../lib/Lib.sol";

interface IBrevisApp {
    function brevisCallback(bytes32 _requestId, bytes32 _appVkHash, bytes calldata _appCircuitOutput) external;

    function brevisBatchCallback(
        bytes32[] calldata _requestIds,
        bytes32[] calldata _appVkHashs,
        bytes[] calldata _appCircuitOutputs
    ) external;
}
