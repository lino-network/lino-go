// Package api initiates a go library API which can be
// used to query data from blockchain and
// broadcast transactions to blockchain.
package api

import (
	"context"
	"encoding/hex"
	goerrors "errors"
	"strings"
	"time"

	"github.com/lino-network/lino-go/broadcast"
	"github.com/lino-network/lino-go/errors"
	"github.com/lino-network/lino-go/model"
	"github.com/lino-network/lino-go/query"
	"github.com/lino-network/lino-go/transport"
	"github.com/lino-network/lino-go/util"
	// "github.com/lino-network/lino/param"
	linotypes "github.com/lino-network/lino/types"
	accmodel "github.com/lino-network/lino/x/account/model"
	"github.com/spf13/viper"
)

// internal errors, not exported.
var (
	errTxWatchTimeout   = goerrors.New("errTxWatchTimeout")
	errSeqChanged       = goerrors.New("errSeqChanged")
	errSeqTxQueryFailed = goerrors.New("errSeqTxQueryFailed")
)

// API is a wrapper of both querying data from blockchain
// and broadcast transactions to blockchain.
type API struct {
	*query.Query
	*broadcast.Broadcast
	checkTxConfirmInterval time.Duration
	timeout                time.Duration
}

// Options is a wrapper of init parameters
type Options struct {
	ChainID                string        `json:"chain_id"`
	NodeURL                string        `json:"node_url"`
	MaxAttempts            int64         `json:"max_attempts"`
	MaxFeeInCoin           int64         `json:"max_fee_in_coin"`
	InitSleepTime          time.Duration `json:"init_sleep_time"`
	Timeout                time.Duration `json:"timeout"`
	ExponentialBackoff     bool          `json:"exponential_back_off"`
	BackoffRandomness      bool          `json:"backoff_randomness"`
	CheckTxConfirmInterval time.Duration `json:"check_tx_confirm_interval"`
}

func (opt *Options) init() {
	if opt.MaxAttempts == 0 {
		opt.MaxAttempts = 3
	}
	if opt.InitSleepTime == 0 {
		opt.InitSleepTime = time.Second * 3
	}
	if opt.Timeout == 0 {
		opt.Timeout = time.Second * 10
	}
	if opt.MaxFeeInCoin == 0 {
		opt.MaxFeeInCoin = linotypes.Decimals
	}
	if opt.CheckTxConfirmInterval == 0 {
		opt.CheckTxConfirmInterval = time.Second
	}
}

// NewLinoAPIFromConfig initiates an instance of API using
// configs from ~/.lino-go/config.json
func NewLinoAPIFromConfig() *API {
	v := viper.New()
	viper.SetConfigType("json")
	v.SetConfigName("config")
	v.AddConfigPath("$HOME/.lino-go/")
	v.AutomaticEnv()
	v.ReadInConfig()

	nodeURL := v.GetString("node_RL")
	chainID := v.GetString("chain_id")
	maxAttempts := v.GetInt64("max_attempts")
	initSleepTime := v.GetInt64("init_sleep_time")
	exponentialBackoff := v.GetBool("exponential_back_off")
	backoffRandomness := v.GetBool("backoff_randomness")
	return NewLinoAPIFromArgs(&Options{
		ChainID:            chainID,
		NodeURL:            nodeURL,
		MaxAttempts:        maxAttempts,
		InitSleepTime:      time.Duration(initSleepTime) * time.Second,
		ExponentialBackoff: exponentialBackoff,
		BackoffRandomness:  backoffRandomness,
	})
}

// NewLinoAPIFromArgs initiates an instance of API using
// chainID and nodeUrl that are passed in.
func NewLinoAPIFromArgs(opt *Options) *API {
	opt.init()
	linotypes.ConfigAndSealCosmosSDKAddress()
	transport := transport.NewTransportFromArgs(opt.ChainID, opt.NodeURL, opt.MaxFeeInCoin)
	return &API{
		Query:                  query.NewQuery(transport),
		Broadcast:              broadcast.NewBroadcast(transport, opt.MaxAttempts, opt.InitSleepTime, opt.Timeout, opt.ExponentialBackoff, opt.BackoffRandomness),
		checkTxConfirmInterval: opt.CheckTxConfirmInterval,
		timeout:                opt.Timeout,
	}
}

// MsgBuilderFunc is usually a closure that return messages bytes for a specific sequence.
type MsgBuilderFunc func(seqs []uint64) ([]byte, errors.Error)

