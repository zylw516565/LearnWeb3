// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract Example {
    //返回多个变量
    function returnMultiple() public pure returns (uint256, bool, uint256[3] memory) {
        return (1, true, [uint256(1), 2, 5]);
    }

    //命名式返回
    function returnNamed() public pure returns (uint256 number_, bool bool_, uint256[3] memory array_) {
        number_ = 2;
        bool_ = true;
        array_ = [uint256(1), 2, 5];
    }

    // 命名式返回，依然支持return
    function returnNamed2() public pure returns(uint256 _number, bool _bool, uint256[3] memory _array){
        return(1, true, [uint256(1),2,5]);
    }

    function returnNamed3() public pure returns (int num1, int num2, int num3) {
        num1 = 1;
        num2 = 2;
        num3 = 3;
    }

    //解构式赋值
    function destructureAssign() public pure  {
        uint256 _number;
        bool _bool;
        uint256[3] memory _array;
        (_number, _bool, _array) = returnNamed();

        (, _bool, ) = returnNamed();
    }

}