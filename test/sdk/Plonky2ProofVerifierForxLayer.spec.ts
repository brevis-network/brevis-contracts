import { Fixture } from 'ethereum-waffle';
import { BigNumber, Wallet } from 'ethers';
import { ethers, waffle } from 'hardhat';
import { Plonky2ProofVerifierForxLayer } from '../../typechain';
import assert from 'assert';
import { hexToBytes } from '../util';
import { expect } from 'chai';

async function deployContract(admin: Wallet) {
  const _factory = await ethers.getContractFactory('Plonky2ProofVerifierForxLayer');
  const _contract = await _factory.connect(admin).deploy();
  return _contract;
}

/*
2024-12-04 03:44:47.970 |INFO | proof.go:65: proofData.A[0]: 0x2c4537e488e3cb6118660dbe93d1a3da817cba90476c0da97b571c82a421814a
2024-12-04 03:44:47.970 |INFO | proof.go:66: proofData.A[1]: 0x0c3367b3277afd31ef0aac035f55af292549526a5a74089a1e4170a64c4a628a
2024-12-04 03:44:47.970 |INFO | proof.go:68: proofData.B[0][0]: 0x019cd8524a22bc985ec5cc43b5e5563754e7d2e8a2ad375cca46209892213d55
2024-12-04 03:44:47.970 |INFO | proof.go:69: proofData.B[0][1]: 0x27dcd502799a20ef8b02c542d2e276b45465430b4046107bb8d44a1702e2a3d6
2024-12-04 03:44:47.970 |INFO | proof.go:70: proofData.B[1][0]: 0x1b7a471a601b8276ab3c9107668af24589b682b0406a408bb438c73df08254d5
2024-12-04 03:44:47.970 |INFO | proof.go:71: proofData.B[1][1]: 0x08c36b4c3ed96c508cf925c4adcfd2aa4c39c1c3d98394cb8b15dd842309f817
2024-12-04 03:44:47.970 |INFO | proof.go:73: proofData.C[0]: 0x0d96799375aafaf3c68551646bb8267df5492ce93c42969214cf6c008324fb4b
2024-12-04 03:44:47.970 |INFO | proof.go:74: proofData.C[1]: 0x026ef48e77bbad429d472a8471d862bf7fe1a44904ec9878b16cffe4822e037d
2024-12-04 03:44:47.970 |INFO | proof.go:76: proofData.Commitment[0]: 0x3e19fac8c20c1e0389930ab7ec164edb5fd7b6947476bfd97639c493d96464
2024-12-04 03:44:47.970 |INFO | proof.go:77: proofData.Commitment[1]: 0x2678093ef2b8aa7985716e0126030207008f990cadd946d36b0e8862a6cc310b
2024-12-04 03:44:47.970 |INFO | proof.go:78: proofData.CommitmentPok[0]: 0x0cee59eb8970ceb6fd4972550ae8185c25832c0fac537cb85dbc1b36cc02ede4
2024-12-04 03:44:47.970 |INFO | proof.go:79: proofData.CommitmentPok[1]: 0x25d1c77af41b48aad6b7e16c708e4c055ba0bfd14f0945d40ef94efeb56c4003
2024-12-04 03:44:47.970 |INFO | proof.go:89: witness_0: 0x12a855366912bffec0958fce5e183797f452cb3370e66d59d3568d780bdc4a1d
2024-12-04 03:44:47.970 |INFO | proof.go:89: witness_1: 0x0000000000000000000000000000000000000000000000000000000000000001
2024-12-04 03:44:47.970 |INFO | proof.go:89: witness_2: 0x0000000000000000000000000000000000000000000000000000000000000001
2024-12-04 03:44:47.970 |INFO | proof.go:89: witness_3: 0x000000000000000000000000000000003385ad72da765c0d628053b3796dee67
2024-12-04 03:44:47.970 |INFO | proof.go:89: witness_4: 0x000000000000000000000000000000008d9cbe83e27b1ccbf86fac0bde7ec26d
2024-12-04 03:44:47.970 |INFO | proof.go:89: witness_5: 0x043a906e6334e53ef2c87b54f8d32d92497a3628cb94c182d859c0aa6e440143
2024-12-04 03:44:47.970 |INFO | proof.go:89: witness_6: 0x2caf5a85b6b3d9ae2d02c69bc8d918e3dcf542bd4a547b83294c77f344a890a7
 */

