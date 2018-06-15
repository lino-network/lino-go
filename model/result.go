package model

import (
	"github.com/cznic/mathutil"
	"github.com/tendermint/go-crypto"
	ttypes "github.com/tendermint/tendermint/types"
)

//
// account related
//
type AccountInfo struct {
	Username       string        `json:"username"`
	CreatedAt      int64         `json:"created_at"`
	MasterKey      crypto.PubKey `json:"master_key"`
	TransactionKey crypto.PubKey `json:"transaction_key"`
	PostKey        crypto.PubKey `json:"post_key"`
}

type AccountBank struct {
	Saving          Coin          `json:"saving"`
	Stake           Coin          `json:"stake"`
	FrozenMoneyList []FrozenMoney `json:"frozen_money_list"`
	NumOfTx         int64         `json:"number_of_transaction"`
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
	Username  string        `json:"username"`
	PubKey    crypto.PubKey `json:"public_key"`
	ExpiresAt int64         `json:"expires_at"`
}

type AccountMeta struct {
	Sequence            int64  `json:"sequence"`
	LastActivityAt      int64  `json:"last_activity_at"`
	TransactionCapacity Coin   `json:"transaction_capacity"`
	JSONMeta            string `json:"json_meta"`
}

type AccountInfraConsumption struct {
	Storage   int64 `json:"storage"`
	Bandwidth int64 `json:"bandwidth"`
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
	OriginalIncome Coin `json:"original_income"`
	FrictionIncome Coin `json:"friction_income"`
	ActualReward   Coin `json:"actual_reward"`
	UnclaimReward  Coin `json:"unclaim_reward"`
}

type Relationship struct {
	DonationTimes int64 `json:"donation_times"`
}

type BalanceHistory struct {
	Details []Detail `json:"details"`
}

type Detail struct {
	DetailType int    `json:"detail"`
	From       string `json:"from"`
	To         string `json:"to"`
	Amount     Coin   `json:"amount"`
	CreatedAt  int64  `json:"created_at"`
	Memo       string `json:"memo"`
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

type PostMeta struct {
	CreatedAt               int64 `json:"created_at"`
	LastUpdatedAt           int64 `json:"last_updated_at"`
	LastActivityAt          int64 `json:"last_activity_at"`
	AllowReplies            bool  `json:"allow_replies"`
	TotalLikeCount          int64 `json:"total_like_count"`
	TotalDonateCount        int64 `json:"total_donate_count"`
	TotalLikeWeight         int64 `json:"total_like_weight"`
	TotalDislikeWeight      int64 `json:"total_dislike_weight"`
	TotalReportStake        Coin  `json:"total_report_stake"`
	TotalUpvoteStake        Coin  `json:"total_upvote_stake"`
	TotalViewCount          int64 `json:"total_view_count"`
	TotalReward             Coin  `json:"reward"`
	RedistributionSplitRate Rat   `json:"redistribution_split_rate"`
}

type Like struct {
	Username  string `json:"username"`
	Weight    int64  `json:"weight"`
	CreatedAt int64  `json:"created_at"`
}

type ReportOrUpvote struct {
	Username  string `json:"username"`
	Stake     Coin   `json:"stake"`
	CreatedAt int64  `json:"created_at"`
	IsReport  bool   `json:"is_report"`
}

type Comment struct {
	Author    string `json:"author"`
	PostID    string `json:"post_key"`
	CreatedAt int64  `json:"created_at"`
}

type View struct {
	Username   string `json:"username"`
	LastViewAt int64  `json:"last_view_at"`
	Times      int64  `jons:"times"`
}

type Donation struct {
	Amount       Coin  `json:"amount"`
	CreatedAt    int64 `json:"created_at"`
	DonationType int   `json:"donation_type"`
}

type Donations struct {
	Username     string     `json:"username"`
	DonationList []Donation `json:"donation_list"`
}

//
// validator related struct
//
type ABCIValidator struct {
	PubKey []byte `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	Power  int64  `protobuf:"varint,2,opt,name=power,proto3" json:"power,omitempty"`
}

type Validator struct {
	ABCIValidator
	Username       string `json:"username"`
	Deposit        Coin   `json:"deposit"`
	AbsentCommit   int    `json:"absent_commit"`
	ProducedBlocks int64  `json:"produced_blocks"`
	Link           string `json:"link"`
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

type DelegateeList struct {
	DelegateeList []string `json:"delegatee_list"`
}

//
// developer related
//
type Developer struct {
	Username       string `json:"username"`
	Deposit        Coin   `json:"deposit"`
	AppConsumption Coin   `json:"app_consumption"`
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
}

type ChangeParamProposal struct {
	ProposalInfo
	Param Parameter `json:"param"`
}

type ContentCensorshipProposal struct {
	ProposalInfo
	PermLink string `json:"perm_link"`
	Reason   string `json:"reason"`
}

type ProtocolUpgradeProposal struct {
	ProposalInfo
	Link string `json:"link"`
}

type ProposalList struct {
	OngoingProposal []string `json:"ongoing_proposal"`
	PastProposal    []string `json:"past_proposal"`
}

//
// others
//
type Coin struct {
	Amount mathutil.Int128 `json:"amount"`
}

type Rat struct {
	Num   int64 `json:"num"`
	Denom int64 `json:"denom"`
}

type Block struct {
	Header     *ttypes.Header      `json:"header"`
	Evidence   ttypes.EvidenceData `json:"evidence"`
	LastCommit *ttypes.Commit      `json:"last_commit"`
	Data       *Data               `json:"data"`
}

type Data struct {
	Txs Txs `json:"txs"`
}

type Txs []Transaction
