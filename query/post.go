package query

import (
	"github.com/lino-network/lino-go/model"
)

//
// Post related query
//
func (query *Query) GetPostInfo(author, postID string) (*model.PostInfo, error) {
	permlink := getPermlink(author, postID)
	resp, err := query.transport.Query(getPostInfoKey(permlink), PostKVStoreKey)
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
	permlink := getPermlink(author, postID)
	resp, err := query.transport.Query(getPostMetaKey(permlink), PostKVStoreKey)
	if err != nil {
		return nil, err
	}
	postMeta := new(model.PostMeta)
	if err := query.transport.Cdc.UnmarshalJSON(resp, postMeta); err != nil {
		return nil, err
	}
	return postMeta, nil
}

func (query *Query) GetPostComment(author, postID, commentpermlink string) (*model.Comment, error) {
	permlink := getPermlink(author, postID)
	resp, err := query.transport.Query(getPostCommentKey(permlink, commentpermlink), PostKVStoreKey)
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
	permlink := getPermlink(author, postID)
	resp, err := query.transport.Query(getPostViewKey(permlink, viewUser), PostKVStoreKey)
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
	permlink := getPermlink(author, postID)
	resp, err := query.transport.Query(getPostDonationKey(permlink, donateUser), PostKVStoreKey)
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
	permlink := getPermlink(author, postID)
	resp, err := query.transport.Query(getPostReportOrUpvoteKey(permlink, user), PostKVStoreKey)
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
	permlink := getPermlink(author, postID)
	resp, err := query.transport.Query(getPostLikeKey(permlink, likeUser), PostKVStoreKey)
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
// Range query
//

func (query *Query) GetUserAllPosts(username string) (map[string]*model.Post, error) {
	resKVs, err := query.transport.QuerySubspace(getUserPostInfoPrefix(username), PostKVStoreKey)
	if err != nil {
		return nil, err
	}

	var permlinkToPostMap = make(map[string]*model.Post)
	for _, KV := range resKVs {
		postInfo := new(model.PostInfo)
		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, postInfo); err != nil {
			return nil, err
		}

		pm, err := query.GetPostMeta(postInfo.Author, postInfo.PostID)
		if err != nil {
			return nil, err
		}

		post := &model.Post{
			PostID:                  postInfo.PostID,
			Title:                   postInfo.Title,
			Content:                 postInfo.Content,
			Author:                  postInfo.Author,
			ParentAuthor:            postInfo.ParentAuthor,
			ParentPostID:            postInfo.ParentPostID,
			SourceAuthor:            postInfo.SourceAuthor,
			SourcePostID:            postInfo.SourcePostID,
			Links:                   postInfo.Links,
			CreatedAt:               pm.CreatedAt,
			LastUpdatedAt:           pm.LastUpdatedAt,
			LastActivityAt:          pm.LastActivityAt,
			AllowReplies:            pm.AllowReplies,
			IsDeleted:               pm.IsDeleted,
			TotalLikeCount:          pm.TotalLikeCount,
			TotalDonateCount:        pm.TotalDonateCount,
			TotalLikeWeight:         pm.TotalLikeWeight,
			TotalDislikeWeight:      pm.TotalDislikeWeight,
			TotalReportStake:        pm.TotalReportStake,
			TotalUpvoteStake:        pm.TotalUpvoteStake,
			TotalViewCount:          pm.TotalViewCount,
			TotalReward:             pm.TotalReward,
			RedistributionSplitRate: pm.RedistributionSplitRate,
		}

		permlinkToPostMap[getSubstringAfterSubstore(KV.Key)] = post
	}

	return permlinkToPostMap, nil
}

func (query *Query) GetAllPostComments(author, postID string) (map[string]*model.Comment, error) {
	permlink := getPermlink(author, postID)
	resKVs, err := query.transport.QuerySubspace(getPostCommentPrefix(permlink), PostKVStoreKey)
	if err != nil {
		return nil, err
	}

	var permlinkToCommentMap = make(map[string]*model.Comment)
	for _, KV := range resKVs {
		comment := new(model.Comment)
		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, comment); err != nil {
			return nil, err
		}
		permlinkToCommentMap[getSubstringAfterKeySeparator(KV.Key)] = comment
	}

	return permlinkToCommentMap, nil
}

