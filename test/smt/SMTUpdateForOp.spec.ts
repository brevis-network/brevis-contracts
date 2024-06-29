import { expect } from 'chai';
import { ContractRunner } from 'ethers';
import { ethers } from 'hardhat';

import { loadFixture } from '@nomicfoundation/hardhat-toolbox/network-helpers';

import {
  SMTUpdateCircuitProofOnOpVerifier,
  SMTUpdateCircuitProofOnOpVerifier__factory,
} from '../../typechain';

// cmd: npx hardhat test SMTUpdateForOp.ts

async function deployContract(admin: ContractRunner) {
  const _factory = new SMTUpdateCircuitProofOnOpVerifier__factory();
  const _contract = await _factory.connect(admin).deploy();
  return _contract;
}

describe('on chain smt proof verifier, for op chain', async () => {
  async function fixture() {
    const [admin] = await ethers.getSigners();
    const contract = await deployContract(admin);
    return { contract };
  }

  let contract: SMTUpdateCircuitProofOnOpVerifier;
  beforeEach(async () => {
    const res = await loadFixture(fixture);
    contract = res.contract;
  });

  it('should pass on true proof brevis', async () => {
    const result = await contract.verifyProof(
      [
        BigInt('21821169791355133933891889452648915920577751791973719959164077605207648843090'),
        BigInt('19878181986039368017902562061275858671296544307198184030353425826968041955033'),

        BigInt('14862322486398405117821446630687218326109074894171495906866175546135817667636'),
        BigInt('15857961825936985071370067783636201104765895992556675858389121155282245950853'),
        BigInt('4829943051097039659915651631804430938205193304759576777037918965577808149583'),
        BigInt('21242044059207098630711902922172372206375292567813642853094242461039023135199'),

        BigInt('11690721679323315209589844301131352051234784611052030082942577328913955523741'),
        BigInt('5585007241503935787816904692650015279192142790213093823530666731459650657613')
      ],
      [
        BigInt('739709195962611198447366149093933090408188983280665092228039241202696244444'),
        BigInt('21311509328172720220289997776041723223716651243076542895396920995868853638603')
      ],
      [
        BigInt('14423296154561051198099808304710346936473447645347848920161415886806080048125'),
        BigInt('19028207970779925965057681527721118789709893733898208249730954658864318649031')
      ],
      [
        BigInt('0x0fc13d55e980de280a92295f523bb563'),
        BigInt('0x91ad967ba4179d4c90543a2307556547'),
        BigInt('0x0b9ccc8f4dfa236c8d3652858fbc40ab'),
        BigInt('0xb3a72b7028708f79fe422f262974e4d6'),
        BigInt('0xf31d2cfb3bef7c0da0d2afc5d27f6c7c'),
        BigInt('0x48e6f121dcae0c690cd24f95374b518d'),
        BigInt('121806079'),
        BigInt('0'),
        BigInt('0')
      ]
    );
    expect(result).eq(true);
  });
});
