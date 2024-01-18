// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import {Hooks} from "@uniswap/v4-core/contracts/libraries/Hooks.sol";
import {IHookFeeManager} from "@uniswap/v4-core/contracts/interfaces/IHookFeeManager.sol";
import {IDynamicFeeManager} from "@uniswap/v4-core/contracts/interfaces/IDynamicFeeManager.sol";
import {FeeLibrary} from "@uniswap/v4-core/contracts/libraries/FeeLibrary.sol";
import {BalanceDelta} from "@uniswap/v4-core/contracts/types/BalanceDelta.sol";
import {PoolId, PoolIdLibrary} from "@uniswap/v4-core/contracts/types/PoolId.sol";
import {PoolKey} from "@uniswap/v4-core/contracts/types/PoolKey.sol";
import {IPoolManager} from "@uniswap/v4-core/contracts/interfaces/IPoolManager.sol";
import {BaseHook} from "./BaseHook.sol";
import {BaseFactory} from "./BaseFactory.sol";
import {UniswapSumVolume} from "../apps/uniswap-sum/UniswapSumVolume.sol";
import {IUniswapSumVolume} from "../apps/uniswap-sum/IUniswapSumVolume.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract TierHook is BaseHook, IHookFeeManager, IDynamicFeeManager {
    using FeeLibrary for uint24;
    using PoolIdLibrary for PoolKey;

    address public sumVolumeAddress;

    event UpdateSumVolumeAddress(address newAddress);
    event Swap(PoolId indexed id, address indexed sender, int128 amount0, int128 amount1);

    uint256 internal constant TIER_NONE = 0;
    uint256 internal constant TIER_SILVER = 1;
    uint256 internal constant TIER_GOLD = 2;

    constructor(IPoolManager _poolManager) BaseHook(_poolManager) {}

    function getHooksCalls() public pure override returns (Hooks.Calls memory) {
        return
            Hooks.Calls({
                beforeInitialize: true,
                afterInitialize: false,
                beforeModifyPosition: false,
                afterModifyPosition: false,
                beforeSwap: false,
                afterSwap: true,
                beforeDonate: false,
                afterDonate: false
            });
    }

    /// @notice The interface for setting a fee on swap or fee on withdraw to the hook
    /// @dev This callback is only made if the Fee.HOOK_SWAP_FEE_FLAG or Fee.HOOK_WITHDRAW_FEE_FLAG in set in the pool's key.fee.
    function getHookFees(PoolKey calldata) external pure returns (uint24 fee) {
        // Swap fee is upper bits.
        // 20% fee as 85 = hex55 which is 5 in both directions. 1/5 = 20%
        // Withdraw fee is lower bits
        // 33% fee as 51 = hex33 which is 3 in both directions. 1/3 = 33%
        fee = 0x5533;
    }

    function getHookWithdrawFee(PoolKey calldata key) external view returns (uint8 fee) {}

    function beforeInitialize(
        address,
        PoolKey calldata key,
        uint160,
        bytes calldata
    ) external pure override returns (bytes4) {
        if (!key.fee.isDynamicFee()) revert MustUseDynamicFee();
        return BaseHook.beforeInitialize.selector;
    }

    function afterSwap(
        address,
        PoolKey calldata key,
        IPoolManager.SwapParams calldata,
        BalanceDelta delta,
        bytes calldata
    ) external override returns (bytes4 selector) {
        emit Swap(key.toId(), tx.origin, delta.amount0(), delta.amount1());

        selector = BaseHook.afterSwap.selector;
    }

    function updateSumVolumeAddress(address _sumVolumeAddress) external onlyOwner {
        sumVolumeAddress = _sumVolumeAddress;
        emit UpdateSumVolumeAddress(_sumVolumeAddress);
    }

    function senderTier(address _sender) internal view returns (uint256) {
        uint256 existingVolume = (IUniswapSumVolume)(sumVolumeAddress).getAttestedSwapSumVolume(_sender); //  uint256(sumVolume.volumes[address(0)]);

        // existingVolume uses 6 as decimals. 1000000000 means $1000
        if (existingVolume > 1000000000) {
            return TIER_GOLD;
        } else if (existingVolume > 0) {
            return TIER_SILVER;
        }
        return TIER_NONE;
    }

    function getFeeBySwapper(address swapper) external view returns (uint24) {
        return calcFee(swapper);
    }

    // invoke by PoolManager contract
    function getFee(
        address,
        PoolKey calldata,
        IPoolManager.SwapParams calldata,
        bytes calldata
    ) external view returns (uint24) {
        return calcFee(tx.origin);
    }

    function calcFee(address user) internal view returns (uint24) {
        uint256 tier = senderTier(user);
        uint24 fee = 10000;
        if (tier == TIER_GOLD) {
            fee = fee / 2;
        } else if (tier == TIER_SILVER) {
            fee = (fee * 4) / 5;
        }
        return fee;
    }
}

contract TierFactory is BaseFactory, Ownable {
    constructor() BaseFactory(address(uint160(Hooks.BEFORE_INITIALIZE_FLAG | Hooks.AFTER_SWAP_FLAG))) {}

    function deploy(IPoolManager poolManager, bytes32 salt) public override returns (address) {
        return address(new TierHook{salt: salt}(poolManager));
    }

    function _hashBytecode(IPoolManager poolManager) internal pure override returns (bytes32 bytecodeHash) {
        bytecodeHash = keccak256(abi.encodePacked(type(TierHook).creationCode, abi.encode(poolManager)));
    }

    function updateHookSumVolumeAddress(TierHook _tierHook, address _sumVolumeAddress) public onlyOwner {
        _tierHook.updateSumVolumeAddress(_sumVolumeAddress);
    }
}
