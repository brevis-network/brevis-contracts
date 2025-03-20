// SPDX-License-Identifier: GPL-3.0-only

pragma solidity ^0.8.18;

import "./proxies/CommonOwnerProxy.sol";
import "./proxies/UpgradeableOwnerProxy.sol";
import "./proxies/BrevisProofOwnerProxy.sol";
import "./proxies/BrevisRequestOwnerProxy.sol";
import "./proxies/SmtOwnerProxy.sol";
import "./proxies/LightClientOwnerProxy.sol";

contract GovernedOwnerProxy is
    CommonOwnerProxy,
    UpgradeableOwnerProxy,
    BrevisProofOwnerProxy,
    BrevisRequestOwnerProxy,
    SmtOwnerProxy,
    LightClientOwnerProxy
{
    constructor(address _initializer) OwnerProxyBase(_initializer) {}
}
