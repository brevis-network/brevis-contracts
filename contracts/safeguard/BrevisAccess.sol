// SPDX-License-Identifier: GPL-3.0-only

pragma solidity ^0.8.18;

import "./Pauser.sol";

// prover and pauser access control using a single map lookup
abstract contract BrevisAccess is Pauser {
    enum ProverState {
        Null,
        Active,
        Paused
    }
    mapping(address => ProverState) public proverStates;
    address[] public provers;

    event ProverAdded(address account);
    event ProverRemoved(address account);

    modifier onlyProver() {
        require(proverStates[msg.sender] != ProverState.Null, "invalid prover");
        _;
    }

    modifier onlyActiveProver() {
        require(proverStates[msg.sender] == ProverState.Active, "invalid prover");
        _;
    }

    function addProvers(address[] memory _accounts) public onlyOwner {
        ProverState state = paused() ? ProverState.Paused : ProverState.Active;
        for (uint256 i = 0; i < _accounts.length; i++) {
            _addProver(_accounts[i], state);
        }
    }

    function removeProvers(address[] memory _accounts) public onlyOwner {
        for (uint256 i = 0; i < _accounts.length; i++) {
            _removeProver(_accounts[i]);
        }
    }

    function pause() public override onlyPauser {
        _pause();
        for (uint256 i = 0; i < provers.length; i++) {
            proverStates[provers[i]] = ProverState.Paused;
        }
    }

    function unpause() public override onlyPauser {
        _unpause();
        for (uint256 i = 0; i < provers.length; i++) {
            proverStates[provers[i]] = ProverState.Active;
        }
    }

    function numProvers() public view returns (uint256) {
        return provers.length;
    }

    function isActiveProver(address _account) public view returns (bool) {
        return proverStates[_account] == ProverState.Active;
    }

    function _addProver(address _account, ProverState _state) private {
        require(proverStates[_account] == ProverState.Null, "account is prover");
        provers.push(_account);
        proverStates[_account] = _state;
        emit ProverAdded(_account);
    }

    function _removeProver(address _account) private {
        require(proverStates[_account] != ProverState.Null, "account is not prover");
        uint256 lastIndex = provers.length - 1;
        for (uint256 i = 0; i < provers.length; i++) {
            if (provers[i] == _account) {
                if (i < lastIndex) {
                    provers[i] = provers[lastIndex];
                }
                provers.pop();
                delete proverStates[_account];
                emit ProverRemoved(_account);
                return;
            }
        }
        revert("prover not found"); // this should never happen
    }
}
