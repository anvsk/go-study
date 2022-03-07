pragma solidity ^0.8.1;

contract Store {
  event ItemSet(bytes32 key, bytes32 value);

  string public version;
  mapping (bytes32 => bytes32) public items;

  constructor(string memory mem) public {
    version = mem;
  }

  function setItem(bytes32 key, bytes32 value) external {
    items[key] = value;
    emit ItemSet(key, value);
  }
}