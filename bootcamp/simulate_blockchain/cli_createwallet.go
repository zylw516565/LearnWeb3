package main

import (
	"fmt"
	"log"
)

func (cli *CLI) createWalletCmd(nodeID string) {
	wallets, _ := NewWallets(nodeID)
	address := wallets.CreateWallet()
	err := wallets.SaveToFile(nodeID)
	if nil != err {
		log.Panic(err)
	}

	fmt.Printf("Your new addres is %s\n", address)
}
