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
<img src="https://juststickers.in/wp-content/uploads/2016/07/go-programming-language.png" alt="Golang" style="width:25%;" />
</div>





# Tx sign

Zadatak: koristeći prethodni kod kreiraj strukturu Transaction koja ima metode:
   - `Sign(key ecdsa.PrivateKey, keyp ecdsa.PublicKey)` potpisuje transakciju
   - `Verify(pubKey ecdsa.PublicKey) bool` - vraća ispravnost popisa transakcije
   - Hint: u `Transaction` dodaj novo polje `Signature []byte` koje nije uključeno kod računanja hash-a

```go
func main() {

 	tx := NewTransaction(GenerateNewAddress(), GenerateNewAddress(), 20)

	pubKeyCurve := elliptic.P256() // http://golang.org/pkg/crypto/elliptic/#P256

	privateKey := new(ecdsa.PrivateKey)
	privateKey, err := ecdsa.GenerateKey(pubKeyCurve, rand.Reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pubKey := privateKey.PublicKey

	// potpiši transakciju
	tx.Sign(*privateKey)


	// verificiraj
	fmt.Println("Valid? ", tx.Verify(pubKey))

}
```

