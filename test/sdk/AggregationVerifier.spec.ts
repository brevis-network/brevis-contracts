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
        BigNumber.from('0x2d6ca9dc5f51406e358a694158054fa8173861cc05bb4cdb444240e5b5888dba'), // a0
        BigNumber.from('0x1b14c6562c6b57abb541adb8b7bcd75a8ce5d1228a28f0228802727a37cd65f8'), // a1

        BigNumber.from('0x011a5cfae06e386a11d80c54f6bb98d5bd04dea2bc0a3fccfa82d8222f76688f'), // b00
        BigNumber.from('0x0b8e2c859bcbbe39c1bc6d3ad437d6900d4a4dc563b3e0cb2d941ee6a70c1690'), // b01
        BigNumber.from('0x061c7739d940495c4cf5902db310dd4af559e03d08a901157d4c32da8d13f391'), // b10
        BigNumber.from('0x1797762d5fc38a017ec90d581ec3e7a0cc459f0225f75bdd609f4068217c1873'), // b11

        BigNumber.from('0x1596537d41a471914900275a70ad364fed4881273aca791881c4075d554368cc'), // c0
        BigNumber.from('0x1c645788f7b85fdff88c155dd32aa957163c86ee6b0331e1fd9b1f1642cfbfea')  // c1
      ],
      [
        BigNumber.from('0x0423485ea151245e4195ff2df2bd089b2c925e1c2ab786ac78aa8eac7b9bffee'), // Commitment 0
        BigNumber.from('0x22a5d19e908e5bb687fa66362b93aaca4b9191a233486c97cc2b8ef4ece93823')  // Commitment 1
      ],
      [
        BigNumber.from('0x1cf83ccf4eb7942493c24a7ce1c7857efe087953dcbbcf085d3cb885d5285b09'), // Query Hash
        BigNumber.from('0x05826920570b43c036b2f33b827350bc'), // SMT Root 0
        BigNumber.from('0x86d8bd2dfe9a234e7ee6a758bf3c6aaf'), // SMT Root 1
        BigNumber.from('0x0cfa2d22152d95266c0ccff01cedc28eafa8105dec9557125dbde2eceb5610dc'), // Aggregation VK Hash
        BigNumber.from('0x26ea6912d14adbed568e5ca38f4ab178'), // App Circuit Output Commitment 0
        BigNumber.from('0x4961cdcf881ce726a8df2f1c86697510'), // App Circuit Output Commitment 1
        BigNumber.from('0x2a60c9696759156b14a859d7636bee5de2f4a19b85e6ac9c060075453ad3704d'), // App Circuit VK Hash
        BigNumber.from('0x117f47f16b02c931f879674b41edfcc319341d7e3a5950a887216a557a1db682') // Commitment Public Input
      ]
    );
    assert.equal(result, true);
  });

  it('should pass on verify raw with real data', async () => {
    const values = [
      BigNumber.from('0x2d6ca9dc5f51406e358a694158054fa8173861cc05bb4cdb444240e5b5888dba'), // a0
      BigNumber.from('0x1b14c6562c6b57abb541adb8b7bcd75a8ce5d1228a28f0228802727a37cd65f8'), // a1

      BigNumber.from('0x011a5cfae06e386a11d80c54f6bb98d5bd04dea2bc0a3fccfa82d8222f76688f'), // b00
      BigNumber.from('0x0b8e2c859bcbbe39c1bc6d3ad437d6900d4a4dc563b3e0cb2d941ee6a70c1690'), // b01
      BigNumber.from('0x061c7739d940495c4cf5902db310dd4af559e03d08a901157d4c32da8d13f391'), // b10
      BigNumber.from('0x1797762d5fc38a017ec90d581ec3e7a0cc459f0225f75bdd609f4068217c1873'), // b11

      BigNumber.from('0x1596537d41a471914900275a70ad364fed4881273aca791881c4075d554368cc'), // c0
      BigNumber.from('0x1c645788f7b85fdff88c155dd32aa957163c86ee6b0331e1fd9b1f1642cfbfea'), // c1

      BigNumber.from('0x0423485ea151245e4195ff2df2bd089b2c925e1c2ab786ac78aa8eac7b9bffee'), // Commitment 0
      BigNumber.from('0x22a5d19e908e5bb687fa66362b93aaca4b9191a233486c97cc2b8ef4ece93823'), // Commitment 1

      BigNumber.from('0x117f47f16b02c931f879674b41edfcc319341d7e3a5950a887216a557a1db682'), // Commitment Public Input
      BigNumber.from('0x1cf83ccf4eb7942493c24a7ce1c7857efe087953dcbbcf085d3cb885d5285b09'), // Query Hash
      BigNumber.from('0x05826920570b43c036b2f33b827350bc86d8bd2dfe9a234e7ee6a758bf3c6aaf'), // SMT Root 
      BigNumber.from('0x0cfa2d22152d95266c0ccff01cedc28eafa8105dec9557125dbde2eceb5610dc'), // Aggregation VK Hash
      BigNumber.from('0x26ea6912d14adbed568e5ca38f4ab1784961cdcf881ce726a8df2f1c86697510'), // App Circuit Output Commitment 
      BigNumber.from('0x2a60c9696759156b14a859d7636bee5de2f4a19b85e6ac9c060075453ad3704d') // App Circuit VK Hash
    ];

    var hexValues = '';

    values.forEach((value) => {
      hexValues += value.toHexString().slice(2).padStart(64, '0');
    });

    const result = await contract.verifyRaw(hexToBytes('0x' + hexValues));

    assert.equal(result, true);
  });
});
