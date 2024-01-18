import { expect } from 'chai';
import { Fixture } from 'ethereum-waffle';
import { BigNumberish, Wallet } from 'ethers';
import { ethers, waffle } from 'hardhat';
import { EthChunkOf128Verifier__factory, VerifierGasReport } from '../../typechain';

async function deployContract(admin: Wallet) {
  const _factory = await ethers.getContractFactory<EthChunkOf128Verifier__factory>('EthChunkOf128Verifier');
  const _contract = await _factory.connect(admin).deploy();
  const factory = await ethers.getContractFactory('VerifierGasReport');
  const contract = (await factory.connect(admin).deploy(_contract.address)) as VerifierGasReport;
  return contract;
}

describe('Eth chunk of 128 proof verifier', async () => {
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

  function getTestProof() {
    // blockCount := 128
    // startBlockNum := 17285000
    // endBlockNum := startBlockNum + blockCount - 1

    const chunkRoot = '1a15768d060e7bc9cfba3c7ed62724adca2b200ef384fb5d07dde6397587b73a';
    const prevHash = 'c11f13fedeb6ed6f7bcad39f96a7280edd6cb8d995484cd3638e6e61cac029fe';
    const endHash = '67e33240456ea4ea73508bbfb3fa25051a07faf9b114f0bade7247c94bc2777e';

    const input = [...splitHash(chunkRoot), ...splitHash(prevHash), ...splitHash(endHash), 17285000, 17285127];
    const a: [BigNumberish, BigNumberish] = [
      '12882591104901069077805259544002660246928880385684811525870849398445345546341',
      '2215135116940610125973457879245914942662591600154515839482387395722498787479'
    ];
    const b: [[BigNumberish, BigNumberish], [BigNumberish, BigNumberish]] = [
      [
        '7317008878237725070147013043790366054022270314661131268810747390122398315507',
        '11988587309652545500925661599800601362600393460992293505292265629324453009166'
      ],
      [
        '13409617315319580831274284109473186124666224920707247084446096933858498455915',
        '3854582558130440260491499516996393500761294334356771985926744062150340489134'
      ]
    ];
    const c: [BigNumberish, BigNumberish] = [
      '19590589426928906396866518501129776522694229017130944452427566408300888632556',
      '1756334569785290077017846840215164680231445088529104707833614834470122996583'
    ];
    const commit: [BigNumberish, BigNumberish] = ['0', '0'];

    return { a, b, c, commit, input };
  }

  it('should pass on true proof', async () => {
    const { a, b, c, commit, input } = getTestProof();
    await expect(contract.ethChunkOf128VerifyProof(a, b, c, commit, input))
      .to.emit(contract, 'ProofVerified')
      .withArgs(true);
  });
  it('should not pass on false proofs', async () => {
    const p = getTestProof();
    p.a[0] = '0';
    await expect(contract.ethChunkOf128VerifyProof(p.a, p.b, p.c, p.commit, p.input)).reverted;
  });
  it('should not pass on false pub input', async () => {
    const p = getTestProof();
    p.input[0] = '0';
    await expect(contract.ethChunkOf128VerifyProof(p.a, p.b, p.c, p.commit, p.input))
      .to.emit(contract, 'ProofVerified')
      .withArgs(false);
  });
});

function splitHash(h: string): BigNumberish[] {
  const a = '0x' + h.substring(0, h.length / 2);
  const b = '0x' + h.substring(h.length / 2, h.length);
  return [a, b];
}
