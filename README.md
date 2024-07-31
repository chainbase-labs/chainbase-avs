
## Chainbase avs 

AVS Core Contracts and Registration-related CLI Tools

Before registering as an AVS, ensure that the operator has already registered with EigenLayer. For reference, see: https://docs.eigenlayer.xyz/eigenlayer/operator-guides/operator-installation


## holesky contract address

| Name | Proxy | Implementation | Notes |
| -------- | -------- | -------- | -------- |
| [`ProxyAdmin`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/proxy/transparent/ProxyAdmin.sol) |  | [`0xdFbD62c5d8C5739852f67F2D7d2148FC5Bf2ce8E`](https://holesky.etherscan.io/address/0xdfbd62c5d8c5739852f67f2d7d2148fc5bf2ce8e) | onwer:0xB3500b9D97C1F26B92f248CACa6906C02b34409A |
| [`AVS`](https://github.com/chainbase-labs/chainbase-avs-contracts/blob/main/src/AVS.sol) |[`0x5e78eff26480a75e06ccdabe88eb522d4d8e1c9d`](https://holesky.etherscan.io/address/0x5e78eff26480a75e06ccdabe88eb522d4d8e1c9d#code) | [`0x0470364dcec9a1da4a011ac23df6f50d9f6da60f`](https://holesky.etherscan.io/address/0x0470364dcec9a1da4a011ac23df6f50d9f6da60f#code) | Proxy: [`TUP@4.7.1`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.1/contracts/proxy/transparent/TransparentUpgradeableProxy.sol) |

## run task

repalce field in avs.toml.example

notice: `operator.yaml`'s `private_key_store_path` shuold points to the corresponding file path in the container.

- ${OPERATOR_CONFIG_PATH:-./operator.yaml}:/opt/operator.yaml
- ${EIGEN_KEY_PATH:-./eigen-test.ecdsa.key.json}:/opt/eigen-test.ecdsa.key.json

```shell
export OPERATOR_CONFIG_PATH=/path/to/operator.yaml  EIGEN_KEY_PATH=/path/to/ecdsa.key.json

docker-compose up --build -d
```

### node health check

Operator can monitor the process alive by checking api: `GET /eigen/node/health`, http status 200 means OK

## register

local build 
```
go mod tidy 
go build -o avs-cli .

./avs-cli register
```


## deployment
 **git submodule**

`git submodule update --init --recursive`


### deploy contract

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

**check avs status**

- check operators

`cast call $AVS_ADDRESS   "operators()(address[])" --rpc-url $RPC_URL`


- check strategyParams

`cast call $AVS_ADDRESS   "strategyParams()((address,uint96)[])" --rpc-url $RPC_URL`



**local release**

`goreleaser release --snapshot --rm-dist`


