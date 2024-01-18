// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../bsc-tendermint/interfaces/IBSCValidatorSet.sol";
import "../interfaces/IEthereumLightClient.sol";

import "./libraries/ECDSA.sol";
import "./libraries/Memory.sol";
import "./libraries/RLPWriter.sol";

// Sample header
// curl --location --request POST 'https://bsc.getblock.io/API_KEY/testnet/'
//  -H "Content-Type: application/json"
//  --data-raw '{"jsonrpc": "2.0", "method": "eth_getBlockByNumber", "params": ["0x68B3", true], "id": "getblock.io"}'
// {
// "difficulty":"0x2",
// "extraData":"0xd983010000846765746889676f312e31322e3137856c696e7578000000000000c3daa60d95817e2789de3eafd44dc354fe804bf5f08059cde7c86bc1215941d022bf9609ca1dee2881baf2144aa93fc80082e6edd0b9f8eac16f327e7d59f16500",
// "gasLimit":"0x1c9c380",
// "gasUsed":"0x0",
// "hash":"0xc3fa2927a8e5b7cfbd575188a30c34994d3356607deb4c10d7fefe0dd5cdcc83",
// "logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
// "miner":"0x35552c16704d214347f29fa77f77da6d75d7c752",
// "mixHash":"0x0000000000000000000000000000000000000000000000000000000000000000",
// "nonce":"0x0000000000000000",
// "number":"0x68b3",
// "parentHash":"0xbf4d16769b8fd946394957049eef29ed938da92454762fc6ac65e0364ea004c7",
// "receiptsRoot":"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
// "sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
// "size":"0x261",
// "stateRoot":"0x7b5a72075082c31ec909afe5c5df032b6e7f19c686a9a408a2cb6b75dec072a3",
// "timestamp":"0x5f080818",
// "totalDifficulty":"0xd167",
// "transactions":[],
// "transactionsRoot":"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
// "uncles":[]
// }

contract PoALightClient is IEthereumLightClient {
    using Memory for bytes;

    struct BNBHeaderInfo {
        bytes32 parentHash;
        bytes32 sha3Uncles;
        address miner;
        bytes32 stateRoot;
        bytes32 transactionsRoot;
        bytes32 receiptsRoot;
        bytes logsBloom;
        uint256 difficulty;
        uint256 number;
        uint64 gasLimit;
        uint64 gasUsed;
        uint64 timestamp;
        bytes extraData;
        bytes32 mixHash;
        bytes8 nonce;
    }

    IBSCValidatorSet public bscValidatorSet;

    uint64 number;
    bytes32 stateRoot;

    constructor(address _bscValidatorSet) {
        bscValidatorSet = IBSCValidatorSet(_bscValidatorSet);
        number = 0;
        stateRoot = hex"0000000000000000000000000000000000000000000000000000000000000000";
    }

    // Fixed number of extra-data prefix bytes reserved for signer vanity.
    // https://eips.ethereum.org/EIPS/eip-225
    uint256 private constant EXTRA_VANITY_LENGTH = 32;

    // Length of signer's signature
    uint256 private constant SIGNATURE_LENGTH = 65;

    uint64 private constant CHAIN_ID = 97;

    function updateHeader(BNBHeaderInfo calldata header) external {
        require(header.number > number, "PoALightClient: invalid block number");

        address signer = _retrieveSignerInfo(header);
        require(bscValidatorSet.isCurrentValidator(signer), "PoALightClient: invalid signer address");

        number = uint64(header.number);
        stateRoot = header.stateRoot;
    }

    function _retrieveSignerInfo(BNBHeaderInfo calldata header) internal pure returns (address signer) {
        bytes memory extraData = header.extraData;

        require(extraData.length > EXTRA_VANITY_LENGTH, "PoALightClient: invalid extra data for vanity");
        require(
            extraData.length >= EXTRA_VANITY_LENGTH + SIGNATURE_LENGTH,
            "PoALightClient: invalid extra data for signature"
        );

        // data: [0, extraData.length - SIGNATURE_LENGTH)
        // signature: [extraData.length - SIGNATURE_LENGTH, extraData.length)
        bytes memory extraDataWithoutSignature = Memory.range(extraData, 0, extraData.length - SIGNATURE_LENGTH);
        bytes memory signature = Memory.range(extraData, extraData.length - SIGNATURE_LENGTH, extraData.length);

        require(signature.length == SIGNATURE_LENGTH, "PoALightClient: signature retrieval failed");
        BNBHeaderInfo memory unsignedHeader = BNBHeaderInfo({
            difficulty: header.difficulty,
            extraData: extraDataWithoutSignature,
            gasLimit: header.gasLimit,
            gasUsed: header.gasUsed,
            logsBloom: header.logsBloom,
            miner: header.miner,
            mixHash: header.mixHash,
            nonce: header.nonce,
            number: header.number,
            parentHash: header.parentHash,
            receiptsRoot: header.receiptsRoot,
            sha3Uncles: header.sha3Uncles,
            stateRoot: header.stateRoot,
            timestamp: header.timestamp,
            transactionsRoot: header.transactionsRoot
        });

        bytes32 message = _hashHeaderWithChainId(unsignedHeader, CHAIN_ID);

        return ECDSA.recover(message, signature);
    }

    function _hashHeaderWithChainId(BNBHeaderInfo memory header, uint64 chainId) internal pure returns (bytes32) {
        bytes[] memory list = new bytes[](16);

        list[0] = RLPWriter.writeUint(chainId);
        list[1] = RLPWriter.writeBytes(abi.encodePacked(header.parentHash));
        list[2] = RLPWriter.writeBytes(abi.encodePacked(header.sha3Uncles));
        list[3] = RLPWriter.writeAddress(header.miner);
        list[4] = RLPWriter.writeBytes(abi.encodePacked(header.stateRoot));
        list[5] = RLPWriter.writeBytes(abi.encodePacked(header.transactionsRoot));
        list[6] = RLPWriter.writeBytes(abi.encodePacked(header.receiptsRoot));
        list[7] = RLPWriter.writeBytes(header.logsBloom);
        list[8] = RLPWriter.writeUint(header.difficulty);
        list[9] = RLPWriter.writeUint(header.number);
        list[10] = RLPWriter.writeUint(header.gasLimit);
        list[11] = RLPWriter.writeUint(header.gasUsed);
        list[12] = RLPWriter.writeUint(header.timestamp);
        list[13] = RLPWriter.writeBytes(header.extraData);
        list[14] = RLPWriter.writeBytes(abi.encodePacked(header.mixHash));
        list[15] = RLPWriter.writeBytes(abi.encodePacked(header.nonce));

        return keccak256(RLPWriter.writeList(list));
    }

    function finalizedExecutionStateRootAndSlot() external view returns (bytes32 root, uint64 slot) {
        return (stateRoot, number);
    }

    function optimisticExecutionStateRootAndSlot() external view returns (bytes32 root, uint64 slot) {
        return (stateRoot, number);
    }
}
