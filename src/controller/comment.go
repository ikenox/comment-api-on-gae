package controller

import (
	"domain"
	"fmt"
	"net/http"
)

// CommentController handles request which is related to comment domain.
type CommentController struct{}

func (c *CommentController) List(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, fmt.Sprintf("%s", &domain.Post{}))
}

func (c *CommentController) Add(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func NewCommentController() *CommentController {
	return &CommentController{}
}
