package model

type Transaction interface{}
type PublishPost struct {
	PostID                  string            `json:"post_id"`
	Title                   string            `json:"title"`
	Content                 string            `json:"content"`
	Author                  string            `json:"author"`
	ParentAuthor            string            `json:"parent_author"`
	ParentPostID            string            `json:"parent_postID"`
	SourceAuthor            string            `json:"source_author"`
	SourcePostID            string            `json:"source_postID"`
	Links                   map[string]string `json:"links"`
	RedistributionSplitRate string            `json:"redistribution_split_rate"`
}

type TransferToAddress struct {
	From      string `json:"from"`
	ToAddress string `json:"to_address"`
	Amount    string `json:"content"`
	Memo      string `json:"memo"`
}

type TransferToUser struct {
	From   string `json:"from"`
	ToUser string `json:"to_user"`
	Amount string `json:"content"`
	Memo   string `json:"memo"`
}
