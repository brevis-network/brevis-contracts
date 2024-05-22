import * as dotenv from 'dotenv';
import { DeployFunction } from 'hardhat-deploy/types';
import { HardhatRuntimeEnvironment } from 'hardhat/types';

dotenv.config();

const deployFunc: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  await deploy('BVN', {
    from: deployer,
    log: true,
    args: ['0xC17f92d12ae0De25C9F89a5c7bcD6aC63AfeF782']
  });
};

deployFunc.tags = ['BVN'];
deployFunc.dependencies = [];
export default deployFunc;
