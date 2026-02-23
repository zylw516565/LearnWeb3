package main

import (
	"fmt"
	"log"
)

func (cli *CLI) sendCmd(from, to string, amount int, nodeID string, mineNow bool) {
	if !ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := NewBlockchain(nodeID)
	defer bc.db.Close()

	wallets, err := NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.GetWallet(from)

	tx := NewUTXOTransaction(&wallet, to, amount, bc)

	if mineNow {
		cbTx := NewCoinbaseTX(from, "")
		txs := []*Transaction{cbTx, tx}

		bc.MineBlock(txs)
		// newBlock := bc.MineBlock(txs)
		//UTXOSet.Update(newBlock)
	} else {
		sendTx(knownNodes[0], tx)
	}

	fmt.Println("Success!")
}
