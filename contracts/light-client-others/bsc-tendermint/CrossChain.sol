// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

import "./interfaces/IApplication.sol";
import "./interfaces/ICrossChain.sol";
import "./interfaces/ITendermintLightClient.sol";
import "./lib/Bytes.sol";
import "./lib/Memory.sol";
import "./lib/Tendermint.sol";
import "./lib/ics23/ics23.sol";
import "./Initializable.sol";
import "./System.sol";
import {PROOFS_PROTO_GLOBAL_ENUMS, CommitmentProof, ProofSpec, InnerSpec, LeafOp, InnerOp} from "./lib/proto/Proofs.sol";

contract CrossChain is Initializable, ICrossChain {
    using Bytes for bytes;
    using Bytes for bytes32;
    using TendermintHelper for TmHeader.Data;
    using TendermintHelper for ConsensusState.Data;
    using TendermintHelper for ValidatorSet.Data;

    // constant variables
    string public constant STORE_NAME = "ibc";
    uint256 public constant CROSS_CHAIN_KEY_PREFIX = 0x01006000; // last 6 bytes
    uint8 public constant SYN_PACKAGE = 0x00;
    uint8 public constant ACK_PACKAGE = 0x01;
    uint8 public constant FAIL_ACK_PACKAGE = 0x02;
    uint256 public constant INIT_BATCH_SIZE = 50;

    // governable parameters
    uint256 public batchSizeForOracle;

    //state variables
    uint256 public previousTxHeight;
    uint256 public txCounter;
    int64 public oracleSequence;
    mapping(uint8 => address) public channelHandlerContractMap;
    mapping(address => mapping(uint8 => bool)) public registeredContractChannelMap;
    mapping(uint8 => uint64) public channelSendSequenceMap;
    mapping(uint8 => uint64) public channelReceiveSequenceMap;
    System private system;

    struct ChannelInit {
        uint8 channelId;
        uint64 sequence;
    }

    ProofSpec.Data private _tmProofSpec =
        ProofSpec.Data({
            leaf_spec: LeafOp.Data({
                hash: PROOFS_PROTO_GLOBAL_ENUMS.HashOp.SHA256,
                prehash_key: PROOFS_PROTO_GLOBAL_ENUMS.HashOp.NO_HASH,
                prehash_value: PROOFS_PROTO_GLOBAL_ENUMS.HashOp.SHA256,
                length: PROOFS_PROTO_GLOBAL_ENUMS.LengthOp.VAR_PROTO,
                prefix: hex"00028cc3f922"
            }),
            inner_spec: InnerSpec.Data({
                child_order: getTmChildOrder(),
                child_size: 32,
                min_prefix_length: 1,
                max_prefix_length: 128,
                empty_child: abi.encodePacked(),
                hash: PROOFS_PROTO_GLOBAL_ENUMS.HashOp.SHA256
            }),
            min_depth: 0,
            max_depth: 0
        });

    // event
    event CrossChainPackage(
        uint16 chainId,
        uint64 indexed oracleSequence,
        uint64 indexed packageSequence,
        uint8 indexed channelId,
        bytes payload
    );
    event ReceivedPackage(uint8 packageType, uint64 indexed packageSequence, uint8 indexed channelId);
    event UnsupportedPackage(uint64 indexed packageSequence, uint8 indexed channelId, bytes payload);
    event UnexpectedRevertInPackageHandler(address indexed contractAddr, string reason);
    event UnexpectedFailureAssertionInPackageHandler(address indexed contractAddr, bytes lowLevelData);

    modifier sequenceInOrder(uint64 _sequence, uint8 _channelID) {
        uint64 expectedSequence = channelReceiveSequenceMap[_channelID];
        require(_sequence == expectedSequence, "sequence not in order");

        channelReceiveSequenceMap[_channelID] = expectedSequence + 1;
        _;
    }

    modifier blockSynced(uint64 _height) {
        require(
            ITendermintLightClient(system.tmLightClient()).isHeaderSynced(_height),
            "light client not sync the block yet"
        );
        _;
    }

    modifier channelSupported(uint8 _channelID) {
        require(channelHandlerContractMap[_channelID] != address(0x0), "channel is not supported");
        _;
    }

    modifier onlyRegisteredContractChannel(uint8 channelId) {
        require(
            registeredContractChannelMap[msg.sender][channelId],
            "the contract and channel have not been registered"
        );
        _;
    }

    // | length   | prefix | sourceChainID| destinationChainID | channelID | sequence |
    // | 32 bytes | 1 byte | 2 bytes      | 2 bytes            |  1 bytes  | 8 bytes  |
    function generateKey(uint64 _sequence, uint8 _channelID) internal pure returns (bytes memory) {
        uint256 fullCROSS_CHAIN_KEY_PREFIX = CROSS_CHAIN_KEY_PREFIX | _channelID;
        bytes memory key = new bytes(14);

        uint256 ptr;
        assembly {
            ptr := add(key, 14)
        }
        assembly {
            mstore(ptr, _sequence)
        }
        ptr -= 8;
        assembly {
            mstore(ptr, fullCROSS_CHAIN_KEY_PREFIX)
        }
        ptr -= 6;
        assembly {
            mstore(ptr, 14)
        }
        return key;
    }

    function init(address _system, ChannelInit[] memory receiveChannelInit) external onlyUninitialized {
        system = System(_system);
        require(system.bscValidatorSet() != address(0), "system uninitialized");

        channelHandlerContractMap[system.STAKING_CHANNEL_ID()] = system.bscValidatorSet();
        registeredContractChannelMap[system.bscValidatorSet()][system.STAKING_CHANNEL_ID()] = true;

        batchSizeForOracle = INIT_BATCH_SIZE;

        oracleSequence = -1;
        previousTxHeight = 0;
        txCounter = 0;

        for (uint256 i = 0; i < receiveChannelInit.length; i++) {
            ChannelInit memory channelInit = receiveChannelInit[i];
            channelReceiveSequenceMap[channelInit.channelId] = channelInit.sequence;
        }

        _initialized = true;
    }

    function encodePayload(
        uint8 packageType,
        uint256 relayFee,
        bytes memory msgBytes
    ) public pure returns (bytes memory) {
        uint256 payloadLength = msgBytes.length + 33;
        bytes memory payload = new bytes(payloadLength);
        uint256 ptr;
        assembly {
            ptr := payload
        }
        ptr += 33;

        assembly {
            mstore(ptr, relayFee)
        }

        ptr -= 32;
        assembly {
            mstore(ptr, packageType)
        }

        ptr -= 1;
        assembly {
            mstore(ptr, payloadLength)
        }

        ptr += 65;
        (uint256 src, ) = Memory.fromBytes(msgBytes);
        Memory.copy(src, ptr, msgBytes.length);

        return payload;
    }

    // | type   | relayFee   |package  |
    // | 1 byte | 32 bytes   | bytes    |
    function decodePayloadHeader(bytes memory payload) internal pure returns (bool, uint8, uint256, bytes memory) {
        if (payload.length < 33) {
            return (false, 0, 0, new bytes(0));
        }

        uint256 ptr;
        assembly {
            ptr := payload
        }

        uint8 packageType;
        ptr += 1;
        assembly {
            packageType := mload(ptr)
        }

        uint256 relayFee;
        ptr += 32;
        assembly {
            relayFee := mload(ptr)
        }

        ptr += 32;
        bytes memory msgBytes = new bytes(payload.length - 33);
        (uint256 dst, ) = Memory.fromBytes(msgBytes);
        Memory.copy(ptr, dst, payload.length - 33);

        return (true, packageType, relayFee, msgBytes);
    }

    function handlePackage(
        bytes calldata payload,
        bytes calldata proof,
        uint64 height,
        uint64 packageSequence,
        uint8 channelId
    )
        external
        onlyInitialized
        sequenceInOrder(packageSequence, channelId)
        blockSynced(height)
        channelSupported(channelId)
    {
        require(msg.sender == system.relayer(), "not relayer");

        bytes memory payloadLocal = payload; // fix error: stack too deep, try removing local variables
        bytes memory proofLocal = proof; // fix error: stack too deep, try removing local variables
        // TODO: Enable after BSC switches to ics-23 proofs
        // require(
        //     verifyMembership(
        //         proofLocal,
        //         ITendermintLightClient(system.tmLightClient()).getAppHash(height).toBytes(),
        //         "",
        //         bytes(generateKey(packageSequence, channelId))),
        //         payloadLocal
        //     )
        // );

        uint8 channelIdLocal = channelId; // fix error: stack too deep, try removing local variables
        (bool success, uint8 packageType, , bytes memory msgBytes) = decodePayloadHeader(payloadLocal);
        if (!success) {
            emit UnsupportedPackage(packageSequence, channelIdLocal, payloadLocal);
            return;
        }
        emit ReceivedPackage(packageType, packageSequence, channelIdLocal);
        if (packageType == SYN_PACKAGE) {
            address handlerContract = channelHandlerContractMap[channelIdLocal];
            try IApplication(handlerContract).handleSynPackage(channelIdLocal, msgBytes) returns (
                bytes memory responsePayload
            ) {
                if (responsePayload.length != 0) {
                    sendPackage(
                        channelSendSequenceMap[channelIdLocal],
                        channelIdLocal,
                        encodePayload(ACK_PACKAGE, 0, responsePayload)
                    );
                    channelSendSequenceMap[channelIdLocal] = channelSendSequenceMap[channelIdLocal] + 1;
                }
            } catch Error(string memory reason) {
                sendPackage(
                    channelSendSequenceMap[channelIdLocal],
                    channelIdLocal,
                    encodePayload(FAIL_ACK_PACKAGE, 0, msgBytes)
                );
                channelSendSequenceMap[channelIdLocal] = channelSendSequenceMap[channelIdLocal] + 1;
                emit UnexpectedRevertInPackageHandler(handlerContract, reason);
            } catch (bytes memory lowLevelData) {
                sendPackage(
                    channelSendSequenceMap[channelIdLocal],
                    channelIdLocal,
                    encodePayload(FAIL_ACK_PACKAGE, 0, msgBytes)
                );
                channelSendSequenceMap[channelIdLocal] = channelSendSequenceMap[channelIdLocal] + 1;
                emit UnexpectedFailureAssertionInPackageHandler(handlerContract, lowLevelData);
            }
        } else if (packageType == ACK_PACKAGE) {
            address handlerContract = channelHandlerContractMap[channelIdLocal];
            try IApplication(handlerContract).handleAckPackage(channelIdLocal, msgBytes) {} catch Error(
                string memory reason
            ) {
                emit UnexpectedRevertInPackageHandler(handlerContract, reason);
            } catch (bytes memory lowLevelData) {
                emit UnexpectedFailureAssertionInPackageHandler(handlerContract, lowLevelData);
            }
        } else if (packageType == FAIL_ACK_PACKAGE) {
            address handlerContract = channelHandlerContractMap[channelIdLocal];
            try IApplication(handlerContract).handleFailAckPackage(channelIdLocal, msgBytes) {} catch Error(
                string memory reason
            ) {
                emit UnexpectedRevertInPackageHandler(handlerContract, reason);
            } catch (bytes memory lowLevelData) {
                emit UnexpectedFailureAssertionInPackageHandler(handlerContract, lowLevelData);
            }
        }
    }

    function sendPackage(uint64 packageSequence, uint8 channelId, bytes memory payload) internal {
        if (block.number > previousTxHeight) {
            oracleSequence++;
            txCounter = 1;
            previousTxHeight = block.number;
        } else {
            txCounter++;
            if (txCounter > batchSizeForOracle) {
                oracleSequence++;
                txCounter = 1;
            }
        }
        emit CrossChainPackage(system.bscChainID(), uint64(oracleSequence), packageSequence, channelId, payload);
    }

    function sendSynPackage(
        uint8 channelId,
        bytes calldata msgBytes,
        uint256 relayFee
    ) external override onlyInitialized onlyRegisteredContractChannel(channelId) {
        uint64 sendSequence = channelSendSequenceMap[channelId];
        sendPackage(sendSequence, channelId, encodePayload(SYN_PACKAGE, relayFee, msgBytes));
        sendSequence++;
        channelSendSequenceMap[channelId] = sendSequence;
    }

    function getTmChildOrder() internal pure returns (int32[] memory) {
        int32[] memory childOrder = new int32[](2);
        childOrder[0] = 0;
        childOrder[1] = 1;

        return childOrder;
    }

    function verifyMembership(
        bytes memory proof,
        bytes memory root,
        bytes memory prefix,
        bytes memory slot,
        bytes memory expectedValue
    ) internal view returns (bool) {
        CommitmentProof.Data memory commitmentProof = CommitmentProof.decode(proof);

        Ics23.VerifyMembershipError vCode = Ics23.verifyMembership(
            _tmProofSpec,
            root,
            commitmentProof,
            slot,
            expectedValue
        );

        return vCode == Ics23.VerifyMembershipError.None;
    }
}
