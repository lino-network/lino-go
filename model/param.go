package model

// parameters can be changed by proposal
type Parameter interface{}

type EvaluateOfContentValueParam struct {
	ConsumptionTimeAdjustBase      int64 `json:"consumption_time_adjust_base"`
	ConsumptionTimeAdjustOffset    int64 `json:"consumption_time_adjust_offset"`
	NumOfConsumptionOnAuthorOffset int64 `json:"num_of_consumption_on_author_offset"`
	TotalAmountOfConsumptionBase   int64 `json:"total_amount_of_consumption_base"`
	TotalAmountOfConsumptionOffset int64 `json:"total_amount_of_consumption_offset"`
	AmountOfConsumptionExponent    Rat   `json:"amount_of_consumption_exponent"`
}

type GlobalAllocationParam struct {
	GlobalGrowthRate         Rat `json:"global_growth_rate"`
	InfraAllocation          Rat `json:"infra_allocation"`
	ContentCreatorAllocation Rat `json:"content_creator_allocation"`
	DeveloperAllocation      Rat `json:"developer_allocation"`
	ValidatorAllocation      Rat `json:"validator_allocation"`
}

type InfraInternalAllocationParam struct {
	StorageAllocation Rat `json:"storage_allocation"`
	CDNAllocation     Rat `json:"CDN_allocation"`
}

type VoteParam struct {
	VoterMinDeposit                Coin  `json:"voter_min_deposit"`
	DelegatorMinWithdraw           Coin  `json:"delegator_min_withdraw"`
	VoterCoinReturnIntervalSec     int64 `json:"voter_coin_return_interval_second"`
	VoterCoinReturnTimes           int64 `json:"voter_coin_return_times"`
	DelegatorCoinReturnIntervalSec int64 `json:"delegator_coin_return_interval_second"`
	DelegatorCoinReturnTimes       int64 `json:"delegator_coin_return_times"`
}

type ProposalParam struct {
	ContentCensorshipDecideSec  int64 `json:"content_censorship_decide_second"`
	ContentCensorshipMinDeposit Coin  `json:"content_censorship_min_deposit"`
	ContentCensorshipPassRatio  Rat   `json:"content_censorship_pass_ratio"`
	ContentCensorshipPassVotes  Coin  `json:"content_censorship_pass_votes"`
	ChangeParamDecideSec        int64 `json:"change_param_decide_second"`
	ChangeParamExecutionSec     int64 `json:"change_param_execution_second"`
	ChangeParamMinDeposit       Coin  `json:"change_param_min_deposit"`
	ChangeParamPassRatio        Rat   `json:"change_param_pass_ratio"`
	ChangeParamPassVotes        Coin  `json:"change_param_pass_votes"`
	ProtocolUpgradeDecideSec    int64 `json:"protocol_upgrade_decide_second"`
	ProtocolUpgradeMinDeposit   Coin  `json:"protocol_upgrade_min_deposit"`
	ProtocolUpgradePassRatio    Rat   `json:"protocol_upgrade_pass_ratio"`
	ProtocolUpgradePassVotes    Coin  `json:"protocol_upgrade_pass_votes"`
}

type DeveloperParam struct {
	DeveloperMinDeposit            Coin  `json:"developer_min_deposit"`
	DeveloperCoinReturnIntervalSec int64 `json:"developer_coin_return_interval_second"`
	DeveloperCoinReturnTimes       int64 `json:"developer_coin_return_times"`
}

type ValidatorParam struct {
	ValidatorMinWithdraw           Coin  `json:"validator_min_withdraw"`
	ValidatorMinVotingDeposit      Coin  `json:"validator_min_voting_deposit"`
	ValidatorMinCommitingDeposit   Coin  `json:"validator_min_commiting_deposit"`
	ValidatorCoinReturnIntervalSec int64 `json:"validator_coin_return_second"`
	ValidatorCoinReturnTimes       int64 `json:"validator_coin_return_times"`
	PenaltyMissVote                Coin  `json:"penalty_miss_vote"`
	PenaltyMissCommit              Coin  `json:"penalty_miss_commit"`
	PenaltyByzantine               Coin  `json:"penalty_byzantine"`
	ValidatorListSize              int64 `json:"validator_list_size"`
	AbsentCommitLimitation         int64 `json:"absent_commit_limitation"`
}

type CoinDayParam struct {
	SecondsToRecoverCoinDay int64 `json:"seconds_to_recover_coin_day"`
}

type BandwidthParam struct {
	SecondsToRecoverBandwidth   int64 `json:"seconds_to_recover_bandwidth"`
	CapacityUsagePerTransaction Coin  `json:"capacity_usage_per_transaction"`
	VirtualCoin                 Coin  `json:"virtual_coin"`
}

type AccountParam struct {
	MinimumBalance               Coin  `json:"minimum_balance"`
	RegisterFee                  Coin  `json:"register_fee"`
	FirstDepositFullCoinDayLimit Coin  `json:"first_deposit_full_coin_day_limit"`
	MaxNumFrozenMoney            int64 `json:"max_num_frozen_money"`
}

type PostParam struct {
	ReportOrUpvoteIntervalSec int64 `json:"report_or_upvote_interval_second"`
	PostIntervalSec           int64 `json:"post_interval_sec"`
}
