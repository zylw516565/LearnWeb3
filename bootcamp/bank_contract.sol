// SPDX-License-Identifier: MIT
pragma solidity ^0.8.31;


contract Bank {
    mapping(address => uint256) private balances_;
    //管理员地址
    address private immutable owner_;

    struct SenderInfo {
        address addr;
        uint    amount;
    }
    SenderInfo[3] private _topDepositors;

    constructor(){
        //在合约部署时，将部署者设置为管理员
        owner_ = msg.sender;
    }

    // 定义事件
    event Received(address sender, uint value);

    receive() external payable {
        emit Received(msg.sender, msg.value);
        balances_[msg.sender] += msg.value;

        SenderInfo memory info;
        info.addr = msg.sender;
        info.amount = balances_[msg.sender];

        updateTopUsers(info);
    }

    event fallbackCalled(address sender, uint value, bytes data);

    fallback() external payable {
        emit fallbackCalled(msg.sender, msg.value, msg.data);
        balances_[msg.sender] += msg.value;

        SenderInfo memory info;
        info.addr = msg.sender;
        info.amount = balances_[msg.sender];
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

        for (uint i = 0; i < _topDepositors.length; i++) {
            // 每当比数组中的用户余额大,就跟数组中的信息交换. 起到每次接收ETH都维护_topDepositors效果
            if(info.amount > _topDepositors[i].amount) {
                tmpinfo.addr = _topDepositors[i].addr;
                tmpinfo.amount = _topDepositors[i].amount;

                _topDepositors[i].addr = info.addr;
                _topDepositors[i].amount = info.amount;

                info.addr = tmpinfo.addr;
                info.amount = tmpinfo.amount;
            }
        }
    }

    function getMyBalance() external view returns (uint256) {
        return balances_[msg.sender];
    }

    function getTopDepositors() external view returns (SenderInfo[3] memory) {
        return _topDepositors;
    }

    /**
     * @dev 内部函数，用于更新存款排行榜。
     * @param user 存款用户的地址。
     * @param newBalance 用户的最新存款余额。
     */
    function _updateTopDepositors(address user, uint256 newBalance) private {
        // 检查用户是否已在排行榜中
        for (uint256 i = 0; i < _topDepositors.length; i++) {
            if (_topDepositors[i].addr == user) {
                // 如果已在榜中，更新其金额
                _topDepositors[i].amount = newBalance;
                // 对排行榜进行排序
                _sortTopDepositors();
                return;
            }
        }

        // 如果用户不在榜中，与当前第三名比较
        if (newBalance > _topDepositors[2].amount) {
            // 如果新金额更大，则替换第三名
            _topDepositors[2] = SenderInfo({addr: user, amount: newBalance});
            _sortTopDepositors();
        }
    }

    /**
     * @dev 内部函数，用于对排行榜进行降序排序。
     */
    function _sortTopDepositors() private {
        // 简单的冒泡排序
        for (uint256 i = 0; i < _topDepositors.length; i++) {
            for (uint256 j = i + 1; j < _topDepositors.length; j++) {
                if (_topDepositors[i].amount < _topDepositors[j].amount) {
                    // 交换位置
                    SenderInfo memory tmpDepositor = _topDepositors[i];
                    _topDepositors[i] = _topDepositors[j];
                    _topDepositors[j] = tmpDepositor;
                }
            }
        }
    }
}