// Register registers a new user on blockchain.
// It composes RegisterMsg and then broadcasts the transaction to blockchain.
// func (api *API) Register(ctx context.Context, referrer, registerFee, username, resetPubKeyHex,
// 	transactionPubKeyHex, appPubKeyHex, referrerPrivKeyHex string) (*model.BroadcastResponse, errors.Error) {
// 	resp, _, err := api.GuaranteeBroadcast(
// 		ctx, util.GetSignerList(referrer),
// 		func(seqs []uint64) ([]byte, errors.Error) {
// 			return api.MakeRegisterMsg(
// 				ctx, referrer, registerFee, username, resetPubKeyHex,
// 				transactionPubKeyHex, appPubKeyHex, referrerPrivKeyHex, seqs[0])
// 		})
// 	return resp, err
// }

// RegisterV2 registers a new user on blockchain.
// It composes RegisterMsg and then broadcasts the transaction to blockchain.
func (api *API) RegisterV2(ctx context.Context, referrer linotypes.AccOrAddr, registerFee, username, newTxAddr, txPubKeyHex,
	signingPubKeyHex, referrerPrivKeyHex, txPrivKeyHex string) (*model.BroadcastResponse, errors.Error) {
	addr, e := hex.DecodeString(newTxAddr)
	if e != nil {
		return nil, errors.InvalidArg("Invalid Transaction Key Address")
	}
	resp, _, err := api.GuaranteeBroadcast(
		ctx, []linotypes.AccOrAddr{referrer, linotypes.NewAccOrAddrFromAddr(addr)},
		func(seqs []uint64) ([]byte, errors.Error) {
			if len(seqs) < 2 {
				return nil, errors.SequenceNumberNotEnoughf("sequence number is not enough. got %d, expect %d", len(seqs), 2)
			}
			return api.MakeRegisterV2Msg(
				ctx, referrer, registerFee, username, txPubKeyHex,
				signingPubKeyHex, referrerPrivKeyHex, txPrivKeyHex, seqs[0], seqs[1])
		})
	return resp, err
}

// Transfer sends a certain amount of LINO token from the sender to the receiver.
// It composes TransferMsg and then broadcasts the transaction to blockchain.
func (api *API) Transfer(
	ctx context.Context, sender, receiver, amount, memo, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(sender), func(seqs []uint64) ([]byte, errors.Error) {
		return api.MakeTransferMsg(sender, receiver, amount, memo, privKeyHex, seqs[0])
	})
	return resp, err
}

// TransferV2 sends a certain amount of LINO token from the sender to the receiver.
// sender and receiver can be address or username
// It composes TransferMsg and then broadcasts the transaction to blockchain.
func (api *API) TransferV2(
	ctx context.Context, sender, receiver linotypes.AccOrAddr, amount, memo, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
	resp, _, err := api.GuaranteeBroadcast(ctx, []linotypes.AccOrAddr{sender}, func(seqs []uint64) ([]byte, errors.Error) {
		return api.MakeTransferV2Msg(sender, receiver, amount, memo, privKeyHex, seqs[0])
	})
	return resp, err
}

func (api *API) UpdateAccountMeta(
	ctx context.Context, username string, meta string, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(username), func(seqs []uint64) ([]byte, errors.Error) {
		return api.MakeUpdateAccountMsg(username, meta, privKeyHex, seqs[0])
	})
	return resp, err
}

// Claim claims rewards of a certain user.
// It composes ClaimMsg and then broadcasts the transaction to blockchain.
// func (api *API) Claim(ctx context.Context, username, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
// 	resp, _, err := api.GuaranteeBroadcast(ctx, username, func(seqs []uint64) ([]byte, errors.Error) {
// 		return api.MakeClaimMsg(username, privKeyHex, seqs[0])
// 	})
// 	return resp, err
// }

// UpdateAccount updates account related info in jsonMeta which are not
// included in AccountInfo or AccountBank.
// It composes UpdateAccountMsg and then broadcasts the transaction to blockchain.
func (api *API) UpdateAccount(
	ctx context.Context, username, jsonMeta, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(username), func(seqs []uint64) ([]byte, errors.Error) {
		return api.MakeUpdateAccountMsg(username, jsonMeta, privKeyHex, seqs[0])
	})
	return resp, err
}

