// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "../../framework/MessageApp.sol";

interface IPeggedToken {
    function mint(address _to, uint256 _amount) external;

    function burnFrom(address _from, uint256 _amount) external;
}

contract PegBridge is MessageApp, Ownable {
    using SafeERC20 for IERC20;

    mapping(bytes32 => bool) public records;
    mapping(address => uint256) public supplies;
    mapping(address => address) public vaultToPegTokens; // vault -> peg
    mapping(address => address) public pegToVaultTokens; // peg -> vault

    mapping(address => uint256) public minBurn;

    uint64 public vaultChain;
    address public vaultAddress;

    event Mint(bytes32 mintId, address account, address token, uint256 amount, bytes32 refId, address depositor);

    event Burn(
        bytes32 burnId,
        address burnAccount,
        address token,
        uint256 amount,
        address withdrawAccount,
        uint64 nonce,
        bytes32 messageId
    );

    event BridgeTokenAdded(address vaultToken, address pegToken);
    event BridgeTokenDeleted(address vaultToken, address pegToken);
    event MinBurnUpdated(address token, uint256 amount);

    constructor(IMessageBridge _messageBridge) MessageApp(_messageBridge) {}

    function _handleMessage(
        uint64 _srcChainId,
        address _sender,
        bytes calldata _message,
        address // execution
    ) internal override {
        require(_srcChainId == vaultChain, "not from vault chain");
        require(_sender == vaultAddress, "sender is not token vault");
        _mint(_message);
    }

    function _mint(bytes calldata _message) private {
        (address vaultToken, uint256 amount, address mintAccount, address depositor, bytes32 depositId) = abi.decode(
            (_message),
            (address, uint256, address, address, bytes32)
        );
        address pegToken = vaultToPegTokens[vaultToken];
        require(pegToken != address(0), "no peg token");
        bytes32 mintId = keccak256(
            abi.encodePacked(pegToken, amount, mintAccount, depositor, depositId, address(this))
        );
        require(records[mintId] == false, "record exists");
        records[mintId] = true;
        IPeggedToken(pegToken).mint(mintAccount, amount);
        supplies[pegToken] += amount;
        emit Mint(mintId, mintAccount, pegToken, amount, depositId, depositor);
    }

    function burn(
        address _token,
        uint256 _amount,
        uint64 _toChainId,
        address _withdrawAccount,
        uint64 _nonce
    ) external returns (bytes32) {
        bytes32 burnId = _burn(_token, _amount, _toChainId, _withdrawAccount, _nonce);
        IPeggedToken(_token).burnFrom(msg.sender, _amount);
        return burnId;
    }

    function _burn(
        address _token,
        uint256 _amount,
        uint64 _toChainId,
        address _withdrawAccount,
        uint64 _nonce
    ) private returns (bytes32) {
        require(_amount > minBurn[_token], "amount too small");
        address vaultToken = pegToVaultTokens[_token];
        require(vaultToken != address(0), "no vault token");
        supplies[_token] -= _amount;
        bytes32 burnId = keccak256(
            abi.encodePacked(
                msg.sender,
                _token,
                _amount,
                _toChainId,
                _withdrawAccount,
                _nonce,
                uint64(block.chainid),
                address(this)
            )
        );
        require(records[burnId] == false, "record exists");
        records[burnId] = true;
        bytes memory message = abi.encode(vaultToken, _amount, _withdrawAccount, msg.sender, burnId);
        bytes32 messageId = _sendMessage(_toChainId, vaultAddress, message);
        emit Burn(burnId, msg.sender, _token, _amount, _withdrawAccount, _nonce, messageId);
        return burnId;
    }

    function setBridgeTokens(address[] calldata _vaultTokens, address[] calldata _pegTokens) external onlyOwner {
        require(_vaultTokens.length == _pegTokens.length, "length mismatch");
        for (uint256 i = 0; i < _vaultTokens.length; i++) {
            vaultToPegTokens[_vaultTokens[i]] = _pegTokens[i];
            pegToVaultTokens[_pegTokens[i]] = _vaultTokens[i];
            emit BridgeTokenAdded(_vaultTokens[i], _pegTokens[i]);
        }
    }

    function deletePegTokens(address[] calldata _pegTokens) external onlyOwner {
        for (uint256 i = 0; i < _pegTokens.length; i++) {
            address pegToken = _pegTokens[i];
            address vaultToken = pegToVaultTokens[pegToken];
            delete vaultToPegTokens[vaultToken];
            delete pegToVaultTokens[pegToken];
            emit BridgeTokenDeleted(vaultToken, pegToken);
        }
    }

    function deleteVaultTokens(address[] calldata _vaultTokens) external onlyOwner {
        for (uint256 i = 0; i < _vaultTokens.length; i++) {
            address vaultToken = _vaultTokens[i];
            address pegToken = vaultToPegTokens[vaultToken];
            delete vaultToPegTokens[vaultToken];
            delete pegToVaultTokens[pegToken];
            emit BridgeTokenDeleted(vaultToken, pegToken);
        }
    }

    function setMinBurn(address[] calldata _tokens, uint256[] calldata _amounts) external onlyOwner {
        require(_tokens.length == _amounts.length, "length mismatch");
        for (uint256 i = 0; i < _tokens.length; i++) {
            minBurn[_tokens[i]] = _amounts[i];
            emit MinBurnUpdated(_tokens[i], _amounts[i]);
        }
    }

    function setTokenVault(uint64 _vaultChain, address _vaultAddress) external onlyOwner {
        vaultChain = _vaultChain;
        vaultAddress = _vaultAddress;
    }
}
