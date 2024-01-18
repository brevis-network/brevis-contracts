// SPDX-License-Identifier: GPL-3.0-only

pragma solidity >=0.8.0;

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

    // https://ethereum.stackexchange.com/a/83577
    // https://github.com/Uniswap/v3-periphery/blob/v1.0.0/contracts/base/Multicall.sol
    function getRevertMsg(bytes memory _returnData) internal pure returns (string memory) {
        // If the _res length is less than 68, then the transaction failed silently (without a revert message)
        if (_returnData.length < 68) return "Transaction reverted silently";
        assembly {
            // Slice the sighash.
            _returnData := add(_returnData, 0x04)
        }
        return abi.decode(_returnData, (string)); // All that remains is the revert string
    }

    function checkRevertMsg(bytes memory _returnData) internal pure returns (string memory) {
        string memory revertMsg = MsgLib.getRevertMsg(_returnData);
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
