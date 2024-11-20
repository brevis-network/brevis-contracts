import * as dotenv from 'dotenv';
import { DeployFunction } from 'hardhat-deploy/types';
import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { verify } from '../utils/utils';

dotenv.config();

const deployFunc: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  // const smt = await deployments.get('MockSMT');
  const args = ['0x6225F3302b3AaF62D588Af396f1f81FC9b974322'];
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
