package main

import (
	"bytes"

	"github.com/btcsuite/btcutil/base58"
)

type TxOutput struct {
	Value      int
	PubKeyHash []byte
}

func (out *TxOutput) Lock(address []byte) {
	fullPayload := base58.Decode(string(address))
	versionedPayload := fullPayload[:len(fullPayload)-addressChecksumLen]
	pubKeyHash := versionedPayload[1:]

	out.PubKeyHash = pubKeyHash
}

func (out *TxOutput) IsLockWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

// NewTXOutput create a new TXOutput
func NewTXOutput(value int, address string) *TxOutput {
	txout := TxOutput{value, nil}
	txout.Lock([]byte(address))

	return &txout
}
