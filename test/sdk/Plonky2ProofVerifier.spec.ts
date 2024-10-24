import { Fixture } from 'ethereum-waffle';
import { BigNumber, Wallet } from 'ethers';
import { ethers, waffle } from 'hardhat';
import { Plonky2ProofVerifier } from '../../typechain';
import assert from 'assert';
import { hexToBytes } from '../util';
import { expect } from 'chai';

async function deployContract(admin: Wallet) {
  const _factory = await ethers.getContractFactory('Plonky2ProofVerifier');
  const _contract = await _factory.connect(admin).deploy();
  return _contract;
}

describe('Plonky2ProofVerifier test', async () => {
  function loadFixture<T>(fixture: Fixture<T>): Promise<T> {
    const provider = waffle.provider;
    return waffle.createFixtureLoader(provider.getWallets(), provider)(fixture);
  }

  async function fixture([admin]: Wallet[]) {
    const contract = await deployContract(admin);
    return { admin, contract };
  }

  let contract: Plonky2ProofVerifier;
  let admin: Wallet;
  beforeEach(async () => {
    const res = await loadFixture(fixture);
    contract = res.contract;
    admin = res.admin;
  });

  it('should pass on true proof', async () => {
    const result = await contract.verifyProof(
      [
        BigNumber.from('0x07226d1bf54c34fd4138d73d9950bff0b46cdf3f1cb23e42ff9bf7c794f7dba4'), // a0
        BigNumber.from('0x288dd90826fb4a52d34982b900962d864c0411df97b0f0819527b4ed318751fa'), // a1

        BigNumber.from('0x07b9d5f4ff9fc7991087d4ae352d2f780efc566cc9fc12ecd8e820d3c9ccdc78'), // b00
        BigNumber.from('0x0992e1f35f68ed4b5aa65a90b8fb754775820c54f8904f9317721375a9d80661'), // b01
        BigNumber.from('0x2e3c2a30bcbf1497723e5c2f94dd9ddc6745d35f6b2e198bdc5aa21fcac14094'), // b10
        BigNumber.from('0x280d9491d5e100788a2b6155aabf1724f10b85d751f769289d7f4eef442877be'), // b11

        BigNumber.from('0x2c74fcdee3a4bdfe216e9fee7f87a4f4a9efe4b7e19ca5ee9c1a728586657ded'), // c0
        BigNumber.from('0x10b7d6d1fe7ac88670f9f59aa3c0bcbb18fc207b8f84c40116029a3d84823f6c') // c1
      ],
      [
        BigNumber.from('0xc3e7285e5c53aa4a73512a066762a84335b18424736c6f2ec276dbbd9fb15c'), // Commitment 0
        BigNumber.from('0x1d658fecdcf74c8ed73511751f8d04bb35837b7e835269e3c11fe6aed35e1989') // Commitment 1
      ],
      [
        BigNumber.from('0x15e54e593fb1f65822f04b76c9e776325fbc27705abe65073c112fae7f4b3bd6'), // Commitment POK0
        BigNumber.from('0x2a318cbb0bca8b760f758e9c011108ce1fea07856d16dcb18d49e818857b22d2') // Commitment POK1
      ],
      [
        BigNumber.from('0x1c80a4a07adc00e26449991f8742a2882f0f89e5e4ef453a118a5c5db1ef22d7'), // Input commitments root
        BigNumber.from('0x00000000000000000000000000000000ca3022d33d006b913ecfad08eb9d6dde'), // SMT Root 0
        BigNumber.from('0x00000000000000000000000000000000985e93fdd80ca6fadebcd2172d28c3f3'), // SMT Root 1
        BigNumber.from('0x0000000000000000000000000000000071712a0c5ec82685a9c47466e00386a7'), // App Circuit Output Commitment 0
        BigNumber.from('0x00000000000000000000000000000000190bc97dbf0fb86d1ff47370783ffa07'), // App Circuit Output Commitment 1
        BigNumber.from('0x091be53f1dc5af98b496abfdc85ef95d9c31b71a3ae28a7d69ba52dd51095718'), // CircuitDigest
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
        BigNumber.from('0x2f7bff88baa27888fc2770d4c708c5e73524858f3af0aa0b1470dd71fe6b0465'), // Input commitments root
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
      BigNumber.from('0x07226d1bf54c34fd4138d73d9950bff0b46cdf3f1cb23e42ff9bf7c794f7dba4'), // a0
      BigNumber.from('0x288dd90826fb4a52d34982b900962d864c0411df97b0f0819527b4ed318751fa'), // a1

      BigNumber.from('0x07b9d5f4ff9fc7991087d4ae352d2f780efc566cc9fc12ecd8e820d3c9ccdc78'), // b00
      BigNumber.from('0x0992e1f35f68ed4b5aa65a90b8fb754775820c54f8904f9317721375a9d80661'), // b01
      BigNumber.from('0x2e3c2a30bcbf1497723e5c2f94dd9ddc6745d35f6b2e198bdc5aa21fcac14094'), // b10
      BigNumber.from('0x280d9491d5e100788a2b6155aabf1724f10b85d751f769289d7f4eef442877be'), // b11

      BigNumber.from('0x2c74fcdee3a4bdfe216e9fee7f87a4f4a9efe4b7e19ca5ee9c1a728586657ded'), // c0
      BigNumber.from('0x10b7d6d1fe7ac88670f9f59aa3c0bcbb18fc207b8f84c40116029a3d84823f6c'), // c1

      BigNumber.from('0xc3e7285e5c53aa4a73512a066762a84335b18424736c6f2ec276dbbd9fb15c'), // Commitment 0
      BigNumber.from('0x1d658fecdcf74c8ed73511751f8d04bb35837b7e835269e3c11fe6aed35e1989'), // Commitment 1

      BigNumber.from('0x15e54e593fb1f65822f04b76c9e776325fbc27705abe65073c112fae7f4b3bd6'), // Commitment POK0
      BigNumber.from('0x2a318cbb0bca8b760f758e9c011108ce1fea07856d16dcb18d49e818857b22d2'), // Commitment POK1

      BigNumber.from('0x1c80a4a07adc00e26449991f8742a2882f0f89e5e4ef453a118a5c5db1ef22d7'), // Input commitments root
      BigNumber.from('0xca3022d33d006b913ecfad08eb9d6dde985e93fdd80ca6fadebcd2172d28c3f3'), // SMT Root
      BigNumber.from('0x71712a0c5ec82685a9c47466e00386a7190bc97dbf0fb86d1ff47370783ffa07'), // App Circuit Output Commitment
      BigNumber.from('0x091be53f1dc5af98b496abfdc85ef95d9c31b71a3ae28a7d69ba52dd51095718'), // CircuitDigest
      BigNumber.from('0x17d1e8686c170d21dfd9870eaba43ff8fbaeefb0aca084a0dd2e308f0e5b3b7e') // DummyCommitment
    ];

    var hexValues = '';

    values.forEach((value) => {
      hexValues += value.toHexString().slice(2).padStart(64, '0');
    });

    console.log('hexValues: ', hexValues);
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
      BigNumber.from('0x0') // DummyCommitment
    ];

    var hexValues = '';

    values.forEach((value) => {
      hexValues += value.toHexString().slice(2).padStart(64, '0');
    });

    const result = contract.verifyRaw(hexToBytes('0x' + hexValues));
    await expect(result).reverted;
  });
});
