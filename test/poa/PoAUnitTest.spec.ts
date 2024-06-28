import { ethers } from 'hardhat';
import { PoALibTest } from '../../typechain';

import { Wallet, encodeRlp, solidityPackedKeccak256, toBeArray } from 'ethers';
import { expect } from 'chai';
import { randomInt } from 'crypto';
import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers';
import { HardhatEthersSigner } from '@nomicfoundation/hardhat-ethers/signers';

describe('PoAUnitTest', async () => {
  async function fixture() {
    const [admin] = await ethers.getSigners();
    const _poaUintTest = await deployLib(admin);
    return { admin, _poaUintTest };
  }

  function hex2Bytes(hexString: string): number[] {
    let hex = hexString;
    const result = [];
    if (hex.substr(0, 2) === '0x') {
      hex = hex.slice(2);
    }
    if (hex.length % 2 === 1) {
      hex = '0' + hex;
    }
    for (let i = 0; i < hex.length; i += 2) {
      result.push(parseInt(hex.substr(i, 2), 16));
    }
    return result;
  }

  let poaUintTest: PoALibTest;
  let admin: HardhatEthersSigner;

  beforeEach(async () => {
    const res = await loadFixture(fixture);
    poaUintTest = res._poaUintTest as PoALibTest;
    admin = res.admin as HardhatEthersSigner;
  });

  async function deployLib(admin: HardhatEthersSigner) {
    const factory = await ethers.getContractFactory('PoALibTest');
    const poaUintTest = (await factory.connect(admin).deploy()) as PoALibTest;
    return poaUintTest;
  }

  it('should pass retrive part of bytes', async () => {
    const testBytes = '0x7894745829abbdfe75814218613412afed1238907ba0';
    const bytesLength = testBytes.length / 2 - 1;

    const inboundEnd = 10;
    const inboundStart = 8;
    expect(await poaUintTest.mockRange(testBytes, inboundStart, inboundEnd)).to.be.eql(
      '0x' + testBytes.replace('0x', '').substring(inboundStart * 2, inboundEnd * 2)
    );

    const greaterStart = 12;
    expect(await poaUintTest.mockRange(testBytes, greaterStart, inboundEnd)).to.be.eql('0x');

    try {
      await poaUintTest.mockRange(testBytes, -1, inboundEnd);
    } catch (error) {
      // assert(JSON.stringify(error).includes('value out-of-bounds (argument="from"'))
    }

    await expect(poaUintTest.mockRange(testBytes, bytesLength, bytesLength + 1)).to.be.revertedWith(
      'Memory: from out of bounds'
    );

    await expect(poaUintTest.mockRange(testBytes, 0, bytesLength + 1)).to.be.revertedWith('Memory: to out of bounds');
  });

  it('should copy part of bytes', async () => {
    expect(await poaUintTest.mockCopy('0xabcd', 1)).to.be.eql('0xab');
  });

  it('should be able to encode uint value', async () => {
    const input = randomInt(0, 19231231314);
    let inputInHex = input.toString(16);
    if (inputInHex.length % 2 == 1) {
      inputInHex = '0' + inputInHex;
    }
    const binaryData = await poaUintTest.mockToBinary(input);
    expect(binaryData).to.be.eql('0x' + inputInHex);
    expect(await poaUintTest.mockWriteUint(input)).to.be.eql(encodeRlp(binaryData));
    expect(await poaUintTest.mockWriteUint(input + 1)).to.not.be.eql(encodeRlp(binaryData));
  });

  it('should be able to convert uint256 max to hex value', async () => {
    expect(await poaUintTest.mockUint256MaxToBinary()).to.be.eql(
      '0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff'
    );
  });

  it('should be able to encode address value', async () => {
    const address = '0xffffffffffffffffffffffffffffffffffffffff';
    expect(await poaUintTest.mockWriteAddress(address)).to.be.eql(encodeRlp(address));
    expect(await poaUintTest.mockWriteAddress(address)).to.be.not.eql(
      encodeRlp('0x0000000000000000000000000000000000000000')
    );
  });

  it('should be able to encode boolean value', async () => {
    expect(await poaUintTest.mockWriteBool(true)).to.be.eql('0x01');
    expect(await poaUintTest.mockWriteBool(false)).to.be.eql('0x80');
  });

  it('should be able to encode string value', async () => {
    let input = 'fdhafaskhfhakfsakfsddasda';
    expect(await poaUintTest.mockWriteString(input)).to.be.eql(encodeRlp(Uint8Array.from(Buffer.from(input))));
  });

  it('should be able to encode bytes value', async () => {
    expect(await poaUintTest.mockWriteBytes('0x')).to.be.eql('0x80');
    let input = '0xabbcdabbcdabbbbcdabbc7381461841341342342d613131231';
    expect(await poaUintTest.mockWriteBytes(input)).to.be.eql(encodeRlp(input));
  });

  it('should be able to encode list value', async () => {
    /// encodeRlp with original value
    /// RLPWriter.sol writes list with encoded value
    const input = [
      '0xffffffffffffffffffffffffffffffffffffffff',
      Uint8Array.from(Buffer.from('fdhafaskhfhakfsakfsddasda')),
      Uint8Array.from(Buffer.from('cat')),
      '0xabbcdabbcdabbbbcdabbc7381461841341342342d613131231'
    ];
    const inputAfterEncoding = input.map((value) => {
      return encodeRlp(value);
    });

    expect(await poaUintTest.mockWriteRLPList(inputAfterEncoding)).to.be.eql(encodeRlp(input));
  });

  it("should be able to recover signer's address", async () => {
    let messageHash = solidityPackedKeccak256(
      [
        'uint256',
        'bytes',
        'uint64',
        'uint64',
        'bytes32',
        'bytes',
        'address',
        'bytes32',
        'bytes8',
        'uint256',
        'bytes32',
        'bytes32',
        'bytes32',
        'uint256',
        'bytes32',
        'uint64',
        'uint256',
        'bytes[]',
        'bytes32',
        'bytes[]'
      ],
      [
        '0x2',
        '0xd983010000846765746889676f312e31322e3137856c696e7578000000000000',
        '0x1c9c380',
        '0x0',
        '0xc3fa2927a8e5b7cfbd575188a30c34994d3356607deb4c10d7fefe0dd5cdcc83',
        '0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000',
        '0x35552c16704d214347f29fa77f77da6d75d7c752',
        '0x0000000000000000000000000000000000000000000000000000000000000000',
        '0x0000000000000000',
        '0x68b3',
        '0xbf4d16769b8fd946394957049eef29ed938da92454762fc6ac65e0364ea004c7',
        '0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421',
        '0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347',
        '0x261',
        '0x7b5a72075082c31ec909afe5c5df032b6e7f19c686a9a408a2cb6b75dec072a3',
        '0x5f080818',
        '0xd167',
        [],
        '0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421',
        []
      ]
    );

//     0x5d1a549fea201ae4074d1bec89f3cd59ad1c2a6e03151f823b11cc4396bac3a1
// 0xff855cda572a31f045a14da70e5ef19b264de861c00eccb42354b38ad899e1c434053ae9b7139f13eaccd5d44124dda9fb1870efd7374f791bd3f79bbd0bcc3c1c
// 0x5e087f154444d71f493312afa4558e3ab12d2494
    let signature = await admin.signMessage(toBeArray(messageHash));
    let signerAddress = await poaUintTest.mockRecoverAddress("0x5d1a549fea201ae4074d1bec89f3cd59ad1c2a6e03151f823b11cc4396bac3a1", "0xff855cda572a31f045a14da70e5ef19b264de861c00eccb42354b38ad899e1c434053ae9b7139f13eaccd5d44124dda9fb1870efd7374f791bd3f79bbd0bcc3c1c");
    expect(signerAddress).to.be.eql(admin.address);
  });
});
