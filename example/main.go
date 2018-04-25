package main

import (
	"encoding/json"
	"fmt"

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

	// transaction example
	transferTx := model.TransferToAddress{}

}
