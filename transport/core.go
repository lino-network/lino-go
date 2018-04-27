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
		sequence: 2,
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
	msgBytes, err := EncodeMsg(msg)
	if err != nil {
		panic(err)
	}
	signMsgBytes, err := EncodeSignMsg(msgBytes, t.chainId, t.sequence)
	if err != nil {
		panic(err)
	}
	// sign
	sig := privKey.Sign(signMsgBytes)

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

// {"msg":[4,{"sender":"Lino","receiver_addr":"89920E0CF4C7910B54AB543B46F30ECAAA19EBF3","amount":"8888888","memo":""}],
// "signatures":[{"pub_key":[1,"4ae7021dc31e64ff63f9f2b3b21b02ad55a3ba82838d686ed37b5ae98d18b5cb"],
// 								"signature":[1,"f6d12a7c63aa595a9e088f51330f3de61b4d8babb2ccd61ed4e164dc017c287a4d89d68a63f65a812269a8c35c0d939c983c490a69fee36bb48b57d6e2941b0b"],
// 								"sequence":1}]}
//
//
// {"msg":[4,{"sender":"Lino","receiver_name":"","receiver_addr":"89920E0CF4C7910B54AB543B46F30ECAAA19EBF3","amount":"8888888","memo":""}],
// "fee":{"Amount":[],"Gas":0},
// "signatures":[{"pub_key":[1,"4AE7021DC31E64FF63F9F2B3B21B02AD55A3BA82838D686ED37B5AE98D18B5CB"],
// 	"signature":[1,"7E9BA04B446F033E15A13C86F9BDE68C5E060BBC7DC11130C99C46C354A19705966CF6CD245E62A9E9BF0027333FBFC05A250096C0AF8B6C23FE27D70A44620D"],"sequence":1}]}
