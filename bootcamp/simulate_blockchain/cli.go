package main

import (
	"flag"
	"fmt"
	"os"
)

type CLI struct {
}

func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	addBlockData := addBlockCmd.String("data", "", "Block data")
	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	createBlockChainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	addressData := createBlockChainCmd.String("address", "", "address")
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	addressGetBalance := getBalanceCmd.String("address", "", "address")
	listAddressCmd := flag.NewFlagSet("listaddress", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	from := sendCmd.String("from", "", "from")
	to := sendCmd.String("to", "", "to")
	amount := sendCmd.Int("amount", 0, "amount")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			cli.printUsage()
			os.Exit(1)
		}

	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			cli.printUsage()
			os.Exit(1)
		}

	case "createwallet":
		err := createWalletCmd.Parse(os.Args[2:])
		if err != nil {
			cli.printUsage()
			os.Exit(1)
		}

	case "createblockchain":
		err := createBlockChainCmd.Parse(os.Args[2:])
		if err != nil {
			cli.printUsage()
			os.Exit(1)
		}

	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			cli.printUsage()
			os.Exit(1)
		}

	case "listaddress":
		err := listAddressCmd.Parse(os.Args[2:])
		if err != nil {
			cli.printUsage()
			os.Exit(1)
		}

	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			cli.printUsage()
			os.Exit(1)
		}

	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if "" == *addBlockData {
			addBlockCmd.Usage()
			os.Exit(1)
		}

		// cli.AddBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

	if createWalletCmd.Parsed() {
		cli.createWalletCmd()
	}

	if createBlockChainCmd.Parsed() {
		if "" == *addressData {
			createBlockChainCmd.Usage()
			os.Exit(1)
		}

		cli.createBlockChainCmd(*addressData)
	}

	if getBalanceCmd.Parsed() {
		if "" == *addressGetBalance {
			getBalanceCmd.Usage()
			os.Exit(1)
		}

		cli.getBalance(*addressGetBalance)
	}

	if listAddressCmd.Parsed() {
		cli.listAddressCmd()
	}

	if sendCmd.Parsed() {
		if "" == *from || "" == *to || 0 == *amount {
			sendCmd.Usage()
			os.Exit(1)
		}

		cli.sendCmd(*from, *to, *amount)
	}

}

func (cli *CLI) validateArgs() {
	if len(os.Args) <= 1 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) printUsage() {
	fmt.Println("  addblock - Addblock to the blockchain, --data with data")
	fmt.Println("  printchain - Print all the blocks of the blockchain")
}

func (cli *CLI) addBlock(data string) {
	// cli.bc.AddBlock(data)
	fmt.Println("Success!")
}
