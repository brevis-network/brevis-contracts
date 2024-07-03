// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

abstract contract BrevisApp {
    address public brevisRequest;
    uint256 public opChallengeWindow = 2 ** 256 - 1; // disable usage of op result by default
    uint8 public opSigOption = 0x01; // require BVN sig by default

    modifier onlyBrevisRequest() {
        require(msg.sender == brevisRequest, "invalid caller");
        _;
    }

    constructor(address _brevisRequest) {
        brevisRequest = _brevisRequest;
    }

    function handleProofResult(bytes32 _vkHash, bytes calldata _appCircuitOutput) internal virtual {
        // to be overrided by custom app
    }

    function handleOpProofResult(bytes32 _vkHash, bytes calldata _appCircuitOutput) internal virtual {
        // to be overrided by custom app
    }

    // app contract can implement logics to set opChallengeWindow if needed
    function _setOpChallengeWindow(uint256 _challangeWindow) internal {
        opChallengeWindow = _challangeWindow;
    }

    // app contract can implement logics to set opSigOption if needed
    function _setOpSigOption(uint8 _opSigOption) internal {
        opSigOption = _opSigOption;
    }

    // app contract can implement logics to update brevisRequest address if needed
    function _setBrevisRequest(address _brevisRequest) internal {
        brevisRequest = _brevisRequest;
    }

    function brevisCallback(bytes32 _appVkHash, bytes calldata _appCircuitOutput) external onlyBrevisRequest {
        handleProofResult(_appVkHash, _appCircuitOutput);
    }

    function brevisBatchCallback(
        bytes32[] calldata _appVkHashes,
        bytes[] calldata _appCircuitOutputs
    ) external onlyBrevisRequest {
        for (uint i = 0; i < _appVkHashes.length; i++) {
            handleProofResult(_appVkHashes[i], _appCircuitOutputs[i]);
        }
    }

    function applyBrevisOpResult(
        bytes32 _proofId,
        uint64 _nonce,
        bytes32 _appVkHash,
        bytes32 _appCommitHash,
        bytes calldata _appCircuitOutput
    ) public {
        require(
            IBrevisRequest(brevisRequest).validateOpAppData(
                _proofId,
                _nonce,
                _appCommitHash,
                _appVkHash,
                opChallengeWindow,
                opSigOption
            ),
            "data not ready to use"
        );
        require(_appCommitHash == keccak256(_appCircuitOutput), "invalid circuit output");
        handleOpProofResult(_appVkHash, _appCircuitOutput);
    }

    function applyBrevisOpResults(
        bytes32[] calldata _proofIds,
        uint64[] calldata _nonces,
        bytes32[] calldata _appVkHashes,
        bytes32[] calldata _appCommitHashes,
        bytes[] calldata _appCircuitOutputs
    ) external {
        for (uint256 i = 0; i < _proofIds.length; i++) {
            applyBrevisOpResult(_proofIds[i], _nonces[i], _appVkHashes[i], _appCommitHashes[i], _appCircuitOutputs[i]);
        }
    }
}

interface IBrevisRequest {
    function validateOpAppData(
        bytes32 _proofId,
        uint64 _nonce,
        bytes32 _appCommitHash,
        bytes32 _appVkHash,
        uint256 _appChallengeWindow,
        uint8 _option
    ) external view returns (bool);
}
