import { expect } from 'chai';
import { ContractRunner } from 'ethers';
import { ethers } from 'hardhat';

import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers';

import { TransactionProofVerifier, VerifierGasReport } from '../../typechain';
import { splitHash } from '../util';

describe('Transaction proof verify', async () => {
  async function fixture() {
    const [admin] = await ethers.getSigners();
    const verifier = await deployLib(admin);
    return { verifier };
  }

  let verifier: VerifierGasReport;

  beforeEach(async () => {
    const res = await loadFixture(fixture);
    verifier = res.verifier;
  });

  async function deployLib(admin: ContractRunner) {
    const _factory = await ethers.getContractFactory('TransactionProofVerifier');
    const _contract = (await _factory.connect(admin).deploy()) as TransactionProofVerifier;
    const address = await _contract.getAddress();
    const factory = await ethers.getContractFactory('VerifierGasReport');
    const contract = (await factory.connect(admin).deploy(address)) as VerifierGasReport;
    return contract;
  }

  it('Verify transaction Proof', async () => {
    const publicInputs = [
      ...splitHash('958c0c028b7a8fe0a3d5961620582cad1f557604937a104f77246246118a24c7'),
      ...splitHash('88bd78528ea4fd5c232978ce51e43f41f0d76ce56e331147c1c9611282308799'),
      BigInt('17086605'),
      BigInt('1681980179')
    ];
    await expect(
      verifier.transactionVerifyProof(
        [
          BigInt('0x2512c06f6094b50e90709f0cbc3f0f455d2c2be86f4d8fe98f230a7f19d66796'),
          BigInt('0x16ee8249067ecd870819b6beae7255584f37d9e3eecee4b749d8101b2e6c07e7')
        ],
        [
          [
            BigInt('0x124ea5e2c0be872ba3209c8fe7c567825c62a6256ada212afa3ecc68b9df2f1f'),
            BigInt('0x161cd7137912abe46c714398421f2ea62797af0eee60d8cdc7296c211a044db5')
          ],
          [
            BigInt('0x11a04204fcfeef8ee2835bd903bd00d434f5820e4fd13de42b0a9853e1bcd337'),
            BigInt('0x041ffd173b9a720e54cd71b0ec58a03bd6067327d82da157ef699e236cbef18a')
          ]
        ],
        [
          BigInt('0x2ed8fb551f5d4facf8abd74ea009049afe861b5010c0f7dd40831b758aa76e6c'),
          BigInt('0x053b9e12ab6d8e115cb639c090534560d8bc348fa84f938c39f8279f6ca1c2d1')
        ],
        [BigInt('0x0'), BigInt('0x0')],
        publicInputs
      )
    )
      .to.emit(verifier, 'ProofVerified')
      .withArgs(true);
  });

  it('Verify transaction Proof failure', async () => {
    const publicInputs = [
      ...splitHash('958c0c028b7a8fe0a3d5961620582cad1f557604937a104f77246246118a24c7'),
      ...splitHash('88bd78528ea4fd5c232978ce51e43f41f0d76ce56e331147c1c9611282308799'),
      BigInt('0'),
      BigInt('1681980179')
    ];
    await expect(
      verifier.transactionVerifyProof(
        [
          BigInt('0x2512c06f6094b50e90709f0cbc3f0f455d2c2be86f4d8fe98f230a7f19d66796'),
          BigInt('0x16ee8249067ecd870819b6beae7255584f37d9e3eecee4b749d8101b2e6c07e7')
        ],
        [
          [
            BigInt('0x124ea5e2c0be872ba3209c8fe7c567825c62a6256ada212afa3ecc68b9df2f1f'),
            BigInt('0x161cd7137912abe46c714398421f2ea62797af0eee60d8cdc7296c211a044db5')
          ],
          [
            BigInt('0x11a04204fcfeef8ee2835bd903bd00d434f5820e4fd13de42b0a9853e1bcd337'),
            BigInt('0x041ffd173b9a720e54cd71b0ec58a03bd6067327d82da157ef699e236cbef18a')
          ]
        ],
        [
          BigInt('0x2ed8fb551f5d4facf8abd74ea009049afe861b5010c0f7dd40831b758aa76e6c'),
          BigInt('0x053b9e12ab6d8e115cb639c090534560d8bc348fa84f938c39f8279f6ca1c2d1')
        ],
        [BigInt('0x0'), BigInt('0x0')],
        publicInputs
      )
    )
      .to.emit(verifier, 'ProofVerified')
      .withArgs(false);
  });

  it('Verify transaction Proof with raw data', async () => {
    const leafHash = divideToTwoString('0x958c0c028b7a8fe0a3d5961620582cad1f557604937a104f77246246118a24c7');
    const blockHash = divideToTwoString('0x88bd78528ea4fd5c232978ce51e43f41f0d76ce56e331147c1c9611282308799');

    const publicInputs = [
      BigInt(leafHash[0]),
      BigInt(leafHash[1]),

      BigInt(blockHash[0]),
      BigInt(blockHash[1]),

      BigInt('17086605'),
      BigInt('1681980179')
    ];

    const a = [
      BigInt('0x2512c06f6094b50e90709f0cbc3f0f455d2c2be86f4d8fe98f230a7f19d66796'),
      BigInt('0x16ee8249067ecd870819b6beae7255584f37d9e3eecee4b749d8101b2e6c07e7')
    ];

    const b = [
      [
        BigInt('0x124ea5e2c0be872ba3209c8fe7c567825c62a6256ada212afa3ecc68b9df2f1f'),
        BigInt('0x161cd7137912abe46c714398421f2ea62797af0eee60d8cdc7296c211a044db5')
      ],
      [
        BigInt('0x11a04204fcfeef8ee2835bd903bd00d434f5820e4fd13de42b0a9853e1bcd337'),
        BigInt('0x041ffd173b9a720e54cd71b0ec58a03bd6067327d82da157ef699e236cbef18a')
      ]
    ];

    const c = [
      BigInt('0x2ed8fb551f5d4facf8abd74ea009049afe861b5010c0f7dd40831b758aa76e6c'),
      BigInt('0x053b9e12ab6d8e115cb639c090534560d8bc348fa84f938c39f8279f6ca1c2d1')
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

    await expect(verifier.verifyRaw(allDataHex)).to.emit(verifier, 'ProofVerified').withArgs(true);
  });

  it('Verify transaction Proof with raw data failure', async () => {
    const leafHash = divideToTwoString('0x958c0c028b7a8fe0a3d5961620582cad1f557604937a104f77246246118a24c7');
    const blockHash = divideToTwoString('0x88bd78528ea4fd5c232978ce51e43f41f0d76ce56e331147c1c9611282308799');

    const publicInputs = [
      BigInt(leafHash[0]),
      BigInt(leafHash[1]),

      BigInt(blockHash[0]),
      BigInt(blockHash[1]),

      BigInt('17086605'),
      BigInt('0')
    ];

    const a = [
      BigInt('0x2512c06f6094b50e90709f0cbc3f0f455d2c2be86f4d8fe98f230a7f19d66796'),
      BigInt('0x16ee8249067ecd870819b6beae7255584f37d9e3eecee4b749d8101b2e6c07e7')
    ];

    const b = [
      [
        BigInt('0x124ea5e2c0be872ba3209c8fe7c567825c62a6256ada212afa3ecc68b9df2f1f'),
        BigInt('0x161cd7137912abe46c714398421f2ea62797af0eee60d8cdc7296c211a044db5')
      ],
      [
        BigInt('0x11a04204fcfeef8ee2835bd903bd00d434f5820e4fd13de42b0a9853e1bcd337'),
        BigInt('0x041ffd173b9a720e54cd71b0ec58a03bd6067327d82da157ef699e236cbef18a')
      ]
    ];

    const c = [
      BigInt('0x2ed8fb551f5d4facf8abd74ea009049afe861b5010c0f7dd40831b758aa76e6c'),
      BigInt('0x053b9e12ab6d8e115cb639c090534560d8bc348fa84f938c39f8279f6ca1c2d1')
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

    await expect(verifier.verifyRaw(allDataHex)).to.emit(verifier, 'ProofVerified').withArgs(false);
  });
});

const divideToTwoString = (input: string) => {
  return [input.slice(0, 34), '0x' + input.slice(34, 66)];
};
