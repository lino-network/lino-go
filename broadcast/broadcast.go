// Pacakge broadcast includes the functionalities to broadcast
// all kinds of transactions to blockchain.
package broadcast

import (
	"context"
	"encoding/hex"
	"strings"
	"time"

	"github.com/lino-network/lino-go/errors"
	"github.com/lino-network/lino-go/model"
	"github.com/lino-network/lino-go/transport"

	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// Broadcast is a wrapper of broadcasting transactions to blockchain.
type Broadcast struct {
	transport *transport.Transport
	timeout   time.Duration
}

// NewBroadcast returns an instance of Broadcast.
func NewBroadcast(transport *transport.Transport, timeout time.Duration) *Broadcast {
	return &Broadcast{
		transport: transport,
		timeout:   timeout,
	}
}

//
// Account related tx
//

// Register registers a new user on blockchain.
// It composes RegisterMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) Register(referrer, registerFee, username, resetPubKeyHex,
	transactionPubKeyHex, appPubKeyHex, referrerPrivKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	resetPubKey, err := transport.GetPubKeyFromHex(resetPubKeyHex)
	if err != nil {
		return nil, errors.FailedToGetPubKeyFromHex("Register: failed to get Reset pub key").AddCause(err)
	}
	txPubKey, err := transport.GetPubKeyFromHex(transactionPubKeyHex)
	if err != nil {
		return nil, errors.FailedToGetPubKeyFromHex("Register: failed to get Tx pub key").AddCause(err)
	}
	appPubKey, err := transport.GetPubKeyFromHex(appPubKeyHex)
	if err != nil {
		return nil, errors.FailedToGetPubKeyFromHex("Register: failed to get App pub key").AddCause(err)
	}

	msg := model.RegisterMsg{
		Referrer:             referrer,
		RegisterFee:          registerFee,
		NewUser:              username,
		NewResetPubKey:       resetPubKey,
		NewTransactionPubKey: txPubKey,
		NewAppPubKey:         appPubKey,
	}
	return broadcast.broadcastTransaction(msg, referrerPrivKeyHex, seq, "")
}

