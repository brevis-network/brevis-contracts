import * as dotenv from 'dotenv';
import { DeployFunction } from 'hardhat-deploy/types';
import { HardhatRuntimeEnvironment } from 'hardhat/types';

dotenv.config();

const deployFunc: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const bscValidatorSet = await deploy('BSCValidatorSet', {
    from: deployer,
    log: true
  });
  await hre.run('verify:verify', { address: bscValidatorSet.address });
};

deployFunc.tags = ['BSCValidatorSet'];
deployFunc.dependencies = [];
export default deployFunc;
