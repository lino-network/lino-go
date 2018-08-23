package query

import (
	"context"

	"github.com/lino-network/lino-go/model"
)

// GetEvaluateOfContentValueParam returns the EvaluateOfContentValueParam.
func (query *Query) GetEvaluateOfContentValueParam(ctx context.Context) (*model.EvaluateOfContentValueParam, error) {
	resp, err := query.transport.Query(ctx, getEvaluateOfContentValueParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.EvaluateOfContentValueParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetGlobalAllocationParam returns the GlobalAllocationParam.
func (query *Query) GetGlobalAllocationParam(ctx context.Context) (*model.GlobalAllocationParam, error) {
	resp, err := query.transport.Query(ctx, getGlobalAllocationParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.GlobalAllocationParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetInfraInternalAllocationParam returns the InfraInternalAllocationParam.
func (query *Query) GetInfraInternalAllocationParam(ctx context.Context) (*model.InfraInternalAllocationParam, error) {
	resp, err := query.transport.Query(ctx, getInfraInternalAllocationParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.InfraInternalAllocationParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetDeveloperParam returns the DeveloperParam.
func (query *Query) GetDeveloperParam(ctx context.Context) (*model.DeveloperParam, error) {
	resp, err := query.transport.Query(ctx, getDeveloperParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.DeveloperParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetVoteParam returns the VoteParam.
func (query *Query) GetVoteParam(ctx context.Context) (*model.VoteParam, error) {
	resp, err := query.transport.Query(ctx, getVoteParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.VoteParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetProposalParam returns the ProposalParam.
func (query *Query) GetProposalParam(ctx context.Context) (*model.ProposalParam, error) {
	resp, err := query.transport.Query(ctx, getProposalParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.ProposalParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetValidatorParam returns the ValidatorParam.
func (query *Query) GetValidatorParam(ctx context.Context) (*model.ValidatorParam, error) {
	resp, err := query.transport.Query(ctx, getValidatorParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.ValidatorParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetCoinDayParam returns the CoinDayParam.
func (query *Query) GetCoinDayParam(ctx context.Context) (*model.CoinDayParam, error) {
	resp, err := query.transport.Query(ctx, getCoinDayParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.CoinDayParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetBandwidthParam returns the BandwidthParam.
func (query *Query) GetBandwidthParam(ctx context.Context) (*model.BandwidthParam, error) {
	resp, err := query.transport.Query(ctx, getBandwidthParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.BandwidthParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetAccountParam returns the AccountParam.
func (query *Query) GetAccountParam(ctx context.Context) (*model.AccountParam, error) {
	resp, err := query.transport.Query(ctx, getAccountParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.AccountParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetPostParam returns the PostParam.
func (query *Query) GetPostParam(ctx context.Context) (*model.PostParam, error) {
	resp, err := query.transport.Query(ctx, getPostParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.PostParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}
