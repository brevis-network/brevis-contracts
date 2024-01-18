// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/access/Ownable.sol";

import "./interfaces/IEthereumLightClient.sol";
import "./interfaces/IAnchorBlocks.sol";

import "./common/Helpers.sol";
import "./common/Constants.sol";
import "./common/Types.sol";

uint256 constant EXECUTION_BLOCK_LEFT_PREFIX_LEN = 4;

contract AnchorBlocks is IAnchorBlocks, Ownable {
    // BlockHashWitness is the RLP code that witnesses the generation of block hash given the ParentHash field
    struct BlockHashWitness {
        bytes left;
        bytes right;
    }

    event AnchorBlockUpdated(uint256 blockNum, bytes32 blockHash);

    IEthereumLightClient public lightClient;
    // execution block number => execution block hash
    mapping(uint256 => bytes32) public blocks;
    uint256 public latestBlockNum;

    constructor(address _lightClient) {
        lightClient = IEthereumLightClient(_lightClient);
    }

    /// @notice Updates an "anchor block" of a specific block number to the contract state
    function processUpdate(LightClientOptimisticUpdate memory hb) external {
        (uint256 blockNum, bytes32 blockHash) = verifyHeadBlock(hb);
        require(blockHash != bytes32(0), "empty blockHash");
        doUpdate(blockNum, blockHash);
    }

    /// @notice Updates an "anchor block" of a specific block number to the contract state
    /// @dev It is possible that an attested block doesn't collect enough sync committee signatures in its corresponding
    /// signature slot and thus cannot be used in an anchor update. In that case, the updater can pick a later block
    /// that has enough sigs, and supply a chainProof to show that the block they want to sync can chain to the head block.
    function processUpdateWithChainProof(
        LightClientOptimisticUpdate memory hb,
        bytes32 blockHash,
        BlockHashWitness[] memory chainProof
    ) external {
        require(chainProof.length > 0, "invalid proof length");
        (uint256 headBlockNum, bytes32 headBlockHash) = verifyHeadBlock(hb);
        uint256 blockNum = headBlockNum - chainProof.length;
        verifyChainProof(blockHash, chainProof, headBlockHash);
        doUpdate(blockNum, blockHash);
    }

    function verifyHeadBlock(LightClientOptimisticUpdate memory hb) private view returns (uint256, bytes32) {
        require(hasSupermajority(hb.syncAggregate.participation), "quorum not reached");
        verifyExecutionPayload(hb.attestedHeader);
        lightClient.verifyCommitteeSignature(hb.signatureSlot, hb.attestedHeader.beacon, hb.syncAggregate);
        HeaderWithExecution memory h = hb.attestedHeader;
        uint256 blockNum = Helpers.revertEndian(uint256(h.execution.blockNumber.leaf));
        return (blockNum, h.execution.blockHash.leaf);
    }

    function verifyExecutionPayload(HeaderWithExecution memory h) private pure {
        bool valid = Helpers.isValidMerkleBranch(h.executionRoot, EXECUTION_PAYLOAD_ROOT_INDEX, h.beacon.bodyRoot);
        require(valid, "bad exec root proof");
        verifyMerkleProof(h.execution.blockNumber, EXECUTION_BLOCK_NUMBER_LOCAL_INDEX, h.executionRoot.leaf);
        verifyMerkleProof(h.execution.blockHash, EXECUTION_BLOCK_HASH_LOCAL_INDEX, h.executionRoot.leaf);
    }

    function doUpdate(uint256 blockNum, bytes32 blockHash) private {
        require(blocks[blockNum] == bytes32(0), "block hash already exists");
        blocks[blockNum] = blockHash;
        if (blockNum > latestBlockNum) {
            latestBlockNum = blockNum;
        }
        emit AnchorBlockUpdated(blockNum, blockHash);
    }

    function verifyChainProof(
        bytes32 blockHash,
        BlockHashWitness[] memory chainProof,
        bytes32 headBlockHash
    ) private pure {
        bytes32 h = blockHash;
        for (uint256 i = 0; i < chainProof.length; i++) {
            // small hack to save some RLP encoding:
            // We only care about whether the given blockHash can somehow combine with something to hash into headBlockHash.
            // The RLP oding of a block always has 3 bytes for total length prefix and 1 byte (0xa0) for bytes32's length
            // prefix; and the ParentHash field is always the first element. So there are always 8 bytes preceding ParentHash.
            require(chainProof[i].left.length == EXECUTION_BLOCK_LEFT_PREFIX_LEN, "invalid left len");
            h = keccak256(bytes.concat(chainProof[i].left, h, chainProof[i].right));
        }
        require(h == headBlockHash, "invalid chainProof");
    }

    function verifyMerkleProof(LeafWithBranch memory proof, uint256 index, bytes32 root) private pure {
        require(Helpers.isValidMerkleBranch(proof, index, root), "bad proof");
    }

    function hasSupermajority(uint64 participation) private pure returns (bool) {
        return participation * 3 >= SYNC_COMMITTEE_SIZE * 2;
    }

    function setLightClient(address _lightClient) external onlyOwner {
        lightClient = IEthereumLightClient(_lightClient);
    }
}
