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
