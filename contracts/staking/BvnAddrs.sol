// SPDX-License-Identifier: GPL-3.0-only

pragma solidity ^0.8.18;

import {DataTypes as dt} from "./DataTypes.sol";
import "./Staking.sol";

/**
 * @title Maintain validator ETH-BVN address mapping
 */
contract BvnAddrs {
    Staking public immutable staking;
    mapping(address => bytes) public addrs;
    event BvnAddrUpdate(address indexed valAddr, bytes oldAddr, bytes newAddr);

    /**
     * @dev Need to deploy Staking contract first before deploying BvnAddr contract
     * @param _staking address of Staking Contract
     */
    constructor(Staking _staking) {
        staking = _staking;
    }

    /**
     * @notice Update bvn address
     * @param _bvnAddr the new address in the BVN
     */
    function updateBvnAddr(bytes calldata _bvnAddr) external {
        address valAddr = msg.sender;
        if (staking.signerVals(msg.sender) != address(0)) {
            valAddr = staking.signerVals(msg.sender);
        }

        dt.ValidatorStatus status = staking.getValidatorStatus(valAddr);
        require(status == dt.ValidatorStatus.Unbonded, "Not unbonded validator");

        bytes memory oldAddr = addrs[valAddr];
        addrs[valAddr] = _bvnAddr;

        staking.validatorNotice(valAddr, "bvn-addr", _bvnAddr);
        emit BvnAddrUpdate(valAddr, oldAddr, _bvnAddr);
    }
}
