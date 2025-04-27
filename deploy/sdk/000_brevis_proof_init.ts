import * as dotenv from 'dotenv';
import { DeployFunction } from 'hardhat-deploy/types';
import { HardhatRuntimeEnvironment } from 'hardhat/types';

dotenv.config();

const deployFunc: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const smt = await deployments.get('MockSMT');
  const args: [string] = [smt.address];

  await deploy('BrevisProof', {
    from: deployer,
    log: true,
    args: args,
    proxy: {
      proxyContract: 'OptimizedTransparentProxy',
      execute: {
        // only called when proxy is deployed, it'll call MessageBus contract.init
        // with proper args
        init: {
          methodName: 'init',
          args: args
        }
      }
    }
  });

  const proxyAdmin = await deployments.get('DefaultProxyAdmin');
  console.log('DefaultProxyAdmin', proxyAdmin.address);
  const proxy = await deployments.get('BrevisProof_Proxy');
  console.log('BrevisProof_Proxy', proxy.address);
  const brevisProof = await deployments.get('BrevisProof_Implementation');
  await hre.run('verify:verify', { address: brevisProof.address, constructorArguments: args });
};

deployFunc.tags = ['BrevisProofInit'];
deployFunc.dependencies = [];
export default deployFunc;
