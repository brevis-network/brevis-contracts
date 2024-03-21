// SPDX-License-Identifier: GPL-2.0-or-later
pragma solidity ^0.8.4;

import '@cryptoalgebra/integral-core/contracts/interfaces/plugin/IAlgebraPlugin.sol';

abstract contract AlgebraPlugin is IAlgebraPlugin {
    /// @inheritdoc IAlgebraPlugin
    function afterInitialize(address, uint160, int24) external virtual override returns (bytes4) {
        revert('Not implemented');
    }

    /// @inheritdoc IAlgebraPlugin
    function beforeModifyPosition(
        address,
        address,
        int24,
        int24,
        int128,
        bytes calldata
    ) external virtual returns (bytes4) {
        revert('Not implemented');
    }

    /// @inheritdoc IAlgebraPlugin
    function afterModifyPosition(
        address,
        address,
        int24,
        int24,
        int128,
        uint256,
        uint256,
        bytes calldata
    ) external virtual returns (bytes4) {
        revert('Not implemented');
    }

    /// @inheritdoc IAlgebraPlugin
    function beforeSwap(
        address,
        address,
        bool,
        int256,
        uint160,
        bool,
        bytes calldata
    ) external virtual returns (bytes4) {
        revert('Not implemented');
    }

    /// @inheritdoc IAlgebraPlugin
    function afterSwap(
        address,
        address,
        bool,
        int256,
        uint160,
        int256,
        int256,
        bytes calldata
    ) external virtual returns (bytes4) {
        revert('Not implemented');
    }

    /// @inheritdoc IAlgebraPlugin
    function beforeFlash(address, address, uint256, uint256, bytes calldata) external virtual returns (bytes4) {
        revert('Not implemented');
    }

    /// @inheritdoc IAlgebraPlugin
    function afterFlash(
        address,
        address,
        uint256,
        uint256,
        uint256,
        uint256,
        bytes calldata
    ) external virtual returns (bytes4) {
        revert('Not implemented');
    }
}
