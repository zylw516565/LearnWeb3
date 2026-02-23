package main

import (
	"fmt"
	"log"

	"github.com/btcsuite/btcutil/base58"
)

func (cli *CLI) getBalance(address, nodeID string) {
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}

	bc := NewBlockchain(nodeID)
	defer bc.db.Close()

	pubKeyHash := base58.Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := bc.FindUTXO(pubKeyHash)

	balance := 0
	for _, utxo := range UTXOs {
		balance += utxo.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}
