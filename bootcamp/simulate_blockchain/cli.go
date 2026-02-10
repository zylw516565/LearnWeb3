package main

import (
	"flag"
	"fmt"
	"os"
)

type CLI struct {
	bc *BlockChain
}

func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	addBlockData := addBlockCmd.String("data", "", "Block data")
	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)

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

func (cli *CLI) printChain() {
	bci := cli.bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		// fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("Pow Validate: %x\n", pow.Validate())
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
