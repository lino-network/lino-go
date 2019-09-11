package query

import (
	"context"

	"github.com/lino-network/lino-go/errors"
	"github.com/lino-network/lino/types"
	"github.com/lino-network/lino/x/validator/model"
)

// GetValidator returns validator info given a validator name from blockchain.
func (query *Query) GetValidator(ctx context.Context, username string) (*model.Validator, error) {
	resp, err := query.transport.Query(ctx, ValidatorKVStoreKey, ValidatorSubStore, []string{username})
	if err != nil {
		linoe, ok := err.(errors.Error)
		if ok && linoe.BlockChainCode() == uint32(types.CodeValidatorNotFound) {
			return nil, errors.EmptyResponse("validator is not found")
		}
		return nil, err
	}
	validator := new(model.Validator)
	if err := query.transport.Cdc.UnmarshalJSON(resp, validator); err != nil {
		return nil, err
	}
	return validator, nil
}

// GetAllValidators returns all oncall validators from blockchain.
func (query *Query) GetAllValidators(ctx context.Context) (*model.ValidatorList, error) {
	resp, err := query.transport.Query(ctx, ValidatorKVStoreKey, ValidatorListSubStore, []string{})
	if err != nil {
		return nil, err
	}

	validatorList := new(model.ValidatorList)
	if err := query.transport.Cdc.UnmarshalJSON(resp, validatorList); err != nil {
		return validatorList, err
	}
	return validatorList, nil
}
