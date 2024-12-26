// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

//https://github.com/Layr-Labs/eigenlayer-contracts?tab=readme-ov-file#current-mainnet-deployment
library EigenMainnetDeployments {
    address internal constant DelegationManager = 0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A;
    address internal constant EigenPodManager = 0x91E677b07F7AF907ec9a428aafA9fc14a0d3A338;
    address internal constant StrategyManager = 0x858646372CC42E1A627fcE94aa7A7033e7CF075A;
    address internal constant Slasher = 0xD92145c07f8Ed1D392c1B88017934E301CC1c3Cd;
    address internal constant AVSDirectory = 0x135DDa560e946695d6f155dACaFC6f1F25C1F5AF;
    address internal constant RewardsCoordinator = 0x7750d328b314EfFa365A0402CcfD489B80B0adda;
    // TODO
    address internal constant AllocationManager = 0x78469728304326CBc65f8f95FA756B0B73164462;

    // strategies
    address internal constant cbETHStrategy = 0x54945180dB7943c0ed0FEE7EdaB2Bd24620256bc;
    address internal constant stETHStrategy = 0x93c4b944D05dfe6df7645A86cd2206016c51564D;
    address internal constant rETHStrategy = 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2;
    address internal constant ETHxStrategy = 0x9d7eD45EE2E8FC5482fa2428f15C971e6369011d;
    address internal constant ankrETHStrategy = 0x13760F50a9d7377e4F20CB8CF9e4c26586c658ff;
    address internal constant OETHStrategy = 0xa4C637e0F704745D182e4D38cAb7E7485321d059;
    address internal constant osETHStrategy = 0x57ba429517c3473B6d34CA9aCd56c0e735b94c02;
    address internal constant swETHStrategy = 0x0Fe4F44beE93503346A3Ac9EE5A26b130a5796d6;
    address internal constant wBETHStrategy = 0x7CA911E83dabf90C90dD3De5411a10F1A6112184;
    address internal constant sfrxETHStrategy = 0x8CA7A5d6f3acd3A7A8bC468a8CD0FB14B6BD28b6;
    address internal constant LsETHStrategy = 0xAe60d8180437b5C34bB956822ac2710972584473;
    address internal constant mETHSTrategy = 0x298aFB19A105D59E74658C4C334Ff360BadE6dd2;
    address internal constant beaconETHStrategy = 0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0;
}

//https://github.com/Layr-Labs/eigenlayer-contracts?tab=readme-ov-file#current-testnet-deployment
library EigenHoleSkyDeployments {
    //core
    address internal constant DelegationManager = 0xA44151489861Fe9e3055d95adC98FbD462B948e7;
    address internal constant EigenPodManager = 0x30770d7E3e71112d7A6b7259542D1f680a70e315;
    address internal constant StrategyManager = 0xdfB5f6CE42aAA7830E94ECFCcAd411beF4d4D5b6;
    address internal constant Slasher = 0xcAe751b75833ef09627549868A04E32679386e7C;
    address internal constant AVSDirectory = 0x055733000064333CaDDbC92763c58BF0192fFeBf;
    address internal constant RewardsCoordinator = 0xAcc1fb458a1317E886dB376Fc8141540537E68fE;
    address internal constant AllocationManager = 0x78469728304326CBc65f8f95FA756B0B73164462;

    // strategies
    address internal constant cbETHStrategy = 0x70EB4D3c164a6B4A5f908D4FBb5a9cAfFb66bAB6;
    address internal constant stETHStrategy = 0x7D704507b76571a51d9caE8AdDAbBFd0ba0e63d3;
    address internal constant rETHStrategy = 0x3A8fBdf9e77DFc25d09741f51d3E181b25d0c4E0;
    address internal constant ETHxStrategy = 0x31B6F59e1627cEfC9fA174aD03859fC337666af7;

    address internal constant ankrETHStrategy = 0x7673a47463F80c6a3553Db9E54c8cDcd5313d0ac;
    address internal constant osETHStrategy = 0x46281E3B7fDcACdBa44CADf069a94a588Fd4C6Ef;
    address internal constant sfrxETHStrategy = 0x9281ff96637710Cd9A5CAcce9c6FAD8C9F54631c;
    address internal constant LsETHStrategy = 0x05037A81BD7B4C9E0F7B430f1F2A22c31a2FD943;
    address internal constant mETHSTrategy = 0xaccc5A86732BE85b5012e8614AF237801636F8e5;
    address internal constant beaconETHStrategy = 0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0;
}
