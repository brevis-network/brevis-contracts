import * as dotenv from 'dotenv';
import hre, { deployments, ethers } from 'hardhat';
import { MessageBridge, MessageBridge__factory } from '../typechain';

dotenv.config();

const params = [
  {
    chainId: 5,
    msgbr: '0x3d3F0b7Ab3efBFF17F2c38d610dF50BD62E5B322',
    lightClient: ''
  },
  {
    chainId: 97,
    msgbr: '0xcB521Fb80c951826A76D3fd7316314cbBa325985',
    lightClient: '0xE654Dd7D7fc226B17DDC587a3aF3B6709C4c5d5A'
  },
  {
    chainId: 43113,
    msgbr: '0x35623AeaaAc916Cd807Fc09613c270DA50534A04',
    lightClient: '0xd3A51e827E3bf96e83e2692356EAF93797447091'
  }
];

const setRemoteMessageBridge = async () => {
  const chainId = parseInt(await hre.getChainId(), 10);
  const [signer] = await ethers.getSigners();
  const dep = await deployments.get('MessageBridge');
  const msgbr = MessageBridge__factory.connect(dep.address, signer);

  for (const param of params.filter((p) => p.chainId !== chainId)) {
    const tx = await msgbr.setRemoteMessageBridge(param.chainId, param.msgbr);
    console.log(`setRemoteMessageBridge(${param.chainId}, ${param.msgbr}) tx:`, tx.hash);
    await tx.wait();
  }
};

const setLightClient = async () => {
  const chainId = parseInt(await hre.getChainId(), 10);
  const [signer] = await ethers.getSigners();

  const dep = await deployments.get('MessageBridge');
  const msgbr = MessageBridge__factory.connect(dep.address, signer);
  
  const param = params.find((p) => p.chainId === chainId);
  if (!param) {
    console.log(`param for chain ${chainId} not found`);
    return;
  }
  if (!param.lightClient) {
    return;
  }
  const tx = await msgbr.setLightClient(5, param.lightClient);
  await tx.wait();
  console.log(`setLightClient(${5}, ${param.lightClient}) tx:`, tx.hash);
};

const run = async () => {
  await setRemoteMessageBridge();
  // await setLightClient();
};

run();
