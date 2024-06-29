// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../../framework/BrevisApp.sol";
import "../../../../safeguard/Ownable.sol";

contract SlotValueExample is BrevisApp, Ownable {
    event PastOwnerAttested(address contractAddr, address ownerAddr, uint64 blockNum);

    bytes32 public vkHash;

    constructor(address _brevisRequest) BrevisApp(_brevisRequest) {}

    // BrevisQuery contract will call our callback once Brevis backend submits the proof.
    // This method is called with once the proof is verified.
    function handleProofResult(bytes32 _vkHash, bytes calldata _circuitOutput) internal override {
        // We need to check if the verifying key that Brevis used to verify the proof
        // generated by our circuit is indeed our designated verifying key. This proves
        // that the _circuitOutput is authentic
        require(vkHash == _vkHash, "invalid vk");
        (address contractAddr, address ownerAddr, uint64 blockNum) = decodeOutput(_circuitOutput);
        emit PastOwnerAttested(contractAddr, ownerAddr, blockNum);
    }

    // In guest circuit we have:
    // api.OutputAddress(s.Contract)
    // api.OutputAddress(owner)
    // api.OutputUint(64, s.BlockNum)
    function decodeOutput(bytes calldata o) internal pure returns (address, address, uint64) {
        address contractAddr = address(bytes20(o[0:20]));
        address ownerAddr = address(bytes20(o[20:40]));
        uint64 blockNum = uint64(bytes8(o[40:48]));
        return (contractAddr, ownerAddr, blockNum);
    }

    function setVkHash(bytes32 _vkHash) external onlyOwner {
        vkHash = _vkHash;
    }
}
