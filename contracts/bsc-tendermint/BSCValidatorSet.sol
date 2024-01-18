// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

import "./Initializable.sol";
import "./System.sol";
import "./interfaces/IApplication.sol";
import "./interfaces/IBSCValidatorSet.sol";
import "./lib/CmnPkg.sol";
import "./lib/Memory.sol";
import "./lib/RLPDecode.sol";

contract BSCValidatorSet is Initializable, IApplication, IBSCValidatorSet {
    using RLPDecode for *;

    uint8 public constant VALIDATORS_UPDATE_MESSAGE_TYPE = 0;

    uint256 public constant EXPIRE_TIME_SECOND_GAP = 1000;
    uint256 public constant MAX_NUM_OF_VALIDATORS = 41;

    uint32 public constant ERROR_UNKNOWN_PACKAGE_TYPE = 101;
    uint32 public constant ERROR_FAIL_CHECK_VALIDATORS = 102;
    uint32 public constant ERROR_LEN_OF_VAL_MISMATCH = 103;

    uint256 public constant INIT_NUM_OF_CABINETS = 21;
    uint256 public constant EPOCH = 200;

    /*********************** state of the contract **************************/
    Validator[] public currentValidatorSet;
    uint256 public expireTimeSecondGap;

    System private system;

    // key is the `consensusAddress` of `Validator`,
    // value is the 1-based index of the element in `currentValidatorSet`.
    mapping(address => uint256) public currentValidatorSetMap;

    struct Validator {
        address consensusAddress;
        address payable feeAddress;
        address BBCFeeAddress;
        uint64 votingPower;
        // only in state
        bool jailed;
        uint256 incoming;
    }

    /*********************** cross chain package **************************/
    struct IbcValidatorSetPackage {
        uint8 packageType;
        Validator[] validatorSet;
    }

    /*********************** events **************************/
    event ValidatorSetUpdated();
    event UnexpectedPackage(uint8 channelId, bytes msgBytes);
    event FailedWithReasonStr(string message);

    /*********************** init **************************/
    function init(address _system, bytes memory _initValidatorSetBytes) external onlyUninitialized {
        (IbcValidatorSetPackage memory validatorSetPkg, bool valid) = decodeValidatorSetSynPackage(
            _initValidatorSetBytes
        );
        require(valid, "failed to parse init validatorSet");
        for (uint256 i; i < validatorSetPkg.validatorSet.length; ++i) {
            currentValidatorSet.push(validatorSetPkg.validatorSet[i]);
            currentValidatorSetMap[validatorSetPkg.validatorSet[i].consensusAddress] = i + 1;
        }
        expireTimeSecondGap = EXPIRE_TIME_SECOND_GAP;

        system = System(_system);

        _initialized = true;
    }

    /*********************** Cross Chain App Implement **************************/
    function handleSynPackage(
        uint8,
        bytes calldata msgBytes
    ) external override onlyInitialized returns (bytes memory responsePayload) {
        require(msg.sender == System(system).crossChain(), "not cross chain contract");

        (IbcValidatorSetPackage memory validatorSetPackage, bool ok) = decodeValidatorSetSynPackage(msgBytes);
        if (!ok) {
            return CmnPkg.encodeCommonAckPackage(system.ERROR_FAIL_DECODE());
        }
        uint32 resCode;
        if (validatorSetPackage.packageType == VALIDATORS_UPDATE_MESSAGE_TYPE) {
            resCode = updateValidatorSet(validatorSetPackage.validatorSet);
        } else {
            resCode = ERROR_UNKNOWN_PACKAGE_TYPE;
        }
        if (resCode == system.CODE_OK()) {
            return new bytes(0);
        } else {
            return CmnPkg.encodeCommonAckPackage(resCode);
        }
    }

    function handleAckPackage(uint8 channelId, bytes calldata msgBytes) external override {
        require(msg.sender == system.crossChain(), "not cross chain contract");

        // should not happen
        emit UnexpectedPackage(channelId, msgBytes);
    }

    function handleFailAckPackage(uint8 channelId, bytes calldata msgBytes) external override {
        require(msg.sender == system.crossChain(), "not cross chain contract");

        // should not happen
        emit UnexpectedPackage(channelId, msgBytes);
    }

    function updateValidatorSet(Validator[] memory validatorSet) internal returns (uint32) {
        {
            // do verify.
            (bool valid, string memory errMsg) = checkValidatorSet(validatorSet);
            if (!valid) {
                emit FailedWithReasonStr(errMsg);
                return ERROR_FAIL_CHECK_VALIDATORS;
            }
        }

        // update validator set state
        doUpdateState(validatorSet);

        emit ValidatorSetUpdated();
        return system.CODE_OK();
    }

    /*********************** Internal Functions **************************/

    function checkValidatorSet(Validator[] memory validatorSet) private pure returns (bool, string memory) {
        if (validatorSet.length > MAX_NUM_OF_VALIDATORS) {
            return (false, "the number of validators exceed the limit");
        }
        for (uint256 i; i < validatorSet.length; ++i) {
            for (uint256 j = 0; j < i; j++) {
                if (validatorSet[i].consensusAddress == validatorSet[j].consensusAddress) {
                    return (false, "duplicate consensus address of validatorSet");
                }
            }
        }
        return (true, "");
    }

    function doUpdateState(Validator[] memory validatorSet) private {
        uint256 n = currentValidatorSet.length;
        uint256 m = validatorSet.length;

        for (uint256 i; i < n; ++i) {
            bool stale = true;
            Validator memory oldValidator = currentValidatorSet[i];
            for (uint256 j = 0; j < m; j++) {
                if (oldValidator.consensusAddress == validatorSet[j].consensusAddress) {
                    stale = false;
                    break;
                }
            }
            if (stale) {
                delete currentValidatorSetMap[oldValidator.consensusAddress];
            }
        }

        if (n > m) {
            for (uint256 i = m; i < n; ++i) {
                currentValidatorSet.pop();
            }
        }
        uint256 k = n < m ? n : m;
        for (uint256 i; i < k; ++i) {
            if (!isSameValidator(validatorSet[i], currentValidatorSet[i])) {
                currentValidatorSetMap[validatorSet[i].consensusAddress] = i + 1;
                currentValidatorSet[i] = validatorSet[i];
            }
        }
    }

    function isSameValidator(Validator memory v1, Validator memory v2) private pure returns (bool) {
        return
            v1.consensusAddress == v2.consensusAddress &&
            v1.feeAddress == v2.feeAddress &&
            v1.BBCFeeAddress == v2.BBCFeeAddress &&
            v1.votingPower == v2.votingPower;
    }

    //rlp encode & decode function
    function decodeValidatorSetSynPackage(
        bytes memory msgBytes
    ) internal pure returns (IbcValidatorSetPackage memory, bool) {
        IbcValidatorSetPackage memory validatorSetPkg;

        RLPDecode.Iterator memory iter = msgBytes.toRLPItem().iterator();
        bool success = false;
        uint256 idx = 0;
        while (iter.hasNext()) {
            if (idx == 0) {
                validatorSetPkg.packageType = uint8(iter.next().toUint());
            } else if (idx == 1) {
                RLPDecode.RLPItem[] memory items = iter.next().toList();
                validatorSetPkg.validatorSet = new Validator[](items.length);
                for (uint256 j; j < items.length; ++j) {
                    (Validator memory val, bool ok) = decodeValidator(items[j]);
                    if (!ok) {
                        return (validatorSetPkg, false);
                    }
                    validatorSetPkg.validatorSet[j] = val;
                }
                success = true;
            } else {
                break;
            }
            idx++;
        }
        return (validatorSetPkg, success);
    }

    function decodeValidator(RLPDecode.RLPItem memory itemValidator) internal pure returns (Validator memory, bool) {
        Validator memory validator;
        RLPDecode.Iterator memory iter = itemValidator.iterator();
        bool success = false;
        uint256 idx = 0;
        while (iter.hasNext()) {
            if (idx == 0) {
                validator.consensusAddress = iter.next().toAddress();
            } else if (idx == 1) {
                validator.feeAddress = payable(iter.next().toAddress());
            } else if (idx == 2) {
                validator.BBCFeeAddress = iter.next().toAddress();
            } else if (idx == 3) {
                validator.votingPower = uint64(iter.next().toUint());
                success = true;
            } else {
                break;
            }
            idx++;
        }
        return (validator, success);
    }

    function isCurrentValidator(address valAddress) external view returns (bool) {
        return currentValidatorSetMap[valAddress] != 0;
    }
}
