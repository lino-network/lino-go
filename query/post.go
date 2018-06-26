package query

import (
	"github.com/lino-network/lino-go/model"
)

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

func (query *Query) GetAllPostComments(author, postID string) ([]*model.Comment, error) {
	postKey := getPostKey(author, postID)
	resKVs, err := query.transport.QuerySubspace(getPostCommentPrefix(postKey), PostKVStoreKey)
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

func (query *Query) GetAllPostViews(author, postID string) ([]*model.View, error) {
	postKey := getPostKey(author, postID)
	resKVs, err := query.transport.QuerySubspace(getPostViewPrefix(postKey), PostKVStoreKey)
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

func (query *Query) GetAllPostDonations(author, postID string) ([]*model.Donation, error) {
	postKey := getPostKey(author, postID)
	resKVs, err := query.transport.QuerySubspace(getPostDonationPrefix(postKey), PostKVStoreKey)
	if err != nil {
		return nil, err
	}

	var donations []*model.Donation
	for _, KV := range resKVs {
		donation := new(model.Donation)
		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, donation); err != nil {
			return nil, err
		}
		donations = append(donations, donation)
	}

	return donations, nil
}

// TODO: how to know the postID?
func (query *Query) GetAllUserDonations(username string) ([]*model.Donation, error) {
	resKVs, err := query.transport.QuerySubspace(getUserDonationPrefix(username), PostKVStoreKey)
	if err != nil {
		return nil, err
	}

	var donations []*model.Donation
	for _, KV := range resKVs {
		donation := new(model.Donation)
		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, donation); err != nil {
			return nil, err
		}
		donations = append(donations, donation)
	}

	return donations, nil
}

func (query *Query) GetAllPostReportOrUpvotes(author, postID string) ([]*model.ReportOrUpvote, error) {
	postKey := getPostKey(author, postID)
	resKVs, err := query.transport.QuerySubspace(getPostReportOrUpvotePrefix(postKey), PostKVStoreKey)
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
func (query *Query) GetAllUserReportOrUpvotes(username string) ([]*model.ReportOrUpvote, error) {
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

func (query *Query) GetAllPostLikes(author, postID string) ([]*model.Like, error) {
	postKey := getPostKey(author, postID)
	resKVs, err := query.transport.QuerySubspace(getPostLikePrefix(postKey), PostKVStoreKey)
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
