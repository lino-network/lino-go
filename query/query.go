// Package query includes the functionalities to query
// data from the blockchain.
package query

import (
	"context"
	"strings"

	auth "github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/lino-network/lino-go/errors"
	"github.com/lino-network/lino-go/model"
	"github.com/lino-network/lino-go/transport"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	ttypes "github.com/tendermint/tendermint/types"
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
func (query *Query) GetBlock(ctx context.Context, height int64) (*ttypes.Block, error) {
	resp, err := query.transport.QueryBlock(ctx, height)
	if err != nil {
		return nil, errors.QueryFailf("GetBlock err").AddCause(err)
	}

	return resp.Block, nil
}

// GetBlockStatus returns the current block status from blockchain.
func (query *Query) GetBlockStatus(ctx context.Context) (*ctypes.ResultStatus, error) {
	resp, err := query.transport.QueryBlockStatus(ctx)
	if err != nil {
		return nil, errors.QueryFailf("GetBlockStatus err").AddCause(err)
	}

	return resp, nil
}

func (query *Query) GetTx(ctx context.Context, hash []byte) (*model.BlockTx, errors.Error) {
	resp, err := query.transport.QueryTx(ctx, hash)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			// if tx already exists in cache
			return nil, errors.QueryTxNotFound().AddCause(err)
		}
		return nil, errors.QueryFailf("GetTx err").AddCause(err)
	}

	var tx auth.StdTx
	if err := query.transport.Cdc.UnmarshalJSON(resp.Tx, &tx); err != nil {
		return nil, errors.QueryFailf("Unmarshal tx err").AddCause(err)
	}

	bt := &model.BlockTx{
		Height: resp.Height,
		Tx:     tx,
		Code:   resp.TxResult.Code,
		Log:    resp.TxResult.Log,
	}

	return bt, nil
}
