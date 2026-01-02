// SPDX-License-Identifier: MIT
pragma solidity ^0.8.31;
struct SenderInfo {
    address addr;
    uint    amount;
}

interface IBank {
    function withdraw() external;
    function getMyBalance() external view returns (uint256);
    function getTopDepositors() external view returns (SenderInfo[3] memory);
}

contract Bank is IBank {
    mapping(address => uint256) internal balances_;
    //管理员地址
    address payable internal owner_;


    SenderInfo[3] private _topDepositors;

    constructor(){
        //在合约部署时，将部署者设置为管理员
        owner_ = payable(msg.sender);
    }

    // 定义事件
    event Received(address sender, uint value);

    receive() external payable virtual  {
        _handleDeposit();
    }

    event fallbackCalled(address sender, uint value, bytes data);

    fallback() external payable {
        _handleDeposit();
    }

    function _handleDeposit() internal {
        emit Received(msg.sender, msg.value);
        balances_[msg.sender] += msg.value;

        SenderInfo memory info;
        info.addr = msg.sender;
        info.amount = balances_[msg.sender];

        _updateTopDepositors(info.addr, info.amount);
    }

    modifier onlyOwner() {
        require(owner_ == msg.sender, "Caller is not owner");
        _;
    }

    function withdraw() public onlyOwner {
        // 获取合约余额
        uint balance = address(this).balance;
        //检查合约中余额
        require(balance > 0, "Insufficient balance");

        (bool success, )= payable(msg.sender).call{value:balance}("");
        require(success, "Transfer failed");
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
    function _updateTopDepositors(address user, uint256 newBalance) internal {
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

contract BigBank is Bank {
    modifier atLeastEth() {
        require(msg.value > 0.001 ether, "At Least 0.001 Eth");
        _;
    }

    function Deposit() external payable atLeastEth {
        _handleDeposit();
    }

    receive() external payable override {
        require(msg.value > 0.001 ether, "At Least 0.001 Eth");
        _handleDeposit();
    }

    function changeAdmin(address payable newAdmin) external onlyOwner {
        require(newAdmin != address(0), "New admin cannot be zero address");
        owner_ = newAdmin;
    }
}

contract Admin {
    //管理员地址
    address private immutable admin_;
    constructor() {
        admin_ = msg.sender;
    }

    modifier onlyOwner {
        require(msg.sender == admin_, "Caller is not owner");
        _;
    }

    function adminWithdraw(IBank bank) external  payable onlyOwner {
        bank.withdraw();
    }

    function withdrawToOwne()external payable onlyOwner {
        // 获取合约余额
        uint balance = address(this).balance;
        //检查合约中余额
        require(balance > 0, "Insufficient balance");

        (bool success, )= payable(admin_).call{value:balance}("");
        require(success, "Transfer failed");
    }
}