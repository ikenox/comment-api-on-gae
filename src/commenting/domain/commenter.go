package domain

import "time"

type CommenterID int
type Commenter struct {
	commenterId CommenterID
	name        string
}

func (c *Commenter) Name() string {
	return c.name
}

func (c *Commenter) CommenterId() CommenterID {
	return c.commenterId
}

func (c *Commenter) NewComment(commentId CommentID, text string, page *Page, commentedAt time.Time) *Comment {
	return NewComment(
		commentId,
		page.pageId,
		text,
		c.commenterId,
		commentedAt,
	)
}

func NewCommenter(commenterId CommenterID, name string) *Commenter {
	return &Commenter{
		commenterId: commenterId,
		name:        name,
	}
}
