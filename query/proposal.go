package query

import (
	"context"

	"github.com/lino-network/lino-go/model"
)

// GetProposalList returns a list of all proposals, including onging
// proposals and past ones.
func (query *Query) GetProposalList(ctx context.Context) (*model.ProposalList, error) {
	resp, err := query.transport.Query(ctx, getProposalListKey(), ProposalKVStoreKey)
	if err != nil {
		return nil, err
	}

	proposalList := new(model.ProposalList)
	if err := query.transport.Cdc.UnmarshalJSON(resp, proposalList); err != nil {
		return nil, err
	}
	return proposalList, nil
}

// GetProposal returns proposal info of a specific proposalID.
func (query *Query) GetProposal(ctx context.Context, proposalID string) (*model.Proposal, error) {
	resp, err := query.transport.Query(ctx, getProposalKey(proposalID), ProposalKVStoreKey)
	if err != nil {
		return nil, err
	}

	proposal := new(model.Proposal)
	if err := query.transport.Cdc.UnmarshalJSON(resp, proposal); err != nil {
		return nil, err
	}
	return proposal, nil
}

// GetOngoingProposal returns all ongoing proposals.
func (query *Query) GetOngoingProposal(ctx context.Context) ([]*model.Proposal, error) {
	proposalList, err := query.GetProposalList(ctx)
	if err != nil {
		return nil, err
	}

	var ongoingProposals []*model.Proposal
	for _, proposalID := range proposalList.OngoingProposal {
		p, err := query.GetProposal(ctx, proposalID)
		if err != nil {
			return nil, err
		}

		ongoingProposals = append(ongoingProposals, p)
	}

	return ongoingProposals, nil
}

// GetExpiredProposal returns all past proposals.
func (query *Query) GetExpiredProposal(ctx context.Context) ([]*model.Proposal, error) {
	proposalList, err := query.GetProposalList(ctx)
	if err != nil {
		return nil, err
	}

	var expiredProposals []*model.Proposal
	for _, proposalID := range proposalList.PastProposal {
		p, err := query.GetProposal(ctx, proposalID)
		if err != nil {
			return nil, err
		}

		expiredProposals = append(expiredProposals, p)
	}

	return expiredProposals, nil
}

// GetProposal returns proposal info of a specific proposalID.
func (query *Query) GetNextProposalID(ctx context.Context) (*model.NextProposalID, error) {
	resp, err := query.transport.Query(ctx, GetNextProposalIDKey(), ProposalKVStoreKey)
	if err != nil {
		return nil, err
	}

	nextProposalID := new(model.NextProposalID)
	if err := query.transport.Cdc.UnmarshalJSON(resp, nextProposalID); err != nil {
		return nil, err
	}
	return nextProposalID, nil
}