// UpdateAccount updates account related info in jsonMeta which are not
// included in AccountInfo or AccountBank.
// It composes UpdateAccountMsg and then broadcasts the transaction to blockchain.
func (api *API) Recover(
	ctx context.Context, username, newTxAddr, newTxPubKeyHex, newSigningPubKeyHex,
	privKeyHex string, newTxPrivKeyHex string) (*model.BroadcastResponse, errors.Error) {
	addr, e := hex.DecodeString(newTxAddr)
	if e != nil {
		return nil, errors.InvalidArg("Invalid Transaction Key Address")
	}
	resp, _, err := api.GuaranteeBroadcast(
		ctx, append(util.GetSignerList(username), linotypes.NewAccOrAddrFromAddr(addr)), func(seqs []uint64) ([]byte, errors.Error) {
			if len(seqs) < 2 {
				return nil, errors.SequenceNumberNotEnoughf("sequence number is not enough. got %d, expect %d", len(seqs), 2)
			}
			return api.MakeRecoverAccountMsg(
				username, newTxPubKeyHex, newSigningPubKeyHex, privKeyHex, newTxPrivKeyHex, seqs[0], seqs[1])
		})
	return resp, err
}

// CreatePost creates a new post on blockchain.
// It composes CreatePostMsg and then broadcasts the transaction to blockchain.
func (api *API) CreatePost(
	ctx context.Context, author, postID, title, content, createdBy string, preauth bool,
	privKeyHex string) (resp *model.BroadcastResponse, err errors.Error) {
	if preauth {
		resp, _, err = api.GuaranteeBroadcast(ctx, util.GetSignerList(author), func(seqs []uint64) ([]byte, errors.Error) {
			return api.MakeCreatePostMsg(author, postID, title, content, createdBy, preauth, privKeyHex, seqs[0])
		})
	} else {
		resp, _, err = api.GuaranteeBroadcast(ctx, util.GetSignerList(createdBy), func(seqs []uint64) ([]byte, errors.Error) {
			return api.MakeCreatePostMsg(author, postID, title, content, createdBy, preauth, privKeyHex, seqs[0])
		})
	}
	return resp, err
}

// Donate adds a money donation to a post by a user.
// It composes DonateMsg and then broadcasts the transaction to blockchain.
func (api *API) Donate(ctx context.Context, username, author,
	amount, postID, fromApp, memo string, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(username), func(seqs []uint64) ([]byte, errors.Error) {
		return api.MakeDonateMsg(username, author, amount, postID, fromApp, memo, privKeyHex, seqs[0])
	})
	return resp, err
}

// DeletePost deletes a post from the blockchain. It doesn't actually
// remove the post from the blockchain, instead it sets IsDeleted to true
// and clears all the other data.
// It composes DeletePostMsg and then broadcasts the transaction to blockchain.
func (api *API) DeletePost(ctx context.Context, author,
	postID string, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(author), func(seqs []uint64) ([]byte, errors.Error) {
		return api.MakeDeleteMsg(author, postID, privKeyHex, seqs[0])
	})
	return resp, err
}

// UpdatePost updates post info with new data.
// It composes UpdatePostMsg and then broadcasts the transaction to blockchain.
func (api *API) UpdatePost(
	ctx context.Context, author, title, postID, content string, links map[string]string,
	privKeyHex string) (*model.BroadcastResponse, errors.Error) {
	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(author), func(seqs []uint64) ([]byte, errors.Error) {
		return api.MakeUpdatePostMsg(author, title, postID, content, links, privKeyHex, seqs[0])
	})
	return resp, err
}

// ValidatorRegister registers validator
// It composes ValidatorDepositMsg and then broadcasts the transaction to blockchain.
func (api *API) ValidatorRegister(ctx context.Context, username,
	validatorPubKey, link, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(username), func(seqs []uint64) ([]byte, errors.Error) {
		return api.MakeValidatorRegisterMsg(username, validatorPubKey, link, privKeyHex, seqs[0])
	})
	return resp, err
}

// ValidatorUpdate registers validator
// It composes ValidatorUpdate and then broadcasts the transaction to blockchain.
func (api *API) ValidatorUpdate(ctx context.Context, username, link, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(username), func(seqs []uint64) ([]byte, errors.Error) {
		return api.MakeValidatorUpdateMsg(username, link, privKeyHex, seqs[0])
	})
	return resp, err
}

// ValidatorRevoke revokes all deposited LINO token of a validator
// so that the user will not be a validator anymore.
// It composes ValidatorRevokeMsg and then broadcasts the transaction to blockchain.
func (api *API) ValidatorRevoke(
	ctx context.Context, username, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(username), func(seqs []uint64) ([]byte, errors.Error) {
		return api.MakeValidatorRevokeMsg(username, privKeyHex, seqs[0])
	})
	return resp, err
}

func (api *API) VoteValidator(
	ctx context.Context, username string, validators []string,
	privKeyHex string) (*model.BroadcastResponse, errors.Error) {
	resp, _, err := api.GuaranteeBroadcast(
		ctx, util.GetSignerList(username), func(seqs []uint64) ([]byte, errors.Error) {
			return api.MakeVoteValidatorMsg(username, validators, privKeyHex, seqs[0])
		})
	return resp, err
}

