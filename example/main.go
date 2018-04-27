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

	privKeyBytes := [64]byte{114, 23, 52, 48, 243, 151, 28, 4, 144, 106, 26, 13, 109, 194, 18, 91, 181, 192, 228, 82, 232, 243, 134, 250, 3, 160, 178, 183, 87, 133, 44, 106, 74, 231, 2, 29, 195,
		30, 100, 255, 99, 249, 242, 179, 178, 27, 2, 173, 85, 163, 186, 130, 131, 141, 104, 110, 211, 123, 90, 233, 141, 24, 181, 203}
	broadcast.BroadcastTransaction(transferTx, privKeyBytes)

}
