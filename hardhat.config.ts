import '@nomiclabs/hardhat-ethers';
import '@nomicfoundation/hardhat-verify';
import '@nomiclabs/hardhat-waffle';
import '@typechain/hardhat';
import 'hardhat-contract-sizer';
import 'hardhat-deploy';
import 'hardhat-gas-reporter';
import '@rumblefishdev/hardhat-kms-signer';

import * as dotenv from 'dotenv';
import { HardhatUserConfig } from 'hardhat/types';

dotenv.config();

const DEFAULT_ENDPOINT = 'http://localhost:8545';
const DEFAULT_PRIVATE_KEY =
  process.env.DEFAULT_PRIVATE_KEY || 'ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff';
const holeskyEndpoint = process.env.HOLESKY_ENDPOINT || DEFAULT_ENDPOINT;
const holeskyPrivateKey = process.env.HOLESKY_PRIVATE_KEY || DEFAULT_PRIVATE_KEY;
const sepoliaEndpoint = process.env.SEPOLIA_ENDPOINT || DEFAULT_ENDPOINT;
const sepoliaPrivateKey = process.env.SEPOLIA_PRIVATE_KEY || DEFAULT_PRIVATE_KEY;
const bscTestEndpoint = process.env.BSC_TEST_ENDPOINT || DEFAULT_ENDPOINT;
const bscTestPrivateKey = process.env.BSC_TEST_PRIVATE_KEY || DEFAULT_PRIVATE_KEY;
const avalancheTestEndpoint = process.env.AVALANCHE_TEST_ENDPOINT || DEFAULT_ENDPOINT;
const avalancheTestPrivateKey = process.env.AVALANCHE_TEST_PRIVATE_KEY || DEFAULT_PRIVATE_KEY;

const config: HardhatUserConfig = {
  defaultNetwork: 'hardhat',
  networks: {
    // Testnets
    hardhat: {
      blockGasLimit: 120_000_000
    },
    localhost: { timeout: 600000 },
    holesky: {
      url: holeskyEndpoint,
      accounts: [`0x${holeskyPrivateKey}`]
    },
    sepolia: {
      url: sepoliaEndpoint,
      accounts: [`0x${sepoliaPrivateKey}`]
    },
    bscTest: {
      url: bscTestEndpoint,
      accounts: [`0x${bscTestPrivateKey}`]
    },
    avalancheTest: {
      url: avalancheTestEndpoint,
      accounts: [`0x${avalancheTestPrivateKey}`]
    }
  },
  namedAccounts: {
    deployer: {
      default: 0
    }
  },
  solidity: {
    version: '0.8.20',
    settings: {
      optimizer: {
        enabled: true,
        runs: 800
      },
      viaIR: true
    }
  },
  gasReporter: {
    enabled: process.env.REPORT_GAS === 'true' ? true : false,
    noColors: true,
    outputFile: 'reports/gas_usage/summary.txt'
  },
  typechain: {
    outDir: 'typechain',
    target: 'ethers-v5'
  },
  etherscan: {
    apiKey: {
      goerli: process.env.ETHERSCAN_API_KEY as string,
      sepolia: process.env.ETHERSCAN_API_KEY as string,
      bscTestnet: process.env.BSCSCAN_API_KEY as string,
      avalancheFujiTestnet: process.env.SNOWTRACE_API_KEY as string
    }
  }
};

if (config.networks?.bscTest) {
  config.networks.bscTest.minMaxPriorityFeePerGas = 3000000000;
  config.networks.bscTest.minMaxFeePerGas = 3000000000;
  config.networks.bscTest.gasPrice = 10000000000;
}

export default config;
