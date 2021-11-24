package v4

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"math/rand"
)

// Kada imamo strukture moramo te strukture "pretočiti" u niz byte-ova
// kako bi mogli izvuči hash iz te strukture

// Za pretvaranje struktura u niz byte-ova koristimo paket encoding/gob

const AddressLength = 4

type Hash [32]byte

type Address [AddressLength]byte


type Transaction struct {
	From Address
	To Address
	Amount int
	Signature []byte
}


// Generiranje nasumične adrese
func GenerateNewAddress() (a Address) {
	addr := make([]byte, 4)
	rand.Read(addr)
	copy(a[:], addr)
	return a
}


// Hash cijele strukture Transaction
func (tx *Transaction) Hash() Hash {

	return HashOf(tx)
}


// Hash samo određenih polja strukture Transaction
func (tx *Transaction) TxHash() Hash  {

	return ValuesHash( tx.From[:],
					   tx.To[:],
					   []byte( fmt.Sprintf("%d", tx.Amount)[:]) )
}



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


func main() {
	a1 := GenerateNewAddress()
	a2 := GenerateNewAddress()

	t := Transaction{From: a1, To: a2, Amount: 22}

	fmt.Println(t.Hash())
	fmt.Println(t.TxHash())
}
