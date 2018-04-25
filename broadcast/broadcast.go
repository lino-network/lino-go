package broadcast

import (
	"encoding/json"

	"github.com/lino-network/lino-go/transport"
	"github.com/tendermint/go-crypto"
)

// type BroadcastCallback func(result, error)

func BroadcastTransaction(transaction interface{}, rawPrivKey string) {
	transport := transport.NewTransportFromViper()
	privKey, _ := crypto.PrivKeyFromBytes([]byte(rawPrivKey))
	txBytes, _ := SignAndBuild(transaction, privKey)
	transport.BroadcastTx(txBytes)
}

func SignAndBuild(transaction interface{}, privKey crypto.PrivKey) ([]byte, error) {
	// build
	signMsg := sdk.StdSignMsg{
		ChainID:   t.ChainID,
		Sequences: []int64{t.Sequence},
		Msg:       msg,
	}
	// sign
	sig := privKey.Sign(signMsg.Bytes())
	sigs := []sdk.StdSignature{{
		PubKey:    privKey.PubKey(),
		Signature: sig,
		Sequence:  t.Sequence,
	}}

	// marshal bytes
	tx := sdk.NewStdTx(signMsg.Msg, signMsg.Fee, sigs)
	txBytes, err := json.Marshal(tx)
	if err != nil {
		return nil, err
	}
	return txBytes, nil
}
