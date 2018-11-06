// Package query includes the functionalities to query
// data from the blockchain.
package query

import (
	"context"

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
func (query *Query) GetBlock(ctx context.Context, height int64) (*model.Block, error) {
	resp, err := query.transport.QueryBlock(ctx, height)
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

// GetBlockStatus returns the current block status from blockchain.
func (query *Query) GetBlockStatus(ctx context.Context) (*model.BlockStatus, error) {
	resp, err := query.transport.QueryBlockStatus(ctx)
	if err != nil {
		return nil, errors.QueryFailf("GetBlockStatus err").AddCause(err)
	}

	bs := &model.BlockStatus{
		LatestBlockHeight: resp.SyncInfo.LatestBlockHeight,
		LatestBlockTime:   resp.SyncInfo.LatestBlockTime,
	}

	return bs, nil
}

func (query *Query) GetTx(ctx context.Context, hash []byte) (*model.BlockTx, error) {
	resp, err := query.transport.QueryTx(ctx, hash)
	if err != nil {
		return nil, errors.QueryFailf("GetTx err").AddCause(err)
	}

	var tx model.Transaction
	if err := query.transport.Cdc.UnmarshalJSON(resp.Tx, &tx); err != nil {
		return nil, err
	}

	bt := &model.BlockTx{
		Height: resp.Height,
		Tx:     tx,
	}

	return bt, nil
}
