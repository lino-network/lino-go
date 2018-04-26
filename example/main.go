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
		From:      "Lino",
		ToAddress: "89920E0CF4C7910B54AB543B46F30ECAAA19EBF3",
		Amount:    "8888888",
	}

	broadcast.BroadcastTransaction(transferTx, "dummy")

}
