package main

import (
	"fmt"
	"log"
)

func (cli *CLI) createWalletCmd() {
	wallets, _ := NewWallets()
	address := wallets.CreateWallet()
	err := wallets.SaveToFile()
	if nil != err {
		log.Panic(err)
	}

	fmt.Printf("Your new addres is %s\n", address)
}
