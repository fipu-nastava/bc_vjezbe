pragma solidity >=0.4.25 <0.7.0;


contract SimpleStorage {
    uint storedData;

    // public function that modifies a memory space (variable)  
    function set(uint x) public {
        storedData = x;
    }

    // public view function that only reads and returns a value 
    function get() public view returns (uint) {
        return storedData;
    }
}