import { expect } from 'chai';
import { Fixture } from 'ethereum-waffle';
import { BigNumber, BigNumberish, Wallet } from 'ethers';
import { ethers, waffle } from 'hardhat';
import { BN254NewVerifier__factory, BN254NewVerifier } from '../../typechain';

async function deployContract(admin: Wallet) {
    const _factory = await ethers.getContractFactory<BN254NewVerifier__factory>('BN254NewVerifier');
    const _contract = await _factory.connect(admin).deploy();
    return _contract;
}

describe('BN254 new proof verifier', async () => {
    function loadFixture<T>(fixture: Fixture<T>): Promise<T> {
        const provider = waffle.provider;
        return waffle.createFixtureLoader(provider.getWallets(), provider)(fixture);
    }

    async function fixture([admin]: Wallet[]) {
        const contract = await deployContract(admin);
        return { admin, contract };
    }

    let contract: BN254NewVerifier;
    let admin: Wallet;
    beforeEach(async () => {
        const res = await loadFixture(fixture);
        contract = res.contract;
        admin = res.admin;
    });

    it('should pass on true proof', async () => {
        const result = await contract.verifyProofWithCommit(
            [
                BigNumber.from('1808400746270349636267254679125757181848956720842592466715612977913086960483'),
                BigNumber.from('15355782967283791563637004912405592165010664300137859647696081913033681811743'),

                BigNumber.from('10454339805569045893655295550513972204170005418132707901249908737873710201572'),
                BigNumber.from('329934161221987699669304799660910125411407030454658304350014424980168156272'),
                BigNumber.from('17465039887431385941184736939969735538331955738619092234180450781178397707323'),
                BigNumber.from('8850514315872903516736605168709237907291678187703590138702594365127196181136'),

                BigNumber.from('8475091034611270602699735232057868525031422654690518324744635341779074631659'),
                BigNumber.from('15161941799735145042656893527596373848509963058156780749551713716948394776487'),
            ],
            [
                BigNumber.from('5421335610452967587599901835807331157677192447190481031355860202789490207366'),
                BigNumber.from('8284992782513961937242349918656600733662266424058595783524968317837430079877'),
            ],
            [
                BigNumber.from('5'),
                BigNumber.from('0x252a6c79712785a9b844884da30096756aaf491667f62398280d8cbf04cef111'),
            ],
        )
        console.log("result", result)
    });

});

function splitHash(h: string): BigNumberish[] {
    const a = '0x' + h.substring(0, h.length / 2);
    const b = '0x' + h.substring(h.length / 2, h.length);
    return [a, b];
}
