import { ethers } from 'hardhat';

import { AnchorBlocks } from '../../typechain';
import { EthereumLightClient } from '../../typechain/contracts/light-client-eth/EthereumLightClient';
import { deployAnchorBlocks, deployLightClient } from './deploy';

export interface LightClientFixture {
  lightClient: EthereumLightClient;
  anchorBlocks: AnchorBlocks;
}

export const lightClientFixture = async (): Promise<LightClientFixture> => {
  const [admin] = await ethers.getSigners();
  const lc = await deployLightClient(admin);
  const lcAddress = await lc.getAddress();
  const ab = await deployAnchorBlocks(admin, lcAddress);
  return { lightClient: lc, anchorBlocks: ab };
};
