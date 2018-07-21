package main

import (
	"fmt"

	"github.com/lino-network/lino-go/api"
	"github.com/lino-network/lino-go/model"
)

//"encoding/hex"

//"strings"

//crypto "github.com/tendermint/go-crypto"

func main() {
	// // query example

	//api := api.NewLinoAPIFromConfig()
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
	// //
	// res7, _ := api.Query.GetAccountMeta("lino")
	// output, _ := json.Marshal(res7)
	// fmt.Println(string(output))
	// //
	// // res8 := query.GetAccountSequence("Lino")
	// // output, _ = json.Marshal(res8)
	// // fmt.Println(string(output))
	// //
	// // res9, _ := query.GetVoter("Lino1")
	// // output, _ = json.Marshal(res9)
	// // fmt.Println(string(output))
	// // //
	// res10, _ := api.Query.GetAccountBank("lino")
	// output, _ = json.Marshal(res10)
	// fmt.Println(string(output))

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
	// broadcast := broadcast.NewBroadcast(transport.NewTransportFromConfig())
	// user := "yukai-tu13"
	// recoveryPriv := crypto.GenPrivKeyEd25519()
	// txPriv := crypto.GenPrivKeyEd25519()
	// fmt.Println("tx priv:", strings.ToUpper(hex.EncodeToString(txPriv.Bytes())))
	// micropaymentPriv := crypto.GenPrivKeyEd25519()
	// postPriv := crypto.GenPrivKeyEd25519()

	// err := broadcast.Register(
	// 	"lino", "100000000", user, hex.EncodeToString(recoveryPriv.PubKey().Bytes()), hex.EncodeToString(txPriv.PubKey().Bytes()),
	// 	hex.EncodeToString(micropaymentPriv.PubKey().Bytes()), hex.EncodeToString(postPriv.PubKey().Bytes()), "E1B0F79A207610DF4B9AA480C78F06C5B505B6F56A9B57A8CA05DCA868A41A95B664E319C9", 22)
	// fmt.Println(err)

	// // time.Sleep(3 * time.Second)
	// // err = broadcast.CreatePost(user, "lino", "title", "content", "", "", "", "", "0", nil, strings.ToUpper(hex.EncodeToString(txPriv.Bytes())), 0)
	// // fmt.Println(err)

	// time.Sleep(3 * time.Second)
	// err = broadcast.GrantPermission(user, "lino", 1000000, model.PostPermission, 10, hex.EncodeToString(txPriv.Bytes()), 0)
	// fmt.Println(err)

	// posts, _ := api.Query.GetUserAllPosts(user)
	// fmt.Println(posts)

	// postPub, _ := api.Query.GetPostPubKey("lino")
	// fmt.Println(postPub)
	// pubKeyToGrantPubKeyMap, _ := api.Query.GetAllGrantPubKeys(user)
	// fmt.Println(pubKeyToGrantPubKeyMap)
	// fmt.Println(pubKeyToGrantPubKeyMap["65623561653938323231303338303430623038666561646563616232623735366664396339323836356136633234653539396663353165613233623235353039343939666663666431303763"])
	// recoveryPub := recoveryPriv.PubKey()
	// txPub := txPriv.PubKey()
	// postPub := postPriv.PubKey()
	api := api.NewLinoAPIFromArgs("test-chain-BgWrtq", "http://18.188.188.164:26657")
	seq, _ := api.GetSeqNumber("test1")
	err := api.GrantPermission("test1", "lino", 7*24*60*60, model.PostPermission, 10, "A32889124085D932085B23628D966E7F98AE4F711282A82C4B9BBD5E38A143C091E3501313A6BF760D89B3AA67E12754420967BC4C4F963921B26B48BC618E9E79B12D2A94", seq)
	if err != nil {
		panic(err)
	}
	pub, _ := api.GetPostPubKey("lino")
	fmt.Println(pub)
	info, _ := api.GetGrantPubKey("test1", "EB5AE98221037BB974CF968EFD294714D01BDF9D848981147BF7FE7432AED3219AA63E307144")
	fmt.Printf("%+v\n", info)
	// addr := recoveryPub.Address()

	// addrHex := strings.ToUpper(hex.EncodeToString(addr))
	// recoveryPrivHex := hex.EncodeToString(recoveryPriv.Bytes())
	// recoveryPubHex := hex.EncodeToString(recoveryPub.Bytes())
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
	// err = broadcast.Register(user, recoveryPubHex, postPubHex, txPubHex, recoveryPrivHex)
	// if err != nil {
	// 	panic(err)
	// }

}
