import { assert } from 'console';
import { Fixture } from 'ethereum-waffle';
import { BigNumber, BigNumberish, Wallet } from 'ethers';
import { keccak256 } from 'ethers/lib/utils';
import { ethers, waffle } from 'hardhat';
import {
  MockBlockChunks__factory,
  MockZkVerifier__factory,
  TxVerifier,
  TxVerifier__factory,
  VerifierGasReport
} from '../../typechain';
import { hexToBytes, splitHash } from '../util';

async function deployTxVerifierContract(admin: Wallet) {
  const syncerFactory = await ethers.getContractFactory<MockBlockChunks__factory>('MockBlockChunks');
  const syncer = await syncerFactory.connect(admin).deploy();
  const factory = await ethers.getContractFactory<TxVerifier__factory>('TxVerifier');
  const contract = await factory.connect(admin).deploy(syncer.address);
  const verifierF = await ethers.getContractFactory<MockZkVerifier__factory>('MockZkVerifier');
  const verifier = await verifierF.connect(admin).deploy();
  await contract.updateVerifierAddress(1, verifier.address);

  const _factory = await ethers.getContractFactory('VerifierGasReport');
  const verifierGasReport = (await _factory.connect(admin).deploy(contract.address)) as VerifierGasReport;
  return { contract, verifierGasReport };
}

function getTestProof(leafHash: string) {
  const mockBlkHash = '0x88bd78528ea4fd5c232978ce51e43f41f0d76ce56e331147c1c9611282308799';
  const input = [...splitHash(leafHash), ...splitHash(mockBlkHash), 17086605, 1681980179];
  const a: [BigNumberish, BigNumberish] = [
    BigNumber.from('0x091712d21a7fb14be9027310e2cbcc7d9d4132d6422598586a4a1e481d69d234'),
    BigNumber.from('0x16c655962badf7228ca62ae8d5674c1bdf10cd4edbd880e039a54ef6e2e55eab')
  ];
  const b: [[BigNumberish, BigNumberish], [BigNumberish, BigNumberish]] = [
    [
      BigNumber.from('0x0798c4c36b7d42124034a55327f8af1a2ec29ecedf1dd7c8b72690164f7d7841'),
      BigNumber.from('0x0398de45e5843c72045fc9d01479c34ea4e6eebfbc8cbb4d13e35f36191c83ca')
    ],
    [
      BigNumber.from('0x1e8324d656f1700f87a9b7f8f06b081f5ed8e7dd363a56fad209997815ea54b6'),
      BigNumber.from('0x231a2b40a5147fcf71ab6d80de168e6a30cb26b87b14ae0ab7c3c9f1bd355513')
    ]
  ];
  const c: [BigNumberish, BigNumberish] = [
    BigNumber.from('0x23d9a9af2e7544e6c0941cdf92115b40ddbd2b0a6bd33ef343823ea4c4e9ec11'),
    BigNumber.from('0x2eb59836c9c43e2a6a6abc07ca138cb9a70588dd72befa183a4d8af4bec4b44c')
  ];
  const commit: [BigNumberish, BigNumberish] = ['0', '0'];
  const allData = [...a];
  allData.push(...b[0], ...b[1]);
  allData.push(...c);
  allData.push(...commit);
  allData.push(...input);

  let allDataHex = '';
  for (let i = 0; i < allData.length; i++) {
    allDataHex = allDataHex + BigNumber.from(allData[i]).toHexString().slice(2).padStart(64, '0');
  }

  allDataHex = allDataHex + 'f901db20b901d7';

  return hexToBytes(allDataHex);
}

function getMockAuxiBlkVerifyInfo() {
  return new Array(8 * 32 + 4).fill(0);
}

