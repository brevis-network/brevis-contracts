import * as dotenv from 'dotenv';
import {DeployFunction} from 'hardhat-deploy/types';
import {HardhatRuntimeEnvironment} from 'hardhat/types';
import {verify} from '../utils/utils';

dotenv.config();

const deployFunc: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre;
  const { deploy } = deployments;
  const { deployer } = await getNamedAccounts();

  const brevisProof = await deployments.get("BrevisProof");
  const args = [brevisProof.address]
  const deployment = await deploy('SlotValueExample', {
    from: deployer,
    log: true,
    args: args,
  });

  await verify(hre, deployment);
};

deployFunc.tags = ['SlotValueExample'];
deployFunc.dependencies = [];
export default deployFunc;
