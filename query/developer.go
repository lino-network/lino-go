package query

import (
	"github.com/lino-network/lino-go/model"
)

//
// Developer related query
//
func (query *Query) GetDeveloper(developerName string) (*model.Developer, error) {
	resp, err := query.transport.Query(getDeveloperKey(developerName), DeveloperKVStoreKey)
	if err != nil {
		return nil, err
	}
	developer := new(model.Developer)
	if err := query.transport.Cdc.UnmarshalJSON(resp, developer); err != nil {
		return nil, err
	}
	return developer, nil
}

func (query *Query) GetDevelopers() (*model.DeveloperList, error) {
	resp, err := query.transport.Query(getDeveloperListKey(), DeveloperKVStoreKey)
	if err != nil {
		return nil, err
	}

	developerList := new(model.DeveloperList)
	if err := query.transport.Cdc.UnmarshalJSON(resp, developerList); err != nil {
		return nil, err
	}
	return developerList, nil
}
