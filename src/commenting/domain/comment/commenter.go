package comment

import "time"

type CommenterID int
type Commenter struct {
	commenterID CommenterID
	name        string
}

func (c *Commenter) Name() string {
	return c.name
}

func (c *Commenter) CommenterId() CommenterID {
	return c.commenterID
}

// commenter経由でcommentが生成されることでドメイン同士の関係性や主述が明確になる
// 一方でcommenterのインスタンス化が必須になるというパフォーマンス上のトレードオフが発生
func (c *Commenter) NewComment(commentId CommentID, text string, page *Page, commentedAt time.Time) *Comment {
	return NewComment(
		commentId,
		page.pageID,
		text,
		c.commenterID,
		commentedAt,
	)
}

func (c *Commenter) NewReply(replyID ReplyID, commentID CommentID, text string, repliedAt time.Time) *Reply {
	return newReply(
		replyID,
		commentID,
		c.commenterID,
		text,
		repliedAt,
	)
}

func NewCommenter(commenterId CommenterID, name string) *Commenter {
	return &Commenter{
		commenterID: commenterId,
		name:        name,
	}
}
