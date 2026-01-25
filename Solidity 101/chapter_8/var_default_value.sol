// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;


enum ActionSet {
    add,
    sub
}

contract VarDefaultValue {
    //值类型初始值
    bool public flag_;
    string public str_;
    int public num_;
    uint public num2_;
    ActionSet public action_;
    address public addr_;

    function f() external {
    }
    //引用类型初始值
    uint[8] public static_array_;
    uint[] public dynamic_array_ = new uint[](10);
    mapping(uint => uint) public map_;
    struct Student {
        uint id;
        uint score;
        string name;
    }
    Student public stu_;
    // delete操作符
    bool public flag2_ = true;
    function deleteVar() external {
        delete flag2_;
    }
}