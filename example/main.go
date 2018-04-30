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
	user1 := "yukai"
	priv1 := crypto.GenPrivKeyEd25519()
	pub1 := priv1.PubKey()
	addr1 := pub1.Address()

	aa := [32]byte(pub1.(crypto.PubKeyEd25519))
	bb := [64]byte(priv1)

	pubHex1 := hex.EncodeToString(append([]byte{0x1}, aa[:]...))
	addrHex1 := strings.ToUpper(hex.EncodeToString(addr1))
	privHex1 := hex.EncodeToString(append([]byte{0x1}, bb[:]...))

	registerTx := model.RegisterMsg{
		NewUser:   user1,
		NewPubKey: pubHex1,
	}
	LinoPrivKey := "016b08e8e10fbd35180a3bc9ede5cfe90ba6578cf6d7e7406e40135e4b09d17526fdf22a3f5bfb5e64648cb09ddf7c069b246a2335a473ed03a3742f6fe6cd0096"
	transferTx := model.TransferMsg{
		Sender:       "Lino",
		ReceiverAddr: addrHex1,
		Amount:       "8888888",
	}

	broadcast.BroadcastTransaction(transferTx, LinoPrivKey)
	broadcast.BroadcastTransaction(registerTx, privHex1)

}
