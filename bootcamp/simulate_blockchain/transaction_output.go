package main

import (
	"bytes"

	"github.com/btcsuite/btcutil/base58"
)

type TxOutput struct {
	Value      int
	PubKeyHash []byte

	ScriptPubKey string
}

func (out *TxOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
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
