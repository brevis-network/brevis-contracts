// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "./BrevisAggProof.sol";
import "../lib/Lib.sol";
import "../../interfaces/ISMT.sol";
import "../../verifiers/interfaces/IZkpVerifier.sol";

contract BrevisProof is BrevisAggProof {
    mapping(uint64 => IZkpVerifier) public verifierAddresses; // chainid => snark verifier contract address
    mapping(bytes32 => bytes32) public proofs; // proofId => keccak256(abi.encodePacked(appCommitHash, appVkHash));
    event VerifierAddressesUpdated(uint64[] chainIds, IZkpVerifier[] newAddresses);

    constructor(ISMT _smtContract) BrevisAggProof(_smtContract) {}

    // To support upgradable deployment.
    // Can only be called once by Proxy via delegateCall, as initOwner will require _owner is 0.
    function init(ISMT _smtContract) external {
        initOwner();
        smtContract = _smtContract;
    }

    /*********************************
     * External and Public Functions *
     *********************************/

    function submitProof(
        uint64 _chainId,
        bytes calldata _proofWithPubInputs
    ) external onlyActiveProver returns (bytes32 proofId, bytes32 appCommitHash, bytes32 appVkHash) {
        require(verifyRaw(_chainId, _proofWithPubInputs), "proof not valid");
        Brevis.ProofData memory data = unpackProofData(_proofWithPubInputs);

        proofId = data.commitHash;
        appCommitHash = data.appCommitHash;
        appVkHash = data.appVkHash;
        require(smtContract.isSmtRootValid(_chainId, data.smtRoot), "smt root not valid");
        proofs[proofId] = keccak256(abi.encodePacked(appCommitHash, appVkHash));
    }

    function validateProofAppData(
        bytes32 _proofId,
        bytes32 _appCommitHash,
        bytes32 _appVkHash
    ) external view returns (bool) {
        require(proofs[_proofId] == keccak256(abi.encodePacked(_appCommitHash, _appVkHash)), "invalid data");
        return true;
    }

    // -------- owner functions --------

    function updateVerifierAddress(
        uint64[] calldata _chainIds,
        IZkpVerifier[] calldata _verifierAddresses
    ) public onlyOwner {
        require(_chainIds.length == _verifierAddresses.length, "length not match");
        for (uint256 i = 0; i < _chainIds.length; i++) {
            verifierAddresses[_chainIds[i]] = _verifierAddresses[i];
        }
        emit VerifierAddressesUpdated(_chainIds, _verifierAddresses);
    }

    /**********************************
     * Internal and Private Functions *
     **********************************/

    function unpackProofData(bytes calldata _proofWithPubInputs) internal pure returns (Brevis.ProofData memory data) {
        data.commitHash = bytes32(_proofWithPubInputs[PUBLIC_BYTES_START_IDX:PUBLIC_BYTES_START_IDX + 32]);
        data.smtRoot = bytes32(_proofWithPubInputs[PUBLIC_BYTES_START_IDX + 32:PUBLIC_BYTES_START_IDX + 2 * 32]);
        //data.vkHash = bytes32(_proofWithPubInputs[PUBLIC_BYTES_START_IDX + 2 * 32:PUBLIC_BYTES_START_IDX + 3 * 32]);
        data.appCommitHash = bytes32(
            _proofWithPubInputs[PUBLIC_BYTES_START_IDX + 3 * 32:PUBLIC_BYTES_START_IDX + 4 * 32]
        );
        data.appVkHash = bytes32(_proofWithPubInputs[PUBLIC_BYTES_START_IDX + 4 * 32:PUBLIC_BYTES_START_IDX + 5 * 32]);
    }

    function verifyRaw(uint64 _chainId, bytes calldata _proofWithPubInputs) private view returns (bool) {
        IZkpVerifier verifier = verifierAddresses[_chainId];
        require(address(verifier) != address(0), "chain verifier not set");
        return verifier.verifyRaw(_proofWithPubInputs);
    }
}
