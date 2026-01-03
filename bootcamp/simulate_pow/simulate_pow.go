package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"hash"
)

const nickname = "peter"

func main() {
	fmt.Println("Begin to simulate pow!")

	nick := []byte(nickname)
	nonce := make([]byte, 16)

	hash := sha256.New()
	fmt.Printf("nonce: %x\n", nonce)
	calcHash(4, hash, nick, nonce)
	calcHash(5, hash, nick, nonce)
}

func calcHash(leadZeroBitwise uint, hash hash.Hash, nick, nonce []byte) {
	for {
		full_name := make([]byte, 0, len(nick)+len(nonce))
		genNonce(nonce)
		full_name = append(full_name, nick...)
		full_name = append(full_name, nonce...)

		hash.Write(full_name)
		result := hash.Sum(nil)

		// fmt.Printf("nonce: %x\t result:%x\t result[0]:%08b\n", nonce, result, result[0])
		if (result[0] >> (8 - leadZeroBitwise)) == 0 {
			break
		}
	}
}
func genNonce(nonce []byte) {
	_, err := rand.Read(nonce)
	if err != nil {
		panic(err)
	}
}
