package domain

import (
	"time"
)

type CommentID int64

type Comment struct {
	commentID   CommentID
	pageId      PageID
	text        string
	name        string
	commenter   *Commenter
	commentedAt time.Time
}

func (c *Comment) Name() string {
	return c.name
}

func NewComment(
	commentID CommentID,
	pageID PageID,
	text string,
	name string,
	commenter *Commenter,
	commentedAt time.Time,
) *Comment {
	return &Comment{
		commentID:   commentID,
		pageId:      pageID,
		text:        text,
		name:        name,
		commenter:   commenter,
		commentedAt: commentedAt,
	}
}

func (c *Comment) Commenter() *Commenter {
	return c.commenter
}

func (c *Comment) CommentID() CommentID {
	return c.commentID
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
