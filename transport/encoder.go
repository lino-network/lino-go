package transport

import (
	"encoding/hex"

	wire "github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/lino-network/lino-go/errors"
	linotypes "github.com/lino-network/lino/types"

	crypto "github.com/tendermint/tendermint/crypto"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
)

// EncodeSignMsg encodes the message to the standard signed message.
func EncodeSignMsg(
	cdc *wire.Codec, msgs []sdk.Msg, chainID string, seq uint64, memo string, maxFeeInCoin int64) []byte {
	return auth.StdSignBytes(
		chainID, 0, seq,
		auth.NewStdFee(0, sdk.NewCoins(
			sdk.NewCoin(linotypes.LinoCoinDenom, sdk.NewInt(maxFeeInCoin))),
		), msgs, memo)
}

// EncodeTx encodes a message to the standard transaction.
func EncodeTx(
	cdc *wire.Codec, msgs []sdk.Msg, pubKey crypto.PubKey, sig []byte,
	seq uint64, memo string, maxFeeInCoin int64) ([]byte, error) {
	stdSig := auth.StdSignature{
		PubKey:    pubKey,
		Signature: sig,
	}

	stdTx := auth.NewStdTx(msgs, auth.NewStdFee(0, sdk.NewCoins(
		sdk.NewCoin(linotypes.LinoCoinDenom, sdk.NewInt(maxFeeInCoin)))), []auth.StdSignature{stdSig}, memo)
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
