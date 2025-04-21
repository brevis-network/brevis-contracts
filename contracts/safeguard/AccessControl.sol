// SPDX-License-Identifier: GPL-3.0-only

pragma solidity ^0.8.18;

import "./Ownable.sol";

abstract contract AccessControl is Ownable {
    mapping(bytes32 role => mapping(address account => bool)) public roles;
    mapping(bytes32 role => address[] accounts) public roleAccounts;

    event RoleGranted(bytes32 role, address account);
    event RoleRevoked(bytes32 role, address account);

    modifier onlyRole(bytes32 role) {
        require(hasRole(role, msg.sender), "unauthorized role");
        _;
    }

    function hasRole(bytes32 role, address account) public view returns (bool) {
        return roles[role][account];
    }

    function numRoleAccounts(bytes32 role) public view returns (uint256) {
        return roleAccounts[role].length;
    }

    function getRoleAccounts(bytes32 role) public view returns (address[] memory accounts) {
        return roleAccounts[role];
    }

    function grantRole(bytes32 role, address account) public onlyOwner {
        _grantRole(role, account);
    }

    function grantRoles(bytes32 role, address[] memory accounts) public onlyOwner {
        for (uint256 i = 0; i < accounts.length; i++) {
            _grantRole(role, accounts[i]);
        }
    }

    function revokeRole(bytes32 role, address account) public onlyOwner {
        _revokeRole(role, account);
    }

    function revokeRoles(bytes32 role, address[] memory accounts) public onlyOwner {
        for (uint256 i = 0; i < accounts.length; i++) {
            _revokeRole(role, accounts[i]);
        }
    }

    function renounceRole(bytes32 role) public {
        _revokeRole(role, msg.sender);
    }

    // -------------- internal functions --------------

    function _grantRole(bytes32 role, address account) internal {
        require(!hasRole(role, account), "already has role");
        roleAccounts[role].push(account);
        roles[role][account] = true;
        emit RoleGranted(role, account);
    }

    function _revokeRole(bytes32 role, address account) internal {
        require(hasRole(role, account), "not has role");
        address[] storage accountList = roleAccounts[role];
        uint256 lastIndex = accountList.length - 1;
        for (uint256 i = 0; i < accountList.length; i++) {
            if (accountList[i] == account) {
                if (i < lastIndex) {
                    accountList[i] = accountList[lastIndex];
                }
                accountList.pop();
                roles[role][account] = false;
                emit RoleRevoked(role, account);
                return;
            }
        }
        revert("role account not found"); // this should never happen
    }
}
