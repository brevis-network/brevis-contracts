// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/access/Ownable.sol";

import "../../../interfaces/ISMT.sol";
import "../../../verifiers/interfaces/IZkpVerifier.sol";
import "./IUniswapSumVolume.sol";

contract UniswapSumVolume is Ownable, IUniswapSumVolume {
    // retrieved from proofData, to align the fields with circuit...
    struct ProofData {
        address recipient;
        uint256 volume;
        bytes32 smtRoot;
        uint256 length;
        uint256 cPub;
        uint256 vkHash;
    }

    uint256[6] public batchTierVkHashes;
    uint32 constant PUBLIC_BYTES_START_IDX = 10 * 32; // the first 10 32bytes are groth16 proof (A/B/C/Commitment)

    mapping(address => uint256) public volumes;
    mapping(uint64 => address) public verifierAddresses; // chainid => snark verifier contract address

    ISMT public smtContract;

    event UpdateVerifierAddress(uint64 chainId, address newAddress);
    event UpdateSmtContract(ISMT smtContract);
    event SumVolume(address user, uint64 fromChain, uint256 volume);

    constructor(ISMT _smtContract) {
        smtContract = _smtContract;
    }

    function submitUniswapSumVolumeProof(uint64 _chainId, bytes calldata _proof) external {
        require(verifyRaw(_chainId, _proof), "proof not valid");

        ProofData memory data = getProofData(_proof);
        require(data.volume > 0, "volume should be larger than 0");
        require(data.vkHash > 0, "vkHash should be larger than 0");
        require(isIn(data.vkHash), "vkHash is not valid");
        require(volumes[data.recipient] == 0, "already proved for this user");
        require(smtContract.isSmtRootValid(_chainId, data.smtRoot), "smt root not valid");

        volumes[data.recipient] = data.volume;
        emit SumVolume(data.recipient, _chainId, data.volume);
    }

    function isIn(uint256 vkHash) internal view returns (bool exists) {
        exists = false;
        for (uint256 i = 0; i < 6; i++) {
            if (vkHash == batchTierVkHashes[i]) {
                exists = true;
                break;
            }
        }
    }

    function verifyRaw(uint64 chainId, bytes calldata proofData) private view returns (bool) {
        require(verifierAddresses[chainId] != address(0), "chain verifier not set");
        return (IZkpVerifier)(verifierAddresses[chainId]).verifyRaw(proofData);
    }

    function getProofData(bytes calldata _proofData) internal pure returns (ProofData memory data) {
        data.cPub = uint256(bytes32(_proofData[PUBLIC_BYTES_START_IDX:PUBLIC_BYTES_START_IDX + 32]));
        data.recipient = address(bytes20(_proofData[PUBLIC_BYTES_START_IDX + 32 + 12:PUBLIC_BYTES_START_IDX + 2 * 32]));
        data.volume = uint256(bytes32(_proofData[PUBLIC_BYTES_START_IDX + 2 * 32:PUBLIC_BYTES_START_IDX + 3 * 32]));
        data.smtRoot = bytes32(
            (uint256(bytes32(_proofData[PUBLIC_BYTES_START_IDX + 3 * 32:PUBLIC_BYTES_START_IDX + 4 * 32])) << 128) |
                uint128(bytes16(_proofData[PUBLIC_BYTES_START_IDX + 4 * 32 + 16:PUBLIC_BYTES_START_IDX + 5 * 32]))
        );
        data.length = uint256(bytes32(_proofData[PUBLIC_BYTES_START_IDX + 5 * 32:PUBLIC_BYTES_START_IDX + 6 * 32]));
        data.vkHash = uint256(bytes32(_proofData[PUBLIC_BYTES_START_IDX + 6 * 32:PUBLIC_BYTES_START_IDX + 7 * 32]));
    }

    function updateSmtContract(ISMT _smtContract) external onlyOwner {
        smtContract = _smtContract;
        emit UpdateSmtContract(smtContract);
    }

    function updateVerifierAddress(uint64 _chainId, address _verifierAddress) external onlyOwner {
        verifierAddresses[_chainId] = _verifierAddress;
        emit UpdateVerifierAddress(_chainId, _verifierAddress);
    }

    function getAttestedSwapSumVolume(address swapper) external view returns (uint256) {
        return volumes[swapper];
    }

    function setBatchTierVkHashes(bytes calldata hashes) external onlyOwner {
        uint256 len = hashes.length / 32;
        require(len <= 6, "exceeds max tiers");
        for (uint256 i = 0; i < len; i++) {
            batchTierVkHashes[i] = uint256(bytes32(hashes[i * 32:(i + 1) * 32]));
        }
    }
}
