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

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lino-network/lino-go/errors"
	"github.com/lino-network/lino-go/model"
	"github.com/lino-network/lino-go/transport"
	"github.com/lino-network/lino/param"
	linotypes "github.com/lino-network/lino/types"
	acctypes "github.com/lino-network/lino/x/account/types"
	devtypes "github.com/lino-network/lino/x/developer/types"
	infratypes "github.com/lino-network/lino/x/infra"
	posttypes "github.com/lino-network/lino/x/post/types"
	proposal "github.com/lino-network/lino/x/proposal"
	valtypes "github.com/lino-network/lino/x/validator"
	votetypes "github.com/lino-network/lino/x/vote"

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
	timeout time.Duration, exponentialBackoff bool, backoffRandomness bool) *Broadcast {
	return &Broadcast{
		transport:          transport,
		maxAttempts:        maxAttempts,
		timeout:            timeout,
		initSleepTime:      initSleepTime,
		exponentialBackoff: exponentialBackoff,
		backoffRandomness:  backoffRandomness,
		FixSequenceNumber:  true,
	}
}

//
// Account related tx
//

// Register registers a new user on blockchain.
// It composes RegisterMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) Register(ctx context.Context, referrer, registerFee, username, resetPubKeyHex,
// 	transactionPubKeyHex, appPubKeyHex, referrerPrivKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	resetPubKey, err := transport.GetPubKeyFromHex(resetPubKeyHex)
// 	if err != nil {
// 		return nil, errors.FailedToGetPubKeyFromHex("Register: failed to get Reset pub key").AddCause(err)
// 	}
// 	txPubKey, err := transport.GetPubKeyFromHex(transactionPubKeyHex)
// 	if err != nil {
// 		return nil, errors.FailedToGetPubKeyFromHex("Register: failed to get Tx pub key").AddCause(err)
// 	}
// 	appPubKey, err := transport.GetPubKeyFromHex(appPubKeyHex)
// 	if err != nil {
// 		return nil, errors.FailedToGetPubKeyFromHex("Register: failed to get App pub key").AddCause(err)
// 	}

