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

func (query *Query) GetPostDonations(author, postID, donateUser string) (*model.Donations, error) {
	permlink := getPermlink(author, postID)
	resp, err := query.transport.Query(getPostDonationKey(permlink, donateUser), PostKVStoreKey)
	if err != nil {
		return nil, err
	}
	donations := new(model.Donations)
	if err := query.transport.Cdc.UnmarshalJSON(resp, donations); err != nil {
		return nil, err
	}
	return donations, nil
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

func (query *Query) GetUserAllPosts(username string) ([]*model.Post, error) {
	resKVs, err := query.transport.QuerySubspace(getUserPostInfoPrefix(username), PostKVStoreKey)
	if err != nil {
		return nil, err
	}

	var posts []*model.Post
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

		posts = append(posts, post)
	}

	return posts, nil
}

func (query *Query) GetPostAllComments(author, postID string) ([]*model.Comment, error) {
	permlink := getPermlink(author, postID)
	resKVs, err := query.transport.QuerySubspace(getPostCommentPrefix(permlink), PostKVStoreKey)
	if err != nil {
		return nil, err
	}

	var comments []*model.Comment
	for _, KV := range resKVs {
		comment := new(model.Comment)
		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, comment); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (query *Query) GetPostAllViews(author, postID string) ([]*model.View, error) {
	permlink := getPermlink(author, postID)
	resKVs, err := query.transport.QuerySubspace(getPostViewPrefix(permlink), PostKVStoreKey)
	if err != nil {
		return nil, err
	}

	var views []*model.View
	for _, KV := range resKVs {
		view := new(model.View)
		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, view); err != nil {
			return nil, err
		}
		views = append(views, view)
	}

	return views, nil
}

// TODO: how to know the postID?
func (query *Query) GetUserAllDonations(username string) ([]*model.Donations, error) {
	resKVs, err := query.transport.QuerySubspace(getUserDonationPrefix(username), PostKVStoreKey)
	if err != nil {
		return nil, err
	}

	var donationsList []*model.Donations
	for _, KV := range resKVs {
		donations := new(model.Donations)
		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, donations); err != nil {
			return nil, err
		}
		donationsList = append(donationsList, donations)
	}

	return donationsList, nil
}

func (query *Query) GetPostAllReportOrUpvotes(author, postID string) ([]*model.ReportOrUpvote, error) {
	permlink := getPermlink(author, postID)
	resKVs, err := query.transport.QuerySubspace(getPostReportOrUpvotePrefix(permlink), PostKVStoreKey)
	if err != nil {
		return nil, err
	}

	var reportOrUpvotes []*model.ReportOrUpvote
	for _, KV := range resKVs {
		reportOrUpvote := new(model.ReportOrUpvote)
		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, reportOrUpvote); err != nil {
			return nil, err
		}
		reportOrUpvotes = append(reportOrUpvotes, reportOrUpvote)
	}

	return reportOrUpvotes, nil
}

// TODO: how to know the postID?
func (query *Query) GetUserAllReportOrUpvotes(username string) ([]*model.ReportOrUpvote, error) {
	resKVs, err := query.transport.QuerySubspace(getUserReportOrUpvotePrefix(username), PostKVStoreKey)
	if err != nil {
		return nil, err
	}

	var reportOrUpvotes []*model.ReportOrUpvote
	for _, KV := range resKVs {
		reportOrUpvote := new(model.ReportOrUpvote)
		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, reportOrUpvote); err != nil {
			return nil, err
		}
		reportOrUpvotes = append(reportOrUpvotes, reportOrUpvote)
	}

	return reportOrUpvotes, nil
}

func (query *Query) GetPostAllLikes(author, postID string) ([]*model.Like, error) {
	permlink := getPermlink(author, postID)
	resKVs, err := query.transport.QuerySubspace(getPostLikePrefix(permlink), PostKVStoreKey)
	if err != nil {
		return nil, err
	}

	var likes []*model.Like
	for _, KV := range resKVs {
		like := new(model.Like)
		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, like); err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}

	return likes, nil
}