// StakeIn deposits a certain amount of LINO token for a user
// in order to become a voter.
// It composes StakeInMsg and then broadcasts the transaction to blockchain.
func (api *API) StakeIn(
	ctx context.Context, username, deposit, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(username), func(seqs []uint64) ([]byte, errors.Error) {
		return api.MakeStakeInMsg(username, deposit, privKeyHex, seqs[0])
	})
	return resp, err
}

// StakeInFor deposits a certain amount of LINO token from sender to receiver
// in order to become a voter.
// It composes StakeInForMsg and then broadcasts the transaction to blockchain.
func (api *API) StakeInFor(
	ctx context.Context, sender, receiver, deposit, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(sender), func(seqs []uint64) ([]byte, errors.Error) {
		return api.MakeStakeInForMsg(sender, receiver, deposit, privKeyHex, seqs[0])
	})
	return resp, err
}

// StakeOut withdraws part of LINO token from a voter's deposit.
// It composes StakeOutMsg and then broadcasts the transaction to blockchain.
func (api *API) StakeOut(
	ctx context.Context, username, amount, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(username), func(seqs []uint64) ([]byte, errors.Error) {
		return api.MakeStakeOutMsg(username, amount, privKeyHex, seqs[0])
	})
	return resp, err
}

// ClaimInterest claims interest of a certain user.
// It composes ClaimInterestMsg and then broadcasts the transaction to blockchain.
func (api *API) ClaimInterest(
	ctx context.Context, username, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(username), func(seqs []uint64) ([]byte, errors.Error) {
		return api.MakeClaimInterestMsg(username, privKeyHex, seqs[0])
	})
	return resp, err
}

// DeveloperRegsiter registers a developer with a certain amount of LINO token on blockchain.
// It composes DeveloperRegisterMsg and then broadcasts the transaction to blockchain.
func (api *API) DeveloperRegister(
	ctx context.Context, username, website, description,
	appMetaData, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(username), func(seqs []uint64) ([]byte, errors.Error) {
		return api.MakeDeveloperRegisterMsg(username, website, description, appMetaData, privKeyHex, seqs[0])
	})
	return resp, err
}

// DeveloperUpdate updates a developer  info on blockchain.
// It composes DeveloperUpdateMsg and then broadcasts the transaction to blockchain.
func (api *API) DeveloperUpdate(
	ctx context.Context, username, website, description, appMetaData,
	privKeyHex string) (*model.BroadcastResponse, errors.Error) {
	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(username), func(seqs []uint64) ([]byte, errors.Error) {
		return api.MakeDeveloperUpdateMsg(username, website, description, appMetaData, privKeyHex, seqs[0])
	})
	return resp, err
}

// DeveloperRevoke updates a developer  info on blockchain.
// It composes DeveloperRevokeMsg and then broadcasts the transaction to blockchain.
func (api *API) DeveloperRevoke(
	ctx context.Context, username, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(username), func(seqs []uint64) ([]byte, errors.Error) {
		return api.MakeDeveloperRevokeMsg(username, privKeyHex, seqs[0])
	})
	return resp, err
}

// GrantPermission grants a certain (e.g. App) permission to
// an authorized app with a certain period of time.
// It composes GrantPermissionMsg and then broadcasts the transaction to blockchain.
// func (api *API) GrantPermission(
// 	ctx context.Context, username, authorizedApp string,
// 	validityPeriodSec int64, grantLevel linotypes.Permission,
// 	amount string, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
// 	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(username), func(seqs []uint64) ([]byte, errors.Error) {
// 		return api.MakeGrantPermissionMsg(
// 			username, authorizedApp, validityPeriodSec, grantLevel, amount, privKeyHex, seqs[0])
// 	})
// 	return resp, err
// }

// GrantAppAndPreAuthPermission grants both app and preauth permission to
// an authorized app with a certain period of time.
// It composes GrantPermissionMsg and then broadcasts the transaction to blockchain.
// func (api *API) GrantAppAndPreAuthPermission(ctx context.Context, username, authorizedApp string,
// 	validityPeriodSec int64, amount string, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
// 	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(username), func(seqs []uint64) ([]byte, errors.Error) {
// 		return api.MakeGrantPermissionMsg(
// 			username, authorizedApp, validityPeriodSec,
// 			linotypes.AppAndPreAuthorizationPermission, amount, privKeyHex, seqs[0])
// 	})
// 	return resp, err
// }

