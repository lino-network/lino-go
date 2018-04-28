package query

import (
	"encoding/json"

	"github.com/lino-network/lino-go/model"
	"github.com/lino-network/lino-go/transport"
)

func GetAllValidators() (*model.ValidatorList, error) {
	transport := transport.NewTransportFromViper()
	resp, err := transport.Query(getValidatorListKey(), ValidatorKVStoreKey)
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
	resp, err := transport.Query(getValidatorKey(username), ValidatorKVStoreKey)
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
	resp, err := transport.Query(getDelegationKey(voter, delegator), VoteKVStoreKey)
	if err != nil {
		return nil, err
	}
	delegation := new(model.Delegation)
	if err := json.Unmarshal(resp, delegation); err != nil {
		return nil, err
	}
	return delegation, nil
}

func GetVoter(voterName string) (*model.Voter, error) {
	transport := transport.NewTransportFromViper()
	resp, err := transport.Query(getVoterKey(voterName), VoteKVStoreKey)
	if err != nil {
		return nil, err
	}
	voter := new(model.Voter)
	if err := json.Unmarshal(resp, voter); err != nil {
		return nil, err
	}
	return voter, nil
}

func GetInfraProvider(providerName string) (*model.InfraProvider, error) {
	transport := transport.NewTransportFromViper()
	resp, err := transport.Query(getInfraProviderKey(providerName), InfraKVStoreKey)
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
	resp, err := transport.Query(getInfraProviderListKey(), InfraKVStoreKey)
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
	resp, err := transport.Query(getDeveloperKey(developerName), DeveloperKVStoreKey)
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
	resp, err := transport.Query(getDeveloperListKey(), DeveloperKVStoreKey)
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
	resp, err := transport.Query(getAccountMetaKey(username), AccountKVStoreKey)
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
	resp, err := transport.Query(getAccountBankKey(address), AccountKVStoreKey)
	if err != nil {
		return nil, err
	}
	bank := new(model.AccountBank)
	if err := json.Unmarshal(resp, bank); err != nil {
		panic(err)
		return nil, err
	}
	return bank, nil
}
