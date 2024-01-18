// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../../framework/BrevisApp.sol";
import "../../../lib/Lib.sol";
import "../../../interface/IBrevisProof.sol";

// Multiple txs: Proof of USDC transfer volume
// Use Brevis to prove address #1’s accumulative USDC transfer volume to address #2.

contract DemoTokenTransferVolume is BrevisApp {
    address public token;
    bytes4 public transferSelector;

    // (sender, receiver) -> volume
    mapping(address => mapping(address => uint256)) public volumes;

    constructor(IBrevisProof _brevisProof, address _token, bytes4 _transferSelector) BrevisApp(_brevisProof) {
        token = _token;
        transferSelector = _transferSelector;
    }

    function submitUserTxs(
        bytes32 _proofRequestId,
        uint64 _chainId,
        bytes32 _smtRoot,
        Brevis.TransactionInfo[] calldata _info,
        bytes[] calldata _txRaw
    ) external {
        require(_info.length == _txRaw.length, "length not match");

        Brevis.ExtractInfos memory info;
        info.smtRoot = _smtRoot;
        info.txs = _info;

        validateRequest(_proofRequestId, _chainId, info);

        address from;
        address receiver;
        uint256 volume;
        for (uint256 i = 0; i < _info.length; i++) {
            bytes memory leafRlp = bytes.concat(_info[i].leafRlpPrefix, _txRaw[i]);
            bytes32 leafHash = keccak256(leafRlp);
            require(leafHash == _info[i].leafHash, "leafHash not match");

            Tx.TxInfo memory txInfo = Tx.decodeTx(_txRaw[i]);
            require(txInfo.to == token, "not a token transfer");
            (bytes4 sig, address recv, uint256 amount) = abiDecodeTransfer(txInfo.data);
            require(sig == transferSelector, "not transfer selector");
            if (i > 0) {
                require(from == txInfo.from, "not same sender");
                require(recv == receiver, "not same receiver");
            } else {
                receiver = recv;
                from = txInfo.from;
            }
            volume = volume + amount;
        }
        volumes[from][receiver] = volume;
    }

    function abiDecodeTransfer(bytes memory _data) private pure returns (bytes4 sig, address receiver, uint256 amount) {
        assembly {
            sig := mload(add(_data, 32))
            receiver := mload(add(_data, 36))
            amount := mload(add(_data, 68))
        }
    }
}
