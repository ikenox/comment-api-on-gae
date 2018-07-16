package presenter

import (
	"comment-api-on-gae/domain"
	"time"
)

type commentJson struct {
	CommentId   int64          `json:"commentId"`
	PageId      string         `json:"pageId"`
	Text        string         `json:"text"`
	CommentedAt time.Time      `json:"commentedAt"`
	Commenter   *commenterJson `json:"commenter"`
}

type commenterJson struct {
	CommenterId int64  `json:"commenterId"`
	Name        string `json:"name"`
}

// TODO: 適切な名前は？CommentPresenterと言いつつCommenterも扱っている
// TODO: ぶっちゃけentityがjson情報持ってたほうが楽か
type CommentPresenter struct{}

func (p *CommentPresenter) Render(comment *domain.Comment, commenter *domain.Commenter) interface{} {
	return &commentJson{
		CommentId:   int64(comment.CommentId()),
		PageId:      string(comment.PageId()),
		Text:        comment.Text(),
		CommentedAt: comment.CommentedAt(),
		Commenter: &commenterJson{
			CommenterId: int64(comment.CommenterId()),
			Name:        commenter.Name(),
		},
	}
}
