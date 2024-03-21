// SPDX-License-Identifier: GPL-2.0-or-later
pragma solidity ^0.8.4;

import {IAlgebraPool} from '@cryptoalgebra/integral-core/contracts/interfaces/IAlgebraPool.sol';
import {PluginConfig} from '../types/PluginConfig.sol';

/// @title Contains logic and constants for easy interacting with the Algebra Integral pool
/// @dev Simplifies interactions with Algebra Integral pool
library PoolInteraction {
    function getSqrtPrice(IAlgebraPool pool) internal view returns (uint160 price) {
        (price, , , ) = getPoolState(pool);
    }

    function getCurrentTick(IAlgebraPool pool) internal view returns (int24 currentTick) {
        (, currentTick, , ) = getPoolState(pool);
    }

    function getFee(IAlgebraPool pool) internal view returns (uint16 fee) {
        (, , fee, ) = getPoolState(pool);
    }

    function getPluginConfig(IAlgebraPool pool) internal view returns (PluginConfig pluginConfig) {
        uint8 _config;
        (, , , _config) = getPoolState(pool);
        return PluginConfig.wrap(_config);
    }

    function getPoolState(
        IAlgebraPool pool
    ) internal view returns (uint160 price, int24 currentTick, uint16 fee, uint8 pluginConfig) {
        (price, currentTick, fee, pluginConfig, , ) = IAlgebraPool(pool).globalState();
    }

    function getPreviousAndNextTicks(IAlgebraPool pool) internal view returns (int24 previousTick, int24 nextTick) {
        previousTick = IAlgebraPool(pool).prevTickGlobal();
        nextTick = IAlgebraPool(pool).nextTickGlobal();
    }

    function changePluginConfigIfNeeded(IAlgebraPool pool, PluginConfig newConfig) internal {
        if (getPluginConfig(pool) != newConfig) {
            IAlgebraPool(pool).setPluginConfig(newConfig.unwrap());
        }
    }

    function changeFeeIfNeeded(IAlgebraPool pool, uint16 newFee) internal {
        if (getFee(pool) != newFee) {
            IAlgebraPool(pool).setFee(newFee);
        }
    }
}
