import * as dotenv from 'dotenv';
import { DeployFunction } from 'hardhat-deploy/types';
import { HardhatRuntimeEnvironment } from 'hardhat/types';

dotenv.config();

const deployFunc: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const bscEd25519Verifier = process.env.BSC_ED25519_VERIFIER;
  const args = [bscEd25519Verifier];

  const tmLightClient = await deploy('TendermintLightClient', {
    from: deployer,
    log: true,
    args: args,
    libraries: {
      Tendermint: process.env.BSC_TENDERMINT as string
    }
  });
  await hre.run('verify:verify', { address: tmLightClient.address, constructorArguments: args });
};

deployFunc.tags = ['BSCTendermintLightClient'];
deployFunc.dependencies = [];
export default deployFunc;
