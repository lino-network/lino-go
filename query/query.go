package query

import (
	"github.com/lino-network/lino-go/model"
	"github.com/lino-network/lino-go/transport"
)

// Account related query
func IsPrivKeyMatchUsername(username string, privKey string) (bool, error) {
	return true, nil
}

func GetAccountSequence(username string) int64 {
	meta, err := GetAccountMeta(username)
	if err != nil {
		return 0
	}
	return meta.Sequence
}

func GetAccountMeta(username string) (*model.AccountMeta, error) {
	transport := transport.NewTransportFromViper()
	resp, err := transport.Query(getAccountMetaKey(username), AccountKVStoreKey)
	if err != nil {
		return nil, err
	}
	meta := new(model.AccountMeta)
	if err := transport.Cdc.UnmarshalJSON(resp, meta); err != nil {
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
	if err := transport.Cdc.UnmarshalJSON(resp, bank); err != nil {
		return nil, err
	}
	return bank, nil
}

func GetAccountInfo(username string) (*model.AccountInfo, error) {
	transport := transport.NewTransportFromViper()
	resp, err := transport.Query(getAccountInfoKey(username), AccountKVStoreKey)
	if err != nil {
		return nil, err
	}
	info := new(model.AccountInfo)
	if err := transport.Cdc.UnmarshalJSON(resp, info); err != nil {
		return nil, err
	}
	return info, nil
}

// Post related query

// Validator related query
func GetValidator(username string) (*model.Validator, error) {
	transport := transport.NewTransportFromViper()
	resp, err := transport.Query(getValidatorKey(username), ValidatorKVStoreKey)
	if err != nil {
		return nil, err
	}
	validator := new(model.Validator)
	if err := transport.Cdc.UnmarshalJSON(resp, validator); err != nil {
		return nil, err
	}
	return validator, nil
}

func GetAllValidators() (*model.ValidatorList, error) {
	transport := transport.NewTransportFromViper()
	resp, err := transport.Query(getValidatorListKey(), ValidatorKVStoreKey)
	if err != nil {
		return nil, err
	}

	validatorList := new(model.ValidatorList)
	if err := transport.Cdc.UnmarshalJSON(resp, validatorList); err != nil {
		return validatorList, err
	}
	return validatorList, nil
}

// Vote related query
func GetDelegation(voter string, delegator string) (*model.Delegation, error) {
	transport := transport.NewTransportFromViper()
	resp, err := transport.Query(getDelegationKey(voter, delegator), VoteKVStoreKey)
	if err != nil {
		return nil, err
	}
	delegation := new(model.Delegation)
	if err := transport.Cdc.UnmarshalJSON(resp, delegation); err != nil {
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
	if err := transport.Cdc.UnmarshalJSON(resp, voter); err != nil {
		return nil, err
	}
	return voter, nil
}

// Developer related query
func GetDeveloper(developerName string) (*model.Developer, error) {
	transport := transport.NewTransportFromViper()
	resp, err := transport.Query(getDeveloperKey(developerName), DeveloperKVStoreKey)
	if err != nil {
		return nil, err
	}
	developer := new(model.Developer)
	if err := transport.Cdc.UnmarshalJSON(resp, developer); err != nil {
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
	if err := transport.Cdc.UnmarshalJSON(resp, developerList); err != nil {
		return nil, err
	}
	return developerList, nil
}

// Infra related query
func GetInfraProvider(providerName string) (*model.InfraProvider, error) {
	transport := transport.NewTransportFromViper()
	resp, err := transport.Query(getInfraProviderKey(providerName), InfraKVStoreKey)
	if err != nil {
		return nil, err
	}
	provider := new(model.InfraProvider)
	if err := transport.Cdc.UnmarshalJSON(resp, provider); err != nil {
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
	if err := transport.Cdc.UnmarshalJSON(resp, providerList); err != nil {
		return nil, err
	}
	return providerList, nil
}
