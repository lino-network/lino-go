package model

import (
	"encoding/json"

	crypto "github.com/tendermint/tendermint/crypto"
)

type Transaction struct {
	Msgs       []Msg       `json:"msg"`
	Fee        Fee         `json:"fee"`
	Signatures []Signature `json:"signatures"`
	Memo       string      `json:"memo"`
}

type Signature struct {
	PubKey        crypto.PubKey `json:"pub_key"`
	Sig           []byte        `json:"signature"`
	AccountNumber int64         `json:"account_number"`
	Sequence      int64         `json:"sequence"`
}

type SignMsg struct {
	AccountNumber int64             `json:"account_number"`
	ChainID       string            `json:"chain_id"`
	Fee           json.RawMessage   `json:"fee"`
	Memo          string            `json:"memo"`
	Msgs          []json.RawMessage `json:"msgs"`
	Sequence      int64             `json:"sequence"`
}

type Fee struct {
	Amount SDKCoins `json:"amount"`
	Gas    int64    `json:"gas"`
}
