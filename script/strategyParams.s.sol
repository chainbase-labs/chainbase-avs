pragma solidity ^0.8.12;

import {IStrategy} from "@eigenlayer/contracts/interfaces/IStrategy.sol";
import {IAVS} from "../src/interface/IAVS.sol";
import {EigenMainnetDeployments, EigenHoleSkyDeployments} from "./EigenDeployments.s.sol";

library StrategyParams {
    uint96 public constant STD_MULTIPLIER = 1e18;

    function holesky() internal pure returns (IAVS.StrategyParam[] memory params) {
        params = new IAVS.StrategyParam[](5);

        params[0] = _stdStrategyParam(EigenHoleSkyDeployments.cbETHStrategy);
        params[1] = _stdStrategyParam(EigenHoleSkyDeployments.stETHStrategy);
        params[2] = _stdStrategyParam(EigenHoleSkyDeployments.rETHStrategy);
        params[3] = _stdStrategyParam(EigenHoleSkyDeployments.ETHxStrategy);
        params[4] = _stdStrategyParam(EigenHoleSkyDeployments.beaconETHStrategy);
    }

    function mainnet() internal pure returns (IAVS.StrategyParam[] memory params) {
        params = new IAVS.StrategyParam[](5);

        params[0] = _stdStrategyParam(EigenMainnetDeployments.cbETHStrategy);
        params[1] = _stdStrategyParam(EigenMainnetDeployments.stETHStrategy);
        params[2] = _stdStrategyParam(EigenMainnetDeployments.rETHStrategy);
        params[3] = _stdStrategyParam(EigenMainnetDeployments.wBETHStrategy);
        params[4] = _stdStrategyParam(EigenMainnetDeployments.beaconETHStrategy);
    }

    function _stdStrategyParam(address strategy) internal pure returns (IAVS.StrategyParam memory) {
        return IAVS.StrategyParam({strategy: IStrategy(strategy), multiplier: STD_MULTIPLIER});
    }
}
