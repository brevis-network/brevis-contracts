import * as dotenv from 'dotenv';
import {DeployFunction} from 'hardhat-deploy/types';
import {HardhatRuntimeEnvironment} from 'hardhat/types';
import {verify} from '../utils/utils';

dotenv.config();

const defaultSmtRoot = '0x02e4b36d42d41af9db0526f588fa0c12113da9ad7db4e3e1a73b006c96a20364';

const deployFunc: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
    const { deployments, getNamedAccounts } = hre;
    const { deploy } = deployments;
    const { deployer } = await getNamedAccounts();

    const anchorProvider = await deployments.get('AnchorBlocks');
    // const anchorProvider = await deployments.get('SameChainAnchorBlocks');
    const verifier = await deployments.get('SMTUpdateCircuitProofVerifier');
    const args = [[1], [anchorProvider.address], [verifier.address], [defaultSmtRoot]];
    const deployment = await deploy('SMT', { from: deployer, log: true, args: args });
    await verify(hre, deployment);
};

deployFunc.tags = ['SMT'];
deployFunc.dependencies = [];
export default deployFunc;
