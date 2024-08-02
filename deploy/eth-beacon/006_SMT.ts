import * as dotenv from 'dotenv';
import { DeployFunction } from 'hardhat-deploy/types';
import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { verify } from '../utils/utils';

dotenv.config();

const defaultSmtRoot = '0x0a7d854baa0e72c5b990c6e58ed2d759ccaae4642f16cac9bc94396c9571a6eb';

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
