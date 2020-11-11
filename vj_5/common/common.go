package common

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
)


// Kada imamo strukture moramo te strukture "pretočiti" u niz byte-ova
// kako bi mogli izvuči hash iz te strukture

// Za pretvaranje struktura u niz byte-ova koristimo paket encoding/gob

type Hash [32]byte
type Address [4]byte

// Generička funkcija koja vraća sha256 od bilo kojeg tipa
func HashOf(a interface{}) (retval Hash){
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
func ValuesHash(a ...[]byte) (retval Hash){
	buffer := &bytes.Buffer{}
	for _, v := range a {
		buffer.Write(v)
	}
	s := buffer.Bytes()
	retval = sha256.Sum256(s)
	return
}

// Generiranje nasumične adrese (za testiranje)
func GenerateNewAddress() (a Address) {
	addr := make([]byte, 4)
	rand.Read(addr)
	copy(a[:], addr)
	return a
}


