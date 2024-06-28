import { expect } from 'chai';
import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers';
import { BigNumberish, ContractRunner, getBytes } from 'ethers';
import { ethers } from 'hardhat';
import {
  AggregationVerifier,
  AggregationVerifier__factory
} from '../../typechain';
import assert from 'assert';
import { hexToBytes } from '../util';

async function deployContract(admin: ContractRunner) {
  const factory = new AggregationVerifier__factory();
  const contract = await factory.connect(admin).deploy();
  return contract;
}

describe('AggregationVerifier', async () => {
  async function fixture() {
    const [admin] = await ethers.getSigners();
    const contract = await deployContract(admin);
    return { admin, contract };
  }

  let contract: AggregationVerifier;
  beforeEach(async () => {
    const res = await loadFixture(fixture);
    contract = res.contract;
  });

  it('should pass on true proof', async () => {
    const result = await contract.verifyProof(
      [
        BigInt('0x069747d573a88d6f27691e2ef1e3aec3c7041f08feaf205a5cb007b62f01f1ce'), // a0
        BigInt('0x249ff6b9a033725ea45da13548240c4ce184c08b86a29ec4a4fae5a0b712f9b0'), // a1

        BigInt('0x234f5b22623222e86ea2479e4ca0f56c2ed492d7eac01694c95ad15f2edc6037'), // b00
        BigInt('0x1530b264da66ad69ff874f65a0af9fda1c99fc3caff354561ba26e58f8a4955a'), // b01
        BigInt('0x020341aec37c6d583b305ea6f9a92882ba99ad9ec683730e47468e198af7b610'), // b10
        BigInt('0x012b06554a5cf087e7aa16b2ef92371480d93a80c1411fe154774083d458d30f'), // b11

        BigInt('0x1e1cc9170e733d9f6c47da67acb48d7756d8f84574e54da2beb76fae79e7706e'), // c0
        BigInt('0x2dd2486c966668c789082d7a195963a0cc3f89f206d8d822f77d7600e4ef8ab0'), // c1
      ],
      [
        BigInt('0x21b2f0c034ea4d93f471a6add8dbad4877856271fe8e7bcf2282b7d399cc0635'), // Commitment 0
        BigInt('0x22dfb97d1ce87049bd43f3aa5c2c2247d6e69c9afabbb903d8755ac85f45f79f'), // Commitment 1
      ],
      [
        BigInt('0x10860b6d156136db85afcc5615b34474dc448f2cbf74dd162552f06f1b687eb4'), // Commitment POK0
        BigInt('0x1a7ca928cc8e20afa2464649f5d62c3b07acb8c3e63ba4f18bb6d2e1d394d675'), // Commitment POK1
      ],
      [
        BigInt('0x21245b6b0756614544af32ea81d9dd81923d13c48afd3aabe1e81aabbd5a93c7'), // Query Hash
        BigInt('0x1166a98a66ffadcc0d211e59f2dadc37'), // SMT Root 0
        BigInt('0x0c38ef8683b3908515283d98227c9c64'), // SMT Root 1
        BigInt('0x22935546fb8b3a680600fcabfe361e5c0a16a99390415abe8aab491838dffca7'), // Aggregation VK Hash
        BigInt('0x16f3086a7b81b13af4d01a78533c686c'), // App Circuit Output Commitment 0
        BigInt('0xcb0c329caafd7509e12d666d0ae1ab69'), // App Circuit Output Commitment 1
        BigInt('0x1b3738642fbaef19b5b7f1d6e516905af845deb4215458037f76a5d435aee13e'), // App Circuit VK Hash
      ]
    );
    assert.equal(result, true);
  });

  it('should pass on verify raw with real data', async () => {
    const values = [
      '0x069747d573a88d6f27691e2ef1e3aec3c7041f08feaf205a5cb007b62f01f1ce', // a0
      '0x249ff6b9a033725ea45da13548240c4ce184c08b86a29ec4a4fae5a0b712f9b0', // a1

      '0x234f5b22623222e86ea2479e4ca0f56c2ed492d7eac01694c95ad15f2edc6037', // b00
      '0x1530b264da66ad69ff874f65a0af9fda1c99fc3caff354561ba26e58f8a4955a', // b01
      '0x020341aec37c6d583b305ea6f9a92882ba99ad9ec683730e47468e198af7b610', // b10
      '0x012b06554a5cf087e7aa16b2ef92371480d93a80c1411fe154774083d458d30f', // b11

      '0x1e1cc9170e733d9f6c47da67acb48d7756d8f84574e54da2beb76fae79e7706e', // c0
      '0x2dd2486c966668c789082d7a195963a0cc3f89f206d8d822f77d7600e4ef8ab0', // c1

      '0x21b2f0c034ea4d93f471a6add8dbad4877856271fe8e7bcf2282b7d399cc0635', // Commitment 0
      '0x22dfb97d1ce87049bd43f3aa5c2c2247d6e69c9afabbb903d8755ac85f45f79f', // Commitment 1

      '0x10860b6d156136db85afcc5615b34474dc448f2cbf74dd162552f06f1b687eb4', // Commitment POK0
      '0x1a7ca928cc8e20afa2464649f5d62c3b07acb8c3e63ba4f18bb6d2e1d394d675', // Commitment POK1

      '0x21245b6b0756614544af32ea81d9dd81923d13c48afd3aabe1e81aabbd5a93c7', // Query Hash
      '0x1166a98a66ffadcc0d211e59f2dadc370c38ef8683b3908515283d98227c9c64', // SMT Root 
      '0x22935546fb8b3a680600fcabfe361e5c0a16a99390415abe8aab491838dffca7', // Aggregation VK Hash
      '0x16f3086a7b81b13af4d01a78533c686ccb0c329caafd7509e12d666d0ae1ab69', // App Circuit Output Commitment 
      '0x1b3738642fbaef19b5b7f1d6e516905af845deb4215458037f76a5d435aee13e', //  App Circuit VK Hash
    ];

    var hexValues = '0x';

    values.forEach((value) => {
      hexValues += value.toString().slice(2).padStart(64, '0');
    });

    const result = await contract.verifyRaw(getBytes(hexValues));

    assert.equal(result, true);
  });
});
