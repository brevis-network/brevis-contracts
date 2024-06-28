import { expect } from 'chai';
import { BytesLike, BigNumberish, ContractRunner } from 'ethers';
import { ethers } from 'hardhat';
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

import { Brevis } from '../../typechain/contracts/sdk/apps/framework/BrevisApp';
import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers';

async function deployTestBrevisAppContract(admin: ContractRunner) {
  const smtFactory = await ethers.getContractFactory('MockSMT');
  const smt = await smtFactory.connect(admin).deploy();
  const smtAddress = await smt.getAddress()
  const brevisProofFactory = await ethers.getContractFactory('BrevisProof');
  const brevisProof = await brevisProofFactory.connect(admin).deploy(smtAddress);
  const brevisProofAddress = await brevisProof.getAddress()
  const verifierF = await ethers.getContractFactory('BatchZkProofVerifier');
  const verifier = await verifierF.connect(admin).deploy();
  const verifierAddress = await verifier.getAddress()
  const verifierFApp = await ethers.getContractFactory('AggregationVerifier');
  const verifierApp = await verifierFApp.connect(admin).deploy();
  const verifierAppAddress = await verifierApp.getAddress()
  const chainVerifier = {
    contractAppZkVerifier: verifierAddress,
    circuitAppZkVerifier: verifierAppAddress
  };
  await brevisProof.updateVerifierAddress([42161], [chainVerifier]);
  const factory = await ethers.getContractFactory('BrevisRequest');
  const app = await factory.connect(admin).deploy(smtAddress, brevisProofAddress);

  return { brevisProof, app };
}


describe('Brevis Reqeust Test', async () => {
  async function fixture() {
    const [admin] = await ethers.getSigners()
    const { brevisProof, app } = await deployTestBrevisAppContract(admin);
    return { admin, brevisProof, app };
  }

  let app: BrevisRequest;
  let brevisProof: BrevisProof;
  before(async () => {
    const res = await loadFixture(fixture);
    app = res.app;
    brevisProof = res.brevisProof;
  });

  it('should pass on biz test', async () => {
    await app.sendRequest(
      "0x206d1ac729747e772cee06d1ac1e99b25eb5060c479e7b0210b10468b5cfaf0a", 
      "0x58b529F9084D7eAA598EB3477Fe36064C5B7bbC1",
      "0x0000000000000000000000000000000000000000"
    );
    await brevisProof.setBatchTierVkHashes(
      ['0x0E73E4069300C6E5A2B15872F8BC3955E6827D3EF2F475BB7CE7B6F967115D89'],
      [512]
    );
    //await brevisProof.submitProof(42161, "0x1731c0fffb965dc056a94bc2a4e6f5083172294b4f98797adbe298caf32181cb2808d862661ef95f6285d22b7b8b4618a313bad3f23134b982bd05d4f111ca230c8828c7431dc9d62b7748f437ebf5d9757854c4d4d3a37ed0ce25d6bf9612d104ccce7100c3f9f96564dd75e0d15e82199a1079658d55b48ae49d603e3f208c18ca475a1e9da41b7115adae00aadf02a11481943527792ef8d78c4646f644760a9e47ec93af3b6c444b8d7f7a74841caf6a52cb5b97e814652adc42096e7e5f2f4f3212cc24959e758d754724fb1f294ba4f1562d98000c3a7bf7271221289c2de56a9f2f48d3a2cb88aba766350c6291a1cea10e85deb25c81291b7c10faab11638dfeb4391ddc9a70b1f4d4149881d78695b3bd5734f1690eed0165adb892260a37c7b9abb6b445a1082b1623078b3b2a95e1c2bec8bcb14937b9574ca1330fb2c3aa55ad38110c03ed37252d173637dc678e89ec8eb6030f9e52a7a53b16206d1ac729747e772cee06d1ac1e99b25eb5060c479e7b0210b10468b5cfaf0a00a7d78ad7ece381c5499a27e03b40fe65ebc96793efb316f6f148257631aa490e73e4069300c6e5a2b15872f8bc3955e6827d3ef2f475bb7ce7b6f967115d89c647635a4aaa5bcd7e024cc3f162cf67bd51eab83bd866b5abc0feb16a5daf2d1414c50ed76ad28bd21b33ddb867b55151b315ef570c0cc5146e85e9332905f4", true);
    await expect(
      app.fulfillRequest(
        "0x206d1ac729747e772cee06d1ac1e99b25eb5060c479e7b0210b10468b5cfaf0a",
        42161,
        "0x1731c0fffb965dc056a94bc2a4e6f5083172294b4f98797adbe298caf32181cb2808d862661ef95f6285d22b7b8b4618a313bad3f23134b982bd05d4f111ca230c8828c7431dc9d62b7748f437ebf5d9757854c4d4d3a37ed0ce25d6bf9612d104ccce7100c3f9f96564dd75e0d15e82199a1079658d55b48ae49d603e3f208c18ca475a1e9da41b7115adae00aadf02a11481943527792ef8d78c4646f644760a9e47ec93af3b6c444b8d7f7a74841caf6a52cb5b97e814652adc42096e7e5f2f4f3212cc24959e758d754724fb1f294ba4f1562d98000c3a7bf7271221289c2de56a9f2f48d3a2cb88aba766350c6291a1cea10e85deb25c81291b7c10faab11638dfeb4391ddc9a70b1f4d4149881d78695b3bd5734f1690eed0165adb892260a37c7b9abb6b445a1082b1623078b3b2a95e1c2bec8bcb14937b9574ca1330fb2c3aa55ad38110c03ed37252d173637dc678e89ec8eb6030f9e52a7a53b16206d1ac729747e772cee06d1ac1e99b25eb5060c479e7b0210b10468b5cfaf0a00a7d78ad7ece381c5499a27e03b40fe65ebc96793efb316f6f148257631aa490e73e4069300c6e5a2b15872f8bc3955e6827d3ef2f475bb7ce7b6f967115d89c647635a4aaa5bcd7e024cc3f162cf67bd51eab83bd866b5abc0feb16a5daf2d1414c50ed76ad28bd21b33ddb867b55151b315ef570c0cc5146e85e9332905f4",
        true,
        "0xad8054d4922c24523f02b2cd04798cff88fbfa33000000000000000000000000000000010000000000002ed100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000173b5e353f7ce00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003"
      )
    ).to.emit(app, 'RequestFulfilled');;
  });
});
