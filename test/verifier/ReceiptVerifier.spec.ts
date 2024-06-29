import { expect } from 'chai';
import { assert } from 'console';
import { BigNumberish, ContractRunner } from 'ethers';
import { ethers } from 'hardhat';

import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers';

import {
  MockBlockChunks__factory,
  MockZkVerifier__factory,
  ReceiptVerifier,
  ReceiptVerifier__factory,
  VerifierGasReport,
} from '../../typechain';
import { VerifierGasReport__factory } from '../../typechain/factories/contracts/test-helper/VerifierGasReport.sol/VerifierGasReport__factory';
import { splitHash } from '../util';

async function deployReceiptVerifierContract(admin: ContractRunner) {
  const syncerFactory = new MockBlockChunks__factory();
  const syncer = await syncerFactory.connect(admin).deploy();
  const factory = new ReceiptVerifier__factory();
  const contract = await factory.connect(admin).deploy(await syncer.getAddress());
  const verifierF = new MockZkVerifier__factory();
  const verifier = await verifierF.connect(admin).deploy();
  await contract.updateVerifierAddress(1, await verifier.getAddress());

  const _factory = new VerifierGasReport__factory();
  const verifierGasReport = (await _factory.connect(admin).deploy(await contract.getAddress())) as VerifierGasReport;
  return { contract, verifierGasReport };
}

function getTestProof(leafHash: string) {
  const mockBlkHash = '0x88bd78528ea4fd5c232978ce51e43f41f0d76ce56e331147c1c9611282308799';
  const input = [...splitHash(leafHash), ...splitHash(mockBlkHash), 17086605, 1681980179];
  const a: [BigNumberish, BigNumberish] = [
    BigInt('0x091712d21a7fb14be9027310e2cbcc7d9d4132d6422598586a4a1e481d69d234'),
    BigInt('0x16c655962badf7228ca62ae8d5674c1bdf10cd4edbd880e039a54ef6e2e55eab')
  ];
  const b: [[BigNumberish, BigNumberish], [BigNumberish, BigNumberish]] = [
    [
      BigInt('0x0798c4c36b7d42124034a55327f8af1a2ec29ecedf1dd7c8b72690164f7d7841'),
      BigInt('0x0398de45e5843c72045fc9d01479c34ea4e6eebfbc8cbb4d13e35f36191c83ca')
    ],
    [
      BigInt('0x1e8324d656f1700f87a9b7f8f06b081f5ed8e7dd363a56fad209997815ea54b6'),
      BigInt('0x231a2b40a5147fcf71ab6d80de168e6a30cb26b87b14ae0ab7c3c9f1bd355513')
    ]
  ];
  const c: [BigNumberish, BigNumberish] = [
    BigInt('0x23d9a9af2e7544e6c0941cdf92115b40ddbd2b0a6bd33ef343823ea4c4e9ec11'),
    BigInt('0x2eb59836c9c43e2a6a6abc07ca138cb9a70588dd72befa183a4d8af4bec4b44c')
  ];
  const commit: [BigNumberish, BigNumberish] = ['0', '0'];
  const allData = [...a];
  allData.push(...b[0], ...b[1]);
  allData.push(...c);
  allData.push(...commit);
  allData.push(...input);

  let allDataHex = '0x';
  for (let i = 0; i < allData.length; i++) {
    allDataHex = allDataHex + BigInt(allData[i]).toString(16).padStart(64, '0');
  }

  // leafRlpPrefix
  allDataHex = allDataHex + '0000000000000001'; // chain id
  allDataHex = allDataHex + 'f9056520b90561';

  return allDataHex;
}

function getMockAuxiBlkVerifyInfo() {
  return '0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000';
}

