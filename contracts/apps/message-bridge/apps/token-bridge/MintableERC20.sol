// SPDX-License-Identifier: GPL-3.0-only

pragma solidity ^0.8.18;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";

/**
 * @title A mintable {ERC20} token.
 */
contract MintableERC20 is ERC20Burnable, Ownable {
    uint8 private _decimals;
    address private _minter;

    /**
     * @dev Constructor that gives msg.sender an initial supply of tokens.
     */
    constructor(
        string memory name_,
        string memory symbol_,
        uint8 decimals_,
        uint256 initialSupply_
    ) ERC20(name_, symbol_) {
        _decimals = decimals_;
        _minter = msg.sender;
        _mint(msg.sender, initialSupply_);
    }

    /**
     * @dev Returns the address of the current owner.
     */
    function minter() public view virtual returns (address) {
        return _minter;
    }

    /**
     * @dev Throws if called by any account other than the minter.
     */
    modifier onlyMinter() {
        require(minter() == msg.sender, "Mintable: caller is not the minter");
        _;
    }

    /**
     * @dev Creates `amount` new tokens for `to`.
     */
    function mint(address to, uint256 amount) public onlyMinter {
        _mint(to, amount);
    }

    function decimals() public view override returns (uint8) {
        return _decimals;
    }

    function setMinter(address minter_) external onlyOwner {
        _minter = minter_;
    }
}
