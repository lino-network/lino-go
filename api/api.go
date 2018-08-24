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
	QueryTimeout     time.Duration // query usually takes less than 1 second, so it can be set to 1 second.
	BroadcastTimeout time.Duration // broadcast usually takes 3 seconds, so it can be set to 4 second.
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
