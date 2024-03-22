// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/access/Ownable.sol";

import "../../framework/BrevisApp.sol";
import "../../../interface/IBrevisProof.sol";
import "./ISwapVolume.sol";

contract IntegralSwapVolume is BrevisApp, Ownable, ISwapVolume {
    event SwapVolumeAttested(address user, uint256 volume);

    bytes32 public vkHash;

    constructor(address brevisProof) BrevisApp(IBrevisProof(brevisProof)) {}

    mapping(address => uint256) public swapVolumes;

    // BrevisQuery contract will call our callback once Brevis backend submits the proof.
    function handleProofResult(
        bytes32 /*_requestId*/,
        bytes32 _vkHash,
        bytes calldata _circuitOutput
    ) internal override {
        require(vkHash == _vkHash, "invalid vk");

        (address userAddr, uint256 swapVolume) = decodeOutput(_circuitOutput);
        swapVolumes[userAddr] = swapVolume;
        emit SwapVolumeAttested(userAddr, swapVolume);
    }

    function decodeOutput(bytes calldata o) internal pure returns (address, uint256) {
        address userAddr = address(bytes20(o[0:20]));
        uint256 swapVolume = uint256(bytes32(o[20:51])); // swapVolume is packed as a uint248 but we cast it to uint256 here
        return (userAddr, swapVolume);
    }

    function setVkHash(bytes32 _vkHash) external onlyOwner {
        vkHash = _vkHash;
    }

    function getAttestedSwapSumVolume(address swapper) external view returns (uint256) {
        return swapVolumes[swapper];
    }
}
