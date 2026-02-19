package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
)

type Wallets struct {
	Wallets map[string]*Wallet
}

type SerializableWallet struct {
	D         *big.Int
	X, Y      *big.Int
	PublicKey []byte
}

func NewWallets() (*Wallets, error) {
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)

	err := wallets.LoadFromFile()
	return &wallets, err
}

// CreateWallet adds a Wallet to Wallets
func (ws *Wallets) CreateWallet() string {
	wallet := NewWallet()
	address := fmt.Sprintf("%s", wallet.GetAddress())
	ws.Wallets[address] = wallet

	return address
}

func (ws *Wallets) GetAddresses() []string {
	var addresses []string
	for addr, _ := range ws.Wallets {
		addresses = append(addresses, addr)
	}

	return addresses
}

func (ws Wallets) GetWallet(address string) Wallet {
	return *ws.Wallets[address]
}

func (ws Wallets) LoadFromFile() error {
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return err
	}

	fileContent, err := ioutil.ReadFile(walletFile)
	if nil != err {
		log.Panic(err)
	}

	var wallets map[string]SerializableWallet

	gob.Register(SerializableWallet{})
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&wallets)
	if nil != err {
		log.Panic(err)
	}

	for k, v := range wallets {
		ws.Wallets[k] = &Wallet{
			PrivateKey: ecdsa.PrivateKey{
				PublicKey: ecdsa.PublicKey{
					Curve: elliptic.P256(),
					X:     v.X,
					Y:     v.Y,
				},
				D: v.D,
			},
			PublicKey: v.PublicKey,
		}
	}

	return nil
}

func (ws Wallets) SaveToFile() error {
	var content bytes.Buffer
	gob.Register(SerializableWallet{})

	wallets := make(map[string]SerializableWallet)
	for k, v := range ws.Wallets {
		wallets[k] = SerializableWallet{
			v.PrivateKey.D,
			v.PrivateKey.X,
			v.PrivateKey.Y,
			v.PublicKey,
		}
	}

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(wallets)
	if nil != err {
		log.Panic(err)
	}

	err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
	if nil != err {
		log.Panic(err)
	}

	return nil
}