// 	msg := model.RegisterMsg{
// 		Referrer:             referrer,
// 		RegisterFee:          registerFee,
// 		NewUser:              username,
// 		NewResetPubKey:       resetPubKey,
// 		NewTransactionPubKey: txPubKey,
// 		NewAppPubKey:         appPubKey,
// 	}
// 	return broadcast.retry(ctx, msg, referrerPrivKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeRegisterMsg return the signed register msg bytes.
func (broadcast *Broadcast) MakeRegisterMsg(ctx context.Context, referrer, registerFee, username, resetPubKeyHex,
	transactionPubKeyHex, appPubKeyHex, referrerPrivKeyHex string, seq uint64) ([]byte, errors.Error) {
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
	msg := acctypes.RegisterMsg{
		Referrer:             linotypes.AccountKey(referrer),
		RegisterFee:          registerFee,
		NewUser:              linotypes.AccountKey(username),
		NewResetPubKey:       resetPubKey,
		NewTransactionPubKey: txPubKey,
		NewAppPubKey:         appPubKey,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, referrerPrivKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

// Transfer sends a certain amount of LINO token from the sender to the receiver.
// It composes TransferMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) Transfer(ctx context.Context, sender, receiver, amount, memo,
// 	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.TransferMsg{
// 		Sender:   sender,
// 		Receiver: receiver,
// 		Amount:   amount,
// 		Memo:     memo,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeTransferMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeTransferMsg(sender, receiver, amount, memo, privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := acctypes.TransferMsg{
		Sender:   linotypes.AccountKey(sender),
		Receiver: linotypes.AccountKey(receiver),
		Amount:   amount,
		Memo:     memo,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

// func (broadcast *Broadcast) DecodeTxBytes(txbytes []byte) (*sdk.Transaction, errors.Error) {
// 	tx := &model.Transaction{}
// 	if err := broadcast.transport.Cdc.UnmarshalJSON(txbytes, tx); err != nil {
// 		return nil, errors.UnmarshaFailed("Unmarshal failed")
// 	}
// 	return tx, nil
// }

// Transfer sends a certain amount of LINO token from the sender to the receiver.
// It composes TransferMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) TransferSync(ctx context.Context, sender, receiver, amount, memo,
// 	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.TransferMsg{
// 		Sender:   sender,
// 		Receiver: receiver,
// 		Amount:   amount,
// 		Memo:     memo,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", true, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// Claim claims rewards of a certain user.
// It composes ClaimMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) Claim(ctx context.Context, username,
// 	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.ClaimMsg{
// 		Username: username,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeClaimMsg return the signed msg bytes.
// func (broadcast *Broadcast) MakeClaimMsg(username, privKeyHex string, seq uint64) ([]byte, errors.Error) {
// 	msg := acctypes.ClaimMsg{
// 		Username: username,
// 	}
// 	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
// 	if buildErr != nil {
// 		return nil, buildErr
// 	}
// 	return txByte, nil
// }

// UpdateAccount updates account related info in jsonMeta which are not
// included in AccountInfo or AccountBank.
// It composes UpdateAccountMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) UpdateAccount(ctx context.Context, username, jsonMeta,
// 	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.UpdateAccountMsg{
// 		Username: username,
// 		JSONMeta: jsonMeta,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeUpdateAccountMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeUpdateAccountMsg(username, jsonMeta,
	privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := acctypes.UpdateAccountMsg{
		Username: linotypes.AccountKey(username),
		JSONMeta: jsonMeta,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

// Recover recovers all keys of a user in case of losing or compromising.
// It composes RecoverMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) Recover(ctx context.Context, username, newResetPubKeyHex,
// 	newTransactionPubKeyHex, newAppPubKeyHex, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	resetPubKey, err := transport.GetPubKeyFromHex(newResetPubKeyHex)
// 	if err != nil {
// 		return nil, errors.FailedToGetPubKeyFromHexf("Recover: failed to get Reset pub key").AddCause(err)
// 	}
// 	txPubKey, err := transport.GetPubKeyFromHex(newTransactionPubKeyHex)
// 	if err != nil {
// 		return nil, errors.FailedToGetPubKeyFromHexf("Recover: failed to get Tx pub key").AddCause(err)
// 	}
// 	appPubKey, err := transport.GetPubKeyFromHex(newAppPubKeyHex)
// 	if err != nil {
// 		return nil, errors.FailedToGetPubKeyFromHexf("Recover: failed to get App pub key").AddCause(err)
// 	}

// 	msg := model.RecoverMsg{
// 		Username:             username,
// 		NewResetPubKey:       resetPubKey,
// 		NewTransactionPubKey: txPubKey,
// 		NewAppPubKey:         appPubKey,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeRecoverAccountMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeRecoverAccountMsg(username, newResetPubKeyHex,
	newTransactionPubKeyHex, newAppPubKeyHex, privKeyHex string, seq uint64) ([]byte, errors.Error) {
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
	msg := acctypes.RecoverMsg{
		Username:             linotypes.AccountKey(username),
		NewResetPubKey:       resetPubKey,
		NewTransactionPubKey: txPubKey,
		NewAppPubKey:         appPubKey,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

//
// Post related tx
//

// CreatePost creates a new post on blockchain.
// It composes CreatePostMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) CreatePost(ctx context.Context, author, postID, title, content,
// 	parentAuthor, parentPostID, sourceAuthor, sourcePostID, redistributionSplitRate string,
// 	links map[string]string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	var mLinks []model.IDToURLMapping
// 	if links == nil || len(links) == 0 {
// 		mLinks = nil
// 	} else {
// 		for k, v := range links {
// 			mLinks = append(mLinks, model.IDToURLMapping{k, v})
// 		}
// 	}

// 	msg := model.CreatePostMsg{
// 		Author:       author,
// 		PostID:       postID,
// 		Title:        title,
// 		Content:      content,
// 		ParentAuthor: parentAuthor,
// 		ParentPostID: parentPostID,
// 		SourceAuthor: sourceAuthor,
// 		SourcePostID: sourcePostID,
// 		Links:        mLinks,
// 		RedistributionSplitRate: redistributionSplitRate,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeCreatePostMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeCreatePostMsg(
	author, postID, title, content, createdBy string,
	preauth bool, privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := posttypes.CreatePostMsg{
		Author:    linotypes.AccountKey(author),
		PostID:    postID,
		Title:     title,
		Content:   content,
		CreatedBy: linotypes.AccountKey(createdBy),
		Preauth:   preauth,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

// CreatePost creates a new post on blockchain.
// It composes CreatePostMsg and then broadcasts the transaction to blockchain return when checkTx pass.
// func (broadcast *Broadcast) CreatePostSync(ctx context.Context, author, postID, title, content,
// 	parentAuthor, parentPostID, sourceAuthor, sourcePostID, redistributionSplitRate string,
// 	links map[string]string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	var mLinks []model.IDToURLMapping
// 	if links == nil || len(links) == 0 {
// 		mLinks = nil
// 	} else {
// 		for k, v := range links {
// 			mLinks = append(mLinks, model.IDToURLMapping{k, v})
// 		}
// 	}

// 	msg := model.CreatePostMsg{
// 		Author:       author,
// 		PostID:       postID,
// 		Title:        title,
// 		Content:      content,
// 		ParentAuthor: parentAuthor,
// 		ParentPostID: parentPostID,
// 		SourceAuthor: sourceAuthor,
// 		SourcePostID: sourcePostID,
// 		Links:        mLinks,
// 		RedistributionSplitRate: redistributionSplitRate,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", true, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// Donate adds a money donation to a post by a user.
// It composes DonateMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) Donate(ctx context.Context, username, author,
// 	amount, postID, fromApp, memo string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.DonateMsg{
// 		Username: username,
// 		Amount:   amount,
// 		Author:   author,
// 		PostID:   postID,
// 		FromApp:  fromApp,
// 		Memo:     memo,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeDonateMsg return signed msg.
func (broadcast *Broadcast) MakeDonateMsg(username, author, amount, postID, fromApp, memo string,
	privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := posttypes.DonateMsg{
		Username: linotypes.AccountKey(username),
		Amount:   amount,
		Author:   linotypes.AccountKey(author),
		PostID:   postID,
		FromApp:  linotypes.AccountKey(fromApp),
		Memo:     memo,
	}

	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

// Donate adds a money donation to a post by a user.
// It composes DonateMsg and then broadcasts the transaction to blockchain return after pass checkTx.
// func (broadcast *Broadcast) DonateSync(ctx context.Context, username, author,
// 	amount, postID, fromApp, memo string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.DonateMsg{
// 		Username: username,
// 		Amount:   amount,
// 		Author:   author,
// 		PostID:   postID,
// 		FromApp:  fromApp,
// 		Memo:     memo,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", true, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// ReportOrUpvote adds a report or upvote action to a post.
// It composes ReportOrUpvoteMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) ReportOrUpvote(ctx context.Context, username, author,
// 	postID string, isReport bool, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.ReportOrUpvoteMsg{
// 		Username: username,
// 		Author:   author,
// 		PostID:   postID,
// 		IsReport: isReport,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeReportOrUpvoteMsg return the signed msg bytes.
// func (broadcast *Broadcast) MakeReportOrUpvoteMsg(username, author,
// 	postID string, isReport bool, privKeyHex string, seq uint64) ([]byte, errors.Error) {
// 	msg := model.ReportOrUpvoteMsg{
// 		Username: username,
// 		Author:   author,
// 		PostID:   postID,
// 		IsReport: isReport,
// 	}
// 	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
// 	if buildErr != nil {
// 		return nil, buildErr
// 	}
// 	return txByte, nil
// }

// DeletePost deletes a post from the blockchain. It doesn't actually
// remove the post from the blockchain, instead it sets IsDeleted to true
// and clears all the other data.
// It composes DeletePostMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) DeletePost(ctx context.Context, author, postID,
// 	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.DeletePostMsg{
// 		Author: author,
// 		PostID: postID,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeDeleteMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeDeleteMsg(author, postID,
	privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := posttypes.DeletePostMsg{
		Author: linotypes.AccountKey(author),
		PostID: postID,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

// View increases the view count of a post by one.
// It composes ViewMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) View(ctx context.Context, username, author, postID,
// 	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.ViewMsg{
// 		Username: username,
// 		Author:   author,
// 		PostID:   postID,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// UpdatePost updates post info with new data.
// It composes UpdatePostMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) UpdatePost(ctx context.Context, author, title, postID, content string,
// 	links map[string]string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	var mLinks []model.IDToURLMapping
// 	if links == nil || len(links) == 0 {
// 		mLinks = nil
// 	} else {
// 		for k, v := range links {
// 			mLinks = append(mLinks, model.IDToURLMapping{k, v})
// 		}
// 	}

// 	msg := model.UpdatePostMsg{
// 		Author:  author,
// 		PostID:  postID,
// 		Title:   title,
// 		Content: content,
// 		Links:   mLinks,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeUpdatePostMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeUpdatePostMsg(author, title, postID, content string,
	links map[string]string, privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := posttypes.UpdatePostMsg{
		Author:  linotypes.AccountKey(author),
		PostID:  postID,
		Title:   title,
		Content: content,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

//
// Validator related tx
//

// ValidatorDeposit deposits a certain amount of LINO token for a user
// in order to become a validator. Before becoming a validator, the user
// has to be a voter.
// It composes ValidatorDepositMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) ValidatorDeposit(ctx context.Context, username, deposit,
// 	validatorPubKey, link, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	valPubKey, err := transport.GetPubKeyFromHex(validatorPubKey)
// 	if err != nil {
// 		return nil, errors.FailedToGetPubKeyFromHexf("ValidatorDeposit: failed to get Val pub key").AddCause(err)
// 	}
// 	msg := model.ValidatorDepositMsg{
// 		Username:  username,
// 		Deposit:   deposit,
// 		ValPubKey: valPubKey,
// 		Link:      link,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeValidatorDepositMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeValidatorDepositMsg(username, deposit,
	validatorPubKey, link, privKeyHex string, seq uint64) ([]byte, errors.Error) {
	valPubKey, err := transport.GetPubKeyFromHex(validatorPubKey)
	if err != nil {
		return nil, errors.FailedToGetPubKeyFromHexf("ValidatorDeposit: failed to get Val pub key").AddCause(err)
	}
	msg := valtypes.ValidatorDepositMsg{
		Username:  linotypes.AccountKey(username),
		Deposit:   deposit,
		ValPubKey: valPubKey,
		Link:      link,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

// ValidatorWithdraw withdraws part of LINO token from a validator's deposit,
// while still keep being a validator.
// It composes ValidatorDepositMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) ValidatorWithdraw(ctx context.Context, username, amount,
// 	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.ValidatorWithdrawMsg{
// 		Username: username,
// 		Amount:   amount,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeValidatorWithdrawMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeValidatorWithdrawMsg(username, amount,
	privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := valtypes.ValidatorWithdrawMsg{
		Username: linotypes.AccountKey(username),
		Amount:   amount,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

// ValidatorRevoke revokes all deposited LINO token of a validator
// so that the user will not be a validator anymore.
// It composes ValidatorRevokeMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) ValidatorRevoke(ctx context.Context, username,
// 	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.ValidatorRevokeMsg{
// 		Username: username,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeValidatorRevokeMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeValidatorRevokeMsg(username, privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := valtypes.ValidatorRevokeMsg{
		Username: linotypes.AccountKey(username),
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

//
// Vote related tx
//

// StakeIn deposits a certain amount of LINO token for a user
// in order to become a voter.
// It composes StakeInMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) StakeIn(ctx context.Context, username, deposit,
// 	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.StakeInMsg{
// 		Username: username,
// 		Deposit:  deposit,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeStakeInMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeStakeInMsg(username, deposit,
	privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := votetypes.StakeInMsg{
		Username: linotypes.AccountKey(username),
		Deposit:  deposit,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

// StakeOut withdraws part of LINO token from a voter's deposit.
// It composes StakeOutMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) StakeOut(ctx context.Context, username, amount,
// 	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.StakeOutMsg{
// 		Username: username,
// 		Amount:   amount,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeStakeOutMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeStakeOutMsg(username, amount, privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := votetypes.StakeOutMsg{
		Username: linotypes.AccountKey(username),
		Amount:   amount,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

// Delegate delegates a certain amount of LINO token of delegator to a voter, so
// the voter will have more voting power.
// It composes DelegateMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) Delegate(ctx context.Context, delegator, voter, amount,
// 	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.DelegateMsg{
// 		Delegator: delegator,
// 		Voter:     voter,
// 		Amount:    amount,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeDelegatetMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeDelegatetMsg(delegator, voter, amount,
	privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := votetypes.DelegateMsg{
		Delegator: linotypes.AccountKey(delegator),
		Voter:     linotypes.AccountKey(voter),
		Amount:    amount,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

// DelegatorWithdraw withdraws part of delegated LINO token of a delegator
// to a voter, while the delegation still exists.
// It composes DelegatorWithdrawMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) DelegatorWithdraw(ctx context.Context, delegator, voter, amount,
// 	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.DelegatorWithdrawMsg{
// 		Delegator: delegator,
// 		Voter:     voter,
// 		Amount:    amount,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeDelegatorWithdrawMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeDelegatorWithdrawMsg(delegator, voter, amount, privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := votetypes.DelegatorWithdrawMsg{
		Delegator: linotypes.AccountKey(delegator),
		Voter:     linotypes.AccountKey(voter),
		Amount:    amount,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

// ClaimInterest claims interest of a certain user.
// It composes ClaimInterestMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) ClaimInterest(ctx context.Context, username,
// 	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.ClaimInterestMsg{
// 		Username: username,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeClaimInterestMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeClaimInterestMsg(username, privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := votetypes.ClaimInterestMsg{
		Username: linotypes.AccountKey(username),
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

//
// Developer related tx
//

// DeveloperRegsiter registers a developer with a certain amount of LINO token on blockchain.
// It composes DeveloperRegisterMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) DeveloperRegister(ctx context.Context, username, deposit, website,
// 	description, appMetaData, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.DeveloperRegisterMsg{
// 		Username:    username,
// 		Deposit:     deposit,
// 		Website:     website,
// 		Description: description,
// 		AppMetaData: appMetaData,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeDeveloperRegisterMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeDeveloperRegisterMsg(username, deposit, website,
	description, appMetaData, privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := devtypes.DeveloperRegisterMsg{
		Username:    linotypes.AccountKey(username),
		Website:     website,
		Description: description,
		AppMetaData: appMetaData,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

// DeveloperUpdate updates a developer  info on blockchain.
// It composes DeveloperUpdateMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) DeveloperUpdate(ctx context.Context, username, website,
// 	description, appMetaData, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.DeveloperUpdateMsg{
// 		Username:    username,
// 		Website:     website,
// 		Description: description,
// 		AppMetaData: appMetaData,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeDeveloperUpdateMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeDeveloperUpdateMsg(username, website,
	description, appMetaData, privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := devtypes.DeveloperUpdateMsg{
		Username:    linotypes.AccountKey(username),
		Website:     website,
		Description: description,
		AppMetaData: appMetaData,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

// DeveloperRevoke reovkes all deposited LINO token of a developer
// so the user will not be a developer anymore.
// It composes DeveloperRevokeMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) DeveloperRevoke(ctx context.Context, username,
// 	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.DeveloperRevokeMsg{
// 		Username: username,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeDeveloperRevokeMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeDeveloperRevokeMsg(username, privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := devtypes.DeveloperRevokeMsg{
		Username: linotypes.AccountKey(username),
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

// GrantPermission grants a certain (e.g. App) permission to
// an authorized app with a certain period of time.
// It composes GrantPermissionMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) GrantPermission(ctx context.Context, username, authorizedApp string,
// 	validityPeriodSec int64, grantLevel model.Permission, amount string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.GrantPermissionMsg{
// 		Username:          username,
// 		AuthorizedApp:     authorizedApp,
// 		ValidityPeriodSec: validityPeriodSec,
// 		GrantLevel:        grantLevel,
// 		Amount:            amount,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeGrantPermissionMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeGrantPermissionMsg(username, authorizedApp string,
	validityPeriodSec int64, grantLevel linotypes.Permission, amount string, privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := devtypes.GrantPermissionMsg{
		Username:          linotypes.AccountKey(username),
		AuthorizedApp:     linotypes.AccountKey(authorizedApp),
		ValidityPeriodSec: validityPeriodSec,
		GrantLevel:        grantLevel,
		Amount:            amount,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

// GrantAppAndPreAuthPermission grants both app and preauth permission to
// an authorized app with a certain period of time.
// It composes GrantPermissionMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) GrantAppAndPreAuthPermission(ctx context.Context, username, authorizedApp string,
// 	validityPeriodSec int64, amount string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.GrantPermissionMsg{
// 		Username:          username,
// 		AuthorizedApp:     authorizedApp,
// 		ValidityPeriodSec: validityPeriodSec,
// 		GrantLevel:        model.AppAndPreAuthorizationPermission,
// 		Amount:            amount,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeGrantAppAndPreAuthPermissionMsg return the signed msg bytes.
// func (broadcast *Broadcast) MakeGrantAppAndPreAuthPermissionMsg(username, authorizedApp string,
// 	validityPeriodSec int64, amount string, privKeyHex string, seq uint64) ([]byte, errors.Error) {
// 	msg := devtypes.GrantPermissionMsg{
// 		Username:          linotypes.AccountKey(username),
// 		AuthorizedApp:     linotypes.AccountKey(authorizedApp),
// 		ValidityPeriodSec: validityPeriodSec,
// 		GrantLevel:        linotypes.AppAndPreAuthorizationPermission,
// 		Amount:            amount,
// 	}
// 	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
// 	if buildErr != nil {
// 		return nil, buildErr
// 	}
// 	return txByte, nil
// }

// PreAuthorizationPermission grants a PreAuthorization permission to
// an authorzied app with a certain period of time.
// It composes GrantPermissionMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) PreAuthorizationPermission(ctx context.Context, username, authorizedApp string,
// 	validityPeriodSec int64, amount string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.GrantPermissionMsg{
// 		Username:          username,
// 		AuthorizedApp:     authorizedApp,
// 		ValidityPeriodSec: validityPeriodSec,
// 		GrantLevel:        model.PreAuthorizationPermission,
// 		Amount:            amount,
// 	}

// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakePreAuthorizationPermissionMsg return the signed msg bytes.
// func (broadcast *Broadcast) MakePreAuthorizationPermissionMsg(username, authorizedApp string,
// 	validityPeriodSec int64, amount string, privKeyHex string, seq uint64) ([]byte, errors.Error) {
// 	msg := devtypes.GrantPermissionMsg{
// 		Username:          linotypes.AccountKey(username),
// 		AuthorizedApp:     linotypes.AccountKey(authorizedApp),
// 		ValidityPeriodSec: validityPeriodSec,
// 		GrantLevel:        model.PreAuthorizationPermission,
// 		Amount:            amount,
// 	}
// 	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
// 	if buildErr != nil {
// 		return nil, buildErr
// 	}
// 	return txByte, nil
// }

// RevokePermission revokes the permission given previously to a app.
// It composes RevokePermissionMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) RevokePermission(ctx context.Context, username, pubKeyHex string,
// 	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	pubKey, err := transport.GetPubKeyFromHex(pubKeyHex)
// 	if err != nil {
// 		return nil, errors.FailedToGetPubKeyFromHex("Register: failed to get pub key").AddCause(err)
// 	}

// 	msg := model.RevokePermissionMsg{
// 		Username: username,
// 		PubKey:   pubKey,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeRevokePermissionPermissionMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeRevokePermissionPermissionMsg(username, revokeFrom string, permission linotypes.Permission,
	privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := devtypes.RevokePermissionMsg{
		Username:   linotypes.AccountKey(username),
		RevokeFrom: linotypes.AccountKey(revokeFrom),
		Permission: permission,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

//
// infra related tx
//

// ProviderReport reports infra usage of a infra provider in order to get infra inflation.
// It composes ProviderReportMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) ProviderReport(ctx context.Context, username string, usage int64,
// 	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.ProviderReportMsg{
// 		Username: username,
// 		Usage:    usage,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeProviderReportMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeProviderReportMsg(username string, usage int64,
	privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := infratypes.ProviderReportMsg{
		Username: linotypes.AccountKey(username),
		Usage:    usage,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

//
// proposal related tx
//

// ChangeGlobalAllocationParam changes GlobalAllocationParam with new value.
// It composes ChangeGlobalAllocationParamMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) ChangeGlobalAllocationParam(ctx context.Context, creator string,
// 	parameter model.GlobalAllocationParam, reason string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.ChangeGlobalAllocationParamMsg{
// 		Creator:   creator,
// 		Parameter: parameter,
// 		Reason:    reason,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeChangeGlobalAllocationParamMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeChangeGlobalAllocationParamMsg(creator string,
	parameter param.GlobalAllocationParam, reason string, privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := proposal.ChangeGlobalAllocationParamMsg{
		Creator:   linotypes.AccountKey(creator),
		Parameter: parameter,
		Reason:    reason,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

// ChangeInfraInternalAllocationParam changes InfraInternalAllocationParam with new value.
// It composes ChangeInfraInternalAllocationParamMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) ChangeInfraInternalAllocationParam(ctx context.Context, creator string,
// 	parameter model.InfraInternalAllocationParam,
// 	reason string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.ChangeInfraInternalAllocationParamMsg{
// 		Creator:   creator,
// 		Parameter: parameter,
// 		Reason:    reason,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeChangeInfraInternalAllocationParamMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeChangeInfraInternalAllocationParamMsg(creator string,
	parameter param.InfraInternalAllocationParam,
	reason string, privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := proposal.ChangeInfraInternalAllocationParamMsg{
		Creator:   linotypes.AccountKey(creator),
		Parameter: parameter,
		Reason:    reason,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

// ChangeVoteParam changes VoteParam with new value.
// It composes ChangeVoteParamMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) ChangeVoteParam(ctx context.Context, creator string,
// 	parameter model.VoteParam, reason string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.ChangeVoteParamMsg{
// 		Creator:   creator,
// 		Parameter: parameter,
// 		Reason:    reason,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeChangeVoteParamMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeChangeVoteParamMsg(creator string,
	parameter param.VoteParam, reason string, privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := proposal.ChangeVoteParamMsg{
		Creator:   linotypes.AccountKey(creator),
		Parameter: parameter,
		Reason:    reason,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

// ChangeProposalParam changes ProposalParam with new value.
// It composes ChangeProposalParamMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) ChangeProposalParam(ctx context.Context, creator string,
// 	parameter model.ProposalParam, reason string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.ChangeProposalParamMsg{
// 		Creator:   creator,
// 		Parameter: parameter,
// 		Reason:    reason,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeChangeProposalParamMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeChangeProposalParamMsg(creator string,
	parameter param.ProposalParam, reason string, privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := proposal.ChangeProposalParamMsg{
		Creator:   linotypes.AccountKey(creator),
		Parameter: parameter,
		Reason:    reason,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

// ChangeDeveloperParam changes DeveloperParam with new value.
// It composes ChangeDeveloperParamMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) ChangeDeveloperParam(ctx context.Context, creator string,
// 	parameter model.DeveloperParam, reason string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.ChangeDeveloperParamMsg{
// 		Creator:   creator,
// 		Parameter: parameter,
// 		Reason:    reason,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeChangeDeveloperParamMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeChangeDeveloperParamMsg(creator string,
	parameter param.DeveloperParam, reason string, privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := proposal.ChangeDeveloperParamMsg{
		Creator:   linotypes.AccountKey(creator),
		Parameter: parameter,
		Reason:    reason,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

// ChangeValidatorParam changes ValidatorParam with new value.
// It composes ChangeValidatorParamMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) ChangeValidatorParam(ctx context.Context, creator string,
// 	parameter model.ValidatorParam, reason string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.ChangeValidatorParamMsg{
// 		Creator:   creator,
// 		Parameter: parameter,
// 		Reason:    reason,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeChangeValidatorParamMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeChangeValidatorParamMsg(creator string,
	parameter param.ValidatorParam, reason string, privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := proposal.ChangeValidatorParamMsg{
		Creator:   linotypes.AccountKey(creator),
		Parameter: parameter,
		Reason:    reason,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

// ChangeBandwidthParam changes BandwidthParam with new value.
// It composes ChangeBandwidthParamMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) ChangeBandwidthParam(ctx context.Context, creator string,
// 	parameter model.BandwidthParam, reason string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.ChangeBandwidthParamMsg{
// 		Creator:   creator,
// 		Parameter: parameter,
// 		Reason:    reason,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeChangeBandwidthParamMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeChangeBandwidthParamMsg(creator string,
	parameter param.BandwidthParam, reason string, privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := proposal.ChangeBandwidthParamMsg{
		Creator:   linotypes.AccountKey(creator),
		Parameter: parameter,
		Reason:    reason,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

// ChangeAccountParam changes AccountParam with new value.
// It composes ChangeAccountParamMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) ChangeAccountParam(ctx context.Context, creator string,
// 	parameter model.AccountParam, reason string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.ChangeAccountParamMsg{
// 		Creator:   creator,
// 		Parameter: parameter,
// 		Reason:    reason,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeChangeAccountParamMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeChangeAccountParamMsg(creator string,
	parameter param.AccountParam, reason string, privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := proposal.ChangeAccountParamMsg{
		Creator:   linotypes.AccountKey(creator),
		Parameter: parameter,
		Reason:    reason,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

// ChangePostParam changes PostParam with new value.
// It composes ChangePostParamMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) ChangePostParam(ctx context.Context, creator string,
// 	parameter model.PostParam, reason string, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.ChangePostParamMsg{
// 		Creator:   creator,
// 		Parameter: parameter,
// 		Reason:    reason,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeChangePostParamMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeChangePostParamMsg(creator string,
	parameter param.PostParam, reason string, privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := proposal.ChangePostParamMsg{
		Creator:   linotypes.AccountKey(creator),
		Parameter: parameter,
		Reason:    reason,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

// MakeDeletePostContentMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeDeletePostContentMsg(creator, postAuthor,
	postID, reason, privKeyHex string, seq uint64) ([]byte, errors.Error) {
	permlink := string(string(postAuthor) + "#" + postID)
	msg := proposal.DeletePostContentMsg{
		Creator:  linotypes.AccountKey(creator),
		Permlink: linotypes.Permlink(permlink),
		Reason:   reason,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

// VoteProposal adds a vote to a certain proposal with agree/disagree.
// It composes VoteProposalMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) VoteProposal(ctx context.Context, voter, proposalID string,
// 	result bool, privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.VoteProposalMsg{
// 		Voter:      voter,
// 		ProposalID: proposalID,
// 		Result:     result,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeVoteProposalMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeVoteProposalMsg(voter, proposalID string,
	result bool, privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := proposal.VoteProposalMsg{
		Voter:      linotypes.AccountKey(voter),
		ProposalID: linotypes.ProposalKey(proposalID),
		Result:     result,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

// UpgradeProtocol upgrades the protocol.
// It composes UpgradeProtocolMsg and then broadcasts the transaction to blockchain.
// func (broadcast *Broadcast) UpgradeProtocol(ctx context.Context, creator, link, reason string,
// 	privKeyHex string, seq uint64) (*model.BroadcastResponse, errors.Error) {
// 	msg := model.UpgradeProtocolMsg{
// 		Creator: creator,
// 		Link:    link,
// 		Reason:  reason,
// 	}
// 	return broadcast.retry(ctx, msg, privKeyHex, seq, "", false, broadcast.maxAttempts, broadcast.initSleepTime)
// }

// MakeUpgradeProtocolMsg return the signed msg bytes.
func (broadcast *Broadcast) MakeUpgradeProtocolMsg(creator, link, reason string,
	privKeyHex string, seq uint64) ([]byte, errors.Error) {
	msg := proposal.UpgradeProtocolMsg{
		Creator: linotypes.AccountKey(creator),
		Link:    link,
		Reason:  reason,
	}
	txByte, buildErr := broadcast.transport.SignAndBuild(msg, privKeyHex, seq, "")
	if buildErr != nil {
		return nil, buildErr
	}
	return txByte, nil
}

func (broadcast *Broadcast) retry(ctx context.Context, msg sdk.Msg, privKeyHex string, seq uint64, memo string, checkTxOnly bool, attempts int64, sleep time.Duration) (*model.BroadcastResponse, errors.Error) {
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

// CalcTxMsgHash return hash bytes
func CalcTxMsgHash(msg []byte) ([]byte, errors.Error) {
	if msg == nil {
		return nil, errors.InvalidArg("CalcTxMsgHash: empty msg bytes")
	}
	return ttypes.Tx(msg).Hash(), nil
}

// CalcTxMsgHashHexString return hex encoded hash string
func CalcTxMsgHashHexString(msg []byte) (string, errors.Error) {
	hash, err := CalcTxMsgHash(msg)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash), nil
}

// ExtractSeqNumberFromErrLog extracts correct sequence number from
func ExtractSeqNumberFromErrLog(log string) *uint64 {
	sub := SubstringAfterStr(log, "seq:")
	i := strings.Index(sub, "\"")
	if i != -1 {
		seqStr := sub[:i]
		correctSeq, err := strconv.ParseUint(seqStr, 10, 64)
		if err != nil {
			return nil
		}
		return &correctSeq
	}
	return nil
}

// BroadcastRawMsgBytesSync broadcast message to CheckTx.
func (broadcast *Broadcast) BroadcastRawMsgBytesSync(ctx context.Context, txBytes []byte, seq uint64) errors.Error {
	var res interface{}
	var err error
	finishCh := make(chan struct{})

	broadcastCtx, cancel := context.WithTimeout(ctx, broadcast.timeout)
	defer cancel()

	go func() {
		defer func() {
			finishCh <- struct{}{}
		}()
		res, err = broadcast.transport.BroadcastTx(txBytes, true)
	}()

	select {
	case <-finishCh:
		break
	case <-ctx.Done():
		return errors.Timeoutf("msg timeout").AddCause(ctx.Err())
	case <-broadcastCtx.Done():
		return errors.BroadcastTimeoutf("broadcast timeout").AddCause(ctx.Err())
	}

	if err != nil {
		return errors.FailedToBroadcastf("broadcast failed, err: %s", err.Error())
	}

	bres, ok := res.(*ctypes.ResultBroadcastTx)
	if !ok {
		return errors.FailedToBroadcast("error to parse the broadcast response")
	}
	code := retrieveCodeFromBlockChainCode(bres.Code)

	// special handling of invalid sequence number
	if code == uint32(errors.CodeUnverifiedBytes) {
		correctSeq := ExtractSeqNumberFromErrLog(bres.Log)
		if correctSeq != nil && seq != *correctSeq {
			return errors.InvalidSequenceNumber("invalid seq").
				AddBlockChainCode(bres.Code).AddBlockChainLog(bres.Log)
		}
	}

	if bres.Code != uint32(0) {
		return errors.CheckTxFail("CheckTx failed!").
			AddBlockChainCode(bres.Code).AddBlockChainLog(bres.Log)
	}
	return nil
}

//
// internal helper functions
//
func (broadcast *Broadcast) broadcastTransaction(ctx context.Context, msg sdk.Msg, privKeyHex string,
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
