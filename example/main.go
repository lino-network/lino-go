package main

import (
	"github.com/lino-network/lino-go/broadcast"
	"github.com/lino-network/lino-go/model"
)

func main() {
	// // query example
	// res, _ := query.GetAllValidators()
	// output, _ := json.MarshalIndent(res, "", "  ")
	// fmt.Println(string(output))
	//
	// res1, _ := query.GetValidator("Lino")
	// output, _ = json.MarshalIndent(res1, "", "  ")
	// fmt.Println(string(output))
	//
	// res2, _ := query.GetDeveloper("Lino")
	// output, _ = json.MarshalIndent(res2, "", "  ")
	// fmt.Println(string(output))
	//
	// res3, _ := query.GetDevelopers()
	// output, _ = json.MarshalIndent(res3, "", "  ")
	// fmt.Println(string(output))
	//
	// res4, _ := query.GetInfraProvider("Lino")
	// output, _ = json.MarshalIndent(res4, "", "  ")
	// fmt.Println(string(output))
	//
	// res5, _ := query.GetInfraProviders()
	// output, _ = json.MarshalIndent(res5, "", "  ")
	// fmt.Println(string(output))
	//
	// res6, _ := query.GetAccountBank("178BDA21C11D86E9F751639092B595DA7EF55B8A")
	// output, _ = json.MarshalIndent(res6, "", "  ")
	// fmt.Println(string(output))
	//
	// res7, _ := query.GetAccountMeta("Lino")
	// output, _ = json.MarshalIndent(res7, "", "  ")
	// fmt.Println(string(output))
	//
	// res8, _ := query.GetAccountSequence("Lino")
	// output, _ = json.MarshalIndent(res8, "", "  ")
	// fmt.Println(string(output))
	//
	// res9, _ := query.GetVoter("Lino")
	// output, _ = json.MarshalIndent(res9, "", "  ")
	// fmt.Println(string(output))
	//

	//broadcast ransaction example
	LinoPrivKey := "016b08e8e10fbd35180a3bc9ede5cfe90ba6578cf6d7e7406e40135e4b09d17526fdf22a3f5bfb5e64648cb09ddf7c069b246a2335a473ed03a3742f6fe6cd0096"
	transferTx := model.TransferMsg{
		Sender:       "Lino",
		ReceiverAddr: "89920E0CF4C7910B54AB543B46F30ECAAA19EBF3",
		Amount:       "8888888",
	}
	broadcast.BroadcastTransaction(transferTx, LinoPrivKey)

}
