// Package api initiates a go library API which can be
// used to query data from blockchain and
// broadcast transactions to blockchain.
package api

import (
	"time"

	"github.com/lino-network/lino-go/broadcast"
	"github.com/lino-network/lino-go/query"
	"github.com/lino-network/lino-go/transport"
	"github.com/spf13/viper"
)

// API is a wrapper of both querying data from blockchain
// and broadcast transactions to blockchain.
type API struct {
	*query.Query
	*broadcast.Broadcast
}

// Options is a wrapper of init parameters
type Options struct {
	ChainID            string        `json:"chain_id"`
	NodeURL            string        `json:"node_url"`
	MaxAttempts        int64         `json:"max_attempts"`
	InitSleepTime      time.Duration `json:"init_sleep_time"`
	ExponentialBackoff bool          `json:"exponential_back_off"`
	BackoffRandomness  bool          `json:"backoff_randomness"`
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
func NewLinoAPIFromArgs(options *Options) *API {
	transport := transport.NewTransportFromArgs(options.ChainID, options.NodeURL)
	return &API{
		Query:     query.NewQuery(transport),
		Broadcast: broadcast.NewBroadcast(transport, options.MaxAttempts, options.InitSleepTime, options.ExponentialBackoff, options.BackoffRandomness),
	}
}
