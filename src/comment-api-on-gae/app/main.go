package app

import (
	"comment-api-on-gae/controller"
	"net/http"
)

func init() {
	commentController := controller.NewCommentController()
	http.HandleFunc("/comment/list", commentController.List)
	http.HandleFunc("/comment/add", commentController.Add)
	http.ListenAndServe(":8080", nil)
}
