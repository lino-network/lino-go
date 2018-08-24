// Package api initiates a go library API which can be
// used to query data from blockchain and
// broadcast transactions to blockchain.
package api

import (
	"time"

	"github.com/lino-network/lino-go/broadcast"
	"github.com/lino-network/lino-go/query"
	"github.com/lino-network/lino-go/transport"
)

// API is a wrapper of both querying data from blockchain
// and broadcast transactions to blockchain.
type API struct {
	*query.Query
	*broadcast.Broadcast
}

type TimeoutOptions struct {
	QueryTimeout     time.Duration
	BroadcastTimeout time.Duration
}

// NewLinoAPIFromConfig initiates an instance of API using
// configs from ~/.lino-go/config.json
func NewLinoAPIFromConfig(options TimeoutOptions) *API {
	transport := transport.NewTransportFromConfig(options.QueryTimeout)
	return &API{
		Query:     query.NewQuery(transport),
		Broadcast: broadcast.NewBroadcast(transport, options.BroadcastTimeout),
	}
}

// NewLinoAPIFromArgs initiates an instance of API using
// chainID and nodeUrl that are passed in.
func NewLinoAPIFromArgs(chainID, nodeUrl string, options TimeoutOptions) *API {
	transport := transport.NewTransportFromArgs(chainID, nodeUrl, options.QueryTimeout)
	return &API{
		Query:     query.NewQuery(transport),
		Broadcast: broadcast.NewBroadcast(transport, options.BroadcastTimeout),
	}
}
