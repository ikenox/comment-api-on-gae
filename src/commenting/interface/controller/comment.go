package controller

import (
	"commenting/interface/presenter"
	"commenting/interface/repository"
	"commenting/usecase"
	"github.com/labstack/echo"
)

type CommentController struct{}

func NewCommentController() *CommentController {
	return &CommentController{}
}

// adapter層なので一意な変換以外のロジックは基本入らなくなる？ガード節とかも基本はいらない？
func (ctl *CommentController) List(c echo.Context) error {
	pageId := c.QueryParam("pageId")

	ctx := c.StdContext()
	u := usecase.NewCommentUseCase(
		repository.NewCommentRepository(ctx),
		repository.NewCommenterRepository(ctx),
		repository.NewPageRepository(ctx),
	)
	data, res := u.GetComments(pageId)

	json := (&presenter.CommentPresenter{}).RenderArray(data)
	return renderJSON(c, json, res)
}

func (ctl *CommentController) PostComment(c echo.Context) error {
	pageId := c.FormValue("pageId")
	name := c.FormValue("name")
	text := c.FormValue("text")

	ctx := c.StdContext()
	u := usecase.NewCommentUseCase(
		repository.NewCommentRepository(ctx),
		repository.NewCommenterRepository(ctx),
		repository.NewPageRepository(ctx),
	)
	data, result := u.PostComment(pageId, name, text)

	json := (&presenter.CommentPresenter{}).Render(data)
	return renderJSON(c, json, result)
}
