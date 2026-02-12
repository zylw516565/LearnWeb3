package main

import (
	"fmt"
	"log"
)

func (cli *CLI) createBlockChainCmd(address string) {
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}

	bc := CreateBlockchain(address)
	bc.db.Close()
	fmt.Println("Done!")
}
