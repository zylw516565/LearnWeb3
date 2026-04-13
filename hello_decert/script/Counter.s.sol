// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script} from "forge-std/Script.sol";
import {Counter} from "../src/Counter.sol";

import {console2} from "forge-std/console2.sol";

contract CounterScript is Script {
    Counter public counter;

    function setUp() public {}

    function run() public {
        string memory mnemonic = vm.envString("MNEMONIC");
        (address deployer, ) = deriveRememberKey(mnemonic, 0);


        vm.startBroadcast(deployer);

        counter = new Counter();
        console2.log("Counter deployed on %s", address(counter));
        vm.stopBroadcast();
    }
}
