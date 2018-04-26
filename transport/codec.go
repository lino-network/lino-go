package transport

import (
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/lino-network/lino-go/model"
	oldwire "github.com/tendermint/go-wire"
)

const (
	msgTypeRegister          = 0x1
	msgTypeFollow            = 0x2
	msgTypeUnfollow          = 0x3
	msgTypeTransfer          = 0x4
	msgTypePost              = 0x5
	msgTypeLike              = 0x6
	msgTypeDonate            = 0x7
	msgTypeValidatorDeposit  = 0x8
	msgTypeValidatorWithdraw = 0x9
	msgTypeValidatorRevoke   = 0x10
	msgTypeClaim             = 0x11
	msgTypeVoterDeposit      = 0x12
	msgTypeVoterRevoke       = 0x13
	msgTypeVoterWithdraw     = 0x14
	msgTypeDelegate          = 0x15
	msgTypeDelegatorWithdraw = 0x16
	msgTypeRevokeDelegation  = 0x17
	msgTypeVote              = 0x18
	msgTypeCreateProposal    = 0x19
	msgTypeDeveloperRegister = 0x20
	msgTypeDeveloperRevoke   = 0x21
	msgTypeProviderReport    = 0x22
)

func MakeCodec() *wire.Codec {

	var _ = oldwire.RegisterInterface(
		struct{ model.Transaction }{},
		oldwire.ConcreteType{model.TransferToAddress{}, msgTypeTransfer},
	)

	cdc := wire.NewCodec()
	return cdc
}
