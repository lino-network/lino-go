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

	//broadcast ransaction example
	transferTx := model.TransferToAddressMsg{
		Sender:       "Lino",
		ReceiverAddr: "89920E0CF4C7910B54AB543B46F30ECAAA19EBF3",
		Amount:       "8888888",
	}

	privKeyBytes := [64]byte{107, 8, 232, 225, 15, 189, 53, 24, 10, 59, 201, 237, 229, 207, 233, 11, 166, 87, 140, 246, 215, 231, 64, 110, 64, 19, 94,
		75, 9, 209, 117, 38, 253, 242, 42, 63, 91, 251, 94, 100, 100, 140, 176, 157, 223, 124, 6, 155, 36, 106, 35, 53, 164, 115, 237, 3, 163, 116, 47, 111, 230, 205, 0, 150}
	broadcast.BroadcastTransaction(transferTx, privKeyBytes)

}
