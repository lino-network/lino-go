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
