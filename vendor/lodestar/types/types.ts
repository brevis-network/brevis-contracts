import { Slot } from './primitive/types';

export { ts as allForks } from './allForks/index';
export { ts as altair } from './altair/index';
export { ts as bellatrix } from './bellatrix/index';
export { ts as capella } from './capella/index';
export { ts as deneb } from './deneb/index';
export { ts as phase0 } from './phase0/index';
export * from './primitive/types';

/** Common non-spec type to represent roots as strings */
export type RootHex = string;

/** Handy enum to represent the block production source */
export enum BlockSource {
  builder = 'builder',
  engine = 'engine'
}

export type SlotRootHex = { slot: Slot; root: RootHex };
export type SlotOptionalRoot = { slot: Slot; root?: RootHex };
