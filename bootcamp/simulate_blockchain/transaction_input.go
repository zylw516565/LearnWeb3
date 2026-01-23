package main

import "bytes"

type TxInput struct {
	Txid      []byte
	Vout      int
	Signature []byte
	PubKey    []byte

	ScriptSig string
}

func (in *TxInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}

func (in *TxInput) UsesKey(key []byte) bool {
	pubKeyHash := HashPubKey(in.PubKey)

	return bytes.Compare(pubKeyHash, key) == 0
}
