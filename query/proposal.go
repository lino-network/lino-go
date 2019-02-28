package query

import (
	"context"

	"github.com/lino-network/lino-go/model"
)

// GetOngoingProposal returns one ongoing proposal.
func (query *Query) GetOngoingProposal(ctx context.Context, proposalID string) (*model.Proposal, error) {
	resp, err := query.transport.Query(ctx, ProposalKVStoreKey, OngoingProposalSubStore, []string{proposalID})
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
func (query *Query) GetOngoingProposalList(ctx context.Context) ([]*model.Proposal, error) {
	resKVs, err := query.transport.QuerySubspace(ctx, getOngoingProposalSubstoreKey(), ProposalKVStoreKey)
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
func (query *Query) GetExpiredProposal(ctx context.Context, proposalID string) (*model.Proposal, error) {
	resp, err := query.transport.Query(ctx, ProposalKVStoreKey, ExpiredProposalSubStore, []string{proposalID})
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
func (query *Query) GetExpiredProposalList(ctx context.Context) ([]*model.Proposal, error) {
	resKVs, err := query.transport.QuerySubspace(ctx, getExpiredProposalSubstoreKey(), ProposalKVStoreKey)
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
func (query *Query) GetNextProposalID(ctx context.Context) (*model.NextProposalID, error) {
	resp, err := query.transport.Query(ctx, ProposalKVStoreKey, NextProposalIDSubStore, []string{})
	if err != nil {
		return nil, err
	}

	nextProposalID := new(model.NextProposalID)
	if err := query.transport.Cdc.UnmarshalJSON(resp, nextProposalID); err != nil {
		return nil, err
	}
	return nextProposalID, nil
}
