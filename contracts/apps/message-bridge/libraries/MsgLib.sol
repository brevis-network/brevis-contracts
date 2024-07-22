// SPDX-License-Identifier: GPL-3.0-only

pragma solidity >=0.8.0;

import "../../../lib/Utils.sol";

library MsgLib {
    string constant ABORT_PREFIX = "MSG::ABORT:";

    function computeMessageId(
        uint64 _nonce,
        address _sender,
        address _receiver,
        uint64 _srcChainId,
        uint64 _dstChainId,
        bytes calldata _message
    ) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(_nonce, _sender, _receiver, _srcChainId, _dstChainId, _message));
    }

    function checkRevertMsg(bytes memory _returnData) internal pure returns (string memory) {
        string memory revertMsg = Utils.getRevertMsg(_returnData);
        checkAbortPrefix(revertMsg);
        return revertMsg;
    }

    function checkAbortPrefix(string memory _revertMsg) private pure {
        bytes memory prefixBytes = bytes(ABORT_PREFIX);
        bytes memory msgBytes = bytes(_revertMsg);
        if (msgBytes.length >= prefixBytes.length) {
            for (uint256 i = 0; i < prefixBytes.length; i++) {
                if (msgBytes[i] != prefixBytes[i]) {
                    return; // prefix not match, return
                }
            }
            revert(_revertMsg); // prefix match, revert
        }
    }
}
