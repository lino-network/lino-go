package broadcast

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/lino-network/lino-go/model"
	"github.com/lino-network/lino-go/transport"
	"github.com/lino-network/lino-go/types"
)

type Broadcast struct {
	transport *transport.Transport
}

func NewBroadcast(transport *transport.Transport) *Broadcast {
	return &Broadcast{
		transport: transport,
	}
}

// Account related tx
func (broadcast *Broadcast) Register(username, masterPubKeyHex, postPubKeyHex, transactionPubKeyHex, masterPrivKeyHex string) error {
	masterPubKey, err := transport.GetPubKeyFromHex(masterPubKeyHex)
	if err != nil {
		return nil
	}
	postPubKey, err := transport.GetPubKeyFromHex(postPubKeyHex)
	if err != nil {
		return nil
	}
	txPubKey, err := transport.GetPubKeyFromHex(transactionPubKeyHex)
	if err != nil {
		return err
	}

	msg := model.RegisterMsg{
		NewUser:              username,
		NewMasterPubKey:      masterPubKey,
		NewPostPubKey:        postPubKey,
		NewTransactionPubKey: txPubKey,
	}
	return broadcast.broadcastTransaction(msg, masterPrivKeyHex)
}

func (broadcast *Broadcast) Transfer(sender, receiverName, receiverAddr, amount, memo, privKeyHex string) error {
	msg := model.TransferMsg{
		Sender:       sender,
		ReceiverName: receiverName,
		ReceiverAddr: receiverAddr,
		Amount:       amount,
		Memo:         memo,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex)
}

func (broadcast *Broadcast) Follow(follower, followee, privKeyHex string) error {
	msg := model.FollowMsg{
		Followee: followee,
		Follower: follower,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex)
}

func (broadcast *Broadcast) Unfollow(follower, followee, privKeyHex string) error {
	msg := model.UnfollowMsg{
		Followee: followee,
		Follower: follower,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex)
}

func (broadcast *Broadcast) Claim(username, privKeyHex string) error {
	msg := model.ClaimMsg{
		Username: username,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex)
}

func (broadcast *Broadcast) SavingToChecking(username, amount, privKeyHex string) error {
	msg := model.SavingToCheckingMsg{
		Username: username,
		Amount:   amount,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex)
}

func (broadcast *Broadcast) CheckingToSaving(username, amount, privKeyHex string) error {
	msg := model.CheckingToSavingMsg{
		Username: username,
		Amount:   amount,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex)
}

// Post related tx
func (broadcast *Broadcast) CreatePost(postID, title, content, author, parentAuthor, parentPostID,
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
	return broadcast.broadcastTransaction(msg, privKeyHex)
}

func (broadcast *Broadcast) Like(username, author, postID, privKeyHex string, weight int64) error {
	msg := model.LikeMsg{
		Username: username,
		Weight:   weight,
		Author:   author,
		PostID:   postID,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex)
}

func (broadcast *Broadcast) Donate(username, author, postID, amount, fromApp, privKeyHex string) error {
	msg := model.DonateMsg{
		Username: username,
		Author:   author,
		PostID:   postID,
		FromApp:  fromApp,
		Amount:   amount,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex)
}

func (broadcast *Broadcast) ReportOrUpvote(username, author, postID, privKeyHex string, isReport bool) error {
	msg := model.ReportOrUpvoteMsg{
		Username: username,
		Author:   author,
		PostID:   postID,
		IsReport: isReport,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex)
}

func (broadcast *Broadcast) View(username, author, postID, privKeyHex string) error {
	msg := model.ViewMsg{
		Username: username,
		Author:   author,
		PostID:   postID,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex)
}

func (broadcast *Broadcast) DeletePost(title, author, postID, privKeyHex string) error {
	msg := model.DeletePostMsg{
		Author: author,
		PostID: postID,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex)
}

func (broadcast *Broadcast) UpdatePost(author, postID, title, content, redistributionSplitRate, privKeyHex string, links map[string]string) error {
	var mLinks []model.IDToURLMapping
	if links == nil || len(links) == 0 {
		mLinks = nil
	} else {
		for k, v := range links {
			mLinks = append(mLinks, model.IDToURLMapping{k, v})
		}
	}

	msg := model.UpdatePostMsg{
		Author:  author,
		PostID:  postID,
		Title:   title,
		Content: content,
		Links:   mLinks,
		RedistributionSplitRate: redistributionSplitRate,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex)
}

