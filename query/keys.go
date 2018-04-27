package query

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
)

func GetValidatorKey(accKey string) []byte {
	return append(ValidatorSubstore, accKey...)
}

func GetValidatorListKey() []byte {
	return ValidatorListSubstore
}

func GetVotePrefix(id string) []byte {
	return append(append(VoteSubstore, id...), KeySeparator...)
}

func GetVoteKey(proposalID string, voter string) []byte {
	return append(GetVotePrefix(proposalID), voter...)
}

func GetProposalKey(proposalID string) []byte {
	return append(ProposalSubstore, proposalID...)
}

func GetProposalListKey() []byte {
	return ProposalListSubStore
}

func GetDelegatorPrefix(me string) []byte {
	return append(append(DelegatorSubstore, me...), KeySeparator...)
}

func GetDelegationKey(me string, myDelegator string) []byte {
	return append(GetDelegatorPrefix(me), myDelegator...)
}

func GetVoterKey(me string) []byte {
	return append(VoterSubstore, me...)
}

func GetDeveloperKey(accKey string) []byte {
	return append(DeveloperSubstore, accKey...)
}

func GetDeveloperListKey() []byte {
	return DeveloperListSubstore
}

func GetInfraProviderKey(accKey string) []byte {
	return append(InfraProviderSubstore, accKey...)
}

func GetInfraProviderListKey() []byte {
	return InfraProviderListSubstore
}
