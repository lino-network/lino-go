package main

import (
	"encoding/hex"
	"strings"

	"github.com/lino-network/lino-go/broadcast"
	"github.com/lino-network/lino-go/model"
	crypto "github.com/tendermint/go-crypto"
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
	// res6, _ := query.GetAccountBank("8F38568B955BBC3E567A570E858292FA77137845")
	// output, _ = json.MarshalIndent(res6, "", "  ")
	// fmt.Println(string(output))
	//
	// res7, _ := query.GetAccountMeta("Lino")
	// output, _ = json.MarshalIndent(res7, "", "  ")
	// fmt.Println(string(output))
	//
	// res8 := query.GetAccountSequence("Lino")
	// output, _ = json.MarshalIndent(res8, "", "  ")
	// fmt.Println(string(output))
	//
	// res9, _ := query.GetVoter("Lino")
	// output, _ = json.MarshalIndent(res9, "", "  ")
	// fmt.Println(string(output))
	//
	// res10, _ := query.GetAccountInfo("Lino")
	// output, _ = json.MarshalIndent(res10, "", "  ")
	// fmt.Println(string(output))

	//broadcast ransaction example
	user1 := "yukai"
	priv1 := crypto.GenPrivKeyEd25519()
	pub1 := priv1.PubKey()
	addr1 := pub1.Address()

	addrHex1 := strings.ToUpper(hex.EncodeToString(addr1))
	privHex1 := hex.EncodeToString(priv1.Bytes())

	registerTx := model.RegisterMsg{
		NewUser:   user1,
		NewPubKey: pub1,
	}
	LinoPrivKey := "a328891240d81fadfd185ff29d0230dd312ff0ded236c15293e635ba1fe3047726546eece62e3126ab8083dc1d845c319ce3002757c036a489818830ceb85b884693940369"
	transferTx := model.TransferMsg{
		Sender:       "Lino",
		ReceiverAddr: addrHex1,
		Amount:       "10000",
	}
	postParam := model.PostCreateParams{
		PostID:       "TestPostID",
		Title:        "this is a test",
		Content:      "dummy content",
		Author:       "Lino",
		ParentAuthor: "",
		ParentPostID: "",
		SourceAuthor: "",
		SourcePostID: "",
		Links:        []model.IDToURLMapping{{Identifier: "id1", URL: "url1"}},
		RedistributionSplitRate: "0",
	}
	postTx := model.CreatePostMsg{
		PostCreateParams: postParam,
	}
	broadcast.BroadcastTransaction(postTx, LinoPrivKey)
	broadcast.BroadcastTransaction(transferTx, LinoPrivKey)
	broadcast.BroadcastTransaction(registerTx, privHex1)

}
