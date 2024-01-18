import { Deployment } from 'hardhat-deploy/types';
import { HardhatRuntimeEnvironment } from 'hardhat/types';
export function sleep(ms: number) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

export const verify = async (hre: HardhatRuntimeEnvironment, deployment: Deployment, args?: any) => {
  console.log('sleeping 5 seconds before verifying contract ' + deployment.address);
  await sleep(5000);
  console.log('verifying contract ' + deployment.address);
  try {
    return hre
      .run('verify:verify', {
        address: deployment.address,
        constructorArguments: args ?? deployment.args
      })
      .then(() => console.log(deployment.address + ' verified'));
  } catch (e: any) {
    if (e.message.toLowerCase().includes('already verified')) {
      console.log(deployment.address + ' already verified');
    } else {
      console.log(e);
    }
  }
};
