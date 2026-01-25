package main

import (
	"encoding/hex"
	"log"

	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const BlocksBucket = "blocks"

const genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

// 区块链结构
type BlockChain struct {
	tip []byte
	db  *bolt.DB
}

func (bc *BlockChain) AddBlock(transactions []*Transaction) {
	var preHash []byte
	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BlocksBucket))
		preHash = b.Get([]byte("l"))

		return nil
	})

	if err != nil {
		log.Fatal("db.View failed !\n")
	}

	newBlock := NewBlock(transactions, preHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BlocksBucket))
		b.Put([]byte("l"), newBlock.Hash)
		b.Put(newBlock.Hash, newBlock.Serialize())
		bc.tip = newBlock.Hash

		return nil
	})
}

func NewBlockChain(adress string) *BlockChain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0666, nil)
	if err != nil {
		log.Fatal("Open %s failed !\n", dbFile)
	}

	err = db.Update(
		func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(BlocksBucket))

			if nil == b {
				cbtx := NewCoinbaseTX(adress, genesisCoinbaseData)
				genesisBlock := NewGenesisBlock(cbtx)
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

func (bc *BlockChain) FindSpendableOutputs(address string, amount int) (int, map[string][]int) {
	unspentOutputs := make(map[string][]int)
	unspentTXs := bc.FindUnspentTransactions(address)
	accumulated := 0

Work:
	for _, tx := range unspentTXs {
		txID := hex.EncodeToString(tx.ID)

		for outIdx, out := range tx.Vout {
			if out.IsLockWithKey(pubKeyHash []byte) && accumulated < amount {
				accumulated += out.Value
				unspentOutputs[txID] = append(unspentOutputs[txID], outIdx)
			}

			if accumulated >= amount {
				break Work
			}
		}
	}

	return 0, map[string][]int{}
}

func (bc *BlockChain) FindUnspentTransactions(address string) []Transaction {
	var unspentTXs []Transaction
	return unspentTXs
}

func (bc *BlockChain) printChain() {

}
