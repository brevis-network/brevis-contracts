// SPDX-License-Identifier: GPL-3.0-only

pragma solidity >=0.8.18;

import "../../safeguard/Ownable.sol";

/**
 * @title Allows the owner to set fee collector and allows fee collectors to collect fees
 */
contract FeeVault is Ownable {
    address public feeCollector;

    event FeeCollected(uint256 amount, address receiver);
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
        emit FeeCollected(_amount, _to);
    }

    function setFeeCollector(address _feeCollector) external onlyOwner {
        address oldFeeCollector = feeCollector;
        feeCollector = _feeCollector;
        emit FeeCollectorUpdated(oldFeeCollector, _feeCollector);
    }

    receive() external payable {}
}
