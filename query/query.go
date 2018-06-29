// Package query includes the functionalities to query
// data from the blockchain.
package query

import (
	"github.com/lino-network/lino-go/errors"
	"github.com/lino-network/lino-go/model"
	"github.com/lino-network/lino-go/transport"
)

// Query is a wrapper of querying data from blockchain.
type Query struct {
	transport *transport.Transport
}

// NewQuery returns an instance of Query.
func NewQuery(transport *transport.Transport) *Query {
	return &Query{
		transport: transport,
	}
}

// GetBlock returns a block at a certain height from blockchain.
func (query *Query) GetBlock(height int64) (*model.Block, error) {
	resp, err := query.transport.QueryBlock(height)
	if err != nil {
		return nil, errors.QueryFailf("GetBlock err").AddCause(err)
	}

	block := new(model.Block)
	block.Header = resp.Block.Header
	block.Evidence = resp.Block.Evidence
	block.LastCommit = resp.Block.LastCommit
	block.Data = new(model.Data)
	block.Data.Txs = []model.Transaction{}
	for _, txBytes := range resp.Block.Data.Txs {
		var tx model.Transaction
		if err := query.transport.Cdc.UnmarshalJSON(txBytes, &tx); err != nil {
			return nil, err
		}
		block.Data.Txs = append(block.Data.Txs, tx)
	}
	return block, nil
}
