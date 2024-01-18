import * as dotenv from 'dotenv';
import { DeployFunction } from 'hardhat-deploy/types';
import { HardhatRuntimeEnvironment } from 'hardhat/types';

dotenv.config();

const deployFunc: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const mockLightClient = await deploy('MockLightClient', {
    from: deployer,
    log: true
    // args: [process.env.ORIGINAL_TOKEN_VAULT_SIGS_VERIFIER]
  });
  await hre.run('verify:verify', { address: mockLightClient.address });

  const mockMessageBridge = await deploy('MockMessageBridge', {
    from: deployer,
    log: true
    // args: [process.env.ORIGINAL_TOKEN_VAULT_SIGS_VERIFIER]
  });
  await hre.run('verify:verify', { address: mockMessageBridge.address });
};

deployFunc.tags = ['MockMessageBridge'];
deployFunc.dependencies = [];
export default deployFunc;
