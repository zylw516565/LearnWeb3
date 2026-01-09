package main

import (
	"bytes"
	"encoding/gob"
	"time"
)

// 区块结构
type Block struct {
	Timestamp     int64
	PrevBlockHash []byte
	Data          []byte
	Hash          []byte
	Nonce         int
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	encoder.Encode(b)

	return result.Bytes()
}

// 区块链结构
type BlockChain struct {
	// hash_list  [][]byte
	// hashBlocks map[string]*Block

	blocks []*Block
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	newBlock := &Block{
		Timestamp:     time.Now().Unix(),
		PrevBlockHash: prevBlockHash,
		Data:          []byte(data),
		Hash:          []byte{},
		Nonce:         0,
	}

	pow := NewProofOfWork(newBlock)
	nonce, hash := pow.Run()
	newBlock.Hash = hash
	newBlock.Nonce = nonce

	return newBlock
}

func (bc *BlockChain) AddBlock(data string) {
	preBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, preBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewGenesisBlock() *Block {
	return NewBlock("GenesisBlock", []byte{})
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}
