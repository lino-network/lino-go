package broadcast

// import (
// 	"github.com/tendermint/go-crypto"
// )

// var (
// 	cdc = app.MakeCodec()
// )

// type BroadcastCallback func(result, error)

// func BroadcastOperation(operation interface{}, rawPrivKey string, broadcastCallback BroadcastCallback) {
// 	privKey, err := crypto.PrivKeyFromBytes([]byte(rawPrivKey))
// 	if err != nil {
// 		broadcastCallback(result, err)
// 		return
// 	}

// 	switch op := operation.(type) {
// 	case TransferToAddress:
// 		amount, err := sdk.NewRatFromDecimal(op.Amount)
// 		if err != nil {
// 			broadcastCallback(result, err)
// 			return
// 		}
// 		msg := acc.NewTransferMsg(op.From, types.LNO(amount), op.Memo, acc.TransferToAddr(sdk.Address(op.ToAddress)))
// 		SignBuildBroadcast(msg, cdc, privKey)
// 	}
// }


