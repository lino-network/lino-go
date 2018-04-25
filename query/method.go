package query

import (
	"encoding/json"
	"fmt"

	"github.com/lino-network/lino-go/model"
	"github.com/lino-network/lino-go/transport"
)

func GetAllValidators() error {
	transport := transport.NewTransportFromViper()
	res, err := transport.Query(GetValidatorListKey(), ValidatorKVStoreKey)
	if err != nil {
		return err
	}

	validatorList := new(model.ValidatorList)
	if err := json.Unmarshal(res, validatorList); err != nil {
		return err
	}

	// print out whole bank
	output, err := json.MarshalIndent(validatorList, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(output))

	return nil
}
