package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const BlocksBucket = "blocks"

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

func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	decoder.Decode(&block)

	return &block
}

// 区块链结构
type BlockChain struct {
	// hash_list  [][]byte
	// hashBlocks map[string]*Block

	tip    []byte
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
	var tip []byte
	db, err := bolt.Open(dbFile, 0666, nil)
	if err != nil {
		log.Fatal("Open %s failed !\n", dbFile)
	}

	err = db.Update(
		func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(BlocksBucket))

			if nil == b {
				genesisBlock := NewGenesisBlock()
				b, err = tx.CreateBucket([]byte(BlocksBucket))
				b.Put(genesisBlock.Hash, genesisBlock.Serialize())
				b.Put([]byte("l"), genesisBlock.Hash)

				tip = genesisBlock.Hash
			} else {
				tip = b.Get([]byte("l"))
			}

			return nil
		})

	return &BlockChain{tip, []*Block{}}
}
