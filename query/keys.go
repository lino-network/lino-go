package query

import (
	"bytes"
	"encoding/hex"
	"strconv"

	crypto "github.com/tendermint/go-crypto"
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
)

var (
	// KeySeparator is the separator of substore key
	KeySeparator = "/"

	// account substore
	accountInfoSubstore              = []byte{0x00}
	accountBankSubstore              = []byte{0x01}
	accountMetaSubstore              = []byte{0x02}
	accountFollowerSubstore          = []byte{0x03}
	accountFollowingSubstore         = []byte{0x04}
	accountRewardSubstore            = []byte{0x05}
	accountPendingStakeQueueSubstore = []byte{0x06}
	accountRelationshipSubstore      = []byte{0x07}
	accountBalanceHistorySubstore    = []byte{0x08}
	accountGrantPubKeySubstore       = []byte{0x09}

	// post substore
	postInfoSubStore           = []byte{0x00}
	postMetaSubStore           = []byte{0x01}
	postLikeSubStore           = []byte{0x02}
	postReportOrUpvoteSubStore = []byte{0x03}
	postCommentSubStore        = []byte{0x04}
	postViewsSubStore          = []byte{0x05}
	postDonationsSubStore      = []byte{0x06}

	// validator substore
	validatorSubstore     = []byte{0x00}
	validatorListSubstore = []byte{0x01}

	// vote substore
	delegationSubstore    = []byte{0x00}
	voterSubstore         = []byte{0x01}
	voteSubstore          = []byte{0x02}
	referenceListSubStore = []byte{0x03}
	delegateeSubStore     = []byte{0x04}

	// developer substore
	developerSubstore     = []byte{0x00}
	developerListSubstore = []byte{0x01}

	// infra provider substore
	infraProviderSubstore     = []byte{0x00}
	infraProviderListSubstore = []byte{0x01}

	// proposal substore
	proposalSubstore       = []byte{0x00}
	proposalListSubStore   = []byte{0x01}
	nextProposalIDSubstore = []byte{0x02}

	// param substore
	allocationParamSubStore              = []byte{0x00}
	infraInternalAllocationParamSubStore = []byte{0x01}
	evaluateOfContentValueParamSubStore  = []byte{0x02}
	developerParamSubStore               = []byte{0x03}
	voteParamSubStore                    = []byte{0x04}
	proposalParamSubStore                = []byte{0x05}
	validatorParamSubStore               = []byte{0x06}
	coinDayParamSubStore                 = []byte{0x07}
	bandwidthParamSubStore               = []byte{0x08}
	accountParamSubstore                 = []byte{0x09}
	postParamSubStore                    = []byte{0x10}
)

func getHexSubstringAfterKeySeparator(key []byte) string {
	return hex.EncodeToString(key[bytes.Index(key, []byte(KeySeparator)):])
}

func getSubstringAfterKeySeparator(key []byte) string {
	return string(key[bytes.Index(key, []byte(KeySeparator)):])
}

func getSubstringAfterSubstore(key []byte) string {
	return string(key[1:])
}

//
// account related
//
func getAccountInfoKey(accKey string) []byte {
	return append(accountInfoSubstore, accKey...)
}

func getAccountBankKey(accKey string) []byte {
	return append(accountBankSubstore, accKey...)
}

func getAccountMetaKey(accKey string) []byte {
	return append(accountMetaSubstore, accKey...)
}

func getFollowerKey(me string, myFollower string) []byte {
	return append(getFollowerPrefix(me), myFollower...)
}

func getFollowerPrefix(me string) []byte {
	return append(append(accountFollowerSubstore, me...), KeySeparator...)
}

func getFollowingKey(me string, myFollowing string) []byte {
	return append(getFollowingPrefix(me), myFollowing...)
}

func getFollowingPrefix(me string) []byte {
	return append(append(accountFollowingSubstore, me...), KeySeparator...)
}

func getRewardKey(accKey string) []byte {
	return append(accountRewardSubstore, accKey...)
}

