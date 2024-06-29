// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../../framework/BrevisApp.sol";
import "../../../../safeguard/Ownable.sol";

contract MyBrevisApp is BrevisApp, Ownable {
    event ReceiveProofResult(bytes32 vkHash, bytes32 outputHash, bool zkAttested);

    constructor(address _brevisRequest) BrevisApp(_brevisRequest) {}

    function handleProofResult(bytes32 _vkHash, bytes calldata _appCircuitOutput) internal override {
        emit ReceiveProofResult(_vkHash, keccak256(_appCircuitOutput), true);
    }

    function handleOpProofResult(bytes32 _vkHash, bytes calldata _appCircuitOutput) internal override {
        emit ReceiveProofResult(_vkHash, keccak256(_appCircuitOutput), false);
    }

    function setOpChallengeWindow(uint256 _challangeWindow) external onlyOwner {
        _setOpChallengeWindow(_challangeWindow);
    }

    function setBrevisRequest(address _brevisRequest) external onlyOwner {
        _setBrevisRequest(_brevisRequest);
    }
}
