package query

import "encoding/hex"

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
