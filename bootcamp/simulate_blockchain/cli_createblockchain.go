package main

import (
	"fmt"
	"log"
)

func (cli *CLI) createBlockChainCmd(address, nodeID string) {
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}

	bc := CreateBlockchain(address, nodeID)
	bc.db.Close()
	fmt.Println("Done!")
}
