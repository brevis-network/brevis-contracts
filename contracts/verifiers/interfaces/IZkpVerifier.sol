// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

interface IZkpVerifier {
    function verifyRaw(bytes calldata proofData) external view returns (bool r);
}
