// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

// post a number at source chain
// to test event and tx verification at dst chain
contract PostNumber {
    event SendNumber(address from, uint256 number);

    function sendNumber(uint256 number) external {
        emit SendNumber(msg.sender, number);
    }
}
