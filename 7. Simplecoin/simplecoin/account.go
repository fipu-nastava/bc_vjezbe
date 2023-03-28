package simplecoin

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

type Account struct {
	Address Address
	PublicKey ecdsa.PublicKey
	PrivateKey ecdsa.PrivateKey
}

type Address [32]byte


func ComputeAddress(key ecdsa.PublicKey) Address {
	buffer := new(bytes.Buffer)
	buffer.Write(key.X.Bytes())
	buffer.Write(key.Y.Bytes())

	return sha256.Sum256(buffer.Bytes())
}


func NewAccount() (a Account) {

	keypair := GenerateKeyPair()

	a.PrivateKey = keypair
	a.PublicKey = keypair.PublicKey

	a.Address = ComputeAddress(a.PublicKey)

	return
}

func (a Account) String() string {
	return fmt.Sprintf("(account:%x)", a.Address[:2])
}

func (a Address) String() string {
	return fmt.Sprintf("(address:%x)", a[:2])
}


func GenerateKeyPair() ecdsa.PrivateKey {
	pubkeyCurve := elliptic.P256()

	keypair := new(ecdsa.PrivateKey)
	keypair, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)

	if err != nil {
		panic(err)
	}

	return *keypair
}
