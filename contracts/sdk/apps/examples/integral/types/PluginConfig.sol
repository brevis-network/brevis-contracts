// SPDX-License-Identifier: GPL-2.0-or-later
pragma solidity ^0.8.4;

import {Plugins} from '@cryptoalgebra/integral-core/contracts/libraries/Plugins.sol';

type PluginConfig is uint8;

using PluginConfigActions for PluginConfig global;
using {eq as ==, neq as !=} for PluginConfig global;

function eq(PluginConfig a, PluginConfig b) pure returns (bool) {
    return PluginConfig.unwrap(a) == PluginConfig.unwrap(b);
}

function neq(PluginConfig a, PluginConfig b) pure returns (bool) {
    return PluginConfig.unwrap(a) != PluginConfig.unwrap(b);
}

/// @title Contains logic and constants for easy interacting with the plugin config in pool
/// @dev Simplifies interactions with plugin configuration
library PluginConfigActions {
    function switchFlag(PluginConfig self, uint256 flag, bool newValue) internal pure returns (PluginConfig) {
        uint8 config = PluginConfig.unwrap(self);
        if (newValue) {
            config |= uint8(flag);
        } else {
            assembly {
                config := and(config, not(flag))
            }
        }
        return PluginConfig.wrap(config);
    }

    function hasFlag(PluginConfig self, uint256 flag) internal pure returns (bool res) {
        return Plugins.hasFlag(PluginConfig.unwrap(self), flag);
    }

    function unwrap(PluginConfig self) internal pure returns (uint8) {
        return PluginConfig.unwrap(self);
    }
}
