import { Wallet } from 'ethers';

import { AnchorBlocks } from '../../typechain';
import { EthereumLightClient } from '../../typechain/contracts/light-client-eth/EthereumLightClient';
import { deployAnchorBlocks, deployLightClient } from './deploy';

export interface LightClientFixture {
  lightClient: EthereumLightClient;
  anchorBlocks: AnchorBlocks;
}

export const lightClientFixture = async ([admin]: Wallet[]): Promise<LightClientFixture> => {
  const lc = await deployLightClient(admin);
  const ab = await deployAnchorBlocks(admin, lc.address);
  return { lightClient: lc, anchorBlocks: ab };
};
