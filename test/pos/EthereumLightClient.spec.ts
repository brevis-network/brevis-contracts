import { expect } from 'chai';
import { zeroPad } from 'ethers/lib/utils';
import { ethers } from 'hardhat';
import { lightClientFixture, LightClientFixture, loadFixture } from './fixture';
import { getSyncCommitteeRoot, getSyncPeriodBySlot, newLightClientUpdate, newOptimisticUpdate } from './helper';
import proof638 from './proof_638.json';
import update637 from './update_637.json';
import update638 from './update_638.json';

let f: LightClientFixture;

describe('EthereumLightClient.processLightClientUpdate()', () => {
  beforeEach(async () => {
    f = await loadFixture(lightClientFixture);
  });

  it('rejects update if finality proof is invalid', async () => {
    const [admin] = await ethers.getSigners();
    const update = newLightClientUpdate(getUpdate638().data, proof638);

    update.finalityBranch = update.finalityBranch.slice(1, -1);
    let tx = f.lightClient.connect(admin).processLightClientUpdate(update);
    await expect(tx).to.be.revertedWith('bad fin proof');

    update.finalityBranch = [];
    tx = f.lightClient.connect(admin).processLightClientUpdate(update);
    await expect(tx).to.be.revertedWith('no fin proof');
  });

  it('rejects update if execution state proof is invalid', async () => {
    const [admin] = await ethers.getSigners();

    let update = newLightClientUpdate(getUpdate638().data, proof638);
    update.finalizedHeader.executionRoot.branch = [];
    let tx = f.lightClient.connect(admin).processLightClientUpdate(update);
    await expect(tx).to.be.revertedWith('bad exec root proof');

    update = newLightClientUpdate(getUpdate638().data, proof638);
    update.finalizedHeader.executionRoot.leaf = zeroPad('0x01', 32);
    tx = f.lightClient.connect(admin).processLightClientUpdate(update);
    await expect(tx).to.be.revertedWith('bad exec root proof');
  });

  it('rejects update if committee root mapping proof is invalid', async () => {
    const [admin] = await ethers.getSigners();
    const root = getSyncCommitteeRoot(
      getUpdate638().data.next_sync_committee.pubkeys,
      getUpdate638().data.next_sync_committee.aggregate_pubkey
    );
    let update = newLightClientUpdate(getUpdate638().data, proof638);
    update.nextSyncCommitteeRootMappingProof.a[0] = '0x0';
    let tx = f.lightClient.connect(admin).processLightClientUpdate(update);
    await expect(tx).reverted;

    update = newLightClientUpdate(getUpdate638().data, proof638);
    update.nextSyncCommitteeRoot = zeroPad('0x01', 32); // wrong root
    tx = f.lightClient.connect(admin).processLightClientUpdate(update);
    await expect(tx).to.be.revertedWith('bad next sync committee proof');

    update = newLightClientUpdate(getUpdate638().data, proof638);
    update.nextSyncCommitteeRoot = root;
    update.nextSyncCommitteeBranch = [zeroPad('0x01', 32)]; // wrong proof
    tx = f.lightClient.connect(admin).processLightClientUpdate(update);
    await expect(tx).to.be.revertedWith('bad next sync committee proof');

    update = newLightClientUpdate(getUpdate638().data, proof638);
    update.nextSyncCommitteeBranch = [];
    tx = f.lightClient.connect(admin).processLightClientUpdate(update);
    await expect(tx).to.be.revertedWith('no next sync committee proof');
  });

  it("rejects update if the update's poseidon root doesn't match light client's poseidon root", async () => {
    const [admin] = await ethers.getSigners();
    const update = newLightClientUpdate(getUpdate638().data, proof638);
    update.syncAggregate.poseidonRoot = zeroPad('0x01', 32);
    const tx = f.lightClient.connect(admin).processLightClientUpdate(update);
    await expect(tx).to.be.revertedWith('bad poseidon root');
  });

  it('rejects update if sig proof is invalid', async () => {
    const [admin] = await ethers.getSigners();
    let update = newLightClientUpdate(getUpdate638().data, proof638);

    update.syncAggregate.proof.a[0] = '0x0';
    const tx = f.lightClient.connect(admin).processLightClientUpdate(update);
    await expect(tx).reverted;
  });

  it('processes finality/committee update', async () => {
    const [admin] = await ethers.getSigners();
    const update = newLightClientUpdate(getUpdate638().data, proof638);
    const tx = await f.lightClient.connect(admin).processLightClientUpdate(update);
    const fh = getUpdate638().data.finalized_header;
    const ah = getUpdate638().data.attested_header;

    // check finalized
    await expect(tx).to.emit(f.lightClient, 'FinalityUpdate').withArgs(fh.beacon.slot, fh.execution.state_root);
    const finSlot = await f.lightClient.connect(admin).finalizedSlot();
    expect(finSlot).equals(fh.beacon.slot);
    const res = await f.lightClient.connect(admin).finalizedExecutionStateRootAndSlot();
    expect(res.root).equals(fh.execution.state_root);
    expect(res.slot).equals(fh.beacon.slot);

    // check next committee
    const root639 = getSyncCommitteeRoot(
      getUpdate638().data.next_sync_committee.pubkeys,
      getUpdate638().data.next_sync_committee.aggregate_pubkey
    );
    await expect(tx)
      .to.emit(f.lightClient, 'SyncCommitteeUpdated')
      .withArgs(getSyncPeriodBySlot(parseInt(ah.beacon.slot)) + 1, root639, update.nextSyncCommitteePoseidonRoot);
    const { currentRoot, nextRoot } = await f.lightClient.connect(admin).latestFinalizedSlotAndCommitteeRoots();
    const root638 = getSyncCommitteeRoot(
      update637.data.next_sync_committee.pubkeys,
      update637.data.next_sync_committee.aggregate_pubkey
    );
    expect(currentRoot).equal(root638);
    expect(nextRoot).equal(root639);
  });

  it('processes optimistic update', async () => {
    const [admin] = await ethers.getSigners();
    const update = newOptimisticUpdate(getUpdate638().data, proof638);
    const tx = await f.lightClient.connect(admin).processLightClientUpdate(update);
    const ah = getUpdate638().data.attested_header;

    // check optimistic
    await expect(tx).to.emit(f.lightClient, 'OptimisticUpdate').withArgs(ah.beacon.slot, ah.execution.state_root);
    const opSlot = await f.lightClient.connect(admin).optimisticSlot();
    expect(opSlot).equals(ah.beacon.slot);
    const res1 = await f.lightClient.connect(admin).optimisticExecutionStateRootAndSlot();
    expect(res1.root).equals(ah.execution.state_root);
    expect(res1.slot).equals(ah.beacon.slot);
  });
});

function getUpdate638() {
  return copy(update638);
}

function copy<T>(o: T): T {
  const j = JSON.stringify(o);
  return JSON.parse(j);
}
