package controller

import (
	"fmt"
	"net/http"

	"comment-api-on-gae/repository"
	"comment-api-on-gae/usecase"
	"encoding/json"
	"google.golang.org/appengine"
)

var count = 0

// CommentController handles request which is related to comment domain.
type CommentController struct{}

func NewCommentController() *CommentController {
	return &CommentController{}
}

func (c *CommentController) List(w http.ResponseWriter, r *http.Request) {
	count++
	fmt.Fprint(w, fmt.Sprintf("%d", count))
}

func (c *CommentController) Add(w http.ResponseWriter, r *http.Request) {
	var params struct {
		pageUrl string
		text    string
		name    string
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		panic(err)
	}

	ctx := appengine.NewContext(r)
	usecase.NewCommentUseCase(
		repository.NewCommentRepository(ctx),
		repository.NewPageRepository(ctx),
		repository.NewCommenterRepository(ctx),
	).PostNewComment(params.pageUrl, params.name, params.text)

	fmt.Fprint(w, fmt.Sprintf("%s, %s", params.pageUrl, params.text))
}
