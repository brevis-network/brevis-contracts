import { Fixture } from 'ethereum-waffle';
import { BigNumber, Wallet } from 'ethers';
import { ethers, waffle } from 'hardhat';
import { Plonky2AggProofVerifier } from '../../typechain';
import assert from 'assert';

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

  /*
2024-10-22 06:33:24.927 |INFO | proof.go:148: proofData.A[0]: 0x085cbf5aff3b193ca425230910fa514fb8799dfa0f4d3c0a7a3ca75ea7c33876
2024-10-22 06:33:24.927 |INFO | proof.go:149: proofData.A[1]: 0x288115ac0ecbbb83b5a93f8680c3dbe3cc0d2d5557b9a5b43510a4c92dfe67a2
2024-10-22 06:33:24.927 |INFO | proof.go:151: proofData.B[0][0]: 0x15744d267afc176d716301ad8ad8c88beb3248dc8f09c1d8380c56d24fe79c08
2024-10-22 06:33:24.927 |INFO | proof.go:152: proofData.B[0][1]: 0x1e1ea03ea68143787edbc16d754835862a00b35aee12942a673a190bab91878f
2024-10-22 06:33:24.927 |INFO | proof.go:153: proofData.B[1][0]: 0x1a65bfe7990e30fe2a3773965ba29487fd2c5fea863560833e01f6c34607e18b
2024-10-22 06:33:24.927 |INFO | proof.go:154: proofData.B[1][1]: 0x0b421a3b82d91c689f0d01c3ab1458770c859f88375b9eaecb763d61da4331c5
2024-10-22 06:33:24.927 |INFO | proof.go:156: proofData.C[0]: 0x2ef4498c44ad3eddb225bd46145bd412f038ab4336f92f9f6899be7ccf7894fd
2024-10-22 06:33:24.927 |INFO | proof.go:157: proofData.C[1]: 0x1f7bff28305f1a70e76b59e70b9bafda0a61ac92c062fc8ac868bce714fd0a08
2024-10-22 06:33:24.927 |INFO | proof.go:159: proofData.Commitment[0]: 0x0614fd03069c0ede511046d71fb1d257bfaf60f44c89bfac856c4eef648fe0d7
2024-10-22 06:33:24.927 |INFO | proof.go:160: proofData.Commitment[1]: 0x0cb356627bca41abd132d8c1207bce6a4c583ae7ffe3e4af0bb2be4b630dd9
2024-10-22 06:33:24.927 |INFO | proof.go:161: proofData.CommitmentPok[0]: 0x14c6cac2599f247fee469d1fe6b814a50f8b5daddc21678dc6db80cd59a0490d
2024-10-22 06:33:24.927 |INFO | proof.go:162: proofData.CommitmentPok[1]: 0x0b10b60132c91d57d0cdd0784d8da35fa718ebb7ef841f5768ae75726c540e7d
2024-10-22 06:33:24.927 |INFO | proof.go:172: witness_0: 0x0000000000000000000000000000000010014c98bf2dd09230f13624e8bcdd6c
2024-10-22 06:33:24.927 |INFO | proof.go:172: witness_1: 0x00000000000000000000000000000000a85a1c5cfe0f7c183eb0c471b6cb9800
2024-10-22 06:33:24.927 |INFO | proof.go:172: witness_2: 0x00000000000000000000000000000000ec2421d3673015d1f77586110882910e
2024-10-22 06:33:24.927 |INFO | proof.go:172: witness_3: 0x000000000000000000000000000000000aeb1f2268c6c016a3954be3ed26798c
   */

  it('should pass on true proof', async () => {
    const result = await contract.verifyProof(
      [
        BigNumber.from('0x085cbf5aff3b193ca425230910fa514fb8799dfa0f4d3c0a7a3ca75ea7c33876'), // a0
        BigNumber.from('0x288115ac0ecbbb83b5a93f8680c3dbe3cc0d2d5557b9a5b43510a4c92dfe67a2'), // a1

        BigNumber.from('0x15744d267afc176d716301ad8ad8c88beb3248dc8f09c1d8380c56d24fe79c08'), // b00
        BigNumber.from('0x1e1ea03ea68143787edbc16d754835862a00b35aee12942a673a190bab91878f'), // b01
        BigNumber.from('0x1a65bfe7990e30fe2a3773965ba29487fd2c5fea863560833e01f6c34607e18b'), // b10
        BigNumber.from('0x0b421a3b82d91c689f0d01c3ab1458770c859f88375b9eaecb763d61da4331c5'), // b11

        BigNumber.from('0x2ef4498c44ad3eddb225bd46145bd412f038ab4336f92f9f6899be7ccf7894fd'), // c0
        BigNumber.from('0x1f7bff28305f1a70e76b59e70b9bafda0a61ac92c062fc8ac868bce714fd0a08') // c1
      ],
      [
        BigNumber.from('0x0614fd03069c0ede511046d71fb1d257bfaf60f44c89bfac856c4eef648fe0d7'), // Commitment 0
        BigNumber.from('0x0cb356627bca41abd132d8c1207bce6a4c583ae7ffe3e4af0bb2be4b630dd9') // Commitment 1
      ],
      [
        BigNumber.from('0x14c6cac2599f247fee469d1fe6b814a50f8b5daddc21678dc6db80cd59a0490d'), // Commitment POK0
        BigNumber.from('0x0b10b60132c91d57d0cdd0784d8da35fa718ebb7ef841f5768ae75726c540e7d') // Commitment POK1
      ],
      [
        BigNumber.from('0x0000000000000000000000000000000010014c98bf2dd09230f13624e8bcdd6c'),
        BigNumber.from('0x00000000000000000000000000000000a85a1c5cfe0f7c183eb0c471b6cb9800'),
        BigNumber.from('0x00000000000000000000000000000000ec2421d3673015d1f77586110882910e'),
        BigNumber.from('0x000000000000000000000000000000000aeb1f2268c6c016a3954be3ed26798c')
      ]
    );
    assert.equal(result, true);
  });
});
