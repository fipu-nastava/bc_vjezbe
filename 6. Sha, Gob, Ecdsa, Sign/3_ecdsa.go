package v4

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"os"
)

// Potpisivanje transakcija
// Paket crypt/ecdsa



func main() {
	// Kreiranje privatnog ključa

	pubKeyCurve := elliptic.P256() // http://golang.org/pkg/crypto/elliptic/#P256

	privateKey := new(ecdsa.PrivateKey)
	privateKey, err := ecdsa.GenerateKey(pubKeyCurve, rand.Reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pubKey := privateKey.PublicKey

	fmt.Println("Private Key :")
	fmt.Printf("%x \n", privateKey)
	fmt.Println("Public Key :")
	fmt.Printf("%x \n", pubKey)


	message := []byte("Poruka koju valja potpisati")

	sha := sha256.New()
	signHash := sha.Sum(message)

	fmt.Println(signHash)

	// Potpisuje hash poruke
	// vraća popis kao dva integera (big.Int)
	r, s, signErr := ecdsa.Sign(rand.Reader, privateKey, signHash)

	if signErr != nil {
		fmt.Println(signErr)
		os.Exit(1)
	}

	signature := append(r.Bytes(), s.Bytes()...)
	fmt.Printf("Signature : %x\n", signature)

	verifyStatus := ecdsa.Verify(&pubKey, signHash, r, s)
	fmt.Println(verifyStatus) // true

}
