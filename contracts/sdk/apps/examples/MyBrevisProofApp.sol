// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../framework/BrevisProofApp.sol";

contract MyBrevisProofApp is BrevisProofApp {
    event ReceiveProofResult(bytes32 _appVkHash, bytes32 _appOutputHash);

    constructor(address _brevisProof) {
        brevisProof = IBrevisProof(_brevisProof);
    }

    function checkBrevisProof(
        uint64 _chainId,
        bytes calldata _proof,
        bytes calldata _appOutput,
        bytes32 _appVkHash
    ) external {
        _checkBrevisProof(_chainId, _proof, _appOutput, _appVkHash);
        bytes32 outputHash = keccak256(_appOutput);
        emit ReceiveProofResult(_appVkHash, outputHash);
    }
}
