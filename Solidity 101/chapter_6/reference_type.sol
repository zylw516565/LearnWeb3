// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract ReferenceType {
    // 固定长度 Array
    uint[8] array1;
    bytes1[5] array2;
    address[6] array3;

    //可变长度数组
    uint[] array4;
    bytes1[] array5;
    address[] array6;
    bytes array7;
    bytes[] array8;

    //memory动态数组
    function test() external pure returns (uint[] memory, address[] memory) {
        uint[] memory array9 = new uint[](5);
        address[] memory array10 = new address[](5);
        return (array9, array10);
    }

    function f()public pure {
        g([uint(1),2,3]);
        // g([1,2,3]); //报错
    }
    function g(uint[3] memory data) public pure {
    }

    function test2() external pure {
        uint[] memory array10 = new uint[](10);
        array10[0] = 1;
        array10[1] = 2;
        array10[2] = 3;
    }

    function arrayPush() public returns (uint[] memory) {
        uint[2] memory a = [uint(1), 2];
        array4 = a;
        array4.push(3);
        array4.push();
        array4.pop();
        uint array_len = array4.length;
        uint array_len2 = array1.length;
        return array4;
    }

    struct Student{
        uint256 id;
        uint256 score;
    }
    Student student;
    //  给结构体赋值
    // 方法1:在函数中创建一个storage的struct引用
    function struct_demo()external  {
        Student storage student_ = student;
        student_.id = 1;
        student_.score = 100;
    }
    // 方法2:直接引用状态变量的struct
    function struct_demo2()external {
        student.id = 2;
        student.score = 200;
    }
    // 方法3:构造函数式
    function struct_demo3()external {
        student = Student(3,9);
    }
    // 方法4:key value
    function struct_demo4()external {
    student = Student({id:3, score: 300});
    }
    //结构体可以包含应用类型, 例如包含数组, 结构体, 映射
    struct DemoStruct{
        uint[] array;
        Student student;
        mapping(uint=>uint) map_;
    }
    //同时结构体也可以作为数组,映射的成员
    Student[] students;
    mapping(uint=>Student) studentMap;
}