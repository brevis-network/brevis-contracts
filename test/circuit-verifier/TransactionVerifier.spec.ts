import { Fixture } from 'ethereum-waffle';
import { TransactionProofVerifier, VerifierGasReport } from '../../typechain';
import { ethers, waffle } from 'hardhat';
import { BigNumber, Wallet } from 'ethers';
import { expect } from 'chai';
import { splitHash, hexToBytes } from '../util';

describe('Transaction proof verify', async () => {
  function loadFixture<T>(fixture: Fixture<T>): Promise<T> {
    const provider = waffle.provider;
    return waffle.createFixtureLoader(provider.getWallets(), provider)(fixture);
  }

  async function fixture([admin]: Wallet[]) {
    const verifier = await deployLib(admin);
    return { admin, verifier };
  }

  let verifier: VerifierGasReport;
  let admin: Wallet;

  beforeEach(async () => {
    const res = await loadFixture(fixture);
    verifier = res.verifier;
    admin = res.admin;
  });

  async function deployLib(admin: Wallet) {
    const _factory = await ethers.getContractFactory('TransactionProofVerifier');
    const _contract = (await _factory.connect(admin).deploy()) as TransactionProofVerifier;
    const factory = await ethers.getContractFactory('VerifierGasReport');
    const contract = (await factory.connect(admin).deploy(_contract.address)) as VerifierGasReport;
    return contract;
  }

  it('Verify transaction Proof', async () => {
    const publicInputs = [
      ...splitHash('958c0c028b7a8fe0a3d5961620582cad1f557604937a104f77246246118a24c7'),
      ...splitHash('88bd78528ea4fd5c232978ce51e43f41f0d76ce56e331147c1c9611282308799'),
      BigNumber.from('17086605'),
      BigNumber.from('1681980179')
    ];
    await expect(
      verifier.transactionVerifyProof(
        [
          BigNumber.from('0x2512c06f6094b50e90709f0cbc3f0f455d2c2be86f4d8fe98f230a7f19d66796'),
          BigNumber.from('0x16ee8249067ecd870819b6beae7255584f37d9e3eecee4b749d8101b2e6c07e7')
        ],
        [
          [
            BigNumber.from('0x124ea5e2c0be872ba3209c8fe7c567825c62a6256ada212afa3ecc68b9df2f1f'),
            BigNumber.from('0x161cd7137912abe46c714398421f2ea62797af0eee60d8cdc7296c211a044db5')
          ],
          [
            BigNumber.from('0x11a04204fcfeef8ee2835bd903bd00d434f5820e4fd13de42b0a9853e1bcd337'),
            BigNumber.from('0x041ffd173b9a720e54cd71b0ec58a03bd6067327d82da157ef699e236cbef18a')
          ]
        ],
        [
          BigNumber.from('0x2ed8fb551f5d4facf8abd74ea009049afe861b5010c0f7dd40831b758aa76e6c'),
          BigNumber.from('0x053b9e12ab6d8e115cb639c090534560d8bc348fa84f938c39f8279f6ca1c2d1')
        ],
        [BigNumber.from('0x0'), BigNumber.from('0x0')],
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
      BigNumber.from('0'),
      BigNumber.from('1681980179')
    ];
    await expect(
      verifier.transactionVerifyProof(
        [
          BigNumber.from('0x2512c06f6094b50e90709f0cbc3f0f455d2c2be86f4d8fe98f230a7f19d66796'),
          BigNumber.from('0x16ee8249067ecd870819b6beae7255584f37d9e3eecee4b749d8101b2e6c07e7')
        ],
        [
          [
            BigNumber.from('0x124ea5e2c0be872ba3209c8fe7c567825c62a6256ada212afa3ecc68b9df2f1f'),
            BigNumber.from('0x161cd7137912abe46c714398421f2ea62797af0eee60d8cdc7296c211a044db5')
          ],
          [
            BigNumber.from('0x11a04204fcfeef8ee2835bd903bd00d434f5820e4fd13de42b0a9853e1bcd337'),
            BigNumber.from('0x041ffd173b9a720e54cd71b0ec58a03bd6067327d82da157ef699e236cbef18a')
          ]
        ],
        [
          BigNumber.from('0x2ed8fb551f5d4facf8abd74ea009049afe861b5010c0f7dd40831b758aa76e6c'),
          BigNumber.from('0x053b9e12ab6d8e115cb639c090534560d8bc348fa84f938c39f8279f6ca1c2d1')
        ],
        [BigNumber.from('0x0'), BigNumber.from('0x0')],
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
      BigNumber.from(leafHash[0]),
      BigNumber.from(leafHash[1]),

      BigNumber.from(blockHash[0]),
      BigNumber.from(blockHash[1]),

      BigNumber.from('17086605'),
      BigNumber.from('1681980179')
    ];

    const a = [
      BigNumber.from('0x2512c06f6094b50e90709f0cbc3f0f455d2c2be86f4d8fe98f230a7f19d66796'),
      BigNumber.from('0x16ee8249067ecd870819b6beae7255584f37d9e3eecee4b749d8101b2e6c07e7')
    ];

    const b = [
      [
        BigNumber.from('0x124ea5e2c0be872ba3209c8fe7c567825c62a6256ada212afa3ecc68b9df2f1f'),
        BigNumber.from('0x161cd7137912abe46c714398421f2ea62797af0eee60d8cdc7296c211a044db5')
      ],
      [
        BigNumber.from('0x11a04204fcfeef8ee2835bd903bd00d434f5820e4fd13de42b0a9853e1bcd337'),
        BigNumber.from('0x041ffd173b9a720e54cd71b0ec58a03bd6067327d82da157ef699e236cbef18a')
      ]
    ];

    const c = [
      BigNumber.from('0x2ed8fb551f5d4facf8abd74ea009049afe861b5010c0f7dd40831b758aa76e6c'),
      BigNumber.from('0x053b9e12ab6d8e115cb639c090534560d8bc348fa84f938c39f8279f6ca1c2d1')
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

    await expect(verifier.verifyRaw(hexToBytes(allDataHex)))
      .to.emit(verifier, 'ProofVerified')
      .withArgs(true);
  });

  it('Verify transaction Proof with raw data failure', async () => {
    const leafHash = divideToTwoString('0x958c0c028b7a8fe0a3d5961620582cad1f557604937a104f77246246118a24c7');
    const blockHash = divideToTwoString('0x88bd78528ea4fd5c232978ce51e43f41f0d76ce56e331147c1c9611282308799');

    const publicInputs = [
      BigNumber.from(leafHash[0]),
      BigNumber.from(leafHash[1]),

      BigNumber.from(blockHash[0]),
      BigNumber.from(blockHash[1]),

      BigNumber.from('17086605'),
      BigNumber.from('0')
    ];

    const a = [
      BigNumber.from('0x2512c06f6094b50e90709f0cbc3f0f455d2c2be86f4d8fe98f230a7f19d66796'),
      BigNumber.from('0x16ee8249067ecd870819b6beae7255584f37d9e3eecee4b749d8101b2e6c07e7')
    ];

    const b = [
      [
        BigNumber.from('0x124ea5e2c0be872ba3209c8fe7c567825c62a6256ada212afa3ecc68b9df2f1f'),
        BigNumber.from('0x161cd7137912abe46c714398421f2ea62797af0eee60d8cdc7296c211a044db5')
      ],
      [
        BigNumber.from('0x11a04204fcfeef8ee2835bd903bd00d434f5820e4fd13de42b0a9853e1bcd337'),
        BigNumber.from('0x041ffd173b9a720e54cd71b0ec58a03bd6067327d82da157ef699e236cbef18a')
      ]
    ];

    const c = [
      BigNumber.from('0x2ed8fb551f5d4facf8abd74ea009049afe861b5010c0f7dd40831b758aa76e6c'),
      BigNumber.from('0x053b9e12ab6d8e115cb639c090534560d8bc348fa84f938c39f8279f6ca1c2d1')
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

    await expect(verifier.verifyRaw(hexToBytes(allDataHex)))
      .to.emit(verifier, 'ProofVerified')
      .withArgs(false);
  });
});

const divideToTwoString = (input: string) => {
  return [input.slice(0, 34), '0x' + input.slice(34, 66)];
};