describe('Plonky2ProofVerifierForxLayer test', async () => {
  function loadFixture<T>(fixture: Fixture<T>): Promise<T> {
    const provider = waffle.provider;
    return waffle.createFixtureLoader(provider.getWallets(), provider)(fixture);
  }

  async function fixture([admin]: Wallet[]) {
    const contract = await deployContract(admin);
    return { admin, contract };
  }

  let contract: Plonky2ProofVerifierForxLayer;
  let admin: Wallet;
  beforeEach(async () => {
    const res = await loadFixture(fixture);
    contract = res.contract;
    admin = res.admin;
  });

  it('should pass on true proof', async () => {
    const result = await contract.verifyProof(
      [
        BigNumber.from('0x2c4537e488e3cb6118660dbe93d1a3da817cba90476c0da97b571c82a421814a'), // a0
        BigNumber.from('0x0c3367b3277afd31ef0aac035f55af292549526a5a74089a1e4170a64c4a628a'), // a1

        BigNumber.from('0x019cd8524a22bc985ec5cc43b5e5563754e7d2e8a2ad375cca46209892213d55'), // b00
        BigNumber.from('0x27dcd502799a20ef8b02c542d2e276b45465430b4046107bb8d44a1702e2a3d6'), // b01
        BigNumber.from('0x1b7a471a601b8276ab3c9107668af24589b682b0406a408bb438c73df08254d5'), // b10
        BigNumber.from('0x08c36b4c3ed96c508cf925c4adcfd2aa4c39c1c3d98394cb8b15dd842309f817'), // b11

        BigNumber.from('0x0d96799375aafaf3c68551646bb8267df5492ce93c42969214cf6c008324fb4b'), // c0
        BigNumber.from('0x026ef48e77bbad429d472a8471d862bf7fe1a44904ec9878b16cffe4822e037d') // c1
      ],
      [
        BigNumber.from('0x3e19fac8c20c1e0389930ab7ec164edb5fd7b6947476bfd97639c493d96464'), // Commitment 0
        BigNumber.from('0x2678093ef2b8aa7985716e0126030207008f990cadd946d36b0e8862a6cc310b') // Commitment 1
      ],
      [
        BigNumber.from('0x0cee59eb8970ceb6fd4972550ae8185c25832c0fac537cb85dbc1b36cc02ede4'), // Commitment POK0
        BigNumber.from('0x25d1c77af41b48aad6b7e16c708e4c055ba0bfd14f0945d40ef94efeb56c4003') // Commitment POK1
      ],
      [
        BigNumber.from('0x12a855366912bffec0958fce5e183797f452cb3370e66d59d3568d780bdc4a1d'), // Input commitments root
        BigNumber.from('0x0000000000000000000000000000000000000000000000000000000000000001'), // SMT Root 0
        BigNumber.from('0x0000000000000000000000000000000000000000000000000000000000000001'), // SMT Root 1
        BigNumber.from('0x000000000000000000000000000000003385ad72da765c0d628053b3796dee67'), // App Circuit Output Commitment 0
        BigNumber.from('0x000000000000000000000000000000008d9cbe83e27b1ccbf86fac0bde7ec26d'), // App Circuit Output Commitment 1
        BigNumber.from('0x043a906e6334e53ef2c87b54f8d32d92497a3628cb94c182d859c0aa6e440143'), // CircuitDigest
        BigNumber.from('0x2caf5a85b6b3d9ae2d02c69bc8d918e3dcf542bd4a547b83294c77f344a890a7') // DummyCommitment
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
      BigNumber.from('0x2c4537e488e3cb6118660dbe93d1a3da817cba90476c0da97b571c82a421814a'), // a0
      BigNumber.from('0x0c3367b3277afd31ef0aac035f55af292549526a5a74089a1e4170a64c4a628a'), // a1

      BigNumber.from('0x019cd8524a22bc985ec5cc43b5e5563754e7d2e8a2ad375cca46209892213d55'), // b00
      BigNumber.from('0x27dcd502799a20ef8b02c542d2e276b45465430b4046107bb8d44a1702e2a3d6'), // b01
      BigNumber.from('0x1b7a471a601b8276ab3c9107668af24589b682b0406a408bb438c73df08254d5'), // b10
      BigNumber.from('0x08c36b4c3ed96c508cf925c4adcfd2aa4c39c1c3d98394cb8b15dd842309f817'), // b11

      BigNumber.from('0x0d96799375aafaf3c68551646bb8267df5492ce93c42969214cf6c008324fb4b'), // c0
      BigNumber.from('0x026ef48e77bbad429d472a8471d862bf7fe1a44904ec9878b16cffe4822e037d'), // c1

      BigNumber.from('0x3e19fac8c20c1e0389930ab7ec164edb5fd7b6947476bfd97639c493d96464'), // Commitment 0
      BigNumber.from('0x2678093ef2b8aa7985716e0126030207008f990cadd946d36b0e8862a6cc310b'), // Commitment 1

      BigNumber.from('0x0cee59eb8970ceb6fd4972550ae8185c25832c0fac537cb85dbc1b36cc02ede4'), // Commitment POK0
      BigNumber.from('0x25d1c77af41b48aad6b7e16c708e4c055ba0bfd14f0945d40ef94efeb56c4003'), // Commitment POK1

      BigNumber.from('0x12a855366912bffec0958fce5e183797f452cb3370e66d59d3568d780bdc4a1d'), // Input commitments root
      BigNumber.from('0x0000000000000000000000000000000100000000000000000000000000000001'), // SMT Root
      BigNumber.from('0x3385ad72da765c0d628053b3796dee678d9cbe83e27b1ccbf86fac0bde7ec26d'), // App Circuit Output Commitment
      BigNumber.from('0x043a906e6334e53ef2c87b54f8d32d92497a3628cb94c182d859c0aa6e440143'), // CircuitDigest
      BigNumber.from('0x2caf5a85b6b3d9ae2d02c69bc8d918e3dcf542bd4a547b83294c77f344a890a7') // DummyCommitment
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
