package query

import (
	"bytes"
)

// Different KV store name
const (
	MainKVStoreKey      = "main"
	AccountKVStoreKey   = "account"
	PostKVStoreKey      = "post"
	ValidatorKVStoreKey = "validator"
	GlobalKVStoreKey    = "global"
	VoteKVStoreKey      = "vote"
	InfraKVStoreKey     = "infra"
	DeveloperKVStoreKey = "developer"
	ParamKVStoreKey     = "param"
	ProposalKVStoreKey  = "proposal"
	ReputationKVStore   = "reputation"
	BandwidthKVStoreKey = "bandwidth"

	AccountInfoSubStore           = "info"
	AccountBankSubStore           = "bank"
	AccountMetaSubStore           = "meta"
	AccountRewardSubStore         = "reward"
	AccountPendingCoinDaySubStore = "pendingCoinDay"
	AccountGrantPubKeySubStore    = "grantPubKey"
	AccountAllGrantPubKeys        = "allGrantPubKey"
	AccountTxAndSequence          = "txAndSeq"

	BandwidthInfoSubStore    = "bandwidthinfo"
	BlockInfoSubStore        = "blockinfo"
	AppBandwidthInfoSubStore = "appinfo"

	DeveloperSubStore     = "dev"
	DeveloperListSubStore = "devList"

	TimeEventListSubStore   = "timeEventList"
	globalMetaSubStore      = "globalMeta"
	inflationPoolSubStore   = "inflationPool"
	consumptionMetaSubStore = "consumptionMeta"
	tpsSubStore             = "tps"
	linoStakeStatSubStore   = "linoStakeStat"

	InfraProviderSubStore = "infra"
	InfraListSubStore     = "infraList"

	PostInfoSubStore           = "info"
	PostMetaSubStore           = "meta"
	PostReportOrUpvoteSubStore = "reportOrUpvote"
	PostCommentSubStore        = "comment"
	PostViewSubStore           = "view"

	NextProposalIDSubStore  = "next"
	OngoingProposalSubStore = "ongoing"
	ExpiredProposalSubStore = "expired"

	ValidatorSubStore     = "validator"
	ValidatorListSubStore = "valList"

	DelegationSubStore    = "delegation"
	VoterSubStore         = "voter"
	VoteSubStore          = "vote"
	ReferenceListSubStore = "refList"
	DelegateeSubStore     = "delegatee"

	AllocationParamSubStore              = "allocation"
	InfraInternalAllocationParamSubStore = "infraInternal"
	DeveloperParamSubStore               = "developer"
	VoteParamSubStore                    = "vote"
	ProposalParamSubStore                = "proposal"
	ValidatorParamSubStore               = "validator"
	CoinDayParamSubStore                 = "coinday"
	BandwidthParamSubStore               = "bandwidth"
	AccountParamSubStore                 = "account"
	PostParamSubStore                    = "post"
	ReputationParamSubStore              = "reputation"
	UserReputation                       = "rep"
)

var (
	// KeySeparator is the separator of substore key
	KeySeparator      = "/"
	PermLinkSeparator = "#"
)

func getHexSubstringAfterKeySeparator(key []byte) string {
	return string(key[bytes.Index(key, []byte(KeySeparator))+1:])
}

func getSubstringAfterKeySeparator(key []byte) string {
	return string(key[bytes.Index(key, []byte(KeySeparator)):])
}

func getSubstringAfterSubstore(key []byte) string {
	return string(key[1:])
}

//
// post related
//
func getPermlink(author string, postID string) string {
	return string(author + PermLinkSeparator + postID)
}
