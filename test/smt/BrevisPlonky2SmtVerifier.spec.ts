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
          BigNumber.from('0x0e387f2b729455c4097e269c20c83748e9b3a1e149c318c252c921721627ed65'), // a0
          BigNumber.from('0x2d3771c636e8c3a9348b9010f3e3c9b8c7e6ad546cdae3afec03816b170171e8'), // a1

          BigNumber.from('0x22c1863a7ac076b464ad34441c60b2cc3dd28967c446498be279a324d3d03e09'), // b00
          BigNumber.from('0x0aa3291b1879abbaf7bd2e70efbd39bf6f48a51eda404b98fd833eff3b4fb065'), // b01
          BigNumber.from('0x0d491305ccb73ec18045c71f59115b9f7fb37a718bc83489c8346d48f6d563b0'), // b10
          BigNumber.from('0x21e627b00b26bebf197d2c58f0c2968419be11e6cbdf31541c7cae99d88d1ad4'), // b11

          BigNumber.from('0x047304cc88e4a6c3315b98ccd1b59a674e25234a31b6d9295940714d6bedf07c'), // c0
          BigNumber.from('0x0375dfb0eb0eba372fcf3d57bbf00db6c798046e8855d91f32e18c0f952a5edd') // c1
        ],
        [
          BigNumber.from('0x15828066f2bcde3528b822aeef1ae54a99352806d7d8e7c7af80ddd79978eaf5'), // Commitment 0
          BigNumber.from('0x1a480d818841b2d93baada593cec0a1808f0f1526b622ff174808199c04eb94e') // Commitment 1
        ],
        [
          BigNumber.from('0x18dd29acec340d08fbbe737655349828559eb5af8a6fd03a5d2e74e46a9248df'), // Commitment POK0
          BigNumber.from('0x26bb05e4dc8886c8e45ef26bff932bcaf63ed561a95c6b9708251cc2642666a8') // Commitment POK1
        ],
        [
          BigNumber.from('0x00000000000000000000000000000000409f3590dcd221e5786a20423cfc8305'),
          BigNumber.from('0x00000000000000000000000000000000738a88d3efaa6210428de2024f75aceb'),
          BigNumber.from('0x0000000000000000000000000000000045187e590779d8dfca61be4ea6c9e888'),
          BigNumber.from('0x0000000000000000000000000000000031b07e7bfefbda1cdb66921cb0d63b86'),
          BigNumber.from('0x00000000000000000000000000000000ca3022d33d006b913ecfad08eb9d6dde'),
          BigNumber.from('0x00000000000000000000000000000000985e93fdd80ca6fadebcd2172d28c3f3'),
          BigNumber.from('0x0000000000000000000000000000000082fc73fc57007f802f958f4b18d7fb6f'),
          BigNumber.from('0x000000000000000000000000000000009f01498afa81edea7ee5404bdb09a283'),
          BigNumber.from('0x06f5391a3225d0639697ec95a506c8b762239a8d1cfed24e9fb83d773ec43717')
        ]
    );
    assert.equal(result, true);
  });
});
