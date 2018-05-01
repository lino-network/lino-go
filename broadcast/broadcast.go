package broadcast

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/lino-network/lino-go/model"
	"github.com/lino-network/lino-go/transport"
	"github.com/lino-network/lino-go/types"
)

// Account related tx
func Register(username, privKeyHex string) error {
	privKey, err := transport.GetPrivKeyFromHex(privKeyHex)
	if err != nil {
		return err
	}

	msg := model.RegisterMsg{
		NewUser:   username,
		NewPubKey: privKey.PubKey(),
	}
	return broadcastTransaction(msg, privKeyHex)
}

func Transfer(sender, receiverName, receiverAddr, amount, memo, privKeyHex string) error {
	msg := model.TransferMsg{
		Sender:       sender,
		ReceiverName: receiverName,
		ReceiverAddr: receiverAddr,
		Amount:       amount,
		Memo:         memo,
	}
	return broadcastTransaction(msg, privKeyHex)
}

func Follow(follower, followee, privKeyHex string) error {
	msg := model.FollowMsg{
		Followee: followee,
		Follower: follower,
	}
	return broadcastTransaction(msg, privKeyHex)
}

func Unfollow(follower, followee, privKeyHex string) error {
	msg := model.UnfollowMsg{
		Followee: followee,
		Follower: follower,
	}
	return broadcastTransaction(msg, privKeyHex)
}

func Claim(username, privKeyHex string) error {
	msg := model.ClaimMsg{
		Username: username,
	}
	return broadcastTransaction(msg, privKeyHex)
}

// Post related tx
func CreatePost(postID, title, content, author, parentAuthor, parentPostID,
	sourceAuthor, sourcePostID, redistributionSplitRate, privKeyHex string, links map[string]string) error {
	var mLinks []model.IDToURLMapping
	if links == nil || len(links) == 0 {
		mLinks = nil
	} else {
		for k, v := range links {
			mLinks = append(mLinks, model.IDToURLMapping{k, v})
		}
	}

	para := model.PostCreateParams{
		PostID:                  postID,
		Title:                   title,
		Content:                 content,
		Author:                  author,
		ParentAuthor:            parentAuthor,
		ParentPostID:            parentPostID,
		SourceAuthor:            sourceAuthor,
		SourcePostID:            sourcePostID,
		RedistributionSplitRate: redistributionSplitRate,
		Links: mLinks,
	}
	msg := model.CreatePostMsg{
		para,
	}
	return broadcastTransaction(msg, privKeyHex)
}

func Like(username, author, postID, privKeyHex string, weight int64) error {
	msg := model.LikeMsg{
		Username: username,
		Weight:   weight,
		Author:   author,
		PostID:   postID,
	}
	return broadcastTransaction(msg, privKeyHex)
}

func Donate(username, author, postID, amount, fromApp, privKeyHex string) error {
	msg := model.DonateMsg{
		Username: username,
		Author:   author,
		PostID:   postID,
		FromApp:  fromApp,
		Amount:   amount,
	}
	return broadcastTransaction(msg, privKeyHex)
}

func ReportOrUpvote(username, author, postID, privKeyHex string, isReport, isRevoke bool) error {
	msg := model.ReportOrUpvoteMsg{
		Username: username,
		Author:   author,
		PostID:   postID,
		IsReport: isReport,
		IsRevoke: isRevoke,
	}
	return broadcastTransaction(msg, privKeyHex)
}

// Validator related tx
func ValidatorDeposit(username, deposit, privKeyHex string) error {
	privKey, err := transport.GetPrivKeyFromHex(privKeyHex)
	if err != nil {
		return err
	}

	msg := model.ValidatorDepositMsg{
		Username:  username,
		Deposit:   deposit,
		ValPubKey: privKey.PubKey(),
	}
	return broadcastTransaction(msg, privKeyHex)
}

func ValidatorWithdraw(username, amount, privKeyHex string) error {
	msg := model.ValidatorWithdrawMsg{
		Username: username,
		Amount:   amount,
	}
	return broadcastTransaction(msg, privKeyHex)
}

func ValidatorRevoke(username, privKeyHex string) error {
	msg := model.ValidatorRevokeMsg{
		Username: username,
	}
	return broadcastTransaction(msg, privKeyHex)
}

