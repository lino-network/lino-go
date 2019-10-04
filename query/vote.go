package query

import (
	"context"

	"github.com/lino-network/lino-go/errors"
	linotypes "github.com/lino-network/lino/types"
	"github.com/lino-network/lino/x/vote/model"
	vote "github.com/lino-network/lino/x/vote/types"
)

// GetDelegation returns the delegation relationship between
// a voter and a delegator from blockchain.
// func (query *Query) GetDelegation(ctx context.Context, voter, delegator string) (*model.Delegation, error) {
// 	resp, err := query.transport.Query(ctx, VoteKVStoreKey, vote.QueryVoter, []string{voter, delegator})
// 	if err != nil {
// 		return nil, err
// 	}
// 	delegation := new(model.Delegation)
// 	if err := query.transport.Cdc.UnmarshalJSON(resp, delegation); err != nil {
// 		return nil, err
// 	}
// 	return delegation, nil
// }

// GetVoter returns voter info given a voter name from blockchain.
func (query *Query) GetVoter(ctx context.Context, voterName string) (*model.Voter, error) {
	resp, err := query.transport.Query(ctx, VoteKVStoreKey, vote.QueryVoter, []string{voterName})
	if err != nil {
		linoe, ok := err.(errors.Error)
		if ok && linoe.BlockChainCode() == uint32(linotypes.CodeVoterNotFound) {
			return nil, errors.EmptyResponse("voter is not found")
		}
		return nil, err
	}
	voter := new(model.Voter)
	if err := query.transport.Cdc.UnmarshalJSON(resp, voter); err != nil {
		return nil, err
	}
	return voter, nil
}
