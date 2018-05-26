package controller

import (
	"fmt"
	"net/http"

)

var count = 0

// CommentController handles request which is related to comment domain.
type CommentController struct{}

func (c *CommentController) List(w http.ResponseWriter, r *http.Request) {
	count++
	//ctx := appengine.NewContext(r)
	//datastore.NewKey()
	fmt.Fprint(w, fmt.Sprintf("%d", count))
}

func (c *CommentController) Add(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func NewCommentController() *CommentController {
	return &CommentController{}
}
