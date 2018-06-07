package query

import (
	"encoding/hex"
	"strconv"
)

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
)

var (
	KeySeparator = "/"

	accountInfoSubstore              = []byte{0x00}
	accountBankSubstore              = []byte{0x01}
	accountMetaSubstore              = []byte{0x02}
	accountFollowerSubstore          = []byte{0x03}
	accountFollowingSubstore         = []byte{0x04}
	accountRewardSubstore            = []byte{0x05}
	accountPendingStakeQueueSubstore = []byte{0x06}
	accountRelationshipSubstore      = []byte{0x07}
	accountGrantListSubstore         = []byte{0x08}
	accountBalanceHistorySubstore    = []byte{0x09}

	postInfoSubStore           = []byte{0x00} // SubStore for all post info
	postMetaSubStore           = []byte{0x01} // SubStore for all post mata info
	postLikeSubStore           = []byte{0x02} // SubStore for all like to post
	postReportOrUpvoteSubStore = []byte{0x03} // SubStore for all like to post
	postCommentSubStore        = []byte{0x04} // SubStore for all comments
	postViewsSubStore          = []byte{0x05} // SubStore for all views
	postDonationsSubStore      = []byte{0x06} // SubStore for all donations

	validatorSubstore     = []byte{0x00}
	validatorListSubstore = []byte{0x01}

	delegatorSubstore = []byte{0x00}
	voterSubstore     = []byte{0x01}
	voteSubstore      = []byte{0x02}

	developerSubstore     = []byte{0x00}
	developerListSubstore = []byte{0x01}

	infraProviderSubstore     = []byte{0x00}
	infraProviderListSubstore = []byte{0x01}

	proposalSubstore     = []byte{0x00}
	proposalListSubStore = []byte{0x01}

	allocationParamSubStore              = []byte{0x00} // SubStore for allocation
	infraInternalAllocationParamSubStore = []byte{0x01} // SubStore for infrat internal allocation
	evaluateOfContentValueParamSubStore  = []byte{0x02} // Substore for evaluate of content value
	developerParamSubStore               = []byte{0x03} // Substore for developer param
	voteParamSubStore                    = []byte{0x04} // Substore for vote param
	proposalParamSubStore                = []byte{0x05} // Substore for proposal param
	validatorParamSubStore               = []byte{0x06} // Substore for validator param
	coinDayParamSubStore                 = []byte{0x07} // Substore for coin day param
	bandwidthParamSubStore               = []byte{0x08} // Substore for bandwidth param
	accountParamSubstore                 = []byte{0x09} // Substore for account param
)

//
// account related
//
func getAccountInfoKey(accKey string) []byte {
	return append(accountInfoSubstore, accKey...)
}

func getAccountBankKey(address string) []byte {
	bz, _ := hex.DecodeString(address)
	return append(accountBankSubstore, []byte(bz)...)
}

func getAccountMetaKey(accKey string) []byte {
	return append(accountMetaSubstore, accKey...)
}
func getGrantKeyListKey(accKey string) []byte {
	return append(accountGrantListSubstore, accKey...)
}

func getBalanceHistoryPrefix(me string) []byte {
	return append(append(accountBalanceHistorySubstore, me...), KeySeparator...)
}
func getBalanceHistoryKey(me string, atWhen int64) []byte {
	return strconv.AppendInt(getBalanceHistoryPrefix(me), atWhen, 10)
}

func getRewardKey(accKey string) []byte {
	return append(accountRewardSubstore, accKey...)
}

func getPendingStakeQueueKey(accKey string) []byte {
	return append(accountPendingStakeQueueSubstore, accKey...)
}

func getRelationshipPrefix(me string) []byte {
	return append(append(accountRelationshipSubstore, me...), KeySeparator...)
}

func getRelationshipKey(me string, other string) []byte {
	return append(getRelationshipPrefix(me), other...)
}

func getFollowerPrefix(me string) []byte {
	return append(append(accountFollowerSubstore, me...), KeySeparator...)
}

func getFollowingPrefix(me string) []byte {
	return append(append(accountFollowingSubstore, me...), KeySeparator...)
}

func getFollowerKey(me string, myFollower string) []byte {
	return append(getFollowerPrefix(me), myFollower...)
}

