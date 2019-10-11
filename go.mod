module github.com/lino-network/lino-go

go 1.12

require (
	github.com/btcsuite/btcutil v0.0.0-20180706230648-ab6388e0c60a
	github.com/cosmos/cosmos-sdk v0.37.0
	github.com/lino-network/lino v0.4.4
	github.com/spf13/viper v1.4.0
	github.com/tendermint/tendermint v0.32.5
	google.golang.org/appengine v1.4.0 // indirect
)

// new version doesn't support https parse. https://github.com/tendermint/tendermint/issues/4051
replace github.com/tendermint/tendermint => github.com/tendermint/tendermint v0.32.3
