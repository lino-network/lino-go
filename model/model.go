package model

type ValidatorList struct {
	OncallValidators   []string `json:"oncall_validators"`
	AllValidators      []string `json:"all_validators"`
	PreBlockValidators []string `json:"pre_block_validators"`
	// LowestPower        types.Coin         `json:"lowest_power"`
	LowestValidator string `json:"lowest_validator"`
}
