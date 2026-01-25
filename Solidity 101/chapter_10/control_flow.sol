// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract ControlFlow {
    function ifElseTest(uint256 num_) public pure returns (bool) {
        if(0 == num_) {
            return (true);
        }else {
            return (false);
        }
    }

    function forLoopTest()public pure returns (uint256){
        uint sum = 0;
        for(uint i=0; i < 10; i++) {
            sum += i;
        }
        return (sum);
    }

    function whileLoopTest() public pure returns(uint256) {
        uint sum = 0;
        uint i = 0;
        while(i < 10) {
            sum += i;
            i++;
        }
        
        return (sum);
    }

    function doWhileTest() public pure returns (uint256) {
        uint sum = 0;
        uint i = 0;

        do {
            sum += i;
            i++;
        } while (i < 10);

        return (sum);
    }

    function ternaryTest(uint256 x, uint256 y) public pure returns (uint256) {
        return x > y ? x : y;
    }

    // 插入排序 错误版
    function insertionSortWrong(uint[] memory nums) public pure returns(uint[] memory) {
        for(uint i=1; i < nums.length; i++) {
            uint value = nums[i];
            uint j = i - 1;
            while(j>=0 && nums[j] > value) {
                nums[j+1] = nums[j];
                j--;
            }
            nums[j+1] = value;
        }

        return nums;
    }

    // 插入排序 
    function insertionSort(uint[] memory nums) public pure returns(uint[] memory) {
        for(uint i=1; i < nums.length; i++) {
            uint value = nums[i];
            uint j = i;
            while(j>=1 && nums[j-1] > value) {
                nums[j] = nums[j-1];
                j--;
            }
            nums[j] = value;
        }

        return nums;
    }
}