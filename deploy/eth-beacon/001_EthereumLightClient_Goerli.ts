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
    1616508000,
    '0x043db0d9a83813551ee2f33450d23797757d430911a9320530ad8a0eabc43efb',
    [0, 36660, 112260, 162304],
    ['0x00001020', '0x01001020', '0x02001020', '0x03001020'],
    5840999,
    '0xe95f332ee309a0491b2b2b02dd7c42df6e97b2823ddf37008615ed744fb72f58', // period 713 sha root
    '0x13818e880443841df46822384794d13cd9d07357260c86e8ffb3960f84c431ae', // period 713 poseidon root
    verifier.address
  ];

  const deployment = await deploy('EthereumLightClient', {
    from: deployer,
    log: true,
    args: args
  });
  await verify(hre, deployment, args);
};

deployFunc.tags = ['EthereumLightClient_Goerli'];
deployFunc.dependencies = [];
export default deployFunc;
