import * as dotenv from 'dotenv';
import { DeployFunction } from 'hardhat-deploy/types';
import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { verify } from '../utils/utils';

dotenv.config();

const defaultSmtRoot = '0x0b12da9a8bff6f32444fafe4f657d51d5cb72b0aa707f3aff1481ff64d6ed805';

const deployFunc: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  // const anchorProvider = await deployments.get('AnchorBlocks');
  const anchorProvider = await deployments.get('SameChainAnchorBlocks');
  const verifier = await deployments.get('SMTUpdateCircuitProofVerifier');
  const args = [[42161], [anchorProvider.address], [verifier.address], [defaultSmtRoot]];
  const deployment = await deploy('TestSMT', { from: deployer, log: true, args: args });
  await verify(hre, deployment);
};

deployFunc.tags = ['TestSMT'];
deployFunc.dependencies = ['SameChainAnchorBlocks', 'SmtVerifier'];
export default deployFunc;
