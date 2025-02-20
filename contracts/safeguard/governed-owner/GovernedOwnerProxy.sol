// SPDX-License-Identifier: GPL-3.0-only

pragma solidity ^0.8.18;

import "./proxies/CommonOwnerProxy.sol";
import "./proxies/UpgradeableOwnerProxy.sol";

contract GovernedOwnerProxy is CommonOwnerProxy, UpgradeableOwnerProxy {
    constructor(address _initializer) OwnerProxyBase(_initializer) {}
}
