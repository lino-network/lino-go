package model

import (
	"time"

	crypto "github.com/tendermint/tendermint/crypto"
	tmtypes "github.com/tendermint/tendermint/types"
)

//
// account related
//
type AccountInfo struct {
	Username       string        `json:"username"`
	CreatedAt      int64         `json:"created_at"`
	ResetKey       crypto.PubKey `json:"reset_key"`
	TransactionKey crypto.PubKey `json:"transaction_key"`
	AppKey         crypto.PubKey `json:"app_key"`
}

type AccountBank struct {
	Saving          Coin          `json:"saving"`
	Stake           Coin          `json:"stake"`
	FrozenMoneyList []FrozenMoney `json:"frozen_money_list"`
	NumOfTx         int64         `json:"number_of_transaction"`
	NumOfReward     int64         `json:"number_of_reward"`
}

type FrozenMoney struct {
	Amount   Coin  `json:"amount"`
	StartAt  int64 `json:"start_at"`
	Times    int64 `json:"times"`
	Interval int64 `json:"interval"`
}

type GrantPubKey struct {
	Username   string     `json:"username"`
	Permission Permission `json:"permission"`
	CreatedAt  int64      `json:"created_at"`
	ExpiresAt  int64      `json:"expires_at"`
	Amount     Coin       `json:"amount"`
}

type AccountMeta struct {
	Sequence             int64  `json:"sequence"`
	LastActivityAt       int64  `json:"last_activity_at"`
	TransactionCapacity  Coin   `json:"transaction_capacity"`
	JSONMeta             string `json:"json_meta"`
	LastReportOrUpvoteAt int64  `json:"last_report_or_upvote_at"`
}

type FollowerMeta struct {
	CreatedAt    int64  `json:"created_at"`
	FollowerName string `json:"follower_name"`
}

type FollowingMeta struct {
	CreatedAt     int64  `json:"created_at"`
	FollowingName string `json:"following_name"`
}

type Reward struct {
	TotalIncome     Coin `json:"total_income"`
	OriginalIncome  Coin `json:"original_income"`
	FrictionIncome  Coin `json:"friction_income"`
	InflationIncome Coin `json:"inflation_income"`
	UnclaimReward   Coin `json:"unclaim_reward"`
}

type RewardDetail struct {
	OriginalDonation Coin   `json:"original_donation"`
	FrictionDonation Coin   `json:"friction_donation"`
	ActualReward     Coin   `json:"actual_reward"`
	Consumer         string `json:"consumer"`
	PostAuthor       string `json:"post_author"`
	PostID           string `json:"post_id`
}

type RewardHistory struct {
	Details []RewardDetail `json:"details"`
}

type Relationship struct {
	DonationTimes int64 `json:"donation_times"`
}

type BalanceHistory struct {
	Details []Detail `json:"details"`
}

type Detail struct {
	DetailType DetailType `json:"detail_type"`
	From       string     `json:"from"`
	To         string     `json:"to"`
	Amount     Coin       `json:"amount"`
	Balance    Coin       `json:"balance"`
	CreatedAt  int64      `json:"created_at"`
	Memo       string     `json:"memo"`
}

//
// post related
//
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

// PostMeta stores tiny and frequently updated fields.
type PostMeta struct {
	CreatedAt               int64  `json:"created_at"`
	LastUpdatedAt           int64  `json:"last_updated_at"`
	LastActivityAt          int64  `json:"last_activity_at"`
	AllowReplies            bool   `json:"allow_replies"`
	IsDeleted               bool   `json:"is_deleted"`
	TotalDonateCount        int64  `json:"total_donate_count"`
	TotalReportStake        Coin   `json:"total_report_stake"`
	TotalUpvoteStake        Coin   `json:"total_upvote_stake"`
	TotalViewCount          int64  `json:"total_view_count"`
	TotalReward             Coin   `json:"total_reward"`
	RedistributionSplitRate string `json:"redistribution_split_rate"`
}

// Post is the combination of PostInfo and PostMeta.
type Post struct {
	PostID                  string           `json:"post_id"`
	Title                   string           `json:"title"`
	Content                 string           `json:"content"`
	Author                  string           `json:"author"`
	ParentAuthor            string           `json:"parent_author"`
	ParentPostID            string           `json:"parent_postID"`
	SourceAuthor            string           `json:"source_author"`
	SourcePostID            string           `json:"source_postID"`
	Links                   []IDToURLMapping `json:"links"`
	CreatedAt               int64            `json:"created_at"`
	LastUpdatedAt           int64            `json:"last_updated_at"`
	LastActivityAt          int64            `json:"last_activity_at"`
	AllowReplies            bool             `json:"allow_replies"`
	IsDeleted               bool             `json:"is_deleted"`
	TotalDonateCount        int64            `json:"total_donate_count"`
	TotalReportStake        Coin             `json:"total_report_stake"`
	TotalUpvoteStake        Coin             `json:"total_upvote_stake"`
	TotalViewCount          int64            `json:"total_view_count"`
	TotalReward             Coin             `json:"reward"`
	RedistributionSplitRate string           `json:"redistribution_split_rate"`
}

