pragma solidity ^0.4.4;

contract Cris {
    string public version = '0.1.0';

    address public admin;

    mapping (string => address) fileOwner;

    function Cris() public {
        admin = msg.sender;
    }

    function store(string hash) public {
      fileOwner[hash] = msg.sender;
    }

    function isOwner(string hash) public returns (bool) {
        return msg.sender == fileOwner[hash];
    }
}
