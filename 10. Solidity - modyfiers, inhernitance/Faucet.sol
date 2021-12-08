pragma solidity ^0.5.0;

import 'Ownable.sol';
import 'Destructible.sol';

contract Faucet is Destructible {

    // Pokloni ether svima koji pitaju
    function withdraw(uint withdraw_amount) public {

        // Limit koliko se može zatražiti
        // Baca Exception ako nije OK
        require(withdraw_amount <= 0.1 ether);

        // Slanje iznosa na adresu koja je zatražila
        msg.sender.transfer(withdraw_amount);
    }
}
