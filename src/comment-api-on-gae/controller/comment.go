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

type PageController struct{}

func NewPageController() *PageController {
	return &PageController{}
}

func (c *PageController) List(w http.ResponseWriter, r *http.Request) {
	var params struct {
		PageUrl string
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
	).GetComments(params.PageUrl)

}

func (c *PageController) Add(w http.ResponseWriter, r *http.Request) {
	var params struct {
		PageUrl string
		Text    string
		Name    string
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
	).PostComment(params.PageUrl, params.Name, params.Text)

	fmt.Fprint(w, fmt.Sprintf("%s, %s", params.PageUrl, params.Text))
}