// RevokePermission revokes the permission given previously to a app.
// It composes RevokePermissionMsg and then broadcasts the transaction to blockchain.
// func (api *API) RevokePermission(
// 	ctx context.Context, username, revokeFrom string, permission linotypes.Permission,
// 	privKeyHex string) (*model.BroadcastResponse, errors.Error) {
// 	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(username), func(seqs []uint64) ([]byte, errors.Error) {
// 		return api.MakeRevokePermissionPermissionMsg(username, revokeFrom, permission, privKeyHex, seqs[0])
// 	})
// 	return resp, err
// }

// IDAIssue issues IDA on the blockchain.
func (api *API) IDAIssue(
	ctx context.Context, username string, IDAPrice int64, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(username), func(seqs []uint64) ([]byte, errors.Error) {
		return api.MakeIDAIssueMsg(username, IDAPrice, privKeyHex, seqs[0])
	})
	return resp, err
}

// IDAMint generates new IDA on the blockchain.
func (api *API) IDAMint(
	ctx context.Context, username, amount string, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(username), func(seqs []uint64) ([]byte, errors.Error) {
		return api.MakeIDAMintMsg(username, amount, privKeyHex, seqs[0])
	})
	return resp, err
}

// IDATransfer moves IDA between accounts.
func (api *API) IDATransfer(
	ctx context.Context, app, amount, from, to, signer, memo string,
	privKeyHex string) (*model.BroadcastResponse, errors.Error) {
	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(signer), func(seqs []uint64) ([]byte, errors.Error) {
		return api.MakeIDATransferMsg(app, amount, from, to, signer, memo, privKeyHex, seqs[0])
	})
	return resp, err
}

// IDADonate moves IDA between accounts.
func (api *API) IDADonate(
	ctx context.Context, username, author, app, amount, postID, signer, memo string,
	privKeyHex string) (*model.BroadcastResponse, errors.Error) {
	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(signer), func(seqs []uint64) ([]byte, errors.Error) {
		return api.MakeIDADonateMsg(username, author, app, amount, postID, signer, memo, privKeyHex, seqs[0])
	})
	return resp, err
}

// IDAAuthorize can set status of user's IDA account.
func (api *API) IDAAuthorize(
	ctx context.Context, username, app string, activate bool,
	privKeyHex string) (*model.BroadcastResponse, errors.Error) {
	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(username), func(seqs []uint64) ([]byte, errors.Error) {
		return api.MakeIDAAuthorizeMsg(username, app, activate, privKeyHex, seqs[0])
	})
	return resp, err
}

// UpdateAffiliated can set affiliate account for app.
func (api *API) UpdateAffiliated(
	ctx context.Context, username, app string, activate bool,
	privKeyHex string) (*model.BroadcastResponse, errors.Error) {
	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(app), func(seqs []uint64) ([]byte, errors.Error) {
		return api.MakeUpdateAffiliatedMsg(username, app, activate, privKeyHex, seqs[0])
	})
	return resp, err
}

// ProviderReport reports infra usage of a infra provider in order to get infra inflation.
// It composes ProviderReportMsg and then broadcasts the transaction to blockchain.
// func (api *API) ProviderReport(ctx context.Context, username string, usage int64, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
// 	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(username), func(seqs []uint64) ([]byte, errors.Error) {
// 		return api.MakeProviderReportMsg(username, usage, privKeyHex, seqs[0])
// 	})
// 	return resp, err
// }

// ChangeGlobalAllocationParam changes GlobalAllocationParam with new value.
// It composes ChangeGlobalAllocationParamMsg and then broadcasts the transaction to blockchain.
// func (api *API) ChangeGlobalAllocationParam(ctx context.Context, creator string,
// 	parameter param.GlobalAllocationParam, reason string, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
// 	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(creator), func(seqs []uint64) ([]byte, errors.Error) {
// 		return api.MakeChangeGlobalAllocationParamMsg(creator, parameter, reason, privKeyHex, seqs[0])
// 	})
// 	return resp, err
// }

// ChangeInfraInternalAllocationParam changes InfraInternalAllocationParam with new value.
// It composes ChangeInfraInternalAllocationParamMsg and then broadcasts the transaction to blockchain.
// func (api *API) ChangeInfraInternalAllocationParam(
// 	ctx context.Context, creator string, parameter param.InfraInternalAllocationParam,
// 	reason string, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
// 	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(creator), func(seqs []uint64) ([]byte, errors.Error) {
// 		return api.MakeChangeInfraInternalAllocationParamMsg(creator, parameter, reason, privKeyHex, seqs[0])
// 	})
// 	return resp, err
// }

