import * as dotenv from 'dotenv';
import {DeployFunction} from 'hardhat-deploy/types';
import {HardhatRuntimeEnvironment} from 'hardhat/types';

dotenv.config();

const deployFunc: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const args = ['USDC', 'USDC', 18, '0'];

  const testToken = await deploy('MintableERC20', {
    from: deployer,
    log: true,
    args: args
  });
  await hre.run('verify:verify', { address: testToken.address, constructorArguments: args });
};

deployFunc.tags = ['MintableERC20'];
export default deployFunc;
