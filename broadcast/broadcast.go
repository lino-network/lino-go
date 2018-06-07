package broadcast

import (
	"fmt"

	"github.com/lino-network/lino-go/model"
	"github.com/lino-network/lino-go/transport"
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
func (broadcast *Broadcast) Register(seq int64, username, masterPubKeyHex, postPubKeyHex, transactionPubKeyHex, masterPrivKeyHex, referrer, registerFee string) error {
	masterPubKey, err := transport.GetPubKeyFromHex(masterPubKeyHex)
	if err != nil {
		return err
	}
	postPubKey, err := transport.GetPubKeyFromHex(postPubKeyHex)
	if err != nil {
		return err
	}
	txPubKey, err := transport.GetPubKeyFromHex(transactionPubKeyHex)
	if err != nil {
		return err
	}

	msg := model.RegisterMsg{
		Referrer:             referrer,
		RegisterFee:          registerFee,
		NewUser:              username,
		NewMasterPubKey:      masterPubKey,
		NewPostPubKey:        postPubKey,
		NewTransactionPubKey: txPubKey,
	}
	return broadcast.broadcastTransaction(seq, msg, masterPrivKeyHex)
}

func (broadcast *Broadcast) Transfer(seq int64, sender, receiver, amount, memo, privKeyHex string) error {
	msg := model.TransferMsg{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
		Memo:     memo,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) Follow(seq int64, follower, followee, privKeyHex string) error {
	msg := model.FollowMsg{
		Followee: followee,
		Follower: follower,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) Unfollow(seq int64, follower, followee, privKeyHex string) error {
	msg := model.UnfollowMsg{
		Followee: followee,
		Follower: follower,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) Claim(seq int64, username, privKeyHex string) error {
	msg := model.ClaimMsg{
		Username: username,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) Recover(seq int64, username, newMasterPubKeyHex, newPostPubKeyHex, newTransactionPubKeyHex, privKeyHex string) error {
	masterPubKey, err := transport.GetPubKeyFromHex(newMasterPubKeyHex)
	if err != nil {
		return nil
	}
	postPubKey, err := transport.GetPubKeyFromHex(newPostPubKeyHex)
	if err != nil {
		return nil
	}
	txPubKey, err := transport.GetPubKeyFromHex(newTransactionPubKeyHex)
	if err != nil {
		return err
	}

	msg := model.RecoverMsg{
		Username:             username,
		NewMasterPubKey:      masterPubKey,
		NewPostPubKey:        postPubKey,
		NewTransactionPubKey: txPubKey,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

// Post related tx
func (broadcast *Broadcast) CreatePost(seq int64, postID, title, content, author, parentAuthor, parentPostID,
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
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) Like(seq int64, username, author, postID, privKeyHex string, weight int64) error {
	msg := model.LikeMsg{
		Username: username,
		Weight:   weight,
		Author:   author,
		PostID:   postID,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) Donate(seq int64, username, author, postID, amount, fromApp, privKeyHex string) error {
	msg := model.DonateMsg{
		Username: username,
		Author:   author,
		PostID:   postID,
		FromApp:  fromApp,
		Amount:   amount,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) ReportOrUpvote(seq int64, username, author, postID, privKeyHex string, isReport bool) error {
	msg := model.ReportOrUpvoteMsg{
		Username: username,
		Author:   author,
		PostID:   postID,
		IsReport: isReport,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) View(seq int64, username, author, postID, privKeyHex string) error {
	msg := model.ViewMsg{
		Username: username,
		Author:   author,
		PostID:   postID,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) DeletePost(seq int64, title, author, postID, privKeyHex string) error {
	msg := model.DeletePostMsg{
		Author: author,
		PostID: postID,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) UpdatePost(seq int64, author, postID, title, content, redistributionSplitRate, privKeyHex string, links map[string]string) error {
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
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

// Validator related tx
func (broadcast *Broadcast) ValidatorDeposit(seq int64, username, deposit, validatorPubKey, link, privKeyHex string) error {
	valPubKey, err := transport.GetPubKeyFromHex(validatorPubKey)
	if err != nil {
		return err
	}
	msg := model.ValidatorDepositMsg{
		Username:  username,
		Deposit:   deposit,
		ValPubKey: valPubKey,
		Link:      link,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) ValidatorWithdraw(seq int64, username, amount, privKeyHex string) error {
	msg := model.ValidatorWithdrawMsg{
		Username: username,
		Amount:   amount,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) ValidatorRevoke(seq int64, username, privKeyHex string) error {
	msg := model.ValidatorRevokeMsg{
		Username: username,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

// Vote related tx
func (broadcast *Broadcast) Vote(seq int64, voter, proposalID, privKeyHex string, result bool) error {
	msg := model.VoteMsg{
		Voter:      voter,
		ProposalID: proposalID,
		Result:     result,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) VoterDeposit(seq int64, username, deposit, privKeyHex string) error {
	msg := model.VoterDepositMsg{
		Username: username,
		Deposit:  deposit,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) VoterWithdraw(seq int64, username, amount, privKeyHex string) error {
	msg := model.VoterWithdrawMsg{
		Username: username,
		Amount:   amount,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) VoterRevoke(seq int64, username, privKeyHex string) error {
	msg := model.VoterRevokeMsg{
		Username: username,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) Delegate(seq int64, delegator, voter, amount, privKeyHex string) error {
	msg := model.DelegateMsg{
		Delegator: delegator,
		Voter:     voter,
		Amount:    amount,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) DelegatorWithdraw(seq int64, delegator, voter, amount, privKeyHex string) error {
	msg := model.DelegatorWithdrawMsg{
		Delegator: delegator,
		Voter:     voter,
		Amount:    amount,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) RevokeDelegation(seq int64, delegator, voter, privKeyHex string) error {
	msg := model.RevokeDelegationMsg{
		Delegator: delegator,
		Voter:     voter,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

// Developer related tx
func (broadcast *Broadcast) DeveloperRegister(seq int64, username, deposit, privKeyHex string) error {
	msg := model.DeveloperRegisterMsg{
		Username: username,
		Deposit:  deposit,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) DeveloperRevoke(seq int64, username, privKeyHex string) error {
	msg := model.DeveloperRegisterMsg{
		Username: username,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) GrantDeveloper(seq int64, username, authenticateApp, privKeyHex string, validityPeriod, grantLevel int64) error {
	msg := model.GrantDeveloperMsg{
		Username:        username,
		AuthenticateApp: authenticateApp,
		ValidityPeriod:  validityPeriod,
		GrantLevel:      grantLevel,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

// infra related tx
func (broadcast *Broadcast) ProviderReport(seq int64, username, privKeyHex string, usage int64) error {
	msg := model.ProviderReportMsg{
		Username: username,
		Usage:    usage,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

// proposal related tx
func (broadcast *Broadcast) ChangeGlobalAllocationParam(seq int64, creator, privKeyHex string, parameter model.GlobalAllocationParam) error {
	msg := model.ChangeGlobalAllocationParamMsg{
		Creator:   creator,
		Parameter: parameter,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) ChangeEvaluateOfContentValueParam(seq int64, creator, privKeyHex string, parameter model.EvaluateOfContentValueParam) error {
	msg := model.ChangeEvaluateOfContentValueParamMsg{
		Creator:   creator,
		Parameter: parameter,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) ChangeInfraInternalAllocationParam(seq int64, creator, privKeyHex string, parameter model.InfraInternalAllocationParam) error {
	msg := model.ChangeInfraInternalAllocationParamMsg{
		Creator:   creator,
		Parameter: parameter,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) ChangeVoteParam(seq int64, creator, privKeyHex string, parameter model.VoteParam) error {
	msg := model.ChangeVoteParamMsg{
		Creator:   creator,
		Parameter: parameter,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) ChangeProposalParam(seq int64, creator, privKeyHex string, parameter model.ProposalParam) error {
	msg := model.ChangeProposalParamMsg{
		Creator:   creator,
		Parameter: parameter,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) ChangeDeveloperParam(seq int64, creator, privKeyHex string, parameter model.DeveloperParam) error {
	msg := model.ChangeDeveloperParamMsg{
		Creator:   creator,
		Parameter: parameter,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) ChangeValidatorParam(seq int64, creator, privKeyHex string, parameter model.ValidatorParam) error {
	msg := model.ChangeValidatorParamMsg{
		Creator:   creator,
		Parameter: parameter,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) ChangeCoinDayParam(seq int64, creator, privKeyHex string, parameter model.CoinDayParam) error {
	msg := model.ChangeCoinDayParamMsg{
		Creator:   creator,
		Parameter: parameter,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) ChangeBandwidthParam(seq int64, creator, privKeyHex string, parameter model.BandwidthParam) error {
	msg := model.ChangeBandwidthParamMsg{
		Creator:   creator,
		Parameter: parameter,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) ChangeAccountParam(seq int64, creator, privKeyHex string, parameter model.AccountParam) error {
	msg := model.ChangeAccountParamMsg{
		Creator:   creator,
		Parameter: parameter,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

func (broadcast *Broadcast) DeletePostContent(seq int64, creator, permLink, privKeyHex string) error {
	msg := model.DeletePostContentMsg{
		Creator:  creator,
		PermLink: permLink,
	}
	return broadcast.broadcastTransaction(seq, msg, privKeyHex)
}

// internal helper functions
func (broadcast *Broadcast) broadcastTransaction(seq int64, transaction interface{}, privKeyHex string) error {
	res, err := broadcast.transport.SignBuildBroadcast(transaction, privKeyHex, seq)
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
