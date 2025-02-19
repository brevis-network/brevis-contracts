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
    11075584, // period finalized slot
    '0xb549827c59146c49262326cbc4747602e1dc45ab6af675fa01c80a47f8a676e8', // period sha root
    '0x23fe301686e7607e52750afd172da048840af65eac8e9cc246a677d0085e8812', // period poseidon root
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
