import { Fixture } from 'ethereum-waffle';
import { BigNumber, Wallet } from 'ethers';
import { ethers, waffle } from 'hardhat';
import { Plonky2AppAggVerifier } from '../../typechain';
import assert from 'assert';

async function deployContract(admin: Wallet) {
  const _factory = await ethers.getContractFactory('Plonky2AppAggVerifier');
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

  let contract: Plonky2AppAggVerifier;
  let admin: Wallet;
  beforeEach(async () => {
    const res = await loadFixture(fixture);
    contract = res.contract;
    admin = res.admin;
  });

  /*
  2024-10-20 13:40:14.375 |INFO | proof.go:148: proofData.A[0]: 0x2c76e5ef5796025eba5f3256c24d7b822aa50701eb174abb7598000750750ae4
2024-10-20 13:40:14.375 |INFO | proof.go:149: proofData.A[1]: 0x0b08cd8e17ce850c5d8eb7c9f4c3607cef90e839f09e59dce423a9847373f878
2024-10-20 13:40:14.375 |INFO | proof.go:151: proofData.B[0][0]: 0x059a4ec05507a983690539904be6a4249b65e04a9e36b7d92275394546ad9558
2024-10-20 13:40:14.375 |INFO | proof.go:152: proofData.B[0][1]: 0x019e079b9541227962fc71e5d9389a00dfbd69d6056e0c418a2d4bb4a194989c
2024-10-20 13:40:14.375 |INFO | proof.go:153: proofData.B[1][0]: 0x15f136ad742f25002b00b7a92d84d4f476d0dd7f7bdefed09fe5a4dea66ac38e
2024-10-20 13:40:14.375 |INFO | proof.go:154: proofData.B[1][1]: 0x15c3314e63664273a530a43d4523ea9c30001f35de22859297f1efcee84f2aa9
2024-10-20 13:40:14.375 |INFO | proof.go:156: proofData.C[0]: 0x24c81b3c244cede4020f56af30871bf82fc400ef35ed8f54ae0b874a090cc8ee
2024-10-20 13:40:14.375 |INFO | proof.go:157: proofData.C[1]: 0x22dd78f58630b1f22b32f6514152abbf40e82ab7226ebede3bc9d228faf8b560
2024-10-20 13:40:14.375 |INFO | proof.go:159: proofData.Commitment[0]: 0x2d9abaebec9d08e7ac2406579f9a7a00aa2ef42a6d72d8500e6feecb449bb199
2024-10-20 13:40:14.375 |INFO | proof.go:160: proofData.Commitment[1]: 0x0bea5b71dd38b3d0287f108c279245cbe58757683a1ef2fd287056c1ee3119ce
2024-10-20 13:40:14.375 |INFO | proof.go:161: proofData.CommitmentPok[0]: 0x08249e2a4217631643225db81c54b04070646eaf98ac734bf92fa457a737a84c
2024-10-20 13:40:14.375 |INFO | proof.go:162: proofData.CommitmentPok[1]: 0x0fed359502fb094b488d269665e2edb843d8509666d9f0bf36c9f1d18cb7c23b
2024-10-20 13:40:14.375 |INFO | proof.go:172: witness_0: 0x0000000000000000000000000000000010014c98bf2dd09230f13624e8bcdd6c
2024-10-20 13:40:14.375 |INFO | proof.go:172: witness_1: 0x00000000000000000000000000000000a85a1c5cfe0f7c183eb0c471b6cb9800
2024-10-20 13:40:14.375 |INFO | proof.go:172: witness_2: 0x00000000000000000000000000000000f3d14b07f1f051a687b6ac4134229cbf
2024-10-20 13:40:14.375 |INFO | proof.go:172: witness_3: 0x0000000000000000000000000000000034b50221e4758c88bb0f587b2c404699
   */

  it('should pass on true proof', async () => {
    const result = await contract.verifyProof(
      [
        BigNumber.from('0x2c76e5ef5796025eba5f3256c24d7b822aa50701eb174abb7598000750750ae4'), // a0
        BigNumber.from('0x0b08cd8e17ce850c5d8eb7c9f4c3607cef90e839f09e59dce423a9847373f878'), // a1

        BigNumber.from('0x059a4ec05507a983690539904be6a4249b65e04a9e36b7d92275394546ad9558'), // b00
        BigNumber.from('0x019e079b9541227962fc71e5d9389a00dfbd69d6056e0c418a2d4bb4a194989c'), // b01
        BigNumber.from('0x15f136ad742f25002b00b7a92d84d4f476d0dd7f7bdefed09fe5a4dea66ac38e'), // b10
        BigNumber.from('0x15c3314e63664273a530a43d4523ea9c30001f35de22859297f1efcee84f2aa9'), // b11

        BigNumber.from('0x24c81b3c244cede4020f56af30871bf82fc400ef35ed8f54ae0b874a090cc8ee'), // c0
        BigNumber.from('0x22dd78f58630b1f22b32f6514152abbf40e82ab7226ebede3bc9d228faf8b560') // c1
      ],
      [
        BigNumber.from('0x2d9abaebec9d08e7ac2406579f9a7a00aa2ef42a6d72d8500e6feecb449bb199'), // Commitment 0
        BigNumber.from('0x0bea5b71dd38b3d0287f108c279245cbe58757683a1ef2fd287056c1ee3119ce') // Commitment 1
      ],
      [
        BigNumber.from('0x08249e2a4217631643225db81c54b04070646eaf98ac734bf92fa457a737a84c'), // Commitment POK0
        BigNumber.from('0x0fed359502fb094b488d269665e2edb843d8509666d9f0bf36c9f1d18cb7c23b') // Commitment POK1
      ],
      [
        BigNumber.from('0x0000000000000000000000000000000010014c98bf2dd09230f13624e8bcdd6c'),
        BigNumber.from('0x00000000000000000000000000000000a85a1c5cfe0f7c183eb0c471b6cb9800'),
        BigNumber.from('0x00000000000000000000000000000000f3d14b07f1f051a687b6ac4134229cbf'),
        BigNumber.from('0x0000000000000000000000000000000034b50221e4758c88bb0f587b2c404699')
      ]
    );
    assert.equal(result, true);
  });
});
