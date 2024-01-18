import { Fixture } from 'ethereum-waffle';
import { ethers, waffle } from 'hardhat';
import { MessageBridge, MsgTest, MockLightClient, MockMessageBridge } from '../../typechain';
import { MockMerkleProofTree } from '../../typechain/MockMerkleProofTree';
import { MessageBridge__factory } from './../../typechain/factories/MessageBridge__factory';
import { MsgTest__factory } from './../../typechain/factories/MsgTest__factory';

import { expect } from 'chai';
import { Wallet } from 'ethers';
import { keccak256 } from 'ethers/lib/utils';
import { generateProof } from './util';

describe('MessageBridge Test', async () => {
  function loadFixture<T>(fixture: Fixture<T>): Promise<T> {
    const provider = waffle.provider;
    return waffle.createFixtureLoader(provider.getWallets(), provider)(fixture);
  }

  async function fixture([admin]: Wallet[]) {
    const { mockMessageBridge, messageBridge, mockLightClient, merkleProofTree, messageTest } = await deployLib(admin);
    return { admin, mockMessageBridge, messageBridge, mockLightClient, merkleProofTree, messageTest };
  }

  let _admin: Wallet;
  let _mockMessageBridge: MockMessageBridge;
  let _mockLightClient: MockLightClient;
  let _messageBridge: MessageBridge;
  let _merkleProofTree: MockMerkleProofTree;
  let _msgTest: MsgTest;
  let _chainId: number;

  beforeEach(async () => {
    const res = await loadFixture(fixture);
    const { admin, mockMessageBridge, messageBridge, mockLightClient, merkleProofTree, messageTest } = res;
    _admin = admin;
    _mockMessageBridge = mockMessageBridge as MockMessageBridge;
    _messageBridge = messageBridge as MessageBridge;
    _mockLightClient = mockLightClient as MockLightClient;
    _merkleProofTree = merkleProofTree as MockMerkleProofTree;
    _msgTest = messageTest as MsgTest;
    _chainId = (await ethers.provider.getNetwork()).chainId;
  });

  async function deployLib(admin: Wallet) {
    const merkleFactory = await ethers.getContractFactory('MockMerkleProofTree');
    const merkleProofTree = (await merkleFactory.connect(admin).deploy()) as MockMerkleProofTree;

    const factory = await ethers.getContractFactory('MockMessageBridge');
    const mockMessageBridge = (await factory.connect(admin).deploy()) as MockMessageBridge;

    const mockLightClientFactory = await ethers.getContractFactory('MockLightClient');
    const mockLightClient = (await mockLightClientFactory.connect(admin).deploy()) as MockLightClient;

    const messageBridgeFactory = await ethers.getContractFactory<MessageBridge__factory>('MessageBridge');
    const messageBridge = (await messageBridgeFactory.connect(admin).deploy()) as MessageBridge;

    const messageTestFactory = await ethers.getContractFactory<MsgTest__factory>('MsgTest');
    const messageTest = (await messageTestFactory.connect(admin).deploy(messageBridge.address)) as MsgTest;

    const chainId = (await ethers.provider.getNetwork()).chainId;
    await messageBridge.connect(admin).setLightClient(chainId, mockLightClient.address);
    await messageBridge.connect(admin).setRemoteMessageBridge(chainId, messageBridge.address);

    return { mockMessageBridge, messageBridge, mockLightClient, merkleProofTree, messageTest };
  }

  it('should pass with execute message with success state', async () => {
    const slot = 1234567;
    const accountAddress = _messageBridge.address;
    const nonce = 32;
    const srcContract = '0xA2B26126ee3E7A26183F4d76837CB6d56bE56637';
    const message = ethers.utils.defaultAbiCoder.encode(['address', 'uint64'], [_admin.address, 66]);
    const { stProof, acntProof } = await generateProof(
      nonce,
      srcContract,
      _msgTest.address,
      _chainId,
      _chainId,
      message,
      accountAddress
    );

    await _mockMessageBridge.initialize(
      slot,
      _messageBridge.address,
      _mockLightClient.address,
      keccak256(acntProof[0])
    );

    await expect(
      _mockMessageBridge.testExecutedMessage(
        _chainId,
        nonce,
        srcContract,
        _msgTest.address,
        message,
        acntProof,
        stProof
      )
    )
      .to.emit(_messageBridge, 'StorageRootVerified')
      .to.emit(_messageBridge, 'MessageExecuted')
      .to.emit(_msgTest, 'MessageReceived')
      .withArgs(_chainId, srcContract, _admin.address, 66);
  });

  it('should pass with execute message with abort', async () => {
    const slot = 1234567;
    const accountAddress = _messageBridge.address;
    const nonce = 32;
    const srcContract = '0xA2B26126ee3E7A26183F4d76837CB6d56bE56637';
    const message = ethers.utils.defaultAbiCoder.encode(['address', 'uint64'], [_admin.address, 1000]);
    const { stProof, acntProof } = await generateProof(
      nonce,
      srcContract,
      _msgTest.address,
      _chainId,
      _chainId,
      message,
      accountAddress
    );

    await _mockMessageBridge.initialize(
      slot,
      _messageBridge.address,
      _mockLightClient.address,
      keccak256(acntProof[0])
    );

    await expect(
      _mockMessageBridge.testExecutedMessage(
        _chainId,
        nonce,
        srcContract,
        _msgTest.address,
        message,
        acntProof,
        stProof
      )
    ).to.be.revertedWith('MSG::ABORT:test abort');
  });

  it('should pass with execute message with failed state', async () => {
    const slot = 1234567;
    const accountAddress = _messageBridge.address;
    const nonce = 32;
    const srcContract = '0xA2B26126ee3E7A26183F4d76837CB6d56bE56637';
    const message = ethers.utils.defaultAbiCoder.encode(['address', 'uint64'], [_admin.address, 1001]);
    const { stProof, acntProof } = await generateProof(
      nonce,
      srcContract,
      _msgTest.address,
      _chainId,
      _chainId,
      message,
      accountAddress
    );

    await _mockMessageBridge.initialize(
      slot,
      _messageBridge.address,
      _mockLightClient.address,
      keccak256(acntProof[0])
    );

    await expect(
      _mockMessageBridge.testExecutedMessage(
        _chainId,
        nonce,
        srcContract,
        _msgTest.address,
        message,
        acntProof,
        stProof
      )
    )
      .to.emit(_messageBridge, 'StorageRootVerified')
      .to.emit(_messageBridge, 'MessageExecuted')
      .to.emit(_messageBridge, 'MessageCallReverted');
  });
});
