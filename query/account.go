package query

import (
	"encoding/hex"
	"math"
	"strings"

	"github.com/lino-network/lino-go/errors"
	"github.com/lino-network/lino-go/model"
	"github.com/lino-network/lino-go/transport"
)

// GetAccountInfo returns account info for a specific user.
func (query *Query) GetAccountInfo(username string) (*model.AccountInfo, error) {
	resp, err := query.transport.Query(getAccountInfoKey(username), AccountKVStoreKey)
	if err != nil {
		return nil, err
	}
	info := new(model.AccountInfo)
	if err := query.transport.Cdc.UnmarshalJSON(resp, info); err != nil {
		return nil, err
	}
	return info, nil
}

// GetTransactionPubKey returns string format transaction public key.
func (query *Query) GetTransactionPubKey(username string) (string, error) {
	resp, err := query.transport.Query(getAccountInfoKey(username), AccountKVStoreKey)
	if err != nil {
		return "", err
	}
	info := new(model.AccountInfo)
	if err := query.transport.Cdc.UnmarshalJSON(resp, info); err != nil {
		return "", err
	}
	return strings.ToUpper(hex.EncodeToString(info.TransactionKey.Bytes())), nil
}

// GetMicropaymentPubKey returns string format micropayment public key.
func (query *Query) GetMicropaymentPubKey(username string) (string, error) {
	resp, err := query.transport.Query(getAccountInfoKey(username), AccountKVStoreKey)
	if err != nil {
		return "", err
	}
	info := new(model.AccountInfo)
	if err := query.transport.Cdc.UnmarshalJSON(resp, info); err != nil {
		return "", err
	}
	return strings.ToUpper(hex.EncodeToString(info.MicropaymentKey.Bytes())), nil
}

// GetPostPubKey returns string format post public key.
func (query *Query) GetPostPubKey(username string) (string, error) {
	resp, err := query.transport.Query(getAccountInfoKey(username), AccountKVStoreKey)
	if err != nil {
		return "", err
	}
	info := new(model.AccountInfo)
	if err := query.transport.Cdc.UnmarshalJSON(resp, info); err != nil {
		return "", err
	}
	return strings.ToUpper(hex.EncodeToString(info.PostKey.Bytes())), nil
}

