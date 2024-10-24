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

  it('should pass on true proof', async () => {
    const result = await contract.verifyProof(
      [
        BigNumber.from('0x055da2f44ec1250fe4943bae59206ddfb2214cd3b8049be31b9915e1dde08148'), // a0
        BigNumber.from('0x27187295f976fce577654bc0c31892937793bbecdab3a1b3c671c087dab3c14b'), // a1

        BigNumber.from('0x0a89b97abff05eeef06762ed638f702de8d70d4551654670f8212a302b54eb23'), // b00
        BigNumber.from('0x194f802b9aa4f0d0a7de06dd054941a610069d4915843fd6e33895ff85269dae'), // b01
        BigNumber.from('0x150c09738284b51714bd0ed516e366847a2e339f3351ff16786d33c1c8609452'), // b10
        BigNumber.from('0x11949044585f0378a0c9ad3dd5c0bc193262536179f1db5288276e863f2d33f9'), // b11

        BigNumber.from('0x427dfbd96c3dccabe35ff12a9ed9bd94bb6c4903bf4085dac280221f625c03'), // c0
        BigNumber.from('0x2d017a3a6ea21f5a744d50bb27facc1fc03e4542c46d08dc1a2393942b27401f') // c1
      ],
      [
        BigNumber.from('0x23d2f8c229ebf6f57203893798cbdf5e70af15f8630f8fd026540451e0bebb7e'), // Commitment 0
        BigNumber.from('0x0bf51700de4ad5378cfbcc2af927820a886e37a5c35b56ba96015458d8e31d89') // Commitment 1
      ],
      [
        BigNumber.from('0x172e94231d85e0f7db939a3ebbdda5e643eecf7aad86f8bf69d4e2327344c2a4'), // Commitment POK0
        BigNumber.from('0x10fdad86e46127a8f6d8c51c84c1aa05aaffd5b8959eee4484aa052f2debfb7f') // Commitment POK1
      ],
      [
        BigNumber.from('0x0000000000000000000000000000000010014c98bf2dd09230f13624e8bcdd6c'), // Merkle Root 0
        BigNumber.from('0x00000000000000000000000000000000a85a1c5cfe0f7c183eb0c471b6cb9800'), // Merkle Root 1
        BigNumber.from('0x00000000000000000000000000000000ec2421d3673015d1f77586110882910e'), // ProofIds Commitment 0
        BigNumber.from('0x000000000000000000000000000000000aeb1f2268c6c016a3954be3ed26798c'), // ProofIds Commitment 1
        BigNumber.from('0x078ab850e8148fc412016972abf837fddbc8c7f87d049e337fcdfdc1a47caca2') // Agg Vk Hash
      ]
    );
    assert.equal(result, true);
  });
});
