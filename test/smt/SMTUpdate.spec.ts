import { expect } from 'chai';
import { ContractRunner, ContractTransaction, ContractTransactionResponse } from 'ethers';
import { ethers } from 'hardhat';
import { SMTUpdateCircuitProofVerifier, SMTUpdateCircuitProofVerifier__factory } from '../../typechain';
import { splitHash } from '../util';
import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers';

async function deployVerifier(admin: ContractRunner) {
  const factory = await ethers.getContractFactory('SMTUpdateCircuitProofVerifier');
  return factory.connect(admin).deploy();
}

async function deployContracts(admin: ContractRunner) {
  const verifier = await deployVerifier(admin);
  return verifier;
}

const update = {
  oldSmtRoot: '0x0d59f44a853523844fced00b333bba974b5188db76e05433edbd8e78f8ac55cf',
  newSmtRoot: '0x0a9bfc09a53172424166b53114b3f05089032f551ce23ed00c5f7e0047fe0021',
  endBlockHash: '0x4ffa00be572eaa4f4f84f38977b3daad3c40a227c8b7cf5107c62ef85eda01d8',
  endBlockNum: 18020863,
  nextChunkMerkleRoot: '0x0000000000000000000000000000000000000000000000000000000000000000',
  commitPub: '0x119faea41b823e7970eeeba1876786922957564800c06a250005afb866aa8f24',
  proof: {
    a: [
      '0x24378675531c7009ace4a994090201c8232baa99ebec1ecc648340e31e51d021',
      '0x0daa3abdcc1f5ef66222f80821eb103a720728cf0f5ee7ec79339a2c59809cad'
    ],
    b: [
      [
        '0x0db5bcad1c7beedac8762c4236836f2ce6f0ba6f9c8713f20f78fe5728fc3415',
        '0x2751918f64ccd75eae1d3e95ef5b373439993806483a90b7670727e87788ec90'
      ],
      [
        '0x0573e4cf558a2805cfe01585980a783390dcb4358f78d7286ba03497323f8890',
        '0x132ca0517d3af5853e2e7efb807f93cafaf8c25eb0e7a30df80e7799c6ed45e3'
      ]
    ],
    c: [
      '0x18bac2d4c173ad2cf47d1b634466e3245f417e6fc1e1cd71dcaf937e5517da0d',
      '0x27ed45d9d41397387185b23e38f678b9acb5af6ee0f0222e31f36125f13a6225'
    ],
    commitment: [
      '0x05f1ddc37f4a2fa282879875bdaf734cd3d7b2d303c81ec0fec11dad905929ba',
      '0x0a0114af2f9e6f71d88ff47f4199b62d2af47259d136196b14703e315491c817'
    ]
  }
};

describe('SMT Update Circuit Contract Verifier', async () => {
  async function fixture() {
    const [admin] = await ethers.getSigners();
    const contract = await deployContracts(admin);
    return { admin, contract };
  }

  let contract: SMTUpdateCircuitProofVerifier;
  beforeEach(async () => {
    const res = await loadFixture(fixture);
    contract = res.contract;
  });

  const publicInputs = [
    ...splitHash(update.oldSmtRoot),
    ...splitHash(update.newSmtRoot),
    ...splitHash(update.endBlockHash),
    BigInt(update.endBlockNum),
    ...splitHash(update.nextChunkMerkleRoot),
    BigInt(update.commitPub)
  ];

  it('update passes on true proofs', async () => {
    const result = await contract.verifyProof(
      [BigInt(update.proof.a[0]), BigInt(update.proof.a[1])],
      [
        [BigInt(update.proof.b[0][0]), BigInt(update.proof.b[0][1])],
        [BigInt(update.proof.b[1][0]), BigInt(update.proof.b[1][1])]
      ],
      [BigInt(update.proof.c[0]), BigInt(update.proof.c[1])],
      [BigInt(update.proof.commitment[0]), BigInt(update.proof.commitment[1])],
      publicInputs
    );
    expect(result).true;
  });

  it('update verify failed on false proofs', async () => {
    const tx = contract.verifyProof(
      [BigInt(0), BigInt(update.proof.a[1])],
      [
        [BigInt(update.proof.b[0][0]), BigInt(update.proof.b[0][1])],
        [BigInt(update.proof.b[1][0]), BigInt(update.proof.b[1][1])]
      ],
      [BigInt(update.proof.c[0]), BigInt(update.proof.c[1])],
      [BigInt(update.proof.commitment[0]), BigInt(update.proof.commitment[1])],
      publicInputs
    );
    await expect(tx).reverted;
  });
});
