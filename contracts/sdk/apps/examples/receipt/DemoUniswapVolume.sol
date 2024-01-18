// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../../framework/SimpleReceiptApp.sol";
import "../../../lib/Lib.sol";
import "../../../interface/IBrevisProof.sol";

// Multiple receipts: Aggregate trading volume in Uniswap pool
// Use Brevis to prove someone’s accumulative trading volume for a given pair in Uniswap.

contract DemoUniswapVolume is SimpleReceiptApp {
    mapping(address => uint256) public volumes;
    event SumVolume(address user, uint64 fromChain, uint256 volume);

    constructor(IBrevisProof _brevisProof) BrevisApp(_brevisProof) {}

    function submitUserVolume(
        bytes32 _proofRequestId,
        uint64 _chainId,
        bytes32 _smtRoot,
        Brevis.ReceiptInfo[] calldata _info
    ) external {
        Brevis.ExtractInfos memory info;
        info.smtRoot = _smtRoot;
        info.receipts = _info;
        validateRequest(_proofRequestId, _chainId, info);

        address swapper = address(bytes20(_info[0].logs[0].value));
        uint256 volume = abs(int256(uint256(_info[0].logs[1].value)));
        if (_info.length > 1) {
            for (uint256 i = 1; i < _info.length; i++) {
                require(swapper == address(bytes20(_info[i].logs[0].value)), "not a unique swapper");
                volume += abs(int256(uint256(_info[i].logs[1].value)));
            }
        }

        volumes[swapper] = volume;
        emit SumVolume(swapper, _chainId, volume);
    }

    function abs(int256 x) private pure returns (uint256) {
        return uint256(x >= 0 ? x : -x);
    }

    function getAttestedSwapSumVolume(address _swapper) external view returns (uint256) {
        return volumes[_swapper];
    }
}
