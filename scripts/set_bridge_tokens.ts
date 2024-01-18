import * as dotenv from 'dotenv';
import hre, {deployments, ethers} from 'hardhat';
import { PegBridge__factory } from '../typechain';
import {PegBridge} from '../typechain/PegBridge';

dotenv.config();

const vaultTokens = ['0x7A5FaC2b0392B1A017Ac720739E2E352934C93a6'];
const bridgeTokens: { [addrs: number]: string[] } = {
  97: ['0x820a9347993763B0710a1ed9A3D40E257f57ebd2'],
  43113: ['0xaBF86F827164D10757fFEF396FEBf4c8631918d7']
};

const init = async () => {
  const [signer] = await ethers.getSigners();
  const chainId = parseInt(await hre.getChainId(), 10);
  
  const dep = await deployments.get('PegBridge');
  const bridge = PegBridge__factory.connect(dep.address, signer);

  const pegTokens = bridgeTokens[chainId];
  if (vaultTokens.length !== pegTokens?.length) {
    console.error('invalid tokens length');
    return;
  }
  const tx = await bridge.setBridgeTokens(vaultTokens, pegTokens);
  console.log(`setBridgeTokens(${vaultTokens}, ${pegTokens}) tx:`, tx.hash);
  await tx.wait();
};

init();
