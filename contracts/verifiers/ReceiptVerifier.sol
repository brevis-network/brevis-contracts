// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/access/Ownable.sol";
import "solidity-rlp/contracts/RLPReader.sol";
import "./interfaces/IReceiptVerifier.sol";
import "./interfaces/IZkpVerifier.sol";
import "../chunk-sync/interfaces/IBlockChunks.sol";

contract ReceiptVerifier is IReceiptVerifier, Ownable {
    using RLPReader for bytes;
    using RLPReader for uint;
    using RLPReader for RLPReader.RLPItem;

    uint32 constant PUBLIC_BYTES_START_IDX = 10 * 32; // the first 10 32bytes are groth16 proof (A/B/C/Commitment)

    // retrieved from proofData, to align the fields with circuit...
    struct ProofData {
        bytes32 leafHash;
        bytes32 blkHash;
        uint32 blkNum;
        uint64 blkTime;
        uint64 chainId;
        bytes leafRlpPrefix; // not public input
    }

    mapping(uint64 => address) public verifierAddresses; // chainid => snark verifier contract address
    address public blockChunks;

    event UpdateVerifierAddress(uint64 chainId, address newAddress);
    event UpdateBlockChunks(address newAddress);
    event VerifiedReceipt(uint64 chainId, bytes32 receiptHash);

    constructor(address _blockChunks) {
        blockChunks = _blockChunks;
    }

    function updateVerifierAddress(uint64 _chainId, address _verifierAddress) external onlyOwner {
        verifierAddresses[_chainId] = _verifierAddress;
        emit UpdateVerifierAddress(_chainId, _verifierAddress);
    }

    function updateBlockChunks(address _blockChunks) external onlyOwner {
        blockChunks = _blockChunks;
        emit UpdateBlockChunks(_blockChunks);
    }

    function verifyReceiptAndLog(
        bytes calldata receiptRaw,
        bytes calldata proofData,
        bytes calldata auxiBlkVerifyInfo
    ) external returns (ReceiptInfo memory info) {
        info = verifyReceipt(receiptRaw, proofData, auxiBlkVerifyInfo);
        // TODO, will emit <chain_id, block_num, tx_index> instead
        emit VerifiedReceipt(info.chainId, keccak256(receiptRaw));
    }

    function verifyReceipt(
        bytes calldata receiptRaw,
        bytes calldata proofData,
        bytes calldata auxiBlkVerifyInfo
    ) public view returns (ReceiptInfo memory info) {
        ProofData memory data = getProofData(proofData);
        require(verifyRaw(data.chainId, proofData), "proof not valid");
        bytes memory leafRlp = bytes.concat(data.leafRlpPrefix, receiptRaw);
        bytes32 leafHash = keccak256(leafRlp);
        require(leafHash == data.leafHash, "leafHash not match");

        (bytes32 prevHash, uint32 numFinal, bytes32[7] memory merkleProof) = getFromAuxiBlkVerifyInfo(
            auxiBlkVerifyInfo
        );

        IBlockChunks.BlockHashWitness memory witness = IBlockChunks.BlockHashWitness({
            chainId: data.chainId,
            blkNum: data.blkNum,
            claimedBlkHash: data.blkHash,
            prevHash: prevHash,
            numFinal: numFinal,
            merkleProof: merkleProof
        });
        require(IBlockChunks(blockChunks).isBlockHashValid(witness), "invalid blkHash");

        info = decodeReceipt(receiptRaw);
        info.blkHash = data.blkHash;
        info.blkTime = data.blkTime;
        info.blkNum = data.blkNum;
        info.chainId = data.chainId;
    }

    function getFromAuxiBlkVerifyInfo(
        bytes calldata auxiBlkVerifyInfo
    ) internal pure returns (bytes32 prevHash, uint32 numFinal, bytes32[7] memory merkleProof) {
        require(auxiBlkVerifyInfo.length == 8 * 32 + 4, "incorrect auxiBlkVerifyInfo");

        prevHash = bytes32(auxiBlkVerifyInfo[:32]);
        numFinal = uint32(bytes4(auxiBlkVerifyInfo[32:36]));
        for (uint8 idx = 0; idx < 6; idx++) {
            merkleProof[idx] = bytes32(auxiBlkVerifyInfo[36 + 32 * idx:36 + 32 * (idx + 1)]);
        }
        merkleProof[6] = bytes32(auxiBlkVerifyInfo[36 + 32 * 6:36 + 32 * (6 + 1)]);
    }

    // support DynamicFeeTxType for now
    function decodeReceipt(bytes calldata receiptRaw) public pure returns (ReceiptInfo memory info) {
        uint8 txType = uint8(receiptRaw[0]);
        require(txType == 2, "not a DynamicFeeTxType");
        bytes memory rlpData = receiptRaw[1:];
        RLPReader.RLPItem[] memory values = rlpData.toRlpItem().toList();
        if (bytes1(values[0].toBytes()) == 0x01) {
            info.success = true;
        }

        RLPReader.RLPItem[] memory rlpLogs = values[3].toList();
        LogInfo[] memory logInfos = new LogInfo[](rlpLogs.length);
        for (uint8 i = 0; i < rlpLogs.length; i++) {
            RLPReader.RLPItem[] memory log = rlpLogs[i].toList();
            //let one = log[0].toBytes();
            logInfos[i].addr = log[0].toAddress();
            RLPReader.RLPItem[] memory topics = log[1].toList();
            logInfos[i].topics = new bytes32[](topics.length);
            for (uint8 j = 0; j < topics.length; j++) {
                logInfos[i].topics[j] = bytes32(topics[j].toBytes());
            }
            logInfos[i].data = log[2].toBytes();
        }
        info.logs = logInfos;
    }

    function verifyRaw(uint64 chainId, bytes calldata proofData) private view returns (bool) {
        require(verifierAddresses[chainId] != address(0), "chain verifier not set");
        return (IZkpVerifier)(verifierAddresses[chainId]).verifyRaw(proofData);
    }

    function getProofData(bytes calldata proofData) internal pure returns (ProofData memory data) {
        data.leafHash = bytes32(
            (uint256(bytes32(proofData[PUBLIC_BYTES_START_IDX:PUBLIC_BYTES_START_IDX + 32])) << 128) |
                uint128(bytes16(proofData[PUBLIC_BYTES_START_IDX + 32 + 16:PUBLIC_BYTES_START_IDX + 2 * 32]))
        );
        data.blkHash = bytes32(
            (uint256(bytes32(proofData[PUBLIC_BYTES_START_IDX + 2 * 32:PUBLIC_BYTES_START_IDX + 3 * 32])) << 128) |
                uint128(bytes16(proofData[PUBLIC_BYTES_START_IDX + 3 * 32 + 16:PUBLIC_BYTES_START_IDX + 4 * 32]))
        );
        data.blkNum = uint32(bytes4(proofData[PUBLIC_BYTES_START_IDX + 5 * 32 - 4:PUBLIC_BYTES_START_IDX + 5 * 32]));
        data.blkTime = uint64(bytes8(proofData[PUBLIC_BYTES_START_IDX + 6 * 32 - 8:PUBLIC_BYTES_START_IDX + 6 * 32]));
        // not public input
        data.chainId = uint64(bytes8(proofData[PUBLIC_BYTES_START_IDX + 6 * 32:PUBLIC_BYTES_START_IDX + 6 * 32 + 8]));
        data.leafRlpPrefix = bytes(proofData[PUBLIC_BYTES_START_IDX + 6 * 32 + 8:]);
    }
}
