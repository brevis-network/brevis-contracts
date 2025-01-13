import {deployments, ethers} from 'hardhat';
import {BrevisDummy__factory} from '../typechain';

const run = async () => {
  const [signer] = await ethers.getSigners();

  const db = await deployments.get('BrevisDummy');
  const brevisDummy =  BrevisDummy__factory.connect(db.address, signer);
  const tx = await brevisDummy.updateDummyEvent(1234567);
  console.log(`updateDummyEvent (${brevisDummy.address}) tx: ${tx.hash}`);
  await tx.wait();
};

run();
