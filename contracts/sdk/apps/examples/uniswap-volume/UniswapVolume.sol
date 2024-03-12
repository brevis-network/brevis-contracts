// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/access/Ownable.sol";

import "../../framework/BrevisApp.sol";
import "../../../interface/IBrevisProof.sol";

contract UniswapVolume is BrevisApp, Ownable {
    event SwapVolumeAttested(address user, uint64 sinceBlockNum, uint256 volume);

    bytes32 public vkHash;

    constructor(address brevisProof) BrevisApp(IBrevisProof(brevisProof)) {}

    // BrevisQuery contract will call our callback once Brevis backend submits the proof.
    function handleProofResult(
        bytes32 /*_requestId*/,
        bytes32 _vkHash,
        bytes calldata _circuitOutput
    ) internal override {
        // We need to check if the verifying key that Brevis used to verify the proof generated by our circuit is indeed
        // our designated verifying key. This proves that the _circuitOutput is authentic
        require(vkHash == _vkHash, "invalid vk");

        (uint256 sumVolume, uint64 minBlockNum, address userAddr) = decodeOutput(_circuitOutput);

        emit SwapVolumeAttested(userAddr, minBlockNum, sumVolume);
    }

    // In guest circuit we have:
    // api.OutputUint256(Salt)
    // api.OutputUint(248, sumVolume)
    // api.OutputUint(64, minBlockNum)
    // api.OutputAddress(c.UserAddr)
    function decodeOutput(bytes calldata o) internal pure returns (uint256, uint64, address) {
        uint256 sumVolume = uint256(bytes32(o[32:63])); // sumVolume is packed as a uint248 but we cast it to uint256 here
        uint64 minBlockNum = uint64(bytes8(o[63:71])); // minBlockNum is uint64 (8 bytes) field in the output
        address userAddr = address(bytes20(o[71:91])); // c.UserAddr is an address (20 bytes) field in the output
        return (sumVolume, minBlockNum, userAddr);
    }

    function setVkHash(bytes32 _vkHash) external onlyOwner {
        vkHash = _vkHash;
    }
}