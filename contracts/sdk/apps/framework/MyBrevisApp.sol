// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "./BrevisApp.sol";

contract MyBrevisApp is BrevisApp {
    constructor(IBrevisProof _brevisProof) BrevisApp(_brevisProof) {}

    function handleProofResult(bytes32 _vkHash, bytes calldata _appCircuitOutput) internal override {}
}
