
## Chainbase avs 

AVS Core Contracts and Registration-related CLI Tools

Before registering as an AVS, ensure that the operator has already registered with EigenLayer. For reference, see: https://docs.eigenlayer.xyz/eigenlayer/operator-guides/operator-installation




## deployment
 **git submodule**

`git submodule update --init --recursive`


### deploy contract

- Deploy AVS

`forge script script/AVS.s.sol:DeployAVS --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast -vvvv`

- Deploy ProxyAdmin

`forge script script/DeplyProxyAdmin.s.sol:DeployProxyAdmin --rpc-url $RPC_URL --private-key $PROXY_ADMIN_DEPLOYER_KEY --broadcast -vvvv`



**Verify & Publish to Etherscan**

`forge verify-contract $AVS_IMPL_ADDRESS  --constructor-args xxx  src/AVS.sol:AVS  --watch -e $ETHERSCAN_API_KEY --rpc-url $RPC_URL`

`forge verify-contract $AVS_PROXY_ADDRESS lib/openzeppelin-contracts/contracts/proxy/transparent/TransparentUpgradeableProxy.sol:TransparentUpgradeableProxy --constructor-args xxx   --watch -e $ETHERSCAN_API_KEY --rpc-url $RPC_URL`

**check avs status**

- check operators

`cast call $AVS_ADDRESS   "operators()(address[])" --rpc-url $RPC_URL`


- check strategyParams

`cast call $AVS_ADDRESS   "strategyParams()((address,uint96)[])" --rpc-url $RPC_URL`



**local release**

`goreleaser release --snapshot --rm-dist`


TODO: 
- offical metadata