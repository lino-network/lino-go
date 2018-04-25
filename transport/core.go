package transport

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	cmn "github.com/tendermint/tmlibs/common"
)

type Transport struct {
	ChainID  string
	NodeURI  string
	Sequence int64
	Client   rpcclient.Client
}

func NewTransportFromViper() Transport {
	var rpc rpcclient.Client
	nodeURI := "localhost:46657"
	if nodeURI != "" {
		rpc = rpcclient.NewHTTP(nodeURI, "/websocket")
	}
	return Transport{
		ChainID:  viper.GetString("test"),
		NodeURI:  nodeURI,
		Sequence: viper.GetInt64("0"),
		Client:   rpc,
	}
}

func (t Transport) Query(key cmn.HexBytes, storeName string) (res []byte, err error) {
	path := fmt.Sprintf("/%s/key", storeName)
	node, err := t.GetNode()
	if err != nil {
		return res, err
	}

	result, err := node.ABCIQuery(path, key)
	if err != nil {
		return res, err
	}
	resp := result.Response
	if resp.Code != uint32(0) {
		return res, errors.Errorf("Query failed: (%d) %s", resp.Code, resp.Log)
	}
	return resp.Value, nil
}

// func (t Transport) BroadcastTx(tx []byte) (*ctypes.ResultBroadcastTxCommit, error) {
// 	node, err := t.GetNode()
// 	if err != nil {
// 		return nil, err
// 	}

// 	res, err := node.BroadcastTxCommit(tx)
// 	if err != nil {
// 		return res, err
// 	}

// 	if res.CheckTx.Code != uint32(0) {
// 		return res, errors.Errorf("CheckTx failed: (%d) %s",
// 			res.CheckTx.Code,
// 			res.CheckTx.Log)
// 	}
// 	if res.DeliverTx.Code != uint32(0) {
// 		return res, errors.Errorf("DeliverTx failed: (%d) %s",
// 			res.DeliverTx.Code,
// 			res.DeliverTx.Log)
// 	}
// 	return res, err
// }

// func (t Transport) SignBuildBroadcast(
// 	msg sdk.Msg, cdc *wire.Codec, privKey crypto.PrivKey) (*ctypes.ResultBroadcastTxCommit, error) {
// 	// build
// 	signMsg := sdk.StdSignMsg{
// 		ChainID:   t.ChainID,
// 		Sequences: []int64{t.Sequence},
// 		Msg:       msg,
// 	}

// 	keybase, err := keys.GetKeyBase()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// sign
// 	bz := signMsg.Bytes()
// 	sig, pubkey, err := keybase.Sign(name, passphrase, bz)
// 	if err != nil {
// 		return nil, err
// 	}

// 	sigs := []sdk.StdSignature{{
// 		PubKey:    pubkey,
// 		Signature: sig,
// 		Sequence:  t.Sequence,
// 	}}

// 	// marshal bytes
// 	tx := sdk.NewStdTx(signMsg.Msg, signMsg.Fee, sigs)
// 	txBytes, err := cdc.MarshalBinary(tx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// broadcast
// 	return ctx.BroadcastTx(txBytes)
// }

func (t Transport) GetNode() (rpcclient.Client, error) {
	if t.Client == nil {
		return nil, errors.New("Must define node URI")
	}
	return t.Client, nil
}
