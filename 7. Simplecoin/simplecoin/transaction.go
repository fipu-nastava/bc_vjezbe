package simplecoin

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type Transaction struct {
	From Address
	To Address
	Amount int
	Signature []byte
	VerificationKey ecdsa.PublicKey
}

func NewTransaction(from Address, to Address, amount int) (t *Transaction) {
	t = &Transaction{}
	t.From = from
	t.To = to
	t.Amount = amount

	return
}
// Hash samo određenih polja strukture Transaction
func (t *Transaction) Hash() Hash  {
	buffer := new(bytes.Buffer)

	buffer.Write(t.From[:])
	buffer.Write(t.To[:])
	buffer.WriteString(fmt.Sprintf("%d", t.Amount))

	txHash := sha256.Sum256(buffer.Bytes())

	return txHash
}



func (t Transaction) ShortString() string {
	return fmt.Sprintf("(txid: %x)", t.Signature[:4])
}

func (t Transaction) String() string {
	return fmt.Sprintf("(txid: %x, desc: %d coins from %s to %s)", t.Signature[:4], t.Amount, t.From, t.To)
}

func (t *Transaction) Sign(key ecdsa.PrivateKey, keyp ecdsa.PublicKey) {
	/*
		Potpisivanje transakcije privatnim ključem.
		Spremanje javnog ključa kao Verification key.
	*/

	txHash := t.Hash()

	r, s, err := ecdsa.Sign(rand.Reader, &key, txHash[:])

	if err != nil {
		panic(err)
	}

	t.Signature = append(r.Bytes()[:32], s.Bytes()[:32]...)
	t.VerificationKey = keyp
}

func (t* Transaction) Verify() bool {
	/*
		Provjera valjanosti transakcije. Transakcija je valjana ako:
			- adresa pošiljatelja izvedena od javnog ključa
		    - mogu verificirati njezin potpis javnim ključem
	*/
	txHash := t.Hash()

	r, s := new(big.Int), new(big.Int)
	r.SetBytes(t.Signature[:32])
	s.SetBytes(t.Signature[32:64])

	isValid := ecdsa.Verify(&t.VerificationKey, txHash[:], r, s)

	sender := t.From
	isValid = isValid && sender == ComputeAddress(t.VerificationKey)

	return isValid
}