// Validator related tx
func (broadcast *Broadcast) ValidatorDeposit(username, deposit, privKeyHex string) error {
	privKey, err := transport.GetPrivKeyFromHex(privKeyHex)
	if err != nil {
		return err
	}

	msg := model.ValidatorDepositMsg{
		Username:  username,
		Deposit:   deposit,
		ValPubKey: privKey.PubKey(),
	}
	return broadcast.broadcastTransaction(msg, privKeyHex)
}

func (broadcast *Broadcast) ValidatorWithdraw(username, amount, privKeyHex string) error {
	msg := model.ValidatorWithdrawMsg{
		Username: username,
		Amount:   amount,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex)
}

func (broadcast *Broadcast) ValidatorRevoke(username, privKeyHex string) error {
	msg := model.ValidatorRevokeMsg{
		Username: username,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex)
}

// Vote related tx
func (broadcast *Broadcast) Vote(voter, proposalID, privKeyHex string, result bool) error {
	msg := model.VoteMsg{
		Voter:      voter,
		ProposalID: proposalID,
		Result:     result,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex)
}

func (broadcast *Broadcast) VoterDeposit(username, deposit, privKeyHex string) error {
	msg := model.VoterDepositMsg{
		Username: username,
		Deposit:  deposit,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex)
}

func (broadcast *Broadcast) VoterWithdraw(username, amount, privKeyHex string) error {
	msg := model.VoterWithdrawMsg{
		Username: username,
		Amount:   amount,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex)
}

func (broadcast *Broadcast) VoterRevoke(username, privKeyHex string) error {
	msg := model.VoterRevokeMsg{
		Username: username,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex)
}

func (broadcast *Broadcast) Delegate(delegator, voter, amount, privKeyHex string) error {
	msg := model.DelegateMsg{
		Delegator: delegator,
		Voter:     voter,
		Amount:    amount,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex)
}

func (broadcast *Broadcast) DelegatorWithdraw(delegator, voter, amount, privKeyHex string) error {
	msg := model.DelegatorWithdrawMsg{
		Delegator: delegator,
		Voter:     voter,
		Amount:    amount,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex)
}

func (broadcast *Broadcast) RevokeDelegation(delegator, voter, privKeyHex string) error {
	msg := model.RevokeDelegationMsg{
		Delegator: delegator,
		Voter:     voter,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex)
}

// Developer related tx
func (broadcast *Broadcast) DeveloperRegister(username, deposit, privKeyHex string) error {
	msg := model.DeveloperRegisterMsg{
		Username: username,
		Deposit:  deposit,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex)
}

func (broadcast *Broadcast) DeveloperRevoke(username, privKeyHex string) error {
	msg := model.DeveloperRegisterMsg{
		Username: username,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex)
}

func (broadcast *Broadcast) GrantDeveloper(username, authenticateApp, privKeyHex string, validityPeriod, grantLevel int64) error {
	msg := model.GrantDeveloperMsg{
		Username:        username,
		AuthenticateApp: authenticateApp,
		ValidityPeriod:  validityPeriod,
		GrantLevel:      grantLevel,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex)
}

// infra related tx
func (broadcast *Broadcast) ProviderReport(username, privKeyHex string, usage int64) error {
	msg := model.ProviderReportMsg{
		Username: username,
		Usage:    usage,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex)
}

// proposal related tx

// internal helper functions
func (broadcast *Broadcast) broadcastTransaction(transaction interface{}, privKeyHex string) error {
	res, err := broadcast.transport.SignBuildBroadcast(transaction, privKeyHex, 0)

	var reg = regexp.MustCompile(`expected (\d+)`)
	var tries = 1

	// keep trying to get newest sequence number
	for err == nil && res.CheckTx.Code == types.InvalidSeqErrCode {
		match := reg.FindString(res.CheckTx.Log)
		seq, err := strconv.ParseInt(match[9:], 10, 64)
		if err != nil || tries == types.BroadcastMaxTries {
			return err
		}
		res, err = broadcast.transport.SignBuildBroadcast(transaction, privKeyHex, seq)
		tries += 1
	}

	if err != nil {
		return err
	}
	if res.CheckTx.Code != uint32(0) {
		return fmt.Errorf("CheckTx failed ! (%d) %s", res.CheckTx.Code, res.CheckTx.Log)
	}
	if res.DeliverTx.Code != uint32(0) {
		return fmt.Errorf("DeliverTx failed ! (%d) %s ", res.DeliverTx.Code, res.DeliverTx.Log)
	}
	fmt.Printf("Committed at block %d. Hash: %s\n", res.Height, res.Hash.String())
	return nil
}
