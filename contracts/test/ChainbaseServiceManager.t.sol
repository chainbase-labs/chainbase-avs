// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import "@eigenlayer/test/mocks/EmptyContract.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@eigenlayer-middleware/src/RegistryCoordinator.sol";

import "./MockAVSDeployer.sol";

contract ChainbaseServiceManagerTest is MockAVSDeployer {
    function setUp() public {
        _setUpMockAVSDeployer();
    }

    function testAddOperatorsToWhitelist() public {
        address[] memory operators = new address[](2);
        operators[0] = address(0x1);
        operators[1] = address(0x2);

        vm.prank(address(0x1));
        vm.expectRevert("Ownable: caller is not the owner");
        chainbaseServiceManagerProxy.addOperatorsToWhitelist(operators);

        chainbaseServiceManagerProxy.addOperatorsToWhitelist(operators);

        assertTrue(chainbaseServiceManagerProxy.operatorWhitelist(address(0x1)));
        assertTrue(chainbaseServiceManagerProxy.operatorWhitelist(address(0x2)));
        assertFalse(chainbaseServiceManagerProxy.operatorWhitelist(address(0x3)));
    }

    function testRemoveOperatorsFromWhitelist() public {
        address[] memory operators = new address[](3);
        operators[0] = address(0x1);
        operators[1] = address(0x2);
        operators[2] = address(0x3);

        chainbaseServiceManagerProxy.addOperatorsToWhitelist(operators);

        address[] memory removedOperators = new address[](2);
        removedOperators[0] = address(0x1);
        removedOperators[1] = address(0x2);

        vm.prank(address(0x1));
        vm.expectRevert("Ownable: caller is not the owner");
        chainbaseServiceManagerProxy.removeOperatorsFromWhitelist(removedOperators);

        chainbaseServiceManagerProxy.removeOperatorsFromWhitelist(removedOperators);

        assertFalse(chainbaseServiceManagerProxy.operatorWhitelist(address(0x1)));
        assertFalse(chainbaseServiceManagerProxy.operatorWhitelist(address(0x2)));
        assertTrue(chainbaseServiceManagerProxy.operatorWhitelist(address(0x3)));
    }

    function testSetWhitelistEnabled() public {
        vm.prank(address(0x1));
        vm.expectRevert("Ownable: caller is not the owner");
        chainbaseServiceManagerProxy.setWhitelistEnabled(true);

        chainbaseServiceManagerProxy.setWhitelistEnabled(true);
        assertTrue(chainbaseServiceManagerProxy.whitelistEnabled());

        vm.prank(address(0x1));
        vm.expectRevert("Ownable: caller is not the owner");
        chainbaseServiceManagerProxy.setWhitelistEnabled(false);

        chainbaseServiceManagerProxy.setWhitelistEnabled(false);
        assertFalse(chainbaseServiceManagerProxy.whitelistEnabled());
    }

    function testRegisterOperatorToAVS() public {
        address operator = address(0x1);
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature;

        // test onlyRegistryCoordinator
        vm.expectRevert("ServiceManagerBase.onlyRegistryCoordinator: caller is not the registry coordinator");
        chainbaseServiceManagerProxy.registerOperatorToAVS(operator, operatorSignature);

        // test onlyWhitelisted modifier
        chainbaseServiceManagerProxy.setWhitelistEnabled(false);
        vm.prank(address(registryCoordinatorProxy));
        vm.expectRevert();
        chainbaseServiceManagerProxy.registerOperatorToAVS(operator, operatorSignature);

        chainbaseServiceManagerProxy.setWhitelistEnabled(true);

        vm.prank(address(registryCoordinatorProxy));
        vm.expectRevert("ChainbaseServiceManager: operator not in whitelist");
        chainbaseServiceManagerProxy.registerOperatorToAVS(operator, operatorSignature);

        chainbaseServiceManagerProxy.addOperatorsToWhitelist(toArray(operator));
        vm.prank(address(registryCoordinatorProxy));
        vm.expectRevert();
        chainbaseServiceManagerProxy.registerOperatorToAVS(operator, operatorSignature);
    }

    function testDeregisterOperatorFromAVS() public {
        address operator = address(0x1);

        // test onlyRegistryCoordinator
        vm.expectRevert("ServiceManagerBase.onlyRegistryCoordinator: caller is not the registry coordinator");
        chainbaseServiceManagerProxy.deregisterOperatorFromAVS(operator);

        // test onlyWhitelisted modifier
        chainbaseServiceManagerProxy.setWhitelistEnabled(false);
        vm.prank(address(registryCoordinatorProxy));
        vm.expectRevert();
        chainbaseServiceManagerProxy.deregisterOperatorFromAVS(operator);

        chainbaseServiceManagerProxy.setWhitelistEnabled(true);

        vm.prank(address(registryCoordinatorProxy));
        vm.expectRevert("ChainbaseServiceManager: operator not in whitelist");
        chainbaseServiceManagerProxy.deregisterOperatorFromAVS(operator);

        chainbaseServiceManagerProxy.addOperatorsToWhitelist(toArray(operator));
        vm.prank(address(registryCoordinatorProxy));
        vm.expectRevert();
        chainbaseServiceManagerProxy.deregisterOperatorFromAVS(operator);
    }

    function toArray(address addr) internal pure returns (address[] memory) {
        address[] memory arr = new address[](1);
        arr[0] = addr;
        return arr;
    }
}
