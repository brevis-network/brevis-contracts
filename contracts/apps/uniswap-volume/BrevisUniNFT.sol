// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract BrevisUniNFT is ERC721URIStorage, Ownable {
    address public minter;
    uint256 public count;
    string public baseTokenURI;

    modifier onlyMinter() {
        require(msg.sender == minter, "caller is not minter");
        _;
    }

    constructor(string memory name_, string memory symbol_, address _minter) ERC721(name_, symbol_) {
        minter = _minter;
    }

    function mint(address to) external onlyMinter {
        _mint(to, count);
        count++;
    }

    function setMinter(address _minter) external onlyOwner {
        minter = _minter;
    }

    function setBaseURI(string memory _uri) external onlyOwner {
        baseTokenURI = _uri;
    }

    function _baseURI() internal view override returns (string memory) {
        return baseTokenURI;
    }
}
