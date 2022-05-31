//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "hardhat/console.sol";

error DoesNotExist();

contract Crud {
  User[] public users;
  uint256 public nextId = 1;
  address admin;

  struct User {
    uint256 id;
    string name;
  }

  constructor() {
    admin = msg.sender;
  }

  modifier adminAccess() {
    require(admin == msg.sender, "No Access Rights");
    _;
  }

  function create(string memory name) public {
    users.push(User(nextId, name));
    nextId++;
  }

  function read(uint256 id) public view returns (uint256, string memory) {
    uint256 i = find(id);
    return (users[i].id, users[i].name);
  }

  function update(uint256 id, string memory newName) public {
    uint256 i = find(id);
    users[i].name = newName;
  }

  function delete_(uint256 id) public adminAccess {
    uint256 i = find(id);
    delete users[i];
  }

  function find(uint256 id) internal view returns (uint256) {
    for (uint256 i = 0; i < users.length; i++) {
      if (users[i].id == id) {
        return i;
      }
    }
    revert DoesNotExist();
  }
}
