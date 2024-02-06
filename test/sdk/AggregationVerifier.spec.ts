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
        BigNumber.from('0x1e82206b07cd2f3ea2571ef9086a941fe266482e1bc1e703b6681461ccafde08'),
        BigNumber.from('0x1284cfcef8e4bfe1affb6ab65444cffde7ef83dce5df55ff12a40257fdcf73a0'),

        BigNumber.from('0x0bb8b2d3ea6b76bbe1f8651c1a74a04a52310c293917f40f1ed1468a14a2f21e'),
        BigNumber.from('0x207f94da832cf3992a2e03a34491330a533e2ffa446ad2d06704ed9f60f9bf17'),
        BigNumber.from('0x252c41c99af3a35e59d347d9642ba8f7e99e191d9aab596a3c592bfa5e22e0b3'),
        BigNumber.from('0x160bd8c557bed58bed0a4284b94660ee5cd2b220d7d93b5a444e18e84c42a623'),

        BigNumber.from('0x15e5ce69173928134057ee3a97157c1d606df1504fecafd8303b00be20ed6ecc'),
        BigNumber.from('0x1ea84a3222cd97f76ad0e4596f1888421fb69ae1242fc41537e8c482dad5f3a2')
      ],
      [
        BigNumber.from('0x01d8fba65a7a43afb2d8a7102d552e9ac4b01048ed2a27e8a1fef67cd063d551'),
        BigNumber.from('0x1f346367962f8ab4d8f6704114144664c3635a61367325aa1c1d8b6c2803da37')
      ],
      [
        BigNumber.from('0x1671e609c54bc73eed306d12add15be48965cf0d232a4368095719ca21dd36df'),
        BigNumber.from('0x05826920570b43c036b2f33b827350bc'),
        BigNumber.from('0x86d8bd2dfe9a234e7ee6a758bf3c6aaf'),
        BigNumber.from('0x1e73236b69ab060e3a7bab6b2a58bb431638fa1c382d51d1e1b72002c55a3649'),
        BigNumber.from('0x2601b828c023e4c35bcd8b761f26aa6a'),
        BigNumber.from('0xb63260d34cdb9bfe7e4bb185c9745580'),
        BigNumber.from('0x1d7f35f3a9b09f723857802db081adfa627b5cb389539ac04eedf6d422a52ed2'),
        BigNumber.from('0x2f578acba6dbd255754829e04de0d0bc539bbd0894be74952bb6da11e3fe57a3')
      ]
    );
    assert.equal(result, true);
  });

  it('should pass on verify raw with real data', async () => {
    const values = [
      BigNumber.from('0x1e82206b07cd2f3ea2571ef9086a941fe266482e1bc1e703b6681461ccafde08'), // a0
      BigNumber.from('0x1284cfcef8e4bfe1affb6ab65444cffde7ef83dce5df55ff12a40257fdcf73a0'), // a1

      BigNumber.from('0x0bb8b2d3ea6b76bbe1f8651c1a74a04a52310c293917f40f1ed1468a14a2f21e'), // b00
      BigNumber.from('0x207f94da832cf3992a2e03a34491330a533e2ffa446ad2d06704ed9f60f9bf17'), // b01
      BigNumber.from('0x252c41c99af3a35e59d347d9642ba8f7e99e191d9aab596a3c592bfa5e22e0b3'), // b10
      BigNumber.from('0x160bd8c557bed58bed0a4284b94660ee5cd2b220d7d93b5a444e18e84c42a623'), // b11

      BigNumber.from('0x15e5ce69173928134057ee3a97157c1d606df1504fecafd8303b00be20ed6ecc'), // c0
      BigNumber.from('0x1ea84a3222cd97f76ad0e4596f1888421fb69ae1242fc41537e8c482dad5f3a2'), // c1

      BigNumber.from('0x01d8fba65a7a43afb2d8a7102d552e9ac4b01048ed2a27e8a1fef67cd063d551'), // Commitment 0
      BigNumber.from('0x1f346367962f8ab4d8f6704114144664c3635a61367325aa1c1d8b6c2803da37'), // Commitment 1

      BigNumber.from('0x2f578acba6dbd255754829e04de0d0bc539bbd0894be74952bb6da11e3fe57a3'), // Commitment Public Input
      BigNumber.from('0x1671e609c54bc73eed306d12add15be48965cf0d232a4368095719ca21dd36df'), // Query Hash
      BigNumber.from('0x05826920570b43c036b2f33b827350bc86d8bd2dfe9a234e7ee6a758bf3c6aaf'), // SMT Root 0 Should Be padded to bytes32
      BigNumber.from('0x1e73236b69ab060e3a7bab6b2a58bb431638fa1c382d51d1e1b72002c55a3649'), // Aggregation VK Hash
      BigNumber.from('0x2601b828c023e4c35bcd8b761f26aa6ab63260d34cdb9bfe7e4bb185c9745580'), // App Circuit Output Commitment 0 Should Be padded to bytes32
      BigNumber.from('0x1d7f35f3a9b09f723857802db081adfa627b5cb389539ac04eedf6d422a52ed2') // App Circuit VK Hash
    ];

    var hexValues = '';

    values.forEach((value) => {
      hexValues += value.toHexString().slice(2).padStart(64, '0');
    });

    const result = await contract.verifyRaw(hexToBytes('0x' + hexValues));

    assert.equal(result, true);
  });
});
