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

  /*
  2025-01-10 10:24:22.940 |INFO | proof.go:65: proofData.A[0]: 0x1e79dabba7f5950609f4b916b6f9805ebe05c3ac59ad5f7ac81a99aa2f0b9f6d
2025-01-10 10:24:22.940 |INFO | proof.go:66: proofData.A[1]: 0x04fc6481756d7f6f51589d30552dcdac71c968b0048b747aab1b4b4ab612348e
2025-01-10 10:24:22.940 |INFO | proof.go:68: proofData.B[0][0]: 0x21d76d06f3c74f27fdf35cb08583ac687f535c45fc645babd7795cc316b663fe
2025-01-10 10:24:22.940 |INFO | proof.go:69: proofData.B[0][1]: 0x1ca45f1beefede0f6d489faa445cdcdd9b142d11e61d7146670ca643ee8582a6
2025-01-10 10:24:22.940 |INFO | proof.go:70: proofData.B[1][0]: 0x16bc296404452b10ff019e05b7e71f0ace3d20058da396485f6213c327eb89df
2025-01-10 10:24:22.940 |INFO | proof.go:71: proofData.B[1][1]: 0x27966ec7387942710b9ce02e6df114112e5daade6b8c2f86f81fbdecc68fd49c
2025-01-10 10:24:22.940 |INFO | proof.go:73: proofData.C[0]: 0x1b986399ff3b14ab68e3992fd8028bd528001b18a1f85bc34d03092a2325d578
2025-01-10 10:24:22.940 |INFO | proof.go:74: proofData.C[1]: 0x22cdeb6d2e969ca450e4651a0608b66676c039b2d79d6d9b505566b745c24f18
2025-01-10 10:24:22.940 |INFO | proof.go:76: proofData.Commitment[0]: 0x1139427aee0dcc664fd1a30f62c06ed24d0d272d588be07c98aa0383c56d61c0
2025-01-10 10:24:22.940 |INFO | proof.go:77: proofData.Commitment[1]: 0x1ba787645286a0388b5d40ed435bf73ebb8ecc99c3264a491148d28d2a213af4
2025-01-10 10:24:22.940 |INFO | proof.go:78: proofData.CommitmentPok[0]: 0x2f7f2a4c4e716223d035f2cfcd110fa142156c7068a2cdd5b51dda7067ad2479
2025-01-10 10:24:22.940 |INFO | proof.go:79: proofData.CommitmentPok[1]: 0x26fe3fd6a6942a8295b5119489a206094d20bb9108dc47f499244512710693cb
2025-01-10 10:24:22.940 |INFO | proof.go:89: witness_0: 0x272af5fc782b1cea7ab2f343436a19357e93a98139d993143e7b63e04a991661
2025-01-10 10:24:22.940 |INFO | proof.go:89: witness_1: 0x00000000000000000000000000000000f3dea08dd248fc49b017e6477aefa2e6
2025-01-10 10:24:22.940 |INFO | proof.go:89: witness_2: 0x000000000000000000000000000000009f8ef7aa194ada537c24db45d1c18e46
2025-01-10 10:24:22.940 |INFO | proof.go:89: witness_3: 0x00000000000000000000000000000000121807149b4dcfd90b8b48396d92b435
2025-01-10 10:24:22.940 |INFO | proof.go:89: witness_4: 0x000000000000000000000000000000004093b58c97789c4ad60aa0b767a73001
2025-01-10 10:24:22.940 |INFO | proof.go:89: witness_5: 0x26566cdd80f2f4d2a2fb0f520603f3669dd842d244e9c75a8d96c8fc3a8504cb
2025-01-10 10:24:22.940 |INFO | proof.go:89: witness_6: 0x27520dda6fc884a5efbd1cdd71cb38e7ab9a36a861aedc4df638d09ee3d705d9
2025-01-10 10:24:22.940 |INFO | proof.go:89: witness_7: 0x08de120895f5885a6fc710817baf78dc8368110c42d8ca7a58152634afb404b6
2025-01-10 10:24:22.940 |INFO | export_agg_all_solidity_test.go:47: res: 0x1e79dabba7f5950609f4b916b6f9805ebe05c3ac59ad5f7ac81a99aa2f0b9f6d,0x04fc6481756d7f6f51589d30552dcdac71c968b0048b747aab1b4b4ab612348e,0x21d76d06f3c74f27fdf35cb08583ac687f535c45fc645babd7795cc316b663fe,0x1ca45f1beefede0f6d489faa445cdcdd9b142d11e61d7146670ca643ee8582a6,0x16bc296404452b10ff019e05b7e71f0ace3d20058da396485f6213c327eb89df,0x27966ec7387942710b9ce02e6df114112e5daade6b8c2f86f81fbdecc68fd49c,0x1b986399ff3b14ab68e3992fd8028bd528001b18a1f85bc34d03092a2325d578,0x22cdeb6d2e969ca450e4651a0608b66676c039b2d79d6d9b505566b745c24f18,0x1139427aee0dcc664fd1a30f62c06ed24d0d272d588be07c98aa0383c56d61c0,0x1ba787645286a0388b5d40ed435bf73ebb8ecc99c3264a491148d28d2a213af4,0x2f7f2a4c4e716223d035f2cfcd110fa142156c7068a2cdd5b51dda7067ad2479,0x26fe3fd6a6942a8295b5119489a206094d20bb9108dc47f499244512710693cb,0x272af5fc782b1cea7ab2f343436a19357e93a98139d993143e7b63e04a991661,0x00000000000000000000000000000000f3dea08dd248fc49b017e6477aefa2e6,0x000000000000000000000000000000009f8ef7aa194ada537c24db45d1c18e46,0x00000000000000000000000000000000121807149b4dcfd90b8b48396d92b435,0x000000000000000000000000000000004093b58c97789c4ad60aa0b767a73001,0x26566cdd80f2f4d2a2fb0f520603f3669dd842d244e9c75a8d96c8fc3a8504cb,0x27520dda6fc884a5efbd1cdd71cb38e7ab9a36a861aedc4df638d09ee3d705d9,0x08de120895f5885a6fc710817baf78dc8368110c42d8ca7a58152634afb404b6,
--- PASS: TestExportSolidityTestData (0.00s)
   */
  it('should pass on true proof', async () => {
    const result = await contract.verifyProof(
      [
        BigNumber.from('0x1e79dabba7f5950609f4b916b6f9805ebe05c3ac59ad5f7ac81a99aa2f0b9f6d'), // a0
        BigNumber.from('0x04fc6481756d7f6f51589d30552dcdac71c968b0048b747aab1b4b4ab612348e'), // a1

        BigNumber.from('0x21d76d06f3c74f27fdf35cb08583ac687f535c45fc645babd7795cc316b663fe'), // b00
        BigNumber.from('0x1ca45f1beefede0f6d489faa445cdcdd9b142d11e61d7146670ca643ee8582a6'), // b01
        BigNumber.from('0x16bc296404452b10ff019e05b7e71f0ace3d20058da396485f6213c327eb89df'), // b10
        BigNumber.from('0x27966ec7387942710b9ce02e6df114112e5daade6b8c2f86f81fbdecc68fd49c'), // b11

        BigNumber.from('0x1b986399ff3b14ab68e3992fd8028bd528001b18a1f85bc34d03092a2325d578'), // c0
        BigNumber.from('0x22cdeb6d2e969ca450e4651a0608b66676c039b2d79d6d9b505566b745c24f18') // c1
      ],
      [
        BigNumber.from('0x1139427aee0dcc664fd1a30f62c06ed24d0d272d588be07c98aa0383c56d61c0'), // Commitment 0
        BigNumber.from('0x1ba787645286a0388b5d40ed435bf73ebb8ecc99c3264a491148d28d2a213af4') // Commitment 1
      ],
      [
        BigNumber.from('0x2f7f2a4c4e716223d035f2cfcd110fa142156c7068a2cdd5b51dda7067ad2479'), // Commitment POK0
        BigNumber.from('0x26fe3fd6a6942a8295b5119489a206094d20bb9108dc47f499244512710693cb') // Commitment POK1
      ],
      [
        BigNumber.from('0x272af5fc782b1cea7ab2f343436a19357e93a98139d993143e7b63e04a991661'), // Input commitments root
        BigNumber.from('0x00000000000000000000000000000000f3dea08dd248fc49b017e6477aefa2e6'), // SMT Root 0
        BigNumber.from('0x000000000000000000000000000000009f8ef7aa194ada537c24db45d1c18e46'), // SMT Root 1
        BigNumber.from('0x00000000000000000000000000000000121807149b4dcfd90b8b48396d92b435'), // App Circuit Output Commitment 0
        BigNumber.from('0x000000000000000000000000000000004093b58c97789c4ad60aa0b767a73001'), // App Circuit Output Commitment 1
        BigNumber.from('0x26566cdd80f2f4d2a2fb0f520603f3669dd842d244e9c75a8d96c8fc3a8504cb'), // CircuitDigest
        BigNumber.from('0x27520dda6fc884a5efbd1cdd71cb38e7ab9a36a861aedc4df638d09ee3d705d9'), // DummyCommitment
        BigNumber.from('0x08de120895f5885a6fc710817baf78dc8368110c42d8ca7a58152634afb404b6') // app vk hash
      ]
    );
    assert.equal(result, true);
  });

  it('should revert on false proof', async () => {
    const result = contract.verifyProof(
        [
          BigNumber.from('0x1e79dabba7f5950609f4b916b6f9805ebe05c3ac59ad5f7ac81a99aa2f0b9f6d'), // a0
          BigNumber.from('0x04fc6481756d7f6f51589d30552dcdac71c968b0048b747aab1b4b4ab612348e'), // a1

          BigNumber.from('0x21d76d06f3c74f27fdf35cb08583ac687f535c45fc645babd7795cc316b663fe'), // b00
          BigNumber.from('0x1ca45f1beefede0f6d489faa445cdcdd9b142d11e61d7146670ca643ee8582a6'), // b01
          BigNumber.from('0x16bc296404452b10ff019e05b7e71f0ace3d20058da396485f6213c327eb89df'), // b10
          BigNumber.from('0x27966ec7387942710b9ce02e6df114112e5daade6b8c2f86f81fbdecc68fd49c'), // b11

          BigNumber.from('0x1b986399ff3b14ab68e3992fd8028bd528001b18a1f85bc34d03092a2325d578'), // c0
          BigNumber.from('0x22cdeb6d2e969ca450e4651a0608b66676c039b2d79d6d9b505566b745c24f18') // c1
        ],
        [
          BigNumber.from('0x1139427aee0dcc664fd1a30f62c06ed24d0d272d588be07c98aa0383c56d61c0'), // Commitment 0
          BigNumber.from('0x1ba787645286a0388b5d40ed435bf73ebb8ecc99c3264a491148d28d2a213af4') // Commitment 1
        ],
        [
          BigNumber.from('0x2f7f2a4c4e716223d035f2cfcd110fa142156c7068a2cdd5b51dda7067ad2479'), // Commitment POK0
          BigNumber.from('0x26fe3fd6a6942a8295b5119489a206094d20bb9108dc47f499244512710693cb') // Commitment POK1
        ],
        [
          BigNumber.from('0x272af5fc782b1cea7ab2f343436a19357e93a98139d993143e7b63e04a991661'), // Input commitments root
          BigNumber.from('0x00000000000000000000000000000000f3dea08dd248fc49b017e6477aefa2e6'), // SMT Root 0
          BigNumber.from('0x000000000000000000000000000000009f8ef7aa194ada537c24db45d1c18e46'), // SMT Root 1
          BigNumber.from('0x00000000000000000000000000000000121807149b4dcfd90b8b48396d92b435'), // App Circuit Output Commitment 0
          BigNumber.from('0x000000000000000000000000000000004093b58c97789c4ad60aa0b767a73001'), // App Circuit Output Commitment 1
          BigNumber.from('0x26566cdd80f2f4d2a2fb0f520603f3669dd842d244e9c75a8d96c8fc3a8504cb'), // CircuitDigest
          BigNumber.from('0x27520dda6fc884a5efbd1cdd71cb38e7ab9a36a861aedc4df638d09ee3d705d9'), // DummyCommitment
          BigNumber.from('0x08de120895f5885a6fc710817baf78dc8368110c42d8ca7a58152634afb404b7') // app vk hash
        ]
    );

    await expect(result).reverted;
  });

  it('should pass on verify raw proof', async () => {
    const values = [
      BigNumber.from('0x1e79dabba7f5950609f4b916b6f9805ebe05c3ac59ad5f7ac81a99aa2f0b9f6d'), // a0
      BigNumber.from('0x04fc6481756d7f6f51589d30552dcdac71c968b0048b747aab1b4b4ab612348e'), // a1

      BigNumber.from('0x21d76d06f3c74f27fdf35cb08583ac687f535c45fc645babd7795cc316b663fe'), // b00
      BigNumber.from('0x1ca45f1beefede0f6d489faa445cdcdd9b142d11e61d7146670ca643ee8582a6'), // b01
      BigNumber.from('0x16bc296404452b10ff019e05b7e71f0ace3d20058da396485f6213c327eb89df'), // b10
      BigNumber.from('0x27966ec7387942710b9ce02e6df114112e5daade6b8c2f86f81fbdecc68fd49c'), // b11

      BigNumber.from('0x1b986399ff3b14ab68e3992fd8028bd528001b18a1f85bc34d03092a2325d578'), // c0
      BigNumber.from('0x22cdeb6d2e969ca450e4651a0608b66676c039b2d79d6d9b505566b745c24f18'), // c1

      BigNumber.from('0x1139427aee0dcc664fd1a30f62c06ed24d0d272d588be07c98aa0383c56d61c0'), // Commitment 0
      BigNumber.from('0x1ba787645286a0388b5d40ed435bf73ebb8ecc99c3264a491148d28d2a213af4'), // Commitment 1

      BigNumber.from('0x2f7f2a4c4e716223d035f2cfcd110fa142156c7068a2cdd5b51dda7067ad2479'), // Commitment POK0
      BigNumber.from('0x26fe3fd6a6942a8295b5119489a206094d20bb9108dc47f499244512710693cb'), // Commitment POK1

      BigNumber.from('0x272af5fc782b1cea7ab2f343436a19357e93a98139d993143e7b63e04a991661'), // Input commitments root
      BigNumber.from('0xf3dea08dd248fc49b017e6477aefa2e69f8ef7aa194ada537c24db45d1c18e46'), // SMT Root
      BigNumber.from('0x121807149b4dcfd90b8b48396d92b4354093b58c97789c4ad60aa0b767a73001'), // App Circuit Output Commitment
      BigNumber.from('0x26566cdd80f2f4d2a2fb0f520603f3669dd842d244e9c75a8d96c8fc3a8504cb'), // CircuitDigest
      BigNumber.from('0x27520dda6fc884a5efbd1cdd71cb38e7ab9a36a861aedc4df638d09ee3d705d9'), // DummyCommitment
      BigNumber.from('0x08de120895f5885a6fc710817baf78dc8368110c42d8ca7a58152634afb404b6') // app vk hash
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
      BigNumber.from('0x0'), // DummyCommitment
      BigNumber.from('0x0') // app vk hash
    ];

    var hexValues = '';

    values.forEach((value) => {
      hexValues += value.toHexString().slice(2).padStart(64, '0');
    });

    const result = contract.verifyRaw(hexToBytes('0x' + hexValues));
    await expect(result).reverted;
  });
});