// ChangeVoteParam changes VoteParam with new value.
// It composes ChangeVoteParamMsg and then broadcasts the transaction to blockchain.
// func (api *API) ChangeVoteParam(
// 	ctx context.Context, creator string, parameter param.VoteParam,
// 	reason string, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
// 	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(creator), func(seqs []uint64) ([]byte, errors.Error) {
// 		return api.MakeChangeVoteParamMsg(creator, parameter, reason, privKeyHex, seqs[0])
// 	})
// 	return resp, err
// }

// ChangeProposalParam changes ProposalParam with new value.
// It composes ChangeProposalParamMsg and then broadcasts the transaction to blockchain.
// func (api *API) ChangeProposalParam(
// 	ctx context.Context, creator string, parameter param.ProposalParam,
// 	reason string, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
// 	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(creator), func(seqs []uint64) ([]byte, errors.Error) {
// 		return api.MakeChangeProposalParamMsg(creator, parameter, reason, privKeyHex, seqs[0])
// 	})
// 	return resp, err
// }

// ChangeDeveloperParam changes DeveloperParam with new value.
// It composes ChangeDeveloperParamMsg and then broadcasts the transaction to blockchain.
// func (api *API) ChangeDeveloperParam(
// 	ctx context.Context, creator string, parameter param.DeveloperParam,
// 	reason string, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
// 	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(creator), func(seqs []uint64) ([]byte, errors.Error) {
// 		return api.MakeChangeDeveloperParamMsg(creator, parameter, reason, privKeyHex, seqs[0])
// 	})
// 	return resp, err
// }

// ChangeValidatorParam changes ValidatorParam with new value.
// It composes ChangeValidatorParamMsg and then broadcasts the transaction to blockchain.
// func (api *API) ChangeValidatorParam(
// 	ctx context.Context, creator string, parameter param.ValidatorParam,
// 	reason string, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
// 	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(creator), func(seqs []uint64) ([]byte, errors.Error) {
// 		return api.MakeChangeValidatorParamMsg(creator, parameter, reason, privKeyHex, seqs[0])
// 	})
// 	return resp, err
// }

// ChangeBandwidthParam changes BandwidthParam with new value.
// It composes ChangeBandwidthParamMsg and then broadcasts the transaction to blockchain.
// func (api *API) ChangeBandwidthParam(
// 	ctx context.Context, creator string, parameter param.BandwidthParam,
// 	reason string, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
// 	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(creator), func(seqs []uint64) ([]byte, errors.Error) {
// 		return api.MakeChangeBandwidthParamMsg(creator, parameter, reason, privKeyHex, seqs[0])
// 	})
// 	return resp, err
// }

// ChangeAccountParam changes AccountParam with new value.
// It composes ChangeAccountParamMsg and then broadcasts the transaction to blockchain.
// func (api *API) ChangeAccountParam(
// 	ctx context.Context, creator string, parameter param.AccountParam,
// 	reason string, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
// 	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(creator), func(seqs []uint64) ([]byte, errors.Error) {
// 		return api.MakeChangeAccountParamMsg(creator, parameter, reason, privKeyHex, seqs[0])
// 	})
// 	return resp, err
// }

// ChangePostParam changes PostParam with new value.
// It composes ChangePostParamMsg and then broadcasts the transaction to blockchain.
// func (api *API) ChangePostParam(
// 	ctx context.Context, creator string, parameter param.PostParam,
// 	reason string, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
// 	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(creator), func(seqs []uint64) ([]byte, errors.Error) {
// 		return api.MakeChangePostParamMsg(creator, parameter, reason, privKeyHex, seqs[0])
// 	})
// 	return resp, err
// }

// DeletePostContent deletes the content of a post on blockchain, which is used
// for content censorship.
// It composes DeletePostContentMsg and then broadcasts the transaction to blockchain.
// func (api *API) DeletePostContent(
// 	ctx context.Context, creator, postAuthor, postID, reason,
// 	privKeyHex string) (*model.BroadcastResponse, errors.Error) {
// 	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(creator), func(seqs []uint64) ([]byte, errors.Error) {
// 		return api.MakeDeletePostContentMsg(creator, postAuthor, postID, reason, privKeyHex, seqs[0])
// 	})
// 	return resp, err
// }

// VoteProposal adds a vote to a certain proposal with agree/disagree.
// It composes VoteProposalMsg and then broadcasts the transaction to blockchain.
// func (api *API) VoteProposal(
// 	ctx context.Context, voter, proposalID string, result bool,
// 	privKeyHex string) (*model.BroadcastResponse, errors.Error) {
// 	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(voter), func(seqs []uint64) ([]byte, errors.Error) {
// 		return api.MakeVoteProposalMsg(voter, proposalID, result, privKeyHex, seqs[0])
// 	})
// 	return resp, err
// }

