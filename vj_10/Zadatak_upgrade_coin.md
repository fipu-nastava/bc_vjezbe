## Zadatak

- Nadogradnja [Coin](vj_10/Coin.sol) ugovora (prethodni zadatak)
  - apstrahirati koncept vlasništva u ugovor **Ownable**
  - apstrahirati koncept povlačenja ugovora u posebnu natklasu **Destructible**
  - Coin ugovor je **Destructible** (samim time i **Ownable**)
  - dodati funkciju **toZero(kome)** kojom **samo vlasnik** može poništiti balance (postaviti na 0) korisnika s adresom **kome** 
  - dodati funkciju **withdraw(iznos)** kojom korisnik može podignuti **iznos**, ali samo ako ima dovoljno sredstava (korisiti modyfier-e)