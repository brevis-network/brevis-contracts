import { ethers } from 'hardhat';
import { getSyncCommitteeRoot } from './helper';
import update638 from './update_638.json';

const root639 = getSyncCommitteeRoot(
  update638.data.next_sync_committee.pubkeys,
  update638.data.next_sync_committee.aggregate_pubkey
);
//console.log('committee 639 root', ethers.utils.hexlify(root639));
