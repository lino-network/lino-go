package model

// Account related messages
type TransferToAddressMsg struct {
	Sender       string `json:"sender"`
	ReceiverName string `json:"receiver_name"`
	ReceiverAddr string `json:"receiver_addr"`
	Amount       string `json:"amount"`
	Memo         string `json:"memo"`
}

type TransferToUsernameMsg struct {
	Sender       string `json:"sender"`
	ReceiverName string `json:"receiver_name"`
	Amount       string `json:"amount"`
	Memo         string `json:"memo"`
}

type FollowMsg struct {
	Follower string `json:"follower"`
	Followee string `json:"followee"`
}

type UnfollowMsg struct {
	Follower string `json:"follower"`
	Followee string `json:"followee"`
}

type ClaimMsg struct {
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

type LikeMsg struct {
	Username string
	Weight   int64
	Author   string
	PostID   string
}

type DonateMsg struct {
	Username string
	Amount   string
	Author   string
	PostID   string
	FromApp  string
}

type ReportOrUpvoteMsg struct {
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

type ValidatorWithdrawMsg struct {
	Username string `json:"username"`
	Amount   string `json:"amount"`
}

type ValidatorRevokeMsg struct {
	Username string `json:"username"`
}

// Vote related messages
type VoteMsg struct {
	Voter      string `json:"voter"`
	ProposalID string `json:"proposal_id"`
	Result     bool   `json:"result"`
}

// type CreateProposal struct {
// 	Creator string `json:"creator"`
// 	model.ChangeParameterDescription
// }

type VoterDepositMsg struct {
	Username string `json:"username"`
	Deposit  string `json:"deposit"`
}

type VoterWithdrawMsg struct {
	Username string `json:"username"`
	Amount   string `json:"amount"`
}

type VoterRevokeMsg struct {
	Username string `json:"username"`
}

type DelegateMsg struct {
	Delegator string `json:"delegator"`
	Voter     string `json:"voter"`
	Amount    string `json:"amount"`
}

type DelegatorWithdrawMsg struct {
	Delegator string `json:"delegator"`
	Voter     string `json:"voter"`
	Amount    string `json:"amount"`
}

type RevokeDelegationMsg struct {
	Delegator string `json:"delegator"`
	Voter     string `json:"voter"`
}

// developer related messages
type DeveloperRegisterMsg struct {
	Username string `json:"username"`
	Deposit  string `json:"deposit"`
}

type DeveloperRevokeMsg struct {
	Username string `json:"username"`
}

// infra related messages
type ProviderReportMsg struct {
	Username string `json:"username"`
	Usage    int64  `json:"usage"`
}
