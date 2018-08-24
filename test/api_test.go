package test

import (
	"testing"
	"time"

	"github.com/lino-network/lino-go/api"
)

var (
	chainID = "lino-testnet"
	// nodeURL = "http://localhost:26657"
	nodeURL = "http://fullnode.linovalidator.io:80"

	referrer          = "lino"
	registerFee       = "10000000"
	referrerTxPrivKey = "E1B0F79B200D44E9B233AB277047A86D4DC3F247E213AEC15185EFE15DF6E1C19B90EB1AEE"

	myUser    = "myuser1"
	txPrivHex = "E1B0F79B20801B33B99F73D134AF828874A2B9716A5F15B562115930414BC398EB96807F7A"

	post1 = "post1"

	testAPI *api.API
)

func setup(t *testing.T) {
	options := api.TimeoutOptions{
		QueryTimeout:     1 * time.Second,
		BroadcastTimeout: 4 * time.Second,
	}

	testAPI = api.NewLinoAPIFromArgs(chainID, nodeURL, options)
}

// func TestBasic(t *testing.T) {
// 	setup(t)

// 	resetPriv := secp256k1.GenPrivKey()
// 	txPriv := secp256k1.GenPrivKey()
// 	appPriv := secp256k1.GenPrivKey()

// 	t.Errorf("reset private key is: %s", strings.ToUpper(hex.EncodeToString(resetPriv.Bytes())))
// 	t.Errorf("transaction private key is: %s", strings.ToUpper(hex.EncodeToString(txPriv.Bytes())))
// 	t.Errorf("app private key is: %s", strings.ToUpper(hex.EncodeToString(appPriv.Bytes())))

// 	resetPub := resetPriv.PubKey()
// 	txPub := txPriv.PubKey()
// 	appPub := appPriv.PubKey()

// 	resetPubHex := hex.EncodeToString(resetPub.Bytes())
// 	txPubHex := hex.EncodeToString(txPub.Bytes())
// 	appPubHex := hex.EncodeToString(appPub.Bytes())

// 	seq, err := testAPI.GetSeqNumber(referrer)
// 	if err != nil {
// 		t.Errorf("failed to get seq: %v", err)
// 	}

// 	err = testAPI.Register(referrer, registerFee, myUser, resetPubHex, txPubHex, appPubHex, referrerTxPrivKey, seq)
// 	if err != nil {
// 		t.Errorf("failed to register: %v", err)
// 	}
// }

func TestAccount(t *testing.T) {
	setup(t)

	// myUserSeq, err := testAPI.GetSeqNumber(myUser)
	// if err != nil {
	// 	t.Errorf("TestAccount: failed to get myUser seq number: %v", err)
	// }
	// err = testAPI.Follow(myUser, "lino", txPrivHex, myUserSeq)
	// if err != nil {
	// 	t.Errorf("TestAccount: failed to broadcast claim msg: %v", err)
	// }

	// _, err := testAPI.GetAllFollowerMeta("lino")
	// if err != nil {
	// 	t.Errorf("TestAccount: failed to get account all relationship: %v", err)
	// }
	// t.Error("err")

	// _, err := testAPI.GetAllFollowingMeta(myUser)
	// if err != nil {
	// 	t.Errorf("TestAccount: failed to get account all relationship: %v", err)
	// }

	// _, err := testAPI.GetAccountInfo("linotv")
	// if err != nil {
	// 	t.Errorf("TestAccount: failed to get account info: %v", err)
	// }
	// t.Errorf("lino ai: %v ", ai)

	// _, err := testAPI.GetAccountBank("lino")
	// if err != nil {
	// 	t.Errorf("TestAccount: failed to get account bank: %v", err)
	// }
	// t.Errorf("ab: %v", ab)

	// 	_, err = testAPI.GetAccountMeta("lino")
	// 	if err != nil {
	// 		t.Errorf("TestAccount: failed to get account meta: %v", err)
	// 	}

	// r1, err := testAPI.GetReward("test1")
	// if err != nil {
	// 	t.Errorf("TestAccount: failed to get reward: %v", err)
	// }
	// fmt.Println(">>r1: ", r1)

	// r2, err := testAPI.GetRewardAtHeight("test1", 10)
	// if err != nil {
	// 	t.Errorf("TestAccount: failed to get reward at height: %v", err)
	// }
	// fmt.Println(">>r2: ", r2)

	// _, err := testAPI.GetAllBalanceHistory(myUser)
	// if err != nil {
	// 	t.Errorf("TestAccount: failed to get balance history: %v", err)
	// }

	// Tested successfully
	//
	// linoSeq, err := testAPI.GetSeqNumber("linowallet")
	// if err != nil {
	// 	t.Errorf("TestAccount: failed to get myUser seq number: %v", err)
	// }
	// fmt.Println("seq: ", linoSeq)
	// err = testAPI.Transfer("linowallet", "jawson", "1", "test", "E1B0F79B20CFD61D55BD37DC67198E722F4B1F2721D5A0AB3B11F9F6F7B975293F083608FE", linoSeq)
	// if err != nil {
	// 	t.Errorf("TestAccount: failed to transfer 1, got err: %v", err)
	// }

	// seq, err := testAPI.GetSeqNumber("lino")
	// if err != nil {
	// 	t.Errorf("TestAccount: failed to get lino seq number: %v", err)
	// }
	// err = testAPI.Claim("lino", referrerTxPrivKey, seq)
	// if err != nil {
	// 	t.Errorf("TestAccount: failed to broadcast claim msg: %v", err)
	// }
}

