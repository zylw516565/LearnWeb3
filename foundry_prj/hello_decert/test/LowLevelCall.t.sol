// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Nox, euint256, externalEuint256} from "@iexec-nox/nox-protocol-contracts/contracts/sdk/Nox.sol"; 

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

  function test_EncodePacked() public pure {
    // 单个参数
    uint a = 10;

    uint256 x = 8;
    console2.log("x= %d", x);

    bytes memory data = abi.encodePacked(a);
    console2.logBytes(data);

    uint8 s = 2; // 占一个字节
    console2.logBytes(abi.encodePacked(s));

    address addr = 0xe74c813e3f545122e88A72FB1dF94052F93B808f;
    console2.logBytes(abi.encodePacked(addr));

    bool b = true;
    // 多个参数
    console2.logBytes(abi.encodePacked(addr, s, b));
  }

  function test_EncodeWithSignature() public pure {
    console2.logBytes(abi.encodeWithSignature("set(uint256)", 10));  //0x60fe47b1000000000000000000000000000000000000000000000000000000000000000a
    uint8 s = 2; // 占一个字节
    bool b = true;
    address addr = 0xe74c813e3f545122e88A72FB1dF94052F93B808f;

    // 参考上方 addr, s, b 的定义
    console2.logBytes(abi.encodeWithSignature("addUser(address,uint8,bool)", addr, s, b));
  }

  function test_EncodeWithSelector() public {
    console2.logBytes(abi.encodeWithSelector(bytes4(keccak256("set(uint256)")), 10));

    uint8 s = 2; // 占一个字节
    bool b = true;
    address addr = 0xe74c813e3f545122e88A72FB1dF94052F93B808f;

    console2.logBytes(abi.encodeWithSelector(0x63f67eb5, addr, s, b));
    // 等价于
    console2.logBytes(abi.encodeWithSelector(bytes4(keccak256("addUser(address,uint8,bool)")), addr, s, b));
  }

  function test_EncodeCall(address _to, uint _value) public {
    console2.logBytes(abi.encodeCall(IERC20.transfer, (_to, _value)));
  }

  function test_Decode() public {
    bytes memory data;
    abi.decode(data, (uint));
  }
}

interface IERC20 {
  function transfer(address recipient, uint amount) external returns (bool);
}