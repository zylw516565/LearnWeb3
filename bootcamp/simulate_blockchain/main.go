package main

import (
	"fmt"
	"strconv"
)

func main() {
	bc := NewBlockChain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
		fmt.Printf("PrevHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		//对工作量证明进行验证
		pow := NewProofOfWork(block)
		fmt.Printf("Pow Validate %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

	//对工作量证明进行验证
	// for _, block := range bc.blocks {
	// 	pow := NewProofOfWork(block)

	// 	fmt.Printf("", pow.Validate())
	// }
}
