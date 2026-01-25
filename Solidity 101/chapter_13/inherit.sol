// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract Yeye {
    event Log(string msg);
    
    function hip() public virtual  {
        emit Log("Yeye");
    }

    function pop() public virtual  {
        emit Log("Yeye");
    }

    function yeye() public {
        emit Log("Yeye");        
    }
}

contract Baba is Yeye{
    function hip() public virtual  override  {
        emit Log("Baba");
    }

    function pop() public virtual  override {
        emit Log("Baba");
    }

    function baba() public {
        emit Log("Baba");        
    }

    //调用父合约的函数
    function callParent()public {
        Yeye.hip();
        super.hip();
    }
}

// contract Erzi is Baba, Yeye{  //继承必须父辈在前
contract Erzi is Yeye, Baba {
    function hip() public override(Yeye, Baba)  {
        emit Log("Erzi");
    }

    function pop() public virtual override(Yeye, Baba) {
        emit Log("Erzi");
    }

    //调用父合约的函数
    function ErziCallParent()public {
        Yeye.hip();
        super.hip();
    }
}

contract Jiuye {
}
contract Person is Yeye,Jiuye {  //多重继承时,无继承关系的父类,书写顺序不敏感
}

//修饰器的继承
contract Base1 {
    modifier exactDividedBy2And3(uint _a) virtual {
        require(_a % 2 == 0 && _a % 3 == 0);
        _;
    }
}

contract Identifier is Base1 {
    modifier exactDividedBy2And3(uint _a) override  {
        require(_a % 2 == 0 && _a % 3 == 0);
        _;
    }

    //计算一个数分别被2除和被3除的值，但是传入的参数必须是2和3的倍数
    function getExactDividedBy2And3(uint _dividend) public exactDividedBy2And3(_dividend) pure returns(uint, uint) {
        return getExactDividedBy2And3WithoutModifier(_dividend);
    }

    //计算一个数分别被2除和被3除的值
    function getExactDividedBy2And3WithoutModifier(uint _dividend) public pure returns(uint, uint){
        uint div2 = _dividend / 2;
        uint div3 = _dividend / 3;
        return (div2, div3);
    }
}

//构造函数的继承
// contract A {
abstract contract A {
    uint public a;
    constructor(uint a_) {
        a = a_;
    }
}

contract B is A(1) {

}

contract C is A {
    constructor(uint c)A(c*c) {
    }
}

//菱形继承
/* 继承树：
  God
 /  \
Adam Eve
 \  /
people
*/

contract God {
    event Log(string message);

    function foo() public virtual {
        emit Log("God.foo called");
    }

    function bar() public virtual {
        emit Log("God.bar called");
    }
}

contract Adam is God {
    function foo() public virtual override {
        emit Log("Adam.foo called");
        super.foo();
    }

    function bar() public virtual override {
        emit Log("Adam.bar called");
        super.bar();
    }
}

contract Eve is God {
    function foo() public virtual override {
        emit Log("Eve.foo called");
        super.foo();
    }

    function bar() public virtual override {
        emit Log("Eve.bar called");
        super.bar();
    }
}

contract people is Adam, Eve {
    function foo() public override(Adam, Eve) {
        super.foo();
    }

    function bar() public override(Adam, Eve) {
        super.bar();
    }
}