import { Fixture } from 'ethereum-waffle';
import { BigNumber, Wallet } from 'ethers';
import { ethers, waffle } from 'hardhat';
import { Bn254Agg16Bn254Verifier } from '../../typechain';
import assert from 'assert';

async function deployContract(admin: Wallet) {
  const _factory = await ethers.getContractFactory('Bn254Agg16Bn254Verifier');
  const _contract = await _factory.connect(admin).deploy();
  return _contract;
}

describe('Aggregation 16 bn254 verifier', async () => {
  function loadFixture<T>(fixture: Fixture<T>): Promise<T> {
    const provider = waffle.provider;
    return waffle.createFixtureLoader(provider.getWallets(), provider)(fixture);
  }

  async function fixture([admin]: Wallet[]) {
    const contract = await deployContract(admin);
    return { admin, contract };
  }

  let contract: Bn254Agg16Bn254Verifier;
  let admin: Wallet;
  beforeEach(async () => {
    const res = await loadFixture(fixture);
    contract = res.contract;
    admin = res.admin;
  });

  it('should pass on true proof', async () => {
    const result = await contract.verifyProof(
      [
        BigNumber.from('11830850729108585455172258155696583116114732859436807323920536533716035746043'), // a0
        BigNumber.from('15254079597827780999204351645423160103922071787202866354566367542277164830544'), // a1

        BigNumber.from('16653330851126253150031607848367488855453137027043618954941348147516164828070'), // b00
        BigNumber.from('187395973282851142812818300055117104584449070294888472755189743255328718237'), // b01
        BigNumber.from('2561618507493324289757197175121106424307926489473486205103964338130515637287'), // b10
        BigNumber.from('17013017322422535754621388735464417826469970866779433642599626359462923223805'), // b11

        BigNumber.from('199341362335784188358407382230085009749434421050018979444983503522532056789'), // c0
        BigNumber.from('10462291537232667351245904513168077236341112261785544481009194723134034279261') // c1
      ],
      [
        BigNumber.from('8281029652804115540779242248394256858262374845244472384326955447842738625749'), // Commitment 0
        BigNumber.from('16291454333950787529071139148857612321350070812453737627848777377269130231635') // Commitment 1
      ],
      [
        BigNumber.from('9119409738732686618655070762577648050360522084447116602775900759116236399278'), // Commitment POK0
        BigNumber.from('20096295095246371279378278495715130754889878869658331075511746223579898895990') // Commitment POK1
      ],
      [
        BigNumber.from('114336509697133943113087290526769312348'),
        BigNumber.from('154082401081700815200070389198426915731'),
        BigNumber.from('134926299864838284709497123807691682185'),
        BigNumber.from('146665854302452622094621623662217484466')
      ]
    );
    assert.equal(result, true);
  });
});
