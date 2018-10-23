package domain

import "time"

type Commenter struct {
	userID string
}

func NewCommenter(userID string) *Commenter {
	return &Commenter{userID: userID}
}

func (c *Commenter) UserID() string {
	return c.userID
}

func (c *Commenter) CreateComment(commentID CommentID, pageID PageID, text, name string, commentedAt time.Time) *Comment {
	return NewComment(commentID, pageID, text, name, c, commentedAt)
}
