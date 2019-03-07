// Package transport implements the functionalities that
// directly interact with the blockchain to query data or broadcast transaction.
package transport

import (
	"context"
	"fmt"

	wire "github.com/cosmos/cosmos-sdk/codec"
	"github.com/lino-network/lino-go/errors"
	"github.com/lino-network/lino-go/model"
	"github.com/spf13/viper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// Transport is a wrapper of tendermint rpc client and codec.
type Transport struct {
	chainId string
	nodeUrl string
	client  rpcclient.Client
	Cdc     *wire.Codec
}

// NewTransportFromConfig initiates an instance of Transport from config files.
func NewTransportFromConfig() *Transport {
	v := viper.New()
	viper.SetConfigType("json")
	v.SetConfigName("config")
	v.AddConfigPath("$HOME/.lino-go/")
	v.AutomaticEnv()
	v.ReadInConfig()

	nodeUrl := v.GetString("node_url")
	if nodeUrl == "" {
		nodeUrl = "localhost:26657"
	}
	rpc := rpcclient.NewHTTP(nodeUrl, "/websocket")
	return &Transport{
		chainId: v.GetString("chain_id"),
		nodeUrl: nodeUrl,
		client:  rpc,
		Cdc:     MakeCodec(),
	}
}

// NewTransportFromArgs initiates an instance of Transport from parameters passed in.
func NewTransportFromArgs(chainID, nodeUrl string) *Transport {
	if nodeUrl == "" {
		nodeUrl = "localhost:26657"
	}
	rpc := rpcclient.NewHTTP(nodeUrl, "/websocket")
	return &Transport{
		chainId: chainID,
		nodeUrl: nodeUrl,
		client:  rpc,
		Cdc:     MakeCodec(),
	}
}

// Query from Tendermint with the provided key and storename
func (t Transport) Query(ctx context.Context, storeName, subStore string, keys []string) (res []byte, err error) {
	finishChan := make(chan bool)
	go func() {
		res, err = t.query(keys, storeName, subStore, 0)
		finishChan <- true
	}()

	select {
	case <-finishChan:
		break
	case <-ctx.Done():
		return nil, errors.Timeout("query timeout").AddCause(ctx.Err())
	}

	return res, err
}

// Query from Tendermint with the provided key and storename at certain height
func (t Transport) QueryAtHeight(ctx context.Context, key cmn.HexBytes, storeName string, height int64) (res []byte, err error) {
	finishChan := make(chan bool)
	go func() {
		res, err = t.queryByKey(key, storeName, "key", height)
		finishChan <- true
	}()

	select {
	case <-finishChan:
		break
	case <-ctx.Done():
		return nil, errors.Timeoutf("query at height %v timeout", height).AddCause(ctx.Err())
	}

	return res, err
}

// Query from Tendermint with the provided subspace and storename
func (t Transport) QuerySubspace(ctx context.Context, subspace []byte, storeName string) (res []sdk.KVPair, err error) {
	var resRaw []byte
	finishChan := make(chan bool)
	go func() {
		resRaw, err = t.query([]string{string(subspace)}, storeName, "subspace", 0)
		finishChan <- true
	}()

	select {
	case <-finishChan:
		break
	case <-ctx.Done():
		return nil, errors.Timeout("query subspace timeout").AddCause(ctx.Err())
	}

	if err != nil {
		return nil, err
	}

	t.Cdc.UnmarshalJSON(resRaw, &res)
	return
}
func (t Transport) queryByKey(key cmn.HexBytes, storeName, substore string, height int64) (res []byte, err error) {
	path := fmt.Sprintf("/store/%s/key", storeName)
	node, err := t.GetNode()
	if err != nil {
		return res, err
	}

	opts := rpcclient.ABCIQueryOptions{
		Height: height,
		Prove:  false,
	}
	result, err := node.ABCIQueryWithOptions(path, key, opts)
	if err != nil {
		return res, err
	}

	resp := result.Response
	if resp.Code != uint32(0) {
		return res, errors.QueryFail("Query failed").AddBlockChainCode(resp.Code).AddBlockChainLog(resp.Log)
	}

	if resp.Value == nil || len(resp.Value) == 0 {
		return nil, errors.EmptyResponse("Empty response!")
	}

	return resp.Value, nil
}

func (t Transport) query(keys []string, storeName, substore string, height int64) (res []byte, err error) {
	path := fmt.Sprintf("/custom/%s/%s", storeName, substore)
	for _, key := range keys {
		path += ("/" + key)
	}
	node, err := t.GetNode()
	if err != nil {
		return res, err
	}

	opts := rpcclient.ABCIQueryOptions{
		Height: height,
		Prove:  false,
	}
	result, err := node.ABCIQueryWithOptions(path, []byte{}, opts)
	if err != nil {
		return res, err
	}

	resp := result.Response
	if resp.Code != uint32(0) {
		return res, errors.QueryFail("Query failed").AddBlockChainCode(resp.Code).AddBlockChainLog(resp.Log)
	}

	if resp.Value == nil || len(resp.Value) == 0 {
		return nil, errors.EmptyResponse("Empty response!")
	}

	return resp.Value, nil
}

// QueryBlock queries a block with a certain height from blockchain.
func (t Transport) QueryBlock(ctx context.Context, height int64) (res *ctypes.ResultBlock, err error) {
	node, err := t.GetNode()
	if err != nil {
		return res, err
	}

	finishChan := make(chan bool)
	go func() {
		res, err = node.Block(&height)
		finishChan <- true
	}()

	select {
	case <-finishChan:
		break
	case <-ctx.Done():
		return nil, errors.Timeout("query block timeout").AddCause(ctx.Err())
	}

	return res, err
}

// QueryBlockStatus queries block status from blockchain.
func (t Transport) QueryBlockStatus(ctx context.Context) (res *ctypes.ResultStatus, err error) {
	node, err := t.GetNode()
	if err != nil {
		return res, err
	}

	finishChan := make(chan bool)
	go func() {
		res, err = node.Status()
		finishChan <- true
	}()

	select {
	case <-finishChan:
		break
	case <-ctx.Done():
		return nil, errors.Timeout("query block status timeout").AddCause(ctx.Err())
	}

	return res, err
}

// QueryTx queries tx from blockchain.
func (t Transport) QueryTx(ctx context.Context, hash []byte) (res *ctypes.ResultTx, err error) {
	node, err := t.GetNode()
	if err != nil {
		return res, err
	}

	finishChan := make(chan bool)
	go func() {
		res, err = node.Tx(hash, false)
		finishChan <- true
	}()

	select {
	case <-finishChan:
		break
	case <-ctx.Done():
		return nil, errors.Timeout("query tx timeout").AddCause(ctx.Err())
	}

	return res, err
}

// BroadcastTx broadcasts a transcation to blockchain.
func (t Transport) BroadcastTx(tx []byte, checkTxOnly bool) (interface{}, error) {
	node, err := t.GetNode()
	if err != nil {
		return nil, err
	}

	if checkTxOnly {
		return node.BroadcastTxSync(tx)
	}
	return node.BroadcastTxCommit(tx)
}

// SignBuildBroadcast signs msg with private key and then broadcasts
// the transaction to blockchain.
func (t Transport) SignBuildBroadcast(msg model.Msg, privKeyHex string, seq int64, memo string, checkTxOnly bool) (interface{}, error) {
	msgs := []model.Msg{msg}

	privKey, err := GetPrivKeyFromHex(privKeyHex)
	if err != nil {
		return nil, err
	}

	signMsgBytes, err := EncodeSignMsg(t.Cdc, msgs, t.chainId, seq)
	if err != nil {
		return nil, err
	}
	// SignatureFromBytes
	sig, err := privKey.Sign(signMsgBytes)
	if err != nil {
		return nil, err
	}

	// build transaction bytes
	txByte, err := EncodeTx(t.Cdc, msgs, privKey.PubKey(), sig, seq, memo)
	if err != nil {
		return nil, err
	}

	// broadcast
	return t.BroadcastTx(txByte, checkTxOnly)
}

// GetNote returns the Tendermint rpc client node.
func (t Transport) GetNode() (rpcclient.Client, error) {
	if t.client == nil {
		return nil, errors.InvalidNodeURL("Must define node URL")
	}
	return t.client, nil
}
