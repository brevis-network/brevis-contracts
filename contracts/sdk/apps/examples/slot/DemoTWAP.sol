// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../../framework/BrevisApp.sol";
import "../../../lib/Lib.sol";
import "../../../interface/IBrevisProof.sol";

// Multiple storage slots: TWAP over an interval in Uniswap
// Use Brevis to prove the time-weighted average price of the WETH-USDC pair over a given block range in Uniswap.

contract DemoTWAP is BrevisApp {
    struct Observation {
        // the block timestamp of the observation
        uint32 blockTimestamp;
        // the tick accumulator, i.e. tick * time elapsed since the pool was first initialized
        int56 tickCumulative;
        // the seconds per liquidity, i.e. seconds elapsed / max(1, liquidity) since the pool was first initialized
        uint160 secondsPerLiquidityCumulativeX128;
        // whether or not the observation is initialized
        bool initialized;
    }

    address public poolAddr; // Uniswap V3 WETH-USDC pool addr
    bytes32 public slotHash; // keccak hash of the index of the Observation slot

    constructor(IBrevisProof _brevisProof, address _poolAddr, bytes32 _slotHash) BrevisApp(_brevisProof) {
        poolAddr = _poolAddr;
        slotHash = _slotHash;
    }

    function submitWETHUSDCPoolStorageInfosAndReturnTWATick(
        bytes32 _proofRequestId,
        uint64 _chainId,
        bytes32 _smtRoot,
        Brevis.StorageInfo[] calldata _info
    ) external view returns (int56 twaTick) {
        require(_info.length == 2, "only needs two uniswap pool observations to calculate twa Tick");
        Brevis.ExtractInfos memory info;
        info.smtRoot = _smtRoot;
        info.stores = _info;
        for (uint256 i = 0; i < 2; i++) {
            require(_info[i].account == poolAddr, "not exepcted pool");
            require(_info[i].slot == slotHash, "not expected slot");
        }
        validateRequest(_proofRequestId, _chainId, info);

        require(_info[1].blockNumber > _info[0].blockNumber, "not right sequence");
        Observation memory o1 = unpackObs(uint256(_info[1].slotValue));
        Observation memory o0 = unpackObs(uint256(_info[0].slotValue));

        return (o1.tickCumulative - o1.tickCumulative) / int56(uint56(o1.blockTimestamp - o0.blockTimestamp));
    }

    function unpackObs(uint256 observation) private pure returns (Observation memory) {
        return
            Observation({
                blockTimestamp: uint32(observation),
                tickCumulative: int56(uint56(observation >> 32)),
                secondsPerLiquidityCumulativeX128: uint160(observation >> 88),
                initialized: true
            });
    }
}
