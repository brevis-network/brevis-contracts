// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

interface IUniswapSumVolume {
    function getAttestedSwapSumVolume(address swapper) external view returns (uint256 volume);
}
