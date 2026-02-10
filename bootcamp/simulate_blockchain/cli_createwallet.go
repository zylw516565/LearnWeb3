package main

import "fmt"

func (cli *CLI) createWalletCmd() {
	wallets, _ := NewWallets()
	address := wallets.CreateWallet()
	wallets.SaveToFile()

	fmt.Printf("Your new addres is %s\n", address)
}
