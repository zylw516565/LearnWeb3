// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract BaseERC20 {
  string public name;
  string public symbol;
  uint8  public decimals;

  uint256 public totalSupply;
  mapping(address => uint256) balances;
  mapping(address => mapping(address => uint256)) allowances;

  event Transfer(address indexed from, address indexed to, uint256 value);
  event Approval(address indexed owner, address indexed spender, uint256 value);

  constructor(){
    name   = "MyToken";
    symbol = "MTK";
    decimals = 18;
    totalSupply = 100000000 * 10 ** uint256(decimals);
    balances[msg.sender] = totalSupply;
  }

  function balanceOf(address owner_)public view returns(uint256 balance) {
    return balances[owner_];
  }

  function transfer(address to_, uint256 value)public returns (bool sucess) {
    require(to_ != address(0), "ERC20: transfer to the zero address");
    require(balances[msg.sender] >= value, "ERC20: transfer amount exceeds balance");

    balances[msg.sender] -= value;
    balances[to_]        += value;

    emit Transfer(msg.sender, to_, value);
    return true;
  }

  function transferFrom(address from_, address to_, uint256 value_) public returns (bool success) {
    require(from_ != address(0), "ERC20: transfer from the zero address");
    require(to_ != address(0), "ERC20: transfer to the zero address");
    require(balances[from_] >= value_, "ERC20: transfer amount exceeds balance");
    require(allowances[from_][msg.sender] >= value_, "ERC20: transfer amount exceeds allowance");

    balances[from_] -= value_;
    balances[to_]   += value_;

    allowances[from_][msg.sender] -= value_;
    emit Transfer(from_, to_, value_);
    return true;
  }

  function approve(address _spender, uint256 _value) public returns (bool success) {
    allowances[msg.sender][_spender] = _value;
    emit Approval(msg.sender, _spender, _value);
    return true;
  }

  function allowance(address _owner, address _spender) public view returns (uint256 remaining) {
    return allowances[msg.sender][_spender];
  }
}