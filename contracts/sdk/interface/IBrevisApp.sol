// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../lib/Lib.sol";

interface IBrevisApp {
    function brevisCallback(bytes32 _requestId, bytes32 _appVkHash, bytes calldata _appCircuitOutput) external;

    function brevisBatchCallback(
        uint64 _chainId,
        Brevis.ProofData[] calldata _proofDataArray,
        bytes[] calldata _appCircuitOutputs
    ) external;
}