describe('Tx Verifier Test', async () => {
  function loadFixture<T>(fixture: Fixture<T>): Promise<T> {
    const provider = waffle.provider;
    return waffle.createFixtureLoader(provider.getWallets(), provider)(fixture);
  }

  async function fixture([admin]: Wallet[]) {
    const { contract, verifierGasReport } = await deployTxVerifierContract(admin);
    return { admin, contract, verifierGasReport };
  }

  let contract: TxVerifier;
  let verifierGasReport: VerifierGasReport;
  let admin: Wallet;
  before(async () => {
    const res = await loadFixture(fixture);
    contract = res.contract;
    verifierGasReport = res.verifierGasReport;
    admin = res.admin;
  });

  it('should pass on decodeTx', async () => {
    const txRaw =
      '0x02f8720183016004808508f6f1387e82523f94ebec795c9c8bbd61ffc14a6662944748f299cacf8801077fd38ccda23c80c001a0df08ec564087b1d4665417df115b4bc0fe56857d6e13089c2c7b423dc4200a2fa06ba2214bfc68ad68bb9f15643a4ff0b13a6536740b8b37dfecac73558591ed7d';
    const tx = await contract.decodeTx(txRaw);
    assert(tx.from == '0x1f9090aaE28b8a3dCeaDf281B0F12828e676c326');
  });
  it('should pass on decodeTx', async () => {
    const txRaw =
      '0x02f903f4018205aa84773594008505197c52d98302a9d394ef1c6e67703c7bd7107eed8303fbe6ec2554bf6b80b9038424856bc30000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000020a000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000001c00000000000000000000000000000000000000000000000000000000000000160000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000ffffffffffffffffffffffffffffffffffffffff0000000000000000000000000000000000000000000000000000000064ae1e120000000000000000000000000000000000000000000000000000000000000002000000000000000000000000ef1c6e67703c7bd7107eed8303fbe6ec2554bf6b000000000000000000000000000000000000000000000000000000006486981a00000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000000413385e9e55f2bb34e442699933afcb89a5f34b1a057dabed18a3b513b83977f1b48f0cc5de7cc7de385117f85545a39431dd747dbd31bfd38fec4edc3774a9e6b1b00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000014fe21f94665f173d99503d254bea54b5081d49d0000000000000000000000000000000000000000000000000b1a2bc2ec50000000000000000000000000000000000000000000000000000000000000528690fa00000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002bc02aaa39b223fe8d0a0e5c4f27ead9083c756cc20001f4a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48000000000000000000000000000000000000000000c080a0c361e12e6a42e94b88c23145019f1d066c4e4c93acbd35649fcb7bca0bbb6312a018285165c3d39313832d1d6223d1178fe45becd961dcd51672d05f20f13ef7de';
    const tx = await contract.decodeTx(txRaw);
    assert(tx.from == '0x14fe21f94665F173d99503d254bEa54B5081d49D');
  });
  it('should pass on verifyTx', async () => {
    const txRaw =
      '0x02f901d30181b48405f5e100850faf9e23de830962b494c36442b4a4522e871399cd717abdd847ab11fe8880b90164883164560000000000000000000000006982508145454ce325ddbe47a25d4ec3d2311933000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb480000000000000000000000000000000000000000000000000000000000002710fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff27660fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9ad400000000000000000000000000000000000000000024c98fcd63f9edd999c2bbe0000000000000000000000000000000000000000000000000000000023c34600000000000000000000000000000000000000000002490b8629cad414ed17c31a0000000000000000000000000000000000000000000000000000000023ac5b490000000000000000000000006880129a290043e85eb6c67c3838d961a85956790000000000000000000000000000000000000000000000000000000064410203c080a0f78c707ba62590c6e4b222ea33c73c585ac7b1179397adf5aa0f80c7c0b63045a01a60736f6bc0effd5005723ae22f65ef92749eebbeed222f7149a8f091ef82bc';
    const leafHash = '0x958c0c028b7a8fe0a3d5961620582cad1f557604937a104f77246246118a24c7';
    const tx = await contract.verifyTx(txRaw, getTestProof(leafHash), getMockAuxiBlkVerifyInfo());
    await verifierGasReport.verifyTx(txRaw, getTestProof(leafHash), getMockAuxiBlkVerifyInfo()); // report gas
    assert(tx.from == '0x6880129a290043e85Eb6c67c3838d961A8595679');
    assert(tx.blkHash == '0x88bd78528ea4fd5c232978ce51e43f41f0d76ce56e331147c1c9611282308799');
  });
});
