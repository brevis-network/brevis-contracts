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
  const sigsVerifier = await deployments.get('BrevisProof');
  const args = ["0xe03B4133fFf5d4023ec125c11167E1ba8d7419DE", proof.address, sigsVerifier.address];
  const deployment = await deploy('BrevisRequest', {
    from: deployer,
    log: true,
    args: args
  });
  await verify(hre, deployment, args);
};

deployFunc.tags = ['BrevisRequest'];
deployFunc.dependencies = [];
export default deployFunc;
