// SPDX-License-Identifier: GPL-3.0-only

pragma solidity ^0.8.18;

import "./OwnerProxyBase.sol";
import "../interfaces/IBrevisProofOwner.sol";
import {SimpleGovernance as sg} from "../SimpleGovernance.sol";
import {OwnerDataTypes as dt} from "./OwnerDataTypes.sol";

abstract contract BrevisProofOwnerProxy is OwnerProxyBase {
    event UpdateSmtContractProposalCreated(uint256 proposalId, address target, address newSmtContract);

    event UpdateVerifierAddressProposalCreated(
        uint256 proposalId,
        address target,
        uint64[] chainIds,
        address[] verifierAddresses
    );

    event UpdateAggProofVerifierAddressesProposalCreated(
        uint256 proposalId,
        address target,
        uint64[] chainIds,
        address[] verifierAddresses
    );

    event SetDummyInputCommitmentsProposalCreated(
        uint256 proposalId,
        address target,
        uint64[] chainIds,
        bytes32[] dummyInputCommitments
    );

    event SetAggVkHashProposalCreated(uint256 proposalId, address target, bytes32 aggVkHash);

    function proposeUpdateSmtContract(address _target, address _newSmtContract) external {
        bytes memory data = abi.encodeWithSelector(IBrevisProofOwner.updateSmtContract.selector, _newSmtContract);
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalDefault);
        emit UpdateSmtContractProposalCreated(proposalId, _target, _newSmtContract);
    }

    function proposeUpdateVerifierAddress(
        address _target,
        uint64[] calldata _chainIds,
        address[] calldata _verifierAddresses
    ) external {
        bytes memory data = abi.encodeWithSelector(
            IBrevisProofOwner.updateAggProofVerifierAddresses.selector,
            _chainIds,
            _verifierAddresses
        );
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalDefault);
        emit UpdateVerifierAddressProposalCreated(proposalId, _target, _chainIds, _verifierAddresses);
    }

    function proposeUpdateAggProofVerifierAddresses(
        address _target,
        uint64[] calldata _chainIds,
        address[] calldata _newVerifierAddresses
    ) external {
        bytes memory data = abi.encodeWithSelector(
            IBrevisProofOwner.updateAggProofVerifierAddresses.selector,
            _chainIds,
            _newVerifierAddresses
        );
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalDefault);
        emit UpdateAggProofVerifierAddressesProposalCreated(proposalId, _target, _chainIds, _newVerifierAddresses);
    }

    function proposeSetDummyInputCommitments(
        address _target,
        uint64[] calldata _chainIds,
        bytes32[] calldata _dummyInputCommitments
    ) external {
        bytes memory data = abi.encodeWithSelector(
            IBrevisProofOwner.setDummyInputCommitments.selector,
            _chainIds,
            _dummyInputCommitments
        );
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalDefault);
        emit SetDummyInputCommitmentsProposalCreated(proposalId, _target, _chainIds, _dummyInputCommitments);
    }

    function proposeSetAggVkHashProposal(address _target, bytes32 _aggVkHash) external {
        bytes memory data = abi.encodeWithSelector(IBrevisProofOwner.setAggVkHash.selector, _aggVkHash);
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalDefault);
        emit SetAggVkHashProposalCreated(proposalId, _target, _aggVkHash);
    }
}
