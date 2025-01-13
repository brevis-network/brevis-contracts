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
        BigNumber.from('0x11a8202d2c5d6164ebcadd1a75e68969e77cb9ddc8c7ad0f3d7e712bfce5a655'), // a0
        BigNumber.from('0x15ec0b6b3dc85c84b808ca794fcff55cfa8d27bc7de1a54d5b05263767dd188b'), // a1

        BigNumber.from('0x2580c251196ebd925d1d64134b28f0264c01cac35211290d334945a55a25f22d'), // b00
        BigNumber.from('0x18fe520347b489968fac45e64647b6174d9559a59239025f17ff7af17f2a0d09'), // b01
        BigNumber.from('0x2198c227a0cb2807a930288e2d8f67533cc41f3add300c4f0c159e36c2b56527'), // b10
        BigNumber.from('0x15124003867aa1096021f0f005df6f7f2c6c3a738ee9decf3468140353cdd816'), // b11

        BigNumber.from('0x17c27a49ddbdeab2ed5705f215f0614cfff2e72775825798d3c9130adbac2aac'), // c0
        BigNumber.from('0x277fc5a919f2240d7c2c428e1b42ec86cefffc3b2cab52655537d0351c11482b') // c1
      ],
      [
        BigNumber.from('0x126b347c94dc59e0082d906b9be09a5080936981b83db32a389674b38aa06a2c'), // Commitment 0
        BigNumber.from('0x0a949c0f7aecf123f7f3c4b1cf2330f487ad5d751c9f714794e95fc4dd532472') // Commitment 1
      ],
      [
        BigNumber.from('0x077b5368fedd971873842f415bcc2b65ab30a7cc7396c9e24f421c94c142e1bf'), // Commitment POK0
        BigNumber.from('0x0f457f3c1babbd9b3824c4e2b3b49ff7dc7b5040c7bb3b8c8dce2190b6a71303') // Commitment POK1
      ],
      [
        BigNumber.from('0x1c80a4a07adc00e26449991f8742a2882f0f89e5e4ef453a118a5c5db1ef22d7'), // Input commitments root
        BigNumber.from('0x00000000000000000000000000000000ca3022d33d006b913ecfad08eb9d6dde'), // SMT Root 0
        BigNumber.from('0x00000000000000000000000000000000985e93fdd80ca6fadebcd2172d28c3f3'), // SMT Root 1
        BigNumber.from('0x0000000000000000000000000000000071712a0c5ec82685a9c47466e00386a7'), // App Circuit Output Commitment 0
        BigNumber.from('0x00000000000000000000000000000000190bc97dbf0fb86d1ff47370783ffa07'), // App Circuit Output Commitment 1
        BigNumber.from('0x17fb56a419275d42fd16d02e2200b706daaaf602031470d809ddecb827ef8cef'), // CircuitDigest
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
      BigNumber.from('0x11a8202d2c5d6164ebcadd1a75e68969e77cb9ddc8c7ad0f3d7e712bfce5a655'), // a0
      BigNumber.from('0x15ec0b6b3dc85c84b808ca794fcff55cfa8d27bc7de1a54d5b05263767dd188b'), // a1

      BigNumber.from('0x2580c251196ebd925d1d64134b28f0264c01cac35211290d334945a55a25f22d'), // b00
      BigNumber.from('0x18fe520347b489968fac45e64647b6174d9559a59239025f17ff7af17f2a0d09'), // b01
      BigNumber.from('0x2198c227a0cb2807a930288e2d8f67533cc41f3add300c4f0c159e36c2b56527'), // b10
      BigNumber.from('0x15124003867aa1096021f0f005df6f7f2c6c3a738ee9decf3468140353cdd816'), // b11

      BigNumber.from('0x17c27a49ddbdeab2ed5705f215f0614cfff2e72775825798d3c9130adbac2aac'), // c0
      BigNumber.from('0x277fc5a919f2240d7c2c428e1b42ec86cefffc3b2cab52655537d0351c11482b'), // c1

      BigNumber.from('0x126b347c94dc59e0082d906b9be09a5080936981b83db32a389674b38aa06a2c'), // Commitment 0
      BigNumber.from('0x0a949c0f7aecf123f7f3c4b1cf2330f487ad5d751c9f714794e95fc4dd532472'), // Commitment 1

      BigNumber.from('0x077b5368fedd971873842f415bcc2b65ab30a7cc7396c9e24f421c94c142e1bf'), // Commitment POK0
      BigNumber.from('0x0f457f3c1babbd9b3824c4e2b3b49ff7dc7b5040c7bb3b8c8dce2190b6a71303'), // Commitment POK1

      BigNumber.from('0x1c80a4a07adc00e26449991f8742a2882f0f89e5e4ef453a118a5c5db1ef22d7'), // Input commitments root
      BigNumber.from('0xca3022d33d006b913ecfad08eb9d6dde985e93fdd80ca6fadebcd2172d28c3f3'), // SMT Root
      BigNumber.from('0x71712a0c5ec82685a9c47466e00386a7190bc97dbf0fb86d1ff47370783ffa07'), // App Circuit Output Commitment
      BigNumber.from('0x17fb56a419275d42fd16d02e2200b706daaaf602031470d809ddecb827ef8cef'), // CircuitDigest
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
