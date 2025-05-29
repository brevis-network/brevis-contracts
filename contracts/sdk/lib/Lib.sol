// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../../lib/RLPReader.sol";

library Brevis {
    // retrieved from proofData, to align the logs with circuit...
    struct ProofData {
        bytes32 commitHash;
        bytes32 appCommitHash; // zk-program computing circuit commit hash
        bytes32 appVkHash; // zk-program computing circuit Verify Key hash
        bytes32 smtRoot;
        bytes32 dummyInputCommitment; // zk-program computing circuit dummy input commitment
    }

    struct ProofAppData {
        bytes32 appCommitHash;
        bytes32 appVkHash;
    }
}