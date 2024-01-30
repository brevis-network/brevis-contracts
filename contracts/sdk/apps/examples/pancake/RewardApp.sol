// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "../../framework/BrevisApp.sol";
import "../../../interface/IBrevisProof.sol";

contract RewardApp is BrevisApp, Ownable {
    using SafeERC20 for IERC20;

    bytes32 public vkHash;
    bytes32 public rewardsMerkleRoot; //leaf (address user, uint64 fromEpoch, uint amount) 
    uint64 public rewardsToEpoch;
    address public rewardToken;
    event Claimed(address indexed user, uint64 fromEpoch, uint64 toEpoch, uint256 amount);

    mapping(address => uint64) userClaimedTo; // user => toEpoch

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

        (rewardsToEpoch, rewardsMerkleRoot) = decodeOutput(_circuitOutput);
    }

    function decodeOutput(bytes calldata o) internal pure returns (uint64 epoch, bytes32 merkleRoot) {
        return (uint64(bytes8(o[0:8])), bytes32(o[8:40]));
    }

    function claim(
        address user, // msg.sender can claim on behalf of user
        uint256 amount,
        uint64 fromEpoch, 
        bytes32[] calldata merkleProof
    ) external {
        // Verifying proof
        bytes32 leaf = keccak256(abi.encode(user, fromEpoch, amount));
        require(_verifyProof(leaf, merkleProof), "not valid proof");
        uint64 lastClaimedEpoch = userClaimedTo[user];
        if (lastClaimedEpoch > 0) {
            require(fromEpoch == lastClaimedEpoch + 1, "illegal claim");
        }

        userClaimedTo[user] = rewardsToEpoch;
        IERC20(rewardToken).safeTransfer(user, amount);
        emit Claimed(user, fromEpoch, rewardsToEpoch, amount);
    }

    function _verifyProof(bytes32 leaf, bytes32[] memory proof) internal view returns (bool) {
        require(rewardsMerkleRoot != bytes32(0), "merkle root not set");
        bytes32 currentHash = leaf;
        uint256 proofLength = proof.length;
        for (uint256 i; i < proofLength; ) {
            if (currentHash < proof[i]) {
                currentHash = keccak256(abi.encode(currentHash, proof[i]));
            } else {
                currentHash = keccak256(abi.encode(proof[i], currentHash));
            }
            unchecked {
                ++i;
            }
        }
        return currentHash == rewardsMerkleRoot;
    }

    function setVkHash(bytes32 _vkHash) external onlyOwner {
        vkHash = _vkHash;
    }

    function setRewardToken(address _rewardToken) external onlyOwner {
        rewardToken = _rewardToken;
    }
}
