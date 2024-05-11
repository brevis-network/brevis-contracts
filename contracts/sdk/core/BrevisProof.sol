// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./BrevisAggProof.sol";
import "../lib/Lib.sol";
import "../../interfaces/ISMT.sol";
import "../../verifiers/interfaces/IZkpVerifier.sol";

contract BrevisProof is BrevisAggProof {
    struct ChainZKVerifier {
        IZkpVerifier contractAppZkVerifier;
        IZkpVerifier circuitAppZkVerifier;
    }
    mapping(uint64 => ChainZKVerifier) public verifierAddresses; // chainid => snark verifier contract address

    mapping(bytes32 => Brevis.ProofData) public proofs; // TODO: store hash of proof data to save gas cost
    mapping(bytes32 => uint256) public vkHashesToBatchSize; // batch tier vk hashes => tier batch size

    event VerifierAddressesUpdated(uint64[] chainIds, ChainZKVerifier[] newAddresses);
    event BatchTierVkHashesUpdated(bytes32[] vkHashes, uint256[] sizes);

    constructor(ISMT _smtContract) BrevisAggProof(_smtContract) {}

    function submitProof(
        uint64 _chainId,
        bytes calldata _proofWithPubInputs,
        bool _withAppProof
    ) external returns (bytes32 _requestId) {
        require(verifyRaw(_chainId, _proofWithPubInputs, _withAppProof), "proof not valid");
        Brevis.ProofData memory data = unpackProofData(_proofWithPubInputs, _withAppProof);
        require(data.vkHash > 0, "vkHash should be larger than 0");
        uint256 batchSize = vkHashesToBatchSize[data.vkHash];
        require(batchSize > 0, "vkHash not valid");

        _requestId = data.commitHash;
        if (_withAppProof) {
            require(smtContract.isSmtRootValid(_chainId, data.smtRoot), "smt root not valid");
            proofs[_requestId].appCommitHash = data.appCommitHash; // save necessary fields only, to save gas
            proofs[_requestId].appVkHash = data.appVkHash;
        } else {
            proofs[_requestId].commitHash = data.commitHash;
        }
    }

    // used by contract app
    function validateRequest(
        bytes32 _requestId,
        uint64 _chainId,
        Brevis.ExtractInfos calldata _extractInfos
    ) external view {
        Brevis.ProofData memory data = proofs[_requestId];
        require(data.commitHash != bytes32(0), "proof not exists");
        require(smtContract.isSmtRootValid(_chainId, _extractInfos.smtRoot), "smt root not valid");

        uint256 itemsLength = _extractInfos.receipts.length + _extractInfos.stores.length + _extractInfos.txs.length;
        require(itemsLength > 0, "empty items");
        uint256 batchSize = vkHashesToBatchSize[data.vkHash];
        require(itemsLength <= batchSize, "item length exceeds batch size");

        bytes memory hashes;

        for (uint256 i = 0; i < _extractInfos.receipts.length; i++) {
            bytes memory fieldInfos;
            for (uint256 j = 0; j < Brevis.NumField; j++) {
                fieldInfos = abi.encodePacked(
                    fieldInfos,
                    _extractInfos.receipts[i].logs[j].logExtraInfo.valueFromTopic,
                    _extractInfos.receipts[i].logs[j].logIndex,
                    _extractInfos.receipts[i].logs[j].logExtraInfo.valueIndex,
                    _extractInfos.receipts[i].logs[j].logExtraInfo.contractAddress,
                    _extractInfos.receipts[i].logs[j].logExtraInfo.logTopic0,
                    _extractInfos.receipts[i].logs[j].value
                );
            }

            hashes = abi.encodePacked(
                hashes,
                keccak256(
                    abi.encodePacked(
                        _extractInfos.smtRoot,
                        _extractInfos.receipts[i].blkNum,
                        _extractInfos.receipts[i].receiptIndex,
                        fieldInfos
                    )
                )
            );
        }

        for (uint256 i = 0; i < _extractInfos.stores.length; i++) {
            hashes = abi.encodePacked(
                hashes,
                keccak256(
                    abi.encodePacked(
                        _extractInfos.smtRoot,
                        _extractInfos.stores[i].blockHash,
                        keccak256(abi.encodePacked(_extractInfos.stores[i].account)),
                        _extractInfos.stores[i].slot,
                        _extractInfos.stores[i].slotValue,
                        _extractInfos.stores[i].blockNumber
                    )
                )
            );
        }
        for (uint256 i = 0; i < _extractInfos.txs.length; i++) {
            hashes = abi.encodePacked(
                hashes,
                keccak256(
                    abi.encodePacked(
                        _extractInfos.smtRoot,
                        _extractInfos.txs[i].leafHash,
                        _extractInfos.txs[i].blockHash,
                        _extractInfos.txs[i].blockNumber,
                        _extractInfos.txs[i].blockTime
                    )
                )
            );
        }

        if (itemsLength < batchSize) {
            bytes32 emptyHash = bytes32(0x0000000000000000000000000000000100000000000000000000000000000001);
            for (uint256 i = itemsLength; i < batchSize; i++) {
                hashes = abi.encodePacked(hashes, emptyHash);
            }
        }
        require(keccak256(hashes) == data.commitHash, "commitHash and info not match");
    }

    function hasProof(bytes32 _requestId) external view returns (bool) {
        return
            proofs[_requestId].commitHash != bytes32(0) ||
            proofs[_requestId].appCommitHash != bytes32(0) ||
            inAgg(_requestId);
    }

    function getProofData(bytes32 _requestId) external view returns (Brevis.ProofData memory) {
        return proofs[_requestId];
    }

    function getProofAppData(bytes32 _requestId) external view returns (bytes32, bytes32) {
        return (proofs[_requestId].appCommitHash, proofs[_requestId].appVkHash);
    }

    function verifyRaw(
        uint64 _chainId,
        bytes calldata _proofWithPubInputs,
        bool _withAppProof
    ) private view returns (bool) {
        IZkpVerifier verifier;
        if (!_withAppProof) {
            verifier = verifierAddresses[_chainId].contractAppZkVerifier;
        } else {
            verifier = verifierAddresses[_chainId].circuitAppZkVerifier;
        }
        require(address(verifier) != address(0), "chain verifier not set");
        return verifier.verifyRaw(_proofWithPubInputs);
    }

    function unpackProofData(
        bytes calldata _proofWithPubInputs,
        bool _withAppProof
    ) internal pure returns (Brevis.ProofData memory data) {
        if (_withAppProof) {
            data.commitHash = bytes32(_proofWithPubInputs[PUBLIC_BYTES_START_IDX:PUBLIC_BYTES_START_IDX + 32]);
            data.smtRoot = bytes32(_proofWithPubInputs[PUBLIC_BYTES_START_IDX + 32:PUBLIC_BYTES_START_IDX + 2 * 32]);
            data.vkHash = bytes32(_proofWithPubInputs[PUBLIC_BYTES_START_IDX + 2 * 32:PUBLIC_BYTES_START_IDX + 3 * 32]);
            data.appCommitHash = bytes32(
                _proofWithPubInputs[PUBLIC_BYTES_START_IDX + 3 * 32:PUBLIC_BYTES_START_IDX + 4 * 32]
            );
            data.appVkHash = bytes32(
                _proofWithPubInputs[PUBLIC_BYTES_START_IDX + 4 * 32:PUBLIC_BYTES_START_IDX + 5 * 32]
            );
        } else {
            data.commitHash = bytes32(_proofWithPubInputs[PUBLIC_BYTES_START_IDX:PUBLIC_BYTES_START_IDX + 32]);
            // data length field in between no need to be unpacked
            data.vkHash = bytes32(_proofWithPubInputs[PUBLIC_BYTES_START_IDX + 2 * 32:PUBLIC_BYTES_START_IDX + 3 * 32]);
        }
    }

    function updateVerifierAddress(
        uint64[] calldata _chainIds,
        ChainZKVerifier[] calldata _verifierAddresses
    ) public onlyOwner {
        require(_chainIds.length == _verifierAddresses.length, "length not match");
        for (uint256 i = 0; i < _chainIds.length; i++) {
            verifierAddresses[_chainIds[i]] = _verifierAddresses[i];
        }
        emit VerifierAddressesUpdated(_chainIds, _verifierAddresses);
    }

    function setBatchTierVkHashes(bytes32[] calldata _vkHashes, uint256[] calldata _sizes) public onlyOwner {
        require(_vkHashes.length == _sizes.length, "length not match");
        for (uint256 i = 0; i < _vkHashes.length; i++) {
            vkHashesToBatchSize[_vkHashes[i]] = _sizes[i];
        }

        emit BatchTierVkHashesUpdated(_vkHashes, _sizes);
    }

    address public brevisRequest;
    event BrevisRequestUpdated(address brevisRequest);
    modifier onlyBrevisRequest() {
        require(brevisRequest == msg.sender, "not brevisRequest");
        _;
    }

    function updateBrevisRequest(address _brevisRequest) public onlyOwner {
        brevisRequest = _brevisRequest;
        emit BrevisRequestUpdated(_brevisRequest);
    }

    function submitOpResult(bytes32 _requestId) external onlyBrevisRequest {
        proofs[_requestId].commitHash = _requestId;
    }
}