import { Fixture } from 'ethereum-waffle';
import { BigNumber, Wallet } from 'ethers';
import { ethers, waffle } from 'hardhat';
import { Plonky2AggProofVerifier } from '../../typechain';
import assert from 'assert';
import {hexToBytes} from "../util";

async function deployContract(admin: Wallet) {
  const _factory = await ethers.getContractFactory('Plonky2AggProofVerifier');
  const _contract = await _factory.connect(admin).deploy();
  return _contract;
}

describe('plonky2 app agg verifier', async () => {
  function loadFixture<T>(fixture: Fixture<T>): Promise<T> {
    const provider = waffle.provider;
    return waffle.createFixtureLoader(provider.getWallets(), provider)(fixture);
  }

  async function fixture([admin]: Wallet[]) {
    const contract = await deployContract(admin);
    return { admin, contract };
  }

  let contract: Plonky2AggProofVerifier;
  let admin: Wallet;
  beforeEach(async () => {
    const res = await loadFixture(fixture);
    contract = res.contract;
    admin = res.admin;
  });

  it('should pass on true proof', async () => {
    const result = await contract.verifyProof(
      [
        BigNumber.from('0x2eda7d34931de170fcda24ca704ba547b06b566f70f66644ecc328333135ef33'), // a0
        BigNumber.from('0x2cb117ead10d2cd391d7bf138b545aea78bec2cbfda2cc84e1fa8d1e065cdd58'), // a1

        BigNumber.from('0x106332160ef67360b694d65791793cdf8975a3cccb332419162108e7d97b83d1'), // b00
        BigNumber.from('0x16c9b729cfcae8294ed4663cbb1359207e4803878cf3fd4c5eb09180284b5b9e'), // b01
        BigNumber.from('0x19de30697cf3ce3aff718bf11fb41acefdf1128087f555bd175161598ec5bbf8'), // b10
        BigNumber.from('0x1a347dac976c0bda6f03c70efe5f8a137decaf4557591eae1d3a21981a95d2a2'), // b11

        BigNumber.from('0x130969ff3ec214e6499a449d6cf019458b518f9b97149d273342371684cfe741'), // c0
        BigNumber.from('0x0e0c71ad98c542a5c65469a944c158c633a0e477e4c8fb2b7e174c552204730b') // c1
      ],
      [
        BigNumber.from('0x098fff58cfc4c1706fb7e4be03f1f2eaebbe9b276859037f98bbce387ce17b60'), // Commitment 0
        BigNumber.from('0x037bbf27d44dd6bf09e2fcf524e39df5fd475b1263d02008c1aa2231b66c13ca') // Commitment 1
      ],
      [
        BigNumber.from('0x13efaaa8bbab7c933da36361cd8449323d07d70213ec24cc52750efbb8207104'), // Commitment POK0
        BigNumber.from('0x1f0ce031346253c1fabd300a9b7fb2989665c8d06d40eb117ddf500bdbd28580') // Commitment POK1
      ],
      [
        BigNumber.from('0x0000000000000000000000000000000058135c8317b5cbf63d7567de0693818f'),
        BigNumber.from('0x00000000000000000000000000000000d7377885e3e26a73ed88f828b44ce437'),
        BigNumber.from('0x00000000000000000000000000000000efd3a1e24fd7fed7c07d0e9c001421fc'),
        BigNumber.from('0x000000000000000000000000000000002185153b9345a5d2e9c0513edcd59a87'),
        BigNumber.from('0x17db68ec6ce43b8a521f8d2b0ac528aa6822f715081ed194a818f3a5daf0f253')
      ]
    );
    assert.equal(result, true);
  });

  it('should pass on verify raw proof', async () => {
    const values = [
      BigNumber.from('0x2eda7d34931de170fcda24ca704ba547b06b566f70f66644ecc328333135ef33'), // a0
      BigNumber.from('0x2cb117ead10d2cd391d7bf138b545aea78bec2cbfda2cc84e1fa8d1e065cdd58'), // a1

      BigNumber.from('0x106332160ef67360b694d65791793cdf8975a3cccb332419162108e7d97b83d1'), // b00
      BigNumber.from('0x16c9b729cfcae8294ed4663cbb1359207e4803878cf3fd4c5eb09180284b5b9e'), // b01
      BigNumber.from('0x19de30697cf3ce3aff718bf11fb41acefdf1128087f555bd175161598ec5bbf8'), // b10
      BigNumber.from('0x1a347dac976c0bda6f03c70efe5f8a137decaf4557591eae1d3a21981a95d2a2'), // b11

      BigNumber.from('0x130969ff3ec214e6499a449d6cf019458b518f9b97149d273342371684cfe741'), // c0
      BigNumber.from('0x0e0c71ad98c542a5c65469a944c158c633a0e477e4c8fb2b7e174c552204730b'), // c1

      BigNumber.from('0x098fff58cfc4c1706fb7e4be03f1f2eaebbe9b276859037f98bbce387ce17b60'), // Commitment 0
      BigNumber.from('0x037bbf27d44dd6bf09e2fcf524e39df5fd475b1263d02008c1aa2231b66c13ca'), // Commitment 1

      BigNumber.from('0x13efaaa8bbab7c933da36361cd8449323d07d70213ec24cc52750efbb8207104'), // Commitment POK0
      BigNumber.from('0x1f0ce031346253c1fabd300a9b7fb2989665c8d06d40eb117ddf500bdbd28580'), // Commitment POK1

      BigNumber.from('0xd7377885e3e26a73ed88f828b44ce43758135c8317b5cbf63d7567de0693818f'),
      BigNumber.from('0x2185153b9345a5d2e9c0513edcd59a87efd3a1e24fd7fed7c07d0e9c001421fc'),
      BigNumber.from('0x17db68ec6ce43b8a521f8d2b0ac528aa6822f715081ed194a818f3a5daf0f253')
    ];

    var hexValues = '';

    values.forEach((value) => {
      hexValues += value.toHexString().slice(2).padStart(64, '0');
    });

    console.log('hexValues: ', hexValues);
    const result = await contract.verifyRaw(hexToBytes('0x' + hexValues));

    assert.equal(result, true);
  });
});
