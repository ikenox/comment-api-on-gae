package main

import (
	"controller"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	r := mux.NewRouter()
	commentController := controller.NewCommentController()
	r.HandleFunc("/comment/list", commentController.List)
	r.HandleFunc("/comment/add", commentController.Add)

	http.Handle("/", r)
}
