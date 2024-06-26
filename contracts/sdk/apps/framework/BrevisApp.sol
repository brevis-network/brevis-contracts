// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/access/Ownable.sol";
import "../../interface/IBrevisProof.sol";
import "../../interface/IBrevisRequest.sol";
import "../../lib/Lib.sol";

abstract contract BrevisApp is Ownable {
    IBrevisProof public brevisProof;
    uint256 public opChallengeWindow;

    modifier onlyBrevisRequest() {
        require(msg.sender == address(brevisProof.getRequestContract()), "invalid caller");
        _;
    }

    constructor(IBrevisProof _brevisProof) {
        brevisProof = _brevisProof;
        opChallengeWindow = 2 ** 256 - 1; // disable usage of op result by default
    }

    function brevisCallback(bytes32 _appVkHash, bytes calldata _appCircuitOutput) external onlyBrevisRequest {
        handleProofResult(_appVkHash, _appCircuitOutput);
    }

    function brevisBatchCallback(
        bytes32[] calldata _appVkHashes,
        bytes[] calldata _appCircuitOutputs
    ) external onlyBrevisRequest {
        for (uint i = 0; i < _appVkHashes.length; i++) {
            handleProofResult(_appVkHashes[i], _appCircuitOutputs[i]);
        }
    }

    // apply proved request
    function applyBrevisProof(
        bytes32 _proofId,
        bytes32 _appVkHash,
        bytes32 _appCommitHash,
        bytes calldata _appCircuitOutput
    ) external {
        brevisProof.validateProofAppData(_proofId, _appCommitHash, _appVkHash);
        require(_appCommitHash == keccak256(_appCircuitOutput), "invalid circuit output");
        handleProofResult(_appVkHash, _appCircuitOutput);
    }

    // apply multiple requests fulfilled through AggProof
    function applyBrevisAggProof(
        uint64 _chainId,
        Brevis.ProofData[] calldata _proofDataArray,
        bytes[] calldata _appCircuitOutputs
    ) external {
        require(_proofDataArray.length == _appCircuitOutputs.length, "length not match");
        brevisProof.validateAggProofData(_chainId, _proofDataArray);
        for (uint i = 0; i < _proofDataArray.length; i++) {
            require(_proofDataArray[i].appCommitHash == keccak256(_appCircuitOutputs[i]), "invalid circuit output");
            handleProofResult(_proofDataArray[i].appVkHash, _appCircuitOutputs[i]);
        }
    }

    // apply single request fulfilled through AggProof
    function applyBrevisAggProof(
        uint64 _chainId,
        Brevis.ProofData calldata _proofData,
        bytes32 _merkleRoot,
        bytes32[] calldata _merkleProof,
        uint8 _nodeIndex,
        bytes calldata _appCircuitOutput
    ) external {
        brevisProof.validateAggProofData(_chainId, _proofData, _merkleRoot, _merkleProof, _nodeIndex);
        require(_proofData.appCommitHash == keccak256(_appCircuitOutput), "invalid circuit output");
        handleProofResult(_proofData.appVkHash, _appCircuitOutput);
    }

    function applyBrevisOpResult(
        bytes32 _proofId,
        uint64 _nonce,
        bytes32 _appVkHash,
        bytes32 _appCommitHash,
        bytes calldata _appCircuitOutput
    ) public {
        require(
            IBrevisRequest(brevisProof.getRequestContract()).validateOpAppData(
                _proofId,
                _nonce,
                _appCommitHash,
                _appVkHash,
                opChallengeWindow
            ),
            "data not ready to use"
        );
        require(_appCommitHash == keccak256(_appCircuitOutput), "invalid circuit output");
        handleOpProofResult(_appVkHash, _appCircuitOutput);
    }

    function applyBrevisOpResults(
        bytes32[] calldata _proofIds,
        uint64[] calldata _nonces,
        bytes32[] calldata _appVkHashes,
        bytes32[] calldata _appCommitHashes,
        bytes[] calldata _appCircuitOutputs
    ) external {
        uint256 len = _proofIds.length;
        require(
            len == _appVkHashes.length && len == _appCommitHashes.length && len == _appCircuitOutputs.length,
            "length mismatch"
        );
        for (uint256 i = 0; i < _proofIds.length; i++) {
            applyBrevisOpResult(_proofIds[i], _nonces[i], _appVkHashes[i], _appCommitHashes[i], _appCircuitOutputs[i]);
        }
    }

    function setOpChallengeWindow(uint256 _challangeWindow) external onlyOwner {
        opChallengeWindow = _challangeWindow;
    }

    function handleProofResult(bytes32 _vkHash, bytes calldata _appCircuitOutput) internal virtual {
        // to be overrided by custom app
    }

    function handleOpProofResult(bytes32 _vkHash, bytes calldata _appCircuitOutput) internal virtual {
        // to be overrided by custom app
    }
}
