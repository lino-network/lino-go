package transport

import (
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/lino-network/lino-go/model"
	crypto "github.com/tendermint/go-crypto"
)

var ZeroFee = Fee{
	Amount: []int64{},
	Gas:    0,
}

func MakeCodec() *wire.Codec {
	cdc := wire.NewCodec()

	cdc.RegisterInterface((*model.Msg)(nil), nil)
	cdc.RegisterConcrete(model.RegisterMsg{}, "register/register", nil)
	cdc.RegisterConcrete(model.FollowMsg{}, "account/follow", nil)
	cdc.RegisterConcrete(model.UnfollowMsg{}, "account/unfollow", nil)
	cdc.RegisterConcrete(model.TransferMsg{}, "account/transfer", nil)
	cdc.RegisterConcrete(model.CreatePostMsg{}, "post/post", nil)
	cdc.RegisterConcrete(model.LikeMsg{}, "post/like", nil)
	cdc.RegisterConcrete(model.DonateMsg{}, "post/donate", nil)
	cdc.RegisterConcrete(model.ValidatorWithdrawMsg{}, "post/withdraw", nil)
	cdc.RegisterConcrete(model.ValidatorRevokeMsg{}, "post/revoke", nil)
	cdc.RegisterConcrete(model.ClaimMsg{}, "account/claim", nil)
	cdc.RegisterConcrete(model.VoterDepositMsg{}, "vote/deposit", nil)
	cdc.RegisterConcrete(model.VoterRevokeMsg{}, "vote/revoke", nil)
	cdc.RegisterConcrete(model.VoterWithdrawMsg{}, "vote/withdraw", nil)
	cdc.RegisterConcrete(model.DelegateMsg{}, "vote/delegate", nil)
	cdc.RegisterConcrete(model.DelegatorWithdrawMsg{}, "vote/delegate/withdraw", nil)
	cdc.RegisterConcrete(model.RevokeDelegationMsg{}, "vote/delegate/revoke", nil)
	cdc.RegisterConcrete(model.VoteMsg{}, "vote/vote", nil)
	cdc.RegisterConcrete(model.DeveloperRegisterMsg{}, "developer/register", nil)
	cdc.RegisterConcrete(model.DeveloperRevokeMsg{}, "developer/revoke", nil)
	cdc.RegisterConcrete(model.ProviderReportMsg{}, "provider/report", nil)
	cdc.RegisterConcrete(model.GrantDeveloperMsg{}, "grant/developer", nil)

	wire.RegisterCrypto(cdc)
	return cdc
}

func EncodeTx(cdc *wire.Codec, msg interface{}, pubKey crypto.PubKey,
	sig crypto.Signature, seq int64) ([]byte, error) {
	stdSig := Signature{
		PubKey:   pubKey,
		Sig:      sig,
		Sequence: seq,
	}

	stdTx := Transaction{
		Msg:  msg,
		Sigs: []Signature{stdSig},
		Fee:  ZeroFee,
	}
	return cdc.MarshalJSON(stdTx)
}

func EncodeSignMsg(cdc *wire.Codec, msg interface{}, chainId string, seq int64) ([]byte, error) {
	feeBytes, err := cdc.MarshalJSON(ZeroFee)
	if err != nil {
		return nil, err
	}
	msgBytes, err := cdc.MarshalJSON(msg)
	if err != nil {
		return nil, err
	}
	stdSignMsg := SignMsg{
		ChainID:   chainId,
		MsgBytes:  msgBytes,
		Sequences: []int64{seq},
		FeeBytes:  feeBytes,
	}
	return cdc.MarshalJSON(stdSignMsg)
}
