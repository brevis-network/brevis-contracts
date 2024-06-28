import { assert } from 'console';
import { BigNumberish, ContractRunner, getBytes, keccak256 } from 'ethers';
import { ethers } from 'hardhat';
import { MerkleTree } from 'merkletreejs';

import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers';

import {
  BlockChunks,
  BlockChunks__factory,
  MockAnchorBlocks,
  MockAnchorBlocks__factory,
  MockZkVerifier__factory
} from '../../typechain';
import { splitHash } from '../util';

async function deployMockAnchorBlocksContract(admin: ContractRunner) {
  const factory = new MockAnchorBlocks__factory();
  const contract = await factory.connect(admin).deploy();
  return contract;
}
async function deployMockZkVerifierContract(admin: ContractRunner) {
  const factory = new MockZkVerifier__factory();
  const contract = await factory.connect(admin).deploy();
  return contract;
}
async function deployBlockChunksContract(admin: ContractRunner) {
  const factory = new BlockChunks__factory();
  const contract = await factory.connect(admin).deploy();
  return contract;
}

function getMockBlkHash(blk: number) {
  const leaf0 = '0x0000000000000000000000000000000000000000000000000000000000000000';
  const blkHex = blk.toString(16);
  return leaf0.slice(0, leaf0.length - blkHex.length) + blkHex;
}

function getChuckRoot(startBlkNum: number) {
  const leaves = [];
  for (let i = 0; i < 128; i++) {
    leaves.push(getMockBlkHash(startBlkNum + i));
  }
  const tree = new MerkleTree(leaves, keccak256);
  const root = tree.getRoot().toString('hex');
  const endHash = leaves[128 - 1];
  const prevHash = startBlkNum > 0 ? getMockBlkHash(startBlkNum - 1) : '0x';
  return { root, endHash, prevHash };
}

function getHexProof(blkNum: number) {
  const startBlkNum = blkNum - (blkNum % 128);
  const leaves = [];
  for (let i = 0; i < 128; i++) {
    leaves.push(getMockBlkHash(startBlkNum + i));
  }
  const tree = new MerkleTree(leaves, keccak256);
  let proof = tree.getHexProof(getMockBlkHash(blkNum));
  return proof;
}

function getTestProof(startBlkNum: number) {
  let { root, endHash, prevHash } = getChuckRoot(startBlkNum);

  const input = [...splitHash(root), ...splitHash(prevHash), ...splitHash(endHash), startBlkNum, startBlkNum + 127];
  const a: [BigNumberish, BigNumberish] = [
    '16217230224774761590414973642073192485520807822804992513061256091576710039093',
    '2002416893759411048918869122785221574310433394721837052525007122310353508563'
  ];
  const b: [[BigNumberish, BigNumberish], [BigNumberish, BigNumberish]] = [
    [
      '10197059906511048117905894151806242600921773365823995195260445360897935348383',
      '17176880858333124367872442540151136901916472062827225492995249573427445734932'
    ],
    [
      '3508508084177307333211815272848724987300711832513791765513593482163273805687',
      '8946113151421411368233424685315381086080512060853567398740775323993107361049'
    ]
  ];
  const c: [BigNumberish, BigNumberish] = [
    '4227181897368161064461503007610136746613299883531807541785600858067518608614',
    '10696546849598832835618070876676453336593868194332417262844290881924036536318'
  ];
  const commit: [BigNumberish, BigNumberish] = ['0', '0'];
  const allData = [...a];
  allData.push(...b[0], ...b[1]);
  allData.push(...c);
  allData.push(...commit);
  allData.push(...input);

  let allDataHex = '0x';
  for (let i = 0; i < allData.length; i++) {
    allDataHex = allDataHex + BigInt(allData[i]).toString(16).padStart(64, '0');
  }

  const proofData = allDataHex;
  return { proofData, endHash };
}

describe('Block Syncer Test', async () => {
  const chainId = 5;

  async function deployContractsFixture() {
    const [admin] = await ethers.getSigners();
    const anchor = await deployMockAnchorBlocksContract(admin);
    const verifier = await deployMockZkVerifierContract(admin);
    const syncer = await deployBlockChunksContract(admin);

    await syncer.updateAnchorBlockProvider(chainId, await anchor.getAddress());
    await syncer.updateVerifierAddress(chainId, await verifier.getAddress());

    return { syncer, anchor };
  }

  let syncer: BlockChunks;
  let anchor: MockAnchorBlocks;
  before(async () => {
    const res = await loadFixture(deployContractsFixture);
    syncer = res.syncer;
    anchor = res.anchor;
  });

  it('should pass on updateRecent', async () => {
    const res = getTestProof(256);
    await anchor.update(383, res.endHash);
    await syncer.updateRecent(chainId, res.proofData);
  });
  it('should pass on updateOld', async () => {
    const { proofData } = getTestProof(128);
    const res = getChuckRoot(256);
    const nextRoot = res.root;
    await syncer.updateOld(chainId, getBytes('0x' + nextRoot), 128, proofData);
  });
  it('should pass on isBlockHashValid', async () => {
    let success = await syncer.isBlockHashValid({
      chainId: chainId,
      blkNum: 253,
      claimedBlkHash: getMockBlkHash(253),
      prevHash: getMockBlkHash(127),
      numFinal: 128,
      merkleProof: getHexProof(253)
    });
    assert(success);
  });
});
