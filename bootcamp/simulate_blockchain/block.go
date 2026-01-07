package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// 区块结构
type Block struct {
	Timestamp     int64
	PrevBlockHash []byte
	Data          []byte
	Hash          []byte
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
	}

	newBlock.SetHash()
	return newBlock
}

func (b *Block) SetHash() {
	timestamp := strconv.FormatInt(b.Timestamp, 10)
	headers := bytes.Join([][]byte{[]byte(timestamp), b.PrevBlockHash, b.Data}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
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
