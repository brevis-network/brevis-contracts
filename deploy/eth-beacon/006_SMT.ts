import * as dotenv from 'dotenv';
import { DeployFunction } from 'hardhat-deploy/types';
import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { verify } from '../utils/utils';

dotenv.config();

const defaultSmtRoot = '0x754023de5458ebe80b59db8a8707d3317d7cbe19b8db8741fa06f531a69dc74a';

const deployFunc: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const anchorProvider = await deployments.get('AnchorBlocks');
  const verifier = await deployments.get('BrevisPlonky2SmtVerifier');
  const args = [[1], [anchorProvider.address], [verifier.address], [defaultSmtRoot]];
  const deployment = await deploy('SMT', { from: deployer, log: true, args: args });
  await verify(hre, deployment);
};

deployFunc.tags = ['SMT'];
deployFunc.dependencies = [];
export default deployFunc;
