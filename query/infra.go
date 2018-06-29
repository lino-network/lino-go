package query

import (
	"github.com/lino-network/lino-go/model"
)

// GetInfraProvider returns the infra provider info such as usage.
func (query *Query) GetInfraProvider(providerName string) (*model.InfraProvider, error) {
	resp, err := query.transport.Query(getInfraProviderKey(providerName), InfraKVStoreKey)
	if err != nil {
		return nil, err
	}
	provider := new(model.InfraProvider)
	if err := query.transport.Cdc.UnmarshalJSON(resp, provider); err != nil {
		return nil, err
	}
	return provider, nil
}

// GetInfraProviders returns a list of all infra providers.
func (query *Query) GetInfraProviders() (*model.InfraProviderList, error) {
	resp, err := query.transport.Query(getInfraProviderListKey(), InfraKVStoreKey)
	if err != nil {
		return nil, err
	}

	providerList := new(model.InfraProviderList)
	if err := query.transport.Cdc.UnmarshalJSON(resp, providerList); err != nil {
		return nil, err
	}
	return providerList, nil
}
