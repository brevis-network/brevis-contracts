import * as dotenv from 'dotenv';
import { DeployFunction } from 'hardhat-deploy/types';
import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { verify } from '../utils/utils';

dotenv.config();

const defaultSmtRoot = '0xca3022d33d006b913ecfad08eb9d6dde985e93fdd80ca6fadebcd2172d28c3f3';

const deployFunc: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const deployment = await deploy('Plonky2ProofVerifier', { from: deployer, log: true });
  await verify(hre, deployment);
};

deployFunc.tags = ['Plonky2ProofVerifier'];
deployFunc.dependencies = [];
export default deployFunc;
