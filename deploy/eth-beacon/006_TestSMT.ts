import * as dotenv from 'dotenv';
import { DeployFunction } from 'hardhat-deploy/types';
import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { verify } from '../utils/utils';

dotenv.config();

const defaultSmtRoot = '0x106e045eb67ba32796174de4fe2848fb89127ba16c398b7d536284ba1ea2b5a9';

const deployFunc: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  // const anchorProvider = await deployments.get('AnchorBlocks');
  const anchorProvider = await deployments.get('SameChainAnchorBlocks');
  const verifier = await deployments.get('SMTUpdateCircuitProofVerifier');
  const args = [[8453], [anchorProvider.address], [verifier.address], [defaultSmtRoot]];
  const deployment = await deploy('TestSMT', { from: deployer, log: true, args: args });
  await verify(hre, deployment);
};

deployFunc.tags = ['TestSMT'];
deployFunc.dependencies = ['SameChainAnchorBlocks', 'SmtVerifier'];
export default deployFunc;
