// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../safeguard/Ownable.sol";

contract BrevisDummy is Ownable {
    event DummyEvent(uint64 data);

    function updateVerifierAddress(uint64 _data) external onlyOwner {
        emit DummyEvent(_data);
    }
}
