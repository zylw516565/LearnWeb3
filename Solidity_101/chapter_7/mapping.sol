// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract MappingType {
    mapping(uint => address) public  idToAddress;
    mapping(address => address) public pairAddr;
    mapping(uint => uint) idToId;

    // 我们定义一个结构体 Struct
    struct Student{
        uint256 id;
        uint256 score; 
    }
    //规则1：映射的_KeyType只能选择Solidity内置的值类型
    // mapping(Student => uint) test;  //会报错, struct类型不能作为mapping的key
    mapping(MappingType => uint) test2;
    // 规则2：映射的存储位置必须是storage;不能用于public函数的参数或返回结果中，因为mapping记录的是一种关系 (key - value pair)。
    function demo_test() external {
        // mapping(uint => uint) memory idToId; //报错, 映射的存储位置必须是storage
    }

    //不能用于public函数的参数或返回结果中
    // function demo_test2(mapping(uint => uint) calldata param) public  {
    //     mapping(uint => uint) storage idToId_ = idToId;
    // }

    //规则3：如果映射声明为public，那么Solidity会自动给你创建一个getter函数，可以通过Key来查询对应的Value。
    mapping(uint => uint) public pubIdToId;
    function initpubIdToId(uint key) external {
        pubIdToId[key] = key;
    }
    // 规则4：给映射新增的键值对的语法为_Var[_Key] = _Value
    function writeMap(uint key, uint value) external {
        pubIdToId[key] = value;
    }
    // function getkeccak256(uint key) external returns (bytes32) {
    //     return pubIdToId.keccak256(h(key) . slot);
    // }
}