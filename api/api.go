// Package api initiates a go library API which can be
// used to query data from blockchain and
// broadcast transactions to blockchain.
package api

import (
	"context"
	"encoding/hex"
	goerrors "errors"
	"time"

	"github.com/lino-network/lino-go/broadcast"
	"github.com/lino-network/lino-go/errors"
	"github.com/lino-network/lino-go/model"
	"github.com/lino-network/lino-go/query"
	"github.com/lino-network/lino-go/transport"
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
	InitSleepTime          time.Duration `json:"init_sleep_time"`
	Timeout                time.Duration `json:"timeout"`
	ExponentialBackoff     bool          `json:"exponential_back_off"`
	BackoffRandomness      bool          `json:"backoff_randomness"`
	FixSequenceNumber      bool          `json:"fix_sequence_number"`
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
	transport := transport.NewTransportFromArgs(opt.ChainID, opt.NodeURL)
	return &API{
		Query:                  query.NewQuery(transport),
		Broadcast:              broadcast.NewBroadcast(transport, opt.MaxAttempts, opt.InitSleepTime, opt.Timeout, opt.ExponentialBackoff, opt.BackoffRandomness, opt.FixSequenceNumber),
		checkTxConfirmInterval: opt.CheckTxConfirmInterval,
		timeout:                opt.Timeout,
	}
}

// GuaranteeBroadcast - gurantee broadcast succ unless ctx timeout, which status is unknown.
func (api *API) GuaranteeBroadcast(ctx context.Context, username string,
	f func(ctx context.Context, seq uint64) (*model.BroadcastResponse, errors.Error),
) (*model.BroadcastResponse, errors.Error) {
	if !api.Broadcast.FixSequenceNumber {
		return nil, errors.GuaranteeBroadcastFail(
			"only fix sequence number can guarantee broadcast")
	}

	var lastHash *string // init: nil
	for {
		resp, txHash, err := func() (*model.BroadcastResponse, *string, error) {
			broadcastCtx, cancel := context.WithTimeout(ctx, api.timeout)
			defer cancel()
			return api.safeBroadcastAndWatch(broadcastCtx, username, lastHash, f)
		}()
		if err == nil {
			return resp, nil
		}
		// The only place that does the retry.
		if err == errTxWatchTimeout || err == errSeqChanged || err == errSeqTxQueryFailed {
			if txHash != nil {
				lastHash = txHash
			}
			continue
		}
		linoErr, ok := err.(errors.Error)
		if ok {
			return resp, linoErr
		}
		// This case shall never happen.
		return resp, errors.GuaranteeBroadcastFail("returned error is not typed: " + linoErr.Error())
	}
}

// this function ensure the safety of making a broadcast by doing a getSeq after getSeq, using
// GetTxAndSequenceNumber, if lastHash is provided.
func (api *API) safeBroadcastAndWatch(ctx context.Context, username string, lastHash *string,
	f func(ctx context.Context, seq uint64) (*model.BroadcastResponse, errors.Error),
) (*model.BroadcastResponse, *string, error) {
	var currentSeq uint64 // 0
	if lastHash == nil {
		var seqErr error
		currentSeq, seqErr = api.Query.GetSeqNumber(ctx, username)
		if seqErr != nil {
			return nil, lastHash, errSeqTxQueryFailed
		}
	} else {
		// XXX(yumin): GetTxAndSequenceNumber does GetSeq then GetTx to ensure that if seq changed,
		// the original tx is not applied, if last hash is not nil.
		txSeq, seqErr := api.Query.GetTxAndSequenceNumber(ctx, username, *lastHash)
		if seqErr != nil {
			return nil, lastHash, errSeqTxQueryFailed
		}

		// alreay succeeded
		if txSeq.Tx != nil {
			return &model.BroadcastResponse{
				Height:     txSeq.Tx.Height,
				CommitHash: txSeq.Tx.Hash,
			}, lastHash, nil
		}
		currentSeq = txSeq.Sequence
	}
	return api.broadcastAndWatch(ctx, currentSeq, lastHash, f)
}

// unsafe, make sure the @p seq is a conservative value that won't do f twice.
func (api *API) broadcastAndWatch(ctx context.Context, seq uint64, lastHash *string,
	f func(ctx context.Context, seq uint64) (*model.BroadcastResponse, errors.Error),
) (*model.BroadcastResponse, *string, error) {
	resp, err := f(ctx, seq)
	if err != nil {
		// can retry.
		if err.CodeType() == errors.CodeInvalidSequenceNumber {
			return nil, lastHash, errSeqChanged
		}
		return nil, lastHash, err
	}

	// check tx commit hash
	commitHash := resp.CommitHash
	commitHashBytes, decodeErr := hex.DecodeString(resp.CommitHash)
	if decodeErr != nil {
		return nil, lastHash, errors.GuaranteeBroadcastFail("commit hash invalid")
	}

	ticker := time.NewTicker(api.checkTxConfirmInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			tx, err := api.GetTx(ctx, commitHashBytes)
			// keep retry
			if err != nil {
				continue
			}
			// if code is not ok (0), report err
			if tx.Code != 0 {
				return nil, &commitHash, errors.DeliverTxFail(
					"deliver tx failed").AddBlockChainCode(tx.Code).AddBlockChainLog(tx.Log)
			}
			resp.Height = tx.Height
			return resp, &commitHash, nil
		case <-ctx.Done():
			// can retry
			return nil, &commitHash, errTxWatchTimeout
		}
	}
}
