package transport

import (
	"encoding/hex"
	"encoding/json"

	wire "github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth"
	txbuilder "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/lino-network/lino-go/errors"
	linotypes "github.com/lino-network/lino/types"
	acctypes "github.com/lino-network/lino/x/account/types"
	devtypes "github.com/lino-network/lino/x/developer/types"
	posttypes "github.com/lino-network/lino/x/post/types"
	proposal "github.com/lino-network/lino/x/proposal"
	valtypes "github.com/lino-network/lino/x/validator"
	votetypes "github.com/lino-network/lino/x/vote"

	crypto "github.com/tendermint/tendermint/crypto"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
)

// MakeCodec returns all interface and messages to Tendermint.
func MakeCodec() *wire.Codec {
	cdc := wire.New()

	acctypes.RegisterWire(cdc)
	posttypes.RegisterCodec(cdc)
	votetypes.RegisterWire(cdc)
	valtypes.RegisterWire(cdc)
	proposal.RegisterWire(cdc)
	devtypes.RegisterWire(cdc)

	wire.RegisterCrypto(cdc)
	sdk.RegisterCodec(cdc)
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
func EncodeSignMsg(cdc *wire.Codec, msgs []sdk.Msg, chainID string, seq uint64, memo string) ([]byte, error) {
	stdSignMsg := txbuilder.StdSignMsg{
		AccountNumber: 0,
		ChainID:       chainID,
		Fee: auth.StdFee{
			Amount: sdk.Coins{},
			Gas:    0,
		},
		Memo:     memo,
		Msgs:     msgs,
		Sequence: seq,
	}

	signMsgBytes, err := cdc.MarshalJSON(stdSignMsg)
	if err != nil {
		return nil, err
	}

	return sdk.MustSortJSON(signMsgBytes), nil
}

// EncodeTx encodes a message to the standard transaction.
func EncodeTx(
	cdc *wire.Codec, msgs []sdk.Msg, pubKey crypto.PubKey, sig []byte,
	seq uint64, memo string, maxFeeInCoin int64) ([]byte, error) {
	stdSig := auth.StdSignature{
		PubKey:    pubKey,
		Signature: sig,
	}

	stdTx := auth.StdTx{
		Msgs: msgs,
		Fee: auth.StdFee{
			Amount: sdk.NewCoins(
				sdk.NewCoin(linotypes.LinoCoinDenom, sdk.NewInt(maxFeeInCoin))),
			Gas: 0,
		},
		Signatures: []auth.StdSignature{stdSig},
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
