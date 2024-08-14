
# Chainbase Contracts 

Chainbase AVS Core Contracts

## Holesky contract address

| Name | Proxy | Implementation | Notes |
| -------- | ------- | -------- | -------- |
| [`ProxyAdmin`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/proxy/transparent/ProxyAdmin.sol) |  | [`0xdFbD62c5d8C5739852f67F2D7d2148FC5Bf2ce8E`](https://holesky.etherscan.io/address/0xdfbd62c5d8c5739852f67f2d7d2148fc5bf2ce8e) | onwer:0x8Ce42d4483dd0A11708F2848F8fdF13E3070713e |
| [`AVS`](https://github.com/chainbase-labs/chainbase-avs-contracts/blob/main/src/AVS.sol) |[`0x5e78eff26480a75e06ccdabe88eb522d4d8e1c9d`](https://holesky.etherscan.io/address/0x5e78eff26480a75e06ccdabe88eb522d4d8e1c9d#code) | [`0x4d6fa2c2a2c64855426f154cee5d661fdbf557a2`](https://holesky.etherscan.io/address/0x4d6fa2c2a2c64855426f154cee5d661fdbf557a2#code) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

## Deployment
 **Install Dependencies**

`forge install`


 **Deploy contract**

- Deploy AVS proxy && impl

`forge script --chain holesky script/AVS.s.sol:DeployAVS --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast -vvvv --verify`

- Deploy ProxyAdmin

`forge script --chain holesky script/DeplyProxyAdmin.s.sol:DeployProxyAdmin --rpc-url $RPC_URL --private-key $PROXY_ADMIN_DEPLOYER_KEY --broadcast -vvvv --verify`

- Just AVS Impl(used for upgrade)

```shell
forge create --chain holesky \
    --constructor-args "$DELEGATION_MANAGER_ADDRESS" "$AVS_DIRECTORY_ADDRESS" \
    --private-key $PRIVATE_KEY \
    -r $RPC_URL \
    --verify \
    src/AVS.sol:AVS
```

**Check avs status**

- check operators

`cast call $AVS_ADDRESS   "operators()(address[])" --rpc-url $RPC_URL`


- check strategyParams

`cast call $AVS_ADDRESS   "strategyParams()((address,uint96)[])" --rpc-url $RPC_URL`


