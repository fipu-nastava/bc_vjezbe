pragma solidity >=0.4.25 <0.7.0;



/*
A smart contract that track flying miles for one customer
*/
contract FrequentFlyerRewardsCalculator
{
     //Set of States
    enum StateType {SetFlyerAndReward, MilesAdded}

    //List of properties
    StateType private  State;
    address private  AirlineRepresentative;
    address private  Flyer;
    uint private RewardsPerMile;
    uint[] private Miles;
    uint private IndexCalculatedUpto;
    uint private TotalRewards;

    // constructor function
    constructor(address flyer, int rewardsPerMile) public
    {
        AirlineRepresentative = msg.sender;
        Flyer = flyer;
        RewardsPerMile = uint(rewardsPerMile);
        IndexCalculatedUpto = 0;
        TotalRewards = 0;
        State = StateType.SetFlyerAndReward;
    }

    // function that allows the contract to be destroyed
    // selfdestruct also transfers any contract balance to the adress passed as parameter
    function destroy() {
        if(AirlineRepresentative == msg.sender){
            selfdestruct(msg.sender);
        }
    }

    // call this function to add miles
    // this is a public function that changes values
    function AddMiles(int[] memory miles) public
    {
        if (Flyer != msg.sender)
        {
            return;
        }

        for (uint i = 0; i < miles.length; i++)
        {
            Miles.push(uint(miles[i]));
        }

        ComputeTotalRewards();

        State = StateType.MilesAdded;
    }

    // private function that changes values
    function ComputeTotalRewards() private
    {
        // make length uint compatible
        uint milesLength = uint(Miles.length);
        for (uint i = IndexCalculatedUpto; i < milesLength; i++)
        {
            // calling CalcReward function too calculates the reward
            TotalRewards += CalcReward(RewardsPerMile, Miles[i]);
            IndexCalculatedUpto++;
        }
    }

    // public function that only reads a value
    function GetMiles() public view returns (uint[] memory) {
        return Miles;
    }

    
    /*
    Pure functions ensure that they not read or modify the state. 
    A function can be declared as pure. 
    The following statements if present in the function are considered reading the state 
    and compiler will throw warning in such cases.
        
        - Reading state variables.
        - Accessing address(this).balance or <address>.balance.
        - Accessing any of the special variable of block, tx, msg (msg.sig and msg.data can be read).
        - Calling any function not marked pure.
        - Using inline assembly that contains certain opcodes.

    Pure functions can use the revert() and require() functions to revert potential state changes if an error occurs.
    */
    function CalcReward(uint rewardsPerMile, uint miles) public pure returns(uint){
      uint reward = rewardsPerMile * miles;
      return reward;
   }
   
   // external view function that can only be called from outside
   function GetRewardsPerMile() external view returns(uint) {
       return RewardsPerMile;
   }
}