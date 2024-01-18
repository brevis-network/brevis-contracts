// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

import "./Initializable.sol";
import "./System.sol";
import "./lib/Tendermint.sol";
import "./interfaces/ITendermintLightClient.sol";
import {GoogleProtobufAny as Any} from "./lib/proto/GoogleProtobufAny.sol";
import {LightHeader, ValidatorSet, ConsensusState, TmHeader} from "./lib/proto/TendermintLight.sol";

contract TendermintLightClient is Initializable, ITendermintLightClient {
    using Bytes for bytes;
    using Bytes for bytes32;
    using TendermintHelper for TmHeader.Data;
    using TendermintHelper for ConsensusState.Data;
    using TendermintHelper for ValidatorSet.Data;

    struct ProtoTypes {
        bytes32 consensusState;
        bytes32 tmHeader;
    }

    ProtoTypes private _pts;
    mapping(uint64 => ConsensusState.Data) public consensusStates;
    mapping(uint64 => bool) public synced;
    uint64 public initialHeight;
    uint64 public latestHeight;
    System private system;
    address ed25519Verifier;

    event ConsensusStateInit(uint64 initialHeight, bytes32 appHash);
    event ConsensusStateSynced(uint64 height, bytes32 appHash);

    constructor(address _ed25519Verifier) {
        ed25519Verifier = _ed25519Verifier;
    }

    function init(address _system, bytes memory _initHeader) external onlyUninitialized {
        _pts = ProtoTypes({
            consensusState: keccak256(abi.encodePacked("/tendermint.types.ConsensusState")),
            tmHeader: keccak256(abi.encodePacked("/tendermint.types.TmHeader"))
        });

        system = System(_system);

        (TmHeader.Data memory tmHeader, bool ok) = unmarshalTmHeader(_initHeader);
        require(ok, "LC: light block is invalid");

        uint64 height = uint64(tmHeader.signed_header.header.height);
        ConsensusState.Data memory cs = tmHeader.toConsensusState();
        consensusStates[height] = cs;

        initialHeight = height;
        latestHeight = height;

        emit ConsensusStateInit(initialHeight, bytes32(cs.root.hash));

        _initialized = true;
    }

    function syncTendermintHeader(
        bytes calldata header,
        uint256[2] memory proofA,
        uint256[2][2] memory proofB,
        uint256[2] memory proofC,
        uint256[2] memory proofCommit,
        uint256 proofCommitPub
    ) external returns (bool) {
        require(msg.sender == system.relayer(), "not relayer");

        (TmHeader.Data memory tmHeader, bool ok) = unmarshalTmHeader(header);
        require(ok, "LC: light block is invalid");

        uint64 height = uint64(tmHeader.signed_header.header.height);
        require(!synced[height], "can't sync duplicated header");
        // assert header height is newer than consensus state
        require(height > latestHeight, "LC: header height not newer than consensus state height");

        checkValidity(consensusStates[latestHeight], tmHeader, proofA, proofB, proofC, proofCommit, proofCommitPub);

        synced[height] = true;

        // Store new cs
        ConsensusState.Data memory cs = tmHeader.toConsensusState();
        consensusStates[height] = cs;

        emit ConsensusStateSynced(height, bytes32(cs.root.hash));

        return true;
    }

    // checkValidity checks if the Tendermint header is valid.
    function checkValidity(
        ConsensusState.Data memory trustedConsensusState,
        TmHeader.Data memory tmHeader,
        uint256[2] memory proofA,
        uint256[2][2] memory proofB,
        uint256[2] memory proofC,
        uint256[2] memory proofCommit,
        uint256 proofCommitPub
    ) private view {
        LightHeader.Data memory lc;
        lc.chain_id = tmHeader.signed_header.header.chain_id;
        lc.height = int64(latestHeight);
        lc.next_validators_hash = trustedConsensusState.next_validators_hash;

        SignedHeader.Data memory trustedHeader;
        trustedHeader.header = lc;

        SignedHeader.Data memory untrustedHeader = tmHeader.signed_header;
        ValidatorSet.Data memory untrustedVals = tmHeader.validator_set;

        bool ok = Tendermint.verify(
            trustedHeader,
            untrustedHeader,
            untrustedVals,
            ed25519Verifier,
            proofA,
            proofB,
            proofC,
            proofCommit,
            proofCommitPub
        );

        require(ok, "LC: failed to verify header");
    }

    function isHeaderSynced(uint64 height) external view override returns (bool) {
        return synced[height] || height == initialHeight;
    }

    function getAppHash(uint64 height) external view override returns (bytes32) {
        return bytes32(consensusStates[height].root.hash);
    }

    function unmarshalTmHeader(bytes memory bz) internal view returns (TmHeader.Data memory header, bool ok) {
        Any.Data memory anyHeader = Any.decode(bz);
        if (keccak256(abi.encodePacked(anyHeader.type_url)) != _pts.tmHeader) {
            return (header, false);
        }
        return (TmHeader.decode(anyHeader.value), true);
    }
}
