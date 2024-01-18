// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

interface ISlotValueVerifier {
    struct SlotInfo {
        uint64 chainId;
        bytes32 addrHash;
        bytes32 blkHash;
        bytes32 slotKeyHash;
        bytes32 slotValue;
        uint32 blkNum;
    }

    /**
     * @notice Called by dApp contracts to verify a slot value
     * @param chainId The source chain ID for which the proof data was generated
     * @param proofData Groth16 proof data, with the appended public inputs.
     * @param blkVerifyInfo Data passed to the BlockSyncer to validate the block in the source chain.
     */
    function verifySlotValue(
        uint64 chainId,
        bytes calldata proofData,
        bytes calldata blkVerifyInfo
    ) external view returns (SlotInfo memory slotInfo);
}
