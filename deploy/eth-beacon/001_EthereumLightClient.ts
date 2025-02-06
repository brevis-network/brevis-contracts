import * as dotenv from 'dotenv';
import { DeployFunction } from 'hardhat-deploy/types';
import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { verify } from '../utils/utils';

dotenv.config();

const deployFunc: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const verifier = await deployments.get('BeaconVerifier');

  const args = [
    1606824023,
    '0x4b363db94e286120d76eb905340fdd4e54bfe9f06bf33ff6cf5ad27f511bfe95',
    [0, 74240, 144896, 194048, 269568],
    ['0x00000000', '0x01000000', '0x02000000', '0x03000000', '0x04000000',],
    10990500,
    '0x99cfd99b21fd1a3a0d77d5096613f309fee2e832eec4d0d8359c49abb16e0192', // period 1341 sha root
    '0x094d7d1fced5abdac11cd4db4a59316f26bbbf5001d9de6cb7ac711a8c6b0890', // period 1341 poseidon root
    verifier.address
  ];

  const deployment = await deploy('EthereumLightClient', {
    from: deployer,
    log: true,
    args: args
  });
  await verify(hre, deployment, args);
};

deployFunc.tags = ['EthereumLightClient'];
deployFunc.dependencies = [];
export default deployFunc;
