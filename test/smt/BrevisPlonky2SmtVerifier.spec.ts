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
          BigNumber.from('0x2f5dfef59cfa952da626e6839cabd03a9d414e37f3d59ff0c4f148f21e8e06a8'), // a0
          BigNumber.from('0x04e41eb5d00ae7a99db6f75d0d3c52592ef38ea567cd97b02ee0d39422384176'), // a1

          BigNumber.from('0x0b9e8df76ae619d9cc39f7249eb01bfda2c1544e4346312fa8e34dc8b6f55f1f'), // b00
          BigNumber.from('0x1b73d50861f9a5ca5bc34a4b6178fe04d96cf3ebd831da0010abefeda6819c25'), // b01
          BigNumber.from('0x2f3ca3a6fa094ba450e75fc6963bb068b7454434adecd2f9d7c51a3d3dc1cba2'), // b10
          BigNumber.from('0x18ac635bd13127131e8143e4051e1216ed0e491c157bf9ac74a75c6705d9c351'), // b11

          BigNumber.from('0x29bfc79e0ddd57d61b85db1b40e5dafc7a6a793cbd3f5364e72c645ba5ee3ea7'), // c0
          BigNumber.from('0x07b0cd6f744358867b79e545337b671ebf87198d1fa69c3b9b68df2dd7f496a8') // c1
        ],
        [
          BigNumber.from('0x1f906ca05aa8f7ced26d4600394dc41982bad9551efb20085075b5ac0ebeb90c'), // Commitment 0
          BigNumber.from('0x193befaf0f4bc0507d255b7494bdc471d1e3c560a7092ecc26c29c1b819dcb34') // Commitment 1
        ],
        [
          BigNumber.from('0x292e5a07d10dfce215a6723b5c5d125f0d477a033fcb003e5bc3b918f4a78061'), // Commitment POK0
          BigNumber.from('0x1ea8c4c6d22d5d78ecb676b654740148415a89b6cd30499dd6fb97bb9872ce9c') // Commitment POK1
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