type ReportOrUpvote struct {
	Username  string `json:"username"`
	Stake     Coin   `json:"stake"`
	CreatedAt int64  `json:"created_at"`
	IsReport  bool   `json:"is_report"`
}

type Comment struct {
	Author    string `json:"author"`
	PostID    string `json:"post_id"`
	CreatedAt int64  `json:"created_at"`
}

type View struct {
	Username   string `json:"username"`
	LastViewAt int64  `json:"last_view_at"`
	Times      int64  `jons:"times"`
}

type Donations struct {
	Username string `json:"username"`
	Times    int64  `json:"times"`
	Amount   Coin   `json:"amount"`
}

//
// validator related struct
//
type PubKey struct {
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

type ABCIValidator struct {
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PubKey  PubKey `protobuf:"bytes,2,opt,name=pub_key,json=pubKey" json:"pub_key"`
	Power   int64  `protobuf:"varint,3,opt,name=power,proto3" json:"power,omitempty"`
}

type Validator struct {
	ABCIValidator
	Username        string `json:"username"`
	Deposit         Coin   `json:"deposit"`
	AbsentCommit    int64  `json:"absent_commit"`
	ByzantineCommit int64  `json:"byzantine_commit"`
	ProducedBlocks  int64  `json:"produced_blocks"`
	Link            string `json:"link"`
}

type ValidatorList struct {
	OncallValidators   []string `json:"oncall_validators"`
	AllValidators      []string `json:"all_validators"`
	PreBlockValidators []string `json:"pre_block_validators"`
	LowestPower        Coin     `json:"lowest_power"`
	LowestValidator    string   `json:"lowest_validator"`
}

//
// vote related struct
//
type Voter struct {
	Username       string `json:"username"`
	Deposit        Coin   `json:"deposit"`
	DelegatedPower Coin   `json:"delegated_power"`
}

type Vote struct {
	Voter       string `json:"voter"`
	VotingPower Coin   `json:"voting_power"`
	Result      bool   `json:"result"`
}

type Delegation struct {
	Delegator string `json:"delegator"`
	Amount    Coin   `json:"amount"`
}

//
// developer related
//
type Developer struct {
	Username       string `json:"username"`
	Deposit        Coin   `json:"deposit"`
	AppConsumption Coin   `json:"app_consumption"`
	Website        string `json:"website"`
	Description    string `json:"description"`
	AppMetaData    string `json:"app_meta_data"`
}

type DeveloperList struct {
	AllDevelopers []string `json:"all_developers"`
}

//
// infra provider related
//
type InfraProvider struct {
	Username string `json:"username"`
	Usage    int64  `json:"usage"`
}

type InfraProviderList struct {
	AllInfraProviders []string `json:"all_infra_providers"`
}

//
// proposal related
//
type Proposal interface{}

type ProposalInfo struct {
	Creator       string `json:"creator"`
	ProposalID    string `json:"proposal_id"`
	AgreeVotes    Coin   `json:"agree_vote"`
	DisagreeVotes Coin   `json:"disagree_vote"`
	Result        int    `json:"result"`
	CreatedAt     int64  `json:"created_at"`
	ExpiredAt     int64  `json:"expired_at"`
	Reason        string `json:"reason"`
}

type NextProposalID struct {
	NextProposalID int64 `json:"next_proposal_id"`
}

type ChangeParamProposal struct {
	ProposalInfo
	Param  Parameter `json:"param"`
	Reason string    `json:"reason"`
}

type ContentCensorshipProposal struct {
	ProposalInfo
	PermLink string `json:"perm_link"`
	Reason   string `json:"reason"`
}

type ProtocolUpgradeProposal struct {
	ProposalInfo
	Link   string `json:"link"`
	Reason string `json:"reason"`
}

type ProposalList struct {
	OngoingProposal []string `json:"ongoing_proposal"`
	PastProposal    []string `json:"past_proposal"`
}

//
// block related
//

type Block struct {
	Header     tmtypes.Header       `json:"header"`
	Evidence   tmtypes.EvidenceData `json:"evidence"`
	LastCommit *tmtypes.Commit      `json:"last_commit"`
	Data       *Data                `json:"data"`
}

type BlockStatus struct {
	LatestBlockHeight int64     `json:"latest_block_height"`
	LatestBlockTime   time.Time `json:"latest_block_time"`
}

type Data struct {
	Txs Txs `json:"txs"`
}

type Txs []Transaction
