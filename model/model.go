package model

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
