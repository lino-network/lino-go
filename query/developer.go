package query

import (
	"context"

	"github.com/lino-network/lino-go/model"
)

// GetDeveloper returns a specific developer info from blockchain.
func (query *Query) GetDeveloper(ctx context.Context, developerName string) (*model.Developer, error) {
	resp, err := query.transport.Query(ctx, DeveloperKVStoreKey, DeveloperSubStore, []string{developerName})
	if err != nil {
		return nil, err
	}
	developer := new(model.Developer)
	if err := query.transport.Cdc.UnmarshalJSON(resp, developer); err != nil {
		return nil, err
	}
	return developer, nil
}

// GetDevelopers returns a list of all developers.
func (query *Query) GetDevelopers(ctx context.Context) (*model.DeveloperList, error) {
	resp, err := query.transport.Query(ctx, DeveloperKVStoreKey, DeveloperListSubStore, []string{})
	if err != nil {
		return nil, err
	}

	developerList := new(model.DeveloperList)
	if err := query.transport.Cdc.UnmarshalJSON(resp, developerList); err != nil {
		return nil, err
	}
	return developerList, nil
}
