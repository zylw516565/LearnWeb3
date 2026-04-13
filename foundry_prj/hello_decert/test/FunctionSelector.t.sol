// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test} from "forge-std/Test.sol";
import {console2} from "forge-std/console2.sol";
import {SelectorExample} from "../src/FunctionSelector.sol";

contract SelectorTest is Test {
  SelectorExample example;
  function setUp() public {

  }

  function test_getTransferSelector() public {
    bytes4 selector = bytes4(keccak256("transfer(address,uint256)"));
    console2.logBytes4(selector);
  }

  function test_getSetSelector() public {
    bytes4 selector = bytes4(keccak256("set(uint256)"));
    console2.logBytes4(selector);
    // 返回: 0x60fe47b1
  }

  function test_UseSelectorProperty() public {
    console2.logBytes4(example.transfer.selector);
  }

}