package model

type TransferToAddress struct {
	Sender       string `json:"sender"`
	ReceiverAddr string `json:"receiver_addr"`
	Amount       string `json:"amount"`
	Memo         string `json:"memo"`
}
