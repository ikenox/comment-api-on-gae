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

// adapter層なので一意な変換以外のロジックは基本入らなくなる？
func (ctl *CommentController) List(c echo.Context) error {
	pageId := c.QueryParam("pageId")

	ctx := c.StdContext()
	u := usecase.NewCommentUseCase(
		repository.NewCommentRepository(ctx),
		repository.NewCommenterRepository(ctx),
		repository.NewPageRepository(ctx),
		repository.NewPublisher(ctx),
		repository.NewLoggingRepository(ctx),
	)
	comments, res := u.GetComments(pageId)

	pr := &presenter.CommentPresenter{}
	json := make([]interface{}, len(comments))
    for i, comment := range comments {
        json[i] = pr.Render(comment)
    }

	return presenter.RenderJSON(c, json, res)
}

func (ctl *CommentController) PostComment(c echo.Context) error {
	var p = &struct {
		PageId  string `json:"pageId"`
		Name    string `json:"name"`
		Text    string `json:"text"`
	}{}
	if err := c.Bind(p); err != nil {
		// 変換エラーはinterface adapter層における異常系
		// リクエストの形式がプロトコルに沿ってないため
		return err
	}

	IDToken := c.Request().Header().Get("IdToken")

	ctx := c.StdContext()
	u := usecase.NewCommentUseCase(
		repository.NewCommentRepository(ctx),
		repository.NewCommenterRepository(ctx),
		repository.NewPageRepository(ctx),
		repository.NewPublisher(ctx),
		repository.NewLoggingRepository(ctx),
	)
	data, result := u.PostComment(IDToken, p.Name, p.PageId, p.Text)
	pr := &presenter.CommentPresenter{}
	json := pr.Render(data)
	// usecaseでのエラーはinterface adapterにとっては正常系であり、エラーかどうかは特に意識せず一意に変換できるのが良いと考えた
	return presenter.RenderJSON(c, json, result)
}

func (ctl *CommentController) Delete(c echo.Context) error {
	commentID := c.Param("id")
	IDToken := c.Request().Header().Get("IdToken")

	ctx := c.StdContext()
	u := usecase.NewCommentUseCase(
		repository.NewCommentRepository(ctx),
		repository.NewCommenterRepository(ctx),
		repository.NewPageRepository(ctx),
		repository.NewPublisher(ctx),
		repository.NewLoggingRepository(ctx),
	)
	result := u.DeleteComment(IDToken, commentID)
	return presenter.RenderJSON(c, nil, result)
}
