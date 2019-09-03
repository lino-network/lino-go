package query

import (
	"context"

	"github.com/lino-network/lino-go/errors"
	"github.com/lino-network/lino/x/infra/model"
)

// GetInfraProvider returns the infra provider info such as usage.
func (query *Query) GetInfraProvider(ctx context.Context, providerName string) (*model.InfraProvider, error) {
	resp, err := query.transport.Query(ctx, InfraKVStoreKey, InfraProviderSubStore, []string{providerName})
	if err != nil {
		linoe, ok := err.(errors.Error)
		if ok && linoe.BlockChainCode() == uint32(errors.CodeInfraProviderNotFound) {
			return nil, errors.EmptyResponse("developer is not found")
		}
		return nil, err
	}
	provider := new(model.InfraProvider)
	if err := query.transport.Cdc.UnmarshalJSON(resp, provider); err != nil {
		return nil, err
	}
	return provider, nil
}

// GetInfraProviders returns a list of all infra providers.
func (query *Query) GetInfraProviders(ctx context.Context) (*model.InfraProviderList, error) {
	resp, err := query.transport.Query(ctx, InfraKVStoreKey, InfraListSubStore, []string{})
	if err != nil {
		return nil, err
	}

	providerList := new(model.InfraProviderList)
	if err := query.transport.Cdc.UnmarshalJSON(resp, providerList); err != nil {
		return nil, err
	}
	return providerList, nil
}
