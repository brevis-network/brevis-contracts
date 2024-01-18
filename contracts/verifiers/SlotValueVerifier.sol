// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./interfaces/ISlotValueVerifier.sol";
import "./interfaces/IZkpVerifier.sol";
import "../chunk-sync/interfaces/IBlockChunks.sol";

contract SlotValueVerifier is ISlotValueVerifier, Ownable {
    uint32 constant PUBLIC_BYTES_START_IDX = 10 * 32;

    // retrieved from proofData, to align the fields with circuit...
    struct ProofData {
        bytes32 blkHash;
        bytes32 addrHash;
        bytes32 slotKeyHash;
        bytes32 slotValue;
        uint32 blkNum;
    }

    mapping(uint64 => address) public verifierAddresses; // chainid => snark verifier contract address
    address public BlockChunks;

    event UpdateVerifierAddress(uint64 chainId, address newAddress);
    event UpdateBlockChunks(address newAddress);

    constructor(address _blocChunks) {
        BlockChunks = _blocChunks;
    }

    function updateVerifierAddress(uint64 _chainId, address _verifierAddress) external onlyOwner {
        verifierAddresses[_chainId] = _verifierAddress;
        emit UpdateVerifierAddress(_chainId, _verifierAddress);
    }

    function updateBlockChunks(address _BlockChunks) external onlyOwner {
        BlockChunks = _BlockChunks;
        emit UpdateBlockChunks(_BlockChunks);
    }

    function verifySlotValue(
        uint64 chainId,
        bytes calldata proofData,
        bytes calldata blkVerifyInfo
    ) external view returns (SlotInfo memory slotInfo) {
        require(verifyRaw(chainId, proofData));

        (bytes32 prevHash, uint32 numFinal, bytes32[7] memory merkleProof) = getFromBlkVerifyInfo(blkVerifyInfo);
        ProofData memory data = getProofData(proofData);

        IBlockChunks.BlockHashWitness memory witness = IBlockChunks.BlockHashWitness({
            chainId: chainId,
            blkNum: data.blkNum,
            claimedBlkHash: data.blkHash,
            prevHash: prevHash,
            numFinal: numFinal,
            merkleProof: merkleProof
        });
        require(IBlockChunks(BlockChunks).isBlockHashValid(witness), "invalid blkHash");

        slotInfo.chainId = chainId;
        slotInfo.blkHash = data.blkHash;
        slotInfo.addrHash = data.addrHash;
        slotInfo.blkNum = data.blkNum;
        slotInfo.slotKeyHash = data.slotKeyHash;
        slotInfo.slotValue = data.slotValue;
    }

    function verifyRaw(uint64 chainId, bytes calldata proofData) private view returns (bool) {
        require(verifierAddresses[chainId] != address(0), "chain verifier not set");
        return (IZkpVerifier)(verifierAddresses[chainId]).verifyRaw(proofData);
    }

    function getFromBlkVerifyInfo(
        bytes calldata blkVerifyInfo
    ) internal pure returns (bytes32 prevHash, uint32 numFinal, bytes32[7] memory merkleProof) {
        require(blkVerifyInfo.length == 8 * 32 + 4, "incorrect blkVerifyInfo");
        prevHash = bytes32(blkVerifyInfo[:32]);
        numFinal = uint32(bytes4(blkVerifyInfo[32:36]));

        for (uint8 idx = 0; idx < 6; idx++) {
            merkleProof[idx] = bytes32(blkVerifyInfo[36 + 32 * idx:36 + 32 * (idx + 1)]);
        }

        merkleProof[6] = bytes32(blkVerifyInfo[36 + 32 * 6:36 + 32 * (6 + 1)]);
    }

    // groth16 proof + public inputs
    // public inputs:
    //  block hash
    //  contractAddrHash
    //  slot key
    //  slot value
    //  block number
    function getProofData(bytes calldata proofData) internal pure returns (ProofData memory data) {
        data.blkHash = bytes32(
            (uint256(bytes32(proofData[PUBLIC_BYTES_START_IDX:PUBLIC_BYTES_START_IDX + 32])) << 128) |
                uint128(bytes16(proofData[PUBLIC_BYTES_START_IDX + 32 + 16:PUBLIC_BYTES_START_IDX + 2 * 32]))
        );
        data.addrHash = bytes32(
            (uint256(bytes32(proofData[PUBLIC_BYTES_START_IDX + 2 * 32:PUBLIC_BYTES_START_IDX + 3 * 32])) << 128) |
                uint128(bytes16(proofData[PUBLIC_BYTES_START_IDX + 3 * 32 + 16:PUBLIC_BYTES_START_IDX + 4 * 32]))
        );
        data.slotKeyHash = bytes32(
            (uint256(bytes32(proofData[PUBLIC_BYTES_START_IDX + 4 * 32:PUBLIC_BYTES_START_IDX + 5 * 32])) << 128) |
                uint128(bytes16(proofData[PUBLIC_BYTES_START_IDX + 5 * 32 + 16:PUBLIC_BYTES_START_IDX + 6 * 32]))
        );
        data.slotValue = bytes32(
            (uint256(bytes32(proofData[PUBLIC_BYTES_START_IDX + 6 * 32:PUBLIC_BYTES_START_IDX + 7 * 32])) << 128) |
                uint128(bytes16(proofData[PUBLIC_BYTES_START_IDX + 7 * 32 + 16:PUBLIC_BYTES_START_IDX + 8 * 32]))
        );
        data.blkNum = uint32(bytes4(proofData[PUBLIC_BYTES_START_IDX + 9 * 32 - 4:PUBLIC_BYTES_START_IDX + 9 * 32]));
    }
}
