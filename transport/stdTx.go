package transport

type Transaction struct {
	Msg  []interface{} `json:"msg"`
	Sigs []interface{} `json:"signatures"`
	Fee  Fee           `json:"fee"`
}

type Signature struct {
	PubKey   []interface{} `json:"pub_key"`
	Sig      []interface{} `json:"signature"`
	Sequence int64         `json:"sequence"`
}

type SignMsg struct {
	ChainID   string  `json:"chain_id"`
	Sequences []int64 `json:"sequences"`
	FeeBytes  []byte  `json:"fee_bytes"`
	MsgBytes  []byte  `json:"msg_bytes"`
	AltBytes  []byte  `json:"alt_bytes"`
}

type Fee struct {
	Amount []int64 `json"amount"`
	Gas    int64   `json"gas"`
}
