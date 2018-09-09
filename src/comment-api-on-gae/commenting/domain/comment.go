package domain

import "time"

type CommentID int64

type Comment struct {
	commentID   CommentID
	pageId      PageID
	text        string
	commenterID CommenterID
	commentedAt time.Time
}

func NewComment(
	commentID CommentID,
	pageID PageID,
	text string,
	commenterID CommenterID,
	commentedAt time.Time,
) *Comment {
	// Commentが絶対守らなくてはならない不変条件はここに
	return &Comment{
		commentID:   commentID,
		pageId:      pageID,
		text:        text,
		commenterID: commenterID,
		commentedAt: commentedAt,
	}
}

func (c *Comment) CommentID() CommentID {
	return c.commentID
}

func (c *Comment) CommenterID() CommenterID {
	return c.commenterID
}

func (c *Comment) CommentedAt() time.Time {
	return c.commentedAt
}

func (c *Comment) Text() string {
	return c.text
}

func (c *Comment) PageID() PageID {
	return c.pageId
}
