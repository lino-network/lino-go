package query

import (
	"github.com/lino-network/lino-go/model"
)

// GetPostInfo returns post info given a permlink(author#postID).
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

// GetPostMeta returns post meta given a permlink.
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

// GetPostComment returns a specific comment of a post given the post permlink
// and comment permlink.
func (query *Query) GetPostComment(author, postID, commentPermlink string) (*model.Comment, error) {
	permlink := getPermlink(author, postID)
	resp, err := query.transport.Query(getPostCommentKey(permlink, commentPermlink), PostKVStoreKey)

	if err != nil {
		return nil, err
	}
	comment := new(model.Comment)
	if err := query.transport.Cdc.UnmarshalJSON(resp, comment); err != nil {
		return nil, err
	}
	return comment, nil
}

// GetPostView returns a view of a post performed by a user.
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

// GetPostDonations returns all donations that a user has given to a post.
func (query *Query) GetPostDonations(author, postID, donateUser string) (*model.Donations, error) {
	permlink := getPermlink(author, postID)
	resp, err := query.transport.Query(getPostDonationsKey(permlink, donateUser), PostKVStoreKey)
	if err != nil {
		return nil, err
	}
	donations := new(model.Donations)
	if err := query.transport.Cdc.UnmarshalJSON(resp, donations); err != nil {
		return nil, err
	}
	return donations, nil
}

// GetPostReportOrUpvote returns report or upvote that a user has given to a post.
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

//
// Range query
//

// GetUserAllPosts returns all posts that a user has created.
func (query *Query) GetUserAllPosts(username string) (map[string]*model.Post, error) {
	resKVs, err := query.transport.QuerySubspace(getUserPostInfoPrefix(username), PostKVStoreKey)
	if err != nil {
		return nil, err
	}

	permlinkToPostMap := make(map[string]*model.Post)
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
			TotalDonateCount:        pm.TotalDonateCount,
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

// GetPostAllComments returns all comments that a post has.
func (query *Query) GetPostAllComments(author, postID string) (map[string]*model.Comment, error) {
	permlink := getPermlink(author, postID)
	resKVs, err := query.transport.QuerySubspace(getPostCommentPrefix(permlink), PostKVStoreKey)
	if err != nil {
		return nil, err
	}

	var permlinkToCommentsMap = make(map[string]*model.Comment)
	for _, KV := range resKVs {
		comment := new(model.Comment)
		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, comment); err != nil {
			return nil, err
		}

		permlinkToCommentsMap[getSubstringAfterKeySeparator(KV.Key)] = comment
	}

	return permlinkToCommentsMap, nil
}

// GetPostAllViews returns all views that a post has.
func (query *Query) GetPostAllViews(author, postID string) (map[string]*model.View, error) {
	permlink := getPermlink(author, postID)
	resKVs, err := query.transport.QuerySubspace(getPostViewPrefix(permlink), PostKVStoreKey)
	if err != nil {
		return nil, err
	}

	userToViewMap := make(map[string]*model.View)
	for _, KV := range resKVs {
		view := new(model.View)
		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, view); err != nil {
			return nil, err
		}
		userToViewMap[getSubstringAfterKeySeparator(KV.Key)] = view
	}

	return userToViewMap, nil
}

// GetPostAllDonations returns all donations that a post has received.
func (query *Query) GetPostAllDonations(author, postID string) (map[string]*model.Donations, error) {
	permlink := getPermlink(author, postID)
	resKVs, err := query.transport.QuerySubspace(getPostDonationsPrefix(permlink), PostKVStoreKey)
	if err != nil {
		return nil, err
	}

	userToDonationsMap := make(map[string]*model.Donations)
	for _, KV := range resKVs {
		donations := new(model.Donations)
		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, donations); err != nil {
			return nil, err
		}
		userToDonationsMap[getSubstringAfterKeySeparator(KV.Key)] = donations
	}

	return userToDonationsMap, nil
}

// GetPostAllReportOrUpvotes returns all reports or upvotes that a post has received.
func (query *Query) GetPostAllReportOrUpvotes(author, postID string) (map[string]*model.ReportOrUpvote, error) {
	permlink := getPermlink(author, postID)
	resKVs, err := query.transport.QuerySubspace(getPostReportOrUpvotePrefix(permlink), PostKVStoreKey)
	if err != nil {
		return nil, err
	}

	userToReportOrUpvotesMap := make(map[string]*model.ReportOrUpvote)
	for _, KV := range resKVs {
		reportOrUpvote := new(model.ReportOrUpvote)
		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, reportOrUpvote); err != nil {
			return nil, err
		}
		userToReportOrUpvotesMap[getSubstringAfterKeySeparator(KV.Key)] = reportOrUpvote
	}

	return userToReportOrUpvotesMap, nil
}
