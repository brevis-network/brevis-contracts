// SPDX-License-Identifier: GPL-3.0-only

pragma solidity ^0.8.18;

import "./OwnerProxyBase.sol";
import "../interfaces/ILightClientOwner.sol";
import {SimpleGovernance as sg} from "../SimpleGovernance.sol";
import {OwnerDataTypes as dt} from "./OwnerDataTypes.sol";

abstract contract LightClientOwnerProxy is OwnerProxyBase {
    event SetLightClientProposalCreated(uint256 proposalId, address target, address lightClient);

    event UpdateForkVersionProposalCreated(uint256 proposalId, address target, uint64 epoch, bytes4 forkVersion);

    event ProcessLightClientForceUpdateProposalCreated(uint256 proposalId, address target);

    function proposeSetLightClient(address _target, address _lightClient) external {
        bytes memory data = abi.encodeWithSelector(ILightClientOwner.setLightClient.selector, _lightClient);
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalDefault);
        emit SetLightClientProposalCreated(proposalId, _target, _lightClient);
    }

    function proposeUpdateForkVersion(address _target, uint64 _epoch, bytes4 _forkVersion) external {
        bytes memory data = abi.encodeWithSelector(ILightClientOwner.updateForkVersion.selector, _epoch, _forkVersion);
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalDefault);
        emit UpdateForkVersionProposalCreated(proposalId, _target, _epoch, _forkVersion);
    }

    function proposeProcessLightClientForceUpdate(address _target) external {
        bytes memory data = abi.encodeWithSelector(ILightClientOwner.processLightClientForceUpdate.selector);
        uint256 proposalId = gov.createProposal(msg.sender, _target, data, sg.ProposalType.ExternalDefault);
        emit ProcessLightClientForceUpdateProposalCreated(proposalId, _target);
    }
}
