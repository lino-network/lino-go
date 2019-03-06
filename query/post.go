package query

import (
	"context"

	"github.com/lino-network/lino-go/errors"
	"github.com/lino-network/lino-go/model"
)

// GetPostInfo returns post info given a permlink(author#postID).
func (query *Query) GetPostInfo(ctx context.Context, author, postID string) (*model.PostInfo, error) {
	permlink := getPermlink(author, postID)
	resp, err := query.transport.Query(ctx, PostKVStoreKey, PostInfoSubStore, []string{permlink})
	if err != nil {
		linoe, ok := err.(errors.Error)
		if ok && linoe.BlockChainCode() == uint32(errors.CodePostNotFound) {
			return nil, errors.EmptyResponse("post is not found")
		}
		return nil, err
	}
	postInfo := new(model.PostInfo)
	if err := query.transport.Cdc.UnmarshalJSON(resp, postInfo); err != nil {
		return nil, err
	}
	return postInfo, nil
}

// GetPostMeta returns post meta given a permlink.
func (query *Query) GetPostMeta(ctx context.Context, author, postID string) (*model.PostMeta, error) {
	permlink := getPermlink(author, postID)
	resp, err := query.transport.Query(ctx, PostKVStoreKey, PostMetaSubStore, []string{permlink})
	if err != nil {
		linoe, ok := err.(errors.Error)
		if ok && linoe.BlockChainCode() == uint32(errors.CodePostMetaNotFound) {
			return nil, errors.EmptyResponse("post meta is not found")
		}
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
func (query *Query) GetPostComment(ctx context.Context, author, postID, commentPermlink string) (*model.Comment, error) {
	permlink := getPermlink(author, postID)
	resp, err := query.transport.Query(ctx, PostKVStoreKey, PostCommentSubStore, []string{permlink})
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
func (query *Query) GetPostView(ctx context.Context, author, postID, viewUser string) (*model.View, error) {
	permlink := getPermlink(author, postID)
	resp, err := query.transport.Query(ctx, PostKVStoreKey, PostViewSubStore, []string{permlink})
	if err != nil {
		return nil, err
	}
	view := new(model.View)
	if err := query.transport.Cdc.UnmarshalJSON(resp, view); err != nil {
		return nil, err
	}
	return view, nil
}

// GetPostReportOrUpvote returns report or upvote that a user has given to a post.
func (query *Query) GetPostReportOrUpvote(ctx context.Context, author, postID, user string) (*model.ReportOrUpvote, error) {
	permlink := getPermlink(author, postID)
	resp, err := query.transport.Query(ctx, PostKVStoreKey, PostReportOrUpvoteSubStore, []string{permlink, user})
	if err != nil {
		linoe, ok := err.(errors.Error)
		if ok && linoe.BlockChainCode() == uint32(errors.CodePostReportOrUpvoteNotFound) {
			return nil, errors.EmptyResponse("post is not found")
		}
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

// // GetUserAllPosts returns all posts that a user has created.
// func (query *Query) GetUserAllPosts(ctx context.Context, username string) (map[string]*model.Post, error) {
// 	resKVs, err := query.transport.QuerySubspace(ctx, append(getUserPostInfoPrefix(username), PermLinkSeparator...), PostKVStoreKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	permlinkToPostMap := make(map[string]*model.Post)
// 	for _, KV := range resKVs {
// 		postInfo := new(model.PostInfo)
// 		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, postInfo); err != nil {
// 			return nil, err
// 		}

// 		pm, err := query.GetPostMeta(ctx, postInfo.Author, postInfo.PostID)
// 		if err != nil {
// 			return nil, err
// 		}

// 		post := &model.Post{
// 			PostID:                  postInfo.PostID,
// 			Title:                   postInfo.Title,
// 			Content:                 postInfo.Content,
// 			Author:                  postInfo.Author,
// 			ParentAuthor:            postInfo.ParentAuthor,
// 			ParentPostID:            postInfo.ParentPostID,
// 			SourceAuthor:            postInfo.SourceAuthor,
// 			SourcePostID:            postInfo.SourcePostID,
// 			Links:                   postInfo.Links,
// 			CreatedAt:               pm.CreatedAt,
// 			LastUpdatedAt:           pm.LastUpdatedAt,
// 			LastActivityAt:          pm.LastActivityAt,
// 			AllowReplies:            pm.AllowReplies,
// 			IsDeleted:               pm.IsDeleted,
// 			TotalDonateCount:        pm.TotalDonateCount,
// 			TotalReportCoinDay:      pm.TotalReportCoinDay,
// 			TotalUpvoteCoinDay:      pm.TotalUpvoteCoinDay,
// 			TotalViewCount:          pm.TotalViewCount,
// 			TotalReward:             pm.TotalReward,
// 			RedistributionSplitRate: pm.RedistributionSplitRate,
// 		}
// 		permlinkToPostMap[getSubstringAfterSubstore(KV.Key)] = post
// 	}

// 	return permlinkToPostMap, nil
// }

// // GetPostAllComments returns all comments that a post has.
// func (query *Query) GetPostAllComments(ctx context.Context, author, postID string) (map[string]*model.Comment, error) {
// 	permlink := getPermlink(author, postID)
// 	resKVs, err := query.transport.QuerySubspace(ctx, getPostCommentPrefix(permlink), PostKVStoreKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var permlinkToCommentsMap = make(map[string]*model.Comment)
// 	for _, KV := range resKVs {
// 		comment := new(model.Comment)
// 		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, comment); err != nil {
// 			return nil, err
// 		}

// 		permlinkToCommentsMap[getSubstringAfterKeySeparator(KV.Key)] = comment
// 	}

// 	return permlinkToCommentsMap, nil
// }

// // GetPostAllViews returns all views that a post has.
// func (query *Query) GetPostAllViews(ctx context.Context, author, postID string) (map[string]*model.View, error) {
// 	permlink := getPermlink(author, postID)
// 	resKVs, err := query.transport.QuerySubspace(ctx, getPostViewPrefix(permlink), PostKVStoreKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	userToViewMap := make(map[string]*model.View)
// 	for _, KV := range resKVs {
// 		view := new(model.View)
// 		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, view); err != nil {
// 			return nil, err
// 		}
// 		userToViewMap[getSubstringAfterKeySeparator(KV.Key)] = view
// 	}

// 	return userToViewMap, nil
// }

// // GetPostAllReportOrUpvotes returns all reports or upvotes that a post has received.
// func (query *Query) GetPostAllReportOrUpvotes(ctx context.Context, author, postID string) (map[string]*model.ReportOrUpvote, error) {
// 	permlink := getPermlink(author, postID)
// 	resKVs, err := query.transport.QuerySubspace(ctx, getPostReportOrUpvotePrefix(permlink), PostKVStoreKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	userToReportOrUpvotesMap := make(map[string]*model.ReportOrUpvote)
// 	for _, KV := range resKVs {
// 		reportOrUpvote := new(model.ReportOrUpvote)
// 		if err := query.transport.Cdc.UnmarshalJSON(KV.Value, reportOrUpvote); err != nil {
// 			return nil, err
// 		}
// 		userToReportOrUpvotesMap[getSubstringAfterKeySeparator(KV.Key)] = reportOrUpvote
// 	}

// 	return userToReportOrUpvotesMap, nil
// }
