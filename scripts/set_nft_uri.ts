import {ethers} from 'hardhat';
import {BrevisUniNFT__factory} from './../typechain/factories/BrevisUniNFT__factory';

const run = async () => {
  const [signer] = await ethers.getSigners();

  const tierNFTs = [
    { contract: '0x12a6c023Db6eC119Fe616cDcd138c65532035d1B', uri: "https://getcelerapp.s3.us-west-1.amazonaws.com/brevis-website/nfts/uni_stone.svg"},
    { contract: '0x3d3F0b7Ab3efBFF17F2c38d610dF50BD62E5B322', uri: "https://getcelerapp.s3.us-west-1.amazonaws.com/brevis-website/nfts/uni_bronze.svg"},
    { contract: '0x159a1B1397C59c5Ab9F1DE43f104E69687935a84', uri: "https://getcelerapp.s3.us-west-1.amazonaws.com/brevis-website/nfts/uni_silver.svg"},
    { contract: '0xd12271CcC26e2745f38A1221284085dC0a6181d5', uri: "https://getcelerapp.s3.us-west-1.amazonaws.com/brevis-website/nfts/uni_gold.svg"},
    { contract: '0x0D743383B75924B83b8596A0E9bb3ec8ad8ddF78', uri: "https://getcelerapp.s3.us-west-1.amazonaws.com/brevis-website/nfts/uni_platinum.svg"},
    { contract: '0x8B01E95d6bB7EE8B04C27281a5c261938cF30eBc', uri: "https://getcelerapp.s3.us-west-1.amazonaws.com/brevis-website/nfts/uni_diamond.svg"}
  ];
  
  for (const nft of tierNFTs) {
    const contract = BrevisUniNFT__factory.connect(nft.contract, signer);
    const tx = await contract.setBaseURI(nft.uri);
    console.log(`setBaseURI(${nft.uri}) tx: ${tx.hash}`);
    await tx.wait();
  }
};

run();
