package query

import (
	"context"

	"github.com/lino-network/lino-go/errors"
	linotypes "github.com/lino-network/lino/types"
	"github.com/lino-network/lino/x/price/model"
	"github.com/lino-network/lino/x/price/types"
)

// GetLastFeed returns the last fed price of @p validator.
func (query *Query) GetLastFeed(ctx context.Context, validator string) (*model.FedPrice, error) {
	resp, err := query.transport.Query(ctx, types.QuerierRoute, types.QueryLastFeed,
		[]string{validator})
	if err != nil {
		linoe, ok := err.(errors.Error)
		if ok && linoe.BlockChainCode() == uint32(linotypes.CodeFedPriceNotFound) {
			return nil, errors.EmptyResponse("last fed price is not found")
		}
		return nil, err
	}
	developer := new(model.FedPrice)
	if err := query.transport.Cdc.UnmarshalJSON(resp, developer); err != nil {
		return nil, err
	}
	return developer, nil
}

// GetCurrentPrice returns the last fed price of @p validator.
func (query *Query) GetCurrentPrice(ctx context.Context) (linotypes.MiniDollar, error) {
	resp, err := query.transport.Query(ctx, types.QuerierRoute, types.QueryPriceCurrent, []string{})
	if err != nil {
		return linotypes.NewMiniDollar(0), err
	}
	rst := new(linotypes.MiniDollar)
	if err := query.transport.Cdc.UnmarshalJSON(resp, rst); err != nil {
		return linotypes.NewMiniDollar(0), err
	}
	return *rst, nil
}

// GetHistoryPrice returns the last fed price of @p validator.
func (query *Query) GetHistoryPrice(ctx context.Context) ([]model.FeedHistory, error) {
	resp, err := query.transport.Query(ctx, types.QuerierRoute, types.QueryPriceHistory, []string{})
	if err != nil {
		return nil, err
	}
	rst := make([]model.FeedHistory, 0)
	if err := query.transport.Cdc.UnmarshalJSON(resp, &rst); err != nil {
		return nil, err
	}
	return rst, nil
}
