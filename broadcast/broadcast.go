package broadcast

import (
	"github.com/lino-network/lino-go/errors"
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

//
// Account related tx
//
func (broadcast *Broadcast) Register(referrer, registerFee, username, masterPubKeyHex, transactionPubKeyHex, postPubKeyHex, referrerPrivKeyHex string, seq int64) errors.Error {
	masterPubKey, err := transport.GetPubKeyFromHex(masterPubKeyHex)
	if err != nil {
		return errors.FailedToGetPubKeyFromHexf("Register: failed to get Master pub key from hex for %v", masterPubKeyHex)
	}
	txPubKey, err := transport.GetPubKeyFromHex(transactionPubKeyHex)
	if err != nil {
		return errors.FailedToGetPubKeyFromHexf("Register: failed to get Tx pub key from hex for %v", transactionPubKeyHex)
	}
	postPubKey, err := transport.GetPubKeyFromHex(postPubKeyHex)
	if err != nil {
		return errors.FailedToGetPubKeyFromHexf("Register: failed to get Post pub key from hex for %v", postPubKeyHex)
	}

	msg := model.RegisterMsg{
		Referrer:             referrer,
		RegisterFee:          registerFee,
		NewUser:              username,
		NewMasterPubKey:      masterPubKey,
		NewTransactionPubKey: txPubKey,
		NewPostPubKey:        postPubKey,
	}
	return broadcast.broadcastTransaction(msg, referrerPrivKeyHex, seq)
}

func (broadcast *Broadcast) Transfer(sender, receiver, amount, memo, privKeyHex string, seq int64) errors.Error {
	msg := model.TransferMsg{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
		Memo:     memo,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) Follow(follower, followee, privKeyHex string, seq int64) errors.Error {
	msg := model.FollowMsg{
		Follower: follower,
		Followee: followee,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) Unfollow(follower, followee, privKeyHex string, seq int64) errors.Error {
	msg := model.UnfollowMsg{
		Follower: follower,
		Followee: followee,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) Claim(username, privKeyHex string, seq int64) errors.Error {
	msg := model.ClaimMsg{
		Username: username,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) UpdateAccount(username, jsonMeta, privKeyHex string, seq int64) errors.Error {
	msg := model.UpdateAccountMsg{
		Username: username,
		JSONMeta: jsonMeta,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) Recover(username, newMasterPubKeyHex, newPostPubKeyHex, newTransactionPubKeyHex, privKeyHex string, seq int64) errors.Error {
	masterPubKey, err := transport.GetPubKeyFromHex(newMasterPubKeyHex)
	if err != nil {
		return errors.FailedToGetPubKeyFromHexf("Recover: failed to get Master pub key from hex for %v", newMasterPubKeyHex)
	}
	postPubKey, err := transport.GetPubKeyFromHex(newPostPubKeyHex)
	if err != nil {
		return errors.FailedToGetPubKeyFromHexf("Recover: failed to get Post pub key from hex for %v", newPostPubKeyHex)
	}
	txPubKey, err := transport.GetPubKeyFromHex(newTransactionPubKeyHex)
	if err != nil {
		return errors.FailedToGetPubKeyFromHexf("Recover: failed to get Tx pub key from hex for %v", newTransactionPubKeyHex)
	}

	msg := model.RecoverMsg{
		Username:             username,
		NewMasterPubKey:      masterPubKey,
		NewPostPubKey:        postPubKey,
		NewTransactionPubKey: txPubKey,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

//
// Post related tx
//
func (broadcast *Broadcast) CreatePost(author, postID, title, content, parentAuthor, parentPostID,
	sourceAuthor, sourcePostID, redistributionSplitRate string, links map[string]string, privKeyHex string, seq int64) errors.Error {
	var mLinks []model.IDToURLMapping
	if links == nil || len(links) == 0 {
		mLinks = nil
	} else {
		for k, v := range links {
			mLinks = append(mLinks, model.IDToURLMapping{k, v})
		}
	}

	msg := model.CreatePostMsg{
		Author:       author,
		PostID:       postID,
		Title:        title,
		Content:      content,
		ParentAuthor: parentAuthor,
		ParentPostID: parentPostID,
		SourceAuthor: sourceAuthor,
		SourcePostID: sourcePostID,
		Links:        mLinks,
		RedistributionSplitRate: redistributionSplitRate,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) Like(username, author string, weight int64, postID, privKeyHex string, seq int64) errors.Error {
	msg := model.LikeMsg{
		Username: username,
		Weight:   weight,
		Author:   author,
		PostID:   postID,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) Donate(username, author, amount, postID, fromApp, memo, privKeyHex string, seq int64) errors.Error {
	msg := model.DonateMsg{
		Username: username,
		Amount:   amount,
		Author:   author,
		PostID:   postID,
		FromApp:  fromApp,
		Memo:     memo,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) ReportOrUpvote(username, author, postID string, isReport bool, privKeyHex string, seq int64) errors.Error {
	msg := model.ReportOrUpvoteMsg{
		Username: username,
		Author:   author,
		PostID:   postID,
		IsReport: isReport,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) DeletePost(author, postID, privKeyHex string, seq int64) errors.Error {
	msg := model.DeletePostMsg{
		Author: author,
		PostID: postID,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) View(username, author, postID, privKeyHex string, seq int64) errors.Error {
	msg := model.ViewMsg{
		Username: username,
		Author:   author,
		PostID:   postID,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) UpdatePost(author, title, postID, content, redistributionSplitRate string, links map[string]string, privKeyHex string, seq int64) errors.Error {
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
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

//
// Validator related tx
//
func (broadcast *Broadcast) ValidatorDeposit(username, deposit, validatorPubKey, link, privKeyHex string, seq int64) errors.Error {
	valPubKey, err := transport.GetPubKeyFromHex(validatorPubKey)
	if err != nil {
		return errors.FailedToGetPubKeyFromHexf("ValidatorDeposit: failed to get Val pub key from hex for %v", validatorPubKey)
	}
	msg := model.ValidatorDepositMsg{
		Username:  username,
		Deposit:   deposit,
		ValPubKey: valPubKey,
		Link:      link,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) ValidatorWithdraw(username, amount, privKeyHex string, seq int64) errors.Error {
	msg := model.ValidatorWithdrawMsg{
		Username: username,
		Amount:   amount,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) ValidatorRevoke(username, privKeyHex string, seq int64) errors.Error {
	msg := model.ValidatorRevokeMsg{
		Username: username,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

//
// Vote related tx
//
func (broadcast *Broadcast) VoterDeposit(username, deposit, privKeyHex string, seq int64) errors.Error {
	msg := model.VoterDepositMsg{
		Username: username,
		Deposit:  deposit,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) VoterWithdraw(username, amount, privKeyHex string, seq int64) errors.Error {
	msg := model.VoterWithdrawMsg{
		Username: username,
		Amount:   amount,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) VoterRevoke(username, privKeyHex string, seq int64) errors.Error {
	msg := model.VoterRevokeMsg{
		Username: username,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) Delegate(delegator, voter, amount, privKeyHex string, seq int64) errors.Error {
	msg := model.DelegateMsg{
		Delegator: delegator,
		Voter:     voter,
		Amount:    amount,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) DelegatorWithdraw(delegator, voter, amount, privKeyHex string, seq int64) errors.Error {
	msg := model.DelegatorWithdrawMsg{
		Delegator: delegator,
		Voter:     voter,
		Amount:    amount,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) RevokeDelegation(delegator, voter, privKeyHex string, seq int64) errors.Error {
	msg := model.RevokeDelegationMsg{
		Delegator: delegator,
		Voter:     voter,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

//
// Developer related tx
//
func (broadcast *Broadcast) DeveloperRegister(username, deposit, privKeyHex string, seq int64) errors.Error {
	msg := model.DeveloperRegisterMsg{
		Username: username,
		Deposit:  deposit,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) DeveloperRevoke(username, privKeyHex string, seq int64) errors.Error {
	msg := model.DeveloperRegisterMsg{
		Username: username,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) GrantDeveloper(username, authenticateApp string, validityPeriod int64, grantLevel int, privKeyHex string, seq int64) errors.Error {
	msg := model.GrantDeveloperMsg{
		Username:        username,
		AuthenticateApp: authenticateApp,
		ValidityPeriod:  validityPeriod,
		GrantLevel:      grantLevel,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

//
// infra related tx
//
func (broadcast *Broadcast) ProviderReport(username string, usage int64, privKeyHex string, seq int64) errors.Error {
	msg := model.ProviderReportMsg{
		Username: username,
		Usage:    usage,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

//
// proposal related tx
//

func (broadcast *Broadcast) ChangeEvaluateOfContentValueParam(creator string, parameter model.EvaluateOfContentValueParam, privKeyHex string, seq int64) errors.Error {
	msg := model.ChangeEvaluateOfContentValueParamMsg{
		Creator:   creator,
		Parameter: parameter,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) ChangeGlobalAllocationParam(creator string, parameter model.GlobalAllocationParam, privKeyHex string, seq int64) errors.Error {
	msg := model.ChangeGlobalAllocationParamMsg{
		Creator:   creator,
		Parameter: parameter,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) ChangeInfraInternalAllocationParam(creator string, parameter model.InfraInternalAllocationParam, privKeyHex string, seq int64) errors.Error {
	msg := model.ChangeInfraInternalAllocationParamMsg{
		Creator:   creator,
		Parameter: parameter,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) ChangeVoteParam(creator string, parameter model.VoteParam, privKeyHex string, seq int64) errors.Error {
	msg := model.ChangeVoteParamMsg{
		Creator:   creator,
		Parameter: parameter,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) ChangeProposalParam(creator string, parameter model.ProposalParam, privKeyHex string, seq int64) errors.Error {
	msg := model.ChangeProposalParamMsg{
		Creator:   creator,
		Parameter: parameter,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) ChangeDeveloperParam(creator string, parameter model.DeveloperParam, privKeyHex string, seq int64) errors.Error {
	msg := model.ChangeDeveloperParamMsg{
		Creator:   creator,
		Parameter: parameter,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) ChangeValidatorParam(creator string, parameter model.ValidatorParam, privKeyHex string, seq int64) errors.Error {
	msg := model.ChangeValidatorParamMsg{
		Creator:   creator,
		Parameter: parameter,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) ChangeCoinDayParam(creator string, parameter model.CoinDayParam, privKeyHex string, seq int64) errors.Error {
	msg := model.ChangeCoinDayParamMsg{
		Creator:   creator,
		Parameter: parameter,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) ChangeBandwidthParam(creator string, parameter model.BandwidthParam, privKeyHex string, seq int64) errors.Error {
	msg := model.ChangeBandwidthParamMsg{
		Creator:   creator,
		Parameter: parameter,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) ChangeAccountParam(creator string, parameter model.AccountParam, privKeyHex string, seq int64) errors.Error {
	msg := model.ChangeAccountParamMsg{
		Creator:   creator,
		Parameter: parameter,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) DeletePostContent(creator, postAuthor, postID, reason, privKeyHex string, seq int64) errors.Error {
	permLink := string(string(postAuthor) + "#" + postID)
	msg := model.DeletePostContentMsg{
		Creator:  creator,
		PermLink: permLink,
		Reason:   reason,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

func (broadcast *Broadcast) VoteProposal(voter, proposalID string, result bool, privKeyHex string, seq int64) errors.Error {
	msg := model.VoteProposalMsg{
		Voter:      voter,
		ProposalID: proposalID,
		Result:     result,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq)
}

//
// internal helper functions
//
func (broadcast *Broadcast) broadcastTransaction(msg interface{}, privKeyHex string, seq int64) errors.Error {
	res, err := broadcast.transport.SignBuildBroadcast(msg, privKeyHex, seq)
	if err != nil {
		return errors.FailedToBroadcastf("failed to broadcast msg: %v, got err: %v", msg, err)
	}

	if err == nil && res.CheckTx.Code == types.InvalidSeqErrCode {
		return errors.InvalidSeqNumberf("msg: %v, details: %v", seq, res.CheckTx.Log)
	}

	if res.CheckTx.Code != uint32(0) {
		return errors.CheckTxFailf("CheckTx failed! (%d) %s", res.CheckTx.Code, res.CheckTx.Log)
	}
	if res.DeliverTx.Code != uint32(0) {
		return errors.DeliverTxFailf("DeliverTx failed! (%d) %s", res.DeliverTx.Code, res.DeliverTx.Log)
	}
	return nil
}
