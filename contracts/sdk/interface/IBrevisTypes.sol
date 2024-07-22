// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

interface IBrevisTypes {
    enum RequestStatus {
        Null,
        ZkPending,
        ZkAttested,
        OpPending,
        OpSubmitted,
        OpDisputing,
        OpDisputed,
        OpAttested,
        Refunded
    }

    struct Request {
        RequestStatus status;
        uint64 timestamp;
        uint8 option;
    }

    struct OnchainRequestInfo {
        bytes32 feeHash; // keccak256(abi.encodePacked(amount, refundee))
        Callback callback;
    }

    struct Callback {
        address target;
        uint64 gas;
    }

    enum DisputeStatus {
        Null,
        WaitingForRequestData,
        RequestDataPosted,
        WaitingForDataAvailabilityProof,
        DataAvailabilityProofPosted,
        WaitingForDataValidityProof,
        DataValidityProofPosted
    }

    struct RequestDataHash {
        bytes32[] hashes;
        bytes32 root;
    }

    struct Dispute {
        DisputeStatus status;
        address challenger;
        RequestDataHash requestDataHash;
        uint256 responseDeadline;
        uint256 deposit;
    }
}
