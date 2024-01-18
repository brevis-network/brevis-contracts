import {deployments, ethers} from 'hardhat';
import {UniswapVolume__factory} from '../typechain';

const params = [
  {
    chainId: 1,
    router: '0xEf1c6E67703c7BD7107eed8303Fbe6EC2554BF6B',
    weth: '0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2',
    usdc: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
    usdcDecimal: 6
  },
  {
    chainId: 137,
    router: '0x4C60051384bd2d3C01bfc845Cf5F4b44bcbE9de5',
    weth: '0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270',
    usdc: '0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174',
    usdcDecimal: 6
  }
];

const tiers = [0, 1, 2, 3, 4, 5];
const tierNFTs = [
  '0x12a6c023Db6eC119Fe616cDcd138c65532035d1B',
  '0x3d3F0b7Ab3efBFF17F2c38d610dF50BD62E5B322',
  '0x159a1B1397C59c5Ab9F1DE43f104E69687935a84',
  '0xd12271CcC26e2745f38A1221284085dC0a6181d5',
  '0x0D743383B75924B83b8596A0E9bb3ec8ad8ddF78',
  '0x8B01E95d6bB7EE8B04C27281a5c261938cF30eBc'
];

const setParams = async () => {
  const [signer] = await ethers.getSigners();

  const dep = await deployments.get('UniswapVolume');
  const uniVol = UniswapVolume__factory.connect(dep.address, signer);

  for (const p of params) {
    let tx = await uniVol.setUniversalRouter(p.chainId, p.router);
    console.log(`setUniversalRouter(${p.chainId}, ${p.router}) tx: ${tx.hash}`);
    await tx.wait();

    tx = await uniVol.setWETH(p.chainId, p.weth);
    console.log(`setWETH(${p.chainId}, ${p.weth}) tx: ${tx.hash}`);
    await tx.wait();

    tx = await uniVol.setUSDC(p.chainId, p.usdc, p.usdcDecimal);
    console.log(`setUSDC(${p.chainId}, ${p.usdc}, ${p.usdcDecimal}) tx: ${tx.hash}`);
    await tx.wait();
  }
};

const setTierNFTs = async () => {
  const [signer] = await ethers.getSigners();

  const dep = await deployments.get('UniswapVolume');
  const uniVol = UniswapVolume__factory.connect(dep.address, signer);

  const tx = await uniVol.setTierNFTs(tiers, tierNFTs);
  console.log(`setTierNFTs(${tiers}, ${tierNFTs}) tx: ${tx.hash}`);
  await tx.wait();
};

const run = async () => {
  await setParams();
  await setTierNFTs();
};

run();
