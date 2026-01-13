package main

type Transaction struct {
	ID   []byte
	Vin  []TxInput
	Vout []TxOutput
}
