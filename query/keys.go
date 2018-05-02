package query

import (
	"encoding/hex"
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
)

var (
	DelegatorSubstore              = []byte{0x00}
	VoterSubstore                  = []byte{0x01}
	ProposalSubstore               = []byte{0x02}
	VoteSubstore                   = []byte{0x03}
	ProposalListSubStore           = []byte{0x04}
	ValidatorReferenceListSubStore = []byte{0x05}
	KeySeparator                   = "/"
	ValidatorSubstore              = []byte("Validator/")
	ValidatorListSubstore          = []byte("ValidatorList/ValidatorListKey")
	DeveloperSubstore              = []byte("Developer/")
	DeveloperListSubstore          = []byte("Developer/DeveloperListKey")
	InfraProviderSubstore          = []byte("InfraProvider/")
	InfraProviderListSubstore      = []byte("InfraProvider/InfraProviderListKey")

	AccountInfoSubstore              = []byte{0x00}
	AccountBankSubstore              = []byte{0x01}
	AccountMetaSubstore              = []byte{0x02}
	AccountFollowerSubstore          = []byte{0x03}
	AccountFollowingSubstore         = []byte{0x04}
	AccountRewardSubstore            = []byte{0x05}
	AccountPendingStakeQueueSubstore = []byte{0x06}
	AccountRelationshipSubstore      = []byte{0x07}
	AccountGrantListSubstore         = []byte{0x08}

	postInfoSubStore           = []byte{0x00} // SubStore for all post info
	postMetaSubStore           = []byte{0x01} // SubStore for all post mata info
	postLikeSubStore           = []byte{0x02} // SubStore for all like to post
	postReportOrUpvoteSubStore = []byte{0x03} // SubStore for all like to post
	postCommentSubStore        = []byte{0x04} // SubStore for all comments
	postViewsSubStore          = []byte{0x05} // SubStore for all views
	postDonationsSubStore      = []byte{0x06} // SubStore for all donations
)

func getValidatorKey(accKey string) []byte {
	return append(ValidatorSubstore, accKey...)
}

func getValidatorListKey() []byte {
	return ValidatorListSubstore
}

func getVotePrefix(id string) []byte {
	return append(append(VoteSubstore, id...), KeySeparator...)
}

func getVoteKey(proposalID string, voter string) []byte {
	return append(getVotePrefix(proposalID), voter...)
}

func getProposalKey(proposalID string) []byte {
	return append(ProposalSubstore, proposalID...)
}

func getProposalListKey() []byte {
	return ProposalListSubStore
}

func getDelegatorPrefix(me string) []byte {
	return append(append(DelegatorSubstore, me...), KeySeparator...)
}

func getDelegationKey(me string, myDelegator string) []byte {
	return append(getDelegatorPrefix(me), myDelegator...)
}

func getVoterKey(me string) []byte {
	return append(VoterSubstore, me...)
}

func getDeveloperKey(accKey string) []byte {
	return append(DeveloperSubstore, accKey...)
}

func getDeveloperListKey() []byte {
	return DeveloperListSubstore
}

func getInfraProviderKey(accKey string) []byte {
	return append(InfraProviderSubstore, accKey...)
}

func getInfraProviderListKey() []byte {
	return InfraProviderListSubstore
}

func getAccountInfoKey(accKey string) []byte {
	return append(AccountInfoSubstore, accKey...)
}

func getAccountBankKey(address string) []byte {
	bz, _ := hex.DecodeString(address)
	return append(AccountBankSubstore, []byte(bz)...)
}

func getAccountMetaKey(accKey string) []byte {
	return append(AccountMetaSubstore, accKey...)
}
func getGrantKeyListKey(accKey string) []byte {
	return append(AccountGrantListSubstore, accKey...)
}

func getRewardKey(accKey string) []byte {
	return append(AccountRewardSubstore, accKey...)
}

func getRelationshipPrefix(me string) []byte {
	return append(append(AccountRelationshipSubstore, me...), KeySeparator...)
}

func getRelationshipKey(me string, other string) []byte {
	return append(getRelationshipPrefix(me), other...)
}

func getFollowerPrefix(me string) []byte {
	return append(append(AccountFollowerSubstore, me...), KeySeparator...)
}

func getFollowingPrefix(me string) []byte {
	return append(append(AccountFollowingSubstore, me...), KeySeparator...)
}

func getFollowerKey(me string, myFollower string) []byte {
	return append(getFollowerPrefix(me), myFollower...)
}

func getFollowingKey(me string, myFollowing string) []byte {
	return append(getFollowingPrefix(me), myFollowing...)
}

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