// UpgradeProtocol upgrades the protocol.
// It composes UpgradeProtocolMsg and then broadcasts the transaction to blockchain.
// func (api *API) UpgradeProtocol(ctx context.Context, creator, link, reason string, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
// 	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(creator), func(seqs []uint64) ([]byte, errors.Error) {
// 		return api.MakeUpgradeProtocolMsg(creator, link, reason, privKeyHex, seqs[0])
// 	})
// 	return resp, err
// }

// FeedPrice report lino price to blockchain.
func (api *API) FeedPrice(
	ctx context.Context, username string, price linotypes.MiniDollar, privKeyHex string) (*model.BroadcastResponse, errors.Error) {
	resp, _, err := api.GuaranteeBroadcast(ctx, util.GetSignerList(username), func(seqs []uint64) ([]byte, errors.Error) {
		return api.MakeFeedPriceMsg(username, price, privKeyHex, seqs[0])
	})
	return resp, err
}

// GuaranteeBroadcast - gurantee broadcast succ unless ctx timeout, which status is unknown.
// return response and an array of tx hash executed.
// WARNING-1: Use on lino fullnode version >= 0.2.10 ONLY!
// on lower version, txs may be executed twice
// WARNING-2: @p f, the MsgBuilderFunc, must be a pure function(no state, deterministic),
// otherwise, tx may be executed twice
func (api *API) GuaranteeBroadcast(ctx context.Context,
	signers []linotypes.AccOrAddr, f MsgBuilderFunc) (*model.BroadcastResponse, []string, errors.Error) {
	hashHistory := make([]string, 0)
	var lastHash *string // init: nil

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	stopped := false
	nRetried := 0
	for ; !stopped; <-ticker.C {
		select {
		case <-ctx.Done():
			stopped = true
		default:
		}

		resp, txHash, err := func() (*model.BroadcastResponse, *string, error) {
			broadcastCtx, cancel := context.WithTimeout(ctx, api.timeout)
			defer cancel()
			return api.safeBroadcastAndWatch(broadcastCtx, signers, lastHash, f)
		}()
		if txHash != nil {
			if lastHash == nil || (lastHash != nil && *txHash != *lastHash) {
				hashHistory = append(hashHistory, *txHash)
			}
		}
		if err == nil {
			return resp, hashHistory, nil
		}
		// The only place that does the retry.
		if err == errTxWatchTimeout || err == errSeqChanged || err == errSeqTxQueryFailed {
			if txHash != nil {
				lastHash = txHash
			}
			nRetried++
			continue
		}
		linoErr, ok := err.(errors.Error)
		if ok {
			return resp, hashHistory, linoErr
		}
		// This case shall never happen.
		return resp, hashHistory, errors.GuaranteeBroadcastFail(
			"returned error is not typed: " + err.Error())
	}
	return nil, hashHistory, errors.BroadcastTimeoutf(
		"GuaranteeBroadcast timeout, retried: %d", nRetried)
}

