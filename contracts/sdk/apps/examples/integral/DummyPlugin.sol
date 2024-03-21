// SPDX-License-Identifier: GPL-2.0-or-later
pragma solidity ^0.8.4;

import {AlgebraPlugin, IAlgebraPlugin} from './base/AlgebraPlugin.sol';
import {IAlgebraPool} from '@cryptoalgebra/integral-core/contracts/interfaces/IAlgebraPool.sol';
import {PoolInteraction} from './libraries/PoolInteraction.sol';
import {PluginConfig, Plugins} from './types/PluginConfig.sol';

contract DummyPlugin is AlgebraPlugin {
    error onlyPoolAllowed();

    PluginConfig private constant _defaultPluginConfig = PluginConfig.wrap(0); // does nothing

    /// @notice the Algebra Integral pool
    IAlgebraPool public immutable pool;

    modifier onlyPool() {
        _checkOnlyPool();
        _;
    }

    constructor(address _pool) {
        pool = IAlgebraPool(_pool);
    }

    function defaultPluginConfig() external pure override returns (uint8 pluginConfig) {
        return _defaultPluginConfig.unwrap();
    }

    /// @inheritdoc IAlgebraPlugin
    function beforeInitialize(address sender, uint160 sqrtPriceX96) external onlyPool returns (bytes4) {
        sender; // suppress warning
        sqrtPriceX96; //suppress warning

        PoolInteraction.changePluginConfigIfNeeded(pool, _defaultPluginConfig);
        return IAlgebraPlugin.beforeInitialize.selector;
    }

    function _checkOnlyPool() internal view {
        if (msg.sender != address(pool)) revert onlyPoolAllowed();
    }
}
