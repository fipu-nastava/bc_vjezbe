// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;


contract Ownable {

    // Owner može primati novac
    address payable owner;

    constructor() {
        owner = payable(msg.sender);
    }

    // Kontrola prava ograničena samo na owner-a
    modifier onlyOwner {
        require(msg.sender == owner, "Nisi vlasnik.");
        _;
    }

    function get_owner() public view returns(address) {
        return owner;
    }

}