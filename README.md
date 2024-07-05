
 **git submodule**

`git submodule update --init --recursive`

`git submodule update --force --recursive --remote`


**deploy**

`forge script script/AVS.s.sol:DeployAVS --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast -vvvv`

**check avs status**

- check operators

`cast call $AVS_ADDRESS   "operators()(address[])" --rpc-url $RPC_URL`


- check strategyParams

`cast call $AVS_ADDRESS   "strategyParams()((address,uint96)[])" --rpc-url $RPC_URL`



**release**

`goreleaser release --snapshot --rm-dist`