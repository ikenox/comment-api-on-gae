package controller

import (
	"comment-api-on-gae/commenting/interface/presenter"
	"comment-api-on-gae/commenting/interface/repository"
	"comment-api-on-gae/commenting/usecase"
	"github.com/labstack/echo"
)

type CommentController struct{}

func NewCommentController() *CommentController {
	return &CommentController{}
}

// adapter層なので一意な変換以外のロジックは基本入らなくなる？
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
	return presenter.RenderJSON(c, json, res)
}

func (ctl *CommentController) PostComment(c echo.Context) error {
	var p = &struct {
		PageId  string `json:"pageId"`
		Name    string `json:"name"`
		Text    string `json:"text"`
		IDToken string `json:"idToken"`
	}{}
	if err := c.Bind(p); err != nil {
		// 変換エラーはinterface adapter層における異常系
		// usecaseでのエラーはinterface adapterにとっては正常系
		return err
	}

	ctx := c.StdContext()
	u := usecase.NewCommentUseCase(
		repository.NewCommentRepository(ctx),
		repository.NewCommenterRepository(ctx),
		repository.NewPageRepository(ctx),
	)
	data, result := u.PostComment(p.IDToken, p.Name, p.PageId, p.Text)
	pr := &presenter.CommentPresenter{}
	json := pr.Render(data)
	return presenter.RenderJSON(c, json, result)
}

func (ctl *CommentController) Delete(c echo.Context) error {
	panic("implement me")
}
