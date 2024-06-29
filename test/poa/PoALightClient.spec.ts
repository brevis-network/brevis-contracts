import { expect } from 'chai';
import { ContractRunner, encodeRlp, getBytes, keccak256 } from 'ethers';
import { ethers } from 'hardhat';

import { HardhatEthersSigner } from '@nomicfoundation/hardhat-ethers/signers';
import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers';

import {
  MockTendermintLightClient,
  MockTendermintLightClient__factory,
  PoALightClient,
  PoALightClient__factory
} from '../../typechain';

describe('PoALightClient Test', async () => {
  async function fixture() {
    const [admin] = await ethers.getSigners();
    const { poaLC, mockTLC } = await deployLib(admin);
    return { admin, poaLC, mockTLC };
  }

  let poaLC: PoALightClient;
  let mockTLC: MockTendermintLightClient;
  let admin: HardhatEthersSigner;

  beforeEach(async () => {
    const res = await loadFixture(fixture);
    admin = res.admin;
    poaLC = res.poaLC;
    mockTLC = res.mockTLC;
  });

  async function deployLib(admin: ContractRunner) {
    const mockTLCFactory = new MockTendermintLightClient__factory();
    const mockTLC = await mockTLCFactory.connect(admin).deploy('0xffffffffffffffffffffffffffffffffffffffff');

    const poaLCFactory = new PoALightClient__factory();
    const mockTLCAddress = await mockTLC.getAddress();
    const poaLC = await poaLCFactory.connect(admin).deploy(mockTLCAddress);

    return { mockTLC, poaLC };
  }

  it('should fail to update header', async () => {
    let extraData =
      '0xd983010000846765746889676f312e31322e3137856c696e7578000000000000c3daa60d95817e2789de3eafd44dc354fe804bf5f08059cde7c86bc1215941d022bf9609ca1dee2881baf2144aa93fc80082e6edd0b9f8eac16f327e7d59f16500';
    const vSig = extraData.substring(extraData.length - 2, extraData.length);

    if (vSig == '00') {
      extraData = extraData.substring(0, extraData.length - 2) + '1b';
    } else if (vSig == '01') {
      extraData = extraData.substring(0, extraData.length - 2) + '1c';
    }

    let header = {
      difficulty: '0x2',
      extraData: extraData,
      gasLimit: '0x1c9c380',
      gasUsed: '0x0',
      hash: '0xc3fa2927a8e5b7cfbd575188a30c34994d3356607deb4c10d7fefe0dd5cdcc83',
      logsBloom:
        '0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000',
      miner: '0x35552c16704d214347f29fa77f77da6d75d7c752',
      mixHash: '0x0000000000000000000000000000000000000000000000000000000000000000',
      nonce: '0x0000000000000000',
      number: '0x68b3',
      parentHash: '0xbf4d16769b8fd946394957049eef29ed938da92454762fc6ac65e0364ea004c7',
      receiptsRoot: '0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421',
      sha3Uncles: '0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347',
      size: '0x261',
      stateRoot: '0x7b5a72075082c31ec909afe5c5df032b6e7f19c686a9a408a2cb6b75dec072a3',
      timestamp: '0x5f080818',
      totalDifficulty: '0xd167',
      transactions: [],
      transactionsRoot: '0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421',
      uncles: []
    };
    await expect(poaLC.updateHeader(header)).to.be.revertedWith('PoALightClient: invalid signer address');
  });

  it('should pass verify header', async () => {
    let rlpInfo = encodeRlp([
      '0x61',
      '0xbf4d16769b8fd946394957049eef29ed938da92454762fc6ac65e0364ea004c7',
      '0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347',
      '0x35552c16704d214347f29fa77f77da6d75d7c752',
      '0x7b5a72075082c31ec909afe5c5df032b6e7f19c686a9a408a2cb6b75dec072a3',
      '0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421',
      '0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421',
      '0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000',
      '0x02',
      '0x68b3',
      '0x01c9c380',
      '0x00',
      '0x5f080818',
      '0xd983010000846765746889676f312e31322e3137856c696e7578000000000000',
      '0x0000000000000000000000000000000000000000000000000000000000000000',
      '0x0000000000000000'
    ]);
    // console.log(`rlpInfo: ${rlpInfo}`, rlpInfo);

    let messageHash = keccak256(rlpInfo);
    // console.log(`messageHash: ${messageHash}`);

    let message = getBytes(messageHash);
    let signature = await admin.signMessage(message);
    // console.log(`signature: ${signature}`);

    let header = {
      parentHash: '0xbf4d16769b8fd946394957049eef29ed938da92454762fc6ac65e0364ea004c7',
      sha3Uncles: '0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347',
      miner: '0x35552c16704d214347f29fa77f77da6d75d7c752',
      stateRoot: '0x7b5a72075082c31ec909afe5c5df032b6e7f19c686a9a408a2cb6b75dec072a3',
      transactionsRoot: '0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421',
      receiptsRoot: '0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421',
      logsBloom:
        '0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000',
      difficulty: '0x02',
      number: '0x68b3',
      gasLimit: '0x1c9c380',
      gasUsed: '0x00',
      timestamp: '0x5f080818',
      extraData: '0xd983010000846765746889676f312e31322e3137856c696e7578000000000000' + signature.replace('0x', ''),
      mixHash: '0x0000000000000000000000000000000000000000000000000000000000000000',
      nonce: '0x0000000000000000'
    };
    await mockTLC.updateSigner(admin.address);

    expect(await poaLC.updateHeader(header));

    const m = await poaLC.finalizedExecutionStateRootAndSlot();
    expect(m[0]).to.be.eqls('0x7b5a72075082c31ec909afe5c5df032b6e7f19c686a9a408a2cb6b75dec072a3');
    expect(m[1]).to.be.eqls(BigInt('0x68b3'));

    await expect(poaLC.updateHeader(header)).to.be.revertedWith('PoALightClient: invalid block number');
  });
});
