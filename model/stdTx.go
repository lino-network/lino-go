package model

import (
	"github.com/tendermint/go-crypto"
)

type Transaction struct {
	Msg  Msg         `json:"msg"`
	Fee  Fee         `json:"fee"`
	Sigs []Signature `json:"signatures"`
}

type Signature struct {
	PubKey        crypto.PubKey    `json:"pub_key"`
	Sig           crypto.Signature `json:"signature"`
	AccountNumber int64            `json:"account_number"`
	Sequence      int64            `json:"sequence"`
}

type SignMsg struct {
	ChainID        string  `json:"chain_id"`
	AccountNumbers []int64 `json:"account_numbers"`
	Sequences      []int64 `json:"sequences"`
	FeeBytes       []byte  `json:"fee_bytes"`
	MsgBytes       []byte  `json:"msg_bytes"`
	AltBytes       []byte  `json:"alt_bytes"`
}

type Fee struct {
	Amount SDKCoins `json:"amount"`
	Gas    int64    `json:"gas"`
}
