// SPDX-License-Identifier: GPL-3.0-only

pragma solidity >=0.8.0;

interface IBrevisRequestOwner {
    function setRequestTimeout(uint256 _timeout) external;

    function setBaseDataURL(string memory _url) external;

    function setBrevisProof(address _brevisProof) external;

    function setBrevisDispute(address _brevisDispute) external;

    function setBvnSigsVerifier(address _bvnSigsVerifier) external;

    function setAvsSigsVerifier(address _avsSigsVerifier) external;

    function setFeeCollector(address _feeCollector) external;

    function setChallengeWindow(uint256 _challengeWindow) external;

    function setResponseTimeout(uint256 _responseTimeout) external;

    function setDisputeDeposits(uint256 _amtAskForData, uint256 _amtAskForProof) external;
}
