package query

import (
	"context"

	"github.com/lino-network/lino-go/errors"
	linotypes "github.com/lino-network/lino/types"
	"github.com/lino-network/lino/x/developer/model"
	"github.com/lino-network/lino/x/developer/types"
)

// GetDeveloper returns a specific developer info from blockchain.
func (query *Query) GetDeveloper(ctx context.Context, developerName string) (*model.Developer, error) {
	resp, err := query.transport.Query(ctx, DeveloperKVStoreKey, types.QueryDeveloper, []string{developerName})
	if err != nil {
		linoe, ok := err.(errors.Error)
		if ok && linoe.BlockChainCode() == uint32(linotypes.CodeDeveloperNotFound) {
			return nil, errors.EmptyResponse("developer is not found")
		}
		return nil, err
	}
	developer := new(model.Developer)
	if err := query.transport.Cdc.UnmarshalJSON(resp, developer); err != nil {
		return nil, err
	}
	return developer, nil
}

// GetIDABalance returns user IDA balance
func (query *Query) GetIDABalance(ctx context.Context, username, app string) (*types.QueryResultIDABalance, error) {
	resp, err := query.transport.Query(ctx, DeveloperKVStoreKey, types.QueryIDABalance, []string{app, username})
	if err != nil {
		linoe, ok := err.(errors.Error)
		if ok && linoe.BlockChainCode() == uint32(linotypes.CodeAccountNotFound) {
			return nil, errors.EmptyResponse("ida bank is not found")
		}
		return nil, err
	}
	bank := new(types.QueryResultIDABalance)
	if err := query.transport.Cdc.UnmarshalJSON(resp, bank); err != nil {
		return nil, err
	}
	return bank, nil
}

// GetIDA returns App IDA info
func (query *Query) GetIDA(ctx context.Context, developerName string) (*model.AppIDA, error) {
	resp, err := query.transport.Query(ctx, DeveloperKVStoreKey, types.QueryIDA, []string{developerName})
	if err != nil {
		linoe, ok := err.(errors.Error)
		if ok && linoe.BlockChainCode() == uint32(linotypes.CodeIDANotFound) {
			return nil, errors.EmptyResponse("ida is not found")
		}
		return nil, err
	}
	ida := new(model.AppIDA)
	if err := query.transport.Cdc.UnmarshalJSON(resp, ida); err != nil {
		return nil, err
	}
	return ida, nil
}

// GetAffiliated returns App affiliated account list
func (query *Query) GetAffiliated(ctx context.Context, developerName string) ([]string, error) {
	resp, err := query.transport.Query(ctx, DeveloperKVStoreKey, types.QueryAffiliated, []string{developerName})
	if err != nil {
		linoe, ok := err.(errors.Error)
		if ok && linoe.BlockChainCode() == uint32(linotypes.CodeDeveloperNotFound) {
			return nil, errors.EmptyResponse("developer is not found")
		}
		return nil, err
	}
	var affiliatedAccs []string
	if err := query.transport.Cdc.UnmarshalJSON(resp, affiliatedAccs); err != nil {
		return nil, err
	}
	return affiliatedAccs, nil
}

// GetReservePool returns App affiliated account list
func (query *Query) GetReservePool(ctx context.Context) (*model.ReservePool, error) {
	resp, err := query.transport.Query(ctx, DeveloperKVStoreKey, types.QueryReservePool, []string{})
	if err != nil {
		return nil, err
	}
	reservePool := new(model.ReservePool)
	if err := query.transport.Cdc.UnmarshalJSON(resp, reservePool); err != nil {
		return nil, err
	}
	return reservePool, nil
}

// GetIDAStats returns App IDA stats
func (query *Query) GetIDAStats(ctx context.Context, developerName string) (*model.AppIDAStats, error) {
	resp, err := query.transport.Query(ctx, DeveloperKVStoreKey, types.QueryIDAStats, []string{developerName})
	if err != nil {
		return nil, err
	}
	IDAStats := new(model.AppIDAStats)
	if err := query.transport.Cdc.UnmarshalJSON(resp, IDAStats); err != nil {
		return nil, err
	}
	return IDAStats, nil
}

// // GetDevelopers returns a list of all developers.
// func (query *Query) GetDevelopers(ctx context.Context) (*model.DeveloperList, error) {
// 	resp, err := query.transport.Query(ctx, DeveloperKVStoreKey, DeveloperListSubStore, []string{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	developerList := new(model.DeveloperList)
// 	if err := query.transport.Cdc.UnmarshalJSON(resp, developerList); err != nil {
// 		return nil, err
// 	}
// 	return developerList, nil
// }
