package presenter

import (
	"comment-api-on-gae/commenting/usecase"
)

type commentWithCommenter struct {
	Comment   *comment   `json:"comment"`
	Commenter *commenter `json:"commenter"`
}

type comment struct {
	CommentId   int64     `json:"commentId"`
	PageId      string    `json:"pageId"`
	Text        string    `json:"text"`
	CommentedAt jsonTime `json:"commentedAt"`
}

type commenter struct {
	UserID      string `json:"userId"`
	CommenterID int64  `json:"commenterId"`
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
	obj := &commentWithCommenter{}
	if d == nil {
		return nil
	}
	if d.Comment != nil {
		obj.Comment = &comment{
			CommentId:   int64(d.Comment.CommentID()),
			PageId:      string(d.Comment.PageID()),
			Text:        d.Comment.Text(),
			CommentedAt: jsonTime{d.Comment.CommentedAt()},
		}
	}
	if d.Commenter != nil {
		obj.Commenter = &commenter{
			UserID:      string(d.Commenter.UserID()),
			CommenterID: int64(d.Commenter.CommenterID()),
			Name:        string(d.Commenter.Name()),
		}
	}
	return obj
}
