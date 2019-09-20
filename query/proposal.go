package query

import (
	"context"

	"github.com/lino-network/lino/x/proposal"
	"github.com/lino-network/lino/x/proposal/model"
)

// GetOngoingProposal returns one ongoing proposal.
func (query *Query) GetOngoingProposal(ctx context.Context, proposalID string) (*model.Proposal, error) {
	resp, err := query.transport.Query(ctx, ProposalKVStoreKey, proposal.QueryOngoingProposal, []string{proposalID})
	if err != nil {
		return nil, err
	}

	proposal := new(model.Proposal)
	if err := query.transport.Cdc.UnmarshalJSON(resp, proposal); err != nil {
		return nil, err
	}
	return proposal, nil
}

// GetExpiredProposal returns one expired past proposal.
func (query *Query) GetExpiredProposal(ctx context.Context, proposalID string) (*model.Proposal, error) {
	resp, err := query.transport.Query(ctx, ProposalKVStoreKey, proposal.QueryExpiredProposal, []string{proposalID})
	if err != nil {
		return nil, err
	}

	proposal := new(model.Proposal)
	if err := query.transport.Cdc.UnmarshalJSON(resp, proposal); err != nil {
		return nil, err
	}
	return proposal, nil
}
