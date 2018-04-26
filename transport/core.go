package transport

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	crypto "github.com/tendermint/go-crypto"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	cmn "github.com/tendermint/tmlibs/common"
)

type Transport struct {
	chainId  string
	nodeUrl  string
	sequence int64
	client   rpcclient.Client
	cdc      *wire.Codec
}

func NewTransportFromViper() Transport {
	var rpc rpcclient.Client
	nodeUrl := "localhost:46657"
	if nodeUrl != "" {
		rpc = rpcclient.NewHTTP(nodeUrl, "/websocket")
	}
	return Transport{
		chainId:  viper.GetString("test"),
		nodeUrl:  nodeUrl,
		sequence: viper.GetInt64("0"),
		client:   rpc,
		cdc:      MakeCodec(),
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

func (t Transport) SignBuildBroadcast(transaction interface{},
	privKey crypto.PrivKey) (*ctypes.ResultBroadcastTxCommit, error) {
	// build sign msg bytes
	signMsgBytes, _ := t.BuildSignMsg(transaction)

	// sign
	sig := privKey.Sign(signMsgBytes)
	sigs := []sdk.StdSignature{{
		PubKey:    privKey.PubKey(),
		Signature: sig,
		Sequence:  t.sequence,
	}}

	// build transaction bytes
	txBytes, _ := t.BulidTx(transaction, sigs)

	// StdTx.Msg is an interface.
	//cdc := wire.NewCodec()
	//txBytes, err := json.Marshal(txBytes)

	// broadcast
	return t.BroadcastTx(txBytes)
}

func (t Transport) BuildSignMsg(transaction interface{}) ([]byte, error) {
	msgBytes, err := json.Marshal(transaction)
	if err != nil {
		panic(err)
	}
	bz, err := json.Marshal(sdk.StdSignDoc{
		ChainID:   t.chainId,
		Sequences: []int64{t.sequence},
		MsgBytes:  msgBytes,
	})
	if err != nil {
		panic(err)
	}
	return bz, nil
}

func (t Transport) BulidTx(transaction interface{}, sigs []sdk.StdSignature) ([]byte, error) {
	tx := sdk.StdTx{
		//Transaction: transaction,
		Signatures: sigs,
	}

	txBytes, err := t.cdc.MarshalBinary(tx)
	if err != nil {
		return nil, err
	}
	return txBytes, nil
}

func (t Transport) GetNode() (rpcclient.Client, error) {
	if t.client == nil {
		return nil, errors.New("Must define node URI")
	}
	return t.client, nil
}

// type StdTx struct {
// 	//Transaction interface{}        `json:"msg2"`
// 	Signatures []sdk.StdSignature `json:"signatures"`
// }
