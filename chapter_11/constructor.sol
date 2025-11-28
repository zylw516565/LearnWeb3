// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract Construct {
    address public owner;
    constructor(address initOwner) {
        owner = initOwner;
    }

    //修饰器
    modifier onlyOwner {
        require(msg.sender == owner);
        _;
    }

    function changeOwner(address newOwner)external onlyOwner {
        owner = newOwner;
    }

    function getOwner()external onlyOwner view returns (address){
        return (owner);
    }
}