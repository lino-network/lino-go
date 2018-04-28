package broadcast

import (
	"fmt"

	"github.com/lino-network/lino-go/query"
	"github.com/lino-network/lino-go/transport"
	"github.com/tendermint/go-crypto"
)

func BroadcastTransaction(transaction interface{}, rawPrivKey [64]byte) {
	transport := transport.NewTransportFromViper()
	privKey := crypto.PrivKeyEd25519(rawPrivKey)
	seq, err := query.GetAccountSequence("Lino")
	if err != nil {
		panic(err)
	}
	res, err := transport.SignBuildBroadcast(transaction, privKey, seq)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Committed at block %d. Hash: %s\n", res.Height, res.Hash.String())
}