// DoesUsernameMatchResetPrivKey returns true if a user has the reset private key.
func (query *Query) DoesUsernameMatchResetPrivKey(username, resetPrivKeyHex string) (bool, error) {
	accInfo, err := query.GetAccountInfo(username)
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
func (query *Query) DoesUsernameMatchTxPrivKey(username, txPrivKeyHex string) (bool, error) {
	accInfo, err := query.GetAccountInfo(username)
	if err != nil {
		return false, err
	}

	txPrivKey, e := transport.GetPrivKeyFromHex(txPrivKeyHex)
	if e != nil {
		return false, e
	}

	return accInfo.TransactionKey.Equals(txPrivKey.PubKey()), nil
}

// DoesUsernameMatchMicropaymentPrivKey returns true if a user has the micropayment private key.
func (query *Query) DoesUsernameMatchMicropaymentPrivKey(username, micropaymentPrivKeyHex string) (bool, error) {
	accInfo, err := query.GetAccountInfo(username)
	if err != nil {
		return false, err
	}

	txPrivKey, e := transport.GetPrivKeyFromHex(micropaymentPrivKeyHex)
	if e != nil {
		return false, e
	}

	return accInfo.MicropaymentKey.Equals(txPrivKey.PubKey()), nil
}

// DoesUsernameMatchPostPrivKey returns true if a user has the post private key.
func (query *Query) DoesUsernameMatchPostPrivKey(username, postPrivKeyHex string) (bool, error) {
	accInfo, err := query.GetAccountInfo(username)
	if err != nil {
		return false, err
	}

	postPrivKey, e := transport.GetPrivKeyFromHex(postPrivKeyHex)
	if e != nil {
		return false, e
	}

	return accInfo.PostKey.Equals(postPrivKey.PubKey()), nil
}

// GetAccountBank returns account bank info for a specific user.
func (query *Query) GetAccountBank(username string) (*model.AccountBank, error) {
	resp, err := query.transport.Query(getAccountBankKey(username), AccountKVStoreKey)
	if err != nil {
		return nil, err
	}
	bank := new(model.AccountBank)
	if err := query.transport.Cdc.UnmarshalJSON(resp, bank); err != nil {
		return nil, err
	}
	return bank, nil
}

// GetAccountMeta returns account meta info for a specific user.
func (query *Query) GetAccountMeta(username string) (*model.AccountMeta, error) {
	resp, err := query.transport.Query(getAccountMetaKey(username), AccountKVStoreKey)
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
func (query *Query) GetSeqNumber(username string) (int64, error) {
	meta, err := query.GetAccountMeta(username)
	if err != nil {
		return 0, err
	}
	return meta.Sequence, nil
}

// GetAllBalanceHistory returns all transaction history related to
// a user's account balance, in reverse-chronological order.
func (query *Query) GetAllBalanceHistory(username string) (*model.BalanceHistory, error) {
	accountBank, err := query.GetAccountBank(username)
	if err != nil {
		return nil, err
	}

	allBalanceHistory := new(model.BalanceHistory)
	if accountBank.NumOfTx == 0 {
		return allBalanceHistory, nil
	}

	bucketSlot := (accountBank.NumOfTx - 1) / 100

	for i := bucketSlot; i >= 0; i-- {
		balanceHistory, err := query.GetBalanceHistory(username, i)
		if err != nil {
			return nil, err
		}

		for index := len(balanceHistory.Details) - 1; index >= 0; index-- {
			allBalanceHistory.Details = append(allBalanceHistory.Details, balanceHistory.Details[index])
		}
	}

	return allBalanceHistory, nil
}

// GetRecentBalanceHistory returns a certain number of recent transaction history
// related to a user's account balance, in reverse-chronological order.
func (query *Query) GetRecentBalanceHistory(username string, numHistory int64) (*model.BalanceHistory, error) {
	if numHistory <= 0 || numHistory > math.MaxInt64 {
		return nil, errors.InvalidArgf("GetRecentBalanceHistory: numHistory is invalid: %v", numHistory)
	}

	accountBank, err := query.GetAccountBank(username)
	if err != nil {
		return nil, err
	}

	allBalanceHistory := new(model.BalanceHistory)
	if accountBank.NumOfTx == 0 {
		return allBalanceHistory, nil
	}

	from := accountBank.NumOfTx - numHistory
	if numHistory > accountBank.NumOfTx {
		from = 0
	}

	to := accountBank.NumOfTx - 1

	return query.GetBalanceHistoryFromTo(username, from, to)
}

// GetBalanceHistoryFromTo returns a list of transaction history in the range of index [from, to]
// related to a user's account balance, in reverse-chronological order.
func (query *Query) GetBalanceHistoryFromTo(username string, from, to int64) (*model.BalanceHistory, error) {
	if from < 0 || from > math.MaxInt64 || to < 0 || to > math.MaxInt64 || from > to {
		return nil, errors.InvalidArgf("GetBalanceHistoryFromTo: from [%v] or to [%v] is invalid", from, to)
	}

	accountBank, err := query.GetAccountBank(username)
	if err != nil {
		return nil, err
	}

	allBalanceHistory := new(model.BalanceHistory)
	if accountBank.NumOfTx == 0 {
		return allBalanceHistory, nil
	}

	if from > accountBank.NumOfTx-1 {
		return allBalanceHistory, errors.InvalidArgf("GetBalanceHistoryFromTo: invalid from [%v] which bigger than total num of tx", from)
	}
	if to > accountBank.NumOfTx-1 {
		to = accountBank.NumOfTx - 1
	}

	// number of banlance history is wanted
	numHistory := to - from + 1

	targetBucketOfTo := to / 100
	bucketSlot := targetBucketOfTo

	// The index of 'to' in the target bucket
	indexOfTo := to % 100

	for bucketSlot > -1 {
		balanceHistory, err := query.GetBalanceHistory(username, bucketSlot)
		if err != nil {
			return nil, err
		}

		var startIndex int64
		if bucketSlot == targetBucketOfTo {
			startIndex = indexOfTo
		} else {
			startIndex = int64(len(balanceHistory.Details) - 1)
		}

		for i := startIndex; i >= 0 && numHistory > 0; i-- {
			allBalanceHistory.Details = append(allBalanceHistory.Details, balanceHistory.Details[i])
			numHistory--
		}

		if numHistory == 0 {
			break
		}

		bucketSlot--
	}

	return allBalanceHistory, nil
}

// GetBalanceHistory returns all balance history in a certain bucket.
func (query *Query) GetBalanceHistory(username string, index int64) (*model.BalanceHistory, error) {
	resp, err := query.transport.Query(getBalanceHistoryKey(username, index), AccountKVStoreKey)
	if err != nil {
		return nil, err
	}
	balanceHistory := new(model.BalanceHistory)
	if err := query.transport.Cdc.UnmarshalJSON(resp, balanceHistory); err != nil {
		return nil, err
	}
	return balanceHistory, nil
}

// GetGrantPubKey returns the specific granted pubkey info of a user
// that has given to the pubKey.
func (query *Query) GetGrantPubKey(username string, pubKeyHex string) (*model.GrantPubKey, error) {
	pubKey, err := transport.GetPubKeyFromHex(pubKeyHex)
	if err != nil {
		return nil, errors.FailedToGetPubKeyFromHex("GetGrantPubKey: failed to get pub key").AddCause(err)
	}

	resp, err := query.transport.Query(getGrantPubKeyKey(username, pubKey), AccountKVStoreKey)
	if err != nil {
		return nil, err
	}

	grantPubKey := new(model.GrantPubKey)
	if err := query.transport.Cdc.UnmarshalJSON(resp, grantPubKey); err != nil {
		return grantPubKey, err
	}
	return grantPubKey, nil
}

// GetReward returns rewards of a user.
func (query *Query) GetReward(username string) (*model.Reward, error) {
	resp, err := query.transport.Query(getRewardKey(username), AccountKVStoreKey)
	if err != nil {
		return nil, err
	}

	reward := new(model.Reward)
	if err := query.transport.Cdc.UnmarshalJSON(resp, reward); err != nil {
		return reward, err
	}
	return reward, nil
}

// GetAllRewardHistory returns all reward history related to
// a user's posts reward, in reverse-chronological order.
func (query *Query) GetAllRewardHistory(username string) (*model.RewardHistory, error) {
	accountBank, err := query.GetAccountBank(username)
	if err != nil {
		return nil, err
	}

	allRewardHistory := new(model.RewardHistory)
	if accountBank.NumOfReward == 0 {
		return allRewardHistory, nil
	}

	bucketSlot := (accountBank.NumOfReward - 1) / 100

	for i := bucketSlot; i >= 0; i-- {
		rewardHistory, err := query.GetRewardHistory(username, i)
		if err != nil {
			return nil, err
		}

		for index := len(rewardHistory.Details) - 1; index >= 0; index-- {
			allRewardHistory.Details = append(allRewardHistory.Details, rewardHistory.Details[index])
		}
	}

	return allRewardHistory, nil
}

// GetRecentRewardHistory returns a certain number of recent reward history
// related to a user's posts reward, in reverse-chronological order.
func (query *Query) GetRecentRewardHistory(username string, numReward int64) (*model.RewardHistory, error) {
	if numReward <= 0 || numReward > math.MaxInt64 {
		return nil, errors.InvalidArgf("GetRecentRewardHistory: numReward is invalid: %v", numReward)
	}

	accountBank, err := query.GetAccountBank(username)
	if err != nil {
		return nil, err
	}

	allRewardHistory := new(model.RewardHistory)
	if accountBank.NumOfReward == 0 {
		return allRewardHistory, nil
	}

	from := accountBank.NumOfReward - numReward
	if numReward > accountBank.NumOfReward {
		from = 0
	}

	to := accountBank.NumOfReward - 1

	return query.GetRewardHistoryFromTo(username, from, to)
}

// GetRewardHistoryFromTo returns a list of reward history in the range of index [from, to]
// related to a user's posts reward, in reverse-chronological order.
func (query *Query) GetRewardHistoryFromTo(username string, from, to int64) (*model.RewardHistory, error) {
	if from < 0 || from > math.MaxInt64 || to < 0 || to > math.MaxInt64 || from > to {
		return nil, errors.InvalidArgf("GetRewardHistoryFromTo: from [%v] or to [%v] is invalid", from, to)
	}

	accountBank, err := query.GetAccountBank(username)
	if err != nil {
		return nil, err
	}

	allRewardHistory := new(model.RewardHistory)
	if accountBank.NumOfReward == 0 {
		return allRewardHistory, nil
	}

	if from > accountBank.NumOfReward-1 {
		return allRewardHistory, errors.InvalidArgf("GetRewardHistoryFromTo: invalid from [%v] which is bigger than total num of reward", from)
	}
	if to > accountBank.NumOfReward-1 {
		to = accountBank.NumOfReward - 1
	}

	// number of reward history is wanted
	numReward := to - from + 1

	targetBucketOfTo := to / 100
	bucketSlot := targetBucketOfTo

	// The index of 'to' in the target bucket
	indexOfTo := to % 100

	for bucketSlot > -1 {
		rewardHistory, err := query.GetRewardHistory(username, bucketSlot)
		if err != nil {
			return nil, err
		}

		var startIndex int64
		if bucketSlot == targetBucketOfTo {
			startIndex = indexOfTo
		} else {
			startIndex = int64(len(rewardHistory.Details) - 1)
		}

		for i := startIndex; i >= 0 && numReward > 0; i-- {
			allRewardHistory.Details = append(allRewardHistory.Details, rewardHistory.Details[i])
			numReward--
		}

		if numReward == 0 {
			break
		}

		bucketSlot--
	}

	return allRewardHistory, nil
}

// GetRewardHistory returns all reward history in a certain bucket
func (query *Query) GetRewardHistory(username string, index int64) (*model.RewardHistory, error) {
	resp, err := query.transport.Query(getRewardHistoryKey(username, index), AccountKVStoreKey)
	if err != nil {
		return nil, err
	}
	rewardHistory := new(model.RewardHistory)
	if err := query.transport.Cdc.UnmarshalJSON(resp, rewardHistory); err != nil {
		return nil, err
	}
	return rewardHistory, nil
}

// GetRelationship returns the donation times of two users.
func (query *Query) GetRelationship(me, other string) (*model.Relationship, error) {
	resp, err := query.transport.Query(getRelationshipKey(me, other), AccountKVStoreKey)
	if err != nil {
		return nil, err
	}

	relationship := new(model.Relationship)
	if err := query.transport.Cdc.UnmarshalJSON(resp, relationship); err != nil {
		return relationship, err
	}
	return relationship, nil
}

// GetFollowerMeta returns the follower meta of two users.
func (query *Query) GetFollowerMeta(me, myFollower string) (*model.FollowerMeta, error) {
	resp, err := query.transport.Query(getFollowerKey(me, myFollower), AccountKVStoreKey)
	if err != nil {
		return nil, err
	}

	followerMeta := new(model.FollowerMeta)
	if err := query.transport.Cdc.UnmarshalJSON(resp, followerMeta); err != nil {
		return followerMeta, err
	}
	return followerMeta, nil
}

// GetFollowingMeta returns the following meta of two users.
func (query *Query) GetFollowingMeta(me, myFollowing string) (*model.FollowingMeta, error) {
	resp, err := query.transport.Query(getFollowerKey(me, myFollowing), AccountKVStoreKey)
	if err != nil {
		return nil, err
	}

	followingMeta := new(model.FollowingMeta)
	if err := query.transport.Cdc.UnmarshalJSON(resp, followingMeta); err != nil {
		return followingMeta, err
	}
	return followingMeta, nil
}

//
// Range Query
//

// GetAllGrantPubKeys returns a list of all granted public keys of a user.
func (query *Query) GetAllGrantPubKeys(username string) (map[string]*model.GrantPubKey, error) {
	resKVs, err := query.transport.QuerySubspace(getGrantPubKeyPrefix(username), AccountKVStoreKey)
	if err != nil {
		return nil, err
	}
	pubKeyToGrantPubKeyMap := make(map[string]*model.GrantPubKey)
	for _, KV := range resKVs {
		grantPubKey := new(model.GrantPubKey)
		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, grantPubKey); err != nil {
			return nil, err
		}
		pubKeyToGrantPubKeyMap[getHexSubstringAfterKeySeparator(KV.Key)] = grantPubKey
	}

	return pubKeyToGrantPubKeyMap, nil
}

// GetAllRelationships returns all donation relationship of a user.
func (query *Query) GetAllRelationships(username string) (map[string]*model.Relationship, error) {
	resKVs, err := query.transport.QuerySubspace(getRelationshipPrefix(username), AccountKVStoreKey)
	if err != nil {
		return nil, err
	}

	userToRelationshipMap := make(map[string]*model.Relationship)
	for _, KV := range resKVs {
		relationship := new(model.Relationship)
		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, relationship); err != nil {
			return nil, err
		}
		userToRelationshipMap[getSubstringAfterKeySeparator(KV.Key)] = relationship
	}

	return userToRelationshipMap, nil
}

