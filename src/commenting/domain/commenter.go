package domain

import "time"

type CommenterId int
type Commenter struct {
	commenterId CommenterId
	name        string
}

func (c *Commenter) Name() string {
	return c.name
}

func (c *Commenter) CommenterId() CommenterId {
	return c.commenterId
}

func (c *Commenter) NewComment(commentId CommentId, text string, page *Page, commentedAt time.Time) *Comment {
	return NewComment(
		commentId,
		page.pageId,
		text,
		c.commenterId,
		commentedAt,
	)
}

func NewCommenter(commenterId CommenterId, name string) *Commenter {
	return &Commenter{
		commenterId: commenterId,
		name:        name,
	}
}
