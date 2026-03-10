// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../interface/IBrevisApp.sol";
import "../interface/IPicoVerifier.sol";
import "../../interfaces/ISMT.sol";
import "../../safeguard/BrevisAccess.sol";

contract SimpleBrevisRequest is BrevisAccess {
    IPicoVerifier public picoVerifier;
    ISMT public smtContract;
    event RequestFulfilled(bytes32 requestKey);
    event RequestCallbackFailed(bytes32 requestKey);
    event PicoVerifierUpdated(IPicoVerifier from, IPicoVerifier to);
    event SmtContractUpdated(ISMT from, ISMT to);

    constructor(IPicoVerifier _picoVerifier, ISMT _smtContract) {
        picoVerifier = _picoVerifier;
        smtContract = _smtContract;
    }

    function fulfillRequest(
        uint64 _chainId,
        bytes32 _requestKey,
        bytes32 _riscvVkey,
        bytes calldata _publicValues,
        uint256[8] calldata _proof,
        address _callbackTarget
    ) external onlyActiveProver {
        require(_publicValues.length > 32, "_publicValues should start with smt root");
        picoVerifier.verifyPicoProof(_riscvVkey, _publicValues, _proof);

        bytes32 smtRoot = bytes32(_publicValues[:32]);
        require(smtContract.isSmtRootValid(_chainId, smtRoot), "smt root not valid");

        bytes memory appBizData = _publicValues[32:];
        bool success = _brevisCallback(_callbackTarget, _riscvVkey, appBizData);
        if (!success) {
            emit RequestCallbackFailed(_requestKey);
        }
        emit RequestFulfilled(_requestKey);
    }

    function _brevisCallback(
        address _callbackTarget,
        bytes32 _riscvVkey,
        bytes memory _appBizData
    ) private returns (bool) {
        if (_callbackTarget != address(0)) {
            uint256 gas = gasleft();
            (bool success, ) = _callbackTarget.call{gas: gas}(
                abi.encodeWithSelector(IBrevisApp.brevisCallback.selector, _riscvVkey, _appBizData)
            );
            if (!success) {
                return false;
            }
        }
        return true;
    }

    function setPicoVerifier(IPicoVerifier _picoVerifier) external onlyOwner {
        require(address(_picoVerifier) != address(0), "zero address");
        IPicoVerifier oldAddr = picoVerifier;
        picoVerifier = _picoVerifier;
        emit PicoVerifierUpdated(oldAddr, _picoVerifier);
    }

    function setSmtContract(ISMT _smtContract) external onlyOwner {
        require(address(_smtContract) != address(0), "zero address");
        ISMT oldAddr = smtContract;
        smtContract = _smtContract;
        emit SmtContractUpdated(oldAddr, _smtContract);
    }
}
