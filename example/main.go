package main

import (
	"encoding/json"
	"fmt"

	"github.com/lino-network/lino-go/broadcast"
	"github.com/lino-network/lino-go/model"
	"github.com/lino-network/lino-go/query"
)

func main() {

	// query example
	res, _ := query.GetAllValidators()
	output, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println(string(output))

	res1, _ := query.GetValidator("Lino")
	output, _ = json.MarshalIndent(res1, "", "  ")
	fmt.Println(string(output))

	// broadcast ransaction example
	transferTx := model.TransferToAddress{
		Sender:       "Lino",
		ReceiverAddr: "89920E0CF4C7910B54AB543B46F30ECAAA19EBF3",
		Amount:       "8888888",
	}

	broadcast.BroadcastTransaction(transferTx, "dummy")

}

//
// {"msg":[4,{"sender":"Lino","receiver_name":"","receiver_addr":"89920E0CF4C7910B54AB543B46F30ECAAA19EBF3","amount":"1000","memo":""}],
// "fee":{"Amount":[],"Gas":0},
// "signatures":[{"pub_key":[1,"DA8327324F5AEC4F02BC13FEDAE406BC8F97291DEAED6EE4D8193B4E714C2686"],
// 							"signature":[1,"28500E66C3435EFC5E82FFD60ACD2FDB0F10BB561B10911F20AAEAF5C372779F66E7E0E487F6C7603D3ED249419E55793C4C48CBDB29840AA0659E542D8B5607"],
// 							"sequence":0}]}
