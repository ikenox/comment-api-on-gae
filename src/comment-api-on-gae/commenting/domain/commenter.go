package domain

import "time"

type UserID string

type CommenterID int64

type Commenter struct {
	userID      UserID
	commenterID CommenterID
	name        string
}

func (c *Commenter) Name() string {
	return c.name
}

func (c *Commenter) CommenterID() CommenterID {
	return c.commenterID
}

func (c *Commenter) UserID() UserID {
	return c.userID
}

// commenter経由でcommentが生成されることでドメイン同士の関係性や主述が明確になる
// 一方でcommenterのインスタンス化が必須になるというパフォーマンス上のトレードオフが発生
func (c *Commenter) MakeComment(commentId CommentID, text string, page *Page, commentedAt time.Time) *Comment {
	return NewComment(
		commentId,
		page.pageID,
		text,
		c.commenterID,
		commentedAt,
	)
}

func NewCommenter(commenterID CommenterID, name string, userID UserID) *Commenter {
	return &Commenter{
		commenterID: commenterID,
		userID:      userID,
		name:        name,
	}
}
