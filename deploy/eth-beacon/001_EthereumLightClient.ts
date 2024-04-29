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
    [0, 74240, 144896, 194048],
    ['0x00000000', '0x01000000', '0x02000000', '0x03000000'],
    8523466,
    '0xc8a45b4fdfda703822b0e654a2ba3fc454952f8b2698b85e8666d6061b8d117c', // period 1040 sha root
    '0x0256ceec42a8c033d79c0977cf6135ff52a432e8293e3f653800e5300b79939d', // period 1040 poseidon root
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