// Transfer sends a certain amount of LINO token from the sender to the receiver.
// It composes TransferMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) Transfer(sender, receiver, amount, memo,
	privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.TransferMsg{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
		Memo:     memo,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// Follow creates a social relationship between follower and followee.
// It composes FollowMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) Follow(follower, followee,
	privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.FollowMsg{
		Follower: follower,
		Followee: followee,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// Unfollow revokes the social relationship between follower and followee.
// It composes UnfollowMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) Unfollow(follower, followee,
	privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.UnfollowMsg{
		Follower: follower,
		Followee: followee,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// Claim claims rewards of a certain user.
// It composes ClaimMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) Claim(username,
	privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.ClaimMsg{
		Username: username,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// UpdateAccount updates account related info in jsonMeta which are not
// included in AccountInfo or AccountBank.
// It composes UpdateAccountMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) UpdateAccount(username, jsonMeta,
	privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.UpdateAccountMsg{
		Username: username,
		JSONMeta: jsonMeta,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// Recover recovers all keys of a user in case of losing or compromising.
// It composes RecoverMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) Recover(username, newResetPubKeyHex,
	newTransactionPubKeyHex, newAppPubKeyHex, privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	resetPubKey, err := transport.GetPubKeyFromHex(newResetPubKeyHex)
	if err != nil {
		return nil, errors.FailedToGetPubKeyFromHexf("Recover: failed to get Reset pub key").AddCause(err)
	}
	txPubKey, err := transport.GetPubKeyFromHex(newTransactionPubKeyHex)
	if err != nil {
		return nil, errors.FailedToGetPubKeyFromHexf("Recover: failed to get Tx pub key").AddCause(err)
	}
	appPubKey, err := transport.GetPubKeyFromHex(newAppPubKeyHex)
	if err != nil {
		return nil, errors.FailedToGetPubKeyFromHexf("Recover: failed to get App pub key").AddCause(err)
	}

	msg := model.RecoverMsg{
		Username:             username,
		NewResetPubKey:       resetPubKey,
		NewTransactionPubKey: txPubKey,
		NewAppPubKey:         appPubKey,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

//
// Post related tx
//

// CreatePost creates a new post on blockchain.
// It composes CreatePostMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) CreatePost(author, postID, title, content, parentAuthor, parentPostID,
	sourceAuthor, sourcePostID, redistributionSplitRate string,
	links map[string]string, privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
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
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// Donate adds a money donation to a post by a user.
// It composes DonateMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) Donate(username, author,
	amount, postID, fromApp, memo string, privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.DonateMsg{
		Username: username,
		Amount:   amount,
		Author:   author,
		PostID:   postID,
		FromApp:  fromApp,
		Memo:     memo,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// ReportOrUpvote adds a report or upvote action to a post.
// It composes ReportOrUpvoteMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ReportOrUpvote(username, author,
	postID string, isReport bool, privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.ReportOrUpvoteMsg{
		Username: username,
		Author:   author,
		PostID:   postID,
		IsReport: isReport,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// DeletePost deletes a post from the blockchain. It doesn't actually
// remove the post from the blockchain, instead it sets IsDeleted to true
// and clears all the other data.
// It composes DeletePostMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) DeletePost(author, postID,
	privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.DeletePostMsg{
		Author: author,
		PostID: postID,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// View increases the view count of a post by one.
// It composes ViewMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) View(username, author, postID,
	privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.ViewMsg{
		Username: username,
		Author:   author,
		PostID:   postID,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// UpdatePost updates post info with new data.
// It composes UpdatePostMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) UpdatePost(author, title, postID, content string,
	links map[string]string, privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
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
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

//
// Validator related tx
//

// ValidatorDeposit deposits a certain amount of LINO token for a user
// in order to become a validator. Before becoming a validator, the user
// has to be a voter.
// It composes ValidatorDepositMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ValidatorDeposit(username, deposit,
	validatorPubKey, link, privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	valPubKey, err := transport.GetPubKeyFromHex(validatorPubKey)
	if err != nil {
		return nil, errors.FailedToGetPubKeyFromHexf("ValidatorDeposit: failed to get Val pub key").AddCause(err)
	}
	msg := model.ValidatorDepositMsg{
		Username:  username,
		Deposit:   deposit,
		ValPubKey: valPubKey,
		Link:      link,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// ValidatorWithdraw withdraws part of LINO token from a validator's deposit,
// while still keep being a validator.
// It composes ValidatorDepositMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ValidatorWithdraw(username, amount,
	privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.ValidatorWithdrawMsg{
		Username: username,
		Amount:   amount,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// ValidatorRevoke revokes all deposited LINO token of a validator
// so that the user will not be a validator anymore.
// It composes ValidatorRevokeMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ValidatorRevoke(username,
	privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.ValidatorRevokeMsg{
		Username: username,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

//
// Vote related tx
//

// StakeIn deposits a certain amount of LINO token for a user
// in order to become a voter.
// It composes StakeInMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) StakeIn(username, deposit,
	privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.StakeInMsg{
		Username: username,
		Deposit:  deposit,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// StakeOut withdraws part of LINO token from a voter's deposit.
// It composes StakeOutMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) StakeOut(username, amount,
	privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.StakeOutMsg{
		Username: username,
		Amount:   amount,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// Delegate delegates a certain amount of LINO token of delegator to a voter, so
// the voter will have more voting power.
// It composes DelegateMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) Delegate(delegator, voter, amount,
	privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.DelegateMsg{
		Delegator: delegator,
		Voter:     voter,
		Amount:    amount,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// DelegatorWithdraw withdraws part of delegated LINO token of a delegator
// to a voter, while the delegation still exists.
// It composes DelegatorWithdrawMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) DelegatorWithdraw(delegator, voter, amount,
	privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.DelegatorWithdrawMsg{
		Delegator: delegator,
		Voter:     voter,
		Amount:    amount,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// ClaimInterest claims interest of a certain user.
// It composes ClaimInterestMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ClaimInterest(username, privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.ClaimInterestMsg{
		Username: username,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

//
// Developer related tx
//

// DeveloperRegsiter registers a developer with a certain amount of LINO token on blockchain.
// It composes DeveloperRegisterMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) DeveloperRegister(username, deposit, website,
	description, appMetaData, privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.DeveloperRegisterMsg{
		Username:    username,
		Deposit:     deposit,
		Website:     website,
		Description: description,
		AppMetaData: appMetaData,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// DeveloperUpdate updates a developer  info on blockchain.
// It composes DeveloperUpdateMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) DeveloperUpdate(username, website,
	description, appMetaData, privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.DeveloperUpdateMsg{
		Username:    username,
		Website:     website,
		Description: description,
		AppMetaData: appMetaData,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// DeveloperRevoke reovkes all deposited LINO token of a developer
// so the user will not be a developer anymore.
// It composes DeveloperRevokeMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) DeveloperRevoke(username, privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.DeveloperRevokeMsg{
		Username: username,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// GrantPermission grants a certain (e.g. App) permission to
// an authorized app with a certain period of time.
// It composes GrantPermissionMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) GrantPermission(username, authorizedApp string,
	validityPeriodSec int64, grantLevel model.Permission, privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.GrantPermissionMsg{
		Username:          username,
		AuthorizedApp:     authorizedApp,
		ValidityPeriodSec: validityPeriodSec,
		GrantLevel:        grantLevel,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// PreAuthorizationPermission grants a PreAuthorization permission to
// an authorzied app with a certain period of time.
// It composes PreAuthorizationMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) PreAuthorizationPermission(username, authorizedApp string,
	validityPeriodSec int64, amount string, privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.PreAuthorizationMsg{
		Username:          username,
		AuthorizedApp:     authorizedApp,
		ValidityPeriodSec: validityPeriodSec,
		Amount:            amount,
	}

	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// RevokePermission revokes the permission given previously to a app.
// It composes RevokePermissionMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) RevokePermission(username, pubKeyHex string,
	privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	pubKey, err := transport.GetPubKeyFromHex(pubKeyHex)
	if err != nil {
		return nil, errors.FailedToGetPubKeyFromHex("Register: failed to get pub key").AddCause(err)
	}

	msg := model.RevokePermissionMsg{
		Username: username,
		PubKey:   pubKey,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

//
// infra related tx
//

// ProviderReport reports infra usage of a infra provider in order to get infra inflation.
// It composes ProviderReportMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ProviderReport(username string, usage int64,
	privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.ProviderReportMsg{
		Username: username,
		Usage:    usage,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

//
// proposal related tx
//

// ChangeEvaluateOfContentValueParam changes EvaluateOfContentValueParam with new value.
// It composes ChangeEvaluateOfContentValueParamMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ChangeEvaluateOfContentValueParam(creator string,
	parameter model.EvaluateOfContentValueParam, reason string, privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.ChangeEvaluateOfContentValueParamMsg{
		Creator:   creator,
		Parameter: parameter,
		Reason:    reason,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// ChangeGlobalAllocationParam changes GlobalAllocationParam with new value.
// It composes ChangeGlobalAllocationParamMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ChangeGlobalAllocationParam(creator string,
	parameter model.GlobalAllocationParam, reason string, privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.ChangeGlobalAllocationParamMsg{
		Creator:   creator,
		Parameter: parameter,
		Reason:    reason,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// ChangeInfraInternalAllocationParam changes InfraInternalAllocationParam with new value.
// It composes ChangeInfraInternalAllocationParamMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ChangeInfraInternalAllocationParam(creator string,
	parameter model.InfraInternalAllocationParam,
	reason string, privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.ChangeInfraInternalAllocationParamMsg{
		Creator:   creator,
		Parameter: parameter,
		Reason:    reason,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// ChangeVoteParam changes VoteParam with new value.
// It composes ChangeVoteParamMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ChangeVoteParam(creator string,
	parameter model.VoteParam, reason string, privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.ChangeVoteParamMsg{
		Creator:   creator,
		Parameter: parameter,
		Reason:    reason,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// ChangeProposalParam changes ProposalParam with new value.
// It composes ChangeProposalParamMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ChangeProposalParam(creator string,
	parameter model.ProposalParam, reason string, privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.ChangeProposalParamMsg{
		Creator:   creator,
		Parameter: parameter,
		Reason:    reason,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// ChangeDeveloperParam changes DeveloperParam with new value.
// It composes ChangeDeveloperParamMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ChangeDeveloperParam(creator string,
	parameter model.DeveloperParam, reason string, privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.ChangeDeveloperParamMsg{
		Creator:   creator,
		Parameter: parameter,
		Reason:    reason,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// ChangeValidatorParam changes ValidatorParam with new value.
// It composes ChangeValidatorParamMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ChangeValidatorParam(creator string,
	parameter model.ValidatorParam, reason string, privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.ChangeValidatorParamMsg{
		Creator:   creator,
		Parameter: parameter,
		Reason:    reason,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// ChangeBandwidthParam changes BandwidthParam with new value.
// It composes ChangeBandwidthParamMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ChangeBandwidthParam(creator string,
	parameter model.BandwidthParam, reason string, privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.ChangeBandwidthParamMsg{
		Creator:   creator,
		Parameter: parameter,
		Reason:    reason,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// ChangeAccountParam changes AccountParam with new value.
// It composes ChangeAccountParamMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ChangeAccountParam(creator string,
	parameter model.AccountParam, reason string, privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.ChangeAccountParamMsg{
		Creator:   creator,
		Parameter: parameter,
		Reason:    reason,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// ChangePostParam changes PostParam with new value.
// It composes ChangePostParamMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ChangePostParam(creator string,
	parameter model.PostParam, reason string, privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.ChangePostParamMsg{
		Creator:   creator,
		Parameter: parameter,
		Reason:    reason,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// DeletePostContent deletes the content of a post on blockchain, which is used
// for content censorship.
// It composes DeletePostContentMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) DeletePostContent(creator, postAuthor,
	postID, reason, privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	permlink := string(string(postAuthor) + "#" + postID)
	msg := model.DeletePostContentMsg{
		Creator:  creator,
		Permlink: permlink,
		Reason:   reason,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// VoteProposal adds a vote to a certain proposal with agree/disagree.
// It composes VoteProposalMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) VoteProposal(voter, proposalID string,
	result bool, privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.VoteProposalMsg{
		Voter:      voter,
		ProposalID: proposalID,
		Result:     result,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

// UpgradeProtocol upgrades the protocol.
// It composes UpgradeProtocolMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) UpgradeProtocol(creator, link, reason string,
	privKeyHex string, seq int64) (*model.BroadcastReponse, error) {
	msg := model.UpgradeProtocolMsg{
		Creator: creator,
		Link:    link,
		Reason:  reason,
	}
	return broadcast.broadcastTransaction(msg, privKeyHex, seq, "")
}

//
// internal helper functions
//
func (broadcast *Broadcast) broadcastTransaction(msg model.Msg, privKeyHex string,
	seq int64, memo string) (*model.BroadcastReponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), broadcast.timeout)
	defer cancel()

	broadcastResp := &model.BroadcastReponse{}

	var res *ctypes.ResultBroadcastTxCommit
	var err error
	finishChan := make(chan bool)
	go func() {
		res, err = broadcast.transport.SignBuildBroadcast(msg, privKeyHex, seq, memo)
		finishChan <- true
	}()

	select {
	case <-finishChan:
		break
	case <-ctx.Done():
		return nil, errors.Timeoutf("msg timeout: %v", msg).AddCause(ctx.Err())
	}

	if err != nil {
		return nil, err
	}

	code := retrieveCodeFromBlockChainCode(res.CheckTx.Code)
	if err == nil && code == model.InvalidSeqErrCode {
		return nil, errors.InvalidSequenceNumber("invalid seq").AddBlockChainCode(res.CheckTx.Code).AddBlockChainLog(res.CheckTx.Log)
	}

	if res.CheckTx.Code != uint32(0) {
		return nil, errors.CheckTxFail("CheckTx failed!").AddBlockChainCode(res.CheckTx.Code).AddBlockChainLog(res.CheckTx.Log)
	}
	if res.DeliverTx.Code != uint32(0) {
		return nil, errors.DeliverTxFail("DeliverTx failed!").AddBlockChainCode(res.DeliverTx.Code).AddBlockChainLog(res.DeliverTx.Log)
	}

	commitHash := hex.EncodeToString(res.Hash)
	broadcastResp.CommitHash = strings.ToUpper(commitHash)

	return broadcastResp, nil
}

func retrieveCodeFromBlockChainCode(bcCode uint32) uint32 {
	return bcCode & 0xff
}
