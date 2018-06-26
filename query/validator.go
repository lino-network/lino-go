package query

import (
	"github.com/lino-network/lino-go/model"
)

//
// Validator related query
//
func (query *Query) GetValidator(username string) (*model.Validator, error) {
	resp, err := query.transport.Query(getValidatorKey(username), ValidatorKVStoreKey)
	if err != nil {
		return nil, err
	}
	validator := new(model.Validator)
	if err := query.transport.Cdc.UnmarshalJSON(resp, validator); err != nil {
		return nil, err
	}
	return validator, nil
}

func (query *Query) GetAllValidators() (*model.ValidatorList, error) {
	resp, err := query.transport.Query(getValidatorListKey(), ValidatorKVStoreKey)
	if err != nil {
		return nil, err
	}

	validatorList := new(model.ValidatorList)
	if err := query.transport.Cdc.UnmarshalJSON(resp, validatorList); err != nil {
		return validatorList, err
	}
	return validatorList, nil
}
