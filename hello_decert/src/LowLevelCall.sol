// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

interface IERC20 {
  function transfer(address to, uint256 amount)external returns (bool);
}

contract LowLevelCall {
  //方法1: 使用encodeWithSignature方法
  function callTransfer1(address token, address to, uint256 amount) public {
    bytes memory data = abi.encodeWithSignature("transfer(address,uint256)", to, amount);
    (bool success, ) = token.call(data);
    require(success, "Transfer failed");
  }

  //方法2: 使用encodeWithSignature方法
  function callTransfer2(address token, address to, uint256 amount) public {
    bytes memory data = abi.encodeWithSelector(bytes4(keccak256("transfer(address,uint256)")), to, amount);
    (bool success, ) = token.call(data);
    require(success, "Transfer failed");
  }

  //方法3: 使用接口的 selector 属性
  function callTransfer3(address token, address to, uint256 amount) public {
    bytes memory data = abi.encodeWithSelector(IERC20.transfer.selector, to, amount);
    (bool success, ) = token.call(data);
    require(success, "Transfer failed");
  }
}