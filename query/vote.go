package query

import (
	"github.com/lino-network/lino-go/model"
)

//
// Vote related query
//

func (query *Query) GetDelegation(voter, delegator string) (*model.Delegation, error) {
	resp, err := query.transport.Query(getDelegationKey(voter, delegator), VoteKVStoreKey)
	if err != nil {
		return nil, err
	}
	delegation := new(model.Delegation)
	if err := query.transport.Cdc.UnmarshalJSON(resp, delegation); err != nil {
		return nil, err
	}
	return delegation, nil
}

func (query *Query) GetVoter(voterName string) (*model.Voter, error) {
	resp, err := query.transport.Query(getVoterKey(voterName), VoteKVStoreKey)
	if err != nil {
		return nil, err
	}
	voter := new(model.Voter)
	if err := query.transport.Cdc.UnmarshalJSON(resp, voter); err != nil {
		return nil, err
	}
	return voter, nil
}

func (query *Query) GetVote(proposalID, voter string) (*model.Vote, error) {
	resp, err := query.transport.Query(getVoteKey(proposalID, voter), VoteKVStoreKey)
	if err != nil {
		return nil, err
	}
	vote := new(model.Vote)
	if err := query.transport.Cdc.UnmarshalJSON(resp, vote); err != nil {
		return nil, err
	}
	return vote, nil
}

//
// Range query
//

func (query *Query) GetVoterAllDelegation(voter string) ([]*model.Delegation, error) {
	resKVs, err := query.transport.QuerySubspace(getDelegationPrefix(voter), VoteKVStoreKey)
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

func (query *Query) GetProposalAllVotes(prposalID string) ([]*model.Vote, error) {
	resKVs, err := query.transport.QuerySubspace(getVotePrefix(prposalID), VoteKVStoreKey)
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

func (query *Query) GetDelegatorAllDelegation(delegatorName string) ([]*model.Delegation, error) {
	resKVs, err := query.transport.QuerySubspace(GetDelegateePrefix(delegatorName), VoteKVStoreKey)
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

	// TODO: add delegatee name in the return slice.
	return delegations, nil
}

// TODO: Get all votes by voter.
