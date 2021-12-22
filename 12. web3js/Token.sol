// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "./Destructible.sol";

contract Token is Destructible {

  mapping(address => uint) balances;
  uint public totalSupply = 128 ether;
  
  constructor () {
      balances[msg.sender] = totalSupply;
  }

  function balanceOf(address _owner) public view returns (uint balance) {
      return balances[_owner];
  }

  function getDonation() public {
      if (balances[msg.sender] == 0) {
          balances[msg.sender] = 20;
          }
  }

}
