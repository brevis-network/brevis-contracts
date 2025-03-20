// SPDX-License-Identifier: GPL-3.0-only

pragma solidity ^0.8.18;

import "./OwnerProxyBase.sol";
import "../interfaces/ICommonOwner.sol";
import {SimpleGovernance as sg} from "../SimpleGovernance.sol";
import {OwnerDataTypes as dt} from "./OwnerDataTypes.sol";

abstract contract CommonOwnerProxy is OwnerProxyBase {
    event TransferOwnershipProposalCreated(uint256 proposalId, address target, address newOwner);

    event UpdatePauserProposalCreated(uint256 proposalId, address target, dt.Action action, address account);

    event UpdatePausersProposalCreated(uint256 proposalId, address target, dt.Action action, address[] accounts);

    event UpdateProversProposalCreated(uint256 proposalId, address target, dt.Action action, address[] accounts);

    event UpdateGovernorProposalCreated(uint256 proposalId, address target, dt.Action action, address account);

    event UpdateGovernorsProposalCreated(uint256 proposalId, address target, dt.Action action, address[] accounts);

    function proposeTransferOwnership(address _target, address _newOwner) external {
        bytes memory data = abi.encodeWithSelector(ICommonOwner.transferOwnership.selector, _newOwner);
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalDefault);
        emit TransferOwnershipProposalCreated(proposalId, _target, _newOwner);
    }

    function proposeUpdatePauser(address _target, dt.Action _action, address _account) external {
        bytes4 selector;
        if (_action == dt.Action.Add) {
            selector = ICommonOwner.addPauser.selector;
        } else if (_action == dt.Action.Remove) {
            selector = ICommonOwner.removePauser.selector;
        } else {
            revert("invalid action");
        }
        bytes memory data = abi.encodeWithSelector(selector, _account);
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalFastPass);
        emit UpdatePauserProposalCreated(proposalId, _target, _action, _account);
    }

    function proposeUpdatePausers(address _target, dt.Action _action, address[] calldata _accounts) external {
        bytes4 selector;
        if (_action == dt.Action.Add) {
            selector = ICommonOwner.addPausers.selector;
        } else if (_action == dt.Action.Remove) {
            selector = ICommonOwner.removePausers.selector;
        } else {
            revert("invalid action");
        }
        bytes memory data = abi.encodeWithSelector(selector, _accounts);
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalFastPass);
        emit UpdatePausersProposalCreated(proposalId, _target, _action, _accounts);
    }

    function proposeUpdateProvers(address _target, dt.Action _action, address[] calldata _accounts) external {
        bytes4 selector;
        if (_action == dt.Action.Add) {
            selector = ICommonOwner.addProvers.selector;
        } else if (_action == dt.Action.Remove) {
            selector = ICommonOwner.removeProvers.selector;
        } else {
            revert("invalid action");
        }
        bytes memory data = abi.encodeWithSelector(selector, _accounts);
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalDefault);
        emit UpdateProversProposalCreated(proposalId, _target, _action, _accounts);
    }

    function proposeUpdateGovernor(address _target, dt.Action _action, address _account) external {
        bytes4 selector;
        if (_action == dt.Action.Add) {
            selector = ICommonOwner.addGovernor.selector;
        } else if (_action == dt.Action.Remove) {
            selector = ICommonOwner.removeGovernor.selector;
        } else {
            revert("invalid action");
        }
        bytes memory data = abi.encodeWithSelector(selector, _account);
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalDefault);
        emit UpdateGovernorProposalCreated(proposalId, _target, _action, _account);
    }

    function proposeUpdateGovernors(address _target, dt.Action _action, address[] calldata _accounts) external {
        bytes4 selector;
        if (_action == dt.Action.Add) {
            selector = ICommonOwner.addGovernors.selector;
        } else if (_action == dt.Action.Remove) {
            selector = ICommonOwner.removeGovernors.selector;
        } else {
            revert("invalid action");
        }
        bytes memory data = abi.encodeWithSelector(selector, _accounts);
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalDefault);
        emit UpdateGovernorsProposalCreated(proposalId, _target, _action, _accounts);
    }
}
