// SPDX-License-Identifier: GPL-3.0-only

pragma solidity >=0.8.18;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @title Allows the owner to set fee collector and allows fee collectors to collect fees
 */
contract FeeVault is Ownable {
    using SafeERC20 for IERC20;

    address public feeCollector;

    event FeeCollectorUpdated(address from, address to);

    constructor(address _feeCollector) {
        feeCollector = _feeCollector;
    }

    modifier onlyFeeCollector() {
        require(msg.sender == feeCollector, "not fee collector");
        _;
    }

    function collectFee(uint256 _amount, address _to) external onlyFeeCollector {
        (bool sent, ) = _to.call{value: _amount, gas: 50000}("");
        require(sent, "send native failed");
    }

    function setFeeCollector(address _feeCollector) external onlyOwner {
        address oldFeeCollector = feeCollector;
        feeCollector = _feeCollector;
        emit FeeCollectorUpdated(oldFeeCollector, _feeCollector);
    }

    receive() external payable {}
}
