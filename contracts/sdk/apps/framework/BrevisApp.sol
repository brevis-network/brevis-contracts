// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/access/Ownable.sol";
import "../../interface/IBrevisProof.sol";
import "../../interface/IBrevisRequest.sol";
import "../../lib/Lib.sol";

abstract contract BrevisApp is Ownable {
    IBrevisProof public brevisProof;
    IBrevisRequest public brevisRequest;
    uint256 public OpChallengeWindow;

    modifier onlyBrevisRequest() {
        require(msg.sender == address(brevisRequest), "invalid caller");
        _;
    }

    constructor(IBrevisProof _brevisProof) {
        brevisProof = _brevisProof;
        brevisRequest = IBrevisRequest(brevisProof.getRequestContract());
    }

    function brevisCallback(
        bytes32 _requestId,
        bytes32 _appVkHash,
        bytes calldata _appCircuitOutput
    ) external onlyBrevisRequest {
        handleProofResult(_requestId, _appVkHash, _appCircuitOutput);
    }

    function brevisBatchCallback(
        bytes32[] calldata _requestIds,
        bytes32[] calldata _appVkHashs,
        bytes[] calldata _appCircuitOutputs
    ) external onlyBrevisRequest {
        for (uint i = 0; i < _requestIds.length; i++) {
            handleProofResult(_requestIds[i], _appVkHashs[i], _appCircuitOutputs[i]);
        }
    }

    // apply proved request
    function applyBrevisProof(bytes32 _requestId, bytes calldata _appCircuitOutput) external {
        (bytes32 appCommitHash, bytes32 appVkHash) = IBrevisProof(brevisProof).getProofAppData(_requestId);
        require(appCommitHash == keccak256(_appCircuitOutput), "failed to open output commitment");
        handleProofResult(_requestId, appVkHash, _appCircuitOutput);
    }

    // apply multiple requests fulfilled through AggProof
    function applyBrevisAggProof(
        uint64 _chainId,
        Brevis.ProofData[] calldata _proofDataArray,
        bytes[] calldata _appCircuitOutputs
    ) external {
        require(_proofDataArray.length == _appCircuitOutputs.length, "length not match");
        IBrevisProof(brevisProof).mustValidateRequests(_chainId, _proofDataArray);
        for (uint i = 0; i < _proofDataArray.length; i++) {
            require(
                _proofDataArray[i].appCommitHash == keccak256(_appCircuitOutputs[i]),
                "failed to open output commitment"
            );
            handleProofResult(_proofDataArray[i].commitHash, _proofDataArray[i].appVkHash, _appCircuitOutputs[i]);
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
        IBrevisProof(brevisProof).mustValidateRequest(_chainId, _proofData, _merkleRoot, _merkleProof, _nodeIndex);
        require(_proofData.appCommitHash == keccak256(_appCircuitOutput), "failed to open output commitment");
        handleProofResult(_proofData.commitHash, _proofData.appVkHash, _appCircuitOutput);
    }

    function applyBrevisOpResult(
        bytes32 _requestId,
        bytes32 _appVkHash,
        bytes32 _appCommitHash,
        bytes calldata _appCircuitOutput
    ) public {
        IBrevisRequest.RequestStatus status = brevisRequest.queryRequestStatus(_requestId, OpChallengeWindow);
        require(
            status == IBrevisRequest.RequestStatus.OpAttested || status == IBrevisRequest.RequestStatus.ZkAttested,
            "invalid status"
        );
        require(brevisRequest.validateRequestOpData(_requestId, _appVkHash, _appCommitHash), "invalid result");
        require(_appCommitHash == keccak256(_appCircuitOutput), "failed to open output commitment");
        handleProofResult(_requestId, _appVkHash, _appCircuitOutput);
    }

    function applyBrevisOpResults(
        bytes32[] calldata _requestIds,
        bytes32[] calldata _appVkHashes,
        bytes32[] calldata _appCommitHashes,
        bytes[] calldata _appCircuitOutputs
    ) external {
        require(
            _requestIds.length == _appVkHashes.length &&
                _requestIds.length == _appCommitHashes.length &&
                _requestIds.length == _appCircuitOutputs.length,
            "length mismatch"
        );
        for (uint256 i = 0; i < _requestIds.length; i++) {
            applyBrevisOpResult(_requestIds[i], _appVkHashes[i], _appCommitHashes[i], _appCircuitOutputs[i]);
        }
    }

    function setOpChallengeWindow(uint256 _challangeWindow) external onlyOwner {
        OpChallengeWindow = _challangeWindow;
    }

    function handleProofResult(bytes32 _requestId, bytes32 _vkHash, bytes calldata _appCircuitOutput) internal virtual {
        // to be overrided by custom app
    }
}
