import { Fixture } from 'ethereum-waffle';
import { BigNumber, Wallet } from 'ethers';
import { ethers, waffle } from 'hardhat';
import { BrevisPlonky2SmtVerifier } from '../../typechain';
import assert from 'assert';
import { hexToBytes } from '../util';
import {expect} from "chai";

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
          BigNumber.from('0x1a6de85e1829528a0963b409da663343c00c138f7582f990d52ff19b5ff5f29a'), // a0
          BigNumber.from('0x218988b65ffad99a70aab1c35e7be3d66565370b07e37723db4734ed3aac5c6d'), // a1

          BigNumber.from('0x19e44d5e50d7ae28d1496263fb6b5ec2202cc6816d9edde31df6c78c9622d152'), // b00
          BigNumber.from('0x11681f4302e49dfa1ffd8875b3edc9b00fcf09bfba9b6b3fe946bcde308b753d'), // b01
          BigNumber.from('0x1b9a79b676a5cbe3966298a987bdd4489816ae2595759afbe65eee7a593a26c8'), // b10
          BigNumber.from('0x269aacd368910231b4e38034a91d74dc28a7443a8a7c25df97dd478f99e4ced2'), // b11

          BigNumber.from('0x08eb00a0bff2d2f28313e41f0ef2441da2218bb271eb9a3d267f9825421574b8'), // c0
          BigNumber.from('0x2aa727997830b038b0c5a24f3720cafa900df677c1100642cd764a2bbf89f994') // c1
        ],
        [
          BigNumber.from('0x1e98f5bf21991b908775d2e7ac6558410fe5f67a14464b2a8699e51cdb0adb25'), // Commitment 0
          BigNumber.from('0x0b0f63cbccf76a457c986722457e5209fd6c5676bfcd15359117a327637e39b9') // Commitment 1
        ],
        [
          BigNumber.from('0x27ad93d284d74b3d09c341f205136699a95c22b7e80a0869af593106c4fb0aa0'), // Commitment POK0
          BigNumber.from('0x10b18b34c2a68dbacd9d008d933350c68431d32bbec17bf741e1b16fa2dc63c7') // Commitment POK1
        ],
        [
          BigNumber.from('0x000000000000000000000000000000004fcf04370e7fff28a61aeb865faa4c97'), // EndBlkHash 0
          BigNumber.from('0x00000000000000000000000000000000201c06eddd475833ba42138333ffc3dd'), // EndBlkHash 1
          BigNumber.from('0x00000000000000000000000000000000a7c945bf248170650ea0370c8bd4dea8'), // NewSmtRoot 0
          BigNumber.from('0x00000000000000000000000000000000301ede52931e4d12a2f0e7a0b4698b6f'), // NewSmtRoot 1
          BigNumber.from('0x00000000000000000000000000000000b244c559122805487f8377c4ad9bc212'), // OldSmtRoot 0
          BigNumber.from('0x000000000000000000000000000000000c34b373ad27a41840940db61ec3c5b7'), // OldSmtRoot 1
          BigNumber.from('0x0000000000000000000000000000000000000000000000000000000000000000'), // NextChunkMerkleRoot 0
          BigNumber.from('0x0000000000000000000000000000000000000000000000000000000000000000'), // NextChunkMerkleRoot 1
          BigNumber.from('0x25d8769b8e3aa1755bc55f71443ae0f9f0b5b3263d51de80ec1c4cddd16dff20')  // CircuitDigest
        ]
    );
    assert.equal(result, true);
  });

    it('should revert on false proof', async () => {
        const result = contract.verifyProof(
            [
                BigNumber.from('0x1a6de85e1829528a0963b409da663343c00c138f7582f990d52ff19b5ff5f29a'), // a0
                BigNumber.from('0x218988b65ffad99a70aab1c35e7be3d66565370b07e37723db4734ed3aac5c6d'), // a1

                BigNumber.from('0x19e44d5e50d7ae28d1496263fb6b5ec2202cc6816d9edde31df6c78c9622d152'), // b00
                BigNumber.from('0x11681f4302e49dfa1ffd8875b3edc9b00fcf09bfba9b6b3fe946bcde308b753d'), // b01
                BigNumber.from('0x1b9a79b676a5cbe3966298a987bdd4489816ae2595759afbe65eee7a593a26c8'), // b10
                BigNumber.from('0x269aacd368910231b4e38034a91d74dc28a7443a8a7c25df97dd478f99e4ced2'), // b11

                BigNumber.from('0x08eb00a0bff2d2f28313e41f0ef2441da2218bb271eb9a3d267f9825421574b8'), // c0
                BigNumber.from('0x2aa727997830b038b0c5a24f3720cafa900df677c1100642cd764a2bbf89f994') // c1
            ],
            [
                BigNumber.from('0x1e98f5bf21991b908775d2e7ac6558410fe5f67a14464b2a8699e51cdb0adb25'), // Commitment 0
                BigNumber.from('0x0b0f63cbccf76a457c986722457e5209fd6c5676bfcd15359117a327637e39b9') // Commitment 1
            ],
            [
                BigNumber.from('0x27ad93d284d74b3d09c341f205136699a95c22b7e80a0869af593106c4fb0aa0'), // Commitment POK0
                BigNumber.from('0x10b18b34c2a68dbacd9d008d933350c68431d32bbec17bf741e1b16fa2dc63c7') // Commitment POK1
            ],
            [
                BigNumber.from('0x000000000000000000000000000000004fcf04370e7fff28a61aeb865faa4c97'), // EndBlkHash 0
                BigNumber.from('0x00000000000000000000000000000000201c06eddd475833ba42138333ffc3dd'), // EndBlkHash 1
                BigNumber.from('0x00000000000000000000000000000000a7c945bf248170650ea0370c8bd4dea8'), // NewSmtRoot 0
                BigNumber.from('0x00000000000000000000000000000000301ede52931e4d12a2f0e7a0b4698b6f'), // NewSmtRoot 1
                BigNumber.from('0x00000000000000000000000000000000b244c559122805487f8377c4ad9bc212'), // OldSmtRoot 0
                BigNumber.from('0x000000000000000000000000000000000c34b373ad27a41840940db61ec3c5b7'), // OldSmtRoot 1
                BigNumber.from('0x0000000000000000000000000000000000000000000000000000000000000001'), // NextChunkMerkleRoot 0
                BigNumber.from('0x0000000000000000000000000000000000000000000000000000000000000000'), // NextChunkMerkleRoot 1
                BigNumber.from('0x25d8769b8e3aa1755bc55f71443ae0f9f0b5b3263d51de80ec1c4cddd16dff20')  // CircuitDigest
            ]
        );

        await expect(result).reverted;
    });
});
