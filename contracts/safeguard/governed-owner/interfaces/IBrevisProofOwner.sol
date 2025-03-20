// SPDX-License-Identifier: GPL-3.0-only

pragma solidity >=0.8.0;

interface IBrevisProofOwner {
    function updateSmtContract(address _smtContract) external;

    function updateAggProofVerifierAddresses(
        uint64[] calldata _chainIds,
        address[] calldata _verifierAddresses
    ) external;

    function setDummyInputCommitments(uint64[] calldata _chainIds, bytes32[] calldata _dummyInputCommitments) external;

    function setAggVkHash(bytes32 _aggVkHash) external;

    function updateVerifierAddress(uint64[] calldata _chainIds, address[] calldata _verifierAddresses) external;
}
