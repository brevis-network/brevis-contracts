import { ethers } from 'hardhat';
import { expect } from 'chai';
import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers';

import { BrevisBn254Verifier, BrevisBn254Verifier__factory } from '../../typechain';

async function deployContract() {
  const [admin] = await ethers.getSigners();
  const contract = await new BrevisBn254Verifier__factory().connect(admin).deploy();
  return { contract };
}

describe('BN254 verifier', async () => {
  let contract: BrevisBn254Verifier;
  beforeEach(async () => {
    const res = await loadFixture(deployContract);
    contract = res.contract;
  });

  it('should pass on true proof', async () => {
    const tx = contract.verifyProof(
      [
        BigInt('0x22519da121779e8eae3ca0dda17820f035ea04dc24e4cca303c039c44ea1dfaf'),
        BigInt('0x2231892ec88e5b80afcc128dda7f67f5670e5bbb04290900cf60f40bb9a84e7a'),
        BigInt('0x16970e8fe1824d2f7c6b6afcf2d93fc4157d72067d296d567626eb47470dce5d'),
        BigInt('0x1e4ee7bccdd202d1b767465709275dd2ca0bc1e48e8dd5897fb062efc2210a9b'),
        BigInt('0x265240a2d3bedf081146666bf1a4997bf361786ece39bdca65a3076eb6c8f239'),
        BigInt('0x101d51e5ba9ded3918cc35a1f0ba12cf5254b8cfbc117adfcde2dccc3bf8d947'),

        BigInt('0x177ef65a6f54c3ed5ef46587486c5863dfde672f5ce7a21030517e4fffbf525a'),
        BigInt('0x07bcfa46276489d9a55f7c75a2bf53645e1b4885dfff681d7a560cd09020775a')
      ],
      [
        BigInt('0x04c9b82e998bfecac61da6da562bfc5730688cea8437466e75b593ca20ec9e49'),
        BigInt('0x22447eb0af135719dc80ffa01e59636308fc7f6ac426a55923cfbfe3ef037f37')
      ],
      [
        BigInt('0x0a556c3eb619dbdfd2942e1d9e458577c2e8e52710156d8a2ac9985b02ed95e9'),
        BigInt('0x184685fbc8c917833963e9dae24bc3c8ac91650e57a9e936fd095c6b27d76e23')
      ],
      [
        BigInt('0x04b64a227c4e2ce2ba210d54ed3969fc2891de5d323ae3c5d6277adc11bdc109'),
        BigInt('0x08646494e28db9b84dbf61f1725e5174'),
        BigInt('0xf7eaee5265971f0d92f33559810e4428'),
        BigInt('0x0f9d2de6cc715cc878ded5252d0ec84d9440b3b60f4b82d8b477944ac1f48999'),
        BigInt('0x00000000000000000000000000000000'),
        BigInt('0x00000000000000000000000000000000'),
        BigInt('0x122e1d04b13e148af66f8fd91acb354519af92b46558fc2b4eb31f2b50cfcb60')
      ]
    );
    await expect(tx).not.reverted
  });

  it('should revert on false proof', async () => {
    const tx = contract.verifyProof(
      [
        BigInt('0'),
        BigInt('0x2231892ec88e5b80afcc128dda7f67f5670e5bbb04290900cf60f40bb9a84e7a'),
        BigInt('0x16970e8fe1824d2f7c6b6afcf2d93fc4157d72067d296d567626eb47470dce5d'),
        BigInt('0x1e4ee7bccdd202d1b767465709275dd2ca0bc1e48e8dd5897fb062efc2210a9b'),
        BigInt('0x265240a2d3bedf081146666bf1a4997bf361786ece39bdca65a3076eb6c8f239'),
        BigInt('0x101d51e5ba9ded3918cc35a1f0ba12cf5254b8cfbc117adfcde2dccc3bf8d947'),

        BigInt('0x177ef65a6f54c3ed5ef46587486c5863dfde672f5ce7a21030517e4fffbf525a'),
        BigInt('0x07bcfa46276489d9a55f7c75a2bf53645e1b4885dfff681d7a560cd09020775a')
      ],
      [
        BigInt('0x04c9b82e998bfecac61da6da562bfc5730688cea8437466e75b593ca20ec9e49'),
        BigInt('0x22447eb0af135719dc80ffa01e59636308fc7f6ac426a55923cfbfe3ef037f37')
      ],
      [
        BigInt('0x0a556c3eb619dbdfd2942e1d9e458577c2e8e52710156d8a2ac9985b02ed95e9'),
        BigInt('0x184685fbc8c917833963e9dae24bc3c8ac91650e57a9e936fd095c6b27d76e23')
      ],
      [
        BigInt('0x04b64a227c4e2ce2ba210d54ed3969fc2891de5d323ae3c5d6277adc11bdc109'),
        BigInt('0x08646494e28db9b84dbf61f1725e5174'),
        BigInt('0xf7eaee5265971f0d92f33559810e4428'),
        BigInt('0x0f9d2de6cc715cc878ded5252d0ec84d9440b3b60f4b82d8b477944ac1f48999'),
        BigInt('0x00000000000000000000000000000000'),
        BigInt('0x00000000000000000000000000000000'),
        BigInt('0x122e1d04b13e148af66f8fd91acb354519af92b46558fc2b4eb31f2b50cfcb60')
      ]
    );
    await expect(tx).reverted;
  });
});