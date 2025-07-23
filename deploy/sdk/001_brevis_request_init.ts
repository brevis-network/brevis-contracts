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


  await deploy('BrevisRequest', {
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
  const proxy = await deployments.get('BrevisRequest_Proxy');
  console.log('BrevisRequest_Proxy', proxy.address);
  const brevisProof = await deployments.get('BrevisRequest_Implementation');
  await hre.run('verify:verify', { address: brevisProof.address, constructorArguments: args });

};

deployFunc.tags = ['BrevisRequestInit'];
deployFunc.dependencies = [];
export default deployFunc;

