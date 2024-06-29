import { expect } from 'chai';
import { ContractRunner } from 'ethers';
import { ethers } from 'hardhat';

import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers';

import {
  EthStorageVerifier__factory,
  MockBlockChunks__factory,
  SlotValueVerifier,
  SlotValueVerifier__factory,
  VerifierGasReport,
  VerifierGasReport__factory,
} from '../../typechain';
import { convertByteArrayToHexString } from '../circuit-verifier/util';

async function deploySlotVerifierContract(admin: ContractRunner) {
  const syncerFactory = new MockBlockChunks__factory();
  const syncer = await syncerFactory.connect(admin).deploy();
  const factory = new SlotValueVerifier__factory();
  const contract = await factory.connect(admin).deploy(await syncer.getAddress());
  const verifierF = new EthStorageVerifier__factory();
  const verifier = await verifierF.connect(admin).deploy();
  await contract.updateVerifierAddress(1, await verifier.getAddress());

  const _factory = new VerifierGasReport__factory();
  const verifierGasReport = await _factory.connect(admin).deploy(await contract.getAddress());
  return { contract, verifierGasReport };
}

function getTestProof() {
  const publicInputs = [
    BigInt(convertByteArrayToHexString([103, 197, 210, 106, 230, 239, 0, 173, 207, 151, 13, 155, 24, 118, 240, 234])),
    BigInt(convertByteArrayToHexString([236, 65, 249, 77, 136, 183, 160, 41, 158, 157, 97, 9, 205, 217, 188, 216])),
    BigInt(convertByteArrayToHexString([230, 66, 26, 191, 243, 181, 187, 60, 128, 126, 39, 8, 155, 41, 116, 25])),
    BigInt(convertByteArrayToHexString([251, 9, 216, 152, 169, 76, 141, 172, 214, 149, 130, 94, 141, 128, 60, 56])),
    BigInt(convertByteArrayToHexString([194, 87, 90, 14, 158, 89, 60, 0, 249, 89, 248, 201, 47, 18, 219, 40])),
    BigInt(convertByteArrayToHexString([105, 195, 57, 90, 59, 5, 2, 208, 94, 37, 22, 68, 111, 113, 248, 91])),
    BigInt(
      convertByteArrayToHexString([255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255])
    ),
    BigInt(
      convertByteArrayToHexString([255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255])
    ),
    BigInt('17037800')
  ];

  const a = [
    BigInt('0x1ff096821d084175748931e8c3bc749e8a6d2c133dcc174535e3af55739510b7'),
    BigInt('0x06eb7bebb8746f6343376fb12eb9f96f8e35c36daadf59a38790bb7c545fb6c1')
  ];

  const b = [
    [
      BigInt('0x08617521d17a02aa3c8289b323eeb7e2698d64e12dc1016b06b14ac470b66cb9'),
      BigInt('0x1587e05e535e3b7292aae17f313a88a4bd4559068663086f5bc0e1efed9fe2bb')
    ],
    [
      BigInt('0x1d0b579ad36cff136336e75b4653bbed97618644d9ce170591e03219250b3933'),
      BigInt('0x03a160d716298b64b0f83d5850998276d1c89ee298d003633a4fd9145a605b5e')
    ]
  ];

  const c = [
    BigInt('0x04517a59525f5a40f70fb80dbe445178513dea08165542deefece2ebc39f6e6a'),
    BigInt('0x031e24b1b11384502dd7e3bad636002b86eada9613b67505d09529d593ca897b')
  ];

  const commitment = [BigInt('0x0'), BigInt('0x0')];

  const allData = [...a];
  allData.push(...b[0], ...b[1]);
  allData.push(...c);
  allData.push(...commitment);
  allData.push(...publicInputs);

  let allDataHex = '0x';
  for (let i = 0; i < allData.length; i++) {
    allDataHex = allDataHex + BigInt(allData[i]).toString(16).padStart(64, '0');
  }
  return allDataHex;
}

function getMockAuxiBlkVerifyInfo() {
  return '0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000';
}

describe('Slot Verifier Test', async () => {
  async function fixture() {
    const [admin] = await ethers.getSigners();
    const { contract, verifierGasReport } = await deploySlotVerifierContract(admin);
    return { admin, contract, verifierGasReport };
  }

  let contract: SlotValueVerifier;
  let verifierGasReport: VerifierGasReport;
  before(async () => {
    const res = await loadFixture(fixture);
    contract = res.contract;
    verifierGasReport = res.verifierGasReport;
  });

  it('should pass on verifySlotValue', async () => {
    const result = await contract.verifySlotValue(BigInt('1'), getTestProof(), getMockAuxiBlkVerifyInfo());
    await verifierGasReport.verifySlotValue(BigInt('1'), getTestProof(), getMockAuxiBlkVerifyInfo());

    expect(result.blkNum).to.equal(17037800);
    expect(result.blkHash).to.equal('0x67c5d26ae6ef00adcf970d9b1876f0eaec41f94d88b7a0299e9d6109cdd9bcd8');
    expect(result.addrHash).to.equal('0xe6421abff3b5bb3c807e27089b297419fb09d898a94c8dacd695825e8d803c38');
    expect(result.slotKeyHash).to.equal('0xc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b');
    expect(result.slotValue).to.equal('0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff');
  });
});
