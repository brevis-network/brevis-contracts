import {deployments, ethers} from 'hardhat';
import {MintableERC20__factory} from '../typechain';

const run = async () => {
  const [signer] = await ethers.getSigners();

  const dep = await deployments.get('MintableERC20');
  const token =  MintableERC20__factory.connect(dep.address, signer);
  const minter = await deployments.get('PegBridge');
  const tx = await token.setMinter(minter.address);
  console.log(`setMinter(${minter.address}) tx: ${tx.hash}`);
  await tx.wait();
};

run();
