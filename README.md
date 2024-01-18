# Contract Repo for Brevis: An Omnichain ZK Data Attestation Platform

## Light Client Contracts

Light client contracts model their corresponding chain's light client logic, replacing places of signature verification with zk proof verification to save on-chain gas cost.

| Path                     | Description                        |
| ------------------------ | ---------------------------------- |
| contracts/light-client   | Ethereum Beacon Chain light client |
| contracts/bsc-tendermint | BNB Beacon Chain light client      |
| contracts/poa            | BNB Smart Chain light client       |

## Application Layer Contracts

| Path                     | Description                                           |
| ------------------------ | ----------------------------------------------------- |
| contracts/message-bridge | arbitrary message passing built on top of brevis      |
| contracts/token-bridge   | peg-model token bridge built on top of message bridge |
