package main

import (
	"advent/day25/crypto"
	"fmt"
)

func main() {
	doorPubKey := int64(18499292)
	cardPubKey := int64(8790390)
	mod := int64(20201227)
	doorPrivKey := crypto.DiscreteLog(int64(7), doorPubKey, mod)
	encryptionKey := crypto.Exp(cardPubKey, doorPrivKey, mod)
	fmt.Printf("Encryption key: %d\n", encryptionKey)
}
