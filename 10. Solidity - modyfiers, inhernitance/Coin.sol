// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

/*
Custom bank/coin for multiple users with options:
- deposit
- transfer
- read balance of a user
*/
contract Coin {
    
    // mapping the balance of each address (user)
    mapping(address => uint) public balances;
    
    // adding the value send by the user to his balance
    function deposit() public payable {
        balances[msg.sender] += msg.value;
    }
    
    // function to transfer a certain amount from one address to other address balance
    function transfer(address to, uint amount) public {
        // Replace this if statement by a modyfier
        if(balances[msg.sender] >= amount){
            balances[msg.sender] -= amount;
            balances[to] += amount;
        }
    }

    // return the balance of the user
    function balanceOf(address user) view public returns (uint){
        return balances[user];
    }
}
