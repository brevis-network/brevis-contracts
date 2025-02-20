// SPDX-License-Identifier: GPL-3.0-only

pragma solidity ^0.8.18;

import "./OwnerProxyBase.sol";
import "../interfaces/IBrevisRequestOwner.sol";
import {SimpleGovernance as sg} from "../SimpleGovernance.sol";
import {OwnerDataTypes as dt} from "./OwnerDataTypes.sol";

abstract contract BrevisRequestOwnerProxy is OwnerProxyBase {
    event SetRequestTimeoutProposalCreated(uint256 proposalId, address target, uint256 timeout);

    event SetBaseDataURLProposalCreated(uint256 proposalId, address target, string url);

    event SetBrevisProofProposalCreated(uint256 proposalId, address target, address brevisProof);

    event SetBrevisDisputeProposalCreated(uint256 proposalId, address target, address brevisDispute);

    event SetBvnSigsVerifierProposalCreated(uint256 proposalId, address target, address bvnSigsVerifier);

    event SetAvsSigsVerifierProposalCreated(uint256 proposalId, address target, address avsSigsVerifier);

    event SetFeeCollectorProposalCreated(uint256 proposalId, address target, address feeCollector);

    event SetChallengeWindowProposalCreated(uint256 proposalId, address target, uint256 challengeWindow);

    event SetResponseTimeoutProposalCreated(uint256 proposalId, address target, uint256 responseTimeout);

    event SetDisputeDepositsProposalCreated(
        uint256 proposalId,
        address target,
        uint256 amtAskForData,
        uint256 amtAskForProof
    );

    function proposeSetRequestTimeout(address _target, uint256 _timeout) external {
        bytes memory data = abi.encodeWithSelector(IBrevisRequestOwner.setRequestTimeout.selector, _timeout);
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalDefault);
        emit SetRequestTimeoutProposalCreated(proposalId, _target, _timeout);
    }

    function proposeSetBaseDataURL(address _target, string calldata _url) external {
        bytes memory data = abi.encodeWithSelector(IBrevisRequestOwner.setBaseDataURL.selector, _url);
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalDefault);
        emit SetBaseDataURLProposalCreated(proposalId, _target, _url);
    }

    function proposeSetBrevisProof(address _target, address _brevisProof) external {
        bytes memory data = abi.encodeWithSelector(IBrevisRequestOwner.setBrevisProof.selector, _brevisProof);
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalDefault);
        emit SetBrevisProofProposalCreated(proposalId, _target, _brevisProof);
    }

    function proposeSetBrevisDispute(address _target, address _brevisDispute) external {
        bytes memory data = abi.encodeWithSelector(IBrevisRequestOwner.setBrevisDispute.selector, _brevisDispute);
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalDefault);
        emit SetBrevisDisputeProposalCreated(proposalId, _target, _brevisDispute);
    }

    function proposeSetBvnSigsVerifier(address _target, address _bvnSigsVerifier) external {
        bytes memory data = abi.encodeWithSelector(IBrevisRequestOwner.setBvnSigsVerifier.selector, _bvnSigsVerifier);
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalDefault);
        emit SetBvnSigsVerifierProposalCreated(proposalId, _target, _bvnSigsVerifier);
    }

    function proposeSetAvsSigsVerifier(address _target, address _avsSigsVerifier) external {
        bytes memory data = abi.encodeWithSelector(IBrevisRequestOwner.setAvsSigsVerifier.selector, _avsSigsVerifier);
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalDefault);
        emit SetAvsSigsVerifierProposalCreated(proposalId, _target, _avsSigsVerifier);
    }

    function proposeSetFeeCollector(address _target, address _feeCollector) external {
        bytes memory data = abi.encodeWithSelector(IBrevisRequestOwner.setFeeCollector.selector, _feeCollector);
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalDefault);
        emit SetFeeCollectorProposalCreated(proposalId, _target, _feeCollector);
    }

    function proposeSetChallengeWindow(address _target, uint256 _challengeWindow) external {
        bytes memory data = abi.encodeWithSelector(IBrevisRequestOwner.setChallengeWindow.selector, _challengeWindow);
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalDefault);
        emit SetChallengeWindowProposalCreated(proposalId, _target, _challengeWindow);
    }

    function proposeSetResponseTimeout(address _target, uint256 _responseTimeout) external {
        bytes memory data = abi.encodeWithSelector(IBrevisRequestOwner.setResponseTimeout.selector, _responseTimeout);
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalDefault);
        emit SetResponseTimeoutProposalCreated(proposalId, _target, _responseTimeout);
    }

    function proposeSetDisputeDeposits(address _target, uint256 _amtAskForData, uint256 _amtAskForProof) external {
        bytes memory data = abi.encodeWithSelector(
            IBrevisRequestOwner.setDisputeDeposits.selector,
            _amtAskForData,
            _amtAskForProof
        );
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalDefault);
        emit SetDisputeDepositsProposalCreated(proposalId, _target, _amtAskForData, _amtAskForProof);
    }
}
