package transport

import (
	"encoding/hex"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/lino-network/lino-go/errors"
	"github.com/lino-network/lino-go/model"

	crypto "github.com/tendermint/tendermint/crypto"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
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
	cdc.RegisterConcrete(model.DeveloperUpdateMsg{}, "lino/devUpdate", nil)
	cdc.RegisterConcrete(model.DeveloperRevokeMsg{}, "lino/devRevoke", nil)
	cdc.RegisterConcrete(model.GrantPermissionMsg{}, "lino/grantPermission", nil)
	cdc.RegisterConcrete(model.RevokePermissionMsg{}, "lino/revokePermission", nil)
	cdc.RegisterConcrete(model.PreAuthorizationMsg{}, "lino/preAuthorizationPermission", nil)

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

	wire.RegisterCrypto(cdc)
	return cdc
}

func sortJSON(toSortJSON []byte) ([]byte, error) {
	var c interface{}
	err := json.Unmarshal(toSortJSON, &c)
	if err != nil {
		return nil, err
	}
	js, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return js, nil
}

// EncodeSignMsg encodes the message to the standard signed message.
func EncodeSignMsg(cdc *wire.Codec, msgs []model.Msg, chainId string, seq int64) ([]byte, error) {
	feeBytes, err := cdc.MarshalJSON(ZeroFee)
	if err != nil {
		return nil, err
	}

	var msgsBytes []json.RawMessage
	for _, msg := range msgs {
		bz, err := cdc.MarshalJSON(msg)
		if err != nil {
			return nil, err
		}

		signBytes, err := sortJSON(bz)
		if err != nil {
			return nil, err
		}

		msgsBytes = append(msgsBytes, json.RawMessage(signBytes))
	}

	stdSignMsg := model.SignMsg{
		AccountNumber: 0,
		ChainID:       chainId,
		Fee:           json.RawMessage(feeBytes),
		Memo:          "",
		Msgs:          msgsBytes,
		Sequence:      seq,
	}

	signMsgBytes, err := cdc.MarshalJSON(stdSignMsg)
	if err != nil {
		return nil, err
	}

	return sortJSON(signMsgBytes)
}

// EncodeTx encodes a message to the standard transaction.
func EncodeTx(cdc *wire.Codec, msgs []model.Msg, pubKey crypto.PubKey,
	sig []byte, seq int64, memo string) ([]byte, error) {
	stdSig := model.Signature{
		PubKey:   pubKey,
		Sig:      sig,
		Sequence: seq,
	}

	stdTx := model.Transaction{
		Msgs:       msgs,
		Fee:        ZeroFee,
		Signatures: []model.Signature{stdSig},
		Memo:       memo,
	}
	return cdc.MarshalJSON(stdTx)
}

// GetPrivKeyFromHex gets private key from private key hex.
func GetPrivKeyFromHex(privHex string) (crypto.PrivKey, error) {
	keyBytes, err := hex.DecodeString(privHex)
	if err != nil {
		return nil, err
	}
	return cryptoAmino.PrivKeyFromBytes(keyBytes)
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
	return cryptoAmino.PubKeyFromBytes(keyBytes)
}
