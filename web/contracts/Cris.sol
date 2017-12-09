pragma solidity ^0.4.4;

contract Cris {
    string public version = '0.1.0';

    address public admin;

    // mapping (string => address) fileOwner;
    mapping (string => bool) hasFile;

    function Cris() public {
        admin = msg.sender;
    }

    function store(string fileHash) {
      // fileOwner[hash] = msg.sender;
      hasFile[fileHash] = true;
    }

    // function isOwner(address sender, string hash) public returns (bool) {
    //     if (sender == fileOwner[hash]) {
    //         return true;
    //     }
    //     return false;
    // }

    function has(string fileHash) constant returns (bool retVal) {
        return hasFile[fileHash];
    }
}