// Vote related tx
func Vote(voter, proposalID, privKeyHex string, result bool) error {
	msg := model.VoteMsg{
		Voter:      voter,
		ProposalID: proposalID,
		Result:     result,
	}
	return broadcastTransaction(msg, privKeyHex)
}

func VoterDeposit(username, deposit, privKeyHex string) error {
	msg := model.VoterDepositMsg{
		Username: username,
		Deposit:  deposit,
	}
	return broadcastTransaction(msg, privKeyHex)
}

func VoterWithdraw(username, amount, privKeyHex string) error {
	msg := model.VoterWithdrawMsg{
		Username: username,
		Amount:   amount,
	}
	return broadcastTransaction(msg, privKeyHex)
}

func VoterRevoke(username, privKeyHex string) error {
	msg := model.VoterRevokeMsg{
		Username: username,
	}
	return broadcastTransaction(msg, privKeyHex)
}

func Delegate(delegator, voter, amount, privKeyHex string) error {
	msg := model.DelegateMsg{
		Delegator: delegator,
		Voter:     voter,
		Amount:    amount,
	}
	return broadcastTransaction(msg, privKeyHex)
}

func DelegatorWithdraw(delegator, voter, amount, privKeyHex string) error {
	msg := model.DelegatorWithdrawMsg{
		Delegator: delegator,
		Voter:     voter,
		Amount:    amount,
	}
	return broadcastTransaction(msg, privKeyHex)
}

func RevokeDelegation(delegator, voter, privKeyHex string) error {
	msg := model.RevokeDelegationMsg{
		Delegator: delegator,
		Voter:     voter,
	}
	return broadcastTransaction(msg, privKeyHex)
}

// Developer related tx
func DeveloperRegister(username, deposit, privKeyHex string) error {
	msg := model.DeveloperRegisterMsg{
		Username: username,
		Deposit:  deposit,
	}
	return broadcastTransaction(msg, privKeyHex)
}

func DeveloperRevoke(username, privKeyHex string) error {
	msg := model.DeveloperRegisterMsg{
		Username: username,
	}
	return broadcastTransaction(msg, privKeyHex)
}

func GrantDeveloper(username, authenticateApp, privKeyHex string, validityPeriod, grantLevel int64) error {
	msg := model.GrantDeveloperMsg{
		Username:        username,
		AuthenticateApp: authenticateApp,
		ValidityPeriod:  validityPeriod,
		GrantLevel:      grantLevel,
	}
	return broadcastTransaction(msg, privKeyHex)
}

// infra related tx
func ProviderReport(username, privKeyHex string, usage int64) error {
	msg := model.ProviderReportMsg{
		Username: username,
		Usage:    usage,
	}
	return broadcastTransaction(msg, privKeyHex)
}

// internal helper functions
func broadcastTransaction(transaction interface{}, privKeyHex string) error {
	transport := transport.NewTransportFromViper()
	res, err := transport.SignBuildBroadcast(transaction, privKeyHex, 0)

	var reg = regexp.MustCompile(`expected (\d+)`)
	var tries = 1

	// keep trying to get newest sequence number
	for err == nil && res.CheckTx.Code == types.InvalidSeqErrCode {
		match := reg.FindString(res.CheckTx.Log)
		seq, err := strconv.ParseInt(match[9:], 10, 64)
		if err != nil || tries == types.BroadcastMaxTries {
			return err
		}
		res, err = transport.SignBuildBroadcast(transaction, privKeyHex, seq)
		tries += 1
	}

	if err != nil {
		fmt.Println("Build and Sign message failed ! ", err)
		return err
	}
	if res.CheckTx.Code != uint32(0) {
		fmt.Println("CheckTx failed ! code: ", res.CheckTx.Code, res.CheckTx.Log)
		return err
	}
	if res.DeliverTx.Code != uint32(0) {
		fmt.Println("DeliverTx failed ! code: ", res.DeliverTx.Code, res.DeliverTx.Log)
		return err
	}
	fmt.Printf("Committed at block %d. Hash: %s\n", res.Height, res.Hash.String())
	return nil
}
