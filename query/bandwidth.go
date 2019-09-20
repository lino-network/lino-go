package query

import (
	"context"
	"github.com/lino-network/lino-go/errors"
	linotypes "github.com/lino-network/lino/types"
	"github.com/lino-network/lino/x/bandwidth/model"
	"github.com/lino-network/lino/x/bandwidth/types"
)

func (query *Query) GetBandwidthInfo(ctx context.Context) (*model.BandwidthInfo, error) {
	resp, err := query.transport.Query(ctx, BandwidthKVStoreKey, types.QueryBandwidthInfo, []string{})
	if err != nil {
		linoe, ok := err.(errors.Error)
		if ok && linoe.BlockChainCode() == uint32(linotypes.CodeBandwidthInfoNotFound) {
			return nil, errors.EmptyResponse("bandwidth info is not found")
		}
		return nil, err
	}
	info := new(model.BandwidthInfo)
	if err := query.transport.Cdc.UnmarshalJSON(resp, info); err != nil {
		return nil, err
	}
	return info, nil
}

func (query *Query) GetBlockInfo(ctx context.Context) (*model.BlockInfo, error) {
	resp, err := query.transport.Query(ctx, BandwidthKVStoreKey, types.QueryBlockInfo, []string{})
	if err != nil {
		linoe, ok := err.(errors.Error)
		if ok && linoe.BlockChainCode() == uint32(linotypes.CodeBlockInfoNotFound) {
			return nil, errors.EmptyResponse("block info is not found")
		}
		return nil, err
	}
	info := new(model.BlockInfo)
	if err := query.transport.Cdc.UnmarshalJSON(resp, info); err != nil {
		return nil, err
	}
	return info, nil
}

func (query *Query) GetAppBandwidthInfo(ctx context.Context, username string) (*model.AppBandwidthInfo, error) {
	resp, err := query.transport.Query(ctx, BandwidthKVStoreKey, types.QueryAppBandwidthInfo, []string{username})
	if err != nil {
		linoe, ok := err.(errors.Error)
		if ok && linoe.BlockChainCode() == uint32(linotypes.CodeAppBandwidthInfoNotFound) {
			return nil, errors.EmptyResponse("app bandwidth info is not found")
		}
		return nil, err
	}
	info := new(model.AppBandwidthInfo)
	if err := query.transport.Cdc.UnmarshalJSON(resp, info); err != nil {
		return nil, err
	}
	return info, nil
}
