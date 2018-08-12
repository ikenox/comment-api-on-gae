package domain

import "time"

type CommentID int64
type Comment struct {
	commentId   CommentID
	pageId      PageID
	text        string
	commenterId CommenterID
	commentedAt time.Time
}

func NewComment(
	commentId CommentID,
	pageId PageID,
	text string,
	commenterId CommenterID,
	commentedAt time.Time,
) *Comment {
	return &Comment{
		commentId:   commentId,
		pageId:      pageId,
		text:        text,
		commenterId: commenterId,
		commentedAt: commentedAt,
	}
}

func (c *Comment) CommentId() CommentID {
	return c.commentId
}

func (c *Comment) CommenterId() CommenterID {
	return c.commenterId
}

func (c *Comment) CommentedAt() time.Time {
	return c.commentedAt
}

func (c *Comment) Text() string {
	return c.text
}

func (c *Comment) PageId() PageID {
	return c.pageId
}
