<table style="caret-color: #000000; font-family: Georgia;" border="0" cellspacing="0" cellpadding="0" >
            <tbody>
              <tr>
                <td valign="center">
                  <a id="logo_a" href="https://fipu.unipu.hr"><img id="logo_img"  src="https://www.unipu.hr/_download/repository/FIPU_horiz_kolor_HR.png" alt="logotip FIPU" title="Fakultet informatike u Puli"></a> 								 </td>
              </tr>
  </tbody>
</table>



<img src="https://solidity.readthedocs.io/en/v0.5.14/_images/logo.svg" alt="solidity" style="zoom:10%;"/>

# [Solidity funkcije](https://solidity.readthedocs.io/en/v0.5.14/)

- <span style="color:blue">addmod(x,y,k) </span>  
  <span style="color:blue">mulmod(x,y,k) </span>   
  Modulo zbrajanje i množenje
- <span style="color:blue">keccak256, sha256, sha3, ripemd160 </span>   
  Izračun različitih hasheva
- <span style="color:blue">ecrecover</span>   
   Ekstrakcija adrese iz potpisa (v, r, s)
- <span style="color:blue">selfdestruct(recipient_address)</span>   
   Brisanje trenutnog ugovora i slanje preostalog iznosa na željenu adresu
- <span style="color:blue">this </span>  
   Adresa trenutno izvršavajućeg ugovora



## Funkcije

- Navode se unutar ugovora

**Deklaracija**

<span style="color:blue">function </span> **NazivFunkcije**([*parameters*])   
	{<span style="color:blue">public </span>|<span style="color:blue">private </span>|<span style="color:blue">internal </span>|<span style="color:blue">external</span>}    
	[<span style="color:blue">pure </span>|<span style="color:blue">view </span>|<span style="color:blue">payable </span>] [*modifiers*]    
	[<span style="color:blue">returns </span> (*return types*)]  



- Parametri
  - Lista parametara (<span style="color:blue">tip </span> + naziv)
- Vidljivost funkcije
  - <span style="color:blue">**public** </span>  
    (izvana - EOA (Externally Owned Accounts) i iznutra)
  - <span style="color:blue">**external** </span>  
    (samo izvana)
  - <span style="color:blue">**internal** </span>  
    (samo iznutra i iz nasljeđenih ugovora) 
  - <span style="color:blue">**private** </span>  
    (samo iznutra i trenutnog ugovora)
- Ponašanje funkcije:
  - <span style="color:blue">**view** </span>   
     (funkcija koja obećaje da neće mijenjate stanje - memorija pridružena ugovoru)
  - <span style="color:blue">**pure** </span>   
     (funkcija koja ne koristi varijable - samo parametre - promiče deklarativno programiranje)
  - <span style="color:blue">**payable** </span>   
     (dozvoljava primanje vrijednosti - wei - uz njezino pozivanje)



## Konstruktor

- Poziva se prilikom kreiranja (*deployment*) ugovora
- Zadužen za postavljanje inicijalnih vrijednosti stanja ugovora

```php
contract BCA {
    // Konstruktor
    constructor() public { 
      		// ...
      }
}
```



## Destruktor

- Ugovor može podržavati i povlačenje, ali to mora biti omogućeno prilikom njegovog postavljanja
- Ne postoji službeni destruktor nego se koristi <span style="color:blue">selfdestruct(recipient_address)</span>  funkcija

```php
// Metoda koja će brisati ugovor, može se nazvati bilo kako
function destroy() public {
		// Poništi ugovor i prebaci preostala sredstva na adresu vlasnika
  	selfdestruct(owner);
}
```





## Zadatak

- Napraviti ugovor za vlastiti **coin** ugovor
- Funkcionalnosti:
  - **balanceOf**(korisnik) => vraća količinu novaca za korisnika
  - **transfer**(kome, koliko) => funkcija prebacuje s jednog računa na drugi
  - **deposit**() => funkcija za pošiljatelja bilježi koliko je poslao novaca