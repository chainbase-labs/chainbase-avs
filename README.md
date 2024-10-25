# Chainbase AVS

Welcome to the Chainbase AVS! This repository contains the source code and other relevant materials for the Chainbase AVS.

## Deployed Contracts

Here you can find the information about the deployed contracts:

### Current Testnet Deployment

The current testnet deployment is on Holesky testnet. You can view the deployed contract addresses below.

| Name                                                                                                       | Proxy                                                                                 | Implementation                                                                        | Notes                                                                                                                                                |
|------------------------------------------------------------------------------------------------------------| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |------------------------------------------------------------------------------------------------------------------------------------------------------|
| `chainbaseServiceManager` | [`0x0fb6b02F8482a06CF1d99558576F111abC377932`](https://holesky.etherscan.io/address/0x0fb6b02F8482a06CF1d99558576F111abC377932) | [`0xE7a1603b13a48aeA9f7Fc31EB42F4941ABece69F`](https://holesky.etherscan.io/address/0xE7a1603b13a48aeA9f7Fc31EB42F4941ABece69F) | Proxy: [`TUP@4.9.6`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.6/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| `registryCoordinator` | [`0xb8B9351a6D21fd9b3249713353ACf430BE2e6bBc`](https://holesky.etherscan.io/address/0xb8B9351a6D21fd9b3249713353ACf430BE2e6bBc) | [`0x606beE41FaeB30618d6378e875175b798c23e0D0`](https://holesky.etherscan.io/address/0x606beE41FaeB30618d6378e875175b798c23e0D0) | Proxy: [`TUP@4.9.6`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.6/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| `blsApkRegistry` | [`0xf9e64168901dbEcA17a3dDe80Bd184e12f46853a`](https://holesky.etherscan.io/address/0xf9e64168901dbEcA17a3dDe80Bd184e12f46853a) | [`0xaf5b82390C1fA6b451c28176e9A60DC1A00529B8`](https://holesky.etherscan.io/address/0xaf5b82390C1fA6b451c28176e9A60DC1A00529B8) | Proxy: [`TUP@4.9.6`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.6/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| `stakeRegistry` | [`0x076B775b2b32740B69ae8DCA1ade674BD8049231`](https://holesky.etherscan.io/address/0x076B775b2b32740B69ae8DCA1ade674BD8049231) | [`0x2dbFAC7bdAE9914EB7FBc8951e4D6807377a0dFf`](https://holesky.etherscan.io/address/0x2dbFAC7bdAE9914EB7FBc8951e4D6807377a0dFf) | Proxy: [`TUP@4.9.6`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.6/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| `indexRegistry` | [`0xB7aCf43E3A614b10d7Ac3E4082f30730DfeC64EE`](https://holesky.etherscan.io/address/0xB7aCf43E3A614b10d7Ac3E4082f30730DfeC64EE) | [`0x2faFA82475b1278DB1445a4Af7ac4e790B854883`](https://holesky.etherscan.io/address/0x2faFA82475b1278DB1445a4Af7ac4e790B854883) | Proxy: [`TUP@4.9.6`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.6/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |
| `proxyAdmin` | -                                                                                     | [`0xa92C5cb01fFa4B077Cc4AeA9d4fc67B622D97feD`](https://holesky.etherscan.io/address/0xa92C5cb01fFa4B077Cc4AeA9d4fc67B622D97feD) |                                                                                                                                                      |
| `pauserRegistry` | -                                                                                     | [`0xf093DBb28ae1CDA2c04A9398A9148ACA7dC3A445`](https://holesky.etherscan.io/address/0xf093DBb28ae1CDA2c04A9398A9148ACA7dC3A445) |                                                                                                                                                      |
| `operatorStateRetriever` | -                                                                                     | [`0x3e302917A4d007eAF367c182226Da217E9639d38`](https://holesky.etherscan.io/address/0x3e302917A4d007eAF367c182226Da217E9639d38) |                                                                                                                                                      |

### Audit Report

The project contracts has been audited by [SlowMist](https://www.slowmist.com/). You can find the audit report [here](https://github.com/chainbase-labs/chainbase-avs/blob/main/contracts/audit/Chainbase%20AVS%20-%20SlowMist%20Audit%20Report.pdf).

## Architecture

The basic architecture of the project is as follows:

<img src="https://raw.githubusercontent.com/chainbase-labs/chainbase-avs/refs/heads/main/doc/image/architecture.png" width="1000" height="600"/>

## Installation Guide

For installation instructions of setting up and running a Chainbase AVS operator node , please refer to the [Installation Guide](link_to_installation_guide).

## Official Website and Community

- **Official Website**: [https://chainbase.com](https://chainbase.com)
- **Discord Community**: Join our discussions and stay updated by joining our [Discord server](https://discord.com/invite/chainbase).

## Contributors

A big thank you to all the contributors who help in making this project better!

<table>
  <tr>
    <td align="center"><a href="https://github.com/lxcong"><img src="https://avatars.githubusercontent.com/u/8024426?v=4" width="100px;" alt=""/><br /><sub><b>lxcong</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/DeKaiju"><img src="https://avatars.githubusercontent.com/u/10413226?v=4" width="100px;" alt=""/><br /><sub><b>DeKaiju</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/ddl-hust"><img src="https://avatars.githubusercontent.com/u/24953789?v=4" width="100px;" alt=""/><br /><sub><b>ddl-hust</b></sub></a><br /></td>
  </tr>
</table>

Feel free to create a pull request to contribute to the project.

---
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.