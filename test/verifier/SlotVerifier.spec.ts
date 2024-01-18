import { assert } from 'console';
import { Fixture } from 'ethereum-waffle';
import { BigNumber, Wallet } from 'ethers';
import { ethers, waffle } from 'hardhat';
import {
  MockBlockChunks__factory,
  MockZkVerifier__factory,
  SlotValueVerifier,
  SlotValueVerifier__factory,
  VerifierGasReport
} from '../../typechain';
import { hexToBytes } from '../util';
import { convertByteArrayToHexString } from '../circuit-verifier/util';
import { expect } from 'chai';

async function deploySlotVerifierContract(admin: Wallet) {
  const syncerFactory = await ethers.getContractFactory<MockBlockChunks__factory>('MockBlockChunks');
  const syncer = await syncerFactory.connect(admin).deploy();
  const factory = await ethers.getContractFactory<SlotValueVerifier__factory>('SlotValueVerifier');
  const contract = await factory.connect(admin).deploy(syncer.address);
  const verifierF = await ethers.getContractFactory<MockZkVerifier__factory>('EthStorageVerifier');
  const verifier = await verifierF.connect(admin).deploy();
  await contract.updateVerifierAddress(1, verifier.address);

  const _factory = await ethers.getContractFactory('VerifierGasReport');
  const verifierGasReport = (await _factory.connect(admin).deploy(contract.address)) as VerifierGasReport;
  return { contract, verifierGasReport };
}

function getTestProof() {
  const publicInputs = [
    BigNumber.from(
      convertByteArrayToHexString([103, 197, 210, 106, 230, 239, 0, 173, 207, 151, 13, 155, 24, 118, 240, 234])
    ),
    BigNumber.from(
      convertByteArrayToHexString([236, 65, 249, 77, 136, 183, 160, 41, 158, 157, 97, 9, 205, 217, 188, 216])
    ),
    BigNumber.from(
      convertByteArrayToHexString([230, 66, 26, 191, 243, 181, 187, 60, 128, 126, 39, 8, 155, 41, 116, 25])
    ),
    BigNumber.from(
      convertByteArrayToHexString([251, 9, 216, 152, 169, 76, 141, 172, 214, 149, 130, 94, 141, 128, 60, 56])
    ),
    BigNumber.from(convertByteArrayToHexString([194, 87, 90, 14, 158, 89, 60, 0, 249, 89, 248, 201, 47, 18, 219, 40])),
    BigNumber.from(convertByteArrayToHexString([105, 195, 57, 90, 59, 5, 2, 208, 94, 37, 22, 68, 111, 113, 248, 91])),
    BigNumber.from(
      convertByteArrayToHexString([255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255])
    ),
    BigNumber.from(
      convertByteArrayToHexString([255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255])
    ),
    BigNumber.from('17037800')
  ];

  const a = [
    BigNumber.from('0x1ff096821d084175748931e8c3bc749e8a6d2c133dcc174535e3af55739510b7'),
    BigNumber.from('0x06eb7bebb8746f6343376fb12eb9f96f8e35c36daadf59a38790bb7c545fb6c1')
  ];

  const b = [
    [
      BigNumber.from('0x08617521d17a02aa3c8289b323eeb7e2698d64e12dc1016b06b14ac470b66cb9'),
      BigNumber.from('0x1587e05e535e3b7292aae17f313a88a4bd4559068663086f5bc0e1efed9fe2bb')
    ],
    [
      BigNumber.from('0x1d0b579ad36cff136336e75b4653bbed97618644d9ce170591e03219250b3933'),
      BigNumber.from('0x03a160d716298b64b0f83d5850998276d1c89ee298d003633a4fd9145a605b5e')
    ]
  ];

  const c = [
    BigNumber.from('0x04517a59525f5a40f70fb80dbe445178513dea08165542deefece2ebc39f6e6a'),
    BigNumber.from('0x031e24b1b11384502dd7e3bad636002b86eada9613b67505d09529d593ca897b')
  ];

  const commitment = [BigNumber.from('0x0'), BigNumber.from('0x0')];

  const allData = [...a];
  allData.push(...b[0], ...b[1]);
  allData.push(...c);
  allData.push(...commitment);
  allData.push(...publicInputs);

  let allDataHex = '';
  for (let i = 0; i < allData.length; i++) {
    allDataHex = allDataHex + BigNumber.from(allData[i]).toHexString().slice(2).padStart(64, '0');
  }
  return hexToBytes(allDataHex);
}

function getMockAuxiBlkVerifyInfo() {
  return new Array(8 * 32 + 4).fill(0);
}

describe('Slot Verifier Test', async () => {
  function loadFixture<T>(fixture: Fixture<T>): Promise<T> {
    const provider = waffle.provider;
    return waffle.createFixtureLoader(provider.getWallets(), provider)(fixture);
  }

  async function fixture([admin]: Wallet[]) {
    const { contract, verifierGasReport } = await deploySlotVerifierContract(admin);
    return { admin, contract, verifierGasReport };
  }

  let contract: SlotValueVerifier;
  let verifierGasReport: VerifierGasReport;
  let admin: Wallet;
  before(async () => {
    const res = await loadFixture(fixture);
    contract = res.contract;
    verifierGasReport = res.verifierGasReport;
    admin = res.admin;
  });

  it('should pass on verifySlotValue', async () => {
    const result = await contract.verifySlotValue(BigNumber.from('1'), getTestProof(), getMockAuxiBlkVerifyInfo());
    await verifierGasReport.verifySlotValue(BigNumber.from('1'), getTestProof(), getMockAuxiBlkVerifyInfo());

    expect(result.blkNum).to.equal(17037800);
    expect(result.blkHash).to.equal('0x67c5d26ae6ef00adcf970d9b1876f0eaec41f94d88b7a0299e9d6109cdd9bcd8');
    expect(result.addrHash).to.equal('0xe6421abff3b5bb3c807e27089b297419fb09d898a94c8dacd695825e8d803c38');
    expect(result.slotKeyHash).to.equal('0xc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b');
    expect(result.slotValue).to.equal('0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff');
  });
});
