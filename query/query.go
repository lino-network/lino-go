package query

import (
	"fmt"

	"github.com/lino-network/lino-go/errors"
	"github.com/lino-network/lino-go/model"
	"github.com/lino-network/lino-go/transport"
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
func (query *Query) GetAccountInfo(username string) (*model.AccountInfo, errors.Error) {
	resp, err := query.transport.Query(getAccountInfoKey(username), AccountKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetAccountInfo err: %v", err)
	}
	info := new(model.AccountInfo)
	if err := query.transport.Cdc.UnmarshalJSON(resp, info); err != nil {
		return nil, errors.UnmarshalFailf("GetAccountInfo err: %v", err)
	}
	return info, nil
}

func (query *Query) DoesUsernameMatchMasterPrivKey(username, masterPrivKeyHex string) (bool, errors.Error) {
	accInfo, err := query.GetAccountInfo(username)
	if err != nil {
		return false, err
	}

	masterPrivKey, e := transport.GetPrivKeyFromHex(masterPrivKeyHex)
	if e != nil {
		return false, errors.FailedToGetPrivKeyFromHexf("DoesUsernameMatchMasterPrivKey failed to get priv key from hex: %v", masterPrivKeyHex)
	}

	return accInfo.MasterKey.Equals(masterPrivKey.PubKey()), nil
}

func (query *Query) DoesUsernameMatchTxPrivKey(username, txPrivKeyHex string) (bool, errors.Error) {
	accInfo, err := query.GetAccountInfo(username)
	if err != nil {
		return false, err
	}

	txPrivKey, e := transport.GetPrivKeyFromHex(txPrivKeyHex)
	if e != nil {
		return false, errors.FailedToGetPrivKeyFromHexf("DoesUsernameMatchTxPrivKey failed to get priv key from hex: %v", txPrivKeyHex)
	}

	return accInfo.TransactionKey.Equals(txPrivKey.PubKey()), nil
}

func (query *Query) DoesUsernameMatchPostPrivKey(username, postPrivKeyHex string) (bool, errors.Error) {
	accInfo, err := query.GetAccountInfo(username)
	if err != nil {
		return false, err
	}

	postPrivKey, e := transport.GetPrivKeyFromHex(postPrivKeyHex)
	if e != nil {
		return false, errors.FailedToGetPrivKeyFromHexf("DoesUsernameMatchPostPrivKey failed to get priv key from hex: %v", postPrivKeyHex)
	}

	return accInfo.PostKey.Equals(postPrivKey.PubKey()), nil
}

func (query *Query) GetAccountBank(username string) (*model.AccountBank, errors.Error) {
	resp, err := query.transport.Query(getAccountBankKey(username), AccountKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetAccountBank err: %v", err)
	}
	bank := new(model.AccountBank)
	if err := query.transport.Cdc.UnmarshalJSON(resp, bank); err != nil {
		return nil, errors.UnmarshalFailf("GetAccountBank err: %v", err)
	}
	return bank, nil
}

func (query *Query) GetAccountMeta(username string) (*model.AccountMeta, errors.Error) {
	resp, err := query.transport.Query(getAccountMetaKey(username), AccountKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetAccountMeta err: %v", err)
	}
	meta := new(model.AccountMeta)
	if err := query.transport.Cdc.UnmarshalJSON(resp, meta); err != nil {
		return nil, errors.UnmarshalFailf("GetAccountMeta err: %v", err)
	}
	return meta, nil
}

func (query *Query) GetSeqNumber(username string) (int64, errors.Error) {
	meta, err := query.GetAccountMeta(username)
	if err != nil {
		return 0, err
	}
	return meta.Sequence, nil
}

func (query *Query) GetAllBalanceHistory(username string) (*model.BalanceHistory, errors.Error) {
	accountBank, err := query.GetAccountBank(username)
	if err != nil {
		return nil, err
	}

	allBalanceHistory := new(model.BalanceHistory)
	bucketSlot := int64(0)
	if accountBank.NumOfTx != 0 {
		bucketSlot = (accountBank.NumOfTx - 1) / 100
	}

	for i := int64(0); i <= bucketSlot; i++ {
		balanceHistory, err := query.GetBalanceHistory(username, i)
		if err != nil {
			return nil, err
		}

		allBalanceHistory.Details = append(allBalanceHistory.Details, balanceHistory.Details...)
	}

	return allBalanceHistory, nil
}

func (query *Query) GetBalanceHistory(username string, index int64) (*model.BalanceHistory, errors.Error) {
	resp, err := query.transport.Query(getBalanceHistoryKey(username, index), AccountKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetBalanceHistory err: %v", err)
	}
	balanceHistory := new(model.BalanceHistory)
	if err := query.transport.Cdc.UnmarshalJSON(resp, balanceHistory); err != nil {
		return nil, errors.UnmarshalFailf("GetBalanceHistory err: %v", err)
	}
	return balanceHistory, nil
}

func (query *Query) GetGrantList(username string) (*model.GrantKeyList, errors.Error) {
	resp, err := query.transport.Query(getGrantKeyListKey(username), AccountKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetGrantList err: %v", err)
	}

	grantKeyList := new(model.GrantKeyList)
	if err := query.transport.Cdc.UnmarshalJSON(resp, grantKeyList); err != nil {
		return grantKeyList, errors.UnmarshalFailf("GetGrantList err: %v", err)
	}
	return grantKeyList, nil
}

func (query *Query) GetReward(username string) (*model.Reward, errors.Error) {
	resp, err := query.transport.Query(getRewardKey(username), AccountKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetReward err: %v", err)
	}

	reward := new(model.Reward)
	if err := query.transport.Cdc.UnmarshalJSON(resp, reward); err != nil {
		return reward, errors.UnmarshalFailf("GetReward err: %v", err)
	}
	return reward, nil
}

func (query *Query) GetRelationship(me, other string) (*model.Relationship, errors.Error) {
	resp, err := query.transport.Query(getRelationshipKey(me, other), AccountKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetRelationship err: %v", err)
	}

	relationship := new(model.Relationship)
	if err := query.transport.Cdc.UnmarshalJSON(resp, relationship); err != nil {
		return relationship, errors.UnmarshalFailf("GetRelationship err: %v", err)
	}
	return relationship, nil
}

func (query *Query) GetFollowerMeta(me, myFollower string) (*model.FollowerMeta, errors.Error) {
	resp, err := query.transport.Query(getFollowerKey(me, myFollower), AccountKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetFollowerMeta err: %v", err)
	}

	followerMeta := new(model.FollowerMeta)
	if err := query.transport.Cdc.UnmarshalJSON(resp, followerMeta); err != nil {
		return followerMeta, errors.UnmarshalFailf("GetFollowerMeta err: %v", err)
	}
	return followerMeta, nil
}

func (query *Query) GetFollowingMeta(me, myFollowing string) (*model.FollowingMeta, errors.Error) {
	resp, err := query.transport.Query(getFollowerKey(me, myFollowing), AccountKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetFollowingMeta err: %v", err)
	}

	followingMeta := new(model.FollowingMeta)
	if err := query.transport.Cdc.UnmarshalJSON(resp, followingMeta); err != nil {
		return followingMeta, errors.UnmarshalFailf("GetFollowingMeta err: %v", err)
	}
	return followingMeta, nil
}

//
// Post related query
//

func (query *Query) GetPostInfo(author, postID string) (*model.PostInfo, errors.Error) {
	postKey := getPostKey(author, postID)
	resp, err := query.transport.Query(getPostInfoKey(postKey), PostKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetPostInfo err: %v", err)
	}
	postInfo := new(model.PostInfo)
	if err := query.transport.Cdc.UnmarshalJSON(resp, postInfo); err != nil {
		return nil, errors.UnmarshalFailf("GetPostInfo err: %v", err)
	}
	return postInfo, nil
}

func (query *Query) GetPostMeta(author, postID string) (*model.PostMeta, errors.Error) {
	postKey := getPostKey(author, postID)
	resp, err := query.transport.Query(getPostMetaKey(postKey), PostKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetPostMeta err: %v", err)
	}
	postMeta := new(model.PostMeta)
	if err := query.transport.Cdc.UnmarshalJSON(resp, postMeta); err != nil {
		return nil, errors.UnmarshalFailf("GetPostMeta err: %v", err)
	}
	return postMeta, nil
}

func (query *Query) GetPostComment(author, postID, commentPostKey string) (*model.Comment, errors.Error) {
	postKey := getPostKey(author, postID)
	resp, err := query.transport.Query(getPostCommentKey(postKey, commentPostKey), PostKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetPostComment err: %v", err)
	}
	comment := new(model.Comment)
	if err := query.transport.Cdc.UnmarshalJSON(resp, comment); err != nil {
		return nil, errors.UnmarshalFailf("GetPostComment err: %v", err)
	}
	return comment, nil
}

func (query *Query) GetPostView(author, postID, viewUser string) (*model.View, errors.Error) {
	postKey := getPostKey(author, postID)
	resp, err := query.transport.Query(getPostViewKey(postKey, viewUser), PostKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetPostView err: %v", err)
	}
	view := new(model.View)
	if err := query.transport.Cdc.UnmarshalJSON(resp, view); err != nil {
		return nil, errors.UnmarshalFailf("GetPostView err: %v", err)
	}
	return view, nil
}

func (query *Query) GetPostDonation(author, postID, donateUser string) (*model.Donation, errors.Error) {
	postKey := getPostKey(author, postID)
	resp, err := query.transport.Query(getPostDonationKey(postKey, donateUser), PostKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetPostDonation err: %v", err)
	}
	donation := new(model.Donation)
	if err := query.transport.Cdc.UnmarshalJSON(resp, donation); err != nil {
		return nil, errors.UnmarshalFailf("GetPostDonation err: %v", err)
	}
	return donation, nil
}

func (query *Query) GetPostReportOrUpvote(author, postID, user string) (*model.ReportOrUpvote, errors.Error) {
	postKey := getPostKey(author, postID)
	resp, err := query.transport.Query(getPostReportOrUpvoteKey(postKey, user), PostKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetPostReportOrUpvote err: %v", err)
	}
	reportOrUpvote := new(model.ReportOrUpvote)
	if err := query.transport.Cdc.UnmarshalJSON(resp, reportOrUpvote); err != nil {
		return nil, errors.UnmarshalFailf("GetPostReportOrUpvote err: %v", err)
	}
	return reportOrUpvote, nil
}

func (query *Query) GetPostLike(author, postID, likeUser string) (*model.Like, errors.Error) {
	postKey := getPostKey(author, postID)
	resp, err := query.transport.Query(getPostLikeKey(postKey, likeUser), PostKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetPostLike err: %v", err)
	}
	like := new(model.Like)
	if err := query.transport.Cdc.UnmarshalJSON(resp, like); err != nil {
		return nil, errors.UnmarshalFailf("GetPostLike err: %v", err)
	}
	return like, nil
}

//
// Validator related query
//
func (query *Query) GetValidator(username string) (*model.Validator, errors.Error) {
	resp, err := query.transport.Query(getValidatorKey(username), ValidatorKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetValidator err: %v", err)
	}
	validator := new(model.Validator)
	if err := query.transport.Cdc.UnmarshalJSON(resp, validator); err != nil {
		return nil, errors.UnmarshalFailf("GetValidator err: %v", err)
	}
	return validator, nil
}

func (query *Query) GetAllValidators() (*model.ValidatorList, errors.Error) {
	resp, err := query.transport.Query(getValidatorListKey(), ValidatorKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetAllValidators err: %v", err)
	}

	validatorList := new(model.ValidatorList)
	if err := query.transport.Cdc.UnmarshalJSON(resp, validatorList); err != nil {
		return validatorList, errors.UnmarshalFailf("GetAllValidators err: %v", err)
	}
	return validatorList, nil
}

//
// Vote related query
//

func (query *Query) GetDelegation(voter, delegator string) (*model.Delegation, errors.Error) {
	resp, err := query.transport.Query(getDelegationKey(voter, delegator), VoteKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetDelegation err: %v", err)
	}
	delegation := new(model.Delegation)
	if err := query.transport.Cdc.UnmarshalJSON(resp, delegation); err != nil {
		return nil, errors.UnmarshalFailf("GetDelegation err: %v", err)
	}
	return delegation, nil
}

func (query *Query) GetVoter(voterName string) (*model.Voter, errors.Error) {
	resp, err := query.transport.Query(getVoterKey(voterName), VoteKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetVoter err: %v", err)
	}
	voter := new(model.Voter)
	if err := query.transport.Cdc.UnmarshalJSON(resp, voter); err != nil {
		return nil, errors.UnmarshalFailf("GetVoter err: %v", err)
	}
	return voter, nil
}

func (query *Query) GetDelegateeList(delegatorName string) (*model.DelegateeList, errors.Error) {
	resp, err := query.transport.Query(GetDelegateeListKey(delegatorName), VoteKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetDelegateeList err: %v", err)
	}
	delegateeList := new(model.DelegateeList)
	if err := query.transport.Cdc.UnmarshalJSON(resp, delegateeList); err != nil {
		return nil, errors.UnmarshalFailf("GetDelegateeList err: %v", err)
	}
	return delegateeList, nil
}

func (query *Query) GetVote(proposalID, voter string) (*model.Vote, errors.Error) {
	resp, err := query.transport.Query(getVoteKey(proposalID, voter), VoteKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetVote err: %v", err)
	}
	vote := new(model.Vote)
	if err := query.transport.Cdc.UnmarshalJSON(resp, vote); err != nil {
		return nil, errors.UnmarshalFailf("GetVote err: %v", err)
	}
	return vote, nil
}

//
// Developer related query
//
func (query *Query) GetDeveloper(developerName string) (*model.Developer, errors.Error) {
	resp, err := query.transport.Query(getDeveloperKey(developerName), DeveloperKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetDeveloper err: %v", err)
	}
	developer := new(model.Developer)
	if err := query.transport.Cdc.UnmarshalJSON(resp, developer); err != nil {
		return nil, errors.UnmarshalFailf("GetDeveloper err: %v", err)
	}
	return developer, nil
}

func (query *Query) GetDevelopers() (*model.DeveloperList, errors.Error) {
	resp, err := query.transport.Query(getDeveloperListKey(), DeveloperKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetDevelopers err: %v", err)
	}

	developerList := new(model.DeveloperList)
	if err := query.transport.Cdc.UnmarshalJSON(resp, developerList); err != nil {
		return nil, errors.UnmarshalFailf("GetDevelopers err: %v", err)
	}
	return developerList, nil
}

//
// Infra related query
//
func (query *Query) GetInfraProvider(providerName string) (*model.InfraProvider, errors.Error) {
	resp, err := query.transport.Query(getInfraProviderKey(providerName), InfraKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetInfraProvider err: %v", err)
	}
	provider := new(model.InfraProvider)
	if err := query.transport.Cdc.UnmarshalJSON(resp, provider); err != nil {
		return nil, errors.UnmarshalFailf("GetInfraProvider err: %v", err)
	}
	return provider, nil
}

func (query *Query) GetInfraProviders() (*model.InfraProviderList, errors.Error) {
	resp, err := query.transport.Query(getInfraProviderListKey(), InfraKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetInfraProviders err: %v", err)
	}

	providerList := new(model.InfraProviderList)
	if err := query.transport.Cdc.UnmarshalJSON(resp, providerList); err != nil {
		return nil, errors.UnmarshalFailf("GetInfraProviders err: %v", err)
	}
	return providerList, nil
}

//
// proposal related query
//
func (query *Query) GetProposalList() (*model.ProposalList, errors.Error) {
	resp, err := query.transport.Query(getProposalListKey(), ProposalKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetProposalList err: %v", err)
	}

	proposalList := new(model.ProposalList)
	if err := query.transport.Cdc.UnmarshalJSON(resp, proposalList); err != nil {
		return nil, errors.UnmarshalFailf("GetProposalList err: %v", err)
	}
	return proposalList, nil
}

func (query *Query) GetProposal(proposalID string) (*model.Proposal, errors.Error) {
	resp, err := query.transport.Query(getProposalKey(proposalID), ProposalKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetProposal err: %v", err)
	}

	proposal := new(model.Proposal)
	if err := query.transport.Cdc.UnmarshalJSON(resp, proposal); err != nil {
		return nil, errors.UnmarshalFailf("GetProposal err: %v", err)
	}
	return proposal, nil
}

func (query *Query) GetOngoingProposal() ([]*model.Proposal, errors.Error) {
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

func (query *Query) GetExpiredProposal() ([]*model.Proposal, errors.Error) {
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
func (query *Query) GetEvaluateOfContentValueParam() (*model.EvaluateOfContentValueParam, errors.Error) {
	resp, err := query.transport.Query(getEvaluateOfContentValueParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetEvaluateOfContentValueParam err: %v", err)
	}

	param := new(model.EvaluateOfContentValueParam)
	fmt.Println("---paramBytes: ", resp)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, errors.UnmarshalFailf("GetEvaluateOfContentValueParam err: %v", err)
	}
	return param, nil
}

func (query *Query) GetGlobalAllocationParam() (*model.GlobalAllocationParam, errors.Error) {
	resp, err := query.transport.Query(getGlobalAllocationParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetGlobalAllocationParam err: %v", err)
	}

	param := new(model.GlobalAllocationParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, errors.UnmarshalFailf("GetGlobalAllocationParam err: %v", err)
	}
	return param, nil
}

func (query *Query) GetInfraInternalAllocationParam() (*model.InfraInternalAllocationParam, errors.Error) {
	resp, err := query.transport.Query(getInfraInternalAllocationParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetInfraInternalAllocationParam err: %v", err)
	}

	param := new(model.InfraInternalAllocationParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, errors.UnmarshalFailf("GetInfraInternalAllocationParam err: %v", err)
	}
	return param, nil
}

func (query *Query) GetDeveloperParam() (*model.DeveloperParam, errors.Error) {
	resp, err := query.transport.Query(getDeveloperParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetDeveloperParam err: %v", err)
	}

	param := new(model.DeveloperParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, errors.UnmarshalFailf("GetDeveloperParam err: %v", err)
	}
	return param, nil
}

func (query *Query) GetVoteParam() (*model.VoteParam, errors.Error) {
	resp, err := query.transport.Query(getVoteParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetVoteParam err: %v", err)
	}

	param := new(model.VoteParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, errors.UnmarshalFailf("GetVoteParam err: %v", err)
	}
	return param, nil
}

func (query *Query) GetProposalParam() (*model.ProposalParam, errors.Error) {
	resp, err := query.transport.Query(getProposalParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetProposalParam err: %v", err)
	}

	param := new(model.ProposalParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, errors.UnmarshalFailf("GetProposalParam err: %v", err)
	}
	return param, nil
}

func (query *Query) GetValidatorParam() (*model.ValidatorParam, errors.Error) {
	resp, err := query.transport.Query(getValidatorParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetValidatorParam err: %v", err)
	}

	param := new(model.ValidatorParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, errors.UnmarshalFailf("GetValidatorParam err: %v", err)
	}
	return param, nil
}

func (query *Query) GetCoinDayParam() (*model.CoinDayParam, errors.Error) {
	resp, err := query.transport.Query(getCoinDayParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetCoinDayParam err: %v", err)
	}

	param := new(model.CoinDayParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, errors.UnmarshalFailf("GetCoinDayParam err: %v", err)
	}
	return param, nil
}

func (query *Query) GetBandwidthParam() (*model.BandwidthParam, errors.Error) {
	resp, err := query.transport.Query(getBandwidthParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetBandwidthParam err: %v", err)
	}

	param := new(model.BandwidthParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, errors.UnmarshalFailf("GetBandwidthParam err: %v", err)
	}
	return param, nil
}

func (query *Query) GetAccountParam() (*model.AccountParam, errors.Error) {
	resp, err := query.transport.Query(getAccountParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, errors.QueryFailf("GetAccountParam err: %v", err)
	}

	param := new(model.AccountParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, errors.UnmarshalFailf("GetAccountParam err: %v", err)
	}
	return param, nil
}

//
// get block
//
func (query *Query) GetBlock(height int64) (*model.Block, errors.Error) {
	resp, err := query.transport.QueryBlock(height)
	if err != nil {
		return nil, errors.QueryFailf("GetBlock err: %v", err)
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
			return nil, errors.UnmarshalFailf("GetBlock err: %v", err)
		}
		block.Data.Txs = append(block.Data.Txs, tx)
	}
	return block, nil
}
