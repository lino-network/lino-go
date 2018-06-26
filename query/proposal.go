package query

import (
	"github.com/lino-network/lino-go/model"
)

//
// proposal related query
//
func (query *Query) GetProposalList() (*model.ProposalList, error) {
	resp, err := query.transport.Query(getProposalListKey(), ProposalKVStoreKey)
	if err != nil {
		return nil, err
	}

	proposalList := new(model.ProposalList)
	if err := query.transport.Cdc.UnmarshalJSON(resp, proposalList); err != nil {
		return nil, err
	}
	return proposalList, nil
}

func (query *Query) GetProposal(proposalID string) (*model.Proposal, error) {
	resp, err := query.transport.Query(getProposalKey(proposalID), ProposalKVStoreKey)
	if err != nil {
		return nil, err
	}

	proposal := new(model.Proposal)
	if err := query.transport.Cdc.UnmarshalJSON(resp, proposal); err != nil {
		return nil, err
	}
	return proposal, nil
}

func (query *Query) GetOngoingProposal() ([]*model.Proposal, error) {
	proposalList, err := query.GetProposalList()
	if err != nil {
		return nil, err
	}

	var ongoingProposals []*model.Proposal
	for _, proposalID := range proposalList.OngoingProposal {
		p, err := query.GetProposal(proposalID)
		if err != nil {
			return nil, err
		}

		ongoingProposals = append(ongoingProposals, p)
	}

	return ongoingProposals, nil
}

func (query *Query) GetExpiredProposal() ([]*model.Proposal, error) {
	proposalList, err := query.GetProposalList()
	if err != nil {
		return nil, err
	}

	var expiredProposals []*model.Proposal
	for _, proposalID := range proposalList.PastProposal {
		p, err := query.GetProposal(proposalID)
		if err != nil {
			return nil, err
		}

		expiredProposals = append(expiredProposals, p)
	}

	return expiredProposals, nil
}
