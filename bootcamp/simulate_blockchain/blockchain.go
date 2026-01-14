package main

import (
	"log"

	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const BlocksBucket = "blocks"

// 区块链结构
type BlockChain struct {
	// hash_list  [][]byte
	// hashBlocks map[string]*Block

	tip []byte
	db  *bolt.DB
}

func (bc *BlockChain) AddBlock(data string) {
	var preHash []byte
	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BlocksBucket))
		preHash = b.Get([]byte("l"))

		return nil
	})

	if err != nil {
		log.Fatal("db.View failed !\n")
	}

	newBlock := NewBlock(data, preHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BlocksBucket))
		b.Put([]byte("l"), newBlock.Hash)
		b.Put(newBlock.Hash, newBlock.Serialize())
		bc.tip = newBlock.Hash

		return nil
	})
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

	return &BlockChain{tip, db}
}

// Iterator returns a BlockchainIterat
func (bc *BlockChain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}

	return bci
}

func (bc *BlockChain) FindSpendableOutputs(from string, amount int) (int, map[string][]int) {

	return 0, map[string][]int{}
}
