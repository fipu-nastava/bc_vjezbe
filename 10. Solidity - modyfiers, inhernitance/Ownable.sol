pragma solidity ^0.5.0;

import 'Payable.sol';

contract Ownable is Payable {

    // Owner može primati novac
    address payable owner;

    constructor() public {
        owner = msg.sender;
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
