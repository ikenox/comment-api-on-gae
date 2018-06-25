package domain

import "time"

type CommentId int
type Comment struct {
	commentId   CommentId
	pageId      PageId
	text        string
	commenterId CommenterId
	commentedAt time.Time
}

func NewComment(
	commentId CommentId,
	pageId PageId,
	text string,
	commenterId CommenterId,
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

func (c *Comment) CommentId() CommentId {
	return c.commentId
}

func (c *Comment) CommenterId() CommenterId {
	return c.commenterId
}

func (c *Comment) CommentedAt() time.Time {
	return c.commentedAt
}

func (c *Comment) Text() string {
	return c.text
}

func (c *Comment) PageId() PageId {
	return c.pageId
}
