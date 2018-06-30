package controller

import (
	"comment-api-on-gae/repository"
	"comment-api-on-gae/usecase"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type CommentController struct{}

func NewCommentController() *CommentController {
	return &CommentController{}
}

type commentsPresenter struct {
	CommentId   int64     `json:"commentId"`
	PageId      string    `json:"pageId"`
	Text        string    `json:"text"`
	CommenterId int64     `json:"commenterId"`
	CommentedAt time.Time `json:"commentedAt"`
}

type applicationErrorPresenter struct {
	Message string `json:"message"`
}

func (ctl *CommentController) List(c echo.Context) error {
	var params struct {
		PageId string
	}
	if err := c.Bind(&params); err != nil {
		return err
	}

	ctx := c.StdContext()
	comments, err := usecase.NewCommentUseCase(
		repository.NewCommentRepository(ctx),
		repository.NewPageRepository(ctx),
		repository.NewCommenterRepository(ctx),
	).GetComments(params.PageId)
	if err != nil {
		return renderErrorJSON(c, err)
	}

	// TODO: 別集約を1つにまとめて返すための正しい方法
	// TODO: Json用structの置き場所やネーミング
	commentsJson := []*commentsPresenter{}
	for _, cm := range comments {
		commentsJson = append(commentsJson, &commentsPresenter{
			CommentId:   int64(cm.CommentId()),
			PageId:      string(cm.PageId()),
			Text:        cm.Text(),
			CommenterId: int64(cm.CommenterId()),
			CommentedAt: cm.CommentedAt(),
		})
	}

	// TODO: レスポンスのデータ構造
	return c.JSON(http.StatusOK, commentsJson)
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
	err := usecase.NewCommentUseCase(
		repository.NewCommentRepository(ctx),
		repository.NewPageRepository(ctx),
		repository.NewCommenterRepository(ctx),
	).PostComment(params.PageId, params.Name, params.Text)
	if err != nil {
		return renderErrorJSON(c, err)
	}

	return c.JSON(http.StatusCreated, struct{}{})
}

func renderErrorJSON(c echo.Context, err *usecase.Error) error {
	var status int
	switch err.Code() {
	case usecase.EINVALID:
		status = http.StatusBadRequest
	case usecase.EINTERNAL:
		status = http.StatusInternalServerError
	case usecase.ENOTFOUND:
		status = http.StatusNotFound
	default:
		panic(fmt.Sprintf("Unknown Error Code '%s'", err.Code()))
	}

	return c.JSON(
		status,
		&applicationErrorPresenter{
			Message: err.Message(),
		},
	)
}
