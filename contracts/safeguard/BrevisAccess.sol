// SPDX-License-Identifier: GPL-3.0-only

pragma solidity ^0.8.18;

import "./Pauser.sol";

// achive prover access and pauser control using a single map lookup
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
        require(proverStates[msg.sender] != ProverState.Null, "not invalid prover");
        _;
    }

    modifier onlyActiveProver() {
        require(proverStates[msg.sender] == ProverState.Active, "not invalid prover");
        _;
    }

    function addProvers(address[] memory _accounts) public onlyOwner {
        for (uint256 i = 0; i < _accounts.length; i++) {
            _addProver(_accounts[i]);
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

    function _addProver(address _account) private {
        require(proverStates[_account] == ProverState.Null, "account is prover");
        provers.push(_account);
        proverStates[_account] = ProverState.Active;
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
                proverStates[_account] = ProverState.Null;
                emit ProverRemoved(_account);
                return;
            }
        }
        revert("prover not found"); // this should never happen
    }
}
