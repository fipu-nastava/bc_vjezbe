pragma solidity >=0.4.25 <0.6.0;


/*

The room thermostat scenario expresses a workflow around thermostat installation and use. 
In this scenario, a person will install a thermostat and indicate who is the intended user for the thermostat. 
The assigned user can do things such as set the target temperature and set the mode for the thermostat.

*/
contract RoomThermostat {
    //Set of States
    enum StateType { Created, InUse}
    
    //List of properties
    StateType public State;
    address public Installer;
    address public User;
    int public TargetTemperature;
    enum ModeEnum {Off, Cool, Heat, Auto}
    ModeEnum public  Mode;
    
    constructor(address thermostatInstaller, address thermostatUser) public
    {
        Installer = thermostatInstaller;
        User = thermostatUser;
        TargetTemperature = 70;
    }

    /*
    Thermostat can only be started by the Installer when the thermostat is created
    */
    function StartThermostat() public
    {
        if (Installer != msg.sender || State != StateType.Created)
        {
            return;
        }

        State = StateType.InUse;
    }

    /*
    Thermostat temperature can only be adjusted by the User when in use
    */
    function SetTargetTemperature(int targetTemperature) public
    {
        if (User != msg.sender || State != StateType.InUse)
        {
            return;
        }
        TargetTemperature = targetTemperature;
    }

    /*
    Thermostat mode can only be adjusted by the User when in use
    */
    function SetMode(ModeEnum mode) public
    {
        if (User != msg.sender || State != StateType.InUse)
        {
            return;
        }
        Mode = mode;
    }
}