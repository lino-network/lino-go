package query

import (
	"github.com/lino-network/lino-go/model"
)

// GetOngoingProposal returns one ongoing proposal.
func (query *Query) GetOngoingProposal(proposalID string) (*model.Proposal, error) {
	resp, err := query.transport.Query(getOngoingProposalKey(proposalID), ProposalKVStoreKey)
	if err != nil {
		return nil, err
	}

	proposal := new(model.Proposal)
	if err := query.transport.Cdc.UnmarshalJSON(resp, proposal); err != nil {
		return nil, err
	}
	return proposal, nil
}

// GetOngoingProposalList returns all ongoing proposals
func (query *Query) GetOngoingProposalList() ([]*model.Proposal, error) {
	resKVs, err := query.transport.QuerySubspace(getOngoingProposalSubstoreKey(), ProposalKVStoreKey)
	if err != nil {
		return nil, err
	}

	var proposals []*model.Proposal
	for _, KV := range resKVs {
		proposal := new(model.Proposal)
		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, proposal); err != nil {
			return nil, err
		}
		proposals = append(proposals, proposal)
	}

	return proposals, nil
}

// GetExpiredProposal returns one expired past proposal.
func (query *Query) GetExpiredProposal(proposalID string) (*model.Proposal, error) {
	resp, err := query.transport.Query(getExpiredProposalKey(proposalID), ProposalKVStoreKey)
	if err != nil {
		return nil, err
	}

	proposal := new(model.Proposal)
	if err := query.transport.Cdc.UnmarshalJSON(resp, proposal); err != nil {
		return nil, err
	}
	return proposal, nil
}

// GetExpiredProposalList returns all expired proposals
func (query *Query) GetExpiredProposalList() ([]*model.Proposal, error) {
	resKVs, err := query.transport.QuerySubspace(getExpiredProposalSubstoreKey(), ProposalKVStoreKey)
	if err != nil {
		return nil, err
	}

	var proposals []*model.Proposal
	for _, KV := range resKVs {
		proposal := new(model.Proposal)
		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, proposal); err != nil {
			return nil, err
		}
		proposals = append(proposals, proposal)
	}

	return proposals, nil
}

// GetProposal returns proposal info of a specific proposalID.
func (query *Query) GetNextProposalID() (*model.NextProposalID, error) {
	resp, err := query.transport.Query(getNextProposalIDKey(), ProposalKVStoreKey)
	if err != nil {
		return nil, err
	}

	nextProposalID := new(model.NextProposalID)
	if err := query.transport.Cdc.UnmarshalJSON(resp, nextProposalID); err != nil {
		return nil, err
	}
	return nextProposalID, nil
}
