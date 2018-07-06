package transport

import (
	"encoding/hex"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/lino-network/lino-go/errors"
	"github.com/lino-network/lino-go/model"

	crypto "github.com/tendermint/go-crypto"
)

// ZeroFee is used in building a standard transaction.
var ZeroFee = model.Fee{
	Amount: model.SDKCoins{},
	Gas:    0,
}

// MakeCodec returns all interface and messages to Tendermint.
func MakeCodec() *wire.Codec {
	cdc := wire.NewCodec()

	cdc.RegisterInterface((*model.Msg)(nil), nil)
	cdc.RegisterInterface((*model.Tx)(nil), nil)

	// account
	cdc.RegisterConcrete(model.RegisterMsg{}, "lino/register", nil)
	cdc.RegisterConcrete(model.FollowMsg{}, "lino/follow", nil)
	cdc.RegisterConcrete(model.UnfollowMsg{}, "lino/unfollow", nil)
	cdc.RegisterConcrete(model.TransferMsg{}, "lino/transfer", nil)
	cdc.RegisterConcrete(model.ClaimMsg{}, "lino/claim", nil)
	cdc.RegisterConcrete(model.RecoverMsg{}, "lino/recover", nil)
	cdc.RegisterConcrete(model.UpdateAccountMsg{}, "lino/updateAcc", nil)

	// post
	cdc.RegisterConcrete(model.CreatePostMsg{}, "lino/createPost", nil)
	cdc.RegisterConcrete(model.UpdatePostMsg{}, "lino/updatePost", nil)
	cdc.RegisterConcrete(model.DeletePostMsg{}, "lino/deletePost", nil)
	cdc.RegisterConcrete(model.LikeMsg{}, "lino/like", nil)
	cdc.RegisterConcrete(model.DonateMsg{}, "lino/donate", nil)
	cdc.RegisterConcrete(model.ViewMsg{}, "lino/view", nil)
	cdc.RegisterConcrete(model.ReportOrUpvoteMsg{}, "lino/reportOrUpvote", nil)

	// validator
	cdc.RegisterConcrete(model.ValidatorDepositMsg{}, "lino/valDeposit", nil)
	cdc.RegisterConcrete(model.ValidatorWithdrawMsg{}, "lino/valWithdraw", nil)
	cdc.RegisterConcrete(model.ValidatorRevokeMsg{}, "lino/valRevoke", nil)

	// vote
	cdc.RegisterConcrete(model.VoterDepositMsg{}, "lino/voteDeposit", nil)
	cdc.RegisterConcrete(model.VoterRevokeMsg{}, "lino/voteRevoke", nil)
	cdc.RegisterConcrete(model.VoterWithdrawMsg{}, "lino/voteWithdraw", nil)
	cdc.RegisterConcrete(model.DelegateMsg{}, "lino/delegate", nil)
	cdc.RegisterConcrete(model.DelegatorWithdrawMsg{}, "lino/delegateWithdraw", nil)
	cdc.RegisterConcrete(model.RevokeDelegationMsg{}, "lino/delegateRevoke", nil)

	// developer
	cdc.RegisterConcrete(model.DeveloperRegisterMsg{}, "lino/devRegister", nil)
	cdc.RegisterConcrete(model.DeveloperRevokeMsg{}, "lino/devRevoke", nil)
	cdc.RegisterConcrete(model.GrantPermissionMsg{}, "lino/grantPermission", nil)
	cdc.RegisterConcrete(model.RevokePermissionMsg{}, "lino/revokePermission", nil)

	// infra provider
	cdc.RegisterConcrete(model.ProviderReportMsg{}, "lino/providerReport", nil)

	// proposal
	cdc.RegisterConcrete(model.VoteProposalMsg{}, "lino/voteProposal", nil)
	cdc.RegisterConcrete(model.DeletePostContentMsg{}, "lino/deletePostContent", nil)
	cdc.RegisterConcrete(model.UpgradeProtocolMsg{}, "lino/upgradeProtocol", nil)
	cdc.RegisterConcrete(model.ChangeGlobalAllocationParamMsg{}, "lino/changeGlobalAllocation", nil)
	cdc.RegisterConcrete(model.ChangeEvaluateOfContentValueParamMsg{}, "lino/changeEvaluation", nil)
	cdc.RegisterConcrete(model.ChangeInfraInternalAllocationParamMsg{}, "lino/changeInfraAllocation", nil)
	cdc.RegisterConcrete(model.ChangeVoteParamMsg{}, "lino/changeVoteParam", nil)
	cdc.RegisterConcrete(model.ChangeProposalParamMsg{}, "lino/changeProposalParam", nil)
	cdc.RegisterConcrete(model.ChangeDeveloperParamMsg{}, "lino/changeDeveloperParam", nil)
	cdc.RegisterConcrete(model.ChangeValidatorParamMsg{}, "lino/changeValidatorParam", nil)
	cdc.RegisterConcrete(model.ChangeBandwidthParamMsg{}, "lino/changeBandwidthParam", nil)
	cdc.RegisterConcrete(model.ChangeAccountParamMsg{}, "lino/changeAccountParam", nil)
	cdc.RegisterConcrete(model.ChangePostParamMsg{}, "lino/changePostParam", nil)

	// // TODO:
	// cdc.RegisterInterface((*model.Proposal)(nil), nil)
	// cdc.RegisterConcrete(&model.ChangeParamProposal{}, "changeParam", nil)
	// cdc.RegisterConcrete(&model.ProtocolUpgradeProposal{}, "upgrade", nil)
	// cdc.RegisterConcrete(&model.ContentCensorshipProposal{}, "censorship", nil)

	// cdc.RegisterInterface((*param.Parameter)(nil), nil)
	// cdc.RegisterConcrete(param.EvaluateOfContentValueParam{}, "param/contentValue", nil)
	// cdc.RegisterConcrete(param.GlobalAllocationParam{}, "param/allocation", nil)
	// cdc.RegisterConcrete(param.InfraInternalAllocationParam{}, "param/infaAllocation", nil)
	// cdc.RegisterConcrete(param.VoteParam{}, "param/vote", nil)
	// cdc.RegisterConcrete(param.ProposalParam{}, "param/proposal", nil)
	// cdc.RegisterConcrete(param.DeveloperParam{}, "param/developer", nil)
	// cdc.RegisterConcrete(param.ValidatorParam{}, "param/validator", nil)
	// cdc.RegisterConcrete(param.CoinDayParam{}, "param/coinDay", nil)
	// cdc.RegisterConcrete(param.BandwidthParam{}, "param/bandwidth", nil)
	// cdc.RegisterConcrete(param.AccountParam{}, "param/account", nil)
	// cdc.RegisterConcrete(param.PostParam{}, "param/post", nil)
	// cdc.RegisterConcrete(param.ProposalIDParam{}, "param/proposalID", nil)

	wire.RegisterCrypto(cdc)
	return cdc
}

