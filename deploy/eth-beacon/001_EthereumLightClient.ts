import * as dotenv from 'dotenv';
import {DeployFunction} from 'hardhat-deploy/types';
import {HardhatRuntimeEnvironment} from 'hardhat/types';
import {verify} from '../utils/utils';

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
    7954695,
    '0xd5b62b78c9ee5326b007ef9a2dedce9bebcef5d63dace04df92dbb4c111508ff', // period 968 sha root
    '0x098c4f5caada71c8736e9a0d6e25a0c7230afb9d63bc0904daa5c75a8fabeefc', // period 968 poseidon root
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
