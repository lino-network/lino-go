package model

import (
	crypto "github.com/tendermint/go-crypto"
)

type Msg interface{}

// Account related messages
type RegisterMsg struct {
	NewUser              string        `json:"new_user"`
	NewMasterPubKey      crypto.PubKey `json:"new_master_public_key"`
	NewPostPubKey        crypto.PubKey `json:"new_post_public_key"`
	NewTransactionPubKey crypto.PubKey `json:"new_transaction_public_key"`
}

type TransferMsg struct {
	Sender       string `json:"sender"`
	ReceiverName string `json:"receiver_name"`
	ReceiverAddr string `json:"receiver_addr"`
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

type RecoverMsg struct {
	Username             string        `json:"username"`
	NewPostPubKey        crypto.PubKey `json:"new_post_public_key"`
	NewTransactionPubKey crypto.PubKey `json:"new_transaction_public_key"`
}

type SavingToCheckingMsg struct {
	Username string `json:"username"`
	Amount   string `json:"amount"`
}

type CheckingToSavingMsg struct {
	Username string `json:"username"`
	Amount   string `json:"amount"`
}

// Post related messages
type CreatePostMsg struct {
	PostCreateParams
}

type PostCreateParams struct {
	PostID                  string           `json:"post_id"`
	Title                   string           `json:"title"`
	Content                 string           `json:"content"`
	Author                  string           `json:"author"`
	ParentAuthor            string           `json:"parent_author"`
	ParentPostID            string           `json:"parent_postID"`
	SourceAuthor            string           `json:"source_author"`
	SourcePostID            string           `json:"source_postID"`
	Links                   []IDToURLMapping `json:"links"`
	RedistributionSplitRate string           `json:"redistribution_split_rate"`
}

type IDToURLMapping struct {
	Identifier string `json:"identifier"`
	URL        string `json:"url"`
}

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
type ValidatorDepositMsg struct {
	Username  string        `json:"username"`
	Deposit   string        `json:"deposit"`
	ValPubKey crypto.PubKey `json:"validator_public_key"`
}

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

type GrantDeveloperMsg struct {
	Username        string `json:"username"`
	AuthenticateApp string `json:"authenticate_app"`
	ValidityPeriod  int64  `json:"validity_period"`
	GrantLevel      int64  `json:"grant_level"`
}

// infra related messages
type ProviderReportMsg struct {
	Username string `json:"username"`
	Usage    int64  `json:"usage"`
}

// proposal related messages
type DeletePostContentMsg struct {
	Creator  string `json:"creator"`
	PermLink string `json:"permLink"`
}

type ChangeGlobalAllocationParamMsg struct {
	Creator   string                `json:"creator"`
	Parameter GlobalAllocationParam `json:"parameter"`
}

type ChangeEvaluateOfContentValueParamMsg struct {
	Creator   string                      `json:"creator"`
	Parameter EvaluateOfContentValueParam `json:"parameter"`
}

type ChangeInfraInternalAllocationParamMsg struct {
	Creator   string                       `json:"creator"`
	Parameter InfraInternalAllocationParam `json:"parameter"`
}

type ChangeVoteParamMsg struct {
	Creator   string    `json:"creator"`
	Parameter VoteParam `json:"parameter"`
}

type ChangeProposalParamMsg struct {
	Creator   string        `json:"creator"`
	Parameter ProposalParam `json:"parameter"`
}

type ChangeDeveloperParamMsg struct {
	Creator   string         `json:"creator"`
	Parameter DeveloperParam `json:"parameter"`
}

type ChangeValidatorParamMsg struct {
	Creator   string         `json:"creator"`
	Parameter ValidatorParam `json:"parameter"`
}

type ChangeCoinDayParamMsg struct {
	Creator   string       `json:"creator"`
	Parameter CoinDayParam `json:"parameter"`
}

type ChangeBandwidthParamMsg struct {
	Creator   string         `json:"creator"`
	Parameter BandwidthParam `json:"parameter"`
}

type ChangeAccountParamMsg struct {
	Creator   string       `json:"creator"`
	Parameter AccountParam `json:"parameter"`
}
