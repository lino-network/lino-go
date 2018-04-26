package model

import (
	"github.com/lino-network/lino/types"
)

// Account related messages
type TransferToAddress struct {
	Sender       string `json:"sender"`
	ReceiverAddr string `json:"receiver_addr"`
	Amount       string `json:"amount"`
	Memo         string `json:"memo"`
}

type TransferToUsername struct {
	Sender       string `json:"sender"`
	ReceiverName string `json:"receiver_name"`
	Amount       string `json:"amount"`
	Memo         string `json:"memo"`
}

type Follow struct {
	Follower string `json:"follower"`
	Followee string `json:"followee"`
}

type Unfollow struct {
	Follower string `json:"follower"`
	Followee string `json:"followee"`
}

type Claim struct {
	Username string `json:"username"`
}

// Post related messages
// type CreatePost struct {
// 	PostID                  string                 `json:"post_id"`
// 	Title                   string                 `json:"title"`
// 	Content                 string                 `json:"content"`
// 	Author                  string                 `json:"author"`
// 	ParentAuthor            string                 `json:"parent_author"`
// 	ParentPostID            string                 `json:"parent_postID"`
// 	SourceAuthor            string                 `json:"source_author"`
// 	SourcePostID            string                 `json:"source_postID"`
// 	Links                   []types.IDToURLMapping `json:"links"`
// 	RedistributionSplitRate sdk.Rat                `json:"redistribution_split_rate"`
// }

type Like struct {
	Username string
	Weight   int64
	Author   string
	PostID   string
}

type Donate struct {
	Username string
	Amount   types.LNO
	Author   string
	PostID   string
	FromApp  string
}

type ReportOrUpvote struct {
	Username string
	Author   string
	PostID   string
	IsReport bool
	IsRevoke bool
}

// Validator related messages
// type ValidatorDeposit struct {
// 	Username  string        `json:"username"`
// 	Deposit   string        `json:"deposit"`
// 	ValPubKey crypto.PubKey `json:"validator_public_key"`
// }

type ValidatorWithdraw struct {
	Username string `json:"username"`
	Amount   string `json:"amount"`
}

type ValidatorRevoke struct {
	Username string `json:"username"`
}

// Vote related messages
type Vote struct {
	Voter      string            `json:"voter"`
	ProposalID types.ProposalKey `json:"proposal_id"`
	Result     bool              `json:"result"`
}

// type CreateProposal struct {
// 	Creator string `json:"creator"`
// 	model.ChangeParameterDescription
// }

type VoterDeposit struct {
	Username string `json:"username"`
	Deposit  string `json:"deposit"`
}

type VoterWithdraw struct {
	Username string `json:"username"`
	Amount   string `json:"amount"`
}

type VoterRevoke struct {
	Username string `json:"username"`
}

type Delegate struct {
	Delegator string `json:"delegator"`
	Voter     string `json:"voter"`
	Amount    string `json:"amount"`
}

type DelegatorWithdraw struct {
	Delegator string `json:"delegator"`
	Voter     string `json:"voter"`
	Amount    string `json:"amount"`
}

type RevokeDelegation struct {
	Delegator string `json:"delegator"`
	Voter     string `json:"voter"`
}

// developer related messages
type DeveloperRegister struct {
	Username string `json:"username"`
	Deposit  string `json:"deposit"`
}

type DeveloperRevoke struct {
	Username string `json:"username"`
}

// infra related messages
type ProviderReport struct {
	Username string `json:"username"`
	Usage    int64  `json:"usage"`
}
