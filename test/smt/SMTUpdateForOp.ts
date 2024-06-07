import { expect } from 'chai';
import { Fixture } from 'ethereum-waffle';
import { BigNumber, BigNumberish, Wallet } from 'ethers';
import { ethers, waffle } from 'hardhat';
import { SMTUpdateCircuitProofOnOpVerifier__factory, SMTUpdateCircuitProofOnOpVerifier } from '../../typechain';

// cmd: npx hardhat test SMTUpdateForOp.ts

async function deployContract(admin: Wallet) {
  const _factory = await ethers.getContractFactory('SMTUpdateCircuitProofOnOpVerifier');
  const _contract = await _factory.connect(admin).deploy();
  return _contract;
}

describe('on chain smt proof verifier, for op chain', async () => {
  function loadFixture<T>(fixture: Fixture<T>): Promise<T> {
    const provider = waffle.provider;
    return waffle.createFixtureLoader(provider.getWallets(), provider)(fixture);
  }

  async function fixture([admin]: Wallet[]) {
    const contract = await deployContract(admin);
    return { admin, contract };
  }

  let contract: SMTUpdateCircuitProofOnOpVerifier;
  let admin: Wallet;
  beforeEach(async () => {
    const res = await loadFixture(fixture);
    contract = res.contract;
    admin = res.admin;
  });

    it('should pass on true proof brevis', async () => {
        const result = await contract.verifyProof(
            [
                BigNumber.from('8216257027028561216502580608860422514212221953842283736672084816773768675847'),
                BigNumber.from('8854285040091176228077914789046506744806603769217785934385781983806081799635'),

                BigNumber.from('21588053076846511636110462216897615959541748832630334292311545751926473725374'),
                BigNumber.from('15399944883833730308619413961960068140402256607949963022272192032630279651142'),
                BigNumber.from('3479616346013860026977833050524063169232552761586418166485098669335751079307'),
                BigNumber.from('14271108628743724592489264284112727172640486726988321734057632445942732568739'),

                BigNumber.from('16631136640609464239804327671067508194135591745597904444436072386144947807613'),
                BigNumber.from('11699260662303439682552929464620458827943327494098726692494478117446457905127')
            ],
            [
                BigNumber.from('21725911230617623619121704871686023049144524661412504705466236207886617871909'),
                BigNumber.from('8913661905551737052849830116650167801370055047143974869736295452201120957426')
            ],
            [
                BigNumber.from('9374538635832710502593755645510863256385512155220141255319371830534366747844'),
                BigNumber.from('17468642930970240000076840587782953920071268956745139549011248520438235431264'),
            ],
            [
                BigNumber.from('19460084310786029646694965716373939580'),
                BigNumber.from('257094756869936700027883792101059331767'),
                BigNumber.from('1963998565058433517617748592832641758'),
                BigNumber.from('260366310900955347496903632087524868466'),
                BigNumber.from('26593845891943262020741027148124405967'),
                BigNumber.from('17953760944934965145003606174584918399'),
                BigNumber.from('120732799'),
                BigNumber.from('0'),
                BigNumber.from('0'),
            ]
        );
        console.log('result4', result);
    });
});
