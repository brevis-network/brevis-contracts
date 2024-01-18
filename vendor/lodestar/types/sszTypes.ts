export { ssz as altair } from './altair/index';
export { ssz as bellatrix } from './bellatrix/index';
export { ssz as capella } from './capella/index';
export { ssz as deneb } from './deneb/index';
export { ssz as phase0 } from './phase0/index';
export * from './primitive/sszTypes';

import { ssz as allForksSsz } from './allForks/index';
export const allForks = allForksSsz.allForks;
export const allForksBlinded = allForksSsz.allForksBlinded;
export const allForksExecution = allForksSsz.allForksExecution;
export const allForksBlobs = allForksSsz.allForksBlobs;
export const allForksLightClient = allForksSsz.allForksLightClient;
