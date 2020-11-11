package main

import (
	. "../common"
	"../core"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"os"
)

/*
	Zadatak: koristeći prethodni kod u strukturu Transaction dodaj metode:
		- Sign(key ecdsa.PrivateKey) - potpisuje transakciju
		- Verify(key ecdsa.PublicKey) bool - vraća ispravnost popisa transakcije

*/



func main() {

	tx := core.NewTransaction(GenerateNewAddress(), GenerateNewAddress(), 20)


	privateKey := GenerateKey()
	pubKey := privateKey.PublicKey

	// potpiši transakciju
	tx.Sign(*privateKey)


	// verificiraj
	fmt.Println("Valid? ", tx.Verify(pubKey))


	// da li je možda neki drugi korisnik potpisao?
	privKey2 := GenerateKey()
	fmt.Println("Valid? ", tx.Verify(privKey2.PublicKey))

}



// Kreiranje privatnog ključa
func GenerateKey() (privKey *ecdsa.PrivateKey)  {
	pubKeyCurve := elliptic.P256() // http://golang.org/pkg/crypto/elliptic/#P256

	privKey, err := ecdsa.GenerateKey(pubKeyCurve, rand.Reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}

// Primjer potpisa transakcije
func sign_example(){
	privKey := GenerateKey()
	pubKey := privKey.PublicKey

	fmt.Println("Private Key :")
	fmt.Printf("%x \n", privKey)
	fmt.Println("Public Key :")
	fmt.Printf("%x \n", pubKey)

	tx := core.NewTransaction(GenerateNewAddress(), GenerateNewAddress(), 20)
	signHash := tx.TxHash()

	// Potpisuje hash transakcije
	r, s, err := ecdsa.Sign(rand.Reader, privKey, signHash[:])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Spaja bytove obih vrijednosti u jedan byte niz
	signature := append(r.Bytes(), s.Bytes()...)
	fmt.Printf("Signature : %x\n", signature)

	// Verificiraj je li korisnik s pubKey potpisao signHash
	status := ecdsa.Verify(&pubKey, signHash[:], r, s)
	fmt.Println(status) // true

}