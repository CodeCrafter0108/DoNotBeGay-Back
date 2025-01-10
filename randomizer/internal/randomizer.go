package internal

import (
	"crypto/rand"
	"log"
	"math/big"
)

var MaxUint256 = new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil), big.NewInt(1))

func generateRandomMaxUint64() *big.Int {
	randomNumber, err := rand.Int(rand.Reader, MaxUint256)
	if err != nil {
		panic(err)
	}
	return randomNumber
}

func RandomCall() {
	randomNumber := generateRandomMaxUint64()
	log.Printf("random number: %v\n", randomNumber)
	// TODO: call method on a contract
}
