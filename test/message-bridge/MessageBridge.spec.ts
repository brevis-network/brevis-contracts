import { expect } from 'chai';
import { AbiCoder, keccak256 } from 'ethers';
import { ethers } from 'hardhat';

import { HardhatEthersSigner } from '@nomicfoundation/hardhat-ethers/signers';
import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers';

import {
  MessageBridge,
  MessageBridge__factory,
  MockLightClient,
  MockLightClient__factory,
  MockMerkleProofTree__factory,
  MockMessageBridge,
  MockMessageBridge__factory,
  MsgTest,
  MsgTest__factory,
} from '../../typechain';
import { MockMerkleProofTree } from '../../typechain/contracts/apps/message-bridge/mock/MockMerkleProofTree';
import { generateProof } from './util';

describe('MessageBridge Test', async () => {
  async function fixture() {
    const { admin, mockMessageBridge, messageBridge, mockLightClient, merkleProofTree, messageTest } =
      await deployLib();
    return { admin, mockMessageBridge, messageBridge, mockLightClient, merkleProofTree, messageTest };
  }

  let _admin: HardhatEthersSigner;
  let _mockMessageBridge: MockMessageBridge;
  let _mockLightClient: MockLightClient;
  let _mockLightClientAddress: string;
  let _messageBridge: MessageBridge;
  let _messageBridgeAddress: string;
  let _merkleProofTree: MockMerkleProofTree;
  let _msgTest: MsgTest;
  let _msgTestAddress: string;
  let _chainId: number;

  beforeEach(async () => {
    const res = await loadFixture(fixture);
    const { admin, mockMessageBridge, messageBridge, mockLightClient, merkleProofTree, messageTest } = res;
    _admin = admin;
    _mockMessageBridge = mockMessageBridge as MockMessageBridge;
    _messageBridge = messageBridge as MessageBridge;
    _messageBridgeAddress = await messageBridge.getAddress();
    _mockLightClient = mockLightClient as MockLightClient;
    _mockLightClientAddress = await _mockLightClient.getAddress();
    _merkleProofTree = merkleProofTree as MockMerkleProofTree;
    _msgTest = messageTest as MsgTest;
    _msgTestAddress = await messageTest.getAddress();
    _chainId = Number((await ethers.provider.getNetwork()).chainId);
  });

  async function deployLib() {
    const [admin] = await ethers.getSigners();

    const merkleFactory = new MockMerkleProofTree__factory();
    const merkleProofTree = await merkleFactory.connect(admin).deploy();

    const factory = new MockMessageBridge__factory();
    const mockMessageBridge = await factory.connect(admin).deploy();

    const mockLightClientFactory = new MockLightClient__factory();
    const mockLightClient = await mockLightClientFactory.connect(admin).deploy();
    const mockLightClientAddress = await mockLightClient.getAddress();

    const messageBridgeFactory = new MessageBridge__factory();
    const messageBridge = await messageBridgeFactory.connect(admin).deploy();
    const messageBridgeAddress = await messageBridge.getAddress();

    const messageTestFactory = new MsgTest__factory();
    const messageTest = await messageTestFactory.connect(admin).deploy(messageBridgeAddress);

    const chainId = (await ethers.provider.getNetwork()).chainId;
    await messageBridge.connect(admin).setLightClient(chainId, mockLightClientAddress);
    await messageBridge.connect(admin).setRemoteMessageBridge(chainId, messageBridgeAddress);

    return { admin, mockMessageBridge, messageBridge, mockLightClient, merkleProofTree, messageTest };
  }

  it('should pass with execute message with success state', async () => {
    const slot = 1234567;
    const accountAddress = _messageBridgeAddress;
    const nonce = 32;
    const srcContract = '0xA2B26126ee3E7A26183F4d76837CB6d56bE56637';
    const message = AbiCoder.defaultAbiCoder().encode(['address', 'uint64'], [_admin.address, 66]);
    const { stProof, acntProof } = await generateProof(
      nonce,
      srcContract,
      _msgTestAddress,
      _chainId,
      _chainId,
      message,
      accountAddress
    );

    await _mockMessageBridge.initialize(slot, _messageBridgeAddress, _mockLightClientAddress, keccak256(acntProof[0]));

    await expect(
      _mockMessageBridge.testExecutedMessage(_chainId, nonce, srcContract, _msgTestAddress, message, acntProof, stProof)
    )
      .to.emit(_messageBridge, 'MessageExecuted')
      .to.emit(_msgTest, 'MessageReceived')
      .withArgs(_chainId, srcContract, _admin.address, 66);
  });

  it('should pass with execute message with abort', async () => {
    const slot = 1234567;
    const accountAddress = await _messageBridge.getAddress();
    const nonce = 32;
    const srcContract = '0xA2B26126ee3E7A26183F4d76837CB6d56bE56637';
    const message = AbiCoder.defaultAbiCoder().encode(['address', 'uint64'], [_admin.address, 1000]);
    const { stProof, acntProof } = await generateProof(
      nonce,
      srcContract,
      _msgTestAddress,
      _chainId,
      _chainId,
      message,
      accountAddress
    );

    await _mockMessageBridge.initialize(slot, _messageBridgeAddress, _mockLightClientAddress, keccak256(acntProof[0]));

    await expect(
      _mockMessageBridge.testExecutedMessage(_chainId, nonce, srcContract, _msgTestAddress, message, acntProof, stProof)
    ).to.be.revertedWith('MSG::ABORT:test abort');
  });

  it('should pass with execute message with failed state', async () => {
    const slot = 1234567;
    const accountAddress = _messageBridgeAddress;
    const nonce = 32;
    const srcContract = '0xA2B26126ee3E7A26183F4d76837CB6d56bE56637';
    const message = AbiCoder.defaultAbiCoder().encode(['address', 'uint64'], [_admin.address, 1001]);
    const { stProof, acntProof } = await generateProof(
      nonce,
      srcContract,
      _msgTestAddress,
      _chainId,
      _chainId,
      message,
      accountAddress
    );

    await _mockMessageBridge.initialize(slot, _messageBridgeAddress, _mockLightClientAddress, keccak256(acntProof[0]));

    await expect(
      _mockMessageBridge.testExecutedMessage(_chainId, nonce, srcContract, _msgTestAddress, message, acntProof, stProof)
    )
      .to.emit(_messageBridge, 'MessageExecuted')
      .to.emit(_messageBridge, 'MessageCallReverted');
  });
});
