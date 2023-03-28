package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"os"
)

/*
	Zadatak: koristeći prethodni kod kreiraj strukturu Transaction koja ima metode:
		- Sign(key ecdsa.PrivateKey) - potpisuje transakciju
		- Verify(key ecdsa.PublicKey) bool - vraća ispravnost popisa transakcije

 */

func GenerateKey() (privKey *ecdsa.PrivateKey)  {
	pubKeyCurve := elliptic.P256() // http://golang.org/pkg/crypto/elliptic/#P256

	privKey, err := ecdsa.GenerateKey(pubKeyCurve, rand.Reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}

func main() {

 	tx := NewTransaction(GenerateNewAddress(), GenerateNewAddress(), 20)


	privateKey := GenerateKey()
	pubKey := privateKey.PublicKey

	// potpiši transakciju
	tx.Sign(*privateKey)


	// verificiraj
	fmt.Println("Valid?", tx.Verify(pubKey))


	// da li je možda neki drugi korisnik potpisao?
	privKey2 := GenerateKey()
	fmt.Println("Valid?", tx.Verify(privKey2.PublicKey))

}