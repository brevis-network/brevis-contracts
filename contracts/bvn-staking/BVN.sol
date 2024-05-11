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
        uint64 timestamp;
    }

    Staking public immutable staking;
    address[] public registeredValidators;
    mapping(address => BrevisValidator) public brevisValidators; // valAddr -> BvnValidator
    mapping(address => address) public signerVals; // signerAddr -> valAddr
    mapping(uint64 => SlashRecord) public slashRecords; // nonce -> SlashRecord

    event BrevisValidatorRegistered(address indexed valAddr, address signer, bytes bvnAddr);
    event BrevisValidatorDeregistered(address indexed valAddr);
    event BrevisValidatorSignerUpdated(address indexed valAddr, address prevSigner, address newSigner);
    event Slash(address indexed valAddr, uint64 nonce, uint64 timestamp, string reason);

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
        require(status == dt.ValidatorStatus.Bonded, "not bonded validator");
        require(_valAddr == msg.sender || _valAddr == staking.signerVals(msg.sender), "unauthorized caller");

        require(signerVals[_signer] == address(0), "signer already used");
        if (_signer != _valAddr) {
            require(brevisValidators[_signer].deregisterTime == 0, "signer is other validator");
        }
        signerVals[_signer] = _valAddr;

        BrevisValidator storage bv = brevisValidators[_valAddr];
        require(bv.deregisterTime < block.timestamp, "already registered validator");
        bv.signer = _signer;
        bv.bvnAddr = _bvnAddr;
        bv.deregisterTime = dt.MAX_INT;
        registeredValidators.push(_valAddr);

        staking.validatorNotice(_valAddr, "register", "");
        emit BrevisValidatorRegistered(_valAddr, _signer, _bvnAddr);
    }

    function updateValidatorSigner(address _signer) external {
        address valAddr = msg.sender;
        BrevisValidator storage bv = brevisValidators[valAddr];
        require(bv.deregisterTime != 0, "unregistered validator");
        require(signerVals[_signer] == address(0), "signer already used");
        if (_signer != valAddr) {
            require(brevisValidators[_signer].deregisterTime == 0, "signer is other validator");
        }
        address prevSigner = bv.signer;
        delete signerVals[bv.signer];
        bv.signer = _signer;
        signerVals[_signer] = valAddr;

        staking.validatorNotice(valAddr, "signer", abi.encodePacked(_signer));
        emit BrevisValidatorSignerUpdated(valAddr, prevSigner, _signer);
    }

    /**
     * @notice Leave BVN
     * @param _valAddr validator eth address
     */
    function deregisterBrevisValidator(address _valAddr) external {
        BrevisValidator storage bv = brevisValidators[_valAddr];
        require(bv.deregisterTime > block.timestamp, "not registered validator");
        if (_valAddr != msg.sender && _valAddr != staking.signerVals(msg.sender) && bv.signer != msg.sender) {
            // if not called by validator itself, require unbonded status
            dt.ValidatorStatus status = staking.getValidatorStatus(_valAddr);
            require(status == dt.ValidatorStatus.Unbonded, "not unbonded validator");
        }
        bv.deregisterTime = block.timestamp;
        delete signerVals[bv.signer];

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
     * @param _timestamp slash triggered time
     * @param _reason slash reason
     * @param _sigs bvn signatures
     */
    function slash(
        address _valAddr,
        uint64 _nonce,
        uint64 _timestamp,
        string calldata _reason,
        bytes[] calldata _sigs
    ) external {
        require(brevisValidators[_valAddr].deregisterTime > block.timestamp, "not registered validator");
        bytes32 domain = keccak256(abi.encodePacked(block.chainid, address(this), "Slash"));
        bytes32 slashHash = keccak256(abi.encodePacked(_valAddr, _nonce, _timestamp, _reason));
        verifySignatures(abi.encodePacked(domain, slashHash), _sigs);

        SlashRecord storage s = slashRecords[_nonce];
        require(s.valAddr == address(0), "used slash nonce");
        s.valAddr = _valAddr;
        s.reason = _reason;
        s.timestamp = _timestamp;
        staking.validatorNotice(_valAddr, "slash", abi.encodePacked(_nonce));
        emit Slash(_valAddr, _nonce, _timestamp, _reason);
    }

    function verifySignatures(bytes memory _msg, bytes[] calldata _sigs) public view returns (bool) {
        bytes32 hash = keccak256(_msg).toEthSignedMessageHash();
        uint256 signedTokens;
        address prev = address(0);
        uint256 quorum = (getBondedTokens() * 2) / 3 + 1;
        for (uint256 i = 0; i < _sigs.length; i++) {
            address signer = hash.recover(_sigs[i]);
            require(signer > prev, "signers not in ascending order");
            prev = signer;

            address valAddr = signerVals[signer];
            require(isBondedValidator(valAddr), "not bonded validator");
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
        return (brevisValidators[_valAddr].deregisterTime > block.timestamp);
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
