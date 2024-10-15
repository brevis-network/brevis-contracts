import { expect } from 'chai';
import { Fixture } from 'ethereum-waffle';
import { ContractTransaction, Wallet } from 'ethers';
import { ethers, waffle } from 'hardhat';
import { ISMT, MockAnchorBlocks__factory, SMT, SMT__factory, TestSmtVerifier__factory } from '../../typechain';

const depth2EmptySmtRoot = '0x070e068f3c2d0058b210baac3410a1b74537b44a5511b189371e9c852db10416';

async function deployAnchorBlocks(admin: Wallet) {
  const factory = await ethers.getContractFactory('MockAnchorBlocks');
  const anchorBlocks = await factory.connect(admin).deploy();
  await anchorBlocks.update(121850623, '0x33a40e4d31779e49311267cedd746360b3ef5eacfd6a7078b135929d5d338b6d');
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
  newSmtRoot: '0x0290e108b334135857000cc4c6254c1d7c07f93eadd59fd35f99fb8048c468a9',
  endBlockHash: '0x33a40e4d31779e49311267cedd746360b3ef5eacfd6a7078b135929d5d338b6d',
  endBlockNum: 121850623,
  nextChunkMerkleRoot: '0x0000000000000000000000000000000000000000000000000000000000000000',
  circuitDigest: "0x00000",
  proof: [
    '5807806351967397868418351699921116242922085858185123616047887050820884996471',
    '19444718462238971532514390638804520037672715311516313809615819236604649104438',
    '3782987856556060089938036329403711264575640744803629636603617850705838812888',
    '3328570114941450486273589426774577317336602527731786139229554529007969305568',
    '9720922898704233363774880950782154961329560611208701799831036238186506382403',
    '446210035996153600590775953350007198581621313124343101976050433146250541743',
    '20348589262760567142723476092780813469593613957578203490570436373472273728927',
    '15910632642884846606452119826850699447433410284647360223246721491871857913590'
  ],
  commit: [
    '17746360320258038327388371538916709431167230043166066574575803531293043135763',
    '18770351197953504710324308428780586436521069948874367943613437808709054204023'
  ],

  knowledgeProof: [
    '2370704969841923008866362083028803213249039676633425612201867923993722869866',
    '7335149012371827295907989140515408787219815165349437088581548537624481255117'
  ]
};

const updateOld: ISMT.SmtUpdateStruct = {
  newSmtRoot: '0x074d7ca9a23757ee1699ebd69467aef8113cac39f9cc01c30aaf92ebb33bf95b',
  endBlockHash: '0x3cb6f635c0359c49293bf4617f0f77adeaa5a5cfb7c193a6c8fdefdce41a5563',
  circuitDigest: "0x00000",
  nextChunkMerkleRoot: '0x0dffe9d9691142bd5df5afcb30c12625198fc7cde364215b92622cce224ad330',
  proof: [
    '11531063991101025000926165418085717051737547065070261499961709449198962642048',
    '12854774381594583006319193887958445085958108573707659413722706864837760546368',
    '19795723019141840639395352318069482489773141357505690820978540023099684390048',
    '3577572054498314567078287276178984913619982548305278731774951944847252197424',
    '20717511832389950260873350048051001796915393244286515747651492833046936924323',
    '6788486040275977191617932568771588844685765705921367170800735471280741421893',
    '4859979885274277453468687450514983859911432565239796763468095881917190279324',
    '14378573002820259113713754339096849723512788079317595027047596863240997556924'
  ],
  commit: [
    '2896101488489045746819075657837107027981145574041180435045881559718068981359',
    '8606797402587866704170647611505268471492404700903528997438939126792104572080'
  ],

  knowledgeProof: [
    '286896185516968390807140493048685098762192864580524699218135871076986127320',
    '2723192933418240579592925214303414831767389526806512792489802796558761900380'
  ]
};

describe('SMT', async () => {
  function loadFixture<T>(fixture: Fixture<T>): Promise<T> {
    const provider = waffle.provider;
    return waffle.createFixtureLoader(provider.getWallets(), provider)(fixture);
  }

  async function fixture([admin]: Wallet[]) {
    const contract = await deployContracts(admin);
    await contract.addProvers([admin.address]);
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
    await expect(tx).to.emit(contract, 'SmtRootUpdated').withArgs(updateNew.newSmtRoot, updateNew.endBlockHash, 1);
    let valid = await contract.isSmtRootValid(1, updateNew.newSmtRoot);
    expect(valid).true;

    tx = contract.updateRoot(1, updateOld);
    await expect(tx).to.emit(contract, 'SmtRootUpdated').withArgs(updateOld.newSmtRoot, updateOld.endBlockHash, 1);
    valid = await contract.isSmtRootValid(1, updateOld.newSmtRoot);
    expect(valid).true;
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
