// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

event Transfer2(address indexed from, address indexed to, uint256 value);

contract EventDemo {
    mapping(address => uint256) _balances;
    event Transfer(address indexed from, address indexed to, uint256 value);

    function Transfer_(address from, address to, uint256 amount)external {
        _balances[from] = 10000000;
        _balances[from] -= amount;
        _balances[to]   += amount;

        emit Transfer(from, to, amount);
    }
}