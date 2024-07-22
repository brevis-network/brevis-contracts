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
        BigNumber.from('21821169791355133933891889452648915920577751791973719959164077605207648843090'),
        BigNumber.from('19878181986039368017902562061275858671296544307198184030353425826968041955033'),

        BigNumber.from('14862322486398405117821446630687218326109074894171495906866175546135817667636'),
        BigNumber.from('15857961825936985071370067783636201104765895992556675858389121155282245950853'),
        BigNumber.from('4829943051097039659915651631804430938205193304759576777037918965577808149583'),
        BigNumber.from('21242044059207098630711902922172372206375292567813642853094242461039023135199'),

        BigNumber.from('11690721679323315209589844301131352051234784611052030082942577328913955523741'),
        BigNumber.from('5585007241503935787816904692650015279192142790213093823530666731459650657613')
      ],
      [
        BigNumber.from('739709195962611198447366149093933090408188983280665092228039241202696244444'),
        BigNumber.from('21311509328172720220289997776041723223716651243076542895396920995868853638603')
      ],
      [
        BigNumber.from('14423296154561051198099808304710346936473447645347848920161415886806080048125'),
        BigNumber.from('19028207970779925965057681527721118789709893733898208249730954658864318649031')
      ],
      [
        BigNumber.from('0x0fc13d55e980de280a92295f523bb563'),
        BigNumber.from('0x91ad967ba4179d4c90543a2307556547'),
        BigNumber.from('0x0b9ccc8f4dfa236c8d3652858fbc40ab'),
        BigNumber.from('0xb3a72b7028708f79fe422f262974e4d6'),
        BigNumber.from('0xf31d2cfb3bef7c0da0d2afc5d27f6c7c'),
        BigNumber.from('0x48e6f121dcae0c690cd24f95374b518d'),
        BigNumber.from('121806079'),
        BigNumber.from('0'),
        BigNumber.from('0')
      ]
    );
    console.log('result4', result);
  });
});
