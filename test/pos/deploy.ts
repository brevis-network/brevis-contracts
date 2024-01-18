import { Wallet } from 'ethers';
import { ethers } from 'hardhat';
import {
  AnchorBlocks,
  AnchorBlocks__factory,
  BeaconVerifier,
  BeaconVerifier__factory,
  EthereumLightClient__factory
} from '../../typechain';
import { EthereumLightClient } from '../../typechain/EthereumLightClient';
import { getSyncCommitteeRoot } from './helper';
import update637 from './update_637.json';

export async function deployLightClient(admin: Wallet): Promise<EthereumLightClient> {
  const zkVerifier = await deployBeaconVerifier(admin);
  const factory = await ethers.getContractFactory<EthereumLightClient__factory>('EthereumLightClient');
  return factory.connect(admin).deploy(
    1616508000,
    '0x043db0d9a83813551ee2f33450d23797757d430911a9320530ad8a0eabc43efb',
    [0, 36660, 112260, 162304],
    ['0x00001020', '0x01001020', '0x02001020', '0x03001020'],
    5226496,
    getSyncCommitteeRoot(
      update637.data.next_sync_committee.pubkeys,
      update637.data.next_sync_committee.aggregate_pubkey
    ),
    '0x0b90f32a0d03c13a909e40b8bfa0dc049a9f3f8fbaab12ae880a0fa540e709ca', // 638 committee poseidon root
    zkVerifier.address
  );
}

export async function deployBeaconVerifier(admin: Wallet): Promise<BeaconVerifier> {
  const factory = await ethers.getContractFactory<BeaconVerifier__factory>('BeaconVerifier');
  return factory.connect(admin).deploy();
}

export async function deployAnchorBlocks(admin: Wallet, lightClient: string): Promise<AnchorBlocks> {
  const factory = await ethers.getContractFactory<AnchorBlocks__factory>('AnchorBlocks');
  return factory.connect(admin).deploy(lightClient);
}
