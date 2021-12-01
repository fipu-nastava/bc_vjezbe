pragma solidity >=0.4.25 <0.6.0;


/*
This is a simple market place auction for one item where the owner publishes the contract,
and the buyers make offers which the owner can reject or accept.

There can only be one offer at a certain time. When rejected, other buyers can offer their price.
*/
contract SimpleMarketplace {
    enum StateType { 
      ItemAvailable,
      OfferPlaced,
      Accepted
    }

    address public InstanceOwner;
    string public Description;
    int public AskingPrice;
    StateType public State;

    address public InstanceBuyer;
    int public OfferPrice;

    constructor(string memory description, int price) public
    {
        InstanceOwner = msg.sender;
        AskingPrice = price;
        Description = description;
        State = StateType.ItemAvailable;
    }

    /* 
    An offer can only be made:
        - if the price is not 0
        - if the item is available
        - if the offer is not made by the owner
    */
    function MakeOffer(int offerPrice) public
    {
        if (offerPrice == 0)
        {
            return;
        }

        if (State != StateType.ItemAvailable)
        {
            return;
        }
        
        if (InstanceOwner == msg.sender)
        {
            return;
        }

        InstanceBuyer = msg.sender;
        OfferPrice = offerPrice;
        State = StateType.OfferPlaced;
    }
    
    /*
    Only the owner can reject the offer when:
        - an offer has been placed
    */
    function Reject() public
    {
        if ( State != StateType.OfferPlaced )
        {
            return;
        }

        if (InstanceOwner != msg.sender)
        {
            return;
        }

        InstanceBuyer = 0x0000000000000000000000000000000000000000;
        State = StateType.ItemAvailable;
    }

    /*
    Only the owner can accept the offer
    */
    function AcceptOffer() public
    {
        if ( msg.sender != InstanceOwner )
        {
            return;
        }

        State = StateType.Accepted;
    }
}