// EncodeTx encodes a message to the standard transaction.
func EncodeTx(cdc *wire.Codec, msg interface{}, pubKey crypto.PubKey,
	sig crypto.Signature, seq int64) ([]byte, error) {
	stdSig := model.Signature{
		PubKey:   pubKey,
		Sig:      sig,
		Sequence: seq,
	}

	stdTx := model.Transaction{
		Msg:  msg,
		Sigs: []model.Signature{stdSig},
		Fee:  ZeroFee,
	}
	return cdc.MarshalJSON(stdTx)
}

// EncodeSignMsg encodes the message to the standard signed message.
func EncodeSignMsg(cdc *wire.Codec, msg interface{}, chainId string, seq int64) ([]byte, error) {
	feeBytes, err := cdc.MarshalJSON(ZeroFee)
	if err != nil {
		return nil, err
	}
	msgBytes, err := cdc.MarshalJSON(msg)
	if err != nil {
		return nil, err
	}
	stdSignMsg := model.SignMsg{
		ChainID:        chainId,
		AccountNumbers: []int64{},
		Sequences:      []int64{seq},
		FeeBytes:       feeBytes,
		MsgBytes:       msgBytes,
	}
	return json.Marshal(stdSignMsg)
}

// GetPrivKeyFromHex gets private key from private key hex.
func GetPrivKeyFromHex(privHex string) (crypto.PrivKey, error) {
	keyBytes, err := hex.DecodeString(privHex)
	if err != nil {
		return nil, err
	}
	return crypto.PrivKeyFromBytes(keyBytes)
}

// GetPubKeyFromHex gets public key from public key hex.
func GetPubKeyFromHex(pubHex string) (crypto.PubKey, error) {
	keyBytes, err := hex.DecodeString(pubHex)
	if err != nil {
		return nil, err
	}

	if keyBytes == nil || len(keyBytes) == 0 {
		return nil, errors.EmptyResponse("Empty bytes !")
	}
	return crypto.PubKeyFromBytes(keyBytes)
}
