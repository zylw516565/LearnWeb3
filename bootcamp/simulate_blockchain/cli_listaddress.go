package main

import (
	"fmt"
	"log"
)

func (cli *CLI) listAddressCmd(nodeID string) {
	wallets, err := NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}

	addresses := wallets.GetAddresses()
	// fmt.Println(addresses)
	for _, address := range addresses {
		fmt.Println(address)
	}
}
