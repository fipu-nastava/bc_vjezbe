package v4

import (
	"crypto/sha256"
	"fmt"
)

// Kriptiranje - generiranje hash-a iz niza byte-ova
// Paket crypto/sha256


func main(){

	data := "Kriptografija je cool!\n"

	sum := sha256.Sum256([]byte(data))

	fmt.Printf("%x", sum)
}