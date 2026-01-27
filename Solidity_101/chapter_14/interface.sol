// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

//抽象合约
// contract InsertionSort {  //含有未实现函数的合约,必须是抽象的(abstract)
abstract contract InsertionSort {
    function insertionSort(uint[] memory a) public pure virtual returns(uint[] memory);
}

abstract contract Base {
    function getAlias() public pure virtual returns(string memory);
}

contract BaseImpl is Base {
    function getAlias() public pure override  returns(string memory){
        return ("BaseImpl");
    }
}

//接口
/*
接口类似于抽象合约，但它不实现任何功能。接口的规则：
不能包含状态变量
不能包含构造函数
不能继承除接口外的其他合约
所有函数都必须是external且不能有函数体
继承接口的非抽象合约必须实现接口定义的所有功能
*/
contract Demo {

}

interface IBase {
}
// interface IName is Demo {  //不能继承除接口外的其他合约
interface IName is IBase {
    // uint id;  //不能包含状态变量
    // constructor(); 不能包含构造函数
    // function getName() public ; 所有函数都必须是external且不能有函数体
    function getName() external returns (string memory);
    function getId() external returns (uint);
}

//非抽象合约必须实现接口中的所有函数声明
// contract Name is  IName{
// }

contract Name is  IName{
    function getName() external returns (string memory) {
        return ("bob");
    }
    function getId() external returns (uint) {
        return (2);
    }
}



















