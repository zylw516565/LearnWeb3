// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

//全局常量?
uint256 constant gState = 0;

contract Demo {
    function fCallData(uint[] calldata x_) public pure returns (uint[] calldata) {
        // x_[0] = 0;  这样修改会报错
        return x_;
    }

    //storage（合约的状态变量）赋值给本地storage（函数里的）时候，会创建引用，改变新变量会影响原变量。例子：
    uint[] x = [1,2,3]; // 状态变量: 数组x
    function fStorage() public {
        uint[] storage xStorage = x;
        xStorage[0] = 100;
    }

    //storage赋值给 memory
    function fMemory() public view {
        uint[] memory xMemory = x;
        xMemory[0] = 100;
        xMemory[1] = 100;
        uint[] memory xMemory2 = x;
        xMemory2[0] = 300;
    }

    //memory赋值给 memory
    function fMemory2Memory() public pure {
        // uint[] memory y = [1,2,3];
        uint[3] memory y = [uint(1),2,3];
        uint[3] memory xMemory = y;
        xMemory[0] = 100;
    }

}

contract Variables {
    //状态变量
    uint public x = 1;
    uint public y;
    string public z;

    function foo() external{
        // 可以在函数里更改状态变量的值
        x = 5;
        y = 2;
        z = "0xAA";
    }

    //局部变量
    function bar() public pure returns(uint) {
        uint xx = 1;
        uint yy = 2;
        uint zz = xx + yy;
        return zz;
    }

    //全局变量
    function global() external view  returns (address, uint, bytes memory) {
        address sender    = msg.sender;
        uint number       = block.number;
        bytes memory data = msg.data;

        return (sender, number, data);
    }
}

contract Unit {
    //4. 全局变量-以太单位与时间单位
    function weiUnit() external pure returns (uint) {
        assert(1 wei == 1e0);
        assert(1 wei == 1);
        return 1 wei;
    }

    function gweiUint() external pure returns (uint) {
        assert(1 gwei == 1e9);
        assert(1 gwei == 1000000000);
        return 1 gwei;
    }

    function etherUint() external pure returns (uint) {
        assert(1 ether == 1e18);
        assert(1 ether == 1000000000000000000);
        return 1 ether;
    }

    //时间单位
    function secondsUint() external pure returns(uint) {
        assert(1 seconds == 1);
        return 1 seconds;
    }

    function minutesUnit() external pure returns (uint) {
        assert(1 minutes == 60);
        assert(1 minutes == 60 seconds);
        return 1 minutes;        
    }

    function hoursUnit() external pure returns (uint) {
        assert(1 hours == 60 * 60);
        assert(1 hours == 60 minutes);
        return 1 hours;
    }

    function daysUnit() external pure returns (uint) {
        assert(1 days == 24 * 60 * 60);
        assert(1 days == 24 * 60 minutes);
        assert(1 days == 24 hours);
        return 1 days;
    }

    function weeksUnit() external pure returns (uint) {
        assert(1 weeks == 7 * 24 * 60 * 60);
        assert(1 weeks == 7 * 24 * 60 minutes);
        assert(1 weeks == 7 * 24 hours);
        assert(1 weeks == 7 days);
        return 1 weeks;
    }
}
