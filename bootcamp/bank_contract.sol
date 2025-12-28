// SPDX-License-Identifier: MIT
pragma solidity ^0.8.31;


contract Bank {
    mapping(address => uint) public leger;
    //管理员地址
    address private immutable owner_;

    struct SenderInfo {
        address addr;
        uint    amount;
    }
    SenderInfo[3] public topUsers_;

    // address[3] private topUsers;

    constructor(){
        //在合约部署时，将部署者设置为管理员
        owner_ = msg.sender;
    }

    // 定义事件
    event Received(address sender, uint value);

    receive() external payable {
        emit Received(msg.sender, msg.value);
        leger[msg.sender] += msg.value;

        SenderInfo memory info;
        info.addr = msg.sender;
        info.amount = leger[msg.sender];

        updateTopUsers(info);
    }

    event fallbackCalled(address sender, uint value, bytes data);

    fallback() external payable {
        emit fallbackCalled(msg.sender, msg.value, msg.data);
        leger[msg.sender] += msg.value;

        SenderInfo memory info;
        info.addr = msg.sender;
        info.amount = leger[msg.sender];
        updateTopUsers(info);
    }

    modifier onlyOwner() {
        require(owner_ == msg.sender, "Caller is not owner");
        _;
    }

    function withdraw(uint withdrawAmount) public onlyOwner {
        //检查合约中余额
        require(address(this).balance >= withdrawAmount, "Insufficient balance");

        (bool success, )= payable(msg.sender).call{value:withdrawAmount}("");
        require(success, "Transfer failed");
    }

    function updateTopUsers(SenderInfo memory info) private {
        SenderInfo memory tmpinfo;

        for (uint i = 0; i < topUsers_.length; i++) {
            // 每当比数组中的用户余额大,就跟数组中的信息交换. 起到每次接收ETH都维护topUsers_效果
            if(info.amount > topUsers_[i].amount) {
                tmpinfo.addr = topUsers_[i].addr;
                tmpinfo.amount = topUsers_[i].amount;

                topUsers_[i].addr = info.addr;
                topUsers_[i].amount = info.amount;

                info.addr = tmpinfo.addr;
                info.amount = tmpinfo.amount;
            }
        }
    }
}