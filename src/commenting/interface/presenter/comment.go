package presenter

import (
	"commenting/domain"
)

type comment struct {
	CommentId   int64      `json:"commentId"`
	PageId      string     `json:"pageId"`
	Text        string     `json:"text"`
	CommentedAt jsonTime   `json:"commentedAt"`
	Name        string     `json:"name"`
	Commenter   *commenter `json:"commenter"`
}

type commenter struct {
	UserID string `json:"userId"`
}

type CommentPresenter struct{}

func (p *CommentPresenter) Render(c *domain.Comment) *comment {
	return &comment{
		CommentId:   int64(c.CommentID()),
		PageId:      string(c.PageID()),
		Text:        c.Text(),
		CommentedAt: jsonTime{c.CommentedAt()},
		Name:        c.Name(),
		Commenter: &commenter{
			UserID: c.Commenter().UserID(),
		},
	}
}
