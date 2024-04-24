// SPDX-License-Identifier: GPL-3.0-only

pragma solidity ^0.8.18;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {DataTypes as dt} from "./DataTypes.sol";
import "./Staking.sol";

/**
 * @title Brevis Validator Network
 */
contract BVN {
    using ECDSA for bytes32;

    struct BrevisValidator {
        address signer;
        bytes bvnAddr;
        uint256 deregisterTime;
    }

    struct SlashRecord {
        address valAddr;
        string reason;
        uint256 time;
    }

    Staking public immutable staking;
    address[] public registeredValidators;
    mapping(address => BrevisValidator) public brevisValidators; // valAddr -> BvnValidator
    mapping(uint64 => SlashRecord) public slashRecords; // nonce -> SlashRecord

    event BrevisValidatorRegistered(address indexed valAddr, address signer, bytes bvnAddr);
    event BrevisValidatorDeregistered(address indexed valAddr);
    event Slash(address indexed valAddr, uint64 nonce, string reason);

    /**
     * @param _staking address of Staking Contract
     */
    constructor(Staking _staking) {
        staking = _staking;
    }

    /**
     * @notice Join BVN
     * @param _valAddr validator eth address
     * @param _signer signer eth address
     * @param _bvnAddr bvn address
     */
    function registerBrevisValidator(address _valAddr, address _signer, bytes calldata _bvnAddr) external {
        dt.ValidatorStatus status = staking.getValidatorStatus(_valAddr);
        require(status == dt.ValidatorStatus.Bonded, "Not bonded validator");
        require(_valAddr == msg.sender || _valAddr == staking.signerVals(msg.sender), "unauthorized caller");

        BrevisValidator storage bv = brevisValidators[_valAddr];
        require(bv.deregisterTime != dt.MAX_INT, "already registered validator");
        bv.signer = _signer;
        bv.bvnAddr = _bvnAddr;
        bv.deregisterTime = dt.MAX_INT;
        registeredValidators.push(_valAddr);

        staking.validatorNotice(_valAddr, "register", "");
        emit BrevisValidatorRegistered(_valAddr, _signer, _bvnAddr);
    }

    /**
     * @notice Leave BVN
     * @param _valAddr validator eth address
     */
    function deregisterBrevisValidator(address _valAddr) external {
        if (_valAddr != msg.sender && _valAddr != staking.signerVals(msg.sender)) {
            // if not called by validator itself, require unbonded status
            dt.ValidatorStatus status = staking.getValidatorStatus(_valAddr);
            require(status == dt.ValidatorStatus.Unbonded, "Not unbonded validator");
        }

        BrevisValidator storage bv = brevisValidators[_valAddr];
        require(bv.deregisterTime == dt.MAX_INT, "not registered validator");
        bv.deregisterTime = block.timestamp;

        staking.validatorNotice(_valAddr, "deregister", "");
        uint256 lastIndex = registeredValidators.length - 1;
        for (uint256 i = 0; i < registeredValidators.length; i++) {
            if (registeredValidators[i] == _valAddr) {
                if (i < lastIndex) {
                    registeredValidators[i] = registeredValidators[lastIndex];
                }
                registeredValidators.pop();
                return;
            }
        }
        revert("validator not found"); // this should never happen
    }

    /**
     * @notice Slash a validator
     * @param _valAddr validator eth address
     * @param _nonce slash nonce
     * @param _reason slash reason
     * @param _sigs bvn signatures
     */
    function slash(address _valAddr, uint64 _nonce, string calldata _reason, bytes[] calldata _sigs) external {
        require(brevisValidators[_valAddr].deregisterTime == dt.MAX_INT, "not registered validator");
        bytes32 domain = keccak256(abi.encodePacked(block.chainid, address(this), "Slash"));
        verifySignatures(abi.encodePacked(domain, _valAddr, _nonce, _reason), _sigs);

        SlashRecord storage s = slashRecords[_nonce];
        require(s.time == 0, "slash record exists");
        s.valAddr = _valAddr;
        s.reason = _reason;
        s.time = block.timestamp;
        staking.validatorNotice(_valAddr, "slash", abi.encodePacked(_nonce));
        emit Slash(_valAddr, _nonce, _reason);
    }

    function verifySignatures(bytes memory _msg, bytes[] calldata _sigs) public view returns (bool) {
        bytes32 hash = keccak256(_msg).toEthSignedMessageHash();
        uint256 signedTokens;
        address prev = address(0);
        uint256 quorum = (getBondedTokens() * 2) / 3 + 1;
        for (uint256 i = 0; i < _sigs.length; i++) {
            address signer = hash.recover(_sigs[i]);
            require(signer > prev, "Signers not in ascending order");
            prev = signer;

            address valAddr = staking.signerVals(signer);
            BrevisValidator storage bv = brevisValidators[valAddr];
            require(bv.deregisterTime == dt.MAX_INT, "not registered validator");

            require(staking.isBondedValidator(valAddr), "not bonded validator");
            // TODO: gas optmization, getValidatorTokens already called in getTotalTokens()
            signedTokens += staking.getValidatorTokens(valAddr);
            if (signedTokens >= quorum) {
                return true;
            }
        }
        revert("Quorum not reached");
    }

    function getBondedTokens() public view returns (uint256) {
        uint256 totalTokens;
        for (uint32 i = 0; i < registeredValidators.length; i++) {
            address valAddr = registeredValidators[i];
            if (staking.isBondedValidator(valAddr)) {
                totalTokens += staking.getValidatorTokens(valAddr);
            }
        }
        return totalTokens;
    }

    function isRegisteredValidator(address _valAddr) public view returns (bool) {
        return (brevisValidators[_valAddr].deregisterTime == dt.MAX_INT);
    }

    function isBondedValidator(address _valAddr) public view returns (bool) {
        return (staking.isBondedValidator(_valAddr) && isRegisteredValidator(_valAddr));
    }

    function getBondedValidatorNum() public view returns (uint256) {
        uint256 num;
        for (uint32 i = 0; i < registeredValidators.length; i++) {
            if (staking.isBondedValidator(registeredValidators[i])) {
                num++;
            }
        }
        return num;
    }
}
