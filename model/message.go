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
	Referrer             string        `json:"referrer"`
	RegisterFee          string        `json:"register_fee"`
	NewUser              string        `json:"new_username"`
	NewResetPubKey       crypto.PubKey `json:"new_reset_public_key"`
	NewTransactionPubKey crypto.PubKey `json:"new_transaction_public_key"`
	NewAppPubKey         crypto.PubKey `json:"new_app_public_key"`
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
	NewResetPubKey       crypto.PubKey `json:"new_reset_public_key"`
	NewTransactionPubKey crypto.PubKey `json:"new_transaction_public_key"`
	NewAppPubKey         crypto.PubKey `json:"new_app_public_key"`
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
	Author  string           `json:"author"`
	PostID  string           `json:"post_id"`
	Title   string           `json:"title"`
	Content string           `json:"content"`
	Links   []IDToURLMapping `json:"links"`
}

type DeletePostMsg struct {
	Author string `json:"author"`
	PostID string `json:"post_id"`
}

type DonateMsg struct {
	Username string `json:"username"`
	Amount   string `json:"amount"`
	Author   string `json:"author"`
	PostID   string `json:"post_id"`
	FromApp  string `json:"from_app"`
	Memo     string `json:"memo"`
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
type StakeInMsg struct {
	Username string `json:"username"`
	Deposit  string `json:"deposit"`
}

type StakeOutMsg struct {
	Username string `json:"username"`
	Amount   string `json:"amount"`
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

type ClaimInterestMsg struct {
	Username string `json:"username"`
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

type DeveloperUpdateMsg struct {
	Username    string `json:"username"`
	Website     string `json:"website"`
	Description string `json:"description"`
	AppMetaData string `json:"app_meta_data"`
}

type DeveloperRevokeMsg struct {
	Username string `json:"username"`
}

type GrantPermissionMsg struct {
	Username          string     `json:"username"`
	AuthorizedApp     string     `json:"authorized_app"`
	ValidityPeriodSec int64      `json:"validity_period_second"`
	GrantLevel        Permission `json:"grant_level"`
	Amount            string     `json:"amount"`
}

type RevokePermissionMsg struct {
	Username string        `json:"username"`
	PubKey   crypto.PubKey `json:"public_key"`
}

type PreAuthorizationMsg struct {
	Username          string `json:"username"`
	AuthorizedApp     string `json:"authorized_app"`
	ValidityPeriodSec int64  `json:"validity_period_second"`
	Amount            string `json:"amount"`
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

type UpgradeProtocolMsg struct {
	Creator string `json:"creator"`
	Link    string `json:"link"`
	Reason  string `json:"reason"`
}

type ChangeGlobalAllocationParamMsg struct {
	Creator   string                `json:"creator"`
	Parameter GlobalAllocationParam `json:"parameter"`
	Reason    string                `json:"reason"`
}

type ChangeInfraInternalAllocationParamMsg struct {
	Creator   string                       `json:"creator"`
	Parameter InfraInternalAllocationParam `json:"parameter"`
	Reason    string                       `json:"reason"`
}

type ChangeVoteParamMsg struct {
	Creator   string    `json:"creator"`
	Parameter VoteParam `json:"parameter"`
	Reason    string    `json:"reason"`
}

type ChangeProposalParamMsg struct {
	Creator   string        `json:"creator"`
	Parameter ProposalParam `json:"parameter"`
	Reason    string        `json:"reason"`
}

type ChangeDeveloperParamMsg struct {
	Creator   string         `json:"creator"`
	Parameter DeveloperParam `json:"parameter"`
	Reason    string         `json:"reason"`
}

type ChangeValidatorParamMsg struct {
	Creator   string         `json:"creator"`
	Parameter ValidatorParam `json:"parameter"`
	Reason    string         `json:"reason"`
}

type ChangeBandwidthParamMsg struct {
	Creator   string         `json:"creator"`
	Parameter BandwidthParam `json:"parameter"`
	Reason    string         `json:"reason"`
}

type ChangeAccountParamMsg struct {
	Creator   string       `json:"creator"`
	Parameter AccountParam `json:"parameter"`
	Reason    string       `json:"reason"`
}

type ChangePostParamMsg struct {
	Creator   string    `json:"creator"`
	Parameter PostParam `json:"parameter"`
	Reason    string    `json:"reason"`
}

type VoteProposalMsg struct {
	Voter      string `json:"voter"`
	ProposalID string `json:"proposal_id"`
	Result     bool   `json:"result"`
}
