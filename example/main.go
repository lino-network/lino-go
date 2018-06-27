package main

import (
	"encoding/json"
	"fmt"

	"github.com/lino-network/lino-go/api"
)

//"encoding/hex"

//"strings"

//crypto "github.com/tendermint/go-crypto"

func main() {
	// // query example

	api := api.NewLinoAPIFromConfig()
	// for {
	// 	res, _ := query.GetAllValidators()
	// 	output, _ := json.Marshal(res)
	// 	fmt.Println(string(output))

	// 	res1, _ := query.GetValidator("Lino1")
	// 	output, _ = json.Marshal(res1)
	// 	fmt.Println(string(output))

	// 	res11, _ := query.GetValidator("Lino2")
	// 	output, _ = json.Marshal(res11)
	// 	fmt.Println(string(output))

	// 	res111, _ := query.GetValidator("Lino3")
	// 	output, _ = json.Marshal(res111)
	// 	fmt.Println(string(output))

	// 	res1111, _ := query.GetValidator("Lino4")
	// 	output, _ = json.Marshal(res1111)
	// 	fmt.Println(string(output))

	// 	fmt.Println("--------------------------------------------------------------------")
	// }
	// res2, _ := query.GetDeveloper("Lino1")
	// output, _ = json.Marshal(res2)
	// fmt.Println(string(output))

	// res3, _ := query.GetDevelopers()
	// output, _ = json.Marshal(res3)
	// fmt.Println(string(output))
	//
	// res4, _ := query.GetInfraProvider("Lino")
	// output, _ = json.Marshal(res4)
	// fmt.Println(string(output))
	//
	// res5, _ := query.GetInfraProviders()
	// output, _ = json.Marshal(res5)
	// fmt.Println(string(output))
	//
	// res6, _ := query.GetAccountBank("61DC349C3802CB3B3884B4450866B2EFF41B75B4")
	// output, _ = json.Marshal(res6)
	// fmt.Println(string(output))
	//
	res7, _ := api.Query.GetAccountMeta("lino")
	output, _ := json.Marshal(res7)
	fmt.Println(string(output))
	//
	// res8 := query.GetAccountSequence("Lino")
	// output, _ = json.Marshal(res8)
	// fmt.Println(string(output))
	//
	// res9, _ := query.GetVoter("Lino1")
	// output, _ = json.Marshal(res9)
	// fmt.Println(string(output))
	// //
	res10, _ := api.Query.GetAccountBank("lino")
	output, _ = json.Marshal(res10)
	fmt.Println(string(output))

	// block, err := api.Query.GetBlock(121)
	// fmt.Println(err)
	// fmt.Println(block)
	//
	// res11, _ := query.GetGrantList("Lino")
	// output, _ = json.Marshal(res11)
	// fmt.Println(string(output))

	// res12, _ := query.GetPostInfo("Lino", "test12")
	// output, _ = json.Marshal(res12)
	// fmt.Println(string(output))
	//
	// res13, _ := query.GetPostMeta("Lino", "test12")
	// output, _ = json.Marshal(res13)
	// fmt.Println(string(output))

	//broadcast ransaction example
	// user := "yukai-tu"
	// masterPriv := crypto.GenPrivKeyEd25519()
	// txPriv := crypto.GenPrivKeyEd25519()
	// postPriv := crypto.GenPrivKeyEd25519()

	// masterPub := masterPriv.PubKey()
	// txPub := txPriv.PubKey()
	// postPub := postPriv.PubKey()

	// addr := masterPub.Address()

	// addrHex := strings.ToUpper(hex.EncodeToString(addr))
	// masterPrivHex := hex.EncodeToString(masterPriv.Bytes())
	// masterPubHex := hex.EncodeToString(masterPub.Bytes())
	// txPubHex := hex.EncodeToString(txPub.Bytes())
	// postPubHex := hex.EncodeToString(postPub.Bytes())

	// linoTxPriv := "A32889124067E8FDA45CB7FC07C4DE02E6E78F46A82A9D40FB41024C219EE5A21852E84F20638D16DEA1030F8DD58270638048080C699DA02342420155693330E1FF4272BD"

	// links := map[string]string{}
	// err := broadcast.CreatePost("test12", "a test", "dummy", "Lino", "", "", "", "", "0", linoTxPriv, links)
	// if err != nil {
	// 	panic(err)
	// }

	// err = broadcast.Transfer("Lino", "", addrHex, "10000", "", linoTxPriv)
	// if err != nil {
	// 	panic(err)
	// }
	// err = broadcast.Register(user, masterPubHex, postPubHex, txPubHex, masterPrivHex)
	// if err != nil {
	// 	panic(err)
	// }

}
