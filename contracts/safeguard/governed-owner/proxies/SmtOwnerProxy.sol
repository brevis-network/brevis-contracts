// SPDX-License-Identifier: GPL-3.0-only

pragma solidity ^0.8.18;

import "./OwnerProxyBase.sol";
import "../interfaces/ISmtOwner.sol";
import {SimpleGovernance as sg} from "../SimpleGovernance.sol";
import {OwnerDataTypes as dt} from "./OwnerDataTypes.sol";

abstract contract SmtOwnerProxy is OwnerProxyBase {
    event SetAnchorProviderProposalCreated(uint256 proposalId, address target, uint64 chainId, address anchorProvider);

    event SetVerifierProposalCreated(uint256 proposalId, address target, uint64 chainId, address verifier);

    event SetCircuitDigestProposalCreated(uint256 proposalId, address target, uint64 chainId, bytes32 circuitDigest);

    event SetRootUpdaterProposalCreated(uint256 proposalId, address target, address rootUpdater);

    function proposeSetAnchorProvider(address _target, uint64 _chainId, address _anchorProvider) external {
        bytes memory data = abi.encodeWithSelector(ISmtOwner.setAnchorProvider.selector, _chainId, _anchorProvider);
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalDefault);
        emit SetAnchorProviderProposalCreated(proposalId, _target, _chainId, _anchorProvider);
    }

    function proposeSetVerifier(address _target, uint64 _chainId, address _verifier) external {
        bytes memory data = abi.encodeWithSelector(ISmtOwner.setVerifier.selector, _chainId, _verifier);
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalDefault);
        emit SetVerifierProposalCreated(proposalId, _target, _chainId, _verifier);
    }

    function proposeSetCircuitDigest(address _target, uint64 _chainId, bytes32 _circuitDigest) external {
        bytes memory data = abi.encodeWithSelector(ISmtOwner.setCircuitDigest.selector, _chainId, _circuitDigest);
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalDefault);
        emit SetCircuitDigestProposalCreated(proposalId, _target, _chainId, _circuitDigest);
    }

    function proposeSetRootUpdater(address _target, address _rootUpdater) external {
        bytes memory data = abi.encodeWithSelector(ISmtOwner.setRootUpdater.selector, _rootUpdater);
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalDefault);
        emit SetRootUpdaterProposalCreated(proposalId, _target, _rootUpdater);
    }
}
