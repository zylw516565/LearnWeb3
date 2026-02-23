package main

import (
	"bytes"

	"github.com/btcsuite/btcutil/base58"
)

type TXOutput struct {
	Value      int
	PubKeyHash []byte
}

func (out *TXOutput) Lock(address []byte) {
	fullPayload := base58.Decode(string(address))
	versionedPayload := fullPayload[:len(fullPayload)-addressChecksumLen]
	pubKeyHash := versionedPayload[1:]

	out.PubKeyHash = pubKeyHash
}

func (out *TXOutput) IsLockWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

// NewTXOutput create a new TXOutput
func NewTXOutput(value int, address string) *TXOutput {
	txout := TXOutput{value, nil}
	txout.Lock([]byte(address))

	return &txout
}
