
## Chainbase avs 

AVS Core Contracts and Registration-related CLI Tools

Before registering as an AVS, ensure that the operator has already registered with EigenLayer. For reference, see: https://docs.eigenlayer.xyz/eigenlayer/operator-guides/operator-installation




## deployment
 **git submodule**

`git submodule update --init --recursive`


**deploy contract**

`forge script script/AVS.s.sol:DeployAVS --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast -vvvv`

**check avs status**

- check operators

`cast call $AVS_ADDRESS   "operators()(address[])" --rpc-url $RPC_URL`


- check strategyParams

`cast call $AVS_ADDRESS   "strategyParams()((address,uint96)[])" --rpc-url $RPC_URL`



**local release**

`goreleaser release --snapshot --rm-dist`