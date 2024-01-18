import * as dotenv from 'dotenv';
import { ethers } from 'hardhat';
import { DeployFunction } from 'hardhat-deploy/types';
import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { TxVerifier } from '../../typechain';
import { verify } from '../utils/utils';

dotenv.config();

const deployFunc: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const chunks = await deployments.get('BlockChunks');

  const args = [chunks.address];
  const deployment = await deploy('TxVerifier', {
    from: deployer,
    log: true,
    args: args
  });

  const [signer] = await ethers.getSigners();
  const txVerfier = await (await ethers.getContract<TxVerifier>('TxVerifier')).connect(signer);
  const verifier = await deployments.get('TransactionProofVerifier');
  const tx = await txVerfier.updateVerifierAddress(1, verifier.address);
  console.log('verifier.address', verifier.address);
  await tx.wait();
  console.log('updateVerifierAddress tx', tx.hash);

  await verify(hre, deployment, args);
};

deployFunc.tags = ['TxVerifier'];
deployFunc.dependencies = [];
export default deployFunc;
