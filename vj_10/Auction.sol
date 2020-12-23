pragma solidity ^0.5.0;

contract Auction {

    enum Stages {
        AcceptingBlindedBids,
        RevealBids
    }

    struct Bid {
        address bidder;
        uint amount;
    }

    Bid[] bids;

    // Trenutno stanje
    Stages public stage = Stages.AcceptingBlindedBids;

    uint public creationTime = now; // ovo se evaluira pri deploymentu

    modifier atStage(Stages _stage) {
        require(stage == _stage, "Poziv nije moguć u ovom trenutku.");
        _;
    }

    modifier minStage(Stages _stage) {
        require(stage <= _stage, "Poziv nije moguć u ovom trenutku.");
        _;
    }

    function nextStage() internal {
        // pretvori enum u integer, uvećaj i pretvori natrag
        stage = Stages(uint(stage) + 1);
    }

    // Modifier koji pazi da smo u ispravnoj fazi ovisno o vremenu
    // Pošto se kod ne može izrvšavati bez transakcija, ovo će se izrvšiti prije transakcije
    modifier timedTransitions() {
        // 10 dana traje prvi krug (5 minuta za testiranje)
        if (stage == Stages.AcceptingBlindedBids && now >= creationTime + 5 minutes) {
            nextStage();
        }
        // zatim 2 dana traje sljedeća faza (10 minuta za testiranje)
        // 10 minuta od početka je zapravo 5 minuta od kraja prve faze
        if (stage == Stages.RevealBids && now >= creationTime + 10 minutes) {
            nextStage();
        }

        _;
    }

    // Paziti na poredak modifiera, prvo želimo da se uveća faza (ako treba)
    // Ako se faza pomakne, atStage ce baciti gresku cime se zaustavlja mogucnost
    // dodavanja novih ponuda. Ako se javila greška, pomaknuta faza se neće zapamtiti na blockchain-u,
    // pa će sljedeći korisnik koji pozove funkciju ponovno okinuti prijelaz faze i ponavlja se sve isto
    // kao i za prvog korisnika
    function bid(uint amount) public payable timedTransitions atStage(Stages.AcceptingBlindedBids)
    {
        bids.push(Bid(msg.sender, amount));
    }

    // Pomoćna funkcija za prelazak u posljednju fazu
    // Poziva se modifier timedTransitions
    // Ako nebi bilo ove funkcije morali bi dodati novu ponudu
    // kako bi se faza pomaknula unaprijed
    function reveal() public timedTransitions atStage(Stages.RevealBids)
    {
    }

    // Funkcije za dohvaćanje ponuda
    // Koriste se kad se završi postavljanje ponuda
    function getBidCount() public view minStage(Stages.RevealBids) returns (uint) {
        return bids.length;
    }

    function getBidAmountAt(uint index) public view minStage(Stages.RevealBids) returns (uint)
    {
        // Provjera da li je index unutar granica
        require(index < bids.length);
        return bids[index].amount;
    }

}