func getFollowingKey(me string, myFollowing string) []byte {
	return append(getFollowingPrefix(me), myFollowing...)
}

//
// post related
//
func getPostInfoKey(postKey string) []byte {
	return append([]byte(postInfoSubStore), postKey...)
}

func getPostKey(author string, postID string) string {
	return string(string(author) + "#" + postID)
}

func getPostMetaKey(postKey string) []byte {
	return append([]byte(postMetaSubStore), postKey...)
}

func getPostLikePrefix(postKey string) []byte {
	return append(append([]byte(postLikeSubStore), postKey...), KeySeparator...)
}

func getPostLikeKey(postKey string, likeUser string) []byte {
	return append(getPostLikePrefix(postKey), likeUser...)
}

func getPostReportOrUpvotePrefix(postKey string) []byte {
	return append(append([]byte(postReportOrUpvoteSubStore), postKey...), KeySeparator...)
}

func getPostReportOrUpvoteKey(postKey string, user string) []byte {
	return append(getPostReportOrUpvotePrefix(postKey), user...)
}

func getPostViewPrefix(postKey string) []byte {
	return append(append([]byte(postViewsSubStore), postKey...), KeySeparator...)
}

func getPostViewKey(postKey string, viewUser string) []byte {
	return append(getPostViewPrefix(postKey), viewUser...)
}

func getPostCommentPrefix(postKey string) []byte {
	return append(append([]byte(postCommentSubStore), postKey...), KeySeparator...)
}

func getPostCommentKey(postKey string, commentPostKey string) []byte {
	return append(getPostCommentPrefix(postKey), commentPostKey...)
}

func getPostDonationPrefix(postKey string) []byte {
	return append(append([]byte(postDonationsSubStore), postKey...), KeySeparator...)
}

func getPostDonationKey(postKey string, donateUser string) []byte {
	return append(getPostDonationPrefix(postKey), donateUser...)
}

//
//  validator related
//
func getValidatorKey(accKey string) []byte {
	return append(validatorSubstore, accKey...)
}

func getValidatorListKey() []byte {
	return validatorListSubstore
}

//
// vote related
//
func getVotePrefix(id string) []byte {
	return append(append(voteSubstore, id...), KeySeparator...)
}

func getVoteKey(proposalID string, voter string) []byte {
	return append(getVotePrefix(proposalID), voter...)
}

func getDelegatorPrefix(me string) []byte {
	return append(append(delegatorSubstore, me...), KeySeparator...)
}

func getDelegationKey(me string, myDelegator string) []byte {
	return append(getDelegatorPrefix(me), myDelegator...)
}

func getVoterKey(me string) []byte {
	return append(voterSubstore, me...)
}

//
// developer related
//
func getDeveloperKey(accKey string) []byte {
	return append(developerSubstore, accKey...)
}

func getDeveloperListKey() []byte {
	return developerListSubstore
}

//
// infra related
//
func getInfraProviderKey(accKey string) []byte {
	return append(infraProviderSubstore, accKey...)
}

func getInfraProviderListKey() []byte {
	return infraProviderListSubstore
}

//
// proposal related
//
func getProposalKey(proposalID string) []byte {
	return append(proposalSubstore, proposalID...)
}

func getProposalListKey() []byte {
	return proposalListSubStore
}

//
// param related
//
func getEvaluateOfContentValueParamKey() []byte {
	return evaluateOfContentValueParamSubStore
}

func getGlobalAllocationParamKey() []byte {
	return allocationParamSubStore
}

func getInfraInternalAllocationParamKey() []byte {
	return infraInternalAllocationParamSubStore
}

func getDeveloperParamKey() []byte {
	return developerParamSubStore
}

func getVoteParamKey() []byte {
	return voteParamSubStore
}

func getValidatorParamKey() []byte {
	return validatorParamSubStore
}

func getProposalParamKey() []byte {
	return proposalParamSubStore
}

func getCoinDayParamKey() []byte {
	return coinDayParamSubStore
}

func getBandwidthParamKey() []byte {
	return bandwidthParamSubStore
}

func getAccountParamKey() []byte {
	return accountParamSubstore
}
