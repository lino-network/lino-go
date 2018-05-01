package broadcast

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"

	"github.com/lino-network/lino-go/transport"
	"github.com/lino-network/lino-go/types"
)

func BroadcastTransaction(transaction interface{}, privKeyHex string, cb types.Callback) {
	go func(transaction interface{}, privKeyHex string, cb types.Callback) {
		transport := transport.NewTransportFromViper()
		res, err := transport.SignBuildBroadcast(transaction, privKeyHex, 0)

		var reg = regexp.MustCompile(`expected (\d+)`)
		var tries = 1

		// keep trying to get newest sequence number
		for err == nil && res.CheckTx.Code == types.InvalidSeqErrCode {
			match := reg.FindString(res.CheckTx.Log)
			seq, err := strconv.ParseInt(match[9:], 10, 64)
			if err != nil || tries == types.BroadcastMaxTries {
				fmt.Println("Get Sequence number failed ! ", err)
				return
			}
			res, err = transport.SignBuildBroadcast(transaction, privKeyHex, seq)
			tries += 1
		}

		if err != nil {
			fmt.Println("Build and Sign message failed ! ", err)
			return
		}
		if res.CheckTx.Code != uint32(0) {
			fmt.Println("CheckTx failed ! code: ", res.CheckTx.Code, res.CheckTx.Log)
			return
		}
		if res.DeliverTx.Code != uint32(0) {
			fmt.Println("DeliverTx failed ! code: ", res.DeliverTx.Code, res.DeliverTx.Log)
			return
		}
		fmt.Printf("Committed at block %d. Hash: %s\n", res.Height, res.Hash.String())

		result, _ := json.Marshal(res)
		cb(err, string(result))
	}(transaction, privKeyHex, cb)

}
