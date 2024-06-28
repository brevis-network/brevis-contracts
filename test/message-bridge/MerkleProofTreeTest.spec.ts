import assert from 'assert';
import { ContractRunner, keccak256, decodeRlp, toBeArray, AbiCoder, solidityPackedKeccak256 } from 'ethers';
import { ethers } from 'hardhat';
import { MockMerkleProofTree } from '../../typechain/contracts/apps/message-bridge/mock/MockMerkleProofTree';
import { computeMessageId, generateProof, hash2bytes } from './util';
import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers';

describe('MerkleProofTree Test', async () => {
  async function fixture() {
    const [admin] = await ethers.getSigners();
    const merkleProofTree = await deployLib(admin);
    return { merkleProofTree };
  }

  let merkleProofTree: MockMerkleProofTree;
  let chainId: number;

  beforeEach(async () => {
    const res = await loadFixture(fixture);
    merkleProofTree = res.merkleProofTree as MockMerkleProofTree;
    chainId = Number((await ethers.provider.getNetwork()).chainId);
  });

  async function deployLib(admin: ContractRunner) {
    const factory = await ethers.getContractFactory('MockMerkleProofTree');
    const merkleProofTree = factory.connect(admin).deploy();
    return merkleProofTree;
  }

  it('should pass with read and verify value from proofs', async () => {
    const accountAddress = '0x00000000000000000000000000000000000000ab';
    const nonce = 32;
    const sender = '0xA2B26126ee3E7A26183F4d76837CB6d56bE56637';
    const receiver = '0xA2B26126ee3E7A26183F4d76837CB6d56bE56637';
    const message = '0xabcd';
    const { stProof, acntProof } = await generateProof(
      nonce,
      sender,
      receiver,
      chainId,
      chainId,
      message,
      accountAddress
    );

    const accountPath = hash2bytes(accountAddress);
    const accountPathValue = await merkleProofTree.mockRead(accountPath, acntProof);
    const storageInfoFromAccountPathValue = decodeRlp(accountPathValue)[2];
    assert.equal(storageInfoFromAccountPathValue, keccak256(stProof[0]));

    const storagePath = hash2bytes(hash2bytes(toBeArray(AbiCoder.defaultAbiCoder().encode(['uint64', 'uint256'], [nonce, 2]))));
    const storagePathValue = await merkleProofTree.mockRead(storagePath, stProof);

    const messageId = computeMessageId(nonce, sender, receiver, chainId, chainId, message).messageId;
    assert.equal(decodeRlp(storagePathValue), messageId);
  });
});
