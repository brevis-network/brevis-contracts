import * as dotenv from 'dotenv';
import { DeployFunction } from 'hardhat-deploy/types';
import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { verify } from '../utils/utils';

dotenv.config();

const deployFunc: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const args = [
    process.env.TEST_TOKEN_NAME,
    process.env.TEST_TOKEN_SYMBOL,
    process.env.TEST_TOKEN_DECIMALS,
    process.env.TEST_TOKEN_INITIAL_SUPPLY,
  ];
  const deployment = await deploy('MintableERC20', {
    from: deployer,
    log: true,
    args: args
  });
  await verify(hre, deployment, args);
};

deployFunc.tags = ['TestToken'];
deployFunc.dependencies = [];
export default deployFunc;
