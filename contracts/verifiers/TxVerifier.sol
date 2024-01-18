// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/access/Ownable.sol";
import "solidity-rlp/contracts/RLPReader.sol";
import "./interfaces/ITxVerifier.sol";
import "./interfaces/IZkpVerifier.sol";
import "../chunk-sync/interfaces/IBlockChunks.sol";

contract TxVerifier is ITxVerifier, Ownable {
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
        bytes leafRlpPrefix;
    }

    mapping(uint64 => address) public verifierAddresses; // chainid => snark verifier contract address
    address public blockChunks;

    event UpdateVerifierAddress(uint64 chainId, address newAddress);
    event UpdateBlockChunks(address newAddress);
    event VerifiedTx(uint64 chainId, bytes32 txHash);

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

    function verifyTxAndLog(
        bytes calldata txRaw,
        bytes calldata proofData,
        bytes calldata auxiBlkVerifyInfo
    ) external returns (TxInfo memory info) {
        info = verifyTx(txRaw, proofData, auxiBlkVerifyInfo);
        emit VerifiedTx(info.chainId, keccak256(txRaw));
    }

    function verifyTx(
        bytes calldata txRaw,
        bytes calldata proofData,
        bytes calldata auxiBlkVerifyInfo
    ) public view returns (TxInfo memory info) {
        info = decodeTx(txRaw);
        require(verifyRaw(info.chainId, proofData), "proof not valid");
        ProofData memory data = getProofData(proofData);
        bytes memory leafRlp = bytes.concat(data.leafRlpPrefix, txRaw);
        bytes32 leafHash = keccak256(leafRlp);
        require(leafHash == data.leafHash, "leafHash not match");

        (bytes32 prevHash, uint32 numFinal, bytes32[7] memory merkleProof) = getFromAuxiBlkVerifyInfo(
            auxiBlkVerifyInfo
        );

        IBlockChunks.BlockHashWitness memory witness = IBlockChunks.BlockHashWitness({
            chainId: info.chainId,
            blkNum: data.blkNum,
            claimedBlkHash: data.blkHash,
            prevHash: prevHash,
            numFinal: numFinal,
            merkleProof: merkleProof
        });
        require(IBlockChunks(blockChunks).isBlockHashValid(witness), "invalid blkHash");

        info.blkHash = data.blkHash;
        info.blkTime = data.blkTime;
        info.blkNum = data.blkNum;
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
    function decodeTx(bytes calldata txRaw) public pure returns (TxInfo memory info) {
        uint8 txType = uint8(txRaw[0]);
        require(txType == 2, "not a DynamicFeeTxType");

        bytes memory rlpData = txRaw[1:];
        RLPReader.RLPItem[] memory values = rlpData.toRlpItem().toList();
        info.chainId = uint64(values[0].toUint());
        info.nonce = uint64(values[1].toUint());
        info.gasTipCap = values[2].toUint();
        info.gasFeeCap = values[3].toUint();
        info.gas = values[4].toUint();
        info.to = values[5].toAddress();
        info.value = values[6].toUint();
        info.data = values[7].toBytes();

        (uint8 v, bytes32 r, bytes32 s) = (
            uint8(values[9].toUint()),
            bytes32(values[10].toBytes()),
            bytes32(values[11].toBytes())
        );
        // remove r,s,v and adjust length field
        bytes memory unsignedTxRaw;
        uint16 unsignedTxRawDataLength;
        uint8 prefix = uint8(txRaw[1]);
        uint8 lenBytes = prefix - 0xf7; // assume lenBytes won't larger than 2, means the tx rlp data size won't exceed 2^16
        if (lenBytes == 1) {
            unsignedTxRawDataLength = uint8(bytes1(txRaw[2:3])) - 67; //67 is the bytes of r,s,v
        } else {
            unsignedTxRawDataLength = uint16(bytes2(txRaw[2:2 + lenBytes])) - 67;
        }
        if (unsignedTxRawDataLength <= 55) {
            unsignedTxRaw = abi.encodePacked(txRaw[:2], txRaw[3:txRaw.length - 67]);
            unsignedTxRaw[1] = bytes1(0xc0 + uint8(unsignedTxRawDataLength));
        } else {
            if (unsignedTxRawDataLength <= 255) {
                unsignedTxRaw = abi.encodePacked(
                    txRaw[0],
                    bytes1(0xf8),
                    bytes1(uint8(unsignedTxRawDataLength)),
                    txRaw[2 + lenBytes:txRaw.length - 67]
                );
            } else {
                unsignedTxRaw = abi.encodePacked(
                    txRaw[0],
                    bytes1(0xf9),
                    bytes2(unsignedTxRawDataLength),
                    txRaw[2 + lenBytes:txRaw.length - 67]
                );
            }
        }
        info.from = recover(keccak256(unsignedTxRaw), r, s, v);
    }

    function recover(bytes32 message, bytes32 r, bytes32 s, uint8 v) internal pure returns (address) {
        if (v < 27) {
            v += 27;
        }
        return ecrecover(message, v, r, s);
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
        data.leafRlpPrefix = bytes(proofData[PUBLIC_BYTES_START_IDX + 6 * 32:]);
    }
}
