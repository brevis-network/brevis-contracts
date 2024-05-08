import '@nomicfoundation/hardhat-verify';
import '@nomiclabs/hardhat-ethers';
import '@nomiclabs/hardhat-waffle';
import '@typechain/hardhat';
import 'hardhat-contract-sizer';
import 'hardhat-deploy';
import 'hardhat-gas-reporter';

import * as dotenv from 'dotenv';
import { HardhatUserConfig } from 'hardhat/types';

dotenv.config();

const privateKey =
  process.env.DEFAULT_PRIVATE_KEY || 'ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff';

const config: HardhatUserConfig = {
  defaultNetwork: 'hardhat',
  networks: {
    // Testnets
    hardhat: {
      blockGasLimit: 120_000_000
    },
    localhost: { timeout: 600000 },
    goerli: {
      url: process.env.GOERLI_ENDPOINT || '',
      accounts: [`0x${privateKey}`]
    },
    sepolia: {
      url: process.env.SEPOLIA_ENDPOINT || '',
      accounts: [`0x${privateKey}`]
    },
    bscTest: {
      url: process.env.BSC_TEST_ENDPOINT || '',
      accounts: [`0x${privateKey}`]
    },
    avalancheTest: {
      url: process.env.AVALANCHE_TEST_ENDPOINT || '',
      accounts: [`0x${privateKey}`]
    },
    // Mainnet
    linea: {
      url: process.env.LINEA_ENDPOINT || '',
      accounts: [`0x${privateKey}`]
    },
    base: {
      url: process.env.BASE_ENDPOINT || '',
      accounts: [`0x${privateKey}`]
    },
    holesky: {
      url: process.env.HOLESKY_ENDPOINT || 'https://holesky.drpc.org',
      accounts: [`0x${privateKey}`]
    },
    arbitrum: {
      url: process.env.ARBITRUM_ENDPOINT || 'https://arbitrum.llamarpc.com',
      accounts: [`0x${privateKey}`]
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
      avalancheFujiTestnet: process.env.SNOWTRACE_API_KEY as string,
      linea: process.env.LINEASCAN_API_KEY as string,
      holesky: process.env.ETHERSCAN_API_KEY as string,
      base: process.env.BASESCAN_API_KEY as string,
      arbitrumOne: process.env.ARBISCAN_API_KEY as string
    },
    customChains: [
      {
        network: 'linea',
        chainId: 59144,
        urls: {
          apiURL: 'https://api.lineascan.build/api',
          browserURL: 'https://lineascan.build/'
        }
      },
      {
        network: 'holesky',
        chainId: 17000,
        urls: {
          apiURL: 'https://api-holesky.etherscan.io/api',
          browserURL: 'https://holesky.etherscan.io/'
        }
      }
    ]
  }
};

export default config;
