// SPDX-License-Identifier: GPL-3.0-only

pragma solidity ^0.8.18;

import "@openzeppelin/contracts/security/Pausable.sol";
import "./Ownable.sol";

abstract contract Pauser is Ownable, Pausable {
    mapping(address => bool) public pausers;
    address[] public pauserList;

    event PauserAdded(address account);
    event PauserRemoved(address account);

    constructor() {
        _addPauser(msg.sender);
    }

    modifier onlyPauser() {
        require(isPauser(msg.sender), "Caller is not pauser");
        _;
    }

    function pause() public virtual onlyPauser {
        _pause();
    }

    function unpause() public virtual onlyPauser {
        _unpause();
    }

    function isPauser(address account) public view returns (bool) {
        return pausers[account];
    }

    function addPauser(address account) public onlyOwner {
        _addPauser(account);
    }

    function addPausers(address[] memory accounts) public onlyOwner {
        for (uint256 i = 0; i < accounts.length; i++) {
            _addPauser(accounts[i]);
        }
    }

    function removePauser(address account) public onlyOwner {
        _removePauser(account);
    }

    function removePausers(address[] memory accounts) public onlyOwner {
        for (uint256 i = 0; i < accounts.length; i++) {
            _removePauser(accounts[i]);
        }
    }

    function renouncePauser() public {
        _removePauser(msg.sender);
    }

    function numPausers() public view returns (uint256) {
        return pauserList.length;
    }

    function _addPauser(address account) private {
        require(!isPauser(account), "Account is already pauser");
        pauserList.push(account);
        pausers[account] = true;
        emit PauserAdded(account);
    }

    function _removePauser(address account) private {
        require(isPauser(account), "Account is not pauser");
        uint256 lastIndex = pauserList.length - 1;
        for (uint256 i = 0; i < pauserList.length; i++) {
            if (pauserList[i] == account) {
                if (i < lastIndex) {
                    pauserList[i] = pauserList[lastIndex];
                }
                pauserList.pop();
                pausers[account] = false;
                emit PauserRemoved(account);
                return;
            }
        }
        revert("pauser not found"); // this should never happen
    }
}
