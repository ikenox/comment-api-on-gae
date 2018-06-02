package controller

import (
	"fmt"
	"net/http"

	"comment-api-on-gae/repository"
	"comment-api-on-gae/usecase"
	"encoding/json"
	"strconv"

	"google.golang.org/appengine"
)

var count = 0

// CommentController handles request which is related to comment domain.
type CommentController struct{}

func (c *CommentController) List(w http.ResponseWriter, r *http.Request) {
	count++
	fmt.Fprint(w, fmt.Sprintf("%d", count))
}

func (c *CommentController) Add(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	length, _ := strconv.Atoi(r.Header.Get("Content-Length"))
	body := make([]byte, length)
	fmt.Fprint(w, fmt.Sprintf("%s", body))
	var params struct {
		pageUrl string
		text    string
	}
	if err := json.Unmarshal(body, &params); err != nil {
		panic(err)
	}


	commentUseCase := usecase.NewCommentUseCase(
		repository.NewPostRepository(ctx),
	)
	commentUseCase.Post(params.pageUrl, params.text)

	fmt.Fprint(w, fmt.Sprintf("%s, %s", params.pageUrl, params.text))
}

func NewCommentController() *CommentController {
	return &CommentController{}
}
