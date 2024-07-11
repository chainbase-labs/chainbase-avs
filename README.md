
## Chainbase avs 

AVS Core Contracts and Registration-related CLI Tools

Before registering as an AVS, ensure that the operator has already registered with EigenLayer. For reference, see: https://docs.eigenlayer.xyz/eigenlayer/operator-guides/operator-installation




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


TODO: 
- offical metadata