func getRelationshipKey(me string, other string) []byte {
	return append(getRelationshipPrefix(me), other...)
}

func getRelationshipPrefix(me string) []byte {
	return append(append(accountRelationshipSubstore, me...), KeySeparator...)
}

func getPendingStakeQueueKey(accKey string) []byte {
	return append(accountPendingStakeQueueSubstore, accKey...)
}

func getBalanceHistoryPrefix(me string) []byte {
	return append(append(accountBalanceHistorySubstore, me...), KeySeparator...)
}
func getBalanceHistoryKey(me string, bucketSlot int64) []byte {
	return strconv.AppendInt(getBalanceHistoryPrefix(me), bucketSlot, 10)
}

func getGrantPubKeyPrefix(me string) []byte {
	return append(append(accountGrantPubKeySubstore, me...), KeySeparator...)
}

func getGrantPubKeyKey(me string, pubKey crypto.PubKey) []byte {
	return append(getGrantPubKeyPrefix(me), pubKey.Bytes()...)
}

//
// post related
//
func getPermlink(author string, postID string) string {
	return string(author + "#" + postID)
}

func getUserPostInfoPrefix(me string) []byte {
	return append(postInfoSubStore, me...)
}

func getPostInfoKey(permlink string) []byte {
	return append(postInfoSubStore, permlink...)
}

func getUserPostMetaPrefix(me string) []byte {
	return append(postMetaSubStore, me...)
}

func getPostMetaKey(permlink string) []byte {
	return append(postMetaSubStore, permlink...)
}

func getPostLikePrefix(permlink string) []byte {
	return append(append(postLikeSubStore, permlink...), KeySeparator...)
}

func getPostLikeKey(permlink string, likeUser string) []byte {
	return append(getPostLikePrefix(permlink), likeUser...)
}

func getPostReportOrUpvotePrefix(permlink string) []byte {
	return append(append(postReportOrUpvoteSubStore, permlink...), KeySeparator...)
}

func getPostReportOrUpvoteKey(permlink string, user string) []byte {
	return append(getPostReportOrUpvotePrefix(permlink), user...)
}

func getPostViewPrefix(permlink string) []byte {
	return append(append(postViewsSubStore, permlink...), KeySeparator...)
}

func getPostViewKey(permlink string, viewUser string) []byte {
	return append(getPostViewPrefix(permlink), viewUser...)
}

func getPostCommentPrefix(permlink string) []byte {
	return append(append(postCommentSubStore, permlink...), KeySeparator...)
}

func getPostCommentKey(permlink string, commentPermlink string) []byte {
	return append(getPostCommentPrefix(permlink), commentPermlink...)
}

func getPostDonationsPrefix(permlink string) []byte {
	return append(append(postDonationsSubStore, permlink...), KeySeparator...)
}

func getPostDonationsKey(permlink string, donateUser string) []byte {
	return append(getPostDonationsPrefix(permlink), donateUser...)
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
func getDelegationPrefix(me string) []byte {
	return append(append(delegationSubstore, me...), KeySeparator...)
}

func getDelegationKey(me string, myDelegator string) []byte {
	return append(getDelegationPrefix(me), myDelegator...)
}

func getVotePrefix(id string) []byte {
	return append(append(voteSubstore, id...), KeySeparator...)
}

func getVoteKey(proposalID string, voter string) []byte {
	return append(getVotePrefix(proposalID), voter...)
}

func getVoterKey(me string) []byte {
	return append(voterSubstore, me...)
}

func getReferenceListKey() []byte {
	return referenceListSubStore
}

func getDelegateePrefix(me string) []byte {
	return append(append(delegateeSubStore, me...), KeySeparator...)
}

func getDelegateeKey(me, delegatee string) []byte {
	return append(getDelegateePrefix(me), delegatee...)
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

func GetNextProposalIDKey() []byte {
	return nextProposalIDSubstore
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

func getPostParamKey() []byte {
	return postParamSubStore
}
