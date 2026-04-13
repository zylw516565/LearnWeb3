// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test} from "forge-std/Test.sol";
import {Counter} from "../src/Counter.sol";

import {console2} from "forge-std/console2.sol";

contract CounterTest is Test {
    Counter public counter;

    function setUp() public {
        counter = new Counter();
        counter.setNumber(0);
    }

    function test_Increment() public {
        counter.increment();
        console2.logUint(counter.number());
        assertEq(counter.number(), 1);
    }

    function testFuzz_SetNumber(uint256 x) public {
        counter.setNumber(x);
        assertEq(counter.number(), x);
    }

    function testSetNumberNonZero(uint256 x) public {
        vm.assume(x != 0);

        counter.setNumber(x);
        assertEq(counter.number(), x);
    }

    function testSetNumberInRange(uint256 x) public {
        uint256 value = bound(x, 1, 100);

        counter.setNumber(value);
        assertEq(counter.number(), value);
        assertTrue(counter.number()>=1 && counter.number() <=100);
    }
}
