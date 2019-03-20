// Pacakge broadcast includes the functionalities to broadcast
// all kinds of transactions to blockchain.
package broadcast

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"encoding/hex"

	"github.com/lino-network/lino-go/errors"
	"github.com/lino-network/lino-go/model"
	"github.com/lino-network/lino-go/transport"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	ttypes "github.com/tendermint/tendermint/types"
)

// Broadcast is a wrapper of broadcasting transactions to blockchain.
type Broadcast struct {
	FixSequenceNumber  bool
	transport          *transport.Transport
	maxAttempts        int64
	initSleepTime      time.Duration
	timeout            time.Duration
	exponentialBackoff bool
	backoffRandomness  bool
}

// NewBroadcast returns an instance of Broadcast.
func NewBroadcast(
	transport *transport.Transport, maxAttempts int64, initSleepTime time.Duration,
	timeout time.Duration, exponentialBackoff bool, backoffRandomness bool, fixSequenceNumber bool) *Broadcast {
	return &Broadcast{
		transport:          transport,
		maxAttempts:        maxAttempts,
		timeout:            timeout,
		initSleepTime:      initSleepTime,
		exponentialBackoff: exponentialBackoff,
		backoffRandomness:  backoffRandomness,
		FixSequenceNumber:  fixSequenceNumber,
	}
}

//
// Account related tx
//

