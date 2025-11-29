// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

error TransferNotOwner();  //自定义error
// error TransferNotOwner(address sender);
contract Exception {
    mapping(uint256 => address) _owners;
    function transferOwner1(uint256 tokenId, address newOwner) public {
        if(_owners[tokenId] != msg.sender) {
            revert TransferNotOwner();
        }

        _owners[tokenId] = newOwner;
    }

    function transferOwner1_require(uint256 tokenId, address newOwner) public {
        require(_owners[tokenId] == msg.sender, "NotOwner");

        _owners[tokenId] = newOwner;
    }

    function transferOwner1_assert(uint256 tokenId, address newOwner) public {
        assert(_owners[tokenId] == msg.sender);

        _owners[tokenId] = newOwner;
    }
}