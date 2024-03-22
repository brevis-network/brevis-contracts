pragma solidity ^0.8.18;

interface ISwapVolume {
    function getAttestedSwapSumVolume(address swapper) external view returns (uint256 volume);
}
