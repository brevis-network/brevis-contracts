import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers';
import { EthStorageVerifier, VerifierGasReport } from '../../typechain';
import { ethers, } from 'hardhat';
import { ContractRunner, getBytes } from 'ethers';
import { convertByteArrayToHexString } from './util';
import { expect } from 'chai';
import { hexToBytes } from '../util';

describe('Eth storage proof verify', async () => {
  async function fixture() {
    const [admin] = await ethers.getSigners();
    const originalVerifier = await deploy(admin)
    const address =  await originalVerifier.getAddress()
    const verifier = await deployLib(admin, address);
    return { admin, originalVerifier, verifier };
  }

  let originalVerifier: EthStorageVerifier;
  let verifier: VerifierGasReport;
  let admin: ContractRunner;

  beforeEach(async () => {
    const res = await loadFixture(fixture);
    verifier = res.verifier;
    originalVerifier = res.originalVerifier;
    admin = res.admin;
  });

  async function deployLib(admin: ContractRunner, originalVerifierAddress: string) {
    const factory = await ethers.getContractFactory('VerifierGasReport');
    const contract = (await factory.connect(admin).deploy(originalVerifierAddress)) as VerifierGasReport;
    return contract;
  }

  async function deploy(admin: ContractRunner) {
    const etherStorageProofFactory = await ethers.getContractFactory('EthStorageVerifier');
    const etherStorageContract = (await etherStorageProofFactory.connect(admin).deploy()) as EthStorageVerifier;
    return etherStorageContract;
  }

  const publicInputs = [
    BigInt(
      convertByteArrayToHexString([103, 197, 210, 106, 230, 239, 0, 173, 207, 151, 13, 155, 24, 118, 240, 234])
    ),
    BigInt(
      convertByteArrayToHexString([236, 65, 249, 77, 136, 183, 160, 41, 158, 157, 97, 9, 205, 217, 188, 216])
    ),
    BigInt(
      convertByteArrayToHexString([230, 66, 26, 191, 243, 181, 187, 60, 128, 126, 39, 8, 155, 41, 116, 25])
    ),
    BigInt(
      convertByteArrayToHexString([251, 9, 216, 152, 169, 76, 141, 172, 214, 149, 130, 94, 141, 128, 60, 56])
    ),
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

  it('Verify Storage Proof', async () => {
    // blockNumber := int64(17037800)
    // accountAddress := "0x5427FEFA711Eff984124bFBB1AB6fbf5E3DA1820"
    // key := "0x03"

    const originalResult = await originalVerifier.verifyProof(
      [
        BigInt('0x1ff096821d084175748931e8c3bc749e8a6d2c133dcc174535e3af55739510b7'),
        BigInt('0x06eb7bebb8746f6343376fb12eb9f96f8e35c36daadf59a38790bb7c545fb6c1')
      ],
      [
        [
          BigInt('0x08617521d17a02aa3c8289b323eeb7e2698d64e12dc1016b06b14ac470b66cb9'),
          BigInt('0x1587e05e535e3b7292aae17f313a88a4bd4559068663086f5bc0e1efed9fe2bb')
        ],
        [
          BigInt('0x1d0b579ad36cff136336e75b4653bbed97618644d9ce170591e03219250b3933'),
          BigInt('0x03a160d716298b64b0f83d5850998276d1c89ee298d003633a4fd9145a605b5e')
        ]
      ],
      [
        BigInt('0x04517a59525f5a40f70fb80dbe445178513dea08165542deefece2ebc39f6e6a'),
        BigInt('0x031e24b1b11384502dd7e3bad636002b86eada9613b67505d09529d593ca897b')
      ],
      [BigInt('0x0'), BigInt('0x0')],
      publicInputs
    );

    expect(originalResult).equal(true);

    await expect(
      verifier.ethStorageVerifyProof(
        [
          BigInt('0x1ff096821d084175748931e8c3bc749e8a6d2c133dcc174535e3af55739510b7'),
          BigInt('0x06eb7bebb8746f6343376fb12eb9f96f8e35c36daadf59a38790bb7c545fb6c1')
        ],
        [
          [
            BigInt('0x08617521d17a02aa3c8289b323eeb7e2698d64e12dc1016b06b14ac470b66cb9'),
            BigInt('0x1587e05e535e3b7292aae17f313a88a4bd4559068663086f5bc0e1efed9fe2bb')
          ],
          [
            BigInt('0x1d0b579ad36cff136336e75b4653bbed97618644d9ce170591e03219250b3933'),
            BigInt('0x03a160d716298b64b0f83d5850998276d1c89ee298d003633a4fd9145a605b5e')
          ]
        ],
        [
          BigInt('0x04517a59525f5a40f70fb80dbe445178513dea08165542deefece2ebc39f6e6a'),
          BigInt('0x031e24b1b11384502dd7e3bad636002b86eada9613b67505d09529d593ca897b')
        ],
        [BigInt('0x0'), BigInt('0x0')],
        publicInputs
      )
    )
      .to.emit(verifier, 'ProofVerified')
      .withArgs(true);
  });

  it('Verify Storage Proof Raw', async () => {
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

    let allDataHex = '';
    for (let i = 0; i < allData.length; i++) {
      allDataHex = allDataHex + BigInt(allData[i]).toString(16).padStart(64, '0');
    }

    allDataHex = "0x" + allDataHex
    const originalResult = await originalVerifier.verifyRaw(allDataHex);
    expect(originalResult).equal(true);
    await expect(verifier.verifyRaw(allDataHex))
      .to.emit(verifier, 'ProofVerified')
      .withArgs(true);
  });

  it('Failed to Verify Storage Proof', async () => {
    // blockNumber := int64(17037800)
    // accountAddress := "0x5427FEFA711Eff984124bFBB1AB6fbf5E3DA1820"
    // key := "0x03"
    publicInputs[0] = BigInt(0);
    await expect(
      verifier.ethStorageVerifyProof(
        [
          BigInt('0x1ff096821d084175748931e8c3bc749e8a6d2c133dcc174535e3af55739510b7'),
          BigInt('0x06eb7bebb8746f6343376fb12eb9f96f8e35c36daadf59a38790bb7c545fb6c1')
        ],
        [
          [
            BigInt('0x08617521d17a02aa3c8289b323eeb7e2698d64e12dc1016b06b14ac470b66cb9'),
            BigInt('0x1587e05e535e3b7292aae17f313a88a4bd4559068663086f5bc0e1efed9fe2bb')
          ],
          [
            BigInt('0x1d0b579ad36cff136336e75b4653bbed97618644d9ce170591e03219250b3933'),
            BigInt('0x03a160d716298b64b0f83d5850998276d1c89ee298d003633a4fd9145a605b5e')
          ]
        ],
        [
          BigInt('0x04517a59525f5a40f70fb80dbe445178513dea08165542deefece2ebc39f6e6a'),
          BigInt('0x031e24b1b11384502dd7e3bad636002b86eada9613b67505d09529d593ca897b')
        ],
        [BigInt('0x0'), BigInt('0x0')],
        publicInputs
      )
    )
      .to.emit(verifier, 'ProofVerified')
      .withArgs(false);
  });
});
