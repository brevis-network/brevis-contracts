// SPDX-License-Identifier: GPL-3.0-only

pragma solidity >=0.8.0;

interface ICommonOwner {
    function transferOwnership(address _newOwner) external;

    function addPauser(address _account) external;

    function addPausers(address[] memory accounts) external;

    function removePauser(address _account) external;

    function removePausers(address[] memory _accounts) external;

    function addProvers(address[] memory _accounts) external;

    function removeProvers(address[] memory _accounts) external;

    function addGovernor(address _account) external;

    function addGovernors(address[] memory _accounts) external;

    function removeGovernor(address _account) external;

    function removeGovernors(address[] memory _accounts) external;
}
