// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../interfaces/ISMT.sol";

contract MockSMT is ISMT {
    function updateRoot(uint64 chainId, SmtUpdate memory u) external {}

    function isSmtRootValid(uint64, bytes32) external pure returns (bool) {
        return true;
    }
}
