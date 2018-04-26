package broadcast

import (
	"github.com/lino-network/lino-go/transport"
	"github.com/tendermint/go-crypto"
)

func BroadcastTransaction(transaction interface{}, rawPrivKey [64]byte) {
	transport := transport.NewTransportFromViper()
	privKey := crypto.PrivKeyEd25519(rawPrivKey)
	_, err := transport.SignBuildBroadcast(transaction, privKey)
	if err != nil {
		panic(err)
	}
}
