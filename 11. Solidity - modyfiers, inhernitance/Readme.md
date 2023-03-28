<div align="center">
<table style="caret-color: #000000; font-family: Georgia;" border="0" cellspacing="0" cellpadding="0" >
            <tbody>
              <tr>
                <td valign="center">
                  <a id="logo_a" href="https://fipu.unipu.hr"><img id="logo_img"  src="https://www.unipu.hr/_download/repository/FIPU_horiz_kolor_HR.png" alt="logotip FIPU" title="Fakultet informatike u Puli"></a> 								 </td>
              </tr>
  </tbody>
</table>
</div>



<div align="center">
  <img src="https://solidity.readthedocs.io/en/v0.5.14/_images/logo.svg" alt="solidity" style="width:25%;"/>
</div>


# [Solidity](https://solidity.readthedocs.io/en/v0.5.14/)

<span style="color:red">Imamo li ovdje problem?</span> 

```php
function destroy() public {
  	selfdestruct(owner);
}
```

- Funkciju <span style="color:blue">destroy()</span>  može pozvati bilo tko i time nam izbrisati ugovor

**Quick fix**

```php
function destroy() public {
  	// Samo owner može izbrisati ugovor (owner je adresa koja je deployala ugovor)
  	if (msg.sender == owner) {
  			selfdestruct(owner);
    }
}
```





## Upravljanje greškama

- ***assert***

  - provjerava uvjet, ukoliko nije istinit, troši preostali gas i vraća stanje ugovora

- ***require***
  - provjerava uvjet, ukoliko nije istinit, vraća preostali gas i stanje ugovora

**Prethodni primjer koristeći** ***require***

```php
function destroy() public {
  	// Samo owner može izbrisati ugovor (owner je adresa koja je deployala ugovor)
  	require (msg.sender == owner, "Poruka greške");

    selfdestruct(owner);

}
```





## Modifiers

- Solidity dopušta posebne dodatke funkcijama za upravljanje greškama
  - posebno se definiraju (jednom) i koriste na željenim funkcijama
  - "vide" parametre i varijable funkcije

```php
// Definicija modifiera za provjeru vlasništva
modifier onlyOwner {
		require(msg.sender == owner, "Nisi vlasnik!"); // Možemo umetniti i poruku greške
		_; // Oznaka za umetanje ostatka funkcije
}

// Metoda koja će brisati ugovor, ako to modifier onlyOwner dopusti
function destroy() public onlyOwner {
  // Poništi ugovor i prebaci preostala sredstva na adresu vlasnika
 	 selfdestruct(owner);
}
```



**Primjer**

```php
uint creationTime = now;

modifier onlyBy(address _account) {
    require(msg.sender == _account, "Sender not authorized.");
		_;
}

modifier onlyAfter(uint _time) {
		require(now >= _time, "Još nije vrijeme.");
  	_;
}

/// Erase ownership information.
/// May only be called 6 weeks after
/// the contract has been created.
function die() public onlyBy(owner) onlyAfter(creationTime + 6 weeks) {
    selfdestruct(owner)
}
```





## Nasljeđivanje

- Ključna riječ **is**

  ```php
  contract Faucet is Ownable {
   	// ...
  }
  ```

- Nasljeđeni ugovor vidi **public** i **internal** atribute i funkcije



## Zadatak

- Nadogradnja Coin ugovora (prethodni zadatak)
  - apstrahirati koncept vlasništva u ugovor **Ownable**
  - apstrahirati koncept povlačenja ugovora u posebnu natklasu **Destructible**
  - Coin ugovor je **Destructible** (samim time i **Ownable**)
  - dodati funkciju **toZero(kome)** kojom samo vlasnik može poništiti balance (postaviti na 0) korisnika s adresom **kome**
  - dodati funkciju **withdraw(iznos)** kojom korisnik može podignuti **iznos**, ali samo ako ima dovoljno sredstava
