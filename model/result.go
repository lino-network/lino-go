package model

import (
	"github.com/tendermint/go-crypto"
)

type Coin struct {
	Amount int64 `json:"amount"`
}

type Rat struct {
	Num   int64 `json:"num"`
	Denom int64 `json:"denom"`
}

type ABCIValidator struct {
	PubKey []byte `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	Power  int64  `protobuf:"varint,2,opt,name=power,proto3" json:"power,omitempty"`
}

// validator related struct
type ValidatorList struct {
	OncallValidators   []string `json:"oncall_validators"`
	AllValidators      []string `json:"all_validators"`
	PreBlockValidators []string `json:"pre_block_validators"`
	LowestPower        Coin     `json:"lowest_power"`
	LowestValidator    string   `json:"lowest_validator"`
}

type Validator struct {
	ABCIValidator
	Username       string `json:"username"`
	Deposit        Coin   `json:"deposit"`
	AbsentCommit   int    `json:"absent_commit"`
	ProducedBlocks int64  `json:"produced_blocks"`
	Link           string `json:"link"`
}

// vote related struct
type Voter struct {
	Username       string `json:"username"`
	Deposit        Coin   `json:"deposit"`
	DelegatedPower Coin   `json:"delegated_power"`
}

type Vote struct {
	Voter  string `json:"voter"`
	Result bool   `json:"result"`
}

type Delegation struct {
	Delegator string `json:"delegator"`
	Amount    Coin   `json:"amount"`
}

// post related
type Comment struct {
	Author  string `json:"author"`
	PostID  string `json:"post_key"`
	Created int64  `json:"created"`
}

type View struct {
	Username string `json:"username"`
	Created  int64  `json:"created"`
	Times    int64  `jons:"times"`
}

type Like struct {
	Username string `json:"username"`
	Weight   int64  `json:"weight"`
	Created  int64  `json:"created"`
}

type Donation struct {
	Amount  Coin  `json:"amount"`
	Created int64 `json:"created"`
}

type ReportOrUpvote struct {
	Username string `json:"username"`
	Stake    Coin   `json:"stake"`
	Created  int64  `json:"created"`
	IsReport bool   `json:"is_report"`
}

type PostInfo struct {
	PostID       string           `json:"post_id"`
	Title        string           `json:"title"`
	Content      string           `json:"content"`
	Author       string           `json:"author"`
	ParentAuthor string           `json:"parent_author"`
	ParentPostID string           `json:"parent_postID"`
	SourceAuthor string           `json:"source_author"`
	SourcePostID string           `json:"source_postID"`
	Links        []IDToURLMapping `json:"links"`
}

type PostMeta struct {
	Created                 int64 `json:"created"`
	LastUpdate              int64 `json:"last_update"`
	LastActivity            int64 `json:"last_activity"`
	AllowReplies            bool  `json:"allow_replies"`
	TotalLikeCount          int64 `json:"total_like_count"`
	TotalDonateCount        int64 `json:"total_donate_count"`
	TotalLikeWeight         int64 `json:"total_like_weight"`
	TotalDislikeWeight      int64 `json:"total_dislike_weight"`
	TotalReportStake        Coin  `json:"total_report_stake"`
	TotalUpvoteStake        Coin  `json:"total_upvote_stake"`
	TotalReward             Coin  `json:"reward"`
	PenaltyScore            Rat   `json:"penalty_score"`
	RedistributionSplitRate Rat   `json:"redistribution_split_rate"`
}

// developer related
type Developer struct {
	Username       string `json:"username"`
	Deposit        Coin   `json:"deposit"`
	AppConsumption Coin   `json:"app_consumption"`
}

type DeveloperList struct {
	AllDevelopers []string `json:"all_developers"`
}

// infra provider related

type InfraProvider struct {
	Username string `json:"username"`
	Usage    int64  `json:"usage"`
}

type InfraProviderList struct {
	AllInfraProviders []string `json:"all_infra_providers"`
}

// account related
type AccountMeta struct {
	Sequence            int64 `json:"sequence"`
	LastActivity        int64 `json:"last_activity"`
	TransactionCapacity Coin  `json:"transaction_capacity"`
}

type AccountInfo struct {
	Username       string        `json:"username"`
	Created        int64         `json:"created"`
	MasterKey      crypto.PubKey `json:"master_key"`
	TransactionKey crypto.PubKey `json:"transaction_key"`
	PostKey        crypto.PubKey `json:"post_key"`
	Address        string        `json:"address"`
}

type AccountBank struct {
	Address         string        `json:"address"`
	Saving          Coin          `json:"saving"`
	Checking        Coin          `json:"checking"`
	Username        string        `json:"username"`
	Stake           Coin          `json:"stake"`
	FrozenMoneyList []FrozenMoney `json:"frozen_money_list"`
}

type FrozenMoney struct {
	Amount   Coin  `json:"amount"`
	StartAt  int64 `json:"start_at"`
	Times    int64 `json:"times"`
	Interval int64 `json:"interval"`
}

type GrantKeyList struct {
	GrantPubKeyList []GrantPubKey `json:"grant_public_key_list"`
}

type GrantPubKey struct {
	Username string        `json:"username"`
	PubKey   crypto.PubKey `json:"public_key"`
	Expire   int64         `json:"expire"`
}

type Reward struct {
	OriginalIncome Coin `json:"original_income"`
	FrictionIncome Coin `json:"friction_income"`
	ActualReward   Coin `json:"actual_reward"`
	UnclaimReward  Coin `json:"unclaim_reward"`
}

type Relationship struct {
	DonationTimes int64 `json:"donation_times"`
}

type FollowerMeta struct {
	CreatedAt    int64  `json:"created_at"`
	FollowerName string `json:"follower_name"`
}

type FollowingMeta struct {
	CreatedAt     int64  `json:"created_at"`
	FollowingName string `json:"following_name"`
}

// proposal related
type ProposalList struct {
	OngoingProposal []string `json:"ongoing_proposal"`
	PastProposal    []string `json:"past_proposal"`
}

type Proposal interface{}

type ProposalInfo struct {
	Creator       string `json:"creator"`
	ProposalID    string `json:"proposal_id"`
	AgreeVotes    Coin   `json:"agree_vote"`
	DisagreeVotes Coin   `json:"disagree_vote"`
	Result        int    `json:"result"`
}

type ChangeParamProposal struct {
	ProposalInfo
	Param Parameter `json:"param"`
}

type ContentCensorshipProposal struct {
	ProposalInfo
	PermLink string `json:"perm_link"`
}

type ProtocolUpgradeProposal struct {
	ProposalInfo
	Link string `json:"link"`
}
