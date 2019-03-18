package query

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/lino-network/lino-go/errors"
	"github.com/lino-network/lino-go/model"
	"github.com/lino-network/lino-go/transport"
	"github.com/tendermint/tendermint/crypto"
)

// GetAccountInfo returns account info for a specific user.
func (query *Query) GetAccountInfo(ctx context.Context, username string) (*model.AccountInfo, error) {
	resp, err := query.transport.Query(ctx, AccountKVStoreKey, AccountInfoSubStore, []string{username})
	if err != nil {
		linoe, ok := err.(errors.Error)
		if ok && linoe.BlockChainCode() == uint32(errors.CodeAccountInfoNotFound) {
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

// GetAppPubKey returns string format app public key.
func (query *Query) GetAppPubKey(ctx context.Context, username string) (string, error) {
	info, err := query.GetAccountInfo(ctx, username)
	if err != nil {
		return "", err
	}

	return strings.ToUpper(hex.EncodeToString(info.AppKey.Bytes())), nil
}

// DoesUsernameMatchResetPrivKey returns true if a user has the reset private key.
func (query *Query) DoesUsernameMatchResetPrivKey(ctx context.Context, username, resetPrivKeyHex string) (bool, error) {
	accInfo, err := query.GetAccountInfo(ctx, username)
	if err != nil {
		return false, err
	}

	resetPrivKey, e := transport.GetPrivKeyFromHex(resetPrivKeyHex)
	if e != nil {
		return false, e
	}

	return accInfo.ResetKey.Equals(resetPrivKey.PubKey()), nil
}

// DoesUsernameMatchTxPrivKey returns true if a user has the transaction private key.
func (query *Query) DoesUsernameMatchTxPrivKey(ctx context.Context, username, txPrivKeyHex string) (bool, error) {
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

// DoesUsernameMatchAppPrivKey returns true if a user has the app private key.
func (query *Query) DoesUsernameMatchAppPrivKey(ctx context.Context, username, appPrivKeyHex string) (bool, error) {
	accInfo, err := query.GetAccountInfo(ctx, username)
	if err != nil {
		return false, err
	}

	appPrivKey, e := transport.GetPrivKeyFromHex(appPrivKeyHex)
	if e != nil {
		return false, e
	}

	return accInfo.AppKey.Equals(appPrivKey.PubKey()), nil
}

// GetAccountBank returns account bank info for a specific user.
func (query *Query) GetAccountBank(ctx context.Context, username string) (*model.AccountBank, error) {
	resp, err := query.transport.Query(ctx, AccountKVStoreKey, AccountBankSubStore, []string{username})
	if err != nil {
		linoe, ok := err.(errors.Error)
		if ok && linoe.BlockChainCode() == uint32(errors.CodeAccountBankNotFound) {
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

// GetAccountBank returns account bank info for a specific user.
func (query *Query) GetPendingCoinDay(ctx context.Context, username string) (*model.PendingCoinDayQueue, error) {
	resp, err := query.transport.Query(ctx, AccountKVStoreKey, AccountPendingCoinDaySubStore, []string{username})
	if err != nil {
		linoe, ok := err.(errors.Error)
		if ok && linoe.BlockChainCode() == uint32(errors.CodePendingCoinDayQueueNotFound) {
			return nil, errors.EmptyResponse("account pending coin day is not found")
		}
		return nil, err
	}
	pendingCoinDay := new(model.PendingCoinDayQueue)
	if err := query.transport.Cdc.UnmarshalJSON(resp, pendingCoinDay); err != nil {
		return nil, err
	}
	return pendingCoinDay, nil
}

// GetAccountMeta returns account meta info for a specific user.
func (query *Query) GetAccountMeta(ctx context.Context, username string) (*model.AccountMeta, error) {
	resp, err := query.transport.Query(ctx, AccountKVStoreKey, AccountMetaSubStore, []string{username})
	if err != nil {
		linoe, ok := err.(errors.Error)
		if ok && linoe.BlockChainCode() == uint32(errors.CodeAccountMetaNotFound) {
			return nil, errors.EmptyResponse("account meta is not found")
		}
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
	meta, err := query.GetAccountMeta(ctx, username)
	if err != nil {
		return 0, err
	}
	return meta.Sequence, nil
}

// GetGrantPubKey returns the specific granted pubkey info of a user
// that has given to the pubKey.
func (query *Query) GetGrantPubKey(ctx context.Context, username string, grantTo string, permission model.Permission) (*model.GrantPubKey, error) {
	resp, err := query.transport.Query(ctx, AccountKVStoreKey, AccountGrantPubKeySubStore, []string{username, grantTo})
	if err != nil {
		return nil, errors.EmptyResponse("grant pubkey is not found or err")
	}

	grantPubKeyList := make([]*model.GrantPubKey, 0)
	if err := query.transport.Cdc.UnmarshalJSON(resp, &grantPubKeyList); err != nil {
		return nil, err
	}
	for _, grantPubKey := range grantPubKeyList {
		if grantPubKey.Permission == permission {
			return grantPubKey, nil
		}
	}
	return nil, errors.EmptyResponse("grant pubkey is not found")
}

// GetReward returns rewards of a user.
func (query *Query) GetReward(ctx context.Context, username string) (*model.Reward, error) {
	resp, err := query.transport.Query(ctx, AccountKVStoreKey, AccountRewardSubStore, []string{username})
	if err != nil {
		linoe, ok := err.(errors.Error)
		if ok && linoe.BlockChainCode() == uint32(errors.CodeRewardNotFound) {
			return nil, errors.EmptyResponse("account reward is not found")
		}
		return nil, err
	}

	reward := new(model.Reward)
	if err := query.transport.Cdc.UnmarshalJSON(resp, reward); err != nil {
		return reward, err
	}
	return reward, nil
}

// GetReward returns rewards of a user.
func (query *Query) GetReputation(ctx context.Context, username string) (*model.Coin, error) {
	resp, err := query.transport.Query(ctx, ReputationKVStore, UserReputation, []string{username})
	if err != nil {
		return nil, err
	}

	reward := new(model.Coin)
	if err := query.transport.Cdc.UnmarshalJSON(resp, reward); err != nil {
		return reward, err
	}
	return reward, nil
}

// GetRewardAtHeight returns rewards of a user at certain height.
func (query *Query) GetRewardAtHeight(ctx context.Context, username string, height int64) (*model.Reward, error) {
	resp, err := query.transport.QueryAtHeight(ctx, getRewardKey(username), AccountKVStoreKey, height)
	if err != nil {
		switch err.(type) {
		case errors.Error:
			vErr := err.(errors.Error)
			if vErr.CodeType() == errors.CodeEmptyResponse {
				return nil, nil
			}
		}
		return nil, err
	}

	reward := new(model.Reward)
	if err := query.transport.Cdc.UnmarshalBinaryLengthPrefixed(resp, reward); err != nil {
		return reward, err
	}
	return reward, nil
}

//
// Range Query
//

// GetAllGrantPubKeys returns a list of all granted public keys of a user.
func (query *Query) GetAllGrantPubKeys(ctx context.Context, username string) ([]*model.GrantPubKey, error) {
	resp, err := query.transport.Query(ctx, AccountKVStoreKey, AccountAllGrantPubKeys, []string{username})
	if err != nil {
		return nil, errors.EmptyResponse("grant pub key is not found")
	}
	grantPubKeyList := make([]*model.GrantPubKey, 0)
	if err := query.transport.Cdc.UnmarshalJSON(resp, &grantPubKeyList); err != nil {
		return nil, err
	}

	return grantPubKeyList, nil
}

// GetAllFollowingMeta returns all following meta of a user.
func (query *Query) SignWithSha256(ctx context.Context, payload string, privKey crypto.PrivKey) ([]byte, error) {
	hasher := sha256.New()
	hasher.Write([]byte(payload))
	signByte := hasher.Sum(nil)
	return privKey.Sign(signByte)
}

// VerifyUserSignatureUsingAppKey verify signature is signed from payload by user's app private key.
func (query *Query) VerifyUserSignatureUsingAppKey(ctx context.Context, username string, payload string, signature string) (bool, error) {
	info, err := query.GetAccountInfo(ctx, username)
	if err != nil {
		return false, err
	}
	sig, err := hex.DecodeString(signature)
	if err != nil {
		return false, err
	}
	return info.AppKey.VerifyBytes([]byte(payload), sig), nil
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

// GetTxAndSequenceNumber verify signature is signed from payload by user's transaction private key.
func (query *Query) GetTxAndSequenceNumber(ctx context.Context, username, hash string) (*model.TxAndSequenceNumber, error) {
	resp, err := query.transport.Query(ctx, AccountKVStoreKey, AccountTxAndSequence, []string{username, hash})
	if err != nil {
		return nil, err
	}

	txAndSeq := new(model.TxAndSequenceNumber)
	if err := query.transport.Cdc.UnmarshalJSON(resp, txAndSeq); err != nil {
		return txAndSeq, err
	}
	return txAndSeq, nil
}
