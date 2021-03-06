package query

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lino-network/lino-go/errors"
	"github.com/lino-network/lino-go/transport"
	linotypes "github.com/lino-network/lino/types"
	"github.com/lino-network/lino/x/account/model"
	"github.com/lino-network/lino/x/account/types"
	"github.com/tendermint/tendermint/crypto"
)

// GetAccountInfo returns account info for a specific user.
func (query *Query) GetAccountInfo(ctx context.Context, username string) (*model.AccountInfo, error) {
	resp, err := query.transport.Query(ctx, AccountKVStoreKey, types.QueryAccountInfo, []string{username})
	if err != nil {
		linoe, ok := err.(errors.Error)
		if ok && linoe.BlockChainCode() == uint32(linotypes.CodeAccountNotFound) {
			return nil, errors.EmptyResponse("account info is not found")
		}
		return nil, err
	}
	info := new(model.AccountInfo)
	if err := query.transport.Cdc.UnmarshalJSON(resp, info); err != nil {
		return nil, err
	}
	return info, nil
}

// GetTransactionPubKey returns string format transaction public key.
func (query *Query) GetTransactionPubKey(ctx context.Context, username string) (string, error) {
	info, err := query.GetAccountInfo(ctx, username)
	if err != nil {
		return "", err
	}
	return strings.ToUpper(hex.EncodeToString(info.TransactionKey.Bytes())), nil
}

// GetSigningPubKey returns string format signing public key.
func (query *Query) GetSigningPubKey(ctx context.Context, username string) (string, error) {
	info, err := query.GetAccountInfo(ctx, username)
	if err != nil {
		return "", err
	}

	return strings.ToUpper(hex.EncodeToString(info.SigningKey.Bytes())), nil
}

// DoesUsernameMatchSigningPrivKey returns true if a user has the correct signing private key.
func (query *Query) DoesUsernameMatchSigningPrivKey(
	ctx context.Context, username, signingPrivKeyHex string) (bool, error) {
	accInfo, err := query.GetAccountInfo(ctx, username)
	if err != nil {
		return false, err
	}

	signingPrivKey, e := transport.GetPrivKeyFromHex(signingPrivKeyHex)
	if e != nil {
		return false, e
	}

	return accInfo.SigningKey.Equals(signingPrivKey.PubKey()), nil
}

// DoesUsernameMatchTransactionPrivKey returns true if a user has the correct transaction private key.
func (query *Query) DoesUsernameMatchTransactionPrivKey(
	ctx context.Context, username, txPrivKeyHex string) (bool, error) {
	accInfo, err := query.GetAccountInfo(ctx, username)
	if err != nil {
		return false, err
	}

	txPrivKey, e := transport.GetPrivKeyFromHex(txPrivKeyHex)
	if e != nil {
		return false, e
	}

	return accInfo.TransactionKey.Equals(txPrivKey.PubKey()), nil
}

// GetAccountBank returns account bank info for a specific user.
func (query *Query) GetAccountBank(ctx context.Context, username string) (*model.AccountBank, error) {
	resp, err := query.transport.Query(ctx, AccountKVStoreKey, types.QueryAccountBank, []string{username})
	if err != nil {
		linoe, ok := err.(errors.Error)
		if ok && linoe.BlockChainCode() == uint32(linotypes.CodeAccountBankNotFound) {
			return nil, errors.EmptyResponse("account bank is not found")
		}
		return nil, err
	}
	bank := new(model.AccountBank)
	if err := query.transport.Cdc.UnmarshalJSON(resp, bank); err != nil {
		return nil, err
	}
	return bank, nil
}

// GetAccountBankByAddress returns account bank info for a specific address.
func (query *Query) GetAccountBankByAddress(ctx context.Context, address string) (*model.AccountBank, error) {
	addr, e := hex.DecodeString(address)
	if e != nil {
		return nil, errors.InvalidArgf("Address %s is not hex string", address)
	}
	resp, err := query.transport.Query(ctx, AccountKVStoreKey, types.QueryAccountBankByAddress, []string{sdk.AccAddress(addr).String()})
	if err != nil {
		linoe, ok := err.(errors.Error)
		if ok && linoe.BlockChainCode() == uint32(linotypes.CodeAccountBankNotFound) {
			return nil, errors.EmptyResponse("account bank is not found")
		}
		return nil, err
	}
	bank := new(model.AccountBank)
	if err := query.transport.Cdc.UnmarshalJSON(resp, bank); err != nil {
		return nil, err
	}
	return bank, nil
}

// GetAccountMeta returns account meta info for a specific user.
func (query *Query) GetAccountMeta(ctx context.Context, username string) (*model.AccountMeta, error) {
	resp, err := query.transport.Query(ctx, AccountKVStoreKey, types.QueryAccountMeta, []string{username})
	if err != nil {
		return nil, err
	}
	meta := new(model.AccountMeta)
	if err := query.transport.Cdc.UnmarshalJSON(resp, meta); err != nil {
		return nil, err
	}
	return meta, nil
}

// GetSeqNumber returns the next sequence number of a user which should
// be used for broadcast.
func (query *Query) GetSeqNumber(ctx context.Context, username string) (uint64, error) {
	bank, err := query.GetAccountBank(ctx, username)
	if err != nil {
		return 0, err
	}
	return bank.Sequence, nil
}

