import * as dotenv from 'dotenv';
import { DeployFunction } from 'hardhat-deploy/types';
import { HardhatRuntimeEnvironment } from 'hardhat/types';

dotenv.config();

const deployFunc: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const ed25519Verifier = await deploy('Ed25519Verifier', {
    from: deployer,
    log: true
  });
  await hre.run('verify:verify', { address: ed25519Verifier.address });
};

deployFunc.tags = ['BSCEd25519Verifier'];
deployFunc.dependencies = [];
export default deployFunc;
