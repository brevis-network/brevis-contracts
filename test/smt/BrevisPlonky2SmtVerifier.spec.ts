import { Fixture } from 'ethereum-waffle';
import { BigNumber, Wallet } from 'ethers';
import { ethers, waffle } from 'hardhat';
import { BrevisPlonky2SmtVerifier } from '../../typechain';
import assert from 'assert';
import { hexToBytes } from '../util';

async function deployContract(admin: Wallet) {
  const _factory = await ethers.getContractFactory('BrevisPlonky2SmtVerifier');
  const _contract = await _factory.connect(admin).deploy();
  return _contract;
}

describe('BrevisPlonky2SmtVerifier test', async () => {
  function loadFixture<T>(fixture: Fixture<T>): Promise<T> {
    const provider = waffle.provider;
    return waffle.createFixtureLoader(provider.getWallets(), provider)(fixture);
  }

  async function fixture([admin]: Wallet[]) {
    const contract = await deployContract(admin);
    return { admin, contract };
  }

  let contract: BrevisPlonky2SmtVerifier;
  let admin: Wallet;
  beforeEach(async () => {
    const res = await loadFixture(fixture);
    contract = res.contract;
    admin = res.admin;
  });

  it('should pass on true proof', async () => {
    const result = await contract.verifyProof(
        [
          BigNumber.from('0x07e2778626964a754b9195def728ea19cf0e3d8b8c5753ba43bc0dd139133736'), // a0
          BigNumber.from('0x042438dd97b6592229539cd15443ee5f9d63b31db2f31481c65594711f3067d2'), // a1

          BigNumber.from('0x236453f47b140a9a05cbe55415e48974299241a2c69469cc63da70e62bc320c4'), // b00
          BigNumber.from('0x28f5c16d64ed7784a8a2f3595362ddc8949848b5ca58820d3309e216fc725e57'), // b01
          BigNumber.from('0x284c0ec5df51b61bac50ec39b335c911b9ede6c8a08aff8f415ae566d59c78e6'), // b10
          BigNumber.from('0x168970a94c85b5a679aaf149141485ac1e2cf3100133c73f47cd299ffa8ae992'), // b11

          BigNumber.from('0x22ff39d4822505452f1668fe7d3e75702007f6a81bc64c1615880756144a6a19'), // c0
          BigNumber.from('0x17065d881740ce2efd4d9494d691cdb52b229a0316d92a751a1d73a9a18c41a2') // c1
        ],
        [
          BigNumber.from('0x1019ebea4a5e5e32c3e9bf00424f3043e29173d00798b616b17a0e1e59994e09'), // Commitment 0
          BigNumber.from('0x153f42fa8353f93f61ea7ae6eca63e917d4cbf38dcf1afd8301828ae42f67ee7') // Commitment 1
        ],
        [
          BigNumber.from('0x0ed223b3b6edc2c3610fd8f865e1a9f910597a49d0882da3753f125fad74b764'), // Commitment POK0
          BigNumber.from('0x1b4d6fff41d141d1fde2abfe75dcb60291cfb97bac4215b23e7d0041d6b76df7') // Commitment POK1
        ],
        [
          BigNumber.from('0x00000000000000000000000000000000409f3590dcd221e5786a20423cfc8305'), // EndBlkHash 0
          BigNumber.from('0x00000000000000000000000000000000738a88d3efaa6210428de2024f75aceb'), // EndBlkHash 1
          BigNumber.from('0x0000000000000000000000000000000045187e590779d8dfca61be4ea6c9e888'), // NewSmtRoot 0
          BigNumber.from('0x0000000000000000000000000000000031b07e7bfefbda1cdb66921cb0d63b86'), // NewSmtRoot 1
          BigNumber.from('0x00000000000000000000000000000000ca3022d33d006b913ecfad08eb9d6dde'), // OldSmtRoot 0
          BigNumber.from('0x00000000000000000000000000000000985e93fdd80ca6fadebcd2172d28c3f3'), // OldSmtRoot 1
          BigNumber.from('0x0000000000000000000000000000000082fc73fc57007f802f958f4b18d7fb6f'), // NextChunkMerkleRoot 0
          BigNumber.from('0x000000000000000000000000000000009f01498afa81edea7ee5404bdb09a283'), // NextChunkMerkleRoot 1
          BigNumber.from('0x1ba0ce182fc9cf0eba3aa35b990b052d0d956d641defdddd8b93473a7d9d6d5e')  // CircuitDigest
        ]
    );
    assert.equal(result, true);
  });
});
