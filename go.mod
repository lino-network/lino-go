module github.com/lino-network/lino-go

go 1.16

// new version doesn't support https parse. https://github.com/tendermint/tendermint/issues/4051
replace github.com/tendermint/tendermint => github.com/tendermint/tendermint v0.32.3

require (
	github.com/cosmos/cosmos-sdk v0.47.3
	github.com/lino-network/lino v0.6.11
	github.com/spf13/viper v1.14.0
	github.com/tendermint/tendermint v0.37.0-rc2
)