// func TestRecentBalanceHistory(t *testing.T) {
// 	setup(t)
// 	// Note: total num of history is 116

// 	// corner case - invalid numHistory
// 	_, err := testAPI.GetRecentBalanceHistory(myUser, -1)
// 	if err == nil {
// 		t.Errorf("GetRecentBalanceHistory should return InvalidArg err: %v", err)
// 	}

// 	// corner case - numHistory is larger than total length
// 	history, err := testAPI.GetRecentBalanceHistory(myUser, 120)
// 	if err != nil {
// 		t.Errorf("GetRecentBalanceHistory fails, err: %v", err)
// 	} else if len(history.Details) != 116 {
// 		t.Errorf("GetRecentBalanceHistory got diff resp, got %v, want 116", len(history.Details))
// 	}

// 	// normal case
// 	history, err = testAPI.GetRecentBalanceHistory(myUser, 2)
// 	if err != nil {
// 		t.Errorf("GetRecentBalanceHistory fails, err: %v", err)
// 	} else if len(history.Details) != 2 {
// 		t.Errorf("GetRecentBalanceHistory got diff resp, got %v, want 2", len(history.Details))
// 	} else if history.Details[0].Memo != "memo49" || history.Details[1].Memo != "memo48" {
// 		t.Errorf("GetRecentBalanceHistory got non-ordered resp, got %v", history.Details)
// 	}

// 	// normal case - get from two bucket slots
// 	history, err = testAPI.GetRecentBalanceHistory(myUser, 51)
// 	if err != nil {
// 		t.Errorf("GetRecentBalanceHistory fails, err: %v", err)
// 	} else if len(history.Details) != 51 {
// 		t.Errorf("GetRecentBalanceHistory got diff resp, got %v, want 116", len(history.Details))
// 	} else if history.Details[50].Memo != "memo59" || history.Details[50].From != "lino" {
// 		t.Errorf("GetRecentBalanceHistory got non-ordered resp, got %v", history.Details[50])
// 	}
// }

// func TestBalanceHistoryFromTo(t *testing.T) {
// 	// setup(t)
// 	// Note: total num of history is 116

// 	// corner case - invalid arg
// 	_, err := testAPI.GetBalanceHistoryFromTo(myUser, -1, -2)
// 	if err == nil {
// 		t.Errorf("GetBalanceHistoryFromTo should return InvalidArg err: %v", err)
// 	}
// 	_, err = testAPI.GetBalanceHistoryFromTo(myUser, 20, 10)
// 	if err == nil {
// 		t.Errorf("GetBalanceHistoryFromTo should return InvalidArg err: %v", err)
// 	}
// 	_, err = testAPI.GetBalanceHistoryFromTo(myUser, 120, 10)
// 	if err == nil {
// 		t.Errorf("GetBalanceHistoryFromTo should return InvalidArg err: %v", err)
// 	}

// 	// normal case
// 	history, err := testAPI.GetBalanceHistoryFromTo(myUser, 114, 115)
// 	if err != nil {
// 		t.Errorf("GetBalanceHistoryFromTo fails, err: %v", err)
// 	} else if len(history.Details) != 2 {
// 		t.Errorf("GetBalanceHistoryFromTo got diff resp, got %v, want 2", len(history.Details))
// 	} else if history.Details[0].Memo != "memo49" || history.Details[1].Memo != "memo48" {
// 		t.Errorf("GetBalanceHistoryFromTo got non-ordered resp, got %v", history.Details)
// 	}

// 	// normal case - to is larger than total length
// 	history, err = testAPI.GetBalanceHistoryFromTo(myUser, 114, 120)
// 	if err != nil {
// 		t.Errorf("GetBalanceHistoryFromTo fails, err: %v", err)
// 	} else if len(history.Details) != 2 {
// 		t.Errorf("GetBalanceHistoryFromTo got diff resp, got %v, want 2", len(history.Details))
// 	} else if history.Details[0].Memo != "memo49" || history.Details[1].Memo != "memo48" {
// 		t.Errorf("GetBalanceHistoryFromTo got non-ordered resp, got %v", history.Details)
// 	}

