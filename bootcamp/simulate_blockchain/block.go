package block

import (
	"bytes"
	"crypto/sha256"
	"strconv"
)

type Block struct {
	Timestamp     int64
	PrevBlockHash []byte
	Data          []byte
	Hash          []byte
}

func (b *Block) SetHash() {
	timestamp := strconv.FormatInt(b.Timestamp, 10)
	headers := bytes.Join([][]byte{[]byte(timestamp), b.PrevBlockHash, b.Data}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}
