package main

import (
	"log"
)

func (cli *CLI) createBlockChainCmd(address string) {
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}

}
