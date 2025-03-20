// SPDX-License-Identifier: GPL-3.0-only

pragma solidity >=0.8.0;

interface ILightClientOwner {
    function setLightClient(address _lightClient) external;

    function updateForkVersion(uint64 epoch, bytes4 forkVersion) external;

    function processLightClientForceUpdate() external;
}
