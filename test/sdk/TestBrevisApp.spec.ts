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
  AggregationVerifier__factory
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
  const factory = await ethers.getContractFactory('TestBrevisApp');
  const app = await factory.connect(admin).deploy(brevisProof.address);
  return { brevisProof, app };
}

describe('Brevis App Test', async () => {
  function loadFixture<T>(fixture: Fixture<T>): Promise<T> {
    const provider = waffle.provider;
    return waffle.createFixtureLoader(provider.getWallets(), provider)(fixture);
  }

  async function fixture([admin]: Wallet[]) {
    const { brevisProof, app } = await deployTestBrevisAppContract(admin);
    return { admin, brevisProof, app };
  }

  let app: TestBrevisApp;
  let brevisProof: BrevisProof;
  let admin: Wallet;
  before(async () => {
    const res = await loadFixture(fixture);
    app = res.app;
    brevisProof = res.brevisProof;
    admin = res.admin;
  });

  it('should pass on biz test', async () => {
    const proof =
      '0x0ec428d3a6be5f691d01860c80d62e528939be9ab8242fdd5db5e6a4d7f2b5840669ac51c672a24bf23fc863dc5aaf7aa63a98f6515f794c9b94d27e082a3ddb2a48541a243b7fef27e2a0777bde1524c0340c647fd741af01a75848eed87153076ca76f689a0a8cc26f08c39b97531fa5dc47a64b7bb59d07099a53ba08acfc2253d84b3adf4aae2b12f38fb87abaab29f3d4518dc9c4d640979edb4895ae44183390fe893725ad64d43799116a439ce85a50283a0b6ff74c755f5c82860fd12a0f344e722fd5e74ac186a590182c1904465f37b6b4305a6784a6d2c5aa36a40d575bb76f25671374e3fe1d1b2d2f8c6d79a49700be7ccf9df0d6bf2af455e10a02bf0f7a88c78b9106e047f636ddb924023a0b9c631d580856d3fe1272a4b623d7e52d7ad1b7538f1925f1ae3cf5b3f88437295a7bde9215ea7b22f35f675e257d69a3df16c62393c8736cee2d88842bea7fd4389245d9bce14db16283196497cc677c7cb46dca48c1b2e181801720910c292e9009bcec63a77a0ed579e844000000000000000000000000000000000000000000000000000000000000000328480547bb8aaa203cf9cadcca995cc6abda65c55f74c706134e6931b72cea8c';
    const isFieldsInTopic: [BigNumberish, BigNumberish, BigNumberish, BigNumberish, BigNumberish] = [1, 0, 0, 0, 0];
    const fieldsIndex: [BigNumberish, BigNumberish, BigNumberish, BigNumberish, BigNumberish] = [2, 0, 0, 0, 0];
    const fieldsLogAddress: [string, string, string, string, string] = [
      '0x8ad599c3a0ff1de082011efddc58f1908eb6e6d8',
      '0x8ad599c3a0ff1de082011efddc58f1908eb6e6d8',
      '0x8ad599c3a0ff1de082011efddc58f1908eb6e6d8',
      '0x8ad599c3a0ff1de082011efddc58f1908eb6e6d8',
      '0x8ad599c3a0ff1de082011efddc58f1908eb6e6d8'
    ];
    const fieldsLogTopic: [string, string, string, string, string] = [
      '0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67',
      '0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67',
      '0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67',
      '0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67',
      '0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67'
    ];
    let logExtraInfo0: Brevis.LogExtraInfoStruct = {
      valueFromTopic: isFieldsInTopic[0],
      valueIndex: fieldsIndex[0],
      contractAddress: fieldsLogAddress[0],
      logTopic0: fieldsLogTopic[0]
    };

    let logExtraInfo1 = {
      valueFromTopic: isFieldsInTopic[1],
      valueIndex: fieldsIndex[1],
      contractAddress: fieldsLogAddress[1],
      logTopic0: fieldsLogTopic[1]
    };

    const fieldsLogIndex: [BigNumberish, BigNumberish, BigNumberish, BigNumberish, BigNumberish] = [4, 4, 4, 4, 4];
    const values: [BytesLike, BytesLike, BytesLike, BytesLike, BytesLike] = [
      '0x0000000000000000000000008df6872be6e53a2ace98b8c4411e052533efa637',
      '0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffffd978d4f',
      '0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffffd978d4f',
      '0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffffd978d4f',
      '0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffffd978d4f'
    ];
    const logs: [
      Brevis.LogInfoStruct,
      Brevis.LogInfoStruct,
      Brevis.LogInfoStruct,
      Brevis.LogInfoStruct,
      Brevis.LogInfoStruct
    ] = [
      { logExtraInfo: logExtraInfo0, logIndex: fieldsLogIndex[0], value: values[0] },
      { logExtraInfo: logExtraInfo1, logIndex: fieldsLogIndex[1], value: values[1] },
      { logExtraInfo: logExtraInfo1, logIndex: fieldsLogIndex[1], value: values[1] },
      { logExtraInfo: logExtraInfo1, logIndex: fieldsLogIndex[1], value: values[1] },
      { logExtraInfo: logExtraInfo1, logIndex: fieldsLogIndex[1], value: values[1] }
    ];

    let receipt = {
      blkNum: 18233760,
      receiptIndex: 75,
      logs: logs
    };

    let storage = {
      blockHash: '0x605c7b6408933648576c558f60459e18fcd26bddf66e28f580315cf137405d55',
      account: '0x881D40237659C251811CEC9c364ef91dC08D300C',
      slot: keccak256('0x0000000000000000000000000000000000000000000000000000000000000001'),
      slotValue: '0x0100000000000000000000000000000000000000000000000000000000000000',
      blockNumber: 18233760
    };
    let tx = {
      leafHash: '0x28df9deaf5d0f2d326b7ed277dbe0f59720e81afa1f684a4dfb6f06a29cfa972',
      blockHash: '0x605c7b6408933648576c558f60459e18fcd26bddf66e28f580315cf137405d55',
      blockNumber: 18233760,
      blockTime: 1695898391,
      leafRlpPrefix: '0xf9016120b9015d'
    };
    await app.setFieldLocations([logExtraInfo0, logExtraInfo1]);

    await brevisProof.setBatchTierVkHashes(
      ['0x28480547bb8aaa203cf9cadcca995cc6abda65c55f74c706134e6931b72cea8c'],
      [20]
    );
    await brevisProof.submitProof(1, proof, false);
    await expect(
      app.submit(
        '0x97cc677c7cb46dca48c1b2e181801720910c292e9009bcec63a77a0ed579e844',
        1,
        '0x0df1656fd4eb26023bb8062183e1059c9f55cdb00efe07199a64ff80b2b9dfcb',
        receipt,
        storage,
        tx
      )
    ).to.emit(app, 'Success');
  });

  it('should pass on custom app proof verification', async () => {
    const proof =
      '0x0028d3939b8a291d8a0aa5c4c25683120eaa4b65c4d914caa2c730e724875a880751827af1c3b5410765da6db4ec4842f68676e17bbb1f8aca82a00ff78c986b17cd79d36dc39121bca789137a55b0472571ceddd0acc17f8a126d150321eb580569bc2d22aaf6f76c7b0369a3d7dc5bbe040e9c7e4be8c98d7989084a762d8a05090b8e4926760ca52d424a903945aa27d8d43f810016cb776bdf0e74c2998a09f285728d8eed0ebd43020b02395aacbc50c3c53188a12f3c2a5447ea444ad11a10075f551f8aab2125c937e8c0a01fb5d3df380b8753d6d90433118e89b92422297a2c6891897be6149fa121ff9bfdf978972a716ccaa1fcd816270afb1d9f03703adff9ac652f748b8318ae5af481d5d215759e133fb9e5a0c48159e7a2ef290a912d0004e1af044dee62c4fd0f04a23e152dc1635a89df7d78b916337b4711a8db876102920b6d51d9ddd535d25981a0be88dc34c46073e9c126190ed73a2026cd953780d0ef6fcc2d69a96bff104bcb661918a4f6bb4c6862cd5e93eb2d08646494e28db9b84dbf61f1725e5174f7eaee5265971f0d92f33559810e442815a9cc96d5579d53c09ebebb1e8289156358a2dbd0568d9a5c634240312df1291468288056310c82aa4c01a7e12a10f8111a0560e72b700555479031b86c357d042989bb107bdd497bf3de4ed66aece2aa1fa9950ae65211f441e6df1b95af9b';

    await brevisProof.setBatchTierVkHashes(
      ['0x15a9cc96d5579d53c09ebebb1e8289156358a2dbd0568d9a5c634240312df129'],
      [20]
    );
    await brevisProof.submitProof(1, proof, true);
  });
});
