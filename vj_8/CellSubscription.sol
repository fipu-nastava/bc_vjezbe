pragma solidity >=0.4.22 <0.7.0;

/*
The subscriber pays a monthly fee in ether to the contract,
and the cell phone company can check if the account is paid in full
*/
contract CellSubscription {
	uint256 monthlyCost;

    constructor(uint256 cost) public {
        monthlyCost = cost;
    }

    // function that allows the subscriber to make a payment towards their account
    function makePayment() payable public {

		}

		// function that allows the subscriber to make a payment towards their account
		// without calling the makePayment function, rather just by simply transfering the funds
		function () external payable {}

    // allows an account to be emptied to the caller of widhdrawBalance
    function withdrawBalance() public {
        msg.sender.transfer(address(this).balance);
    }
    // functionality that allows the cell phone company to check the status of the account on a given date
    function isBalanceCurrent(uint256 monthsElapsed) public view returns (bool) {
        return monthlyCost * monthsElapsed == address(this).balance;
    }
}
