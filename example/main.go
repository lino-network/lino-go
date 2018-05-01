package main

import (
	"encoding/hex"
	"strings"

	"github.com/lino-network/lino-go/broadcast"
	crypto "github.com/tendermint/go-crypto"
)

func main() {
	// // query example
	// res, _ := query.GetAllValidators()
	// output, _ := json.Marshal(res)
	// fmt.Println(string(output))
	//
	// res1, _ := query.GetValidator("Lino")
	// output, _ = json.Marshal(res1)
	// fmt.Println(string(output))
	//
	// res2, _ := query.GetDeveloper("Lino")
	// output, _ = json.Marshal(res2)
	// fmt.Println(string(output))
	//
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
	// res6, _ := query.GetAccountBank("B15F5C3647F2962DCEB45B700A13DFF2B5E97CA4")
	// output, _ = json.Marshal(res6)
	// fmt.Println(string(output))
	//
	// res7, _ := query.GetAccountMeta("Lino")
	// output, _ = json.Marshal(res7)
	// fmt.Println(string(output))
	//
	// res8 := query.GetAccountSequence("Lino")
	// output, _ = json.Marshal(res8)
	// fmt.Println(string(output))
	//
	// res9, _ := query.GetVoter("Lino")
	// output, _ = json.Marshal(res9)
	// fmt.Println(string(output))
	//
	// res10, _ := query.GetAccountInfo("Lino")
	// output, _ = json.Marshal(res10)
	// fmt.Println(string(output))

	//broadcast ransaction example
	user1 := "yukai-tu"
	priv1 := crypto.GenPrivKeyEd25519()
	pub1 := priv1.PubKey()
	addr1 := pub1.Address()

	addrHex1 := strings.ToUpper(hex.EncodeToString(addr1))
	privHex1 := hex.EncodeToString(priv1.Bytes())
	linoPrivHex := "a328891240d81fadfd185ff29d0230dd312ff0ded236c15293e635ba1fe3047726546eece62e3126ab8083dc1d845c319ce3002757c036a489818830ceb85b884693940369"

	links := map[string]string{}
	err := broadcast.CreatePost("test8", "a test", "dummy", "Lino", "", "", "", "", "0", linoPrivHex, links)
	if err != nil {
		panic(err)
	}
	err = broadcast.Transfer("Lino", "", addrHex1, "10000", "", linoPrivHex)
	if err != nil {
		panic(err)
	}
	err = broadcast.Register(user1, privHex1)
	if err != nil {
		panic(err)
	}

}
