import { Fixture } from 'ethereum-waffle';
import { Wallet } from 'ethers';
import { waffle } from 'hardhat';
import { AnchorBlocks } from '../../typechain';
import { EthereumLightClient } from '../../typechain/EthereumLightClient';
import { deployAnchorBlocks, deployLightClient } from './deploy';
export interface LightClientFixture {
  lightClient: EthereumLightClient;
  anchorBlocks: AnchorBlocks;
}

export function loadFixture<T>(fixture: Fixture<T>): Promise<T> {
  const provider = waffle.provider;
  return waffle.createFixtureLoader(provider.getWallets(), provider)(fixture);
}

export const lightClientFixture = async ([admin]: Wallet[]): Promise<LightClientFixture> => {
  const lc = await deployLightClient(admin);
  const ab = await deployAnchorBlocks(admin, lc.address);
  return { lightClient: lc, anchorBlocks: ab };
};
