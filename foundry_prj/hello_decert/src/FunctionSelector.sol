// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

contract SelectorExample {
  // 函数签名格式：functionName(paramType1,paramType2,...)
  // 注意：不包含参数名，不包含空格，不包含返回值类型

  function transfer(address to, uint256 amount) public returns (bool) {
      // transfer 的选择器是 0xa9059cbb
      // 计算方式：bytes4(keccak256("transfer(address,uint256)"))
      return true;
  }

  // 手动计算函数选择器
  function getTransferSelector() public pure returns (bytes4) {
    bytes4 selector = bytes4(keccak256("transfer(address,uint256)"));
    return selector;
  }
}

contract SelectorProperty {
  function transfer(address to, uint256 amount) public returns (bool) {
      return true;
  }

  function getSelector() public pure returns (bytes4) {
      // 使用 .selector 属性直接获取
      return this.transfer.selector;
      // 返回: 0xa9059cbb
  }
}