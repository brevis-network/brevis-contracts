// SPDX-License-Identifier: GPL-3.0-only

pragma solidity >=0.8.0;

interface IBvnSigsVerifier {
    /**
     * @notice Verifies that a message is signed by a quorum among the signers.
     * @param _msg signed message
     * @param _sigs list of signatures sorted by signer addresses in ascending order
     * @param _signers sorted list of current signers
     * @param _powers powers of current signers
     */
    function verifySigs(
        bytes memory _msg,
        bytes[] calldata _sigs,
        address[] calldata _signers,
        uint256[] calldata _powers
    ) external view;

    /**
     * @notice Verifies that a message is signed by a quorum among the signers.
     * @param _msgHash hash of signed message
     * @param _sigs list of signatures sorted by signer addresses in ascending order
     * @param _signers sorted list of current signers
     * @param _powers powers of current signers
     */
    function verifySigs(
        bytes32 _msgHash,
        bytes[] calldata _sigs,
        address[] calldata _signers,
        uint256[] calldata _powers
    ) external view;

    struct SigInfo {
        bytes[] sigs;
        address[] signers;
        uint256[] powers;
    }
}

interface IAvsSigsVerifier {
    struct BN254_G1Point {
        uint256 X;
        uint256 Y;
    }

    // Encoding of field elements is: X[1] * i + X[0]
    struct BN254_G2Point {
        uint256[2] X;
        uint256[2] Y;
    }

    struct NonSignerStakesAndSignature {
        uint32[] nonSignerQuorumBitmapIndices; // is the indices of all nonsigner quorum bitmaps
        BN254_G1Point[] nonSignerPubkeys; // is the G1 pubkeys of all nonsigners
        BN254_G1Point[] quorumApks; // is the aggregate G1 pubkey of each quorum
        BN254_G2Point apkG2; // is the aggregate G2 pubkey of all signers
        BN254_G1Point sigma; // is the aggregate G1 signature of all signers
        uint32[] quorumApkIndices; // is the indices of each quorum aggregate pubkey
        uint32[] totalStakeIndices; // is the indices of each quorums total stake
        uint32[][] nonSignerStakeIndices; // is the indices of each non signers stake within a quorum
    }

    struct SigInfo {
        uint64 blockNum;
        NonSignerStakesAndSignature params;
    }

    function verifySigs(bytes32 msgHash, uint64 blockNum, NonSignerStakesAndSignature calldata params) external view;
}
