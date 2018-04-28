package main

import (
	"github.com/lino-network/lino-go/broadcast"
	"github.com/lino-network/lino-go/model"
)

func main() {
	privKeyBytes := [64]byte{107, 8, 232, 225, 15, 189, 53, 24, 10, 59, 201, 237, 229, 207, 233, 11, 166, 87, 140, 246, 215, 231, 64, 110, 64, 19, 94,
		75, 9, 209, 117, 38, 253, 242, 42, 63, 91, 251, 94, 100, 100, 140, 176, 157, 223, 124, 6, 155, 36, 106, 35, 53, 164, 115, 237, 3, 163, 116, 47,
		111, 230, 205, 0, 150}
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
	transferTx := model.TransferMsg{
		Sender:       "Lino",
		ReceiverAddr: "89920E0CF4C7910B54AB543B46F30ECAAA19EBF3",
		Amount:       "8888888",
	}
	broadcast.BroadcastTransaction(transferTx, privKeyBytes)
	broadcast.BroadcastTransaction(transferTx, privKeyBytes)

	// followTx := model.FollowMsg{
	// 	Follower: "Lino",
	// 	Followee: "Zhimao",
	// }
	// broadcast.BroadcastTransaction(followTx, privKeyBytes)
	//
	// unfollowTx := model.UnfollowMsg{
	// 	Follower: "Lino",
	// 	Followee: "Zhimao",
	// }
	// broadcast.BroadcastTransaction(unfollowTx, privKeyBytes)

}
