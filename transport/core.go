package transport

import (
	"fmt"

	"github.com/tendermint/go-crypto"

	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	cmn "github.com/tendermint/tmlibs/common"
)

type Transport struct {
	chainId string
	nodeUrl string
	client  rpcclient.Client
	Cdc     *wire.Codec
}

func NewTransportFromViper() Transport {
	v := viper.New()
	viper.SetConfigType("json")
	v.SetConfigName("config")
	v.AddConfigPath("$GOPATH/src/github.com/lino-network/lino-go/")
	v.AutomaticEnv()
	v.ReadInConfig()

	var rpc rpcclient.Client
	nodeUrl := v.GetString("node_url")
	if nodeUrl != "" {
		rpc = rpcclient.NewHTTP(nodeUrl, "/websocket")
	}
	return Transport{
		chainId: v.GetString("chain_id"),
		nodeUrl: nodeUrl,
		client:  rpc,
		Cdc:     MakeCodec(),
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

func (t Transport) BroadcastTx(tx []byte) (*ctypes.ResultBroadcastTxCommit, error) {
	node, err := t.GetNode()
	if err != nil {
		return nil, err
	}

	res, err := node.BroadcastTxCommit(tx)
	if err != nil {
		return res, err
	}
	return res, err
}

func (t Transport) SignBuildBroadcast(msg interface{},
	privKey crypto.PrivKey, seq int64) (*ctypes.ResultBroadcastTxCommit, error) {
	signMsgBytes, err := EncodeSignMsg(t.Cdc, msg, t.chainId, seq)
	if err != nil {
		return nil, err
	}
	// sign
	sig := privKey.Sign(signMsgBytes)

	// build transaction bytes
	txBytes, err := EncodeTx(t.Cdc, msg, privKey.PubKey(), sig, seq)
	if err != nil {
		return nil, err
	}
	// broadcast
	return t.BroadcastTx(txBytes)
}

func (t Transport) GetNode() (rpcclient.Client, error) {
	if t.client == nil {
		return nil, errors.New("Must define node URI")
	}
	return t.client, nil
}
