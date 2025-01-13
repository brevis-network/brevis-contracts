import '@nomicfoundation/hardhat-verify';
import '@nomiclabs/hardhat-ethers';
import '@nomiclabs/hardhat-waffle';
import '@typechain/hardhat';
import 'hardhat-contract-sizer';
import 'hardhat-deploy';
import 'hardhat-gas-reporter';
import '@rumblefishdev/hardhat-kms-signer';

import * as dotenv from 'dotenv';
import { HardhatUserConfig } from 'hardhat/types';

dotenv.config();

const privateKey =
  process.env.DEFAULT_PRIVATE_KEY || 'ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff';
const DEFAULT_PRIVATE_KEY = privateKey;
const goerliEndpoint = process.env.GOERLI_ENDPOINT || process.env.DEFAULT_ENDPOINT;
const goerliPrivateKey = process.env.GOERLI_PRIVATE_KEY || DEFAULT_PRIVATE_KEY;
const sepoliaEndpoint = process.env.SEPOLIA_ENDPOINT || process.env.DEFAULT_ENDPOINT;
const sepoliaPrivateKey = process.env.SEPOLIA_PRIVATE_KEY || DEFAULT_PRIVATE_KEY;
const bscTestEndpoint = process.env.BSC_TEST_ENDPOINT || process.env.DEFAULT_ENDPOINT;
const bscTestPrivateKey = process.env.BSC_TEST_PRIVATE_KEY || DEFAULT_PRIVATE_KEY;
const avalancheTestEndpoint = process.env.AVALANCHE_TEST_ENDPOINT || process.env.DEFAULT_ENDPOINT;
const avalancheTestPrivateKey = process.env.AVALANCHE_TEST_PRIVATE_KEY || DEFAULT_PRIVATE_KEY;
const holeskyEndpoint = process.env.HOLESKY_ENDPOINT || 'https://holesky.drpc.org';
const holeskyPrivateKey = process.env.HOLESKY_PRIVATE_KEY || DEFAULT_PRIVATE_KEY;
const optimismEndpoint = process.env.OPTIMISM_ENDPOINT ||  process.env.DEFAULT_ENDPOINT;
const optimismPrivateKey = process.env.OPTIMISM_PRIVATE_KEY || DEFAULT_PRIVATE_KEY;
const arbitrumSepoliaEndpoint = process.env.ARBITRUM_SEPOLIA_ENDPOINT ||  process.env.DEFAULT_ENDPOINT;
const arbitrumSepoliaPrivateKey = process.env.ARBITRUM_SEPOLIA_PRIVATE_KEY || DEFAULT_PRIVATE_KEY;
const baseSepoliaEndpoint = process.env.BASE_SEPOLIA_ENDPOINT ||  process.env.DEFAULT_ENDPOINT;
const baseSepoliaPrivateKey = process.env.BASE_SEPOLIA_PRIVATE_KEY || DEFAULT_PRIVATE_KEY;
const modeEndpoint = process.env.MODE_ENDPOINT ||  process.env.DEFAULT_ENDPOINT;
const modePrivateKey = process.env.MODE_PRIVATE_KEY || DEFAULT_PRIVATE_KEY;
const soneiumEndpoint= process.env.SONEIUM_ENDPOINT || process.env.DEFAULT_ENDPOINT;
const soneiumPrivateKey = process.env.MODE_PRIVATE_KEY || DEFAULT_PRIVATE_KEY;


const config: HardhatUserConfig = {
  defaultNetwork: 'hardhat',
  networks: {
    // Testnets
    hardhat: {
      blockGasLimit: 120_000_000
    },
    localhost: { timeout: 600000 },
    goerli: {
      url: goerliEndpoint || '',
      accounts: [`0x${goerliPrivateKey}`]
    },
    sepolia: {
      url: sepoliaEndpoint || '',
      accounts: [`0x${sepoliaPrivateKey}`]
    },
    bscTest: {
      url: bscTestEndpoint || '',
      accounts: [`0x${bscTestPrivateKey}`]
    },
    avalancheTest: {
      url: avalancheTestEndpoint || '',
      accounts: [`0x${avalancheTestPrivateKey}`]
    },
    holesky: {
      url: holeskyEndpoint,
      accounts: [`0x${holeskyPrivateKey}`]
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
    arbitrum: {
      url: process.env.ARBITRUM_ENDPOINT || 'https://arbitrum.llamarpc.com',
      accounts: [`0x${privateKey}`]
    },
    optimism: {
      url: optimismEndpoint || '',
      accounts: [`0x${optimismPrivateKey}`]
    },
    arbitrumSepolia: {
      url: arbitrumSepoliaEndpoint,
      accounts: [`0x${arbitrumSepoliaPrivateKey}`]
    },
    baseSepolia: {
      url: baseSepoliaEndpoint,
      accounts: [`0x${baseSepoliaPrivateKey}`]
    },
    mode: {
      url: modeEndpoint,
      accounts: [`0x${modePrivateKey}`],
      chainId: 34443,
    },
    soneium: {
      url: soneiumEndpoint,
      accounts: [`0x${soneiumPrivateKey}`],
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
  sourcify: {
    enabled: true,
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
      arbitrumOne: process.env.ARBISCAN_API_KEY as string,
      optimisticEthereum: process.env.OPSCAN_API_KEY as string,
      arbitrumSepolia: process.env.ARBISCAN_API_KEY as string,
      baseSepolia: process.env.BASESCAN_API_KEY as string,
      mode: "mode",
      soneium: process.env.SONEIUM_API_KEY as string,
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
      },
      {
        network: 'arbitrumSepolia',
        chainId: 421614,
        urls: {
          apiURL: 'https://api-sepolia.arbiscan.io/api',
          browserURL: 'https://sepolia.arbiscan.io/'
        }
      },
      {
        network: 'baseSepolia',
        chainId: 84532,
        urls: {
          apiURL: 'https://api-sepolia.basescan.org/api',
          browserURL: 'https://sepolia.basescan.io/'
        }
      },
      {
        network: "mode",
        chainId: 34443,
        urls: {
          apiURL: "https://api.routescan.io/v2/network/mainnet/evm/34443/etherscan",
          browserURL: "https://modescan.io"
        }
      },
      {
        network: "soneium",
        chainId: 1868,
        urls: {
          apiURL: "https://soneium.blockscout.com/api",
          browserURL: "https://soneium.blockscout.com/",
        },
      },
    ]
  }
};

if (config.networks?.bscTest) {
  config.networks.bscTest.minMaxPriorityFeePerGas = 3000000000;
  config.networks.bscTest.minMaxFeePerGas = 3000000000;
  config.networks.bscTest.gasPrice = 10000000000;
}

export default config;
