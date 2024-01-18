import * as dotenv from 'dotenv';
import {ethers} from 'hardhat';
import {DeployFunction} from 'hardhat-deploy/types';
import {HardhatRuntimeEnvironment} from 'hardhat/types';
import {BlockChunks} from '../../typechain';
import {verify} from '../utils/utils';

dotenv.config();

const deployFunc: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const deployment = await deploy('BlockChunks', {
    from: deployer,
    log: true
  });
  const [signer] = await ethers.getSigners();
  const contract = await (await ethers.getContract<BlockChunks>('BlockChunks')).connect(signer);
  const lc = await deployments.get('AnchorBlocks');
  await contract.updateAnchorBlockProvider(1, lc.address);
  const verifier = await deployments.get('EthChunkOf128Verifier');
  await contract.updateVerifierAddress(1, verifier.address);
  await verify(hre, deployment);
};

deployFunc.tags = ['BlockChunks'];
deployFunc.dependencies = [];
export default deployFunc;
