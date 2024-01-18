// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "../../bsc-tendermint/interfaces/IBSCValidatorSet.sol";

contract MockTendermintLightClient is IBSCValidatorSet {
    address public signerAddress;

    function updateSigner(address _newSigner) external {
        signerAddress = _newSigner;
    }

    constructor(address _signer) {
        signerAddress = _signer;
    }

    function isCurrentValidator(address _signer) external view returns (bool valid) {
        return signerAddress == _signer;
    }
}
