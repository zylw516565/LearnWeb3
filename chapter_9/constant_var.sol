// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

// constant变量必须在声明的时候初始化，之后不能改变

contract ConstantVar {
    uint256  public constant num1 = 1;
    string public constant name = "bob";
    bytes public constant data1 = "WTF";
    address public constant addr = address(0);

    uint256 immutable num2;
    // string immutable name2 = "bob";  //string 不能声明成immutable
    // bytes immutable data2;  bytes 不能声明成immutable
    address immutable addr2 = address(0);

    // immutable变量可以在constructor里初始化，之后不能改变
    uint256 public immutable imm_num = 1;
    constructor() {
        imm_num = 2;
        IMMUTABLE_TEST = test();
    }

    function f() external {
        // imm_num = 3;  //immutable变量初始化之后不能改变
    }
    // 在`Solidity v8.0.21`以后,下列变量数值暂为初始值
    address public immutable IMMUTABLE_ADDRESS; 
    uint256 public immutable IMMUTABLE_BLOCK;
    uint public immutable IMMUTABLE_TEST;

    //你可以使用全局变量例如address(this)，block.number 或者自定义的函数给immutable变量初始化。
    address public immutable IMM_ADDR = address(this);
    uint public immutable IMM_BLOCK_NUM = block.number;

    function test() private pure returns (uint) {
        return 9;
    }

    uint public immutable IMM_TEST = test();

}