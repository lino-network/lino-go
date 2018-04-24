package broadcast

import (
	"github.com/lino-network/lino/app"

	"github.com/tendermint/go-crypto"
)

var (
	cdc = app.MakeCodec()
)

type BroadcastCallback func(result, error)

func NewLinoProxy() core.CoreContext {
	nodeURI := viper.GetString(client.FlagNode)
	var rpc rpcclient.Client
	if nodeURI != "" {
		rpc = rpcclient.NewHTTP(nodeURI, "/websocket")
	}
	return core.CoreContext{
		ChainID:         viper.GetString(client.FlagChainID),
		Height:          viper.GetInt64(client.FlagHeight),
		TrustNode:       viper.GetBool(client.FlagTrustNode),
		FromAddressName: viper.GetString(client.FlagName),
		NodeURI:         nodeURI,
		Sequence:        viper.GetInt64(client.FlagSequence),
		Client:          rpc,
	}
}

func BroadcastOperation(operation interface{}, rawPrivKey string, broadcastCallback BroadcastCallback) {
	privKey, err := crypto.PrivKeyFromBytes([]byte(rawPrivKey))
	if err != nil {
		broadcastCallback(result, err)
		return
	}

	switch op := operation.(type) {
	case TransferToAddress:
		amount, err := sdk.NewRatFromDecimal(op.Amount)
		if err != nil {
			broadcastCallback(result, err)
			return
		}
		msg := acc.NewTransferMsg(op.From, types.LNO(amount), op.Memo, acc.TransferToAddr(sdk.Address(op.ToAddress)))
		SignBuildBroadcast(msg, cdc, privKey)
	}
}

func (ctx CoreContext) SignBuildBroadcast(
	msg sdk.Msg, cdc *wire.Codec, privKey crypto.PrivKey) (*ctypes.ResultBroadcastTxCommit, error) {

	node, err := ctx.GetNode()
	if err != nil {
		return nil, err
	}

	res, err := node.BroadcastTxCommit(tx)
	if err != nil {
		return res, err
	}

	if res.CheckTx.Code != uint32(0) {
		return res, errors.Errorf("CheckTx failed: (%d) %s",
			res.CheckTx.Code,
			res.CheckTx.Log)
	}
	if res.DeliverTx.Code != uint32(0) {
		return res, errors.Errorf("DeliverTx failed: (%d) %s",
			res.DeliverTx.Code,
			res.DeliverTx.Log)
	}
	return res, err
}
