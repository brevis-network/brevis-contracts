import { Fixture } from 'ethereum-waffle';
import { ReceiptCircuitProofVerifier, VerifierGasReport } from '../../typechain';
import { ethers, waffle } from 'hardhat';
import { BigNumber, Wallet } from 'ethers';
import { expect } from 'chai';
import { splitHash, hexToBytes } from '../util';

describe('Receipt circuit proof verification', async () => {
  function loadFixture<T>(fixture: Fixture<T>): Promise<T> {
    const provider = waffle.provider;
    return waffle.createFixtureLoader(provider.getWallets(), provider)(fixture);
  }

  async function fixture([admin]: Wallet[]) {
    const originalVerifier = await deployOriginalVerifier(admin);
    const gasReporter = await deployGasReporter(admin, originalVerifier.address);
    return { admin, originalVerifier, gasReporter };
  }

  let gasReporter: VerifierGasReport;

  beforeEach(async () => {
    const res = await loadFixture(fixture);
    gasReporter = res.gasReporter;
  });

  async function deployOriginalVerifier(admin: Wallet) {
    const factory = await ethers.getContractFactory('ReceiptCircuitProofVerifier');
    const contract = (await factory.connect(admin).deploy()) as ReceiptCircuitProofVerifier;
    return contract;
  }

  async function deployGasReporter(admin: Wallet, originalVerifierAddress: string) {
    const factory = await ethers.getContractFactory('VerifierGasReport');
    const contract = (await factory.connect(admin).deploy(originalVerifierAddress)) as VerifierGasReport;
    return contract;
  }

  it('Verify receipt circuit Proof', async () => {
    const publicInputs = [
      ...splitHash('ec3384944ee3756aba922025ae1805096022e11d0abbb25be199fc918e4e7765'),
      ...splitHash('a3f5f903ac37f86fa7ff562892d94aa31e65dda2a2a356efe693fef0e35ec313'),
      BigNumber.from('17490377'),
      BigNumber.from('1686893999')
    ];

    await expect(
      gasReporter.receiptVerifyProof(
        [
          BigNumber.from('0x2c518faac1930a047fe3f409647301be1b8c717c05e7be57c7364eeec64e0afe'),
          BigNumber.from('0x161833dca478bc0dbf7fd208364caa3b67708116e7c5544e7ac30761d2ac892a')
        ],
        [
          [
            BigNumber.from('0x2ce8e956869c5ab9557aafd1766c96fae726e4b8790c2d8fa2ccdd8beba33080'),
            BigNumber.from('0x06d9ea8ee23d1df99329e5e1920b3b26e85926056d9903bff9ee08814fd5c486')
          ],
          [
            BigNumber.from('0x17dfef6c5c408b5e4af92c19e5f160430e16acf350fe22aa4af3772278a00b6f'),
            BigNumber.from('0x235f8dc87868313c5e0c5e9b5953155938dd0efc371f8b08af8a15592eda73f1')
          ]
        ],
        [
          BigNumber.from('0x1e0bd8018b4f177864db0f618fa26d4f19f2a97f392e916ce7b29feaef569969'),
          BigNumber.from('0x24ab8b849877e48500297733e25819fada1f2c45c5e7933620383ad302468c13')
        ],
        [BigNumber.from('0x0'), BigNumber.from('0x0')],
        publicInputs
      )
    )
      .to.emit(gasReporter, 'ProofVerified')
      .withArgs(true);
  });

  it('Verify receipt Proof failure', async () => {
    const publicInputs = [
      ...splitHash('ec3384944ee3756aba922025ae1805096022e11d0abbb25be199fc918e4e7765'),
      ...splitHash('a3f5f903ac37f86fa7ff562892d94aa31e65dda2a2a356efe693fef0e35ec313'),
      BigNumber.from('17490377'),
      BigNumber.from('0') /// Change this for mock failure
    ];

    await expect(
      gasReporter.receiptVerifyProof(
        [
          BigNumber.from('0x2c518faac1930a047fe3f409647301be1b8c717c05e7be57c7364eeec64e0afe'),
          BigNumber.from('0x161833dca478bc0dbf7fd208364caa3b67708116e7c5544e7ac30761d2ac892a')
        ],
        [
          [
            BigNumber.from('0x2ce8e956869c5ab9557aafd1766c96fae726e4b8790c2d8fa2ccdd8beba33080'),
            BigNumber.from('0x06d9ea8ee23d1df99329e5e1920b3b26e85926056d9903bff9ee08814fd5c486')
          ],
          [
            BigNumber.from('0x17dfef6c5c408b5e4af92c19e5f160430e16acf350fe22aa4af3772278a00b6f'),
            BigNumber.from('0x235f8dc87868313c5e0c5e9b5953155938dd0efc371f8b08af8a15592eda73f1')
          ]
        ],
        [
          BigNumber.from('0x1e0bd8018b4f177864db0f618fa26d4f19f2a97f392e916ce7b29feaef569969'),
          BigNumber.from('0x24ab8b849877e48500297733e25819fada1f2c45c5e7933620383ad302468c13')
        ],
        [BigNumber.from('0x0'), BigNumber.from('0x0')],
        publicInputs
      )
    )
      .to.emit(gasReporter, 'ProofVerified')
      .withArgs(false);
  });

  it('Verify receipt Proof with raw data', async () => {
    const leafHash = divideToTwoString('0xec3384944ee3756aba922025ae1805096022e11d0abbb25be199fc918e4e7765');
    const blockHash = divideToTwoString('0xa3f5f903ac37f86fa7ff562892d94aa31e65dda2a2a356efe693fef0e35ec313');

    const publicInputs = [
      BigNumber.from(leafHash[0]),
      BigNumber.from(leafHash[1]),

      BigNumber.from(blockHash[0]),
      BigNumber.from(blockHash[1]),

      BigNumber.from('17490377'),
      BigNumber.from('1686893999')
    ];

    const a = [
      BigNumber.from('0x2c518faac1930a047fe3f409647301be1b8c717c05e7be57c7364eeec64e0afe'),
      BigNumber.from('0x161833dca478bc0dbf7fd208364caa3b67708116e7c5544e7ac30761d2ac892a')
    ];

    const b = [
      [
        BigNumber.from('0x2ce8e956869c5ab9557aafd1766c96fae726e4b8790c2d8fa2ccdd8beba33080'),
        BigNumber.from('0x06d9ea8ee23d1df99329e5e1920b3b26e85926056d9903bff9ee08814fd5c486')
      ],
      [
        BigNumber.from('0x17dfef6c5c408b5e4af92c19e5f160430e16acf350fe22aa4af3772278a00b6f'),
        BigNumber.from('0x235f8dc87868313c5e0c5e9b5953155938dd0efc371f8b08af8a15592eda73f1')
      ]
    ];

    const c = [
      BigNumber.from('0x1e0bd8018b4f177864db0f618fa26d4f19f2a97f392e916ce7b29feaef569969'),
      BigNumber.from('0x24ab8b849877e48500297733e25819fada1f2c45c5e7933620383ad302468c13')
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

    await expect(gasReporter.verifyRaw(hexToBytes(allDataHex)))
      .to.emit(gasReporter, 'ProofVerified')
      .withArgs(true);
  });

  it('Verify receipt proof with raw data failure', async () => {
    const leafHash = divideToTwoString('0xec3384944ee3756aba922025ae1805096022e11d0abbb25be199fc918e4e7765');
    const blockHash = divideToTwoString('0xa3f5f903ac37f86fa7ff562892d94aa31e65dda2a2a356efe693fef0e35ec313');

    const publicInputs = [
      BigNumber.from(leafHash[0]),
      BigNumber.from(leafHash[1]),

      BigNumber.from(blockHash[0]),
      BigNumber.from(blockHash[1]),

      BigNumber.from('17490377'),
      BigNumber.from('0')
    ];

    const a = [
      BigNumber.from('0x2c518faac1930a047fe3f409647301be1b8c717c05e7be57c7364eeec64e0afe'),
      BigNumber.from('0x161833dca478bc0dbf7fd208364caa3b67708116e7c5544e7ac30761d2ac892a')
    ];

    const b = [
      [
        BigNumber.from('0x2ce8e956869c5ab9557aafd1766c96fae726e4b8790c2d8fa2ccdd8beba33080'),
        BigNumber.from('0x06d9ea8ee23d1df99329e5e1920b3b26e85926056d9903bff9ee08814fd5c486')
      ],
      [
        BigNumber.from('0x17dfef6c5c408b5e4af92c19e5f160430e16acf350fe22aa4af3772278a00b6f'),
        BigNumber.from('0x235f8dc87868313c5e0c5e9b5953155938dd0efc371f8b08af8a15592eda73f1')
      ]
    ];

    const c = [
      BigNumber.from('0x1e0bd8018b4f177864db0f618fa26d4f19f2a97f392e916ce7b29feaef569969'),
      BigNumber.from('0x24ab8b849877e48500297733e25819fada1f2c45c5e7933620383ad302468c13')
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

    await expect(gasReporter.verifyRaw(hexToBytes(allDataHex)))
      .to.emit(gasReporter, 'ProofVerified')
      .withArgs(false);
  });
});

const divideToTwoString = (input: string) => {
  return [input.slice(0, 34), '0x' + input.slice(34, 66)];
};