describe('Receipt Verifier Test', async () => {
  async function fixture() {
    const [admin] = await ethers.getSigners();
    const { contract, verifierGasReport } = await deployReceiptVerifierContract(admin);
    return { contract, verifierGasReport };
  }

  let contract: ReceiptVerifier;
  let verifierGasReport: VerifierGasReport;
  before(async () => {
    const res = await loadFixture(fixture);
    contract = res.contract;
    verifierGasReport = res.verifierGasReport;
  });

  it('should pass on decodeReceipt', async () => {
    const receiptRaw =
      '0x02f9055d018388ee3cb90100008000000000000000001000000000000000000000000000000000000000010000000000000000000000008000000000000101000000000004000000002000000000000a0000800008000008000000000000000000000000000048000000000000000000020000000200000000000800000000000800000000000010000000010000000008000000000000000000000000002000010000000020000200000000020000000000200001000000000000000000020000200000000000000000000240000002000000000000081000000000000000000000000000001000000020020010000001000000000000000000000000000000000000400000000040000800f90452f89b946982508145454ce325ddbe47a25d4ec3d2311933f863a0ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa00000000000000000000000006880129a290043e85eb6c67c3838d961a8595679a0000000000000000000000000cee31c846cbf003f4ceb5bbd234cba03c6e940c7a00000000000000000000000000000000000000000024c98fcd60a663566c4ef68f89b946982508145454ce325ddbe47a25d4ec3d2311933f863a08c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925a00000000000000000000000006880129a290043e85eb6c67c3838d961a8595679a0000000000000000000000000c36442b4a4522e871399cd717abdd847ab11fe88a0000000000000000000000000000000000000000002af5ddc25ade6bdae93ddb7f89b94a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48f863a0ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa00000000000000000000000006880129a290043e85eb6c67c3838d961a8595679a0000000000000000000000000cee31c846cbf003f4ceb5bbd234cba03c6e940c7a00000000000000000000000000000000000000000000000000000000023c34600f9011d94cee31c846cbf003f4ceb5bbd234cba03c6e940c7f884a07a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bdea0000000000000000000000000c36442b4a4522e871399cd717abdd847ab11fe88a0fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff27660a0fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9ad40b880000000000000000000000000c36442b4a4522e871399cd717abdd847ab11fe880000000000000000000000000000000000000000000000000e1db05eef5230540000000000000000000000000000000000000000024c98fcd60a663566c4ef680000000000000000000000000000000000000000000000000000000023c34600f89c94c36442b4a4522e871399cd717abdd847ab11fe88f884a0ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa00000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000006880129a290043e85eb6c67c3838d961a8595679a0000000000000000000000000000000000000000000000000000000000007847d80f8bb94c36442b4a4522e871399cd717abdd847ab11fe88f842a03067048beee31b25b2f1681f88dac838c8bba36af25bfb2b7cf7473a5847e35fa0000000000000000000000000000000000000000000000000000000000007847db8600000000000000000000000000000000000000000000000000e1db05eef5230540000000000000000000000000000000000000000024c98fcd60a663566c4ef680000000000000000000000000000000000000000000000000000000023c34600';
    const receiptInfo = await contract.decodeReceipt(receiptRaw);
    assert(receiptInfo.logs[0].addr == '0x6982508145454Ce325dDbE47a25d4ec3d2311933');
  });

  it('should pass on verifyReceipt', async () => {
    const receiptRaw =
      '0x02f9055d018388ee3cb90100008000000000000000001000000000000000000000000000000000000000010000000000000000000000008000000000000101000000000004000000002000000000000a0000800008000008000000000000000000000000000048000000000000000000020000000200000000000800000000000800000000000010000000010000000008000000000000000000000000002000010000000020000200000000020000000000200001000000000000000000020000200000000000000000000240000002000000000000081000000000000000000000000000001000000020020010000001000000000000000000000000000000000000400000000040000800f90452f89b946982508145454ce325ddbe47a25d4ec3d2311933f863a0ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa00000000000000000000000006880129a290043e85eb6c67c3838d961a8595679a0000000000000000000000000cee31c846cbf003f4ceb5bbd234cba03c6e940c7a00000000000000000000000000000000000000000024c98fcd60a663566c4ef68f89b946982508145454ce325ddbe47a25d4ec3d2311933f863a08c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925a00000000000000000000000006880129a290043e85eb6c67c3838d961a8595679a0000000000000000000000000c36442b4a4522e871399cd717abdd847ab11fe88a0000000000000000000000000000000000000000002af5ddc25ade6bdae93ddb7f89b94a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48f863a0ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa00000000000000000000000006880129a290043e85eb6c67c3838d961a8595679a0000000000000000000000000cee31c846cbf003f4ceb5bbd234cba03c6e940c7a00000000000000000000000000000000000000000000000000000000023c34600f9011d94cee31c846cbf003f4ceb5bbd234cba03c6e940c7f884a07a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bdea0000000000000000000000000c36442b4a4522e871399cd717abdd847ab11fe88a0fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff27660a0fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9ad40b880000000000000000000000000c36442b4a4522e871399cd717abdd847ab11fe880000000000000000000000000000000000000000000000000e1db05eef5230540000000000000000000000000000000000000000024c98fcd60a663566c4ef680000000000000000000000000000000000000000000000000000000023c34600f89c94c36442b4a4522e871399cd717abdd847ab11fe88f884a0ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa00000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000006880129a290043e85eb6c67c3838d961a8595679a0000000000000000000000000000000000000000000000000000000000007847d80f8bb94c36442b4a4522e871399cd717abdd847ab11fe88f842a03067048beee31b25b2f1681f88dac838c8bba36af25bfb2b7cf7473a5847e35fa0000000000000000000000000000000000000000000000000000000000007847db8600000000000000000000000000000000000000000000000000e1db05eef5230540000000000000000000000000000000000000000024c98fcd60a663566c4ef680000000000000000000000000000000000000000000000000000000023c34600';
    const leafHash = '0x60144e033b2f048cab71d95b8ed3d0d0efc93de5797ab5369f14c52ec4b1f2e2';
    const receipt = await contract.verifyReceipt(receiptRaw, getTestProof(leafHash), getMockAuxiBlkVerifyInfo());
    await verifierGasReport.verifyReceipt(receiptRaw, getTestProof(leafHash), getMockAuxiBlkVerifyInfo()); // report gas
    assert(receipt.logs[0].addr == '0x6982508145454Ce325dDbE47a25d4ec3d2311933');
    expect(receipt).to.emit(contract, 'VerifiedReceipt');
  });
});
