package main

type Wallets struct {
	Wallets map[string]*Wallet
}

func NewWallets() (*Wallets, error) {
	return &Wallets{}, nil
}

func (ws Wallets) GetWallet(address string) Wallet {
	return Wallet{}
}
