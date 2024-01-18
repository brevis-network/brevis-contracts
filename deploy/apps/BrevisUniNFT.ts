import * as dotenv from 'dotenv';
import { DeployFunction } from 'hardhat-deploy/types';
import { HardhatRuntimeEnvironment } from 'hardhat/types';
import {verify} from '../utils/utils';

dotenv.config();

const deployFunc: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const uniVol = await deployments.get('UniswapVolume');
  const args = ['Uniswap Stone-Tier Trader', 'UniSTONE', uniVol.address]; // name, symbol, minter
  // const args = ['Uniswap Bronze-Tier Trader', 'UniBRONZE', uniVol.address];
  // const args = ['Uniswap Silver-Tier Trader', 'UniSILVER', uniVol.address];
  // const args = ['Uniswap Gold-Tier Trader', 'UniGOLD', uniVol.address];
  // const args = ['Uniswap Platinum-Tier Trader', 'UniPLAT', uniVol.address];
  // const args = ['Uniswap Diamond-Tier Trader', 'UniDIAMOND', uniVol.address];
  const deployment = await deploy('BrevisUniNFT', {
    from: deployer,
    log: true,
    args: args
  });
  await verify(hre, deployment, args);
};

deployFunc.tags = ['BrevisUniNFT'];
deployFunc.dependencies = [];
export default deployFunc;
