package transport

import (
	"fmt"

	"github.com/pkg/errors"
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
}

func NewTransportFromViper() Transport {
	var rpc rpcclient.Client
	nodeUrl := "localhost:46657"
	if nodeUrl != "" {
		rpc = rpcclient.NewHTTP(nodeUrl, "/websocket")
	}
	return Transport{
		chainId:  "test-chain-iiMiw7",
		nodeUrl:  nodeUrl,
		sequence: 4,
		client:   rpc,
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

func (t Transport) SignBuildBroadcast(msg interface{},
	privKey crypto.PrivKey) (*ctypes.ResultBroadcastTxCommit, error) {
	// build sign msg bytes
	// msgBytes, err := EncodeMsg(msg)
	// if err != nil {
	// 	panic(err)
	// }
	signMsgBytes, err := EncodeSignMsg(msg, t.chainId, t.sequence)
	if err != nil {
		panic(err)
	}
	// sign
	sig := privKey.Sign(signMsgBytes)

	fmt.Println("sign message after marshal json ", string(signMsgBytes))
	fmt.Println("sig is ", sig)

	// build transaction bytes
	txBytes, err := EncodeTx(msg, privKey.PubKey(), sig, t.sequence)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(txBytes))
	// broadcast
	return t.BroadcastTx(txBytes)
}

func (t Transport) GetNode() (rpcclient.Client, error) {
	if t.client == nil {
		return nil, errors.New("Must define node URI")
	}
	return t.client, nil
}

//right
// message is  {"chain_id":"test-chain-iiMiw7","sequences":[2],
// 	"fee_bytes":"eyJBbW91bnQiOltdLCJHYXMiOjB9",
// 	"msg_bytes":"eyJzZW5kZXIiOiJMaW5vIiwicmVjZWl2ZXJfbmFtZSI6IiIsInJlY2VpdmVyX2FkZHIiOiI4OTkyMEUwQ0Y0Qzc5MTBCNTRBQjU0M0I0NkYzMEVDQUFBMTlFQkYzIiwiYW1vdW50IjoiODg4ODg4OCIsIm1lbW8iOiIifQ==",
// 	"alt_bytes":null}
// sig is  {/70A3C9D6D81D.../}

// me

// me
