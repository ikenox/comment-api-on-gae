package controller

import (
	"comment-api-on-gae/repository"
	"comment-api-on-gae/usecase"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type PageController struct{}

func NewPageController() *PageController {
	return &PageController{}
}

type commentsPresenter struct {
	CommentId   int64     `json:"commentId"`
	PageId      int64     `json:"pageId"`
	Text        string    `json:"text"`
	CommenterId int64     `json:"commenterId"`
	CommentedAt time.Time `json:"commentedAt"`
}

func (ctl *PageController) List(c echo.Context) error {
	var params struct {
		Url string
	}
	if err := c.Bind(&params); err != nil {
		return err
	}

	ctx := c.StdContext()
	comments := usecase.NewCommentUseCase(
		repository.NewCommentRepository(ctx),
		repository.NewPageRepository(ctx),
		repository.NewCommenterRepository(ctx),
	).GetComments(params.Url)

	// TODO: 別集約を1つにまとめて返すための正しい方法
	// TODO: Json用structの置き場所やネーミング
	commentsJson := []*commentsPresenter{}
	for _, cm := range comments {
		commentsJson = append(commentsJson, &commentsPresenter{
			CommentId:   int64(cm.CommentId()),
			PageId:      int64(cm.PageId()),
			Text:        cm.Text(),
			CommenterId: int64(cm.CommenterId()),
			CommentedAt: cm.CommentedAt(),
		})
	}

	// TODO: レスポンスのデータ構造
	return c.JSON(http.StatusOK, commentsJson)
}

func (ctl *PageController) PostComment(c echo.Context) error {
	var params struct {
		Url  string
		Name string
		Text string
	}
	if err := c.Bind(&params); err != nil {
		return err
	}

	ctx := c.StdContext()
	usecase.NewCommentUseCase(
		repository.NewCommentRepository(ctx),
		repository.NewPageRepository(ctx),
		repository.NewCommenterRepository(ctx),
	).PostComment(params.Url, params.Name, params.Text)

	return c.JSON(http.StatusCreated, struct{}{})
}
