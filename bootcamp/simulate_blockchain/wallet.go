package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

const (
	Verion             = byte(0x00)
	walletFile         = "wallet.dat"
	addressChecksumLen = 4
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func NewWallet() *Wallet {
	private, public := newKeyPair()
	return &Wallet{private, public}
}

func newKeyPair() (ecdsa.PrivateKey, []byte) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if nil != err {
		return ecdsa.PrivateKey{}, []byte{}
	}

	pubKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)
	return *privateKey, pubKey
}

func (w Wallet) GetAddress() []byte {
	pkHash := HashPubKey(w.PublicKey)

	versionedPayload := append([]byte{Verion}, pkHash...)
	checksum := checksum(versionedPayload)
	fullPayload := append(versionedPayload, checksum...)
	address := base58.Encode(fullPayload)

	return []byte(address)
}

func HashPubKey(pubKey []byte) []byte {
	pubKeySHA256 := sha256.Sum256(pubKey)
	ripemd160Hash := ripemd160.New()
	ripemd160Hash.Write(pubKeySHA256[:])
	pubKehHash := ripemd160Hash.Sum(nil)

	return pubKehHash
}

func checksum(payload []byte) []byte {
	firstHash := sha256.Sum256(payload)
	secondHash := sha256.Sum256(firstHash[:])

	return secondHash[:addressChecksumLen]
}
