package transport

import (
	"encoding/hex"
	"encoding/json"

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
	cdc.RegisterConcrete(model.RegisterMsg{}, "register", nil)
	cdc.RegisterConcrete(model.FollowMsg{}, "follow", nil)
	cdc.RegisterConcrete(model.UnfollowMsg{}, "unfollow", nil)
	cdc.RegisterConcrete(model.TransferMsg{}, "transfer", nil)
	cdc.RegisterConcrete(model.ClaimMsg{}, "claim", nil)
	cdc.RegisterConcrete(model.CreatePostMsg{}, "post", nil)
	cdc.RegisterConcrete(model.LikeMsg{}, "like", nil)
	cdc.RegisterConcrete(model.DonateMsg{}, "donate", nil)
	cdc.RegisterConcrete(model.ReportOrUpvoteMsg{}, "reportOrUpvote", nil)
	cdc.RegisterConcrete(model.ValidatorDepositMsg{}, "val/deposit", nil)
	cdc.RegisterConcrete(model.ValidatorWithdrawMsg{}, "val/withdraw", nil)
	cdc.RegisterConcrete(model.ValidatorRevokeMsg{}, "val/revoke", nil)
	cdc.RegisterConcrete(model.VoterDepositMsg{}, "vote/deposit", nil)
	cdc.RegisterConcrete(model.VoterRevokeMsg{}, "vote/revoke", nil)
	cdc.RegisterConcrete(model.VoterWithdrawMsg{}, "vote/withdraw", nil)
	cdc.RegisterConcrete(model.DelegateMsg{}, "delegate", nil)
	cdc.RegisterConcrete(model.DelegatorWithdrawMsg{}, "delegate/withdraw", nil)
	cdc.RegisterConcrete(model.RevokeDelegationMsg{}, "delegate/revoke", nil)
	cdc.RegisterConcrete(model.VoteMsg{}, "vote", nil)
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
	feeBytes, err := json.Marshal(ZeroFee)
	if err != nil {
		return nil, err
	}
	msgBytes, err := json.Marshal(msg)
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

func GetPrivKeyFromHex(privHex string) (crypto.PrivKey, error) {
	keyBytes, err := hex.DecodeString(privHex)
	if err != nil {
		return nil, err
	}
	return crypto.PrivKeyFromBytes(keyBytes)
}

func GetPubKeyFromHex(pubHex string) (crypto.PubKey, error) {
	keyBytes, err := hex.DecodeString(pubHex)
	if err != nil {
		return nil, err
	}
	return crypto.PubKeyFromBytes(keyBytes)
}
