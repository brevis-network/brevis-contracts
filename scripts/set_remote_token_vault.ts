import * as dotenv from 'dotenv';
import {deployments, ethers} from 'hardhat';
import { PegBridge__factory } from '../typechain';

dotenv.config();

const init = async () => {
  const [signer] = await ethers.getSigners();

  const dep = await deployments.get('PegBridge');
  const bridge = PegBridge__factory.connect(dep.address, signer);

  const tx = await bridge.setTokenVault(5, '0xc63b5873A3dA07A34f52224aEFA27a1f7FD7AB7A');
  await tx.wait();
  console.log('setRemoteTokenVault() tx:', tx.hash);
};

init();
