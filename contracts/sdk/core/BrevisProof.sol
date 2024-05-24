// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./BrevisAggProof.sol";
import "../lib/Lib.sol";
import "../../interfaces/ISMT.sol";
import "../../verifiers/interfaces/IZkpVerifier.sol";

contract BrevisProof is BrevisAggProof {
    mapping(uint64 => IZkpVerifier) public verifierAddresses; // chainid => snark verifier contract address
    mapping(bytes32 => Brevis.ProofData) public proofs; // TODO: store hash of proof data to save gas cost

    event VerifierAddressesUpdated(uint64[] chainIds, IZkpVerifier[] newAddresses);

    constructor(ISMT _smtContract) BrevisAggProof(_smtContract) {}

    function submitProof(uint64 _chainId, bytes calldata _proofWithPubInputs) external returns (bytes32 _requestId) {
        require(verifyRaw(_chainId, _proofWithPubInputs), "proof not valid");
        Brevis.ProofData memory data = unpackProofData(_proofWithPubInputs);

        _requestId = data.commitHash;
        require(smtContract.isSmtRootValid(_chainId, data.smtRoot), "smt root not valid");
        proofs[_requestId].appCommitHash = data.appCommitHash; // save necessary fields only, to save gas
        proofs[_requestId].appVkHash = data.appVkHash;
    }

    function hasProof(bytes32 _requestId) external view returns (bool) {
        return
            proofs[_requestId].commitHash != bytes32(0) ||
            proofs[_requestId].appCommitHash != bytes32(0) ||
            inAgg(_requestId);
    }

    function getProofAppData(bytes32 _requestId) external view returns (bytes32, bytes32) {
        return (proofs[_requestId].appCommitHash, proofs[_requestId].appVkHash);
    }

    function verifyRaw(uint64 _chainId, bytes calldata _proofWithPubInputs) private view returns (bool) {
        IZkpVerifier verifier = verifierAddresses[_chainId];
        require(address(verifier) != address(0), "chain verifier not set");
        return verifier.verifyRaw(_proofWithPubInputs);
    }

    function unpackProofData(bytes calldata _proofWithPubInputs) internal pure returns (Brevis.ProofData memory data) {
        data.commitHash = bytes32(_proofWithPubInputs[PUBLIC_BYTES_START_IDX:PUBLIC_BYTES_START_IDX + 32]);
        data.smtRoot = bytes32(_proofWithPubInputs[PUBLIC_BYTES_START_IDX + 32:PUBLIC_BYTES_START_IDX + 2 * 32]);
        data.vkHash = bytes32(_proofWithPubInputs[PUBLIC_BYTES_START_IDX + 2 * 32:PUBLIC_BYTES_START_IDX + 3 * 32]);
        data.appCommitHash = bytes32(
            _proofWithPubInputs[PUBLIC_BYTES_START_IDX + 3 * 32:PUBLIC_BYTES_START_IDX + 4 * 32]
        );
        data.appVkHash = bytes32(_proofWithPubInputs[PUBLIC_BYTES_START_IDX + 4 * 32:PUBLIC_BYTES_START_IDX + 5 * 32]);
    }

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