// this function ensure the safety of making a broadcast by doing a getSeq after getSeq, using
// GetTxAndSequenceNumber, if lastHash is provided.
// The safaty is guaranteed by that, seq number can advance IFF last tx does not exist in
// GetTxAndSequenceNumber.
func (api *API) safeBroadcastAndWatch(
	ctx context.Context, signers []linotypes.AccOrAddr, lastHash *string,
	f MsgBuilderFunc) (*model.BroadcastResponse, *string, error) {
	currentSeqs := make([]uint64, len(signers))
	if lastHash == nil {
		for i, signer := range signers {
			var seq uint64
			var err error
			if signer.IsAddr {
				addr := hex.EncodeToString(signer.Addr)
				seq, err = api.Query.GetSeqNumberByAddress(ctx, addr)
			} else {
				seq, err = api.Query.GetSeqNumber(ctx, string(signer.AccountKey))
			}
			if err != nil {
				return nil, lastHash, errSeqTxQueryFailed
			}
			currentSeqs[i] = seq
			if i > 0 && checkEqual(signers[i], signers[i-1]) {
				currentSeqs[i] += 1
			}
		}
	} else {
		// XXX(yumin): GetTxAndSequenceNumber does GetSeq then GetTx to ensure that if seq changed,
		// the original tx is not applied, if last hash is not nil.
		for i, signer := range signers {
			var txSeq *accmodel.TxAndSequenceNumber
			var err error
			if signer.IsAddr {
				addr := hex.EncodeToString(signer.Addr)
				txSeq, err = api.Query.GetTxAndSequenceNumberByAddress(ctx, addr, *lastHash)
				if err != nil {
					return nil, lastHash, errSeqTxQueryFailed
				}
			} else {
				txSeq, err = api.Query.GetTxAndSequenceNumberByUsername(ctx, string(signer.AccountKey), *lastHash)
				if err != nil {
					return nil, lastHash, errSeqTxQueryFailed
				}
			}

			// alreay succeeded
			if txSeq.Tx != nil {
				return &model.BroadcastResponse{
					Height:     txSeq.Tx.Height,
					CommitHash: txSeq.Tx.Hash,
				}, lastHash, nil
			}
			currentSeqs[i] = txSeq.Sequence
			if i > 0 && checkEqual(signers[i], signers[i-1]) {
				currentSeqs[i] += 1
			}
		}
	}

	msgBytes, err := f(currentSeqs)
	if err != nil {
		return nil, lastHash, err
	}
	newHash, err := broadcast.CalcTxMsgHashHexString(msgBytes)
	if err != nil {
		return nil, lastHash, err
	}

	// XXX(yumin): so bad that GetTxAndSequenceNumber is broken for now because txinder
	// is async. In this case, we will redo the query N times, and allow it only
	// when stay the same for that many times of queries.
	// Note that, it Guarantee NOTHING, only likely to be just ok. DO NOT use in
	// important txs.
	if lastHash != nil && *lastHash != newHash {
		N := 5
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for i := 0; i < N; i++ {
			for index, signer := range signers {
				var txSeq *accmodel.TxAndSequenceNumber
				var err error
				if signer.IsAddr {
					addr := hex.EncodeToString(signer.Addr)
					txSeq, err = api.Query.GetTxAndSequenceNumberByAddress(ctx, addr, *lastHash)
					if err != nil {
						return nil, lastHash, errSeqTxQueryFailed
					}
				} else {
					txSeq, err = api.Query.GetTxAndSequenceNumberByUsername(ctx, string(signer.AccountKey), *lastHash)
					if err != nil {
						return nil, lastHash, errSeqTxQueryFailed
					}
				}
				// not stabled.
				if txSeq.Sequence != currentSeqs[index] {
					return nil, lastHash, errSeqTxQueryFailed
				}
				// well it actually succeeded.
				if txSeq.Tx != nil {
					return &model.BroadcastResponse{
						Height:     txSeq.Tx.Height,
						CommitHash: txSeq.Tx.Hash,
					}, lastHash, nil
				}
			}
			<-ticker.C
		}
	}

	bres, berr := api.broadcastAndWatch(ctx, msgBytes, currentSeqs[0])
	if berr != nil {
		return nil, &newHash, berr
	}
	return bres, &newHash, nil
}

// unsafe, make sure the @p seq is a conservative value that won't do f twice.
func (api *API) broadcastAndWatch(ctx context.Context, msg []byte, seq uint64) (*model.BroadcastResponse, error) {
	err := api.Broadcast.BroadcastRawMsgBytesSync(ctx, msg, seq)
	if err != nil {
		// can retry.
		if err.CodeType() == errors.CodeInvalidSequenceNumber {
			return nil, errSeqChanged
		}

		// only in case that (tx in cache) or (timeout), we continue polling.
		if err.CodeType() == errors.CodeFailedToBroadcast &&
			strings.Contains(err.Error(), "Tx already exists in cache") {
			// do nothing and start to polling.
		} else if err.CodeType() == errors.CodeTimeout ||
			err.CodeType() == errors.CodeBroadcastTimeout {
			// no-op, fallthrough to polling.
		} else {
			return nil, err
		}
	}

	// polling tx commit hash
	hashBytes, _ := broadcast.CalcTxMsgHash(msg) // msg passed checktx won't trigger error here.
	ticker := time.NewTicker(api.checkTxConfirmInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			tx, err := api.GetTx(ctx, hashBytes)
			// keep retry
			if err != nil {
				continue
			}
			// if code is not ok (0), report err
			if tx.Code != 0 {
				return nil, errors.DeliverTxFail("deliver tx failed").
					AddBlockChainCode(tx.Code).AddBlockChainLog(tx.Log)
			}
			return &model.BroadcastResponse{
				CommitHash: hex.EncodeToString(hashBytes),
				Height:     tx.Height,
			}, nil
		case <-ctx.Done():
			// can retry
			return nil, errTxWatchTimeout
		}
	}
}

func checkEqual(a1, a2 linotypes.AccOrAddr) bool {
	if a1.IsAddr != a2.IsAddr {
		return false
	}
	if !(a1.Addr.Equals(a2.Addr)) {
		return false
	}
	if a1.AccountKey != a2.AccountKey {
		return false
	}
	return true
}
