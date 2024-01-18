// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./Initializable.sol";

contract System is Ownable, Initializable {
    uint32 public constant CODE_OK = 0;
    uint32 public constant ERROR_FAIL_DECODE = 100;

    uint8 public constant STAKING_CHANNEL_ID = 0x08;

    address public bscValidatorSet;
    address public tmLightClient;
    address public crossChain;

    uint16 public bscChainID;
    address public relayer;

    function init(
        uint16 _bscChainID,
        address _relayer,
        address _bscValidatorSet,
        address _tmLightClient,
        address _crossChain
    ) external onlyUninitialized onlyOwner {
        bscChainID = _bscChainID;
        relayer = _relayer;
        bscValidatorSet = _bscValidatorSet;
        tmLightClient = _tmLightClient;
        crossChain = _crossChain;

        _initialized = true;
    }

    function setRelayer(address _relayer) external onlyOwner {
        relayer = _relayer;
    }
}
