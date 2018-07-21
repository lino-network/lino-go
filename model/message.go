// Package model includes all message, params, storage struct,
// standard transactions and consts same as on Lino blockchain.
package model

import (
	crypto "github.com/tendermint/tendermint/crypto"
)

type Msg interface{}

type Tx interface{}

//
// Account related messages
//
type RegisterMsg struct {
	Referrer              string        `json:"referrer"`
	RegisterFee           string        `json:"register_fee"`
	NewUser               string        `json:"new_username"`
	NewResetPubKey        crypto.PubKey `json:"new_reset_public_key"`
	NewTransactionPubKey  crypto.PubKey `json:"new_transaction_public_key"`
	NewMicropaymentPubKey crypto.PubKey `json:"new_micropayment_public_key"`
	NewPostPubKey         crypto.PubKey `json:"new_post_public_key"`
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
	Username              string        `json:"username"`
	NewResetPubKey        crypto.PubKey `json:"new_reset_public_key"`
	NewTransactionPubKey  crypto.PubKey `json:"new_transaction_public_key"`
	NewMicropaymentPubKey crypto.PubKey `json:"new_micropayment_public_key"`
	NewPostPubKey         crypto.PubKey `json:"new_post_public_key"`
}

type TransferMsg struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Amount   string `json:"amount"`
	Memo     string `json:"memo"`
}

type UpdateAccountMsg struct {
	Username string `json:"username"`
	JSONMeta string `json:"json_meta"`
}

//
// Post related messages
//
type CreatePostMsg struct {
	Author                  string           `json:"author"`
	PostID                  string           `json:"post_id"`
	Title                   string           `json:"title"`
	Content                 string           `json:"content"`
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

type UpdatePostMsg struct {
	Author                  string           `json:"author"`
	PostID                  string           `json:"post_id"`
	Title                   string           `json:"title"`
	Content                 string           `json:"content"`
	Links                   []IDToURLMapping `json:"links"`
	RedistributionSplitRate string           `json:"redistribution_split_rate"`
}

type DeletePostMsg struct {
	Author string `json:"author"`
	PostID string `json:"post_id"`
}

type LikeMsg struct {
	Username string `json:"username"`
	Weight   int64  `json:"weight"`
	Author   string `json:"author"`
	PostID   string `json:"post_id"`
}

type DonateMsg struct {
	Username       string `json:"username"`
	Amount         string `json:"amount"`
	Author         string `json:"author"`
	PostID         string `json:"post_id"`
	FromApp        string `json:"from_app"`
	Memo           string `json:"memo"`
	IsMicroPayment bool   `json:"is_micropayment"`
}

type ViewMsg struct {
	Username string `json:"username"`
	Author   string `json:"author"`
	PostID   string `json:"post_id"`
}

type ReportOrUpvoteMsg struct {
	Username string `json:"username"`
	Author   string `json:"author"`
	PostID   string `json:"post_id"`
	IsReport bool   `json:"is_report"`
}

//
// Validator related messages
//
type ValidatorDepositMsg struct {
	Username  string        `json:"username"`
	Deposit   string        `json:"deposit"`
	ValPubKey crypto.PubKey `json:"validator_public_key"`
	Link      string        `json:"link"`
}

type ValidatorWithdrawMsg struct {
	Username string `json:"username"`
	Amount   string `json:"amount"`
}

type ValidatorRevokeMsg struct {
	Username string `json:"username"`
}

//
// Vote related messages
//
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

//
// developer related messages
//
type DeveloperRegisterMsg struct {
	Username    string `json:"username"`
	Deposit     string `json:"deposit"`
	Website     string `json:"website"`
	Description string `json:"description"`
	AppMetaData string `json:"app_meta_data"`
}

type DeveloperRevokeMsg struct {
	Username string `json:"username"`
}

type GrantPermissionMsg struct {
	Username        string     `json:"username"`
	AuthenticateApp string     `json:"authenticate_app"`
	ValidityPeriod  int64      `json:"validity_period"`
	GrantLevel      Permission `json:"grant_level"`
	Times           int64      `json:"times"`
}

type RevokePermissionMsg struct {
	Username   string        `json:"username"`
	PubKey     crypto.PubKey `json:"public_key"`
	GrantLevel Permission    `json:"grant_level"`
}

//
// infra related messages
//
type ProviderReportMsg struct {
	Username string `json:"username"`
	Usage    int64  `json:"usage"`
}

//
// proposal related messages
//
type DeletePostContentMsg struct {
	Creator  string `json:"creator"`
	Permlink string `json:"permlink"`
	Reason   string `json:"reason"`
}

type VoteProposalMsg struct {
	Voter      string `json:"voter"`
	ProposalID string `json:"proposal_id"`
	Result     bool   `json:"result"`
}

type UpgradeProtocolMsg struct {
	Creator string `json:"creator"`
	Link    string `json:"link"`
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

type ChangePostParamMsg struct {
	Creator   string    `json:"creator"`
	Parameter PostParam `json:"parameter"`
}
