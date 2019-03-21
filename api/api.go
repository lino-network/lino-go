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

// MsgBuilderFunc is usually a closure that return messages bytes for a specific sequence.
type MsgBuilderFunc func(seq uint64) ([]byte, errors.Error)

// GuaranteeBroadcast - gurantee broadcast succ unless ctx timeout, which status is unknown.
// XXX(yumin): BROKEN now, not recommended to use.
func (api *API) GuaranteeBroadcast(ctx context.Context,
	username string, f MsgBuilderFunc) (*model.BroadcastResponse, errors.Error) {
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
			nRetried++
			continue
		}
		linoErr, ok := err.(errors.Error)
		if ok {
			return resp, linoErr
		}
		// This case shall never happen.
		return resp, errors.GuaranteeBroadcastFail("returned error is not typed: " + err.Error())
	}
	return nil, errors.BroadcastTimeoutf("GuaranteeBroadcast timeout, retried: %d", nRetried)
}

// this function ensure the safety of making a broadcast by doing a getSeq after getSeq, using
// GetTxAndSequenceNumber, if lastHash is provided.
// The safaty is guaranteed by that, seq number can advance IFF last tx does not exist in
// GetTxAndSequenceNumber.
func (api *API) safeBroadcastAndWatch(ctx context.Context, username string, lastHash *string,
	f MsgBuilderFunc) (*model.BroadcastResponse, *string, error) {
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
	msgBytes, err := f(currentSeq)
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
			txSeqCheck, seqErr := api.Query.GetTxAndSequenceNumber(ctx, username, *lastHash)
			if seqErr != nil {
				return nil, lastHash, errSeqTxQueryFailed
			}
			// not stabled.
			if txSeqCheck.Sequence != currentSeq {
				return nil, lastHash, errSeqTxQueryFailed
			}
			// well it actually succeeded.
			if txSeqCheck.Tx != nil {
				return &model.BroadcastResponse{
					Height:     txSeqCheck.Tx.Height,
					CommitHash: txSeqCheck.Tx.Hash,
				}, lastHash, nil
			}
			<-ticker.C
		}
	}

	bres, berr := api.broadcastAndWatch(ctx, msgBytes, currentSeq)
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
	hashBytes, _ := broadcast.CalcTxMsgHash(msg) // msg passed chectx won't meet error here.
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
