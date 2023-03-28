// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "./Ownable.sol";

contract Destructible is Ownable {

  // Samo owner može pozvati brisanje ugovora
  function destroy() onlyOwner public {
    selfdestruct(owner);
  }

  // Moguče je uništiti ugovor i prenesti novac s ugovora na neki drugi račun
  function destroyAndSend(address payable _recipient) onlyOwner public {
    selfdestruct(_recipient);
  }
}