// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test} from "forge-std/Test.sol";
import {console2} from "forge-std/console2.sol";
import {LowLevelCall} from "../src/LowLevelCall.sol";

contract LowLevelCallTest is Test {
  function setUp() public {
  }

  function test_Set() public {
    bytes memory data = abi.encodeWithSignature("set(uint256)", 10);
    console2.logBytes(data);

    uint a = 10;
    bytes memory data2 = abi.encode(a);
    console2.logBytes(data2);

    uint8 s = 2;
    bytes memory data3 = abi.encode(s);
    console2.logBytes(data3);

    address addr = 0xe74c813e3f545122e88A72FB1dF94052F93B808f;
    console2.logBytes(abi.encode(addr));

    // 多个参数
    console2.logBytes(abi.encode(addr, a));

    bool b = true;
    console2.logBytes(abi.encode(addr, a, b));
  }
}