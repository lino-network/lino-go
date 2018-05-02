package model

import "github.com/tendermint/go-crypto"

// validator related struct
type ValidatorList struct {
	OncallValidators   []string         `json:"oncall_validators"`
	AllValidators      []string         `json:"all_validators"`
	PreBlockValidators []string         `json:"pre_block_validators"`
	LowestPower        map[string]int64 `json:"lowest_power"`
	LowestValidator    string           `json:"lowest_validator"`
}

type Validator struct {
	Username     string           `json:"username"`
	Deposit      map[string]int64 `json:"deposit"`
	AbsentCommit int              `json:"absent_commit"`
}

// vote related struct
type Voter struct {
	Username       string           `json:"username"`
	Deposit        map[string]int64 `json:"deposit"`
	DelegatedPower map[string]int64 `json:"delegated_power"`
}

type Vote struct {
	Voter  string `json:"voter"`
	Result bool   `json:"result"`
}

type Delegation struct {
	Delegator string           `json:"delegator"`
	Amount    map[string]int64 `json:"amount"`
}

// post related

// developer related
type Developer struct {
	Username       string           `json:"username"`
	Deposit        map[string]int64 `json:"deposit"`
	AppConsumption map[string]int64 `json:"app_consumption"`
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
	Sequence            int64            `json:"sequence"`
	LastActivity        int64            `json:"last_activity"`
	TransactionCapacity map[string]int64 `json:"transaction_capacity"`
}

type AccountInfo struct {
	Username string        `json:"username"`
	Created  int64         `json:"created"`
	PostKey  crypto.PubKey `json:"post_key"`
	OwnerKey crypto.PubKey `json:"owner_key"`
	Address  string        `json:"address"`
}

type AccountBank struct {
	Address  string           `json:"address"`
	Balance  map[string]int64 `json:"balance"`
	Username string           `json:"username"`
	Stake    map[string]int64 `json:"stake"`
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
	OriginalIncome map[string]int64 `json:"original_income"`
	FrictionIncome map[string]int64 `json:"friction_income"`
	ActualReward   map[string]int64 `json:"actual_reward"`
	UnclaimReward  map[string]int64 `json:"unclaim_reward"`
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
