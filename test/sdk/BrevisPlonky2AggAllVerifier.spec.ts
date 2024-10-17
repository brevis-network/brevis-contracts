import { Fixture } from 'ethereum-waffle';
import { BigNumber, Wallet } from 'ethers';
import { ethers, waffle } from 'hardhat';
import { BrevisPlonky2AggAllVerifier } from '../../typechain';
import assert from 'assert';
import { hexToBytes } from '../util';
import { expect } from 'chai';

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
        BigNumber.from('0x067c89cf20e457d1d5008139e388c95dfaef74ec927f648bc7048116ad879df9'), // a0
        BigNumber.from('0x2b490bb46d91cad7667623bf7cecac1c7f6ae16712f82e87148157aeda32c046'), // a1

        BigNumber.from('0x0b3445bc871b565e532630b0b2e793a3a7270b83f8bbcb131fcff19330c7ff64'), // b00
        BigNumber.from('0x163b2971ee0eb0ebf63b16c4cb2502a36d735913fe0268fc8b424a60b6dc5c83'), // b01
        BigNumber.from('0x2134bc80369d748a29a6d032454354ac362018f3179fd7411c24825ef2afffc8'), // b10
        BigNumber.from('0x2583b3a3f1eb5def4920674a065e7d8eafaaeccbee6322bfd8a6e2620d07c1bb'), // b11

        BigNumber.from('0x1eda53e3a3c1eee7a66b84a1bca1f29cb6a1f3fd5d4d2ac48a69cf2bff39efce'), // c0
        BigNumber.from('0x17cdaeb9dc2dc92354098cf5dfb431538db95b6d8ff005ba46516c3819699cf5') // c1
      ],
      [
        BigNumber.from('0x34a56eb26a24a6174354063f680660f0a0f7e710ce6fdcff66b111d401f271'), // Commitment 0
        BigNumber.from('0x2c4e04b63a814d8e801a8ae0002ea66bde542f4fa151ee572628aa0267b5fffb') // Commitment 1
      ],
      [
        BigNumber.from('0x2daf786e373ddd39333c53a3965bc98e78d7259096167340e41db1a46546024f'), // Commitment POK0
        BigNumber.from('0x0abee73cf9c9eaaaf22818c48d6a247740996d02aa53a61d5838d63abea22df8') // Commitment POK1
      ],
      [
        BigNumber.from('0x1c80a4a07adc00e26449991f8742a2882f0f89e5e4ef453a118a5c5db1ef22d7'), // Query Hash
        BigNumber.from('0x00000000000000000000000000000000ca3022d33d006b913ecfad08eb9d6dde'), // SMT Root 0
        BigNumber.from('0x00000000000000000000000000000000985e93fdd80ca6fadebcd2172d28c3f3'), // SMT Root 1
        BigNumber.from('0x0000000000000000000000000000000071712a0c5ec82685a9c47466e00386a7'), // App Circuit Output Commitment 0
        BigNumber.from('0x00000000000000000000000000000000190bc97dbf0fb86d1ff47370783ffa07'), // App Circuit Output Commitment 1
        BigNumber.from('0x2c86965e6e0d3878f8bc667d9e3ff0043f07a125d9ea24c3f83df9519e7d57eb'), // CircuitDigest
        BigNumber.from('0x17d1e8686c170d21dfd9870eaba43ff8fbaeefb0aca084a0dd2e308f0e5b3b7e') // DummyCommitment
      ]
    );
    assert.equal(result, true);
  });

  it('should revert on false proof', async () => {
    const result = contract.verifyProof(
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
        BigNumber.from('0x0000000000000000000000000000000071712a0c5ec82685a9c47466e00386a7'), // App Circuit Output Commitment 0
        BigNumber.from('0x00000000000000000000000000000000190bc97dbf0fb86d1ff47370783ffa07'), // App Circuit Output Commitment 1
        BigNumber.from('0x140390b7ddcb67336f3cc7269b552519940d8dbc54b5e24d4d7b772b0cb9b53b'), // CircuitDigest
        BigNumber.from('0x0') // DummyCommitment
      ]
    );

    await expect(result).reverted;
  });

  it('should pass on verify raw proof', async () => {
    const values = [
      BigNumber.from('0x067c89cf20e457d1d5008139e388c95dfaef74ec927f648bc7048116ad879df9'), // a0
      BigNumber.from('0x2b490bb46d91cad7667623bf7cecac1c7f6ae16712f82e87148157aeda32c046'), // a1

      BigNumber.from('0x0b3445bc871b565e532630b0b2e793a3a7270b83f8bbcb131fcff19330c7ff64'), // b00
      BigNumber.from('0x163b2971ee0eb0ebf63b16c4cb2502a36d735913fe0268fc8b424a60b6dc5c83'), // b01
      BigNumber.from('0x2134bc80369d748a29a6d032454354ac362018f3179fd7411c24825ef2afffc8'), // b10
      BigNumber.from('0x2583b3a3f1eb5def4920674a065e7d8eafaaeccbee6322bfd8a6e2620d07c1bb'), // b11

      BigNumber.from('0x1eda53e3a3c1eee7a66b84a1bca1f29cb6a1f3fd5d4d2ac48a69cf2bff39efce'), // c0
      BigNumber.from('0x17cdaeb9dc2dc92354098cf5dfb431538db95b6d8ff005ba46516c3819699cf5'), // c1

      BigNumber.from('0x34a56eb26a24a6174354063f680660f0a0f7e710ce6fdcff66b111d401f271'), // Commitment 0
      BigNumber.from('0x2c4e04b63a814d8e801a8ae0002ea66bde542f4fa151ee572628aa0267b5fffb'), // Commitment 1

      BigNumber.from('0x2daf786e373ddd39333c53a3965bc98e78d7259096167340e41db1a46546024f'), // Commitment POK0
      BigNumber.from('0x0abee73cf9c9eaaaf22818c48d6a247740996d02aa53a61d5838d63abea22df8'), // Commitment POK1

      BigNumber.from('0x1c80a4a07adc00e26449991f8742a2882f0f89e5e4ef453a118a5c5db1ef22d7'), // Query Hash
      BigNumber.from('0xca3022d33d006b913ecfad08eb9d6dde985e93fdd80ca6fadebcd2172d28c3f3'), // SMT Root
      BigNumber.from('0x71712a0c5ec82685a9c47466e00386a7190bc97dbf0fb86d1ff47370783ffa07'), // App Circuit Output Commitment
      BigNumber.from('0x2c86965e6e0d3878f8bc667d9e3ff0043f07a125d9ea24c3f83df9519e7d57eb'), // CircuitDigest
      BigNumber.from('0x17d1e8686c170d21dfd9870eaba43ff8fbaeefb0aca084a0dd2e308f0e5b3b7e')  // DummyCommitment
    ];

    var hexValues = '';

    values.forEach((value) => {
      hexValues += value.toHexString().slice(2).padStart(64, '0');
    });

    const result = await contract.verifyRaw(hexToBytes('0x' + hexValues));

    assert.equal(result, true);
  });

  it('should revert on verify raw proof', async () => {
    const values = [
      BigNumber.from('0x26e54ead75a42369aef8a23c06ee360cd05aaa4b87093d5c2f49fbcad94d3fa1'), // a0
      BigNumber.from('0x2e88f5d72cd34f0e66e61636800b2aba46d905b77abba325725e0b62ee4a0f36'), // a1

      BigNumber.from('0x227da9309e298a6e0dba976d2943fba3e21c4938ec362ad2d195566b3552053f'), // b00
      BigNumber.from('0x12b26382fccc31d74331354ebede47deb1db24083639a40ad13f8713c4aa2d52'), // b01
      BigNumber.from('0x2bd221f23136de6b7843f5b91406604ba4c6bb759f6fc55b3eb2bd675be8fe94'), // b10
      BigNumber.from('0x2927236f491819b6a944ce421a9a0dd7e6b5a7d8bbecf329da4bb7fa1c9cfce8'), // b11

      BigNumber.from('0x1934eb7c3d161670aa57d2103b97b5f2fd834a34a0694abd2c19a784f7e5f398'), // c0
      BigNumber.from('0x19a4cc8afb7066a34406ede7ccab49b75b9fda64ba914b6d84f500b2b74d02b2'), // c1

      BigNumber.from('0x1bc9c2612678d3999b68752d9ae7f4e4ab5fcb1a4517de7b7579206c77134944'), // Commitment 0
      BigNumber.from('0x12a8ae56b61b25eaaf77566751e6f5a8c182000fdb4cf1fa6e8236d80757ef43'), // Commitment 1

      BigNumber.from('0x2754f60697d0ddf5a337d13e984f95d2497479ce86bc0e04f217d0408c7a9680'), // Commitment POK0
      BigNumber.from('0x217366aada9048eb697c3da4507b5e13eaa5897e09459d70e2e02bbf203d28b5'), // Commitment POK1

      BigNumber.from('0x2f7bff88baa27888fc2770d4c708c5e73524858f3af0aa0b1470dd71fe6b0465'), // Query Hash
      BigNumber.from('0xca3022d33d006b913ecfad08eb9d6dde985e93fdd80ca6fadebcd2172d28c3f3'), // SMT Root
      BigNumber.from('0x71712a0c5ec82685a9c47466e00386a7190bc97dbf0fb86d1ff47370783ffa07'), // App Circuit Output Commitment
      BigNumber.from('0x140390b7ddcb67336f3cc7269b552519940d8dbc54b5e24d4d7b772b0cb9b53b'), // CircuitDigest
      BigNumber.from('0x0')  // DummyCommitment
    ];

    var hexValues = '';

    values.forEach((value) => {
      hexValues += value.toHexString().slice(2).padStart(64, '0');
    });

    const result = contract.verifyRaw(hexToBytes('0x' + hexValues));
    await expect(result).reverted;
  });
});
