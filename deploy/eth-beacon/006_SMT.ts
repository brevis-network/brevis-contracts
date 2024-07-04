import * as dotenv from 'dotenv';
import { DeployFunction } from 'hardhat-deploy/types';
import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { verify } from '../utils/utils';

dotenv.config();

const defaultSmtRoot = '0x1138d0ac4aedc38fa8d5f64859de3cd2837a6e5267203da1b7ff0fbd85e1faec';

const deployFunc: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  // const anchorProvider = await deployments.get('AnchorBlocks');
  const anchorProvider = await deployments.get('SameChainAnchorBlocks');
  const verifier = await deployments.get('SMTUpdateCircuitProofOnOpVerifier');
  const args = [[10], [anchorProvider.address], [verifier.address], [defaultSmtRoot]];
  const deployment = await deploy('SMT', { from: deployer, log: true, args: args });
  await verify(hre, deployment);
};

deployFunc.tags = ['SMT'];
deployFunc.dependencies = [];
export default deployFunc;
