import * as dotenv from 'dotenv';
import { DeployFunction } from 'hardhat-deploy/types';
import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { verify } from '../utils/utils';

dotenv.config();

const deployFunc: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const proof = await deployments.get('BrevisProof');
  const totalFeeApp = await deployments.get('TotalFeeApp');
  const args = [proof.address, totalFeeApp.address];
  const deployment = await deploy('SingleRewardApp', {
    from: deployer,
    log: true,
    args: args
  });
  await verify(hre, deployment, args);
};

deployFunc.tags = ['SingleRewardApp'];
deployFunc.dependencies = [];
export default deployFunc;