// 	// normal case - get from two bucket slots
// 	history, err = testAPI.GetBalanceHistoryFromTo(myUser, 64, 114)
// 	if err != nil {
// 		t.Errorf("GetBalanceHistoryFromTo fails, err: %v", err)
// 	} else if len(history.Details) != 51 {
// 		t.Errorf("GetBalanceHistoryFromTo got diff resp, got %v, want 51", len(history.Details))
// 	} else if history.Details[0].Memo != "memo48" || history.Details[0].From != myUser {
// 		t.Errorf("GetRecentBalanceHistory got non-ordered resp, got %v", history.Details[0])
// 	} else if history.Details[50].Memo != "memo58" || history.Details[50].From != "lino" {
// 		t.Errorf("GetRecentBalanceHistory got non-ordered resp, got %v", history.Details[50])
// 	}
// }

// func TestTransfer(t *testing.T) {
// testAPI := setup(t)

// Test successfully
//

// seq, err := testAPI.GetSeqNumber("lino")
// if err != nil {
// 	t.Errorf("TestAccount: failed to get lino seq number: %v", err)
// }
// for i := 0; i < 60; i++ {
// 	memo := "memo" + strconv.Itoa(i)
// 	amount := strconv.Itoa(i + 1)
// 	err = testAPI.Transfer("lino", myUser, amount, memo, referrerTxPrivKey, seq)
// 	if err != nil {
// 		t.Errorf("TestAccount: failed to transfer 1 to myUser: %v", err)
// 	}
// 	seq++
// }

// seq, err := testAPI.GetSeqNumber(myUser)
// if err != nil {
// 	t.Errorf("TestAccount: failed to get myUser seq number: %v", err)
// }
// for i := 0; i < 50; i++ {
// 	memo := "memo" + strconv.Itoa(i)
// 	amount := strconv.Itoa(i + 1)
// 	err = testAPI.Transfer(myUser, "lino", amount, memo, txPrivHex, seq)
// 	if err != nil {
// 		t.Errorf("TestAccount: failed to transfer 1 to lino")
// 	}
// 	seq++
// }
// }

// func TestPost(t *testing.T) {
// 	setup(t)

// Tested successfully
//
// seq, err := testAPI.GetSeqNumber(myUser)
// if err != nil {
// 	t.Errorf("TestAccount: failed to get seq number: %v", err)
// }
// links := map[string]string{}
// err = testAPI.CreatePost(myUser, post1, "mytitle", "mycontent", "", "", "", "", "0.2", links, txPrivHex, seq)
// if err != nil {
// 	t.Errorf("TestPost: failed to broadcast CreatePost msg: %v", err)
// }

// _, err := testAPI.GetPostInfo(myUser, post1)
// if err != nil {
// 	t.Errorf("TestPost: failed to get post info: %v", err)
// }

// 	_, err = testAPI.GetPostMeta(myUser, post1)
// 	if err != nil {
// 		t.Errorf("TestPost: failed to get post meta: %v", err)
// 	}

// seq, err := testAPI.GetSeqNumber(myUser)
// if err != nil {
// 	t.Errorf("TestAccount: failed to get myUser seq number: %v", err)
// }

// 	links := map[string]string{}
// 	err = testAPI.UpdatePost(myUser, "newTitle", post1, "newContent", "0.3", links, txPrivHex, seq)
// 	if err != nil {
// 		t.Errorf("TestPost: failed to broadcast UpdatePost: %v", err)
// 	}

// Tested successfully
//
// seq++
// err = testAPI.DeletePost(myUser, post1, txPrivHex, seq)
// if err != nil {
// 	t.Errorf("TestPost: failed to broadcast DeletePost: %v", err)
// }
// }

// func TestVoting(t *testing.T) {

// _, err := testAPI.GetVoter("lino")
// if err != nil {
// 	t.Errorf("TestVoting: failed to get voter lino: %v", err)
// }

// _, err = testAPI.GetAllDelegation(myUser)
// if err != nil {
// 	t.Errorf("TestVoting: failed to get all delegation: %v", err)
// }

// seq, err := testAPI.GetSeqNumber(myUser)
// if err != nil {
// 	t.Errorf("TestVoting: failed to get myUser seq number: %v", err)
// }

// err = testAPI.VoterWithdraw(myUser, "500", txPrivHex, seq)
// if err != nil {
// 	t.Errorf("TestVoting: failed to broadcast VoterWithdraw: %v", err)
// }

// err = testAPI.Delegate(myUser, "lino", "1000", txPrivHex, seq)
// if err != nil {
// 	t.Errorf("TestVoting: failed to broadcast Delegate: %v", err)
// }

