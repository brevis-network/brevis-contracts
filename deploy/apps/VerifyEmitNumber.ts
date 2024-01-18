import * as dotenv from 'dotenv';
import { DeployFunction } from 'hardhat-deploy/types';
import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { verify } from '../utils/utils';

dotenv.config();

const deployFunc: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
    const { deployments, getNamedAccounts } = hre;
    const { deploy } = deployments;
    const { deployer } = await getNamedAccounts();

    const verifier = await deployments.get('ReceiptVerifier');

    const args = [verifier.address];
    const deployment = await deploy('VerifyEmitNumber', {
        from: deployer,
        log: true,
        args: args
    });
    await verify(hre, deployment, args);
};

deployFunc.tags = ['VerifyEmitNumber'];
deployFunc.dependencies = [];
export default deployFunc;
