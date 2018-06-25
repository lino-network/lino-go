package query

import (
	"math"

	"github.com/lino-network/lino-go/errors"
	"github.com/lino-network/lino-go/model"
	"github.com/lino-network/lino-go/transport"

	crypto "github.com/tendermint/go-crypto"
)

type Query struct {
	transport *transport.Transport
}

func NewQuery(transport *transport.Transport) *Query {
	return &Query{
		transport: transport,
	}
}

//
// Account related query
//
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

func (query *Query) DoesUsernameMatchMasterPrivKey(username, masterPrivKeyHex string) (bool, error) {
	accInfo, err := query.GetAccountInfo(username)
	if err != nil {
		return false, err
	}

	masterPrivKey, e := transport.GetPrivKeyFromHex(masterPrivKeyHex)
	if e != nil {
		return false, e
	}

	return accInfo.MasterKey.Equals(masterPrivKey.PubKey()), nil
}

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

func (query *Query) GetSeqNumber(username string) (int64, error) {
	meta, err := query.GetAccountMeta(username)
	if err != nil {
		return 0, err
	}
	return meta.Sequence, nil
}

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

	bucketSlot := (accountBank.NumOfTx - 1) / 100

	for bucketSlot > -1 {
		balanceHistory, err := query.GetBalanceHistory(username, bucketSlot)
		if err != nil {
			return nil, err
		}
		for i := len(balanceHistory.Details) - 1; i >= 0 && numHistory > 0; i-- {
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

	if from > accountBank.NumOfTx {
		return allBalanceHistory, errors.InvalidArgf("GetBalanceHistoryFromTo: invalid from [%v] which bigger than total num of tx", from)
	}
	if to > accountBank.NumOfTx {
		to = accountBank.NumOfTx
	}

	// number of banlance history is wanted
	numHistory := to - from + 1

	targetBucketOfTo := (to - 1) / 100
	bucketSlot := targetBucketOfTo

	// The index of 'to' in the target bucket
	indexOfTo := (to - 1) % 100

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

func (query *Query) GetGrantUser(username string, pubKey crypto.PubKey) (*model.GrantUser, error) {
	resp, err := query.transport.Query(getGrantUserKey(username, pubKey), AccountKVStoreKey)
	if err != nil {
		return nil, err
	}

	grantUser := new(model.GrantUser)
	if err := query.transport.Cdc.UnmarshalJSON(resp, grantUser); err != nil {
		return grantUser, err
	}
	return grantUser, nil
}

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
// Post related query
//

func (query *Query) GetPostInfo(author, postID string) (*model.PostInfo, error) {
	postKey := getPostKey(author, postID)
	resp, err := query.transport.Query(getPostInfoKey(postKey), PostKVStoreKey)
	if err != nil {
		return nil, err
	}
	postInfo := new(model.PostInfo)
	if err := query.transport.Cdc.UnmarshalJSON(resp, postInfo); err != nil {
		return nil, err
	}
	return postInfo, nil
}

func (query *Query) GetPostMeta(author, postID string) (*model.PostMeta, error) {
	postKey := getPostKey(author, postID)
	resp, err := query.transport.Query(getPostMetaKey(postKey), PostKVStoreKey)
	if err != nil {
		return nil, err
	}
	postMeta := new(model.PostMeta)
	if err := query.transport.Cdc.UnmarshalJSON(resp, postMeta); err != nil {
		return nil, err
	}
	return postMeta, nil
}

func (query *Query) GetPostComment(author, postID, commentPostKey string) (*model.Comment, error) {
	postKey := getPostKey(author, postID)
	resp, err := query.transport.Query(getPostCommentKey(postKey, commentPostKey), PostKVStoreKey)
	if err != nil {
		return nil, err
	}
	comment := new(model.Comment)
	if err := query.transport.Cdc.UnmarshalJSON(resp, comment); err != nil {
		return nil, err
	}
	return comment, nil
}

func (query *Query) GetPostView(author, postID, viewUser string) (*model.View, error) {
	postKey := getPostKey(author, postID)
	resp, err := query.transport.Query(getPostViewKey(postKey, viewUser), PostKVStoreKey)
	if err != nil {
		return nil, err
	}
	view := new(model.View)
	if err := query.transport.Cdc.UnmarshalJSON(resp, view); err != nil {
		return nil, err
	}
	return view, nil
}

func (query *Query) GetPostDonation(author, postID, donateUser string) (*model.Donation, error) {
	postKey := getPostKey(author, postID)
	resp, err := query.transport.Query(getPostDonationKey(postKey, donateUser), PostKVStoreKey)
	if err != nil {
		return nil, err
	}
	donation := new(model.Donation)
	if err := query.transport.Cdc.UnmarshalJSON(resp, donation); err != nil {
		return nil, err
	}
	return donation, nil
}

func (query *Query) GetPostReportOrUpvote(author, postID, user string) (*model.ReportOrUpvote, error) {
	postKey := getPostKey(author, postID)
	resp, err := query.transport.Query(getPostReportOrUpvoteKey(postKey, user), PostKVStoreKey)
	if err != nil {
		return nil, err
	}
	reportOrUpvote := new(model.ReportOrUpvote)
	if err := query.transport.Cdc.UnmarshalJSON(resp, reportOrUpvote); err != nil {
		return nil, err
	}
	return reportOrUpvote, nil
}

func (query *Query) GetPostLike(author, postID, likeUser string) (*model.Like, error) {
	postKey := getPostKey(author, postID)
	resp, err := query.transport.Query(getPostLikeKey(postKey, likeUser), PostKVStoreKey)
	if err != nil {
		return nil, err
	}
	like := new(model.Like)
	if err := query.transport.Cdc.UnmarshalJSON(resp, like); err != nil {
		return nil, err
	}
	return like, nil
}

//
// Validator related query
//
func (query *Query) GetValidator(username string) (*model.Validator, error) {
	resp, err := query.transport.Query(getValidatorKey(username), ValidatorKVStoreKey)
	if err != nil {
		return nil, err
	}
	validator := new(model.Validator)
	if err := query.transport.Cdc.UnmarshalJSON(resp, validator); err != nil {
		return nil, err
	}
	return validator, nil
}

func (query *Query) GetAllValidators() (*model.ValidatorList, error) {
	resp, err := query.transport.Query(getValidatorListKey(), ValidatorKVStoreKey)
	if err != nil {
		return nil, err
	}

	validatorList := new(model.ValidatorList)
	if err := query.transport.Cdc.UnmarshalJSON(resp, validatorList); err != nil {
		return validatorList, err
	}
	return validatorList, nil
}

//
// Vote related query
//

func (query *Query) GetDelegation(voter, delegator string) (*model.Delegation, error) {
	resp, err := query.transport.Query(getDelegationKey(voter, delegator), VoteKVStoreKey)
	if err != nil {
		return nil, err
	}
	delegation := new(model.Delegation)
	if err := query.transport.Cdc.UnmarshalJSON(resp, delegation); err != nil {
		return nil, err
	}
	return delegation, nil
}

func (query *Query) GetVoter(voterName string) (*model.Voter, error) {
	resp, err := query.transport.Query(getVoterKey(voterName), VoteKVStoreKey)
	if err != nil {
		return nil, err
	}
	voter := new(model.Voter)
	if err := query.transport.Cdc.UnmarshalJSON(resp, voter); err != nil {
		return nil, err
	}
	return voter, nil
}

func (query *Query) GetDelegateeList(delegatorName string) (*model.DelegateeList, error) {
	resp, err := query.transport.Query(GetDelegateeListKey(delegatorName), VoteKVStoreKey)
	if err != nil {
		return nil, err
	}
	delegateeList := new(model.DelegateeList)
	if err := query.transport.Cdc.UnmarshalJSON(resp, delegateeList); err != nil {
		return nil, err
	}
	return delegateeList, nil
}

func (query *Query) GetVote(proposalID, voter string) (*model.Vote, error) {
	resp, err := query.transport.Query(getVoteKey(proposalID, voter), VoteKVStoreKey)
	if err != nil {
		return nil, err
	}
	vote := new(model.Vote)
	if err := query.transport.Cdc.UnmarshalJSON(resp, vote); err != nil {
		return nil, err
	}
	return vote, nil
}

//
// Developer related query
//
func (query *Query) GetDeveloper(developerName string) (*model.Developer, error) {
	resp, err := query.transport.Query(getDeveloperKey(developerName), DeveloperKVStoreKey)
	if err != nil {
		return nil, err
	}
	developer := new(model.Developer)
	if err := query.transport.Cdc.UnmarshalJSON(resp, developer); err != nil {
		return nil, err
	}
	return developer, nil
}

func (query *Query) GetDevelopers() (*model.DeveloperList, error) {
	resp, err := query.transport.Query(getDeveloperListKey(), DeveloperKVStoreKey)
	if err != nil {
		return nil, err
	}

	developerList := new(model.DeveloperList)
	if err := query.transport.Cdc.UnmarshalJSON(resp, developerList); err != nil {
		return nil, err
	}
	return developerList, nil
}

//
// Infra related query
//
func (query *Query) GetInfraProvider(providerName string) (*model.InfraProvider, error) {
	resp, err := query.transport.Query(getInfraProviderKey(providerName), InfraKVStoreKey)
	if err != nil {
		return nil, err
	}
	provider := new(model.InfraProvider)
	if err := query.transport.Cdc.UnmarshalJSON(resp, provider); err != nil {
		return nil, err
	}
	return provider, nil
}

func (query *Query) GetInfraProviders() (*model.InfraProviderList, error) {
	resp, err := query.transport.Query(getInfraProviderListKey(), InfraKVStoreKey)
	if err != nil {
		return nil, err
	}

	providerList := new(model.InfraProviderList)
	if err := query.transport.Cdc.UnmarshalJSON(resp, providerList); err != nil {
		return nil, err
	}
	return providerList, nil
}

//
// proposal related query
//
func (query *Query) GetProposalList() (*model.ProposalList, error) {
	resp, err := query.transport.Query(getProposalListKey(), ProposalKVStoreKey)
	if err != nil {
		return nil, err
	}

	proposalList := new(model.ProposalList)
	if err := query.transport.Cdc.UnmarshalJSON(resp, proposalList); err != nil {
		return nil, err
	}
	return proposalList, nil
}

func (query *Query) GetProposal(proposalID string) (*model.Proposal, error) {
	resp, err := query.transport.Query(getProposalKey(proposalID), ProposalKVStoreKey)
	if err != nil {
		return nil, err
	}

	proposal := new(model.Proposal)
	if err := query.transport.Cdc.UnmarshalJSON(resp, proposal); err != nil {
		return nil, err
	}
	return proposal, nil
}

func (query *Query) GetOngoingProposal() ([]*model.Proposal, error) {
	proposalList, err := query.GetProposalList()
	if err != nil {
		return nil, err
	}

	var ongoingProposals []*model.Proposal
	for _, proposalID := range proposalList.OngoingProposal {
		p, err := query.GetProposal(proposalID)
		if err != nil {
			return nil, err
		}

		ongoingProposals = append(ongoingProposals, p)
	}

	return ongoingProposals, nil
}

func (query *Query) GetExpiredProposal() ([]*model.Proposal, error) {
	proposalList, err := query.GetProposalList()
	if err != nil {
		return nil, err
	}

	var expiredProposals []*model.Proposal
	for _, proposalID := range proposalList.PastProposal {
		p, err := query.GetProposal(proposalID)
		if err != nil {
			return nil, err
		}

		expiredProposals = append(expiredProposals, p)
	}

	return expiredProposals, nil
}

//
// param related query
//
func (query *Query) GetEvaluateOfContentValueParam() (*model.EvaluateOfContentValueParam, error) {
	resp, err := query.transport.Query(getEvaluateOfContentValueParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.EvaluateOfContentValueParam)
	// fmt.Println("---paramBytes: ", resp)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

func (query *Query) GetGlobalAllocationParam() (*model.GlobalAllocationParam, error) {
	resp, err := query.transport.Query(getGlobalAllocationParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.GlobalAllocationParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

func (query *Query) GetInfraInternalAllocationParam() (*model.InfraInternalAllocationParam, error) {
	resp, err := query.transport.Query(getInfraInternalAllocationParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.InfraInternalAllocationParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

func (query *Query) GetDeveloperParam() (*model.DeveloperParam, error) {
	resp, err := query.transport.Query(getDeveloperParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.DeveloperParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

func (query *Query) GetVoteParam() (*model.VoteParam, error) {
	resp, err := query.transport.Query(getVoteParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.VoteParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

func (query *Query) GetProposalParam() (*model.ProposalParam, error) {
	resp, err := query.transport.Query(getProposalParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.ProposalParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

func (query *Query) GetValidatorParam() (*model.ValidatorParam, error) {
	resp, err := query.transport.Query(getValidatorParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.ValidatorParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

func (query *Query) GetCoinDayParam() (*model.CoinDayParam, error) {
	resp, err := query.transport.Query(getCoinDayParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.CoinDayParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

func (query *Query) GetBandwidthParam() (*model.BandwidthParam, error) {
	resp, err := query.transport.Query(getBandwidthParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.BandwidthParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

func (query *Query) GetAccountParam() (*model.AccountParam, error) {
	resp, err := query.transport.Query(getAccountParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.AccountParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

//
// get block
//
func (query *Query) GetBlock(height int64) (*model.Block, error) {
	resp, err := query.transport.QueryBlock(height)
	if err != nil {
		return nil, errors.QueryFailf("GetBlock err").AddCause(err)
	}

	block := new(model.Block)
	block.Header = resp.Block.Header
	block.Evidence = resp.Block.Evidence
	block.LastCommit = resp.Block.LastCommit
	block.Data = new(model.Data)
	block.Data.Txs = []model.Transaction{}
	for _, txBytes := range resp.Block.Data.Txs {
		var tx model.Transaction
		if err := query.transport.Cdc.UnmarshalJSON(txBytes, &tx); err != nil {
			return nil, err
		}
		block.Data.Txs = append(block.Data.Txs, tx)
	}
	return block, nil
}
