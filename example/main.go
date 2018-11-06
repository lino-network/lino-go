package main

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/lino-network/lino-go/api"
	"github.com/tendermint/tendermint/crypto/secp256k1"
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
	// resetPriv := crypto.GenPrivKeyEd25519()
	// txPriv := crypto.GenPrivKeyEd25519()
	// fmt.Println("tx priv:", strings.ToUpper(hex.EncodeToString(txPriv.Bytes())))
	// appPriv := crypto.GenPrivKeyEd25519()

	// err := broadcast.Register(
	// 	"lino", "100000000", user, hex.EncodeToString(resetPriv.PubKey().Bytes()), hex.EncodeToString(txPriv.PubKey().Bytes()),
	// 	 hex.EncodeToString(appPriv.PubKey().Bytes()), "E1B0F79A207610DF4B9AA480C78F06C5B505B6F56A9B57A8CA05DCA868A41A95B664E319C9", 22)
	// fmt.Println(err)

	// // time.Sleep(3 * time.Second)
	// // err = broadcast.CreatePost(user, "lino", "title", "content", "", "", "", "", "0", nil, strings.ToUpper(hex.EncodeToString(txPriv.Bytes())), 0)
	// // fmt.Println(err)

	// time.Sleep(3 * time.Second)
	// err = broadcast.GrantPermission(user, "lino", 1000000, model.AppPermission, 10, hex.EncodeToString(txPriv.Bytes()), 0)
	// fmt.Println(err)

	// posts, _ := api.Query.GetUserAllPosts(user)
	// fmt.Println(posts)

	// appPub, _ := api.Query.GetAppPubKey("lino")
	// fmt.Println(appPub)
	// pubKeyToGrantPubKeyMap, _ := api.Query.GetAllGrantPubKeys(user)
	// fmt.Println(pubKeyToGrantPubKeyMap)
	// fmt.Println(pubKeyToGrantPubKeyMap["65623561653938323231303338303430623038666561646563616232623735366664396339323836356136633234653539396663353165613233623235353039343939666663666431303763"])
	// resetPub := resetPriv.PubKey()
	// txPub := txPriv.PubKey()
	// appPub := appPriv.PubKey()
	newUserResetKey := secp256k1.GenPrivKey()
	newUserTxKey := secp256k1.GenPrivKey()
	newUserAppKey := secp256k1.GenPrivKey()
	newUser := "ffsffssds"

	api := api.NewLinoAPIFromArgs("lino-staging", "http://18.222.11.221:26657")
	seq, _ := api.GetSeqNumber(context.Background(), "lino")
	resp, err := api.Register(context.Background(), "lino", "100", newUser, hex.EncodeToString(newUserResetKey.PubKey().Bytes()), hex.EncodeToString(newUserTxKey.PubKey().Bytes()), hex.EncodeToString(newUserAppKey.PubKey().Bytes()), "E1B0F79B202FDC4DB4ED428384A06E9A6562527A0A0E85203508700E1BFA96CAB458D899B1", seq)
	if err != nil {
		panic(err)
	}

	fmt.Println(">>resp: ", resp.CommitHash)

	// _, err = api.GrantPermission(newUser, "lino", 7*24*60*60, model.AppPermission, hex.EncodeToString(newUserTxKey.Bytes()), 0)
	// if err != nil {
	// 	panic(err)
	// }
	// pub, _ := api.GetAppPubKey("lino")
	// fmt.Println(pub)
	// info, _ := api.GetGrantPubKey(newUser, pub)
	// fmt.Printf("%+v\n", info)
	// privKey, _ := transport.GetPrivKeyFromHex("E1B0F79B20490005A517EB5CA5C8BE22FB7865ADD64F01AAF9797440DE18F0260A2421E633")
	// sig, err := api.Query.SignWithSha256("wpbqekqjaa", privKey)
	// fmt.Println(sig)
	// res, err := api.Query.VerifyUserSignatureUsingAppKey("lino", "qdgnouryic", "3045022100c359dd4753ff29ce5a67dbabc14ae5ecacdb4ac8d0a4ca944b766b0922dc2fd602203899e2e5f41f740859b58685d2d48284d41fa8daf480d2172877a74b86933794")
	// fmt.Printf("verify sig result: %+v, %+v\n", res, err)

	// addr := resetPub.Address()

	// addrHex := strings.ToUpper(hex.EncodeToString(addr))
	// resetPrivHex := hex.EncodeToString(resetPriv.Bytes())
	// resetPubHex := hex.EncodeToString(resetPub.Bytes())
	// txPubHex := hex.EncodeToString(txPub.Bytes())
	// appPubHex := hex.EncodeToString(appPub.Bytes())

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
	// err = broadcast.Register(user, resetPubHex, appPubHex, txPubHex, resetPrivHex)
	// if err != nil {
	// 	panic(err)
	// }
}
