import { Fixture } from 'ethereum-waffle';
import { ethers, waffle } from 'hardhat';
import { PoALightClient, MockTendermintLightClient } from '../../typechain';

import { BigNumber, Wallet } from 'ethers';
import { expect } from 'chai';
import { keccak256, RLP } from 'ethers/lib/utils';

describe('PoALightClient Test', async () => {
  function loadFixture<T>(fixture: Fixture<T>): Promise<T> {
    const provider = waffle.provider;
    return waffle.createFixtureLoader(provider.getWallets(), provider)(fixture);
  }

  async function fixture([admin]: Wallet[]) {
    const { poaLC, mockTLC } = await deployLib(admin);
    return { admin, poaLC, mockTLC };
  }

  let poaLC: PoALightClient;
  let mockTLC: MockTendermintLightClient;
  let admin: Wallet;

  beforeEach(async () => {
    const res = await loadFixture(fixture);
    admin = res.admin as Wallet;
    poaLC = res.poaLC as PoALightClient;
    mockTLC = res.mockTLC as MockTendermintLightClient;
  });

  async function deployLib(admin: Wallet) {
    const mockTLCFactory = await ethers.getContractFactory('MockTendermintLightClient');
    const mockTLC = (await mockTLCFactory
      .connect(admin)
      .deploy('0xffffffffffffffffffffffffffffffffffffffff')) as MockTendermintLightClient;

    const poaLCFactory = await ethers.getContractFactory('PoALightClient');
    const poaLC = (await poaLCFactory.connect(admin).deploy(mockTLC.address)) as PoALightClient;

    return { mockTLC, poaLC };
  }
  /*
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
    let rlpInfo = RLP.encode([
      '0x02',
      '0xd983010000846765746889676f312e31322e3137856c696e7578000000000000',
      '0x01c9c380',
      '0x00',
      '0xc3fa2927a8e5b7cfbd575188a30c34994d3356607deb4c10d7fefe0dd5cdcc83',
      '0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000',
      '0x35552c16704d214347f29fa77f77da6d75d7c752',
      '0x0000000000000000000000000000000000000000000000000000000000000000',
      '0x0000000000000000',
      '0x68b3',
      '0xbf4d16769b8fd946394957049eef29ed938da92454762fc6ac65e0364ea004c7',
      '0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421',
      '0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347',
      '0x0261',
      '0x7b5a72075082c31ec909afe5c5df032b6e7f19c686a9a408a2cb6b75dec072a3',
      '0x5f080818',
      '0xd167',
      [],
      '0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421',
      []
    ]);
    // console.log(hexlify(message), signature)

    let messageHash = keccak256(rlpInfo);

    let message = ethers.utils.arrayify(messageHash);
    let signature = await admin.signMessage(message);

    let header = {
      difficulty: '0x2',
      extraData: '0xd983010000846765746889676f312e31322e3137856c696e7578000000000000' + signature.replace('0x', ''),
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
    await mockTLC.updateSigner(admin.address);

    expect(await poaLC.updateHeader(header));

    const m = await poaLC.finalizedExecutionStateRootAndSlot();
    expect(m[0]).to.be.eqls('0x7b5a72075082c31ec909afe5c5df032b6e7f19c686a9a408a2cb6b75dec072a3');
    expect(m[1]).to.be.eqls(BigNumber.from('0x68b3'));

    await expect(poaLC.updateHeader(header)).to.be.revertedWith('PoALightClient: invalid block number');
  });
  */
});
