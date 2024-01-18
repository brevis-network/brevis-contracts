// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../verifiers/interfaces/IZkpVerifier.sol";

contract MockZkVerifier is IZkpVerifier {
    function verifyRaw(bytes calldata) external pure returns (bool r) {
        return true;
    }
}
