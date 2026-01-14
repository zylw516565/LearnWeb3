package main

import (
	"encoding/hex"
	"log"
)

type Transaction struct {
	ID   []byte
	Vin  []TxInput
	Vout []TxOutput
}

func NewUTXOTransaction(from, to string, amount int, bc *BlockChain) *Transaction {
	var inputs []TxInput
	var outputs []TxOutput

	acc, validOutputs := bc.FindSpendableOutputs(from, amount)

	if acc < amount {
		log.Panic("ERROR: Not enough funds")
	}

	for txid, outs := range validOutputs {
		txID, _ := hex.DecodeString(txid)

		for _, out := range outs {
			inputs = append(inputs, TxInput{txID, out, from})
		}
	}

	return &Transaction{}
}
