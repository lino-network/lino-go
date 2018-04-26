package broadcast

import (
	"github.com/lino-network/lino-go/transport"
	"github.com/tendermint/go-crypto"
)

func BroadcastTransaction(transaction interface{}, rawPrivKey string) {
	keyBytes := [64]byte{64, 54, 172, 112, 137, 204, 17, 93, 138, 33, 150, 34, 13, 26, 206, 98, 121,
		142, 98, 243, 170, 131, 83, 248, 49, 121, 109, 20, 216, 134, 175, 170, 218, 131, 39, 50, 79, 90, 236,
		79, 2, 188, 19, 254, 218, 228, 6, 188, 143, 151, 41, 29, 234, 237, 110, 228, 216, 25, 59, 78, 113, 76, 38, 134}

	transport := transport.NewTransportFromViper()
	privKey := crypto.PrivKeyEd25519(keyBytes)
	_, err := transport.SignBuildBroadcast(transaction, privKey)
	if err != nil {
		panic(err)
	}
}
