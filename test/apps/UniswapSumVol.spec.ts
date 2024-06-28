import { expect } from 'chai';
import { Wallet } from 'ethers';
import { ethers, } from 'hardhat';

import { loadFixture, time } from '@nomicfoundation/hardhat-toolbox/network-helpers';

import {
  MockSMT__factory,
  UniswapSumVolume,
  UniswapSumVolume__factory,
  UniswapSumVolumeVerifier__factory
} from '../../typechain';

async function deployUniswapSumContract() {
  const [admin] = await ethers.getSigners();
  const smtFactory = await ethers.getContractFactory('MockSMT');
  const smt = await smtFactory.connect(admin).deploy();
  const smtAddress = await smt.getAddress()
  const factory = await ethers.getContractFactory('UniswapSumVolume');
  const contract = await factory.connect(admin).deploy(smtAddress);
  const verifierF = await ethers.getContractFactory('UniswapSumVolumeVerifier');
  const verifier = await verifierF.connect(admin).deploy();
  const verifierAddress = await verifier.getAddress()
  await contract.updateVerifierAddress(1, verifierAddress);

  return contract;
}

describe('Uniswap Sum Volume App Test', async () => {
  async function fixture() {
    const contract = await deployUniswapSumContract();
    return { admin, contract };
  }

  let contract: UniswapSumVolume;
  let admin: Wallet;
  before(async () => {
    const res = await loadFixture(fixture);
    contract = res.contract;
    admin = res.admin;
  });

  it('should pass on submit proof', async () => {
    const proof =
      '0x138609a2cb8a8372320273aafd2e671cbc6528d7e947e4bd3ccc56245cb0ac6d13f3374447477a80e1fb1a10a958e42140e624952415c932ad5cf9964a589912173fcc3d8b119d5de6af781d039a8e912102f17c670d867a9a6c9d119b0288312672733e9826848198abd3d6bd13fa98629fa19b1694ee207923dd626ed3076c2e679f577b5b39f2a7300982e9559cae870440c226576899c1269e046d317629144e88462fad391d3682fc9974327738d9dade8421f4ea5cf1d9fdee7a0a938221802eab2992d1e92e4412e6a4945def8dda907299d4381b2380fc0d24ddcc0c13a0493f586f8ab8b20d0dd516830031123dd1908a0c86cc8466590c83d8b2ac2dbe2c93a0bf79c3415cbc2bf0ee501c5847657b68a5414a54a34f2334b50141240a0875e1bdfaba0744be85ac8f498d1618d10c0646cf59efa0d801ce8266b61578125e53474179266a6a4f5c92aa69d8448aab76b715201a3e801ddd52ae89000000000000000000000000c925403763b9ebd6700ac23c90510f0ff174dfc3000000000000000000000000000000000000000000000000000000028fa6ae00000000000000000000000000000000000c8132219bedc56454178d3a03840e7700000000000000000000000000000000667d05b2324f6f8a4a3146500a0584f0000000000000000000000000000000000000000000000000000000000000000218dbb7af7ced04f8e3bb519ed400b662290d409b49d5a2cc48ff5a088396b6e6';
    await contract.setBatchTierVkHashes('0x18dbb7af7ced04f8e3bb519ed400b662290d409b49d5a2cc48ff5a088396b6e6');
    await expect(contract.submitUniswapSumVolumeProof(1, proof)).to.emit(contract, 'SumVolume');
  });
});
