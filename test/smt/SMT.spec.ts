import { expect } from 'chai';
import { Fixture } from 'ethereum-waffle';
import { ContractTransaction, Wallet } from 'ethers';
import { ethers, waffle } from 'hardhat';
import { ISMT, MockAnchorBlocks__factory, SMT, SMT__factory, TestSmtVerifier__factory } from '../../typechain';

const depth2EmptySmtRoot = '0x079088add0fe0111ea10854a338f56a0c95bbef74af24cb880a834136e4e719f';

async function deployAnchorBlocks(admin: Wallet) {
  const factory = await ethers.getContractFactory('MockAnchorBlocks');
  const anchorBlocks = await factory.connect(admin).deploy();
  await anchorBlocks.update(3, '0xfe2cc7043ae4f2b1648e9e88ba21fc909e701e2cff3148ca71f77459278b0876');
  return anchorBlocks;
}

async function deployVerifier(admin: Wallet) {
  const factory = await ethers.getContractFactory('TestSmtVerifier');
  return factory.connect(admin).deploy();
}

async function deployContracts(admin: Wallet) {
  const anchorBlocks = await deployAnchorBlocks(admin);
  const verifier = await deployVerifier(admin);
  const factory = await ethers.getContractFactory('SMT');
  return factory.connect(admin).deploy([1], [anchorBlocks.address], [verifier.address], [depth2EmptySmtRoot]);
}

const updateNew: ISMT.SmtUpdateStruct = {
  newSmtRoot: '0x01f425b428440bed0d29b1cf0ca9220dcdef0ff15e7ab61fd3833fa60c1f2bde',
  endBlockHash: '0xfe2cc7043ae4f2b1648e9e88ba21fc909e701e2cff3148ca71f77459278b0876',
  endBlockNum: 3,
  nextChunkMerkleRoot: '0x0000000000000000000000000000000000000000000000000000000000000000',
  commitPub: '0x23011508a19a36cc972594549eff4d7406f343793512b71e81cdb2cef95008ba',
  proof: {
    a: [
      '1348325766103565444276701016369451933705891687920345169994759433957300802996',
      '13943460908745183759092205081867638806267643204664929307209808299533302607211'
    ],
    b: [
      [
        '3217423846190919386986672941482542071075422235166187877942925759861129761667',
        '4675866305523457611370531063081053480410072291558564302283879934995574726205'
      ],
      [
        '1214388506424066661482902927759066765798777208334241665985193165794518531427',
        '21075191256196888779616464471319297766754208445417590240391609451786405270399'
      ]
    ],
    c: [
      '2479249280338148560603144652506588563575823593059417439884362637839756591843',
      '4831904693067925608058255871260751542431464648967497660241794507935679781514'
    ],
    commitment: [
      '12453155494853432631634824824178736470571814963087487601514605997283287232234',
      '2197942387692741143914033121376638768005645968997132824111194584405436652350'
    ]
  }
};

const updateOld: ISMT.SmtUpdateStruct = {
  newSmtRoot: '0x07b6cbd08df3f48178a40fa81103de10c494f439d83d1ef7ad7c7c8993941e94',
  endBlockHash: '0xc0238a9ff29ca00cac6814b8b89a354ae4bbdde687d49aade9e71cefa48d509d',
  endBlockNum: 1,
  nextChunkMerkleRoot: '0x04f47a02321dab1378f10c950c5f341d355b9bab7c7c53105e157a8c0c422182',
  commitPub: '0x21e254fabaf8481a21c7e4c71b456e0b3b2b379f6009f589281ee224024ce63e',
  proof: {
    a: [
      '3784961427403605596159122738352772403418869948069183603507934760709947621497',
      '7264194677928719525161587507676273350766590954292441115963750839876804734838'
    ],
    b: [
      [
        '12709619191839562308143724433351351937366997122848810617875333791953321239646',
        '17286146002646853864021043160559164032207725069887973441967408944186067846748'
      ],
      [
        '10273896804875088108792721921889084376801689215548945694063911514104802640664',
        '5696199705846166327142130606867804367803789088201111043121206532206630199860'
      ]
    ],
    c: [
      '17723507174724319324676143841521374564442500862971063278998726694622946646034',
      '1048752224077563743237833960664273455425256660945919114613062892668825390222'
    ],
    commitment: [
      '17274298116617061391457416488185453853350724120481981712654459293655824940174',
      '1312116012315038367768070603732790912446481217429689348574779417503974828838'
    ]
  }
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
    await expect(tx).to.emit(contract, 'SmtRootUpdated').withArgs(updateNew.newSmtRoot, updateNew.endBlockNum);
    let valid = await contract.isSmtRootValid(1, updateNew.newSmtRoot);
    expect(valid).true;

    tx = contract.updateRoot(1, updateOld);
    await expect(tx).to.emit(contract, 'SmtRootUpdated').withArgs(updateOld.newSmtRoot, updateOld.endBlockNum);
    valid = await contract.isSmtRootValid(1, updateOld.newSmtRoot);
    expect(valid).true;
  });
  it('update reverts on false proofs', async () => {
    let tx: Promise<ContractTransaction>;
    const tmp = updateNew.proof.a[0];
    updateNew.proof.a[0] = '0';
    tx = contract.updateRoot(1, updateNew);
    await expect(tx).reverted;

    updateNew.proof.a[0] = tmp;
    updateNew.newSmtRoot = '0x0000000000000000000000000000000000000000000000000000000000000000';
    tx = contract.updateRoot(1, updateNew);
    await expect(tx).reverted;
  });
});
