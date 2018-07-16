package controller

import (
	"comment-api-on-gae/interface/presenter"
	"comment-api-on-gae/interface/repository"
	"comment-api-on-gae/usecase"
	"github.com/labstack/echo"
)

type CommentController struct{}

func NewCommentController() *CommentController {
	return &CommentController{}
}

// adapter層なのでガード節などの変換関係ないビジネスロジックは基本入らなくなる？
func (ctl *CommentController) List(c echo.Context) error {
	var params struct {
		PageId string
	}
	// 変換失敗時点でreturn
	// アプリケーションエラーとは別？
	// アプリケーションエラーはプロトコル守っている前提で発生したエラー
	// この場合は指定のプロトコル守ってないので変換エラー
	if err := c.Bind(&params); err != nil {
		return err
	}

	ctx := c.StdContext()
	u := usecase.NewCommentUseCase(
		repository.NewCommentRepository(ctx),
		repository.NewCommenterRepository(ctx),
		repository.NewPageRepository(ctx),
	)
	comments, commenters, res := u.GetComments(params.PageId)

	// TODO: 別集約を1つにまとめて返すための正しい方法
	// TODO: Json用structの置き場所やネーミング
	var json []interface{}
	if len(comments) > 0 {
		json = make([]interface{}, len(comments))
		p := &presenter.CommentPresenter{}
		for i := 0; i < len(comments); i++ {
			if commenters[i] != nil && comments[i] != nil {
				json[i] = p.Render(comments[i], commenters[i])
			}
		}
	}

	// TODO: レスポンスのデータ構造
	return renderJSON(c, json, res)
}

func (ctl *CommentController) PostComment(c echo.Context) error {
	var params struct {
		PageId string
		Name   string
		Text   string
	}
	if err := c.Bind(&params); err != nil {
		return err
	}

	ctx := c.StdContext()
	u := usecase.NewCommentUseCase(
		repository.NewCommentRepository(ctx),
		repository.NewCommenterRepository(ctx),
		repository.NewPageRepository(ctx),
	)
	comment, commenter, result := u.PostComment(params.PageId, params.Name, params.Text)

	var json interface{}
	if comment != nil && commenter != nil {
		p := &presenter.CommentPresenter{}
		json = p.Render(comment, commenter)
	}
	return renderJSON(c, json, result)
}
