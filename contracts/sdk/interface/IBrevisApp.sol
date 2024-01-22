// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../lib/Lib.sol";

interface IBrevisApp {
    function brevisCallback(bytes32 _requestId, bytes calldata _appCircuitOutput) external;
}