func (query *Query) GetAllPostViews(author, postID string) (map[string]*model.View, error) {
	permlink := getPermlink(author, postID)
	resKVs, err := query.transport.QuerySubspace(getPostViewPrefix(permlink), PostKVStoreKey)
	if err != nil {
		return nil, err
	}

	var userToViewMap = make(map[string]*model.View)
	for _, KV := range resKVs {
		view := new(model.View)
		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, view); err != nil {
			return nil, err
		}
		userToViewMap[getSubstringAfterKeySeparator(KV.Key)] = view
	}

	return userToViewMap, nil
}

func (query *Query) GetAllPostDonations(author, postID string) (map[string]*model.Donations, error) {
	permlink := getPermlink(author, postID)
	resKVs, err := query.transport.QuerySubspace(getPostDonationsPrefix(permlink), PostKVStoreKey)
	if err != nil {
		return nil, err
	}

	var userToDonationMap = make(map[string]*model.Donations)
	for _, KV := range resKVs {
		donations := new(model.Donations)
		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, donations); err != nil {
			return nil, err
		}
		userToDonationMap[getSubstringAfterKeySeparator(KV.Key)] = donations
	}

	return userToDonationMap, nil
}

// // TODO: how to know the postID?
// func (query *Query) GetAllUserDonations(username string) ([]*model.Donation, error) {
// 	resKVs, err := query.transport.QuerySubspace(getUserDonationsPrefix(username), PostKVStoreKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var donations []*model.Donation
// 	for _, KV := range resKVs {
// 		donation := new(model.Donation)
// 		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, donation); err != nil {
// 			return nil, err
// 		}
// 		donations = append(donations, donation)
// 	}

// 	return donations, nil
// }

func (query *Query) GetAllPostReportOrUpvotes(author, postID string) (map[string]*model.ReportOrUpvote, error) {
	permlink := getPermlink(author, postID)
	resKVs, err := query.transport.QuerySubspace(getPostReportOrUpvotePrefix(permlink), PostKVStoreKey)
	if err != nil {
		return nil, err
	}

	var userToReportOrUpvotesMap = make(map[string]*model.ReportOrUpvote)
	for _, KV := range resKVs {
		reportOrUpvote := new(model.ReportOrUpvote)
		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, reportOrUpvote); err != nil {
			return nil, err
		}
		userToReportOrUpvotesMap[getSubstringAfterKeySeparator(KV.Key)] = reportOrUpvote
	}

	return userToReportOrUpvotesMap, nil
}

// // TODO: how to know the postID?
// func (query *Query) GetAllUserReportOrUpvotes(username string) ([]*model.ReportOrUpvote, error) {
// 	resKVs, err := query.transport.QuerySubspace(getUserReportOrUpvotePrefix(username), PostKVStoreKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var reportOrUpvotes []*model.ReportOrUpvote
// 	for _, KV := range resKVs {
// 		reportOrUpvote := new(model.ReportOrUpvote)
// 		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, reportOrUpvote); err != nil {
// 			return nil, err
// 		}
// 		reportOrUpvotes = append(reportOrUpvotes, reportOrUpvote)
// 	}

// 	return reportOrUpvotes, nil
// }

func (query *Query) GetAllPostLikes(author, postID string) (map[string]*model.Like, error) {
	permlink := getPermlink(author, postID)
	resKVs, err := query.transport.QuerySubspace(getPostLikePrefix(permlink), PostKVStoreKey)
	if err != nil {
		return nil, err
	}

	var userToLikeMap = make(map[string]*model.Like)
	for _, KV := range resKVs {
		like := new(model.Like)
		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, like); err != nil {
			return nil, err
		}
		userToLikeMap[getSubstringAfterKeySeparator(KV.Key)] = like
	}

	return userToLikeMap, nil
}
