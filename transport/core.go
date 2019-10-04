// Package transport implements the functionalities that
// directly interact with the blockchain to query data or broadcast transaction.
package transport

import (
	"context"
	"fmt"

	"github.com/lino-network/lino-go/errors"
	linoapp "github.com/lino-network/lino/app"
	"github.com/spf13/viper"

	wire "github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	crypto "github.com/tendermint/tendermint/crypto"
	cmn "github.com/tendermint/tendermint/libs/common"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// Transport is a wrapper of tendermint rpc client and codec.
type Transport struct {
	chainId      string
	nodeUrl      string
	maxFeeInCoin int64
	client       rpcclient.Client
	Cdc          *wire.Codec
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
		Cdc:     linoapp.MakeCodec(),
	}
}

// NewTransportFromArgs initiates an instance of Transport from parameters passed in.
func NewTransportFromArgs(chainID, nodeUrl string, maxFeeInCoin int64) *Transport {
	if nodeUrl == "" {
		nodeUrl = "localhost:26657"
	}
	rpc := rpcclient.NewHTTP(nodeUrl, "/websocket")
	return &Transport{
		chainId:      chainID,
		nodeUrl:      nodeUrl,
		client:       rpc,
		Cdc:          linoapp.MakeCodec(),
		maxFeeInCoin: maxFeeInCoin,
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
// func (t Transport) SignBuildBroadcast(msg model.Msg, privKeyHex string, seq uint64, memo string, checkTxOnly bool) (*model.BroadcastResponse, errors.Error) {
// 	msgs := []model.Msg{msg}

// 	privKey, err := GetPrivKeyFromHex(privKeyHex)
// 	if err != nil {
// 		return nil, errors.FailedToBroadcastf("error to get private key from public key, err: %s", err.Error())
// 	}

// 	signMsgBytes, err := EncodeSignMsg(t.Cdc, msgs, t.chainId, seq, memo)
// 	if err != nil {
// 		return nil, errors.FailedToBroadcastf("error to encode sign msg, err: %s", err.Error())
// 	}
// 	// SignatureFromBytes
// 	sig, err := privKey.Sign(signMsgBytes)
// 	if err != nil {
// 		return nil, errors.FailedToBroadcastf("error to sign the msg, err: %s", err.Error())
// 	}

// 	// build transaction bytes
// 	txByte, err := EncodeTx(t.Cdc, msgs, privKey.PubKey(), sig, seq, memo)
// 	if err != nil {
// 		return nil, errors.FailedToBroadcastf("error to encode transaction, err: %s", err.Error())
// 	}

// 	broadcastResp := &model.BroadcastResponse{
// 		CommitHash: hex.EncodeToString(ttypes.Tx(txByte).Hash()),
// 	}

// 	res, err := t.BroadcastTx(txByte, checkTxOnly)
// 	if err != nil {
// 		return broadcastResp, errors.FailedToBroadcastf("broadcast tx failed, err: %s", err.Error())
// 	}

// 	if checkTxOnly {
// 		res, ok := res.(*ctypes.ResultBroadcastTx)
// 		if !ok {
// 			return broadcastResp, errors.FailedToBroadcast("error to parse the broadcast response")
// 		}
// 		code := retrieveCodeFromBlockChainCode(res.Code)
// 		if err == nil && code == model.InvalidSeqErrCode {
// 			return broadcastResp, errors.InvalidSequenceNumber("invalid seq").AddBlockChainCode(res.Code).AddBlockChainLog(res.Log)
// 		}

// 		if res.Code != uint32(0) {
// 			return broadcastResp, errors.CheckTxFail("CheckTx failed!").AddBlockChainCode(res.Code).AddBlockChainLog(res.Log)
// 		}
// 		if res.Code != uint32(0) {
// 			return broadcastResp, errors.DeliverTxFail("DeliverTx failed!").AddBlockChainCode(res.Code).AddBlockChainLog(res.Log)
// 		}
// 	} else {
// 		res, ok := res.(*ctypes.ResultBroadcastTxCommit)
// 		if !ok {
// 			return broadcastResp, errors.FailedToBroadcast("error to parse the broadcast response")
// 		}
// 		code := retrieveCodeFromBlockChainCode(res.CheckTx.Code)
// 		if err == nil && code == model.InvalidSeqErrCode {
// 			return nil, errors.InvalidSequenceNumber("invalid seq").AddBlockChainCode(res.CheckTx.Code).AddBlockChainLog(res.CheckTx.Log)
// 		}

// 		if res.CheckTx.Code != uint32(0) {
// 			return nil, errors.CheckTxFail("CheckTx failed!").AddBlockChainCode(res.CheckTx.Code).AddBlockChainLog(res.CheckTx.Log)
// 		}
// 		if res.DeliverTx.Code != uint32(0) {
// 			return nil, errors.DeliverTxFail("DeliverTx failed!").AddBlockChainCode(res.DeliverTx.Code).AddBlockChainLog(res.DeliverTx.Log)
// 		}
// 		broadcastResp.Height = res.Height
// 	}

// 	return broadcastResp, nil
// }

// SignAndBuild signs msg with private key and return tx bytes
func (t Transport) SignAndBuild(msg sdk.Msg, privKeyHex string, seq uint64, memo string) ([]byte, errors.Error) {
	msgs := []sdk.Msg{msg}

	privKey, err := GetPrivKeyFromHex(privKeyHex)
	if err != nil {
		return nil, errors.FailedToBroadcastf("error to get private key from public key, err: %s", err.Error())
	}

	signMsgBytes := EncodeSignMsg(t.Cdc, msgs, t.chainId, seq, memo, t.maxFeeInCoin)
	// SignatureFromBytes
	sig, err := privKey.Sign(signMsgBytes)
	if err != nil {
		return nil, errors.FailedToBroadcastf("error to sign the msg, err: %s", err.Error())
	}

	// build transaction bytes
	txByte, err := EncodeTx(t.Cdc, msgs, []crypto.PubKey{privKey.PubKey()}, [][]byte{sig}, memo, t.maxFeeInCoin)
	if err != nil {
		return nil, errors.FailedToBroadcastf("error to encode transaction, err: %s", err.Error())
	}
	return txByte, nil
}

// SignAndBuildMultiSig signs msg with multiple private key and return tx bytes
func (t Transport) SignAndBuildMultiSig(msg sdk.Msg, privKeyHexs []string, seqs []uint64, memo string) ([]byte, errors.Error) {
	msgs := []sdk.Msg{msg}

	pubKeys := []crypto.PubKey{}
	sigs := [][]byte{}
	for i, privKeyHex := range privKeyHexs {
		privKey, err := GetPrivKeyFromHex(privKeyHex)
		if err != nil {
			return nil, errors.FailedToBroadcastf("error to get private key from public key, err: %s", err.Error())
		}

		signMsgBytes := EncodeSignMsg(t.Cdc, msgs, t.chainId, seqs[i], memo, t.maxFeeInCoin)
		// SignatureFromBytes
		sig, err := privKey.Sign(signMsgBytes)
		if err != nil {
			return nil, errors.FailedToBroadcastf("error to sign the msg, err: %s", err.Error())
		}
		pubKeys = append(pubKeys, privKey.PubKey())
		sigs = append(sigs, sig)
	}

	// build transaction bytes
	txByte, err := EncodeTx(t.Cdc, msgs, pubKeys, sigs, memo, t.maxFeeInCoin)
	if err != nil {
		return nil, errors.FailedToBroadcastf("error to encode transaction, err: %s", err.Error())
	}
	return txByte, nil
}

// GetNote returns the Tendermint rpc client node.
func (t Transport) GetNode() (rpcclient.Client, error) {
	if t.client == nil {
		return nil, errors.InvalidNodeURL("Must define node URL")
	}
	return t.client, nil
}

func retrieveCodeFromBlockChainCode(bcCode uint32) uint32 {
	return bcCode & 0xff
}
