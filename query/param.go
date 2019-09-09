package query

import (
	"context"

	"github.com/lino-network/lino/param"
)

// GetGlobalAllocationParam returns the GlobalAllocationParam.
func (query *Query) GetGlobalAllocationParam(ctx context.Context) (*param.GlobalAllocationParam, error) {
	resp, err := query.transport.Query(ctx, ParamKVStoreKey, AllocationParamSubStore, []string{})
	if err != nil {
		return nil, err
	}

	param := new(param.GlobalAllocationParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetInfraInternalAllocationParam returns the InfraInternalAllocationParam.
func (query *Query) GetInfraInternalAllocationParam(ctx context.Context) (*param.InfraInternalAllocationParam, error) {
	resp, err := query.transport.Query(ctx, ParamKVStoreKey, InfraInternalAllocationParamSubStore, []string{})
	if err != nil {
		return nil, err
	}

	param := new(param.InfraInternalAllocationParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetDeveloperParam returns the DeveloperParam.
func (query *Query) GetDeveloperParam(ctx context.Context) (*param.DeveloperParam, error) {
	resp, err := query.transport.Query(ctx, ParamKVStoreKey, DeveloperParamSubStore, []string{})
	if err != nil {
		return nil, err
	}

	param := new(param.DeveloperParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetVoteParam returns the VoteParam.
func (query *Query) GetVoteParam(ctx context.Context) (*param.VoteParam, error) {
	resp, err := query.transport.Query(ctx, ParamKVStoreKey, VoteParamSubStore, []string{})
	if err != nil {
		return nil, err
	}

	param := new(param.VoteParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetProposalParam returns the ProposalParam.
func (query *Query) GetProposalParam(ctx context.Context) (*param.ProposalParam, error) {
	resp, err := query.transport.Query(ctx, ParamKVStoreKey, ProposalParamSubStore, []string{})
	if err != nil {
		return nil, err
	}

	param := new(param.ProposalParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetValidatorParam returns the ValidatorParam.
func (query *Query) GetValidatorParam(ctx context.Context) (*param.ValidatorParam, error) {
	resp, err := query.transport.Query(ctx, ParamKVStoreKey, ValidatorParamSubStore, []string{})
	if err != nil {
		return nil, err
	}

	param := new(param.ValidatorParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetCoinDayParam returns the CoinDayParam.
func (query *Query) GetCoinDayParam(ctx context.Context) (*param.CoinDayParam, error) {
	resp, err := query.transport.Query(ctx, ParamKVStoreKey, CoinDayParamSubStore, []string{})
	if err != nil {
		return nil, err
	}

	param := new(param.CoinDayParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetBandwidthParam returns the BandwidthParam.
func (query *Query) GetBandwidthParam(ctx context.Context) (*param.BandwidthParam, error) {
	resp, err := query.transport.Query(ctx, ParamKVStoreKey, BandwidthParamSubStore, []string{})
	if err != nil {
		return nil, err
	}

	param := new(param.BandwidthParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetAccountParam returns the AccountParam.
func (query *Query) GetAccountParam(ctx context.Context) (*param.AccountParam, error) {
	resp, err := query.transport.Query(ctx, ParamKVStoreKey, AccountParamSubStore, []string{})
	if err != nil {
		return nil, err
	}

	param := new(param.AccountParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetPostParam returns the PostParam.
func (query *Query) GetPostParam(ctx context.Context) (*param.PostParam, error) {
	resp, err := query.transport.Query(ctx, ParamKVStoreKey, PostParamSubStore, []string{})
	if err != nil {
		return nil, err
	}

	param := new(param.PostParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}
