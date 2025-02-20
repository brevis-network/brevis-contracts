// SPDX-License-Identifier: GPL-3.0-only

pragma solidity >=0.8.0;

interface ISmtOwner {
    function setAnchorProvider(uint64 chainId, address anchorProvider) external;

    function setVerifier(uint64 chainId, address verifier) external;

    function setCircuitDigest(uint64 chainId, bytes32 _circuitDigest) external;

    function setRootUpdater(address _rootUpdater) external;
}
