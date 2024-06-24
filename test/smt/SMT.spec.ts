import { expect } from 'chai';
import { Fixture } from 'ethereum-waffle';
import { ContractTransaction, Wallet } from 'ethers';
import { ethers, waffle } from 'hardhat';
import { ISMT, MockAnchorBlocks__factory, SMT, SMT__factory, TestSmtVerifier__factory } from '../../typechain';

const depth2EmptySmtRoot = '0x0fc13d55e980de280a92295f523bb56391ad967ba4179d4c90543a2307556547';

async function deployAnchorBlocks(admin: Wallet) {
  const factory = await ethers.getContractFactory('MockAnchorBlocks');
  const anchorBlocks = await factory.connect(admin).deploy();
  await anchorBlocks.update(121806079, '0xf31d2cfb3bef7c0da0d2afc5d27f6c7c48e6f121dcae0c690cd24f95374b518d');
  return anchorBlocks;
}

async function deployVerifier(admin: Wallet) {
  const factory = await ethers.getContractFactory('SMTUpdateCircuitProofOnOpVerifier');
  return factory.connect(admin).deploy();
}

async function deployContracts(admin: Wallet) {
  const anchorBlocks = await deployAnchorBlocks(admin);
  const verifier = await deployVerifier(admin);
  const factory = await ethers.getContractFactory('SMT');
  return factory.connect(admin).deploy([1], [anchorBlocks.address], [verifier.address], [depth2EmptySmtRoot]);
}

const updateNew: ISMT.SmtUpdateStruct = {
  newSmtRoot: '0x0b9ccc8f4dfa236c8d3652858fbc40abb3a72b7028708f79fe422f262974e4d6',
  endBlockHash: '0xf31d2cfb3bef7c0da0d2afc5d27f6c7c48e6f121dcae0c690cd24f95374b518d',
  endBlockNum: 121806079,
  nextChunkMerkleRoot: '0x0000000000000000000000000000000000000000000000000000000000000000',
  proof: [
    '21821169791355133933891889452648915920577751791973719959164077605207648843090',
    '19878181986039368017902562061275858671296544307198184030353425826968041955033',
    '14862322486398405117821446630687218326109074894171495906866175546135817667636',
    '15857961825936985071370067783636201104765895992556675858389121155282245950853',
    '4829943051097039659915651631804430938205193304759576777037918965577808149583',
    '21242044059207098630711902922172372206375292567813642853094242461039023135199',
    '11690721679323315209589844301131352051234784611052030082942577328913955523741',
    '5585007241503935787816904692650015279192142790213093823530666731459650657613',
  ],
  commit: [
    '739709195962611198447366149093933090408188983280665092228039241202696244444', 
    '21311509328172720220289997776041723223716651243076542895396920995868853638603',
  ],

  knowledgeProof: [
    '14423296154561051198099808304710346936473447645347848920161415886806080048125', 
    '19028207970779925965057681527721118789709893733898208249730954658864318649031',
  ],
};

const updateOld: ISMT.SmtUpdateStruct = {
  newSmtRoot: '0x04c83b2fdff05bf817a05c048729cb4d05556158c34eaae00876896bebbc27ac',
  endBlockHash: '0x97018e37ba9055b307c4e2721f5810e1198aac4c903865a41df10362dbec57c9',
  endBlockNum: 121680511,
  nextChunkMerkleRoot: '0x0000000000000000000000000000000000000000000000000000000000000000',
  proof: [
    '239817638144552681721916866575576872653226085684361068742504232711724822432',
    '13943460908745183759092205081867638806267643204664929307209808299533302607211',
    '3692190010637739787388349945922910587940245465963984672717761834747731500881',
    '4763759920509881571974987867100305459187871726036322828165748604554900784425',
    '21009580099009264720689974819204809225082117639356866216592064355635728743268',
    '6367024708557821601329572227511560630501908895341927720245242361531213375965',
    '7121626639567398836621650107912346764251371323818705382205456214430413651849',
    '172669705171825628136287955348324292558477888836225703899369327801960013121',
  ],
  commit: [
    '10244783296106980424375341033110534732398731246657464619371249519647297078754', 
    '21700570861839820137369934097716805653566525299639534689608947640894445143703',
  ],

  knowledgeProof: ['0', '0'],
};

describe('SMT', async () => {
  function loadFixture<T>(fixture: Fixture<T>): Promise<T> {
    const provider = waffle.provider;
    return waffle.createFixtureLoader(provider.getWallets(), provider)(fixture);
  }

  async function fixture([admin]: Wallet[]) {
    const contract = await deployContracts(admin);
    return { admin, contract };
  }

  let contract: SMT;
  let admin: Wallet;
  beforeEach(async () => {
    const res = await loadFixture(fixture);
    contract = res.contract;
    admin = res.admin;
  });

  it('update passes on true proofs', async () => {
    let tx: Promise<ContractTransaction>;
    tx = contract.updateRoot(1, updateNew);
    await expect(tx).to.emit(contract, 'SmtRootUpdated').withArgs(updateNew.newSmtRoot, updateNew.endBlockNum, 1);
    let valid = await contract.isSmtRootValid(1, updateNew.newSmtRoot);
    expect(valid).true;

    // tx = contract.updateRoot(1, updateOld);
    // await expect(tx).to.emit(contract, 'SmtRootUpdated').withArgs(updateOld.newSmtRoot, updateOld.endBlockNum, 2);
    // valid = await contract.isSmtRootValid(1, updateOld.newSmtRoot);
    // expect(valid).true;
  });
  it('update reverts on false proofs', async () => {
    let tx: Promise<ContractTransaction>;
    const tmp = updateNew.proof[0];
    updateNew.proof[0] = '0';
    tx = contract.updateRoot(1, updateNew);
    await expect(tx).reverted;

    updateNew.proof[0] = tmp;
    updateNew.newSmtRoot = '0x0000000000000000000000000000000000000000000000000000000000000000';
    tx = contract.updateRoot(1, updateNew);
    await expect(tx).reverted;
  });
});
