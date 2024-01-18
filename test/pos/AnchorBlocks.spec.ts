import { expect } from 'chai';
import { lightClientFixture, LightClientFixture, loadFixture } from './fixture';
import header_8687478_rlp from './header_8687478_rlp.json';
import { LightClientUpdate, newHeaderWithExecution, newSyncAggregate, ZERO_BYTES_32 } from './helper';
import proof from './proof_638.json';
import update638 from './update_638.json';

let f: LightClientFixture;

describe('AnchorBlocks.processUpdate()', () => {
  beforeEach(async () => {
    f = await loadFixture(lightClientFixture);
  });

  it('rejects if quorum not reached', async () => {
    const head = newHeadBlock(update638.data);
    head.syncAggregate.participation = 123;
    const tx = f.anchorBlocks.processUpdate(head);
    await expect(tx).revertedWith('quorum not reached');
  });

  it('rejects if execution root is invalid', async () => {
    const head = newHeadBlock(update638.data);
    head.attestedHeader.executionRoot.leaf = ZERO_BYTES_32;
    const tx = f.anchorBlocks.processUpdate(head);
    await expect(tx).revertedWith('bad exec root proof');
  });

  it('rejects if execution blockNumber/blockHash is invalid', async () => {
    let head = newHeadBlock(update638.data);
    head.attestedHeader.execution.blockHash.leaf = ZERO_BYTES_32;
    let tx = f.anchorBlocks.processUpdate(head);
    await expect(tx).revertedWith('bad proof');
    head = newHeadBlock(update638.data);
    head.attestedHeader.execution.blockNumber.leaf = ZERO_BYTES_32;
    tx = f.anchorBlocks.processUpdate(head);
    await expect(tx).revertedWith('bad proof');
  });

  it('rejects if head block sig is invalid', async () => {
    const head = newHeadBlock(update638.data);
    head.syncAggregate.proof.a[0] = 0;
    const tx = f.anchorBlocks.processUpdate(head);
    await expect(tx).reverted;
  });

  it('processes update', async () => {
    const head = newHeadBlock(update638.data);
    const tx = f.anchorBlocks.processUpdate(head);
    const exec = update638.data.attested_header.execution;
    await expect(tx).to.emit(f.anchorBlocks, 'AnchorBlockUpdated').withArgs(exec.block_number, exec.block_hash);
    const blockHash = await f.anchorBlocks.blocks(exec.block_number);
    await expect(blockHash).to.equal(exec.block_hash);
  });
});

describe('AnchorBlocks.processUpdateWithChainProof()', () => {
  beforeEach(async () => {
    f = await loadFixture(lightClientFixture);
  });

  it('rejects if chain proof length is invalid', async () => {
    const head = newHeadBlock(update638.data);
    const parent = update638.data.attested_header.execution.parent_hash;
    const tx = f.anchorBlocks.processUpdateWithChainProof(head, parent, []);
    await expect(tx).revertedWith('invalid proof length');
  });

  it('rejects if block hash witness left part length is invalid', async () => {
    const head = newHeadBlock(update638.data);
    const exec = update638.data.attested_header.execution;
    const parent = exec.parent_hash;
    const [, right] = header_8687478_rlp.rlp.split(parent.replace('0x', ''));
    const w = { left: '0x00', right: '0x' + right };
    const tx = f.anchorBlocks.processUpdateWithChainProof(head, parent, [w]);
    await expect(tx).revertedWith('invalid left len');
  });

  it('processes update with chain proof', async () => {
    const head = newHeadBlock(update638.data);
    const exec = update638.data.attested_header.execution;
    const parent = exec.parent_hash;
    const [left, right] = header_8687478_rlp.rlp.split(parent.replace('0x', ''));
    const w = { left: '0x' + left, right: '0x' + right };
    const tx = f.anchorBlocks.processUpdateWithChainProof(head, parent, [w]);
    await expect(tx)
      .to.emit(f.anchorBlocks, 'AnchorBlockUpdated')
      .withArgs(parseInt(exec.block_number) - 1, exec.parent_hash);
    const blockHash = await f.anchorBlocks.blocks(parseInt(exec.block_number) - 1);
    await expect(blockHash).to.equal(exec.parent_hash);
  });
});

function newHeadBlock(u: LightClientUpdate) {
  return {
    attestedHeader: newHeaderWithExecution(u.attested_header),
    syncAggregate: newSyncAggregate(u.sync_aggregate.sync_committee_bits, proof.bls_sig_proof),
    signatureSlot: u.signature_slot
  };
}
