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
	transferTx := model.TransferToAddressMsg{
		Sender:       "Lino",
		ReceiverAddr: "89920E0CF4C7910B54AB543B46F30ECAAA19EBF3",
		Amount:       "8888888",
	}

	keyBytes := [64]byte{64, 54, 172, 112, 137, 204, 17, 93, 138, 33, 150, 34, 13, 26, 206, 98, 121,
		142, 98, 243, 170, 131, 83, 248, 49, 121, 109, 20, 216, 134, 175, 170, 218, 131, 39, 50, 79, 90, 236,
		79, 2, 188, 19, 254, 218, 228, 6, 188, 143, 151, 41, 29, 234, 237, 110, 228, 216, 25, 59, 78, 113, 76, 38, 134}
	broadcast.BroadcastTransaction(transferTx, keyBytes)

}
