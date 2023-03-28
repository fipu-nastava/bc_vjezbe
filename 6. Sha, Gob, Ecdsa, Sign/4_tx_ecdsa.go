package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"math/big"
)


// Ovo je samo iskopirani kod

const AddressLength = 4

type Hash [32]byte

type Address [AddressLength]byte


// Generička funkcija koja vraća sha256 od bilo kojeg tipa
func HashOf(a interface{}) (retval Hash) {

	buffer := &bytes.Buffer{}
	enc := gob.NewEncoder(buffer)

	err := enc.Encode(a)
	if err != nil {
		panic(err)
	}

	s := buffer.Bytes()
	retval = sha256.Sum256(s)

	return
}


// Generička funkcija koja vraća sha256 od svih predanih vrijednosti
func ValuesHash(a ...[]byte) (retval Hash)  {

	buffer := new(bytes.Buffer)

	for _, val := range a {
		buffer.Write(val)
	}

	retval = sha256.Sum256(buffer.Bytes())

	return
}

// Generiranje nasumične adrese
func GenerateNewAddress() (a Address) {
	addr := make([]byte, 4)
	rand.Read(addr)
	copy(a[:], addr)
	return a
}

// Kraj iskorištenog iskopiranog koda


type transaction struct {
	From Address
	To Address
	Amount int
	Signature []byte
}

func NewTransaction(from Address, to Address, amount int) (t *transaction) {
	t = &transaction{}
	t.From = from
	t.To = to
	t.Amount = amount

	return
}


// Hash samo određenih polja strukture Transaction
func (t *transaction) TxHash() Hash  {

	return ValuesHash( t.From[:],
	    				t.To[:],
		    			[]byte( fmt.Sprintf("%d", t.Amount)[:]) )
}



func (t *transaction) Sign(key ecdsa.PrivateKey) {
	/*
	   Potpisivanje transakcije privatnim ključem.
	*/

	txHash := t.TxHash()
	// Potpiši transakciju
	r, s, err := ecdsa.Sign(rand.Reader, &key, txHash[:])

	if err != nil {
		panic(err)
	}

	// Spremi r, s u signature
	t.Signature = append(r.Bytes()[:32], s.Bytes()[:32]...)
}

func (t* transaction) Verify(key ecdsa.PublicKey) bool {

	txHash := t.TxHash()

	r, s := new(big.Int), new(big.Int)

	r.SetBytes(t.Signature[:32])
	s.SetBytes(t.Signature[32:64])

	isValid := ecdsa.Verify(&key, txHash[:], r, s)

	return isValid
}
