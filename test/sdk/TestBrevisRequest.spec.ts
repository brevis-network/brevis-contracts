import { expect } from 'chai';
import { Fixture } from 'ethereum-waffle';
import { BytesLike, BigNumberish, Wallet } from 'ethers';
import { ethers, waffle } from 'hardhat';
import {
  MockSMT__factory,
  BatchZkProofVerifier__factory,
  TestBrevisApp__factory,
  TestBrevisApp,
  BrevisProof__factory,
  BrevisProof,
  AggregationVerifier__factory,
  BrevisRequest,
  BrevisRequest__factory,
} from '../../typechain';

import { Brevis } from '../../typechain/TestBrevisApp';

import { keccak256 } from 'ethers/lib/utils';

async function deployTestBrevisAppContract(admin: Wallet) {
  const smtFactory = await ethers.getContractFactory('MockSMT');
  const smt = await smtFactory.connect(admin).deploy();
  const brevisProofFactory = await ethers.getContractFactory('BrevisProof');
  const brevisProof = await brevisProofFactory.connect(admin).deploy(smt.address);
  const verifierF = await ethers.getContractFactory('BatchZkProofVerifier');
  const verifier = await verifierF.connect(admin).deploy();
  const verifierFApp = await ethers.getContractFactory('AggregationVerifier');
  const verifierApp = await verifierFApp.connect(admin).deploy();
  const chainVerifier = {
    contractAppZkVerifier: verifier.address,
    circuitAppZkVerifier: verifierApp.address
  };
  await brevisProof.updateVerifierAddress([1], [chainVerifier]);
  const factory = await ethers.getContractFactory('BrevisRequest');
  const app = await factory.connect(admin).deploy(smt.address, brevisProof.address);
  return { brevisProof, app };
}

describe('Brevis Reqeust Test', async () => {
  function loadFixture<T>(fixture: Fixture<T>): Promise<T> {
    const provider = waffle.provider;
    return waffle.createFixtureLoader(provider.getWallets(), provider)(fixture);
  }

  async function fixture([admin]: Wallet[]) {
    const { brevisProof, app } = await deployTestBrevisAppContract(admin);
    return { admin, brevisProof, app };
  }

  let app: BrevisRequest;
  let brevisProof: BrevisProof;
  let admin: Wallet;
  before(async () => {
    const res = await loadFixture(fixture);
    app = res.app;
    brevisProof = res.brevisProof;
    admin = res.admin;
  });

  it('should pass on biz test', async () => {
    await app.sendRequest(
      "0x2c9fb082f0e3873b1f63da45a2963e9bb2746baae3dab553e70c3e856e2cbac2", 
      "0x58b529F9084D7eAA598EB3477Fe36064C5B7bbC1",
      "0x0000000000000000000000000000000000000000"
    );
    await brevisProof.setBatchTierVkHashes(
      ['0x22935546FB8B3A680600FCABFE361E5C0A16A99390415ABE8AAB491838DFFCA7'],
      [512]
    );
    await expect(
      app.fulfillRequest(
        "0x2c9fb082f0e3873b1f63da45a2963e9bb2746baae3dab553e70c3e856e2cbac2",
        1,
        "0x1f95570c92409a8419043262fd46a2c820ac263ed273452105ca92972e01d58a1f3c7d6c9b595e3825b9cca38d19141ffa85c570b65dba60b3f5e5ae7ebd5aec09b842ef2fa441494d5daf259e5a6c4f5e271692555e11bf65a6f9294f1879eb2e72c9962c5d804ac8797f30ae5481f9643cdb7518ba16a732c2150d622307071c0de49a565e5d7253ef82837eb0bcb9e6017852ab926a785ed62dacede5fedc2301479cb600ef44ef07dd8f53a1c3e650c49ecc56437f8369812c06b28b42ad16dd473643ab2da5a6fb0f40b7f0d4afcff97b87dbe169fbd5e81d61621b5b0620afac7859d76859d18d4de22dac2917ec644f5cbd5b5f0564f9525c83c0771513060312e9c6aea8b59a3e6077dc74ba44f9b41ef7dd49ea6296d7c2170d9de204b9b67636b872a82ff9987d016b4b7447625bcad160eb13a298f8b9f15647730132bb32abd473a644efa71b637c406e0eebde1e9a2021faba4ab360c3743e1f2c9fb082f0e3873b1f63da45a2963e9bb2746baae3dab553e70c3e856e2cbac21166a98a66ffadcc0d211e59f2dadc370c38ef8683b3908515283d98227c9c6422935546fb8b3a680600fcabfe361e5c0a16a99390415abe8aab491838dffca715d4d00fe9fb519e4add27fd4b4a6f0de337bc0aab7114b21e6b57eabdc6dee81b3738642fbaef19b5b7f1d6e516905af845deb4215458037f76a5d435aee13e0e90df65798e3f2463a52f8b308b89406744fb5a74f1926366e5508ee0463f4a",
        true,
        "0x6c2843ba78feb261798be1aac579d1a4ae2c64b400000000012308dc"
      )
    ).to.emit(app, 'RequestFulfilled');;
  });
});
