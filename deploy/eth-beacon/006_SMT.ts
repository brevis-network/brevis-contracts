import * as dotenv from 'dotenv';
import { DeployFunction } from 'hardhat-deploy/types';
import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { verify } from '../utils/utils';

dotenv.config();

const defaultSmtRoot = '0x8be7b2f3bb10dbfbacf4f3b2da07bbacab804cd4b58addd5ab609520a65dbd0b';

const deployFunc: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  // const anchorProvider = await deployments.get('AnchorBlocks');
  const anchorProvider = await deployments.get('MockAnchorBlocks');
  const verifier = await deployments.get('BrevisPlonky2SmtVerifier');
  const args = [[1], [anchorProvider.address], [verifier.address], [defaultSmtRoot]];
  const deployment = await deploy('SMT', { from: deployer, log: true, args: args });
  await verify(hre, deployment);
};

deployFunc.tags = ['SMT'];
deployFunc.dependencies = [];
export default deployFunc;
