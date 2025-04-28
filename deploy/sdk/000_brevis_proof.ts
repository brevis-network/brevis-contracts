import * as dotenv from 'dotenv';
import { DeployFunction } from 'hardhat-deploy/types';
import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { verify } from '../utils/utils';

dotenv.config();

const deployFunc: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  // const smt = await deployments.get('SMT');
  const args = ['0x7636e517daf7eB3C7bf74BB982E9E0e55554e0D1'];
  const deployment = await deploy('BrevisProof', {
    from: deployer,
    log: true,
    args: args
  });
  await verify(hre, deployment, args);
};

deployFunc.tags = ['BrevisProof'];
deployFunc.dependencies = [];
export default deployFunc;
