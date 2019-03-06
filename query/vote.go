package query

import (
	"context"

	"github.com/lino-network/lino-go/errors"
	"github.com/lino-network/lino-go/model"
)

// GetDelegation returns the delegation relationship between
// a voter and a delegator from blockchain.
func (query *Query) GetDelegation(ctx context.Context, voter, delegator string) (*model.Delegation, error) {
	resp, err := query.transport.Query(ctx, VoteKVStoreKey, ValidatorListSubStore, []string{voter, delegator})
	if err != nil {
		return nil, err
	}
	delegation := new(model.Delegation)
	if err := query.transport.Cdc.UnmarshalJSON(resp, delegation); err != nil {
		return nil, err
	}
	return delegation, nil
}

// GetVoterAllDelegation returns all delegations that are delegated to a voter.
func (query *Query) GetVoterAllDelegation(ctx context.Context, voter string) ([]*model.Delegation, error) {
	resKVs, err := query.transport.QuerySubspace(ctx, getDelegationPrefix(voter), VoteKVStoreKey)
	if err != nil {
		return nil, err
	}

	var delegations []*model.Delegation
	for _, KV := range resKVs {
		delegation := new(model.Delegation)
		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, delegation); err != nil {
			return nil, err
		}
		delegations = append(delegations, delegation)
	}

	return delegations, nil
}

// GetDelegatorAllDelegation returns all delegations that a delegator has delegated to.
func (query *Query) GetDelegatorAllDelegation(ctx context.Context, delegatorName string) (map[string]*model.Delegation, error) {
	resKVs, err := query.transport.QuerySubspace(ctx, getDelegateePrefix(delegatorName), VoteKVStoreKey)
	if err != nil {
		return nil, err
	}

	delegateeToDelegations := make(map[string]*model.Delegation)
	for _, KV := range resKVs {
		delegation := new(model.Delegation)
		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, delegation); err != nil {
			return nil, err
		}
		delegateeToDelegations[getSubstringAfterKeySeparator(KV.Key)] = delegation
	}

	return delegateeToDelegations, nil
}

// GetVoter returns voter info given a voter name from blockchain.
func (query *Query) GetVoter(ctx context.Context, voterName string) (*model.Voter, error) {
	resp, err := query.transport.Query(ctx, VoteKVStoreKey, VoterSubStore, []string{voterName})
	if err != nil {
		linoe, ok := err.(errors.Error)
		if ok && linoe.BlockChainCode() == uint32(errors.CodeVoterNotFound) {
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

// GetVote returns a vote performed by a voter for a given proposal.
func (query *Query) GetVote(ctx context.Context, proposalID, voter string) (*model.Vote, error) {
	resp, err := query.transport.Query(ctx, VoteKVStoreKey, VoteSubStore, []string{proposalID, voter})
	if err != nil {
		linoe, ok := err.(errors.Error)
		if ok && linoe.BlockChainCode() == uint32(errors.CodeVoteNotFound) {
			return nil, errors.EmptyResponse("voter is not found")
		}
		return nil, err
	}
	vote := new(model.Vote)
	if err := query.transport.Cdc.UnmarshalJSON(resp, vote); err != nil {
		return nil, err
	}
	return vote, nil
}

// GetProposalAllVotes returns all votes of a given proposal.
func (query *Query) GetProposalAllVotes(ctx context.Context, prposalID string) ([]*model.Vote, error) {
	resKVs, err := query.transport.QuerySubspace(ctx, getVotePrefix(prposalID), VoteKVStoreKey)
	if err != nil {
		return nil, err
	}

	var votes []*model.Vote
	for _, KV := range resKVs {
		vote := new(model.Vote)
		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, vote); err != nil {
			return nil, err
		}
		votes = append(votes, vote)
	}

	return votes, nil
}
