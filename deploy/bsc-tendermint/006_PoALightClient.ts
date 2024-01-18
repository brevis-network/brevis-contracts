import * as dotenv from 'dotenv';
import { DeployFunction } from 'hardhat-deploy/types';
import { HardhatRuntimeEnvironment } from 'hardhat/types';

dotenv.config();

const deployFunc: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const bscValidatorSet = process.env.BSC_VALIDATOR_SET;
  const args = [bscValidatorSet];

  const poaLightClient = await deploy('PoALightClient', {
    from: deployer,
    log: true,
    args: args
  });
  await hre.run('verify:verify', { address: poaLightClient.address, constructorArguments: args });
};

deployFunc.tags = ['BSCPoALightClient'];
deployFunc.dependencies = [];
export default deployFunc;
