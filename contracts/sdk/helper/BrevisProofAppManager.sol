// SPDX-License-Identifier: GPL-3.0-only

pragma solidity ^0.8.18;

import "../../safeguard/Whitelist.sol";
import "../interface/IBrevisProof.sol";

contract BrevisProofAppManager is Whitelist {
    IBrevisProof public brevisProof;

    event BrevisProofUpdated(address from, address to);

    constructor( IBrevisProof _brevisProof) {
        brevisProof = _brevisProof;
    }

    function submitProof(
        uint64 _chainId,
        bytes calldata _proofWithPubInputs
    ) external onlyWhitelisted returns (bytes32 requestId, bytes32 appCommitHash, bytes32 appVkHash) {
        return brevisProof.submitProof(_chainId, _proofWithPubInputs);
    }

    function validateProofAppData(
        bytes32 _requestId,
        bytes32 _appCommitHash,
        bytes32 _appVkHash
    ) external view returns (bool) {
        return brevisProof.validateProofAppData(_requestId, _appCommitHash, _appVkHash);
    }

    function submitAggProof(
        uint64 _chainId,
        bytes32[] calldata _requestIds,
        bytes calldata _proofWithPubInputs
    ) external onlyWhitelisted {
        return brevisProof.submitAggProof(_chainId, _requestIds, _proofWithPubInputs);
    }

    function validateAggProofData(uint64 _chainId, Brevis.ProofData[] calldata _proofDataArray) external view {
        return brevisProof.validateAggProofData(_chainId, _proofDataArray);
    }

    function setBrevisProof(address _brevisProof) external onlyOwner {
        address oldAddr = address(brevisProof);
        brevisProof = IBrevisProof(_brevisProof);
        emit BrevisProofUpdated(oldAddr, _brevisProof);
    }
}