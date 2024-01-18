import * as dotenv from 'dotenv';
import { DeployFunction } from 'hardhat-deploy/types';
import { HardhatRuntimeEnvironment } from 'hardhat/types';

dotenv.config();

const deployFunc: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const system = await deploy('System', {
    from: deployer,
    log: true
  });
  await hre.run('verify:verify', { address: system.address });
};

deployFunc.tags = ['BSCSystem'];
deployFunc.dependencies = [];
export default deployFunc;
