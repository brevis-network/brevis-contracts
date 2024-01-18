import * as dotenv from 'dotenv';
import { DeployFunction } from 'hardhat-deploy/types';
import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { verify } from '../utils/utils';

dotenv.config();

const deployFunc: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const deployment = await deploy('TransactionProofVerifier', {
    from: deployer,
    log: true
  });
  await verify(hre, deployment);
};

deployFunc.tags = ['TransactionProofVerifier'];
deployFunc.dependencies = [];
export default deployFunc;