// GetSeqNumberByAddress returns the next sequence number of an address which should
// be used for broadcast.
func (query *Query) GetSeqNumberByAddress(ctx context.Context, address string) (uint64, error) {
	bank, err := query.GetAccountBankByAddress(ctx, address)
	if err != nil {
		linoe, ok := err.(errors.Error)
		if ok && linoe.CodeType() == errors.CodeEmptyResponse {
			return 0, nil
		}
		return 0, err
	}
	return bank.Sequence, nil
}

// GetReputation returns rewards of a user.
// func (query *Query) GetReputation(ctx context.Context, username string) (*linotypes.Coin, error) {
// 	resp, err := query.transport.Query(ctx, ReputationKVStore, types, []string{username})
// 	if err != nil {
// 		return nil, err
// 	}

// 	reward := new(linotypes.Coin)
// 	if err := query.transport.Cdc.UnmarshalJSON(resp, reward); err != nil {
// 		return reward, err
// 	}
// 	return reward, nil
// }

// GetRewardAtHeight returns rewards of a user at certain height.
// func (query *Query) GetRewardAtHeight(ctx context.Context, username string, height int64) (*model.Reward, error) {
// 	resp, err := query.transport.QueryAtHeight(ctx, getRewardKey(username), AccountKVStoreKey, height)
// 	if err != nil {
// 		switch err.(type) {
// 		case errors.Error:
// 			vErr := err.(errors.Error)
// 			if vErr.CodeType() == errors.CodeEmptyResponse {
// 				return nil, nil
// 			}
// 		}
// 		return nil, err
// 	}

// 	reward := new(model.Reward)
// 	if err := query.transport.Cdc.UnmarshalBinaryLengthPrefixed(resp, reward); err != nil {
// 		return reward, err
// 	}
// 	return reward, nil
// }

//
// Range Query
//

// GetAllFollowingMeta returns all following meta of a user.
func (query *Query) SignWithSha256(ctx context.Context, payload string, privKey crypto.PrivKey) ([]byte, error) {
	hasher := sha256.New()
	hasher.Write([]byte(payload))
	signByte := hasher.Sum(nil)
	return privKey.Sign(signByte)
}

// VerifyUserSignature verify signature is signed from payload by user's signing or tx private key.
func (query *Query) VerifyUserSignature(
	ctx context.Context, username string, payload string, signature string) (bool, error) {
	info, err := query.GetAccountInfo(ctx, username)
	if err != nil {
		return false, err
	}
	sig, err := hex.DecodeString(signature)
	if err != nil {
		return false, err
	}
	if info.SigningKey.VerifyBytes([]byte(payload), sig) || info.TransactionKey.VerifyBytes([]byte(payload), sig) {
		return true, nil
	}
	return false, nil
}

// VerifyUserSignatureUsingSigningKey verify signature is signed from payload by user's app private key.
func (query *Query) VerifyUserSignatureUsingSigningKey(
	ctx context.Context, username string, payload string, signature string) (bool, error) {
	info, err := query.GetAccountInfo(ctx, username)
	if err != nil {
		return false, err
	}
	sig, err := hex.DecodeString(signature)
	if err != nil {
		return false, err
	}
	return info.SigningKey.VerifyBytes([]byte(payload), sig), nil
}

// VerifyUserSignatureUsingTxKey verify signature is signed from payload by user's transaction private key.
func (query *Query) VerifyUserSignatureUsingTxKey(ctx context.Context, username string, payload string, signature string) (bool, error) {
	info, err := query.GetAccountInfo(ctx, username)
	if err != nil {
		return false, err
	}
	sig, err := hex.DecodeString(signature)
	if err != nil {
		return false, err
	}
	return info.TransactionKey.VerifyBytes([]byte(payload), sig), nil
}

// GetTxAndSequenceNumberByUsername get sequence from remote then check transaction is valid or not by username.
func (query *Query) GetTxAndSequenceNumberByUsername(ctx context.Context, username, hash string) (*model.TxAndSequenceNumber, error) {
	resp, err := query.transport.Query(ctx, AccountKVStoreKey, types.QueryTxAndAccountSequence, []string{username, hash, "false"})
	if err != nil {
		return nil, err
	}

	txAndSeq := new(model.TxAndSequenceNumber)
	if err := query.transport.Cdc.UnmarshalJSON(resp, txAndSeq); err != nil {
		return txAndSeq, err
	}
	return txAndSeq, nil
}

// GetTxAndSequenceNumberByAddress get sequence from remote then check transaction is valid or not by address.
func (query *Query) GetTxAndSequenceNumberByAddress(ctx context.Context, address, hash string) (*model.TxAndSequenceNumber, error) {
	addr, e := hex.DecodeString(address)
	if e != nil {
		return nil, errors.InvalidArgf("Address %s is not hex string", address)
	}
	resp, err := query.transport.Query(ctx, AccountKVStoreKey, types.QueryTxAndAccountSequence, []string{sdk.AccAddress(addr).String(), hash, "true"})
	if err != nil {
		return nil, err
	}

	txAndSeq := new(model.TxAndSequenceNumber)
	if err := query.transport.Cdc.UnmarshalJSON(resp, txAndSeq); err != nil {
		return txAndSeq, err
	}
	return txAndSeq, nil
}
