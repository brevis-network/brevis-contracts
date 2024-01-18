// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/access/Ownable.sol";
import "../../lib/Lib.sol";
import "./BrevisApp.sol";

abstract contract SimpleReceiptApp is BrevisApp, Ownable {
    Brevis.LogExtraInfo[] private logExtraInfos;

    event LogExtraInfoUpdated(Brevis.LogExtraInfo[] logExtraInfos);

    function validateRequest(
        bytes32 _requestId,
        uint64 _chainId,
        Brevis.ExtractInfos memory _extractInfos
    ) public view override returns (bool) {
        brevisProof.validateRequest(_requestId, _chainId, _extractInfos);
        if (_extractInfos.receipts.length > 0) {
            _validateLogExtraInfo(_extractInfos.receipts);
        }
        return true;
    }

    // default implementation, can be overwritten
    function _validateLogExtraInfo(Brevis.ReceiptInfo[] memory _infos) internal view virtual {
        Brevis.LogExtraInfo[] memory logExtras = logExtraInfos;
        uint256 receiptFieldsNum = logExtraInfos.length;
        for (uint256 i = 0; i < _infos.length; i++) {
            for (uint256 j = 0; j < receiptFieldsNum; j++) {
                Brevis.LogExtraInfo memory logExtra = _infos[i].logs[j].logExtraInfo;
                require(logExtras[j].valueFromTopic == logExtra.valueFromTopic, "wrong valueFromTopic");
                require(logExtras[j].valueIndex == logExtra.valueIndex, "wrong valueIndex");
                require(logExtras[j].contractAddress == logExtra.contractAddress, "wrong contractAddress");
                require(logExtras[j].logTopic0 == logExtra.logTopic0, "wrong logTopic0");
            }
        }
    }

    function setFieldLocations(Brevis.LogExtraInfo[] calldata _logExtraInfos) public onlyOwner {
        require(_logExtraInfos.length <= Brevis.NumField);
        logExtraInfos = _logExtraInfos;
        emit LogExtraInfoUpdated(_logExtraInfos);
    }

    function getFieldLocations() public view returns (Brevis.LogExtraInfo[] memory) {
        return logExtraInfos;
    }
}
