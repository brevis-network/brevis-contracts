import { expect } from 'chai';
import { Fixture } from 'ethereum-waffle';
import { BigNumber, BigNumberish, Wallet } from 'ethers';
import { ethers, waffle } from 'hardhat';
import {
  BN254NewVerifier__factory,
  BN254NewVerifier,
  AggregationVerifier,
  AggregationVerifier__factory
} from '../../typechain';
import assert from 'assert';
import { hexToBytes } from '../util';

async function deployContract(admin: Wallet) {
  const _factory = await ethers.getContractFactory('AggregationVerifier');
  const _contract = await _factory.connect(admin).deploy();
  return _contract;
}

describe('BN254 final_emulate proof verifier', async () => {
  function loadFixture<T>(fixture: Fixture<T>): Promise<T> {
    const provider = waffle.provider;
    return waffle.createFixtureLoader(provider.getWallets(), provider)(fixture);
  }

  async function fixture([admin]: Wallet[]) {
    const contract = await deployContract(admin);
    return { admin, contract };
  }

  let contract: AggregationVerifier;
  let admin: Wallet;
  beforeEach(async () => {
    const res = await loadFixture(fixture);
    contract = res.contract;
    admin = res.admin;
  });

  it('should pass on true proof', async () => {
    const result = await contract.verifyProofWithCommit(
      [
        BigNumber.from('7098988087689144641203879580397093154573435494056885401628087621279772152016'),
        BigNumber.from('2619442256419406363560033246421835527980947896074630699814006358960491764643'),

        BigNumber.from('4782692870294396253325213900836483075573237374545999396032284488094394593819'),
        BigNumber.from('4538103055553572184945752732535192213272333755912635851796079758059720180653'),
        BigNumber.from('11755698893767106647302652438947051394260170597388223244940562848499762418848'),
        BigNumber.from('16445248470260456720129520237821147436418596306747812564920016577065208685617'),

        BigNumber.from('3473228272257922598007303562952158401623568452125572421614604451909359183703'),
        BigNumber.from('12323358119480857658089219478054552232809411076351206536072169945611742328465')
      ],
      [
        BigNumber.from('0x2400930cab3fe30f8e9a820f3bd6753d61d6b57b5b010d1bbca8b635d0be92b4'),
        BigNumber.from('0x0b6adb08eaecfbb1fc3045ec69cafa43fcbc2c8274653a391761193d605a6518')
      ],
      [
        BigNumber.from('0x1188defa0ffaa65debe1dcb9a21f8b9b805d2b792f91232c1c3cb59e0a604132'),
        BigNumber.from('0x0df1656fd4eb26023bb8062183e1059c'),
        BigNumber.from('0x9f55cdb00efe07199a64ff80b2b9dfcb'),
        BigNumber.from('0x2fa2588cf088bd167879f3de8cd4dfd65f84a5550b0236a900f5f0ce551ad08f'),
        BigNumber.from('0x1468288056310c82aa4c01a7e12a10f8'),
        BigNumber.from('0x111a0560e72b700555479031b86c357d'),
        BigNumber.from('0x2a3c81f1a247ebf31c3141747737158c4c89f7c420c8647018d312be693bf638'),
        BigNumber.from('0x275f76130e82f53a8b06a062ad6ccf1e4e07e7a6310decc300fea93712d3b381')
      ]
    );
    assert.equal(result, true);
  });

  it('should pass on verify raw with real data', async () => {
    const values = [
      BigNumber.from('0x28d3939b8a291d8a0aa5c4c25683120eaa4b65c4d914caa2c730e724875a88'), // a0
      BigNumber.from('0x0751827af1c3b5410765da6db4ec4842f68676e17bbb1f8aca82a00ff78c986b'), // a1

      BigNumber.from('0x17cd79d36dc39121bca789137a55b0472571ceddd0acc17f8a126d150321eb58'), // b00
      BigNumber.from('0x0569bc2d22aaf6f76c7b0369a3d7dc5bbe040e9c7e4be8c98d7989084a762d8a'), // b01
      BigNumber.from('0x05090b8e4926760ca52d424a903945aa27d8d43f810016cb776bdf0e74c2998a'), // b10
      BigNumber.from('0x09f285728d8eed0ebd43020b02395aacbc50c3c53188a12f3c2a5447ea444ad1'), // b11

      BigNumber.from('0x1a10075f551f8aab2125c937e8c0a01fb5d3df380b8753d6d90433118e89b924'), // c0
      BigNumber.from('0x22297a2c6891897be6149fa121ff9bfdf978972a716ccaa1fcd816270afb1d9f'), // c1

      BigNumber.from('0x03703adff9ac652f748b8318ae5af481d5d215759e133fb9e5a0c48159e7a2ef'), // Commitment 0
      BigNumber.from('0x290a912d0004e1af044dee62c4fd0f04a23e152dc1635a89df7d78b916337b47'), // Commitment 1

      BigNumber.from('0x11a8db876102920b6d51d9ddd535d25981a0be88dc34c46073e9c126190ed73a'), // Commitment Public Input
      BigNumber.from('0x2026cd953780d0ef6fcc2d69a96bff104bcb661918a4f6bb4c6862cd5e93eb2d'), // Query Hash
      BigNumber.from('0x08646494e28db9b84dbf61f1725e5174f7eaee5265971f0d92f33559810e4428'), // SMT Root 0 Should Be padded to bytes32
      BigNumber.from('0x15a9cc96d5579d53c09ebebb1e8289156358a2dbd0568d9a5c634240312df129'), // Aggregation VK Hash
      BigNumber.from('0x1468288056310c82aa4c01a7e12a10f8111a0560e72b700555479031b86c357d'), // App Circuit Output Commitment 0 Should Be padded to bytes32
      BigNumber.from('0x042989bb107bdd497bf3de4ed66aece2aa1fa9950ae65211f441e6df1b95af9b') // App Circuit VK Hash
    ];

    var hexValues = '';

    values.forEach((value) => {
      hexValues += value.toHexString().slice(2).padStart(64, '0');
    });

    const result = await contract.verifyRaw(hexToBytes('0x' + hexValues));

    assert.equal(result, true);
  });
});
