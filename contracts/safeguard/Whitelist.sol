// SPDX-License-Identifier: GPL-3.0-only

pragma solidity ^0.8.18;

import "./Ownable.sol";

abstract contract Whitelist is Ownable {
    mapping(address => bool) public whitelist;
    bool public whitelistEnabled;

    event WhitelistedAdded(address account);
    event WhitelistedRemoved(address account);

    modifier onlyWhitelisted() {
        if (whitelistEnabled) {
            require(isWhitelisted(msg.sender), "Caller is not whitelisted");
        }
        _;
    }

    /**
     * @notice Set whitelistEnabled
     */
    function setWhitelistEnabled(bool _whitelistEnabled) public onlyOwner {
        whitelistEnabled = _whitelistEnabled;
    }

    /**
     * @notice Add an account to whitelist
     */
    function addWhitelistedAccount(address _account) public onlyOwner {
        require(!isWhitelisted(_account), "Already whitelisted");
        whitelist[_account] = true;
        emit WhitelistedAdded(_account);
    }

    /**
     * @notice Add multiple accounts to whitelist
     */
    function addWhitelistedAccounts(address[] memory _accounts) public onlyOwner {
        for (uint256 i = 0; i < _accounts.length; i++) {
            addWhitelistedAccount(_accounts[i]);
        }
    }

    /**
     * @notice Remove an account from whitelist
     */
    function removeWhitelistedAccount(address _account) public onlyOwner {
        require(isWhitelisted(_account), "Not whitelisted");
        whitelist[_account] = false;
        emit WhitelistedRemoved(_account);
    }

    /**
     * @notice Remove multiple accounts from whitelist
     */
    function removeWhitelistedAccounts(address[] memory _accounts) public onlyOwner {
        for (uint256 i = 0; i < _accounts.length; i++) {
            removeWhitelistedAccount(_accounts[i]);
        }
    }

    /**
     * @return is account whitelisted
     */
    function isWhitelisted(address account) public view returns (bool) {
        return whitelist[account];
    }
}
