// SPDX-License-Identifier: MIT
pragma solidity ^0.8.2;

import "./Bytes.sol";
import {TENDERMINTLIGHT_PROTO_GLOBAL_ENUMS, SignedHeader, BlockID, Timestamp, ValidatorSet, Duration, Fraction, Commit, Validator, CommitSig, CanonicalVote, Vote} from "./proto/TendermintLight.sol";
import "./proto/TendermintHelper.sol";
import "./proto/Encoder.sol";
import "../Ed25519Verifier.sol";
import {BytesLib} from "solidity-bytes-utils/contracts/BytesLib.sol";

library Tendermint {
    using Bytes for bytes;
    using Bytes for bytes32;
    using TendermintHelper for ValidatorSet.Data;
    using TendermintHelper for SignedHeader.Data;
    using TendermintHelper for Timestamp.Data;
    using TendermintHelper for BlockID.Data;
    using TendermintHelper for Commit.Data;
    using TendermintHelper for Vote.Data;

    // TODO: Change visibility to public for deployment. For some reason have to use internal for abigen.
    function verify(
        SignedHeader.Data memory trustedHeader,
        SignedHeader.Data memory untrustedHeader,
        ValidatorSet.Data memory untrustedVals,
        address verifier,
        uint256[2] memory proofA,
        uint256[2][2] memory proofB,
        uint256[2] memory proofC,
        uint256[2] memory proofCommit,
        uint256 proofCommitPub
    ) internal view returns (bool) {
        verifyNewHeaderAndVals(untrustedHeader, untrustedVals, trustedHeader);

        // Check the validator hashes are the same
        require(
            untrustedHeader.header.validators_hash.toBytes32() == trustedHeader.header.next_validators_hash.toBytes32(),
            "expected old header next validators to match those from new header"
        );

        // Ensure that +2/3 of new validators signed correctly.
        bool ok = verifyCommitLight(
            untrustedVals,
            trustedHeader.header.chain_id,
            untrustedHeader.commit.block_id,
            untrustedHeader.header.height,
            untrustedHeader.commit,
            verifier,
            proofA,
            proofB,
            proofC,
            proofCommit,
            proofCommitPub
        );

        return ok;
    }

    function verifyNewHeaderAndVals(
        SignedHeader.Data memory untrustedHeader,
        ValidatorSet.Data memory untrustedVals,
        SignedHeader.Data memory trustedHeader
    ) internal pure {
        // SignedHeader validate basic
        require(
            keccak256(abi.encodePacked(untrustedHeader.header.chain_id)) ==
                keccak256(abi.encodePacked(trustedHeader.header.chain_id)),
            "header belongs to another chain"
        );
        require(untrustedHeader.commit.height == untrustedHeader.header.height, "header and commit height mismatch");

        bytes32 untrustedHeaderBlockHash = untrustedHeader.hash();
        // TODO: Fix block hash
        // require(
        //     untrustedHeaderBlockHash == untrustedHeader.commit.block_id.hash.toBytes32(),
        //     "commit signs signs block failed"
        // );

        require(
            untrustedHeader.header.height > trustedHeader.header.height,
            "expected new header height to be greater than one of old header"
        );

        // Skip time verification for now

        bytes32 validatorsHash = untrustedVals.hash();
        // TODO: Fix validators hash
        // require(
        //     untrustedHeader.header.validators_hash.toBytes32() == validatorsHash,
        //     "expected new header validators to match those that were supplied at height XX"
        // );
    }

    // VerifyCommitLight
    // Proof of concept header verification with batch signature SNARK proof
    function verifyCommitLight(
        ValidatorSet.Data memory vals,
        string memory chainID,
        BlockID.Data memory blockID,
        int64 height,
        Commit.Data memory commit,
        address verifier,
        uint256[2] memory proofA,
        uint256[2][2] memory proofB,
        uint256[2] memory proofC,
        uint256[2] memory proofCommit,
        uint256 proofCommitPub
    ) internal view returns (bool) {
        require(vals.validators.length == commit.signatures.length, "invalid commit signatures");
        require(commit.signatures.length > 8, "insufficient signatures");

        require(height == commit.height, "invalid commit height");

        require(commit.block_id.isEqual(blockID), "invalid commit -- wrong block ID");

        bytes[8] memory pubkeys;
        bytes[8] memory messages;
        uint256 sigCount;
        for (uint256 i = 0; i < commit.signatures.length; i++) {
            // no need to verify absent or nil votes.
            if (
                commit.signatures[i].block_id_flag !=
                TENDERMINTLIGHT_PROTO_GLOBAL_ENUMS.BlockIDFlag.BLOCK_ID_FLAG_COMMIT
            ) {
                continue;
            }

            pubkeys[sigCount] = vals.validators[i].pub_key.ed25519;
            messages[sigCount] = Encoder.encodeDelim(voteSignBytes(commit, chainID, i));

            sigCount++;
            if (sigCount == 8) {
                break;
            }
        }

        uint256[57] memory input = prepareInput(pubkeys, messages, proofCommitPub);
        return Ed25519Verifier(verifier).verifyProof(proofA, proofB, proofC, proofCommit, input);
    }

    function prepareInput(
        bytes[8] memory pubkeys,
        bytes[8] memory messages,
        uint256 proofCommitPub
    ) private pure returns (uint256[57] memory input) {
        for (uint256 i = 0; i < 8; i++) {
            bytes memory messagePart0 = BytesLib.slice(messages[i], 0, 25);
            bytes memory messagePart1 = BytesLib.slice(messages[i], 25, 25);
            bytes memory messagePart2 = BytesLib.slice(messages[i], 50, 25);
            bytes memory messagePart3 = BytesLib.slice(messages[i], 75, 25);
            bytes memory messagePart4 = BytesLib.slice(messages[i], 100, 22);
            input[5 * i] = uint256(uint200(bytes25(messagePart0)));
            input[5 * i + 1] = uint256(uint200(bytes25(messagePart1)));
            input[5 * i + 2] = uint256(uint200(bytes25(messagePart2)));
            input[5 * i + 3] = uint256(uint200(bytes25(messagePart3)));
            input[5 * i + 4] = uint256(uint176(bytes22(messagePart4)));
            bytes memory pubkeyHigh = BytesLib.slice(pubkeys[i], 0, 16);
            bytes memory pubkeyLow = BytesLib.slice(pubkeys[i], 16, 16);
            input[2 * i + 40] = uint256(uint128(bytes16(pubkeyHigh)));
            input[2 * i + 1 + 40] = uint256(uint128(bytes16(pubkeyLow)));
        }
        input[56] = proofCommitPub;
        return input;
    }

    function voteSignBytes(
        Commit.Data memory commit,
        string memory chainID,
        uint256 idx
    ) internal pure returns (bytes memory) {
        Vote.Data memory vote;
        vote = commit.toVote(idx);

        return CanonicalVote.encode(vote.toCanonicalVote(chainID));
    }

    function voteSignBytesDelim(
        Commit.Data memory commit,
        string memory chainID,
        uint256 idx
    ) internal pure returns (bytes memory) {
        return Encoder.encodeDelim(voteSignBytes(commit, chainID, idx));
    }
}
