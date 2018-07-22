package presenter

import (
	"commenting/usecase"
	"time"
)

type commentWithCommenter struct {
	Comment   *comment   `json:"comment"`
	Commenter *commenter `json:"commenter"`
}

type comment struct {
	CommentId   int64     `json:"commentId"`
	PageId      string    `json:"pageId"`
	Text        string    `json:"text"`
	CommentedAt time.Time `json:"commentedAt"`
}

type commenter struct {
	CommenterId int64  `json:"commenterId"`
	Name        string `json:"name"`
}

type CommentPresenter struct{}

func (p *CommentPresenter) RenderArray(d []*usecase.CommentWithCommenter) []*commentWithCommenter {
	json := make([]*commentWithCommenter, len(d))
	for i, c := range d {
		json[i] = p.Render(c)
	}
	return json
}

func (p *CommentPresenter) Render(d *usecase.CommentWithCommenter) *commentWithCommenter {
	return &commentWithCommenter{
		Comment: &comment{
			CommentId:   int64(d.Comment.CommentId()),
			PageId:      string(d.Comment.PageId()),
			Text:        d.Comment.Text(),
			CommentedAt: d.Comment.CommentedAt(),
		},
		Commenter: &commenter{
			CommenterId: int64(d.Commenter.CommenterId()),
			Name:        string(d.Commenter.Name()),
		},
	}
}
