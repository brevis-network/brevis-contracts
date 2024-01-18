import { expect } from 'chai';
import { Fixture } from 'ethereum-waffle';
import { BigNumberish, Wallet } from 'ethers';
import { ethers, waffle } from 'hardhat';
import { EthChunkOf4Verifier__factory, VerifierGasReport } from '../../typechain';

async function deployContract(admin: Wallet) {
  const _factory = await ethers.getContractFactory<EthChunkOf4Verifier__factory>('EthChunkOf4Verifier');
  const _contract = await _factory.connect(admin).deploy();
  const factory = await ethers.getContractFactory('VerifierGasReport');
  const contract = (await factory.connect(admin).deploy(_contract.address)) as VerifierGasReport;
  return contract;
}
function getTestProof() {
  const chunkRoot = 'ddc08833ebef5364d5cdedc989770becc949caf6cbdc72bba0e842d442136c06';
  const prevHash = '0301010101010101010101010101010101010101010101010101010101010102';
  const endHash = '931cf1c54d8ab666b200f2f88748a6a6ec03d3795a3a453895df015b2013b96d';

  const input = [...splitHash(chunkRoot), ...splitHash(prevHash), ...splitHash(endHash), 1, 4];
  const a: [BigNumberish, BigNumberish] = [
    '16217230224774761590414973642073192485520807822804992513061256091576710039093',
    '2002416893759411048918869122785221574310433394721837052525007122310353508563'
  ];
  const b: [[BigNumberish, BigNumberish], [BigNumberish, BigNumberish]] = [
    [
      '10197059906511048117905894151806242600921773365823995195260445360897935348383',
      '17176880858333124367872442540151136901916472062827225492995249573427445734932'
    ],
    [
      '3508508084177307333211815272848724987300711832513791765513593482163273805687',
      '8946113151421411368233424685315381086080512060853567398740775323993107361049'
    ]
  ];
  const c: [BigNumberish, BigNumberish] = [
    '4227181897368161064461503007610136746613299883531807541785600858067518608614',
    '10696546849598832835618070876676453336593868194332417262844290881924036536318'
  ];
  const commit: [BigNumberish, BigNumberish] = ['0', '0'];
  return { a, b, c, commit, input };
}

describe('Eth chunk of 4 proof verifier', async () => {
  function loadFixture<T>(fixture: Fixture<T>): Promise<T> {
    const provider = waffle.provider;
    return waffle.createFixtureLoader(provider.getWallets(), provider)(fixture);
  }

  async function fixture([admin]: Wallet[]) {
    const contract = await deployContract(admin);
    return { admin, contract };
  }

  let contract: VerifierGasReport;
  let admin: Wallet;
  beforeEach(async () => {
    const res = await loadFixture(fixture);
    contract = res.contract;
    admin = res.admin;
  });

  it('should pass on true proof', async () => {
    const { a, b, c, commit, input } = getTestProof();
    await expect(contract.ethChunkOf4VerifyProof(a, b, c, commit, input))
      .to.emit(contract, 'ProofVerified')
      .withArgs(true);
  });
  it('should not pass on false proofs', async () => {
    const p = getTestProof();
    p.a[0] = '0';
    await expect(contract.ethChunkOf4VerifyProof(p.a, p.b, p.c, p.commit, p.input)).reverted;
  });
  it('should not pass on false pub input', async () => {
    const p = getTestProof();
    p.input[0] = '0';
    await expect(contract.ethChunkOf4VerifyProof(p.a, p.b, p.c, p.commit, p.input))
      .to.emit(contract, 'ProofVerified')
      .withArgs(false);
  });
});

function splitHash(h: string): BigNumberish[] {
  const a = '0x' + h.substring(0, h.length / 2);
  const b = '0x' + h.substring(h.length / 2, h.length);
  return [a, b];
}