// Register registers a new user on blockchain.
// It composes RegisterMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) Register(ctx context.Context, referrer, registerFee, username, resetPubKeyHex,
	transactionPubKeyHex, appPubKeyHex, referrerPrivKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
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
	return broadcast.retry(ctx, msg, referrerPrivKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// Transfer sends a certain amount of LINO token from the sender to the receiver.
// It composes TransferMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) Transfer(ctx context.Context, sender, receiver, amount, memo,
	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.TransferMsg{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
		Memo:     memo,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// Transfer sends a certain amount of LINO token from the sender to the receiver.
// It composes TransferMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) TransferSync(ctx context.Context, sender, receiver, amount, memo,
	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.TransferMsg{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
		Memo:     memo,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, strconv.FormatInt(time.Now().Unix(), 10), true, broadcast.maxAttempts, broadcast.initSleepTime)
}

// Follow creates a social relationship between follower and followee.
// It composes FollowMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) Follow(ctx context.Context, follower, followee,
	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.FollowMsg{
		Follower: follower,
		Followee: followee,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// Unfollow revokes the social relationship between follower and followee.
// It composes UnfollowMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) Unfollow(ctx context.Context, follower, followee,
	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.UnfollowMsg{
		Follower: follower,
		Followee: followee,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// Claim claims rewards of a certain user.
// It composes ClaimMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) Claim(ctx context.Context, username,
	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.ClaimMsg{
		Username: username,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// UpdateAccount updates account related info in jsonMeta which are not
// included in AccountInfo or AccountBank.
// It composes UpdateAccountMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) UpdateAccount(ctx context.Context, username, jsonMeta,
	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.UpdateAccountMsg{
		Username: username,
		JSONMeta: jsonMeta,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// Recover recovers all keys of a user in case of losing or compromising.
// It composes RecoverMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) Recover(ctx context.Context, username, newResetPubKeyHex,
	newTransactionPubKeyHex, newAppPubKeyHex, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
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
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

//
// Post related tx
//

// CreatePost creates a new post on blockchain.
// It composes CreatePostMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) CreatePost(ctx context.Context, author, postID, title, content,
	parentAuthor, parentPostID, sourceAuthor, sourcePostID, redistributionSplitRate string,
	links map[string]string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
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
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// CreatePost creates a new post on blockchain.
// It composes CreatePostMsg and then broadcasts the transaction to blockchain return when checkTx pass.
func (broadcast *Broadcast) CreatePostSync(ctx context.Context, author, postID, title, content,
	parentAuthor, parentPostID, sourceAuthor, sourcePostID, redistributionSplitRate string,
	links map[string]string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
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
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", true, broadcast.maxAttempts, broadcast.initSleepTime)
}

// Donate adds a money donation to a post by a user.
// It composes DonateMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) Donate(ctx context.Context, username, author,
	amount, postID, fromApp, memo string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.DonateMsg{
		Username: username,
		Amount:   amount,
		Author:   author,
		PostID:   postID,
		FromApp:  fromApp,
		Memo:     memo,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// Donate adds a money donation to a post by a user.
// It composes DonateMsg and then broadcasts the transaction to blockchain return after pass checkTx.
func (broadcast *Broadcast) DonateSync(ctx context.Context, username, author,
	amount, postID, fromApp, memo string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.DonateMsg{
		Username: username,
		Amount:   amount,
		Author:   author,
		PostID:   postID,
		FromApp:  fromApp,
		Memo:     memo,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, strconv.FormatInt(time.Now().Unix(), 10), true, broadcast.maxAttempts, broadcast.initSleepTime)
}

// ReportOrUpvote adds a report or upvote action to a post.
// It composes ReportOrUpvoteMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ReportOrUpvote(ctx context.Context, username, author,
	postID string, isReport bool, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.ReportOrUpvoteMsg{
		Username: username,
		Author:   author,
		PostID:   postID,
		IsReport: isReport,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// DeletePost deletes a post from the blockchain. It doesn't actually
// remove the post from the blockchain, instead it sets IsDeleted to true
// and clears all the other data.
// It composes DeletePostMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) DeletePost(ctx context.Context, author, postID,
	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.DeletePostMsg{
		Author: author,
		PostID: postID,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// View increases the view count of a post by one.
// It composes ViewMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) View(ctx context.Context, username, author, postID,
	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.ViewMsg{
		Username: username,
		Author:   author,
		PostID:   postID,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// UpdatePost updates post info with new data.
// It composes UpdatePostMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) UpdatePost(ctx context.Context, author, title, postID, content string,
	links map[string]string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
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
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

//
// Validator related tx
//

// ValidatorDeposit deposits a certain amount of LINO token for a user
// in order to become a validator. Before becoming a validator, the user
// has to be a voter.
// It composes ValidatorDepositMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ValidatorDeposit(ctx context.Context, username, deposit,
	validatorPubKey, link, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
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
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// ValidatorWithdraw withdraws part of LINO token from a validator's deposit,
// while still keep being a validator.
// It composes ValidatorDepositMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ValidatorWithdraw(ctx context.Context, username, amount,
	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.ValidatorWithdrawMsg{
		Username: username,
		Amount:   amount,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// ValidatorRevoke revokes all deposited LINO token of a validator
// so that the user will not be a validator anymore.
// It composes ValidatorRevokeMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ValidatorRevoke(ctx context.Context, username,
	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.ValidatorRevokeMsg{
		Username: username,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

//
// Vote related tx
//

// StakeIn deposits a certain amount of LINO token for a user
// in order to become a voter.
// It composes StakeInMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) StakeIn(ctx context.Context, username, deposit,
	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.StakeInMsg{
		Username: username,
		Deposit:  deposit,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// StakeOut withdraws part of LINO token from a voter's deposit.
// It composes StakeOutMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) StakeOut(ctx context.Context, username, amount,
	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.StakeOutMsg{
		Username: username,
		Amount:   amount,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// Delegate delegates a certain amount of LINO token of delegator to a voter, so
// the voter will have more voting power.
// It composes DelegateMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) Delegate(ctx context.Context, delegator, voter, amount,
	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.DelegateMsg{
		Delegator: delegator,
		Voter:     voter,
		Amount:    amount,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// DelegatorWithdraw withdraws part of delegated LINO token of a delegator
// to a voter, while the delegation still exists.
// It composes DelegatorWithdrawMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) DelegatorWithdraw(ctx context.Context, delegator, voter, amount,
	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.DelegatorWithdrawMsg{
		Delegator: delegator,
		Voter:     voter,
		Amount:    amount,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// ClaimInterest claims interest of a certain user.
// It composes ClaimInterestMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ClaimInterest(ctx context.Context, username,
	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.ClaimInterestMsg{
		Username: username,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

//
// Developer related tx
//

// DeveloperRegsiter registers a developer with a certain amount of LINO token on blockchain.
// It composes DeveloperRegisterMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) DeveloperRegister(ctx context.Context, username, deposit, website,
	description, appMetaData, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.DeveloperRegisterMsg{
		Username:    username,
		Deposit:     deposit,
		Website:     website,
		Description: description,
		AppMetaData: appMetaData,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// DeveloperUpdate updates a developer  info on blockchain.
// It composes DeveloperUpdateMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) DeveloperUpdate(ctx context.Context, username, website,
	description, appMetaData, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.DeveloperUpdateMsg{
		Username:    username,
		Website:     website,
		Description: description,
		AppMetaData: appMetaData,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// DeveloperRevoke reovkes all deposited LINO token of a developer
// so the user will not be a developer anymore.
// It composes DeveloperRevokeMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) DeveloperRevoke(ctx context.Context, username,
	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.DeveloperRevokeMsg{
		Username: username,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// GrantPermission grants a certain (e.g. App) permission to
// an authorized app with a certain period of time.
// It composes GrantPermissionMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) GrantPermission(ctx context.Context, username, authorizedApp string,
	validityPeriodSec int64, grantLevel model.Permission, amount string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.GrantPermissionMsg{
		Username:          username,
		AuthorizedApp:     authorizedApp,
		ValidityPeriodSec: validityPeriodSec,
		GrantLevel:        grantLevel,
		Amount:            amount,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// GrantAppAndPreAuthPermission grants both app and preauth permission to
// an authorized app with a certain period of time.
// It composes GrantPermissionMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) GrantAppAndPreAuthPermission(ctx context.Context, username, authorizedApp string,
	validityPeriodSec int64, amount string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.GrantPermissionMsg{
		Username:          username,
		AuthorizedApp:     authorizedApp,
		ValidityPeriodSec: validityPeriodSec,
		GrantLevel:        model.AppAndPreAuthorizationPermission,
		Amount:            amount,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// PreAuthorizationPermission grants a PreAuthorization permission to
// an authorzied app with a certain period of time.
// It composes GrantPermissionMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) PreAuthorizationPermission(ctx context.Context, username, authorizedApp string,
	validityPeriodSec int64, amount string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.GrantPermissionMsg{
		Username:          username,
		AuthorizedApp:     authorizedApp,
		ValidityPeriodSec: validityPeriodSec,
		GrantLevel:        model.PreAuthorizationPermission,
		Amount:            amount,
	}

	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// RevokePermission revokes the permission given previously to a app.
// It composes RevokePermissionMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) RevokePermission(ctx context.Context, username, pubKeyHex string,
	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	pubKey, err := transport.GetPubKeyFromHex(pubKeyHex)
	if err != nil {
		return nil, errors.FailedToGetPubKeyFromHex("Register: failed to get pub key").AddCause(err)
	}

	msg := model.RevokePermissionMsg{
		Username: username,
		PubKey:   pubKey,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

//
// infra related tx
//

// ProviderReport reports infra usage of a infra provider in order to get infra inflation.
// It composes ProviderReportMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ProviderReport(ctx context.Context, username string, usage int64,
	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.ProviderReportMsg{
		Username: username,
		Usage:    usage,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

//
// proposal related tx
//

// ChangeGlobalAllocationParam changes GlobalAllocationParam with new value.
// It composes ChangeGlobalAllocationParamMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ChangeGlobalAllocationParam(ctx context.Context, creator string,
	parameter model.GlobalAllocationParam, reason string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.ChangeGlobalAllocationParamMsg{
		Creator:   creator,
		Parameter: parameter,
		Reason:    reason,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// ChangeInfraInternalAllocationParam changes InfraInternalAllocationParam with new value.
// It composes ChangeInfraInternalAllocationParamMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ChangeInfraInternalAllocationParam(ctx context.Context, creator string,
	parameter model.InfraInternalAllocationParam,
	reason string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.ChangeInfraInternalAllocationParamMsg{
		Creator:   creator,
		Parameter: parameter,
		Reason:    reason,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// ChangeVoteParam changes VoteParam with new value.
// It composes ChangeVoteParamMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ChangeVoteParam(ctx context.Context, creator string,
	parameter model.VoteParam, reason string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.ChangeVoteParamMsg{
		Creator:   creator,
		Parameter: parameter,
		Reason:    reason,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// ChangeProposalParam changes ProposalParam with new value.
// It composes ChangeProposalParamMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ChangeProposalParam(ctx context.Context, creator string,
	parameter model.ProposalParam, reason string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.ChangeProposalParamMsg{
		Creator:   creator,
		Parameter: parameter,
		Reason:    reason,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// ChangeDeveloperParam changes DeveloperParam with new value.
// It composes ChangeDeveloperParamMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ChangeDeveloperParam(ctx context.Context, creator string,
	parameter model.DeveloperParam, reason string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.ChangeDeveloperParamMsg{
		Creator:   creator,
		Parameter: parameter,
		Reason:    reason,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// ChangeValidatorParam changes ValidatorParam with new value.
// It composes ChangeValidatorParamMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ChangeValidatorParam(ctx context.Context, creator string,
	parameter model.ValidatorParam, reason string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.ChangeValidatorParamMsg{
		Creator:   creator,
		Parameter: parameter,
		Reason:    reason,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// ChangeBandwidthParam changes BandwidthParam with new value.
// It composes ChangeBandwidthParamMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ChangeBandwidthParam(ctx context.Context, creator string,
	parameter model.BandwidthParam, reason string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.ChangeBandwidthParamMsg{
		Creator:   creator,
		Parameter: parameter,
		Reason:    reason,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// ChangeAccountParam changes AccountParam with new value.
// It composes ChangeAccountParamMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ChangeAccountParam(ctx context.Context, creator string,
	parameter model.AccountParam, reason string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.ChangeAccountParamMsg{
		Creator:   creator,
		Parameter: parameter,
		Reason:    reason,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// ChangePostParam changes PostParam with new value.
// It composes ChangePostParamMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) ChangePostParam(ctx context.Context, creator string,
	parameter model.PostParam, reason string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.ChangePostParamMsg{
		Creator:   creator,
		Parameter: parameter,
		Reason:    reason,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// DeletePostContent deletes the content of a post on blockchain, which is used
// for content censorship.
// It composes DeletePostContentMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) DeletePostContent(ctx context.Context, creator, postAuthor,
	postID, reason, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	permlink := string(string(postAuthor) + "#" + postID)
	msg := model.DeletePostContentMsg{
		Creator:  creator,
		Permlink: permlink,
		Reason:   reason,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// VoteProposal adds a vote to a certain proposal with agree/disagree.
// It composes VoteProposalMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) VoteProposal(ctx context.Context, voter, proposalID string,
	result bool, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.VoteProposalMsg{
		Voter:      voter,
		ProposalID: proposalID,
		Result:     result,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

// UpgradeProtocol upgrades the protocol.
// It composes UpgradeProtocolMsg and then broadcasts the transaction to blockchain.
func (broadcast *Broadcast) UpgradeProtocol(ctx context.Context, creator, link, reason string,
	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
	msg := model.UpgradeProtocolMsg{
		Creator: creator,
		Link:    link,
		Reason:  reason,
	}
	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
}

func (broadcast *Broadcast) retry(ctx context.Context, msg model.Msg, privKeyHex string, seq uint64, memo string, checkTxOnly bool, attempts int64, sleep time.Duration) (*model.BroadcastResponse, errors.Error) {
	res, err := broadcast.broadcastTransaction(ctx, msg, privKeyHex, seq, memo, checkTxOnly)
	if err != nil {
		if attempts--; attempts > 0 {
			if strings.Contains(err.Error(), "Tx already exists in cache") || err.CodeType() == errors.CodeTimeout {
				// if tx already exists in cache
				return res, err
			}
			if err.CodeType() == errors.CodeCheckTxFail ||
				err.CodeType() == errors.CodeDeliverTxFail {
				if err.BlockChainCode() != 155 {
					return res, err
				} else {
					// sign byte error, replace sequence number with correct one
					lo := err.BlockChainLog()
					sub := SubstringAfterStr(lo, "seq:")
					i := strings.Index(sub, "\"")
					if i != -1 {
						seqStr := sub[:i]
						correctSeq, err := strconv.ParseUint(seqStr, 10, 64)
						if err != nil {
							return res, errors.InvalidArg("invalid sequence number format")
						}
						if correctSeq == seq {
							return res, errors.InvalidSignature("invalid signature")
						}

						if broadcast.FixSequenceNumber {
							return res, errors.InvalidSequenceNumber(fmt.Sprintf("sequence number error, use %v, expect: %v", seq, correctSeq))
						}
						seq = correctSeq
					}
				}
			}
			time.Sleep(sleep)
			if broadcast.backoffRandomness {
				jitter := time.Duration(rand.Int63n(int64(sleep)))
				sleep = sleep + jitter/2
			}
			if broadcast.exponentialBackoff {
				sleep += sleep
			}

			// Add some randomness to prevent creating a Thundering Herd
			return broadcast.retry(ctx, msg, privKeyHex, seq, memo, checkTxOnly, attempts, sleep)
		}
	}
	return res, err
}

//
// internal helper functions
//
func (broadcast *Broadcast) broadcastTransaction(ctx context.Context, msg model.Msg, privKeyHex string,
	seq uint64, memo string, checkTxOnly bool) (*model.BroadcastResponse, errors.Error) {
	var res interface{}
	var err error
	finishChan := make(chan bool)

	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, memo)
	if buildErr != nil {
		return nil, buildErr
	}

	broadcastCtx, cancel := context.WithTimeout(context.Background(), broadcast.timeout)
	defer cancel()

	response := &model.BroadcastResponse{
		CommitHash: hex.EncodeToString(ttypes.Tx(txByte).Hash()),
	}

	go func() {
		res, err = broadcast.transport.BroadcastTx(txByte, checkTxOnly)
		finishChan <- true
	}()

	select {
	case <-finishChan:
		break
	case <-ctx.Done():
		return response, errors.Timeoutf("msg timeout: %v", msg).AddCause(ctx.Err())
	case <-broadcastCtx.Done():
		return response, errors.BroadcastTimeoutf("broadcast timeout: %v", msg).AddCause(ctx.Err())
	}

	if err != nil {
		return response, errors.FailedToBroadcastf("broadcast failed, err: %s", err.Error())
	}

	if checkTxOnly {
		res, ok := res.(*ctypes.ResultBroadcastTx)
		if !ok {
			return response, errors.FailedToBroadcast("error to parse the broadcast response")
		}
		code := retrieveCodeFromBlockChainCode(res.Code)
		if err == nil && code == model.InvalidSeqErrCode {
			return response, errors.InvalidSequenceNumber("invalid seq").AddBlockChainCode(res.Code).AddBlockChainLog(res.Log)
		}

		if res.Code != uint32(0) {
			return response, errors.CheckTxFail("CheckTx failed!").AddBlockChainCode(res.Code).AddBlockChainLog(res.Log)
		}
		if res.Code != uint32(0) {
			return response, errors.DeliverTxFail("DeliverTx failed!").AddBlockChainCode(res.Code).AddBlockChainLog(res.Log)
		}
	} else {
		res, ok := res.(*ctypes.ResultBroadcastTxCommit)
		if !ok {
			return response, errors.FailedToBroadcast("error to parse the broadcast response")
		}
		code := retrieveCodeFromBlockChainCode(res.CheckTx.Code)
		if err == nil && code == model.InvalidSeqErrCode {
			return response, errors.InvalidSequenceNumber("invalid seq").AddBlockChainCode(res.CheckTx.Code).AddBlockChainLog(res.CheckTx.Log)
		}

		if res.CheckTx.Code != uint32(0) {
			return response, errors.CheckTxFail("CheckTx failed!").AddBlockChainCode(res.CheckTx.Code).AddBlockChainLog(res.CheckTx.Log)
		}
		if res.DeliverTx.Code != uint32(0) {
			return response, errors.DeliverTxFail("DeliverTx failed!").AddBlockChainCode(res.DeliverTx.Code).AddBlockChainLog(res.DeliverTx.Log)
		}
		response.Height = res.Height
	}

	return response, nil
}

func retrieveCodeFromBlockChainCode(bcCode uint32) uint32 {
	return bcCode & 0xff
}

func SubstringAfterStr(value, a string) string {
	// Get substring after a string.
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:len(value)]
}
