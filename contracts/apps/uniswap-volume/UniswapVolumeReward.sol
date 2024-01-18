// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import {BytesLib} from "solidity-bytes-utils/contracts/BytesLib.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

import "./libraries/Commands.sol";
import "./libraries/Path.sol";
import "../../verifiers/interfaces/ITxVerifier.sol";

interface IBrevisUniNFT {
    function mint(address to) external;
}

contract UniswapVolumeReward is Ownable {
    ITxVerifier public txVerifier;

    mapping(uint64 => address) public universalRouter; // chainId -> router address
    mapping(uint64 => address) public weth; // chainId -> WETH address
    mapping(uint64 => address) public usdc; // chainId -> USDC address
    mapping(uint64 => uint8) usdcDecimal; // chainId -> USDC decimal

    bytes4 private executeSelector = bytes4(keccak256(bytes("execute(bytes,bytes[],uint256)")));
    bytes4 private executeNoDealineSelector = bytes4(keccak256(bytes("execute(bytes,bytes[])")));

    enum TierName {
        Null,
        Stone,
        Bronze,
        Silver,
        Gold,
        Platinum,
        Diamond
    }
    mapping(TierName => address) public tierNFTs;
    mapping(address => TierName) public userTier; // user -> tier

    event VerifiedSwap(uint64 chainId, address trader, uint64 timestamp, uint256 usdcAmount, TierName tier);

    constructor(ITxVerifier _txVerifier) {
        txVerifier = _txVerifier;
    }

    function submitUniswapTxProof(
        bytes calldata _tx,
        bytes calldata _proof,
        bytes calldata _auxiBlkVerifyInfo
    ) external {
        ITxVerifier.TxInfo memory txInfo = txVerifier.verifyTxAndLog(_tx, _proof, _auxiBlkVerifyInfo);
        require(txInfo.to == universalRouter[txInfo.chainId], "invalid to address");
        require(userTier[txInfo.from] == TierName.Null, "swap already proved for this user");

        (uint256 amount, TierName tier) = usdcSwapAmount(txInfo.chainId, txInfo.data);
        userTier[txInfo.from] = tier;
        IBrevisUniNFT(tierNFTs[tier]).mint(txInfo.from);

        emit VerifiedSwap(txInfo.chainId, txInfo.from, txInfo.blkTime, amount, tier);
    }

    function setUniversalRouter(uint64 _chainId, address _router) external onlyOwner {
        universalRouter[_chainId] = _router;
    }

    function setWETH(uint64 _chainId, address _weth) external onlyOwner {
        weth[_chainId] = _weth;
    }

    function setUSDC(uint64 _chainId, address _usdc, uint8 _decimal) external onlyOwner {
        usdc[_chainId] = _usdc;
        usdcDecimal[_chainId] = _decimal;
    }

    function setTierNFTs(TierName[] calldata _names, address[] calldata _nfts) external onlyOwner {
        require(_names.length == _nfts.length, "length mismatch");
        for (uint256 i = 0; i < _names.length; i++) {
            tierNFTs[_names[i]] = _nfts[i];
        }
    }

    function setTxVerifier(ITxVerifier _txVerifier) external onlyOwner {
        txVerifier = _txVerifier;
    }

    function usdcSwapAmount(uint64 _chainId, bytes memory _data) public view returns (uint256 amount, TierName tier) {
        bytes4 method;
        assembly {
            method := mload(add(_data, 32))
        }
        require(method == executeSelector || method == executeNoDealineSelector, "wrong method");
        bytes memory argdata = BytesLib.slice(_data, 4, _data.length - 4);
        (bytes memory commands, bytes[] memory inputs, ) = abi.decode(argdata, (bytes, bytes[], uint256));
        // assume tx succeeded, so no need to check deadline or command/input length match

        uint256 command = uint8(commands[0] & Commands.COMMAND_TYPE_MASK);
        bytes memory input;
        if (command == Commands.WRAP_ETH || command == Commands.PERMIT2_PERMIT) {
            command = uint8(commands[1] & Commands.COMMAND_TYPE_MASK);
            input = inputs[1];
        } else {
            input = inputs[0];
        }
        require(command == Commands.V3_SWAP_EXACT_IN || command == Commands.V3_SWAP_EXACT_OUT, "unsupported command");
        (, uint256 amountA, uint256 amountB, bytes memory path, ) = abi.decode(
            input,
            (address, uint256, uint256, bytes, bool)
        );
        (address tokenA, address tokenB, ) = Path.decodeFirstPool(path);
        if (tokenA == weth[_chainId]) {
            require(tokenB == usdc[_chainId], "unsupported pair");
            amount = amountB;
        } else if (tokenA == usdc[_chainId]) {
            require(tokenB == weth[_chainId], "unsupported pair");
            amount = amountA;
        } else {
            revert("unsupported pair");
        }
        require(amount > 0, "zero usdc amount");
        tier = getAmountTier(_chainId, amount);
    }

    function getAmountTier(uint64 _chainId, uint256 _amount) private view returns (TierName tier) {
        uint256 decimal = 10 ** usdcDecimal[_chainId];
        if (_amount >= 1000000 * decimal) {
            return TierName.Diamond;
        } else if (_amount >= 100000 * decimal) {
            return TierName.Platinum;
        } else if (_amount >= 10000 * decimal) {
            return TierName.Gold;
        } else if (_amount >= 1000 * decimal) {
            return TierName.Silver;
        } else if (_amount >= 100 * decimal) {
            return TierName.Bronze;
        }
        return TierName.Stone;
    }
}
