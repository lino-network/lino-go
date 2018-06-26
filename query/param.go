package query

import (
	"github.com/lino-network/lino-go/model"
)

//
// param related query
//
func (query *Query) GetEvaluateOfContentValueParam() (*model.EvaluateOfContentValueParam, error) {
	resp, err := query.transport.Query(getEvaluateOfContentValueParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.EvaluateOfContentValueParam)
	// fmt.Println("---paramBytes: ", resp)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

func (query *Query) GetGlobalAllocationParam() (*model.GlobalAllocationParam, error) {
	resp, err := query.transport.Query(getGlobalAllocationParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.GlobalAllocationParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

func (query *Query) GetInfraInternalAllocationParam() (*model.InfraInternalAllocationParam, error) {
	resp, err := query.transport.Query(getInfraInternalAllocationParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.InfraInternalAllocationParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

func (query *Query) GetDeveloperParam() (*model.DeveloperParam, error) {
	resp, err := query.transport.Query(getDeveloperParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.DeveloperParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

func (query *Query) GetVoteParam() (*model.VoteParam, error) {
	resp, err := query.transport.Query(getVoteParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.VoteParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

func (query *Query) GetProposalParam() (*model.ProposalParam, error) {
	resp, err := query.transport.Query(getProposalParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.ProposalParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

func (query *Query) GetValidatorParam() (*model.ValidatorParam, error) {
	resp, err := query.transport.Query(getValidatorParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.ValidatorParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

func (query *Query) GetCoinDayParam() (*model.CoinDayParam, error) {
	resp, err := query.transport.Query(getCoinDayParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.CoinDayParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

func (query *Query) GetBandwidthParam() (*model.BandwidthParam, error) {
	resp, err := query.transport.Query(getBandwidthParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.BandwidthParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

func (query *Query) GetAccountParam() (*model.AccountParam, error) {
	resp, err := query.transport.Query(getAccountParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.AccountParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}
