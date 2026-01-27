// SPDX-License-Identifier: MIT
pragma solidity ^0.8.31;

contract Fallback {
    // 定义事件
    event Received(address addr, uint value);

    // 接收ETH时释放Received事件
    receive() external payable {
        emit Received(msg.sender, msg.value);
    }

    event fallbackCalled(address addr, uint value, bytes data);

    fallback() external payable {
        emit fallbackCalled(msg.sender, msg.value, msg.data);
    }
}