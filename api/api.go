package api

import (
	"github.com/lino-network/lino-go/broadcast"
	"github.com/lino-network/lino-go/query"
	"github.com/lino-network/lino-go/transport"
)

type API struct {
	*query.Query
	*broadcast.Broadcast
}

func DefaultLinoAPI() *API {
	transport := transport.NewTransportFromViper()
	return &API{
		Query:     query.NewQuery(transport),
		Broadcast: broadcast.NewBroadcast(transport),
	}
}

func NewLionAPI(chainID, nodeUrl string) *API {
	transport := transport.NewTransportFromArgs(chainID, nodeUrl)
	return &API{
		Query:     query.NewQuery(transport),
		Broadcast: broadcast.NewBroadcast(transport),
	}
}
