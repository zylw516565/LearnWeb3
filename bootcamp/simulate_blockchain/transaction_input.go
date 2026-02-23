package main

import "bytes"

type TXInput struct {
	Txid      []byte
	Vout      int
	Signature []byte
	PubKey    []byte
}

func (in *TXInput) UsesKey(key []byte) bool {
	pubKeyHash := HashPubKey(in.PubKey)

	return bytes.Compare(pubKeyHash, key) == 0
}