// GetAllFollowerMeta returns all follower meta of a user.
func (query *Query) GetAllFollowerMeta(username string) (map[string]*model.FollowerMeta, error) {
	resKVs, err := query.transport.QuerySubspace(getFollowerPrefix(username), AccountKVStoreKey)
	if err != nil {
		return nil, err
	}

	followerToMetaMap := make(map[string]*model.FollowerMeta)
	for _, KV := range resKVs {
		followerMeta := new(model.FollowerMeta)
		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, followerMeta); err != nil {
			return nil, err
		}
		followerToMetaMap[getSubstringAfterKeySeparator(KV.Key)] = followerMeta
	}

	return followerToMetaMap, nil
}

// GetAllFollowingMeta returns all following meta of a user.
func (query *Query) GetAllFollowingMeta(username string) (map[string]*model.FollowingMeta, error) {
	resKVs, err := query.transport.QuerySubspace(getFollowingPrefix(username), AccountKVStoreKey)
	if err != nil {
		return nil, err
	}

	followingMetas := make(map[string]*model.FollowingMeta)
	for _, KV := range resKVs {
		followingMeta := new(model.FollowingMeta)
		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, followingMeta); err != nil {
			return nil, err
		}
		followingMetas[getSubstringAfterKeySeparator(KV.Key)] = followingMeta
	}

	return followingMetas, nil
}
