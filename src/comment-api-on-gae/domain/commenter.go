package domain

import "time"

type CommenterId int
type Commenter struct {
	Entity
	commenterId CommenterId
	name        string
}

func (c *Commenter) NewComment(text string, page *Page, commentedAt time.Time) *Comment {
	return &Comment{
		commenterId: c.commenterId,
		pageId:      page.pageId,
		text:        text,
		commentedAt: commentedAt,
	}
}

func NewCommenter(commenterId CommenterId, name string) *Commenter{
	return &Commenter{
		commenterId: commenterId,
		name: name,
	}
}
