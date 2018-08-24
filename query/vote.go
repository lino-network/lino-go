package query

import (
	"github.com/lino-network/lino-go/model"
)

// GetDelegation returns the delegation relationship between
// a voter and a delegator from blockchain.
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

// GetVoterAllDelegation returns all delegations that are delegated to a voter.
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

// GetDelegatorAllDelegation returns all delegations that a delegator has delegated to.
func (query *Query) GetDelegatorAllDelegation(delegatorName string) (map[string]*model.Delegation, error) {
	resKVs, err := query.transport.QuerySubspace(getDelegateePrefix(delegatorName), VoteKVStoreKey)
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

// GetVote returns a vote performed by a voter for a given proposal.
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

// GetProposalAllVotes returns all votes of a given proposal.
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

// TODO: Get all votes by voter.