// // Tested successfully
// err = testAPI.VoterDeposit(myUser, "320000", txPrivHex, seq)
// if err != nil {
// 	t.Errorf("TestVoting: failed to broadcast VoterDeposit: %v", err)
// }
// }

// func TestValidators(t *testing.T) {

// 	_, err := testAPI.GetAllValidators()
// 	if err != nil {
// 		t.Errorf("TestValidators: failed to get all validators: %v", err)
// 	}

// 	_, err = testAPI.GetValidator("lino")
// 	if err != nil {
// 		t.Errorf("TestValidators: failed to get validator lino: %v", err)
// 	}
// }

// func TestDevelopers(t *testing.T) {

// 	_, err := testAPI.GetDevelopers()
// 	if err != nil {
// 		t.Errorf("TestDevelopers: failed to get all developers: %v", err)
// 	}

// 	_, err = testAPI.GetDeveloper("lino")
// 	if err != nil {
// 		t.Errorf("TestDevelopers: failed to get developer lino: %v", err)
// 	}
// }

// func TestInfra(t *testing.T) {

// 	_, err := testAPI.GetInfraProviders()
// 	if err != nil {
// 		t.Errorf("TestInfra: failed to get all infra providers: %v", err)
// 	}

// 	_, err = testAPI.GetInfraProvider("lino")
// 	if err != nil {
// 		t.Errorf("TestInfra: failed to get infra provider lino: %v", err)
// 	}
// }

// func TestBlocks(t *testing.T) {

// 	_, err := testAPI.GetBlock(1)
// 	if err != nil {
// 		t.Errorf("TestBlocks: failed to get block at height 1: %v", err)
// 	}

// 	bs, err := testAPI.GetBlockStatus()
// 	if err != nil {
// 		t.Errorf("TestBlocks: failed to get block status: %v", err)
// 	}
// 	t.Errorf(">> bs: %v", bs)
// }

// func TestProposal(t *testing.T) {

// 	_, err := testAPI.GetProposalList()
// 	if err != nil {
// 		t.Errorf("TestProposal: failed to get all proposals: %v", err)
// 	}

// _, err = testAPI.GetProposal("1")
// if err != nil {
// 	t.Errorf("TestProposal: failed to get proposal: %v", err)
// }

// 	_, err = testAPI.GetOngoingProposal()
// 	if err != nil {
// 		t.Errorf("TestProposal: failed to get onging proposal: %v", err)
// 	}

// 	_, err = testAPI.GetExpiredProposal()
// 	if err != nil {
// 		t.Errorf("TestProposal: failed to get expired proposal: %v", err)
// 	}
// }

// func TestParams(t *testing.T) {
// testAPI := setup(t)

// seq, err := testAPI.GetSeqNumber(myUser)
// if err != nil {
// 	t.Errorf("TestAccount: failed to get myUser seq number: %v", err)
// }

// p := model.CoinDayParam{
// 	DaysToRecoverCoinDayStake:    int64(10),
// 	SecondsToRecoverCoinDayStake: int64(7 * 24 * 3600),
// }
// err = testAPI.ChangeCoinDayParam(myUser, p, txPrivHex, seq)
// if err != nil {
// 	t.Errorf("TestParams: failed to broadcast ChangeCoinDayParam: %v", err)
// }

// _, err := testAPI.GetEvaluateOfContentValueParam()
// if err != nil {
// 	t.Errorf("TestParams: failed to get evaluate of content value param: %v", err)
// }
// t.Errorf("TestPar")

// _, err = testAPI.GetGlobalAllocationParam()
// if err != nil {
// 	t.Errorf("TestParams: failed to get global allocation param: %v", err)
// }

// _, err = testAPI.GetInfraInternalAllocationParam()
// if err != nil {
// 	t.Errorf("TestParams: failed to get infra internal allocation param: %v", err)
// }

// _, err = testAPI.GetDeveloperParam()
// if err != nil {
// 	t.Errorf("TestParams: failed to get developer param: %v", err)
// }

// _, err = testAPI.GetVoteParam()
// if err != nil {
// 	t.Errorf("TestParams: failed to get vote param: %v", err)
// }

// _, err = testAPI.GetProposalParam()
// if err != nil {
// 	t.Errorf("TestParams: failed to get proposal param: %v", err)
// }

// _, err = testAPI.GetValidatorParam()
// if err != nil {
// 	t.Errorf("TestParams: failed to get validator param: %v", err)
// }

// _, err = testAPI.GetCoinDayParam()
// if err != nil {
// 	t.Errorf("TestParams: failed to get coin day param: %v", err)
// }

// _, err = testAPI.GetBandwidthParam()
// if err != nil {
// 	t.Errorf("TestParams: failed to get bandwidth param: %v", err)
// }

// _, err = testAPI.GetAccountParam()
// if err != nil {
// 	t.Errorf("TestParams: failed to get account param: %v", err)
// }
// }
