package main

import (
	"comment"
	"net/http"
)

func init() {
	commentHandler := new(comment.Handler)
	http.HandleFunc("/", commentHandler.List)
}
