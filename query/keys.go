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

var ValidatorSubstore = []byte("Validator/")
var ValidatorListSubstore = []byte("ValidatorList/ValidatorListKey")

func GetValidatorKey(accKey string) []byte {
	return append(ValidatorSubstore, accKey...)
}

func GetValidatorListKey() []byte {
	return ValidatorListSubstore
}
