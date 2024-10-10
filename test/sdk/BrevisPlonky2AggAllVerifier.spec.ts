import { Fixture } from 'ethereum-waffle';
import { BigNumber, Wallet } from 'ethers';
import { ethers, waffle } from 'hardhat';
import { BrevisPlonky2AggAllVerifier } from '../../typechain';
import assert from 'assert';
import { hexToBytes } from '../util';

async function deployContract(admin: Wallet) {
  const _factory = await ethers.getContractFactory('BrevisPlonky2AggAllVerifier');
  const _contract = await _factory.connect(admin).deploy();
  return _contract;
}

describe('BrevisPlonky2AggAllVerifier test', async () => {
  function loadFixture<T>(fixture: Fixture<T>): Promise<T> {
    const provider = waffle.provider;
    return waffle.createFixtureLoader(provider.getWallets(), provider)(fixture);
  }

  async function fixture([admin]: Wallet[]) {
    const contract = await deployContract(admin);
    return { admin, contract };
  }

  let contract: BrevisPlonky2AggAllVerifier;
  let admin: Wallet;
  beforeEach(async () => {
    const res = await loadFixture(fixture);
    contract = res.contract;
    admin = res.admin;
  });

  it('should pass on true proof', async () => {
    const result = await contract.verifyProof(
      [
        BigNumber.from('0x26e54ead75a42369aef8a23c06ee360cd05aaa4b87093d5c2f49fbcad94d3fa1'), // a0
        BigNumber.from('0x2e88f5d72cd34f0e66e61636800b2aba46d905b77abba325725e0b62ee4a0f36'), // a1

        BigNumber.from('0x227da9309e298a6e0dba976d2943fba3e21c4938ec362ad2d195566b3552053f'), // b00
        BigNumber.from('0x12b26382fccc31d74331354ebede47deb1db24083639a40ad13f8713c4aa2d52'), // b01
        BigNumber.from('0x2bd221f23136de6b7843f5b91406604ba4c6bb759f6fc55b3eb2bd675be8fe94'), // b10
        BigNumber.from('0x2927236f491819b6a944ce421a9a0dd7e6b5a7d8bbecf329da4bb7fa1c9cfce8'), // b11

        BigNumber.from('0x1934eb7c3d161670aa57d2103b97b5f2fd834a34a0694abd2c19a784f7e5f398'), // c0
        BigNumber.from('0x19a4cc8afb7066a34406ede7ccab49b75b9fda64ba914b6d84f500b2b74d02b2') // c1
      ],
      [
        BigNumber.from('0x1bc9c2612678d3999b68752d9ae7f4e4ab5fcb1a4517de7b7579206c77134944'), // Commitment 0
        BigNumber.from('0x12a8ae56b61b25eaaf77566751e6f5a8c182000fdb4cf1fa6e8236d80757ef43') // Commitment 1
      ],
      [
        BigNumber.from('0x2754f60697d0ddf5a337d13e984f95d2497479ce86bc0e04f217d0408c7a9680'), // Commitment POK0
        BigNumber.from('0x217366aada9048eb697c3da4507b5e13eaa5897e09459d70e2e02bbf203d28b5') // Commitment POK1
      ],
      [
        BigNumber.from('0x2f7bff88baa27888fc2770d4c708c5e73524858f3af0aa0b1470dd71fe6b0465'), // Query Hash
        BigNumber.from('0x00000000000000000000000000000000ca3022d33d006b913ecfad08eb9d6dde'), // SMT Root 0
        BigNumber.from('0x00000000000000000000000000000000985e93fdd80ca6fadebcd2172d28c3f3'), // SMT Root 1
        BigNumber.from('0x0000000000000000000000000000000071712a0c5ec82685a9c47466e00386a7'), // Aggregation VK Hash
        BigNumber.from('0x00000000000000000000000000000000190bc97dbf0fb86d1ff47370783ffa07'), // App Circuit Output Commitment 0
        BigNumber.from('0x140390b7ddcb67336f3cc7269b552519940d8dbc54b5e24d4d7b772b0cb9b53b'), // App Circuit Output Commitment 1
        BigNumber.from('0x017a4240393a0eeffb0c2f9be5620a8d9080f8079c7df8064174b09341dbfe2b') // App Circuit VK Hash
      ]
    );
    assert.equal(result, true);
  });
});
