package query

import (
	"encoding/json"

	"github.com/lino-network/lino-go/model"
	"github.com/lino-network/lino-go/transport"
)

func GetAllValidators() (*model.ValidatorList, error) {
	transport := transport.NewTransportFromViper()
	resp, err := transport.Query(GetValidatorListKey(), ValidatorKVStoreKey)
	if err != nil {
		return nil, err
	}

	validatorList := new(model.ValidatorList)
	if err := json.Unmarshal(resp, validatorList); err != nil {
		return validatorList, err
	}
	return validatorList, nil
}

func GetValidator(username string) (*model.Validator, error) {
	transport := transport.NewTransportFromViper()
	resp, err := transport.Query(GetValidatorKey(username), ValidatorKVStoreKey)
	if err != nil {
		return nil, err
	}
	validator := new(model.Validator)
	if err := json.Unmarshal(resp, validator); err != nil {
		return nil, err
	}
	return validator, nil
}

func GetDelegation(voter string, delegator string) (*model.Delegation, error) {
	transport := transport.NewTransportFromViper()
	resp, err := transport.Query(GetDelegationKey(voter, delegator), VoteKVStoreKey)
	if err != nil {
		return nil, err
	}
	delegation := new(model.Delegation)
	if err := json.Unmarshal(resp, delegation); err != nil {
		return nil, err
	}
	return delegation, nil
}

func GetInfraProvider(providerName string) (*model.InfraProvider, error) {
	transport := transport.NewTransportFromViper()
	resp, err := transport.Query(GetInfraProviderKey(providerName), InfraKVStoreKey)
	if err != nil {
		return nil, err
	}
	provider := new(model.InfraProvider)
	if err := json.Unmarshal(resp, provider); err != nil {
		return nil, err
	}
	return provider, nil
}

func GetInfraProviders() (*model.InfraProviderList, error) {
	transport := transport.NewTransportFromViper()
	resp, err := transport.Query(GetInfraProviderListKey(), InfraKVStoreKey)
	if err != nil {
		return nil, err
	}

	providerList := new(model.InfraProviderList)
	if err := json.Unmarshal(resp, providerList); err != nil {
		return nil, err
	}
	return providerList, nil
}

func GetDeveloper(developerName string) (*model.Developer, error) {
	transport := transport.NewTransportFromViper()
	resp, err := transport.Query(GetDeveloperKey(developerName), DeveloperKVStoreKey)
	if err != nil {
		return nil, err
	}
	developer := new(model.Developer)
	if err := json.Unmarshal(resp, developer); err != nil {
		return nil, err
	}
	return developer, nil
}

func GetDevelopers() (*model.DeveloperList, error) {
	transport := transport.NewTransportFromViper()
	resp, err := transport.Query(GetDeveloperListKey(), DeveloperKVStoreKey)
	if err != nil {
		return nil, err
	}

	developerList := new(model.DeveloperList)
	if err := json.Unmarshal(resp, developerList); err != nil {
		return nil, err
	}
	return developerList, nil
}

func GetAccountSequence(username string) (int64, error) {
	meta, err := GetAccountMeta(username)
	if err != nil {
		return -1, err
	}
	return meta.Sequence, nil
}

func GetAccountMeta(username string) (*model.AccountMeta, error) {
	transport := transport.NewTransportFromViper()
	resp, err := transport.Query(GetAccountMetaKey(username), AccountKVStoreKey)
	if err != nil {
		return nil, err
	}
	meta := new(model.AccountMeta)
	if err := json.Unmarshal(resp, meta); err != nil {
		return nil, err
	}
	return meta, nil
}

func GetAccountBank(address string) (*model.AccountBank, error) {
	transport := transport.NewTransportFromViper()
	resp, err := transport.Query(GetAccountBankKey(address), AccountKVStoreKey)
	if err != nil {
		return nil, err
	}
	bank := new(model.AccountBank)
	if err := json.Unmarshal(resp, bank); err != nil {
		return nil, err
	}
	return bank, nil
}
