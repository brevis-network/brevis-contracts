import { expect } from 'chai';
import { Fixture } from 'ethereum-waffle';
import { BigNumber, BigNumberish, Wallet } from 'ethers';
import { ethers, waffle } from 'hardhat';
import { OnChainBn254Verifier__factory, OnChainBn254Verifier } from '../../typechain';

// cmd: npx hardhat test BN254TestNewVerifier_v2.spec.ts

async function deployContract(admin: Wallet) {
  const _factory = await ethers.getContractFactory('OnChainBn254Verifier');
  const _contract = await _factory.connect(admin).deploy();
  return _contract;
}

describe('on chain BN254 new proof verifier', async () => {
  function loadFixture<T>(fixture: Fixture<T>): Promise<T> {
    const provider = waffle.provider;
    return waffle.createFixtureLoader(provider.getWallets(), provider)(fixture);
  }

  async function fixture([admin]: Wallet[]) {
    const contract = await deployContract(admin);
    return { admin, contract };
  }

  let contract: OnChainBn254Verifier;
  let admin: Wallet;
  beforeEach(async () => {
    const res = await loadFixture(fixture);
    contract = res.contract;
    admin = res.admin;
  });

  it('should pass on true proof', async () => {
    const result = await contract.verifyProofWithCommit(
      [
        BigNumber.from('3505869154799487274154051604654523367127013872211077091443895260083316795207'),
        BigNumber.from('2805644029761297169904529793920707726243698880576351859729035269841904808874'),

        BigNumber.from('16521177833142659863001123880535820738623215657184256828438741498072680831353'),
        BigNumber.from('14267077936896938193564222549455155955887373380261667467733385795185555907376'),
        BigNumber.from('5432281320634779813389792125626487044546670774364462261978502067260157529986'),
        BigNumber.from('21234743648584160058390584746734595875166819528060243891302587384381026765836'),

        BigNumber.from('21417637665821155503556325882148149683395693510299584638801772375626144804614'),
        BigNumber.from('16431307499276306440397413239470758753669703580340525861795048872578054826680')
      ],
      [
        BigNumber.from('5039644548309958745715662588982734324900835737516956278773683562757983235741'),
        BigNumber.from('2140778688527627774636786249530059499519005355064052246681841432388774061328')
      ],
      [
        BigNumber.from('15'),
        BigNumber.from('19550185910341262643857283842989960554652042465157786208272582640339747673526')
      ],
    );
    console.log('result', result);
  });

    it('should pass on true proof2', async () => {
        const result = await contract.verifyCommitmentCommitted(
            [
                BigNumber.from('5039644548309958745715662588982734324900835737516956278773683562757983235741'),
                BigNumber.from('2140778688527627774636786249530059499519005355064052246681841432388774061328'),
            ],
            [
                BigNumber.from('8701162704901251142275075132932648865428909294607467747956804274074166165655'),
                BigNumber.from('14511416323448129028502569305320958408972136914447242784996559978787213430794'),
            ],
        );
        console.log('result2', result);
    });

    it('should pass on true proof4', async () => {
        const result = await contract.verifyProofWithCommitAll(
            [
                BigNumber.from('3505869154799487274154051604654523367127013872211077091443895260083316795207'),
                BigNumber.from('2805644029761297169904529793920707726243698880576351859729035269841904808874'),

                BigNumber.from('16521177833142659863001123880535820738623215657184256828438741498072680831353'),
                BigNumber.from('14267077936896938193564222549455155955887373380261667467733385795185555907376'),
                BigNumber.from('5432281320634779813389792125626487044546670774364462261978502067260157529986'),
                BigNumber.from('21234743648584160058390584746734595875166819528060243891302587384381026765836'),

                BigNumber.from('21417637665821155503556325882148149683395693510299584638801772375626144804614'),
                BigNumber.from('16431307499276306440397413239470758753669703580340525861795048872578054826680')
            ],
            [
                BigNumber.from('5039644548309958745715662588982734324900835737516956278773683562757983235741'),
                BigNumber.from('2140778688527627774636786249530059499519005355064052246681841432388774061328')
            ],
            [
                BigNumber.from('15'),
                BigNumber.from('19550185910341262643857283842989960554652042465157786208272582640339747673526')
            ],
            [
                BigNumber.from('8701162704901251142275075132932648865428909294607467747956804274074166165655'),
                BigNumber.from('14511416323448129028502569305320958408972136914447242784996559978787213430794'),
            ]
        );
        console.log('result4', result);
    });
});
